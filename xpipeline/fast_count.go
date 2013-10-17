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
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
)

type FastCount struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	bucket                catalog.Bucket
	index                 catalog.CountIndex
	downstreamStopChannel misc.StopChannel
	ranges                plan.ScanRanges
	expr                  ast.Expression
	query                 network.Query
}

func NewFastCount(bucket catalog.Bucket, index catalog.CountIndex, expr ast.Expression, ranges plan.ScanRanges) *FastCount {
	return &FastCount{
		expr:           expr,
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		bucket:         bucket,
		index:          index,
		ranges:         ranges,
	}
}

func (this *FastCount) SetSource(source Operator) {}

func (this *FastCount) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *FastCount) Run(stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	// this MUST be here so that it runs before the channels are closed
	defer this.RecoverPanic()

	clog.To(CHANNEL, "fastcount operator starting")

	if this.ranges == nil {
		if this.index == nil {
			count, err := this.bucket.Count()
			if err != nil {
				this.SendError(query.NewError(err, "Error counting values in bucket"))
			} else {
				groupDoc := dparval.NewValue(map[string]interface{}{})
				aggregates := map[string]interface{}{"COUNT-true-<nil>-false": dparval.NewValue(float64(count))}
				groupDoc.SetAttachment("aggregates", aggregates)
				this.SendItem(groupDoc)
			}
		} else {
			count, err := this.index.ValueCount()
			if err != nil {
				this.SendError(query.NewError(err, "Error counting values in index"))
			} else {
				groupDoc := dparval.NewValue(map[string]interface{}{})
				aggkey := fmt.Sprintf("COUNT-false-%v-false", this.expr)
				aggregates := map[string]interface{}{aggkey: dparval.NewValue(float64(count))}
				groupDoc.SetAttachment("aggregates", aggregates)
				this.SendItem(groupDoc)
			}
		}
	}
	// FIXME eventually we can support counting ranges

	clog.To(CHANNEL, "fastcount operator finished")
}

func (this *FastCount) processItem(item *dparval.Value) bool {
	return true
}

func (this *FastCount) afterItems() {}

func (this *FastCount) SendItem(item *dparval.Value) bool {
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

func (this *FastCount) SendError(err query.Error) bool {
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

func (this *FastCount) RecoverPanic() {
	r := recover()
	if r != nil {
		clog.Error(fmt.Errorf("Query Execution Panic: %v\n%s", r, debug.Stack()))
		this.SendError(query.NewError(nil, "Panic In Exeuction Pipeline"))
	}
}

func (this *FastCount) SetQuery(q network.Query) {
	this.query = q
}
