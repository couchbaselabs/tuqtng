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
)

type Scan struct {
	itemChannel query.ItemChannel
	errChannel  query.ErrorChannel
	warnChannel query.ErrorChannel
	bucket      catalog.Bucket
	scanner     catalog.Scanner
}

func NewScan(bucket catalog.Bucket, scanner catalog.Scanner) *Scan {
	return &Scan{
		itemChannel: make(query.ItemChannel),
		errChannel:  make(query.ErrorChannel),
		warnChannel: make(query.ErrorChannel),
		bucket:      bucket,
		scanner:     scanner,
	}
}

func (this *Scan) SetSource(source Operator) {
	panic("Cannot set source for a datasource")
}

func (this *Scan) GetChannels() (query.ItemChannel, query.ErrorChannel, query.ErrorChannel) {
	return this.itemChannel, this.warnChannel, this.errChannel
}

func (this *Scan) Run() {
	defer close(this.itemChannel)
	defer close(this.errChannel)
	defer close(this.warnChannel)
	defer this.bucket.Release()

	scannerItemChannel := make(query.ItemChannel)
	scannerWarnChannel := make(query.ErrorChannel)
	scannerErrorChannel := make(query.ErrorChannel)
	var item query.Item
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
				this.warnChannel <- warn
			}
		case err, ok = <-scannerErrorChannel:
			if err != nil {
				this.errChannel <- err
			}
		}
	}
}
