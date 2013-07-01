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

	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

type Filter struct {
	Source      Operator
	Expr        ast.Expression
	itemChannel query.ItemChannel
}

func NewFilter(expr ast.Expression) *Filter {
	return &Filter{
		Expr:        expr,
		itemChannel: make(query.ItemChannel),
	}
}

func (this *Filter) SetSource(source Operator) {
	this.Source = source
}

func (this *Filter) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Filter) Run() {
	defer close(this.itemChannel)

	go this.Source.Run()

	for item := range this.Source.GetItemChannel() {
		val, err := this.Expr.Evaluate(item)
		if err == nil {
			boolVal := ast.ValueInBooleanContext(val)
			switch boolVal := boolVal.(type) {
			case bool:
				if boolVal {
					this.itemChannel <- item
				}
			}
		} else {
			switch err := err.(type) {
			case *query.Undefined:
			//ignore these
			default:
				log.Printf("Error evaluating filter: %v", err)
			}
		}
	}
}
