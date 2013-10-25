//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

const FETCH_BATCH_SIZE = 1000

type Fetch struct {
	Base        *BaseOperator
	bucket      catalog.Bucket
	batch       dparval.ValueCollection
	projection  ast.Expression
	as          string
	ids         []string
	rowsFetched int
}

func NewFetch(bucket catalog.Bucket, projection ast.Expression, as string) *Fetch {
	return &Fetch{
		Base:       NewBaseOperator(),
		bucket:     bucket,
		batch:      make(dparval.ValueCollection, 0, FETCH_BATCH_SIZE),
		projection: projection,
		as:         as,
	}
}

func (this *Fetch) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *Fetch) SetIds(ids []string) {
	this.ids = ids
}

func (this *Fetch) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *Fetch) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "fetch operator starting")
	if this.Base.Source != nil {
		this.Base.RunOperator(this, stopChannel)
	} else {
		defer close(this.Base.itemChannel)
		defer close(this.Base.supportChannel)
		defer close(this.Base.upstreamStopChannel)
		for _, id := range this.ids {
			doc := dparval.NewValue(map[string]interface{}{})
			doc.SetAttachment("meta", map[string]interface{}{"id": id})
			this.processItem(doc)
		}
		this.afterItems()
	}
	clog.To(CHANNEL, "fetch operator finished, fetched %d", this.rowsFetched)
}

func (this *Fetch) processItem(item *dparval.Value) bool {
	// add this item to the batch
	this.batch = append(this.batch, item)

	// if the batch is full, do a fetch
	if len(this.batch) >= FETCH_BATCH_SIZE {
		return this.flushBatch()
	}
	return true
}

func (this *Fetch) afterItems() {
	// source may have closed with a parital batch
	this.flushBatch()
}

func (this *Fetch) flushBatch() bool {

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

func (this *Fetch) SetQuery(q network.Query) {
	this.Base.SetQuery(q)
}
