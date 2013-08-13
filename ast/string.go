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

// ****************************************************************************
// String Concatenation
// ****************************************************************************

type StringConcatenateOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewStringConcatenateOperator(left, right Expression) *StringConcatenateOperator {
	return &StringConcatenateOperator{
		"string_concat",
		BinaryOperator{
			operator: "||",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *StringConcatenateOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothString, err := this.EvaluateBothRequireString(context)
	if err != nil {
		return nil, err
	}

	if bothString {
		result = lv + rv
	}
	return dparval.NewValue(result), nil
}

func (this *StringConcatenateOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
