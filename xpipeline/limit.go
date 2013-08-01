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
	"github.com/mschoch/dparval"
)

type Limit struct {
	Source         Operator
	Limit          int
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	count          int
}

func NewLimit(limit int) *Limit {
	return &Limit{
		Limit:          limit,
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
	}
}

func (this *Limit) SetSource(source Operator) {
	this.Source = source
}

func (this *Limit) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Limit) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	this.count = 0

	go this.Source.Run()

	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	ok := true
	for ok && this.count < this.Limit {
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
}

func (this *Limit) processItem(item *dparval.Value) {
	this.itemChannel <- item
	this.count++
}
