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
	"github.com/couchbaselabs/dparval"
)

type Filter struct {
	Source         Operator
	Expr           ast.Expression
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	ok             bool
}

func NewFilter(expr ast.Expression) *Filter {
	return &Filter{
		Expr:           expr,
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
	}
}

func (this *Filter) SetSource(source Operator) {
	this.Source = source
}

func (this *Filter) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Filter) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	go this.Source.Run()

	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	this.ok = true
	for this.ok {
		select {
		case item, this.ok = <-sourceItemChannel:
			if this.ok {
				this.processItem(item)
			}
		case obj, this.ok = <-supportChannel:
			if this.ok {
				switch obj := obj.(type) {
				case query.Error:
					this.supportChannel <- obj
					if obj.IsFatal() {
						return
					}
				default:
					this.supportChannel <- obj
				}
			}
		}
	}
}

func (this *Filter) processItem(item *dparval.Value) {
	val, err := this.Expr.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			//ignore these
			return
		default:
			this.supportChannel <- query.NewError(err, "error evaluating filter")
			this.ok = false
			return
		}
	}

	valval := val.Value()
	boolVal := ast.ValueInBooleanContext(valval)
	switch boolVal := boolVal.(type) {
	case bool:
		if boolVal {
			this.itemChannel <- item
		}
	}

}
