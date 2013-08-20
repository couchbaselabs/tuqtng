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
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/query"
)

type BaseOperator struct {
	Source                Operator
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	upstreamStopChannel   misc.StopChannel
	downstreamStopChannel misc.StopChannel
}

func NewBaseOperator() *BaseOperator {
	return &BaseOperator{
		itemChannel:         make(dparval.ValueChannel),
		supportChannel:      make(PipelineSupportChannel),
		upstreamStopChannel: make(misc.StopChannel),
	}
}

func (this *BaseOperator) SetSource(source Operator) {
	this.Source = source
}

func (this *BaseOperator) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *BaseOperator) SendItem(item *dparval.Value) bool {
	ok := true
	for ok {
		select {
		case this.itemChannel <- item:
			return true
		case _, ok = <-this.downstreamStopChannel:
			// someone closed the stop channel
		}
	}
	return ok
}

func (this *BaseOperator) SendError(err query.Error) bool {
	ok := true
	for ok {
		select {
		case this.supportChannel <- err:
			if err.IsFatal() {
				return false
			}
			return true
		case _, ok = <-this.downstreamStopChannel:
			// someone closed the stop channel
		}
	}
	return false
}

func (this *BaseOperator) SendOther(obj interface{}) bool {
	ok := true
	for ok {
		select {
		case this.supportChannel <- obj:
			return true
		case _, ok = <-this.downstreamStopChannel:
			// someone closed the stop channel
		}
	}
	return false
}

// func (this *BaseOperator) StopUpstream() {
// 	close(this.stopChannel)
// }

func (this *BaseOperator) RunOperator(oper Operator, stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	defer close(this.upstreamStopChannel)

	this.downstreamStopChannel = stopChannel

	go this.Source.Run(this.upstreamStopChannel)

	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	ok := true
	for ok {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				ok = oper.processItem(item)
			}
		case obj, ok = <-supportChannel:
			if ok {
				switch obj := obj.(type) {
				case query.Error:
					ok = this.SendError(obj)
				default:
					ok = this.SendOther(obj)
				}
			}
		case _, ok = <-stopChannel:
			// downstream has asked us to stop
		}
	}

	oper.afterItems()
}
