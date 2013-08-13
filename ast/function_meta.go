//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"github.com/couchbaselabs/dparval"
)

type FunctionCallMeta struct {
	FunctionCall
}

func NewFunctionCallMeta(operands FunctionArgExpressionList) Expression {
	return &FunctionCallMeta{
		FunctionCall{
			Type:     "function",
			Name:     "META",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallMeta) Evaluate(item *dparval.Value) (*dparval.Value, error) {

	// FIXME the commented code below wont work until we fix how we store meta

	// av, err := arguments[0].Expr.Evaluate(item)
	// if err != nil {
	// 	return nil, err
	// }

	// meta := av.Meta()
	// if meta != nil {
	// 	metaData, err := meta.Path("meta")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return metaData, nil
	// }

	meta := item.GetAttachment("meta")
	metaValue := dparval.NewValue(meta)
	return metaValue, nil
}

func (this *FunctionCallMeta) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallValue struct {
	FunctionCall
}

func NewFunctionCallValue(operands FunctionArgExpressionList) Expression {
	return &FunctionCallValue{
		FunctionCall{
			Type:     "function",
			Name:     "VALUE",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallValue) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	if len(this.Operands) > 0 {
		// first evaluate the argument
		av, err := this.Operands[0].Expr.Evaluate(item)
		if err != nil {
			return nil, err
		}
		return av, nil
	} else {
		// this mode is still relied up for projecting in the FROM clause
		// review for cleanup
		return item, nil
	}
}

func (this *FunctionCallValue) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
