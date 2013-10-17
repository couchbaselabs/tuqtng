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
	"runtime/debug"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

type DropIndex struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	index                 catalog.Index
	downstreamStopChannel misc.StopChannel
	query                 network.Query
}

func NewDropIndex(index catalog.Index) *DropIndex {
	return &DropIndex{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		index:          index,
	}
}

func (this *DropIndex) SetSource(source Operator) {}

func (this *DropIndex) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *DropIndex) Run(stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	// this MUST be here so that it runs before the channels are closed
	defer this.RecoverPanic()

	this.downstreamStopChannel = stopChannel
	clog.To(CHANNEL, "drop_index operator starting")
	err := this.index.Drop()
	if err != nil {
		this.SendError(err)
	} else {
		item := dparval.NewValue(map[string]interface{}{})
		item.SetAttachment("projection", map[string]interface{}{
			"dropped": true,
		})
		this.SendItem(item)
	}
	clog.To(CHANNEL, "drop_index operator finished")
}

func (this *DropIndex) processItem(item *dparval.Value) bool {
	return true
}

func (this *DropIndex) afterItems() {}

func (this *DropIndex) SendItem(item *dparval.Value) bool {
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

func (this *DropIndex) SendError(err query.Error) bool {
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

func (this *DropIndex) RecoverPanic() {
	r := recover()
	if r != nil {
		clog.Error(fmt.Errorf("Query Execution Panic: %v\n%s", r, debug.Stack()))
		this.SendError(query.NewError(nil, "Panic In Exeuction Pipeline"))
	}
}

func (this *DropIndex) SetQuery(q network.Query) {
	this.query = q
}
