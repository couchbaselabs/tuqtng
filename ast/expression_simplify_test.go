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
	"math"
	"reflect"
	"testing"
)

func TestExpressionSimplify(t *testing.T) {

	tests := []struct {
		input  Expression
		output Expression
	}{
		// simplify arithmetic on literal numbers
		{NewPlusOperator(NewLiteralNumber(5.0), NewLiteralNumber(5.0)), NewLiteralNumber(10.0)},
		{NewSubtractOperator(NewLiteralNumber(5.0), NewLiteralNumber(5.0)), NewLiteralNumber(0.0)},
		{NewMultiplyOperator(NewLiteralNumber(5.0), NewLiteralNumber(5.0)), NewLiteralNumber(25.0)},
		{NewDivideOperator(NewLiteralNumber(5.0), NewLiteralNumber(5.0)), NewLiteralNumber(1.0)},

		// divide by 0
		// FIXME - still subject to debate
		{NewDivideOperator(NewLiteralNumber(5.0), NewLiteralNumber(0.0)), NewLiteralNumber(math.Inf(1))},

		// string concat on literal strings
		{NewStringConcatenateOperator(NewLiteralString("cat"), NewLiteralString("dog")), NewLiteralString("catdog")},

		// arithmetic on mixed datatypes
		{NewPlusOperator(NewLiteralNumber(5.0), NewLiteralString("cat")), NewLiteralNull()},

		// string concat on mixed datatypes
		{NewStringConcatenateOperator(NewLiteralString("cat"), NewLiteralNumber(5.0)), NewLiteralNull()},

		// comparison operators
		{NewGreaterThanOperator(NewLiteralNumber(5.0), NewLiteralNumber(3.0)), NewLiteralBool(true)},
		{NewLessThanOperator(NewLiteralString("b"), NewLiteralString("a")), NewLiteralBool(false)},

		// function call with literal args
		{NewFunctionCall("length", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("marty"))}), NewLiteralNumber(5.0)},

		// a more complex expression
		{NewGreaterThanOperator(NewFunctionCall("length", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("marty"))}), NewLiteralNumber(5.0)), NewLiteralBool(false)},

		// things that should not be simplified at all
		{NewPlusOperator(NewProperty("path"), NewLiteralNumber(5.0)), NewPlusOperator(NewProperty("path"), NewLiteralNumber(5.0))},
	}

	es := NewExpressionSimplifier()
	for _, test := range tests {
		actual, err := test.input.Accept(es)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(test.output, actual) {
			t.Errorf("Expected %v, got %v, for %v", test.output, actual, test.input)
		}
	}

}
