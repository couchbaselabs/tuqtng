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
	"fmt"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/dparval"
)

const FETCH_BATCH_SIZE = 1000

type Fetch struct {
	Source         Operator
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	bucket         catalog.Bucket
	batch          dparval.ValueCollection
	ok             bool
}

func NewFetch(bucket catalog.Bucket) *Fetch {
	return &Fetch{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		bucket:         bucket,
		batch:          make(dparval.ValueCollection, 0, FETCH_BATCH_SIZE),
	}
}

func (this *Fetch) SetSource(source Operator) {
	this.Source = source
}

func (this *Fetch) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Fetch) Run() {
	defer this.bucket.Release()
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	go this.Source.Run()

	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	this.ok = true
	for this.ok {
		select {
		case item, this.ok = <-sourceItemChannel:
			if this.ok {
				this.processItem(item)
			}
		case obj, this.ok = <-supportChannel:
			if this.ok {
				switch obj := obj.(type) {
				case query.Error:
					this.supportChannel <- obj
					if obj.IsFatal() {
						return
					}
				default:
					this.supportChannel <- obj
				}
			}
		}
	}

	// source may have closed with a parital batch
	this.flushBatch()
}

func (this *Fetch) processItem(item *dparval.Value) {
	// add this item to the batch
	this.batch = append(this.batch, item)

	// if the batch is full, do a fetch
	if len(this.batch) >= FETCH_BATCH_SIZE {
		this.flushBatch()
	}
}

func (this *Fetch) flushBatch() {

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
				this.supportChannel <- query.NewError(nil, "asked to fetch an item without a key")
				this.ok = false
			} else {
				ids = append(ids, id.(string)) //FIXME assert ids always strings
			}
		} else {
			this.supportChannel <- query.NewError(nil, "asked to fetch an item without a meta")
			this.ok = false
		}
	}

	// now do a bulk fetch
	bulkResponse, err := this.bucket.BulkFetch(ids)
	if err != nil {
		this.supportChannel <- query.NewError(err, "error getting bulk response")
		this.ok = false
		return
	}

	// now we need to emit the bulk fetched items in the correct order (from the id list)
	for _, v := range ids {
		item, ok := bulkResponse[v]
		if !ok {
			this.supportChannel <- query.NewError(nil, fmt.Sprintf("missing value bulk response for key `%s`", v))
			this.ok = false
		} else {
			this.itemChannel <- item
		}
	}
}
