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
	"log"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type Scan struct {
	itemChannel query.ItemChannel
	scanner     catalog.Scanner
}

func NewScan(scanner catalog.Scanner) *Scan {
	return &Scan{
		itemChannel: make(query.ItemChannel),
		scanner:     scanner,
	}
}

func (this *Scan) SetSource(source Operator) {
	panic("Cannot set source for a datasource")
}

func (this *Scan) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Scan) Run() {
	defer close(this.itemChannel)

	scannerItemChannel := make(query.ItemChannel)
	scannerErrorChannel := make(query.ErrorChannel)
	var item query.Item
	var err query.Error

	go this.scanner.ScanAll(scannerItemChannel, scannerErrorChannel)

	ok := true
	for ok {
		select {
		case item, ok = <-scannerItemChannel:
			if ok {
				this.itemChannel <- item
			}
		case err, ok = <-scannerErrorChannel:
			if err != nil {
				log.Printf("underlying scanner gave error: %v", err)
				// FIXME proper error handling
			}
		}
	}
}
