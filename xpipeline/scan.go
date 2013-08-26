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
	"github.com/couchbaselabs/tuqtng/query"
)

type Scan struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	bucket                catalog.Bucket
	index                 catalog.ScanIndex
	downstreamStopChannel misc.StopChannel
}

func NewScan(bucket catalog.Bucket, index catalog.ScanIndex) *Scan {
	return &Scan{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		bucket:         bucket,
		index:          index,
	}
}

func (this *Scan) SetSource(source Operator) {}

func (this *Scan) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Scan) Run(stopChannel misc.StopChannel) {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	// this MUST be here so that it runs before the channels are closed
	defer this.RecoverPanic()

	this.downstreamStopChannel = stopChannel
	clog.To(CHANNEL, "scan operator starting")

	indexItemChannel := make(dparval.ValueChannel)
	indexWarnChannel := make(query.ErrorChannel)
	indexErrorChannel := make(query.ErrorChannel)
	var item *dparval.Value
	var warn query.Error
	var err query.Error

	go this.index.ScanEntries(indexItemChannel, indexWarnChannel, indexErrorChannel)

	ok := true
	for ok {
		select {
		case item, ok = <-indexItemChannel:
			if ok {
				this.SendItem(item)
			}
		case warn, ok = <-indexWarnChannel:
			if warn != nil {
				this.SendError(warn)
			}
		case err, ok = <-indexErrorChannel:
			if err != nil {
				this.SendError(warn)
			}
		case _, ok = <-stopChannel:
			// downstream has asked us to stop
		}
	}
	clog.To(CHANNEL, "scan operator finished")
}

func (this *Scan) processItem(item *dparval.Value) bool {
	return true
}

func (this *Scan) afterItems() {}

func (this *Scan) SendItem(item *dparval.Value) bool {
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

func (this *Scan) SendError(err query.Error) bool {
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

func (this *Scan) RecoverPanic() {
	r := recover()
	if r != nil {
		clog.Error(fmt.Errorf("Query Execution Panic: %v\n%s", r, debug.Stack()))
		this.SendError(query.NewError(nil, "Panic In Exeuction Pipeline"))
	}
}
