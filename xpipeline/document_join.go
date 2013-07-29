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
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

type DocumentJoin struct {
	Source      Operator
	Over        ast.Expression
	As          string
	itemChannel query.ItemChannel
	errChannel  query.ErrorChannel
	warnChannel query.ErrorChannel
}

func NewDocumentJoin(over ast.Expression, as string) *DocumentJoin {
	return &DocumentJoin{
		Over:        over,
		As:          as,
		itemChannel: make(query.ItemChannel),
		errChannel:  make(query.ErrorChannel),
		warnChannel: make(query.ErrorChannel),
	}
}

func (this *DocumentJoin) SetSource(source Operator) {
	this.Source = source
}

func (this *DocumentJoin) GetChannels() (query.ItemChannel, query.ErrorChannel, query.ErrorChannel) {
	return this.itemChannel, this.warnChannel, this.errChannel
}

func (this *DocumentJoin) Run() {
	defer close(this.itemChannel)
	defer close(this.errChannel)
	defer close(this.warnChannel)

	go this.Source.Run()

	var item query.Item
	var warn query.Error
	var err query.Error
	sourceItemChannel, sourceWarnChannel, sourceErrorChannel := this.Source.GetChannels()
	ok := true
	for ok {
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

func (this *DocumentJoin) processItem(item query.Item) {
	val, err := this.Over.Evaluate(item)
	if err == nil {
		switch val := val.(type) {
		case []query.Value:
			// over expression evaluted to array
			// now walk the array and join
			for _, v := range val {
				itemValue := item.GetValue()
				newValue := map[string]query.Value{}
				switch itemValue := itemValue.(type) {
				case map[string]query.Value:
					for itemK, itemV := range itemValue {
						newValue[itemK] = itemV
					}
					newValue[this.As] = v
				case map[string]interface{}:
					for itemK, itemV := range itemValue {
						newValue[itemK] = itemV
					}
					newValue[this.As] = v
				}
				itemMeta := item.GetMeta()
				finalItem := query.NewParsedItem(newValue, itemMeta)
				this.itemChannel <- finalItem
			}
		}
	} else {
		switch err := err.(type) {
		case *query.Undefined:
		//ignore these
		default:
			this.errChannel <- query.NewError(err, "Internal Error")
			return
		}
	}
}
