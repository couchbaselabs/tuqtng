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

type ProjectInline struct {
	Base   *BaseOperator
	Result *ast.ResultExpression
}

func NewProjectInline(result *ast.ResultExpression) *ProjectInline {
	return &ProjectInline{
		Base:   NewBaseOperator(),
		Result: result,
	}
}

func (this *ProjectInline) SetSource(source Operator) {
	this.Base.SetSource(source)
}

func (this *ProjectInline) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.Base.GetChannels()
}

func (this *ProjectInline) Run() {
	this.Base.RunOperator(this)
}

func (this *ProjectInline) processItem(item *dparval.Value) bool {
	var res interface{}

	val, err := projectedValueOfResultExpression(item, this.Result)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined contributes nothing to the result map
			// but otherwise is NOT an error
			// FIXME review if this should be a warning
			return true
		default:
			return this.Base.SendError(query.NewError(err, "unexpected error projecting expression"))
		}
	}
	res = val

	// create the actual result Item
	finalItem := dparval.NewValue(res)
	itemMeta := item.GetAttachment("meta")
	if itemMeta != nil {
		finalItem.SetAttachment("meta", itemMeta)
	}

	// write this to the output
	return this.Base.SendItem(finalItem)
}

func (this *ProjectInline) afterItems() {}
