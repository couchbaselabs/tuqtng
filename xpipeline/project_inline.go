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
	"github.com/mschoch/dparval"
)

type ProjectInline struct {
	Source         Operator
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	Result         *ast.ResultExpression
	ok             bool
}

func NewProjectInline(result *ast.ResultExpression) *ProjectInline {
	return &ProjectInline{
		Result:         result,
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
	}
}

func (this *ProjectInline) SetSource(source Operator) {
	this.Source = source
}

func (this *ProjectInline) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *ProjectInline) Run() {
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

func (this *ProjectInline) processItem(item *dparval.Value) {
	var err error
	var res interface{}

	if this.Result.Star {
		if this.Result.Expr != nil {
			// evaluate this expression first
			val, err := this.Result.Expr.Evaluate(item)
			if err != nil {
				switch err := err.(type) {
				case *dparval.Undefined:
					// undefined contributes nothing to the result
					// but otherwise is NOT an error
					// FIXME review if this should be a warning
					return
				default:
					this.supportChannel <- query.NewError(err, "unexpected error projecting dot star expression")
					this.ok = false
					return
				}
			}
			if val.Type() == dparval.OBJECT {
				valval := val.Value()
				switch valval := valval.(type) {
				case map[string]interface{}:
					// then if the result was an object
					// then project it
					res = valval
				}
			}
		} else {
			// just a star, make item the result
			res = item.Value()
		}
	} else if this.Result.Expr != nil {
		// evaluate the expression
		val, err := this.Result.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// undefined contributes nothing to the result map
				// but otherwise is NOT an error
				// FIXME review if this should be a warning
				return
			default:
				this.supportChannel <- query.NewError(err, "unexpected error projecting expression")
				this.ok = false
				return
			}
		}
		res = val
	}

	// create the actual result Item
	finalItem := dparval.NewValue(res)
	itemMetaVal := item.Meta()
	itemMetaData, err := itemMetaVal.Path("meta")
	if err != nil {
		this.supportChannel <- query.NewError(err, "unable to find item metadata")
		this.ok = false
		return
	}
	finalItem.AddMeta("meta", itemMetaData)

	// write this to the output
	this.itemChannel <- finalItem
}
