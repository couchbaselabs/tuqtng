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

type Limit struct {
	Base  *BaseOperator
	Limit int
	count int
}

func NewLimit(limit int) *Limit {
	return &Limit{
		Base:  NewBaseOperator(),
		Limit: limit,
	}
}

func (this *Limit) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *Limit) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *Limit) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "limit operator starting")
	this.Base.RunOperator(this, stopChannel)
	clog.To(CHANNEL, "limit operator finished")
}

func (this *Limit) processItem(item *dparval.Value) bool {
	this.count++
	if this.count > this.Limit {
		return false
	}
	this.Base.SendItem(item)
	return true
}

func (this *Limit) afterItems() {}
