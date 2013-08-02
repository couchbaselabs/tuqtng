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
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/dparval"
)

// this is a terrible implementation to remove duplicates
// it stores the entire result set in memory
// and compares each document against the remaining documents
// FIXME improve implementation

type EliminateDuplicates struct {
	Source         Operator
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	buffer         dparval.ValueCollection
}

func NewEliminateDuplicates() *EliminateDuplicates {
	return &EliminateDuplicates{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		buffer:         make(dparval.ValueCollection, 0),
	}
}

func (this *EliminateDuplicates) SetSource(source Operator) {
	this.Source = source
}

func (this *EliminateDuplicates) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *EliminateDuplicates) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	go this.Source.Run()

	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	ok := true
	for ok {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				this.processItem(item)
			}
		case obj, ok = <-supportChannel:
			if ok {
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

	// write the output
	for pos, item := range this.buffer {
		// we will nil out duplicates and then skip over those entries in the buffer
		if item != nil {
			if pos < len(this.buffer) {
				// look to see if the exact same item appears later in the buffer
				for nextpos, nextitem := range this.buffer[pos+1:] {
					itemVal := item.Value()
					nextItemVal := nextitem.Value()
					comp := ast.CollateJSON(itemVal, nextItemVal)
					if comp == 0 {
						this.buffer[nextpos+1] = nil
					}
				}
			}
			this.itemChannel <- item
		}
	}
}

func (this *EliminateDuplicates) processItem(item *dparval.Value) {
	this.buffer = append(this.buffer, item)
}
