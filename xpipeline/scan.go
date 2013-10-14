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
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
)

type Scan struct {
	itemChannel           dparval.ValueChannel
	supportChannel        PipelineSupportChannel
	bucket                catalog.Bucket
	index                 catalog.ScanIndex
	downstreamStopChannel misc.StopChannel
	ranges                plan.ScanRanges
	as                    string
}

func NewScan(bucket catalog.Bucket, index catalog.ScanIndex, ranges plan.ScanRanges, as string) *Scan {
	return &Scan{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		bucket:         bucket,
		index:          index,
		ranges:         ranges,
		as:             as,
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

	clog.To(CHANNEL, "scan operator starting")

	if this.ranges == nil {
		this.scanRange(nil)
	} else {
		for _, scanRange := range this.ranges {
			ok := this.scanRange(scanRange)
			if !ok {
				break
			}
		}
	}

	clog.To(CHANNEL, "scan operator finished")
}

func (this *Scan) scanRange(scanRange *plan.ScanRange) bool {

	indexItemChannel := make(catalog.EntryChannel)
	indexWarnChannel := make(query.ErrorChannel)
	indexErrorChannel := make(query.ErrorChannel)

	clog.To(CHANNEL, "scanning range %v", scanRange)
	if scanRange == nil {
		go this.index.ScanEntries(0, indexItemChannel, indexWarnChannel, indexErrorChannel)
	} else {
		rangeScan, ok := this.index.(catalog.RangeIndex)
		if ok {
			go rangeScan.ScanRange(scanRange.Low, scanRange.High, scanRange.Inclusion, scanRange.Limit, indexItemChannel, indexWarnChannel, indexErrorChannel)
		} else {
			this.SendError(query.NewError(nil, "Cannot range scan this"))
			return false
		}
	}

	var item *catalog.IndexEntry
	var warn query.Error
	var err query.Error

	ok := true
	for ok {
		select {
		case item, ok = <-indexItemChannel:
			if ok {
				// rematerialize an object from the data returned by this index entry
				doc := dparval.NewValue(map[string]interface{}{})

				hackIndex, ok := this.index.(catalog.Index)
				if ok {
					if hackIndex.Key() != nil && item.EntryKey != nil && len(hackIndex.Key()) == len(item.EntryKey) {
						for i, key := range hackIndex.Key() {
							entry := item.EntryKey[i]
							doc = this.buildValue(key, entry)
						}
					}
				}

				// attach metadata
				doc.SetAttachment("meta", map[string]interface{}{"id": item.PrimaryKey})

				if this.as != "" {
					this.SendItem(dparval.NewValue(map[string]interface{}{this.as: doc}))
				} else {
					this.SendItem(doc)
				}
			}
		case warn, ok = <-indexWarnChannel:
			if warn != nil {
				this.SendError(warn)
			}
		case err, ok = <-indexErrorChannel:
			if err != nil {
				this.SendError(err)
				return false
			}
		case _, ok = <-this.downstreamStopChannel:
			// downstream has asked us to stop
			return false
		}
	}

	return true
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

func (this *Scan) buildValue(key ast.Expression, val *dparval.Value) *dparval.Value {
	doc := dparval.NewValue(map[string]interface{}{})
	switch key := key.(type) {
	case *ast.LiteralNumber:
		doc = dparval.NewValue([]interface{}{})
		index := int(key.Val)
		doc.SetIndex(index, val)
	case *ast.Property:
		doc.SetPath(key.Path, val)
	case *ast.DotMemberOperator:
		doc = this.buildValue(key.Left, this.buildValue(key.Right, val))
	case *ast.BracketMemberOperator:
		doc = this.buildValue(key.Left, this.buildValue(key.Right, val))
	default:
		clog.Error(fmt.Errorf("Unsupported key type %T encountered uncovering query", key))
	}
	return doc
}
