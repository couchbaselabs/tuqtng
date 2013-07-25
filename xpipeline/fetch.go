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
	"log"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

const FETCH_BATCH_SIZE = 1000

type Fetch struct {
	Source      Operator
	itemChannel query.ItemChannel
	bucket      catalog.Bucket
	batch       []query.Item
}

func NewFetch(bucket catalog.Bucket) *Fetch {
	return &Fetch{
		itemChannel: make(query.ItemChannel),
		bucket:      bucket,
		batch:       make([]query.Item, 0, FETCH_BATCH_SIZE),
	}
}

func (this *Fetch) SetSource(source Operator) {
	this.Source = source
}

func (this *Fetch) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Fetch) Run() {
	defer this.bucket.Release()
	defer close(this.itemChannel)

	go this.Source.Run()

	for item := range this.Source.GetItemChannel() {

		// add this item to the batch
		this.batch = append(this.batch, item)

		// if the batch is full, do a fetch
		if len(this.batch) >= FETCH_BATCH_SIZE {
			this.flushBatch()
		}

	}

	// source may have closed with a parital batch
	this.flushBatch()
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
			log.Printf("asked to fetch an item without an id")
		} else {
			ids = append(ids, id.(string)) //FIXME assert ids always strings
		}
	}

	// now do a bulk fetch

	bulkResponse, err := this.bucket.BulkFetch(ids)
	if err != nil {
		// FIXME proper error handling
		log.Printf("error getting bulk response %v", err)
		return
	}

	// now we need to emit the bulk fetched items in the correct order (from the id list)
	for _, v := range ids {
		item, ok := bulkResponse[v]
		if !ok {
			log.Printf("missing item in the bulk response %v", v)
		} else {
			this.itemChannel <- item
		}
	}
}
