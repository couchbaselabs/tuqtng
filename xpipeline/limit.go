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
	"github.com/couchbaselabs/tuqtng/query"
)

type Limit struct {
	Source      Operator
	Limit       int
	itemChannel query.ItemChannel
	errChannel  query.ErrorChannel
	warnChannel query.ErrorChannel
	count       int
}

func NewLimit(limit int) *Limit {
	return &Limit{
		Limit:       limit,
		itemChannel: make(query.ItemChannel),
		errChannel:  make(query.ErrorChannel),
		warnChannel: make(query.ErrorChannel),
	}
}

func (this *Limit) SetSource(source Operator) {
	this.Source = source
}

func (this *Limit) GetChannels() (query.ItemChannel, query.ErrorChannel, query.ErrorChannel) {
	return this.itemChannel, this.warnChannel, this.errChannel
}

func (this *Limit) Run() {
	defer close(this.itemChannel)
	defer close(this.errChannel)
	defer close(this.warnChannel)

	this.count = 0

	go this.Source.Run()

	var item query.Item
	var warn query.Error
	var err query.Error
	sourceItemChannel, sourceWarnChannel, sourceErrorChannel := this.Source.GetChannels()
	ok := true
	for ok && this.count < this.Limit {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				this.processItem(item)
			}
		case warn, ok = <-sourceWarnChannel:
			// propogate the warning
			if warn != nil {
				this.warnChannel <- warn
			}
		case err, ok = <-sourceErrorChannel:
			// propogate the error and return
			if err != nil {
				this.errChannel <- err
				return
			}
		}
	}
}

func (this *Limit) processItem(item query.Item) {
	this.itemChannel <- item
	this.count++
}
