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
	"github.com/couchbaselabs/tuqtng/query"
)

type Limit struct {
	Source      Operator
	Limit       int
	itemChannel query.ItemChannel
}

func NewLimit(limit int) *Limit {
	return &Limit{
		Limit:       limit,
		itemChannel: make(query.ItemChannel),
	}
}

func (this *Limit) SetSource(source Operator) {
	this.Source = source
}

func (this *Limit) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Limit) Run() {
	defer close(this.itemChannel)

	count := 0

	// start the source
	go this.Source.Run()
	for item := range this.Source.GetItemChannel() {
		this.itemChannel <- item
		count++
		if count >= this.Limit {
			break
		}
	}
}
