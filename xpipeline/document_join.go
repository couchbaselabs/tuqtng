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

func (this *DocumentJoin) Run(stopChannel misc.StopChannel) {
	clog.To(CHANNEL, "document join operator starting")
	this.Base.RunOperator(this, stopChannel)
	clog.To(CHANNEL, "document join operator finished")
}

func (this *DocumentJoin) processItem(item *dparval.Value) bool {

	val, err := this.Base.Evaluate(this.Over, item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return true
		default:
			return this.Base.SendError(query.NewError(err, "Internal Error"))
		}
	}

	if val.Type() == dparval.ARRAY {
		ok := true
		index := 0
		for ok {
			inner, err := val.Index(index)
			index = index + 1
			if err != nil {
				switch err := err.(type) {
				case *dparval.Undefined:
					ok = false
				default:
					this.Base.SendError(query.NewError(err, "Internal Error"))
					return false
				}
			} else {
				newItem := item.Duplicate()
				newItem.SetPath(this.As, inner)
				this.Base.SendItem(newItem)
			}
		}
	}

	return true
}

func (this *DocumentJoin) afterItems() {}

func (this *DocumentJoin) SetQuery(q network.Query) {
	this.Base.SetQuery(q)
}
