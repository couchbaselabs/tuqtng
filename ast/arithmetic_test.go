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
	"testing"

	"github.com/couchbaselabs/dparval"
)

func TestArithmetic(t *testing.T) {

	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)
	numberNegativeSeven := NewLiteralNumber(-7.0)
	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")
	dneProperty := NewProperty("foo")

	tests := ExpressionTestSet{
		{NewPlusOperator(stringCouchbase, stringServer), nil, nil}, // no longer support string concatenation, uses different operator

		{NewPlusOperator(numberSeven, numberSeven), 14.0, nil},
		{NewPlusOperator(numberSeven, stringCouchbase), nil, nil},
		{NewPlusOperator(stringCouchbase, numberSeven), nil, nil},
		{NewPlusOperator(dneProperty, numberSeven), nil, &dparval.Undefined{"foo"}},
		{NewPlusOperator(numberSeven, dneProperty), nil, &dparval.Undefined{"foo"}},

		{NewSubtractOperator(numberSeven, numberSeven), 0.0, nil},
		{NewSubtractOperator(numberSeven, stringCouchbase), nil, nil},
		{NewSubtractOperator(stringCouchbase, numberSeven), nil, nil},
		{NewSubtractOperator(dneProperty, numberSeven), nil, &dparval.Undefined{"foo"}},
		{NewSubtractOperator(numberSeven, dneProperty), nil, &dparval.Undefined{"foo"}},

		{NewMultiplyOperator(numberSeven, numberSeven), 49.0, nil},
		{NewMultiplyOperator(numberSeven, stringCouchbase), nil, nil},
		{NewMultiplyOperator(stringCouchbase, numberSeven), nil, nil},
		{NewMultiplyOperator(dneProperty, numberSeven), nil, &dparval.Undefined{"foo"}},
		{NewMultiplyOperator(numberSeven, dneProperty), nil, &dparval.Undefined{"foo"}},

		{NewDivideOperator(numberSeven, numberSeven), 1.0, nil},
		{NewDivideOperator(numberSeven, stringCouchbase), nil, nil},
		{NewDivideOperator(stringCouchbase, numberSeven), nil, nil},
		{NewDivideOperator(dneProperty, numberSeven), nil, &dparval.Undefined{"foo"}},
		{NewDivideOperator(numberSeven, dneProperty), nil, &dparval.Undefined{"foo"}},

		{NewModuloOperator(numberSeven, numberSix), 1.0, nil},
		{NewModuloOperator(stringCouchbase, numberSix), nil, nil},
		{NewModuloOperator(numberSeven, stringCouchbase), nil, nil},
		{NewModuloOperator(dneProperty, numberSix), nil, &dparval.Undefined{"foo"}},
		{NewModuloOperator(numberSeven, dneProperty), nil, &dparval.Undefined{"foo"}},

		{NewChangeSignOperator(numberSeven), -7.0, nil},
		{NewChangeSignOperator(numberNegativeSeven), 7.0, nil},
		{NewChangeSignOperator(stringCouchbase), nil, nil},
		{NewChangeSignOperator(dneProperty), nil, &dparval.Undefined{"foo"}},
	}

	tests.Run(t)
}

func TestArithmeticStringRepresentation(t *testing.T) {
	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)

	tests := ExpressionStringTestSet{
		{NewPlusOperator(numberSeven, numberSeven), `7 + 7`},
		{NewSubtractOperator(numberSeven, numberSeven), `7 - 7`},
		{NewMultiplyOperator(numberSeven, numberSeven), `7 * 7`},
		{NewDivideOperator(numberSeven, numberSeven), `7 / 7`},
		{NewModuloOperator(numberSeven, numberSix), `7 % 6`},
		{NewChangeSignOperator(numberSeven), `-7`},
	}

	tests.Run(t)
}

func TestArithmeticValidate(t *testing.T) {

	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)

	tests := ExpressionValidateTestSet{
		{NewPlusOperator(numberSeven, numberSeven), nil},
		{NewSubtractOperator(numberSeven, numberSeven), nil},
		{NewMultiplyOperator(numberSeven, numberSeven), nil},
		{NewDivideOperator(numberSeven, numberSeven), nil},
		{NewModuloOperator(numberSeven, numberSix), nil},
		{NewChangeSignOperator(numberSeven), nil},
		// first arg invalid
		{NewPlusOperator(notValidExpression, numberSeven), notValidExpressionError},
		{NewSubtractOperator(notValidExpression, numberSeven), notValidExpressionError},
		{NewMultiplyOperator(notValidExpression, numberSeven), notValidExpressionError},
		{NewDivideOperator(notValidExpression, numberSeven), notValidExpressionError},
		{NewModuloOperator(notValidExpression, numberSix), notValidExpressionError},
		{NewChangeSignOperator(notValidExpression), notValidExpressionError},
		// second arg invalid
		{NewPlusOperator(numberSeven, notValidExpression), notValidExpressionError},
		{NewSubtractOperator(numberSeven, notValidExpression), notValidExpressionError},
		{NewMultiplyOperator(numberSeven, notValidExpression), notValidExpressionError},
		{NewDivideOperator(numberSeven, notValidExpression), notValidExpressionError},
		{NewModuloOperator(numberSeven, notValidExpression), notValidExpressionError},
	}

	tests.Run(t)
}

