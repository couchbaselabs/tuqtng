//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

// this operator performs a document join. It needs a bucket from where the keys
// will be fetched along with a Key expression that will extract the list of keys
// to be fetched from the downstream data channel

import (
	"fmt"
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

type KeyJoin struct {
	Base        *BaseOperator
	bucket      catalog.Bucket
	Projection  ast.Expression
	Keys        ast.KeyExpression
	joinType    int
	As          string
	batch       []string
	rowsFetched int
}

func NewKeyJoin(bucket catalog.Bucket, projection ast.Expression, keys ast.KeyExpression, as string) *KeyJoin {
	return &KeyJoin{
		Base:       NewBaseOperator(),
		bucket:     bucket,
		Keys:       keys,
		Projection: projection,
		As:         as,
	}
}

func (this *KeyJoin) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *KeyJoin) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *KeyJoin) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "key join operator starting")
	if this.Base.Source != nil {
		this.Base.RunOperator(this, stopChannel)
	} else {
		this.Base.SendError(query.NewError(fmt.Errorf("missing source operator"), ""))
	}

	clog.To(CHANNEL, "key operator finished, fetched %d", this.rowsFetched)
}

// evaluate the key expression from the item and fetch the keys from the bucket

func (this *KeyJoin) processItem(item *dparval.Value) bool {
	// add this item to the batch
	val, err := this.Base.Evaluate(this.Keys.Expr, item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return true
		default:
			return this.Base.SendError(query.NewError(err, "Internal error in KeyJoin"))
		}
	}

	if val.Type() == dparval.STRING {
		if this.Keys.Type == "KEYS" {
			this.Base.SendError(query.NewError(fmt.Errorf("KEYS expression should evaluate to an array"), ""))
			return false
		}
		fetch_id := val.Value().(string)
		keyItem, err := this.bucket.Fetch(fetch_id)
		if err != nil {
			this.Base.SendError(query.NewError(err, "Unable to fetch key"))
		}

		return this.joinItems(item, keyItem)
	}
	if val.Type() == dparval.ARRAY {
		ok := true
		index := 0

		if this.Keys.Type == "KEY" {
			this.Base.SendError(query.NewError(fmt.Errorf("KEY used with an array argument"), ""))
			return false
		}
		for ok {
			id, err := val.Index(index)
			index = index + 1
			if err != nil {
				return true
			}
			fetch_id := (*id).Value().(string)
			keyItem, err := this.bucket.Fetch(fetch_id)
			if err != nil {
				this.Base.SendError(query.NewError(err, "Unable to fetch key"))
			}
			ok = this.joinItems(item, keyItem)
			this.rowsFetched += 1
		}
	}

	return true
}

func (this *KeyJoin) joinItems(item *dparval.Value, keyItem *dparval.Value) bool {

	newItem := item.Duplicate()
	/* join the item and ship it */
	if this.Projection != nil {
		keyProj, Error := this.Base.Evaluate(this.Projection, keyItem)
		if Error != nil {
			switch err := Error.(type) {
			case *dparval.Undefined:
				return true
			default:
				return this.Base.SendError(query.NewError(err, "Internal error in KeyJoin"))
			}

		}
		newItem.SetPath(this.As, keyProj)
	} else {
		newItem.SetPath(this.As, keyItem)
	}
	this.Base.SendItem(newItem)
	return true
}

func (this *KeyJoin) afterItems() {}

/*
func (this *KeyJoin) flushBatch() bool {

	defer func() {
		// no matter what hapens in this function
		// clear out the batch and start a new one
		this.batch = make(dparval.ValueCollection, 0, FETCH_BATCH_SIZE)
	}()

	// gather the ids
	ids := make([]string, 0, FETCH_BATCH_SIZE)
	for _, v := range this.batch {
		meta := v.GetAttachment("meta")
		if meta != nil {
			id, ok := meta.(map[string]interface{})["id"]
			if !ok {
				return this.Base.SendError(query.NewError(nil, "asked to fetch an item without a key"))
			} else {
				ids = append(ids, id.(string)) //FIXME assert ids always strings
			}
		} else {
			return this.Base.SendError(query.NewError(nil, "asked to fetch an item without a meta"))
		}
	}

	// now do a bulk fetch
	bulkResponse, err := this.bucket.BulkFetch(ids)
	if err != nil {
		return this.Base.SendError(query.NewError(err, "error getting bulk response"))
	}

	// now we need to emit the bulk fetched items in the correct order (from the id list)
	for _, v := range ids {
		item, ok := bulkResponse[v]
		if ok {
			if this.projection != nil {
				projectedVal, err := this.Base.projectedValueOfResultExpression(item, ast.NewResultExpression(this.projection))
				if err != nil {
					switch err := err.(type) {
					case *dparval.Undefined:
						// undefined contributes nothing to the result map
						continue
					default:
						return this.Base.SendError(query.NewError(err, "unexpected error projecting fetch expression"))
					}
				} else {
					if this.as != "" {
						this.Base.SendItem(dparval.NewValue(map[string]interface{}{this.as: projectedVal}))
					} else {
						this.Base.SendItem(projectedVal)
					}
				}
			} else {
				if this.as != "" {
					this.Base.SendItem(dparval.NewValue(map[string]interface{}{this.as: item}))
				} else {
					this.Base.SendItem(item)
				}
			}
			this.rowsFetched += 1
		}
	}
	return true
}
*/

func (this *KeyJoin) SetQuery(q network.Query) {
	this.Base.SetQuery(q)
}
