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
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

type Filter struct {
	Base *BaseOperator
	Expr ast.Expression
}

func NewFilter(expr ast.Expression) *Filter {
	return &Filter{
		Base: NewBaseOperator(),
		Expr: expr,
	}
}

func (this *Filter) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *Filter) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *Filter) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "filter operator starting")
	this.Base.RunOperator(this, stopChannel)
	clog.To(CHANNEL, "filter operator finished")
}

func (this *Filter) processItem(item *dparval.Value) bool {
	val, err := this.Base.Evaluate(this.Expr, item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			//ignore these
			return true
		default:
			return this.Base.SendError(query.NewError(err, "error evaluating filter"))
		}
	}

	valval := val.Value()
	boolVal := ast.ValueInBooleanContext(valval)
	switch boolVal := boolVal.(type) {
	case bool:
		if boolVal {
			return this.Base.SendItem(item)
		}
	}
	return true
}

func (this *Filter) afterItems() {}

func (this *Filter) SetQuery(q network.Query) {
	this.Base.SetQuery(q)
}
