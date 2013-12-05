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

// ***********************************************************************************
// Array Utility Functions
// ***********************************************************************************

type FunctionCallArrayConcat struct {
	FunctionCall
}

func NewFunctionCallArrayConcat(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallArrayConcat{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_CONCAT",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallArrayConcat) Copy() Expression {
	return &FunctionCallArrayConcat{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_CONCAT",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallArrayConcat) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothArray, err := this.EvaluateBothRequireArray(context)
	if err != nil {
		return nil, err
	}

	if bothArray {
		result = append(lv, rv...)
		return dparval.NewValue(result), nil
	}
	return nil, err
}

func (this *FunctionCallArrayConcat) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// append an element to the array. Only the first operand should be an array
type FunctionCallArrayAppend struct {
	FunctionCall
}

func NewFunctionCallArrayAppend(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallArrayAppend{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_APPEND",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallArrayAppend) Copy() Expression {
	return &FunctionCallArrayAppend{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_APPEND",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallArrayAppend) Evaluate(context *dparval.Value) (*dparval.Value, error) {

	var result interface{} = nil

	lv, rv, ok, err := this.EvaluateOperandsForArrayAppend(context)

	if err != nil {
		return nil, err
	}

	if ok {
		result = append(lv, rv)
		return dparval.NewValue(result), nil
	}

	return nil, err
}

func (this *FunctionCallArrayAppend) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// prepend an element to the array. Only the second operand should be an array

type FunctionCallArrayPrepend struct {
	FunctionCall
}

func NewFunctionCallArrayPrepend(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallArrayPrepend{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_PREPEND",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallArrayPrepend) Copy() Expression {
	return &FunctionCallArrayPrepend{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_PREPEND",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallArrayPrepend) Evaluate(context *dparval.Value) (*dparval.Value, error) {

	var result []interface{} = nil

	lv, rv, ok, err := this.EvaluateOperandsForArrayPrepend(context)

	if err != nil {
		return nil, err
	}

	if ok {
		result = make([]interface{}, 1)
		result[0] = lv
		result = append(result, rv...)
		return dparval.NewValue(result), nil
	}

	return nil, err
}

func (this *FunctionCallArrayPrepend) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
