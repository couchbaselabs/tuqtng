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
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/mschoch/dparval"
)

type Scan struct {
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	bucket         catalog.Bucket
	scanner        catalog.Scanner
}

func NewScan(bucket catalog.Bucket, scanner catalog.Scanner) *Scan {
	return &Scan{
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		bucket:         bucket,
		scanner:        scanner,
	}
}

func (this *Scan) SetSource(source Operator) {
	panic("Cannot set source for a datasource")
}

func (this *Scan) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Scan) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)
	defer this.bucket.Release()

	scannerItemChannel := make(dparval.ValueChannel)
	scannerWarnChannel := make(query.ErrorChannel)
	scannerErrorChannel := make(query.ErrorChannel)
	var item *dparval.Value
	var warn query.Error
	var err query.Error

	go this.scanner.ScanAll(scannerItemChannel, scannerWarnChannel, scannerErrorChannel)

	ok := true
	for ok {
		select {
		case item, ok = <-scannerItemChannel:
			if ok {
				this.itemChannel <- item
			}
		case warn, ok = <-scannerWarnChannel:
			if warn != nil {
				this.supportChannel <- warn
			}
		case err, ok = <-scannerErrorChannel:
			if err != nil {
				this.supportChannel <- err
			}
		}
	}
}
