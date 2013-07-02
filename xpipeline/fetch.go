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

type Fetch struct {
	Source      Operator
	itemChannel query.ItemChannel
	bucket      catalog.Bucket
}

func NewFetch(bucket catalog.Bucket) *Fetch {
	return &Fetch{
		itemChannel: make(query.ItemChannel),
		bucket:      bucket,
	}
}

func (this *Fetch) SetSource(source Operator) {
	this.Source = source
}

func (this *Fetch) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Fetch) Run() {
	defer close(this.itemChannel)

	go this.Source.Run()

	for item := range this.Source.GetItemChannel() {
		meta := item.GetMeta()
		id, ok := meta["id"]

		if ok {
			fetchedItem, err := this.bucket.Fetch(id.(string)) //FIXME assert ids always strings
			if err != nil {
				log.Printf("error fetching from bucket %v", err)
				// FIXME proper error handling
			}
			this.itemChannel <- fetchedItem
		} else {
			log.Printf("fetch got an item without an id??")
			// FIXME proper error handling
		}
	}
}
