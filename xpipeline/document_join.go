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
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

type DocumentJoin struct {
	Base *BaseOperator
	Over ast.Expression
	As   string
}

func NewDocumentJoin(over ast.Expression, as string) *DocumentJoin {
	return &DocumentJoin{
		Base: NewBaseOperator(),
		Over: over,
		As:   as,
	}
}

func (this *DocumentJoin) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *DocumentJoin) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *DocumentJoin) Run() {
	this.Base.RunOperator(this)
}

func (this *DocumentJoin) processItem(item *dparval.Value) bool {
	val, err := this.Over.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return true
		default:
			return this.Base.SendError(query.NewError(err, "Internal Error"))
		}
	}

	if val.Type() == dparval.ARRAY {
		overval := val.Value()
		switch overval := overval.(type) {
		case []interface{}:
			// FIXME major cleanup after full converstion to dparval
			// over expression evaluted to array
			// now walk the array and join
			for _, v := range overval {
				itemValue := item.Value()
				newValue := map[string]interface{}{}
				switch itemValue := itemValue.(type) {
				case map[string]interface{}:
					for itemK, itemV := range itemValue {
						newValue[itemK] = itemV
					}
					newValue[this.As] = v
				}
				itemMeta := item.GetAttachment("meta")
				finalItem := dparval.NewValue(newValue)
				if itemMeta != nil {
					finalItem.SetAttachment("meta", itemMeta)
				}
				this.Base.SendItem(finalItem)
			}
		}
	}

	return true
}

func (this *DocumentJoin) afterItems() {}
