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
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/query"
)

type CreateIndex struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	bucket                catalog.Bucket
	name                  string
	index_type            string
	on                    ast.ExpressionList
	downstreamStopChannel misc.StopChannel
}

func NewCreateIndex(bucket catalog.Bucket, name string, index_type string, on ast.ExpressionList) *CreateIndex {
	return &CreateIndex{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		bucket:         bucket,
		name:           name,
		index_type:     index_type,
		on:             on,
	}
}

func (this *CreateIndex) SetSource(source Operator) {}

func (this *CreateIndex) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *CreateIndex) Run(stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	// this MUST be here so that it runs before the channels are closed
	defer this.RecoverPanic()

	this.downstreamStopChannel = stopChannel
	clog.To(CHANNEL, "create_index operator starting")
	index, err := this.bucket.CreateIndex(this.name, this.on, this.index_type)
	if err != nil {
		this.SendError(err)
	} else {
		this.SendItem(dparval.NewValue(map[string]interface{}{
			"id":   index.Id(),
			"name": index.Name(),
		}))
	}
	clog.To(CHANNEL, "create_index operator finished")
}

func (this *CreateIndex) processItem(item *dparval.Value) bool {
	return true
}

func (this *CreateIndex) afterItems() {}

func (this *CreateIndex) SendItem(item *dparval.Value) bool {
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

func (this *CreateIndex) SendError(err query.Error) bool {
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

func (this *CreateIndex) RecoverPanic() {
	r := recover()
	if r != nil {
		clog.Error(fmt.Errorf("Query Execution Panic: %v\n%s", r, debug.Stack()))
		this.SendError(query.NewError(nil, "Panic In Exeuction Pipeline"))
	}
}
