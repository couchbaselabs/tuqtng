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
	"reflect"
	"testing"
)

func TestExpressionNNF(t *testing.T) {

	bt := NewLiteralBool(true)
	bf := NewLiteralBool(false)

	tests := []struct {
		input  Expression
		output Expression
	}{
		// distribute NOT over AND
		{
			NewNotOperator(NewAndOperator(ExpressionList{bt, bt})),
			NewOrOperator(ExpressionList{NewNotOperator(bt), NewNotOperator(bt)}),
		},
		// distribute NOT over OR
		{
			NewNotOperator(NewOrOperator(ExpressionList{bt, bf})),
			NewAndOperator(ExpressionList{NewNotOperator(bt), NewNotOperator(bf)}),
		},
		// eliminiate double NOT
		{
			NewNotOperator(NewNotOperator(bt)),
			bt,
		},
		// reduce 3x NOT to NOT
		{
			NewNotOperator(NewNotOperator(NewNotOperator(bt))),
			NewNotOperator(bt),
		},
		// eliminate a double NOT that arises from the transformation
		{
			NewNotOperator(NewAndOperator(ExpressionList{NewNotOperator(bt), bf})),
			NewOrOperator(ExpressionList{bt, NewNotOperator(bf)}),
		},
		// test works across nested structures as well
		{
			NewAndOperator(ExpressionList{NewNotOperator(NewAndOperator(ExpressionList{NewNotOperator(bt), bf})), bt}),
			NewAndOperator(ExpressionList{NewOrOperator(ExpressionList{bt, NewNotOperator(bf)}), bt}),
		},
	}

	ennf := NewExpressionNNF()
	for _, test := range tests {
		actual, err := test.input.Accept(ennf)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(test.output, actual) {
			t.Errorf("Expected %v, got %v, for %v", test.output, actual, test.input)
		}
	}

}
