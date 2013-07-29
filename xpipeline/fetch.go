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
)

const FETCH_BATCH_SIZE = 1000

type Fetch struct {
	Source      Operator
	itemChannel query.ItemChannel
	errChannel  query.ErrorChannel
	warnChannel query.ErrorChannel
	bucket      catalog.Bucket
	batch       []query.Item
	ok          bool
}

func NewFetch(bucket catalog.Bucket) *Fetch {
	return &Fetch{
		itemChannel: make(query.ItemChannel),
		errChannel:  make(query.ErrorChannel),
		warnChannel: make(query.ErrorChannel),
		bucket:      bucket,
		batch:       make([]query.Item, 0, FETCH_BATCH_SIZE),
	}
}

func (this *Fetch) SetSource(source Operator) {
	this.Source = source
}

func (this *Fetch) GetChannels() (query.ItemChannel, query.ErrorChannel, query.ErrorChannel) {
	return this.itemChannel, this.warnChannel, this.errChannel
}

func (this *Fetch) Run() {
	defer this.bucket.Release()
	defer close(this.itemChannel)
	defer close(this.errChannel)
	defer close(this.warnChannel)

	go this.Source.Run()

	var item query.Item
	var warn query.Error
	var err query.Error
	sourceItemChannel, sourceWarnChannel, sourceErrorChannel := this.Source.GetChannels()
	this.ok = true
	for this.ok {
		select {
		case item, this.ok = <-sourceItemChannel:
			if this.ok {
				this.processItem(item)
			}
		case warn, this.ok = <-sourceWarnChannel:
			// propogate the warning
			if warn != nil {
				this.warnChannel <- warn
			}
		case err, this.ok = <-sourceErrorChannel:
			// propogate the error and return
			if err != nil {
				this.errChannel <- err
				return
			}
		}
	}

	// source may have closed with a parital batch
	this.flushBatch()
}

func (this *Fetch) processItem(item query.Item) {
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
		this.batch = make([]query.Item, 0, FETCH_BATCH_SIZE)
	}()

	// gather the ids
	ids := make([]string, 0, FETCH_BATCH_SIZE)
	for _, v := range this.batch {
		meta := v.GetMeta()
		id, ok := meta["id"]
		if !ok {
			this.errChannel <- query.NewError(nil, "asked to fetch an item without a key")
			this.ok = false
		} else {
			ids = append(ids, id.(string)) //FIXME assert ids always strings
		}
	}

	// now do a bulk fetch

	bulkResponse, err := this.bucket.BulkFetch(ids)
	if err != nil {
		this.errChannel <- query.NewError(err, "error getting bulk response")
		this.ok = false
		return
	}

	// now we need to emit the bulk fetched items in the correct order (from the id list)
	for _, v := range ids {
		item, ok := bulkResponse[v]
		if !ok {
			this.errChannel <- query.NewError(nil, fmt.Sprintf("missing value bulk response for key %s", v))
			this.ok = false
		} else {
			this.itemChannel <- item
		}
	}
}
