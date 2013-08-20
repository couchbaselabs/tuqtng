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
	"github.com/couchbaselabs/tuqtng/misc"
)

type ExpressionEvaluatorSource struct {
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
}

func NewExpressionEvaluatorSource() *ExpressionEvaluatorSource {
	return &ExpressionEvaluatorSource{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
	}
}

func (this *ExpressionEvaluatorSource) SetSource(source Operator) {}

func (this *ExpressionEvaluatorSource) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *ExpressionEvaluatorSource) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "expression evaluator source operator starting")
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	item := dparval.NewValue(map[string]interface{}{})
	this.itemChannel <- item
	clog.To(CHANNEL, "expression evaluator source operator finished")
}

func (this *ExpressionEvaluatorSource) processItem(item *dparval.Value) bool {
	return true
}

func (this *ExpressionEvaluatorSource) afterItems() {}
