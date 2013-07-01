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

type Offset struct {
	Source      Operator
	Offset      int
	itemChannel query.ItemChannel
}

func NewOffset(offset int) *Offset {
	return &Offset{
		Offset:      offset,
		itemChannel: make(query.ItemChannel),
	}
}

func (this *Offset) SetSource(source Operator) {
	this.Source = source
}

func (this *Offset) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Offset) Run() {
	defer close(this.itemChannel)

	count := 0

	// start the source
	go this.Source.Run()
	for item := range this.Source.GetItemChannel() {
		count++
		if count <= this.Offset {
			continue
		}
		this.itemChannel <- item
	}
}
