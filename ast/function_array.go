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
		result = append(lv, rv)
		return dparval.NewValue(result), nil
	}
	return nil, err
}

func (this *FunctionCallArrayConcat) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