func TestArithmeticVerifyFormalNotation(t *testing.T) {

	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)

	tests := ExpressionVerifyFormalNotationTestSet{
		{NewPlusOperator(numberSeven, numberSeven), NewPlusOperator(numberSeven, numberSeven), nil},
		{NewSubtractOperator(numberSeven, numberSeven), NewSubtractOperator(numberSeven, numberSeven), nil},
		{NewMultiplyOperator(numberSeven, numberSeven), NewMultiplyOperator(numberSeven, numberSeven), nil},
		{NewDivideOperator(numberSeven, numberSeven), NewDivideOperator(numberSeven, numberSeven), nil},
		{NewModuloOperator(numberSeven, numberSix), NewModuloOperator(numberSeven, numberSix), nil},
		{NewChangeSignOperator(numberSeven), NewChangeSignOperator(numberSeven), nil},

		// first arg not formal
		{NewPlusOperator(notFormalExpression, numberSeven), NewPlusOperator(notFormalExpression, numberSeven), notFormalExpressionError},
		{NewSubtractOperator(notFormalExpression, numberSeven), NewSubtractOperator(notFormalExpression, numberSeven), notFormalExpressionError},
		{NewMultiplyOperator(notFormalExpression, numberSeven), NewMultiplyOperator(notFormalExpression, numberSeven), notFormalExpressionError},
		{NewDivideOperator(notFormalExpression, numberSeven), NewDivideOperator(notFormalExpression, numberSeven), notFormalExpressionError},
		{NewModuloOperator(notFormalExpression, numberSix), NewModuloOperator(notFormalExpression, numberSix), notFormalExpressionError},
		{NewChangeSignOperator(notFormalExpression), NewChangeSignOperator(notFormalExpression), notFormalExpressionError},

		// second arg not formal
		{NewPlusOperator(numberSeven, notFormalExpression), NewPlusOperator(numberSeven, notFormalExpression), notFormalExpressionError},
		{NewSubtractOperator(numberSeven, notFormalExpression), NewSubtractOperator(numberSeven, notFormalExpression), notFormalExpressionError},
		{NewMultiplyOperator(numberSeven, notFormalExpression), NewMultiplyOperator(numberSeven, notFormalExpression), notFormalExpressionError},
		{NewDivideOperator(numberSeven, notFormalExpression), NewDivideOperator(numberSeven, notFormalExpression), notFormalExpressionError},
		{NewModuloOperator(numberSeven, notFormalExpression), NewModuloOperator(numberSeven, notFormalExpression), notFormalExpressionError},
	}

	tests.Run(t, []string{}, []string{"bucket", "child"}, "")

	// same tests again with one alias
	tests = ExpressionVerifyFormalNotationTestSet{
		// first arg not formal
		{NewPlusOperator(notFormalExpression, numberSeven), NewPlusOperator(notFormalExpressionAfter, numberSeven), nil},
		{NewSubtractOperator(notFormalExpression, numberSeven), NewSubtractOperator(notFormalExpressionAfter, numberSeven), nil},
		{NewMultiplyOperator(notFormalExpression, numberSeven), NewMultiplyOperator(notFormalExpressionAfter, numberSeven), nil},
		{NewDivideOperator(notFormalExpression, numberSeven), NewDivideOperator(notFormalExpressionAfter, numberSeven), nil},
		{NewModuloOperator(notFormalExpression, numberSix), NewModuloOperator(notFormalExpressionAfter, numberSix), nil},
		{NewChangeSignOperator(notFormalExpression), NewChangeSignOperator(notFormalExpressionAfter), nil},

		// second arg not formal
		{NewPlusOperator(numberSeven, notFormalExpression), NewPlusOperator(numberSeven, notFormalExpressionAfter), nil},
		{NewSubtractOperator(numberSeven, notFormalExpression), NewSubtractOperator(numberSeven, notFormalExpressionAfter), nil},
		{NewMultiplyOperator(numberSeven, notFormalExpression), NewMultiplyOperator(numberSeven, notFormalExpressionAfter), nil},
		{NewDivideOperator(numberSeven, notFormalExpression), NewDivideOperator(numberSeven, notFormalExpressionAfter), nil},
		{NewModuloOperator(numberSeven, notFormalExpression), NewModuloOperator(numberSeven, notFormalExpressionAfter), nil},
	}

	tests.Run(t, []string{}, []string{"bucket"}, "bucket")
}

func TestArithmeticEquivalentTo(t *testing.T) {

	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)

	tests := ExpressionEquivalenceTestSet{
		{NewPlusOperator(numberSix, numberSeven), NewPlusOperator(numberSix, numberSeven), true},
		{NewPlusOperator(numberSix, numberSeven), NewPlusOperator(numberSeven, numberSix), true},
		{NewSubtractOperator(numberSix, numberSeven), NewSubtractOperator(numberSix, numberSeven), true},
		{NewSubtractOperator(numberSix, numberSeven), NewSubtractOperator(numberSeven, numberSix), false},
		{NewMultiplyOperator(numberSix, numberSeven), NewMultiplyOperator(numberSix, numberSeven), true},
		{NewMultiplyOperator(numberSix, numberSeven), NewMultiplyOperator(numberSeven, numberSix), true},
		{NewDivideOperator(numberSix, numberSeven), NewDivideOperator(numberSix, numberSeven), true},
		{NewDivideOperator(numberSix, numberSeven), NewDivideOperator(numberSeven, numberSix), false},
		{NewModuloOperator(numberSix, numberSeven), NewModuloOperator(numberSix, numberSeven), true},
		{NewModuloOperator(numberSix, numberSeven), NewModuloOperator(numberSeven, numberSix), false},
		{NewChangeSignOperator(NewLiteralNumber(7.5)), NewChangeSignOperator(NewLiteralNumber(7.5)), true},
		{NewChangeSignOperator(NewLiteralNumber(7.5)), NewChangeSignOperator(NewLiteralNumber(4.5)), false},
		{NewChangeSignOperator(NewLiteralNumber(7.5)), NewLiteralNumber(4.5), false},
	}

	tests.Run(t)
}
