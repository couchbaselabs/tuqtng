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
	"github.com/couchbaselabs/tuqtng/misc"
)

// this is a terrible implementation to remove duplicates
// it stores the entire result set in memory
// and compares each document against the remaining documents
// FIXME improve implementation

type EliminateDuplicates struct {
	Base   *BaseOperator
	buffer dparval.ValueCollection
}

func NewEliminateDuplicates() *EliminateDuplicates {
	return &EliminateDuplicates{
		Base:   NewBaseOperator(),
		buffer: make(dparval.ValueCollection, 0),
	}
}

func (this *EliminateDuplicates) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *EliminateDuplicates) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *EliminateDuplicates) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "eliminate duplicates operator starting")
	this.Base.RunOperator(this, stopChannel)
	clog.To(CHANNEL, "eliminate duplicates operator finished")
}

func (this *EliminateDuplicates) processItem(item *dparval.Value) bool {
	this.buffer = append(this.buffer, item)
	return true
}

func (this *EliminateDuplicates) afterItems() {
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
			this.Base.SendItem(item)
		}
	}
}
