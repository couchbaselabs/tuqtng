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
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
)

type Explain struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	downstreamStopChannel misc.StopChannel
	Plan                  plan.PlanElement
}

func NewExplain(plan plan.PlanElement) *Explain {
	return &Explain{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		Plan:           plan,
	}
}

func (this *Explain) SetSource(source Operator) {}

func (this *Explain) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Explain) Run(stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	// this MUST be here so that it runs before the channels are closed
	defer this.RecoverPanic()

	this.downstreamStopChannel = stopChannel
	clog.To(CHANNEL, "explain operator starting")
	item := dparval.NewValue(map[string]interface{}{})

	planBytes, err := json.Marshal(this.Plan)
	if err != nil {
		this.SendError(query.NewError(err, "error serializing plan to JSON"))
	} else {
		projection := dparval.NewValueFromBytes(planBytes)
		item.SetAttachment("projection", projection)
		this.SendItem(item)
	}

	clog.To(CHANNEL, "explain operator finished")
}

func (this *Explain) processItem(item *dparval.Value) bool {
	return true
}

func (this *Explain) afterItems() {}

func (this *Explain) SendItem(item *dparval.Value) bool {
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

func (this *Explain) SendError(err query.Error) bool {
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

func (this *Explain) RecoverPanic() {
	r := recover()
	if r != nil {
		clog.Error(fmt.Errorf("Query Execution Panic: %v\n%s", r, debug.Stack()))
		this.SendError(query.NewError(nil, "Panic In Exeuction Pipeline"))
	}
}
