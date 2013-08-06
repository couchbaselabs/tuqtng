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

func TestCompare(t *testing.T) {

	nonExistantProperty := NewProperty("dne")
	null := NewLiteralNull()
	numberSixty := NewLiteralNumber(60.0)
	numberNine := NewLiteralNumber(9.0)
	stringBob := NewLiteralString("bob")
	stringCat := NewLiteralString("cat")
	patternMatchSingle := NewLiteralString("b_b")
	patternMatchMulti := NewLiteralString("b%")
	patternNoMatch := NewLiteralString("____")

	tests := ExpressionTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), false, nil},
		{NewGreaterThanOperator(numberSixty, numberNine), true, nil},
		{NewGreaterThanOperator(numberNine, numberSixty), false, nil},
		{NewGreaterThanOperator(null, numberSixty), nil, nil},
		{NewGreaterThanOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewGreaterThanOperator(stringBob, numberSixty), false, nil},

		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), true, nil},
		{NewGreaterThanOrEqualOperator(numberSixty, numberNine), true, nil},
		{NewGreaterThanOrEqualOperator(numberNine, numberSixty), false, nil},
		{NewGreaterThanOrEqualOperator(null, numberSixty), nil, nil},
		{NewGreaterThanOrEqualOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewGreaterThanOrEqualOperator(stringBob, numberSixty), false, nil},

		{NewLessThanOperator(numberSixty, numberSixty), false, nil},
		{NewLessThanOperator(numberSixty, numberNine), false, nil},
		{NewLessThanOperator(numberNine, numberSixty), true, nil},
		{NewLessThanOperator(null, numberSixty), nil, nil},
		{NewLessThanOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewLessThanOperator(stringBob, numberSixty), false, nil},

		{NewLessThanOrEqualOperator(numberSixty, numberSixty), true, nil},
		{NewLessThanOrEqualOperator(numberSixty, numberNine), false, nil},
		{NewLessThanOrEqualOperator(numberNine, numberSixty), true, nil},
		{NewLessThanOrEqualOperator(null, numberSixty), nil, nil},
		{NewLessThanOrEqualOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewLessThanOrEqualOperator(stringBob, numberSixty), false, nil},

		{NewEqualToOperator(numberSixty, numberSixty), true, nil},
		{NewEqualToOperator(numberSixty, numberNine), false, nil},
		{NewEqualToOperator(numberNine, numberSixty), false, nil},
		{NewEqualToOperator(null, numberSixty), nil, nil},
		{NewEqualToOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewEqualToOperator(stringBob, numberSixty), false, nil},

		{NewNotEqualToOperator(numberSixty, numberSixty), false, nil},
		{NewNotEqualToOperator(numberSixty, numberNine), true, nil},
		{NewNotEqualToOperator(numberNine, numberSixty), true, nil},
		{NewNotEqualToOperator(null, numberSixty), nil, nil},
		{NewNotEqualToOperator(nonExistantProperty, numberSixty), nil, &dparval.Undefined{"dne"}},
		{NewNotEqualToOperator(stringBob, numberSixty), false, nil},

		{NewLikeOperator(stringBob, patternMatchSingle), true, nil},
		{NewLikeOperator(stringCat, patternMatchSingle), false, nil},
		{NewLikeOperator(stringBob, patternMatchMulti), true, nil},
		{NewLikeOperator(stringCat, patternMatchMulti), false, nil},
		{NewLikeOperator(stringBob, patternNoMatch), false, nil},
		{NewLikeOperator(stringBob, numberNine), nil, nil},
		{NewLikeOperator(numberNine, patternMatchSingle), nil, nil},
		{NewLikeOperator(stringBob, nonExistantProperty), nil, &dparval.Undefined{"dne"}},
		{NewLikeOperator(nonExistantProperty, stringBob), nil, &dparval.Undefined{"dne"}},

		{NewNotLikeOperator(stringBob, patternMatchSingle), false, nil},
		{NewNotLikeOperator(stringCat, patternMatchSingle), true, nil},
		{NewNotLikeOperator(stringBob, patternMatchMulti), false, nil},
		{NewNotLikeOperator(stringCat, patternMatchMulti), true, nil},
		{NewNotLikeOperator(stringBob, patternNoMatch), true, nil},
		{NewNotLikeOperator(stringBob, numberNine), nil, nil},
		{NewNotLikeOperator(numberNine, patternMatchSingle), nil, nil},
		{NewNotLikeOperator(stringBob, nonExistantProperty), nil, &dparval.Undefined{"dne"}},
		{NewNotLikeOperator(nonExistantProperty, stringBob), nil, &dparval.Undefined{"dne"}},

		// these tests all conform to the table in the specification
		{NewIsNullOperator(stringBob), false, nil},
		{NewIsNullOperator(null), true, nil},
		{NewIsNullOperator(nonExistantProperty), false, nil},

		{NewIsNotNullOperator(stringBob), true, nil},
		{NewIsNotNullOperator(null), false, nil},
		{NewIsNotNullOperator(nonExistantProperty), false, nil},

		{NewIsMissingOperator(stringBob), false, nil},
		{NewIsMissingOperator(null), false, nil},
		{NewIsMissingOperator(nonExistantProperty), true, nil},

		{NewIsNotMissingOperator(stringBob), true, nil},
		{NewIsNotMissingOperator(null), true, nil},
		{NewIsNotMissingOperator(nonExistantProperty), false, nil},

		{NewIsValuedOperator(stringBob), true, nil},
		{NewIsValuedOperator(null), false, nil},
		{NewIsValuedOperator(nonExistantProperty), false, nil},

		{NewIsNotValuedOperator(stringBob), false, nil},
		{NewIsNotValuedOperator(null), true, nil},
		{NewIsNotValuedOperator(nonExistantProperty), false, nil},

		// new tests to check ordering of missing/null values
		{NewLessThanOperator(numberSixty, nonExistantProperty), nil, &dparval.Undefined{"dne"}},
		{NewLessThanOperator(numberSixty, null), nil, nil},

		//comparing booleans
		{NewEqualToOperator(NewLiteralBool(true), NewLiteralBool(false)), false, nil},
		{NewEqualToOperator(NewLiteralBool(false), NewLiteralBool(true)), false, nil},
	}

	tests.Run(t)

}

func TestCompareStringRepresentation(t *testing.T) {

	numberSixty := NewLiteralNumber(60.0)
	stringBob := NewLiteralString("bob")

	tests := ExpressionStringTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), "60 > 60"},
		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), "60 >= 60"},
		{NewLessThanOperator(numberSixty, numberSixty), "60 < 60"},
		{NewLessThanOrEqualOperator(numberSixty, numberSixty), "60 <= 60"},
		{NewEqualToOperator(numberSixty, numberSixty), "60 = 60"},
		{NewNotEqualToOperator(numberSixty, numberSixty), "60 != 60"},
		{NewLikeOperator(stringBob, stringBob), `"bob" LIKE "bob"`},
		{NewNotLikeOperator(stringBob, stringBob), `"bob" NOT LIKE "bob"`},

		{NewIsNullOperator(stringBob), `"bob" IS NULL`},
		{NewIsNotNullOperator(stringBob), `"bob" IS NOT NULL`},
		{NewIsMissingOperator(stringBob), `"bob" IS MISSING`},
		{NewIsNotMissingOperator(stringBob), `"bob" IS NOT MISSING`},
		{NewIsValuedOperator(stringBob), `"bob" IS VALUED`},
		{NewIsNotValuedOperator(stringBob), `"bob" IS NOT VALUED`},
	}

	tests.Run(t)
}

func TestCompareValidate(t *testing.T) {

	numberSixty := NewLiteralNumber(60.0)
	stringBob := NewLiteralString("bob")

	tests := ExpressionValidateTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), nil},
		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), nil},
		{NewLessThanOperator(numberSixty, numberSixty), nil},
		{NewLessThanOrEqualOperator(numberSixty, numberSixty), nil},
		{NewEqualToOperator(numberSixty, numberSixty), nil},
		{NewNotEqualToOperator(numberSixty, numberSixty), nil},
		{NewLikeOperator(stringBob, stringBob), nil},
		{NewNotLikeOperator(stringBob, stringBob), nil},

		{NewIsNullOperator(stringBob), nil},
		{NewIsNotNullOperator(stringBob), nil},
		{NewIsMissingOperator(stringBob), nil},
		{NewIsNotMissingOperator(stringBob), nil},
		{NewIsValuedOperator(stringBob), nil},
		{NewIsNotValuedOperator(stringBob), nil},

		// first arg not valid
		{NewGreaterThanOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewGreaterThanOrEqualOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewLessThanOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewLessThanOrEqualOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewEqualToOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewNotEqualToOperator(notValidExpression, numberSixty), notValidExpressionError},
		{NewLikeOperator(notValidExpression, stringBob), notValidExpressionError},
		{NewNotLikeOperator(notValidExpression, stringBob), notValidExpressionError},

		{NewIsNullOperator(notValidExpression), notValidExpressionError},
		{NewIsNotNullOperator(notValidExpression), notValidExpressionError},
		{NewIsMissingOperator(notValidExpression), notValidExpressionError},
		{NewIsNotMissingOperator(notValidExpression), notValidExpressionError},
		{NewIsValuedOperator(notValidExpression), notValidExpressionError},
		{NewIsNotValuedOperator(notValidExpression), notValidExpressionError},

		// second arg not valid
		{NewGreaterThanOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewGreaterThanOrEqualOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewLessThanOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewLessThanOrEqualOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewEqualToOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewNotEqualToOperator(numberSixty, notValidExpression), notValidExpressionError},
		{NewLikeOperator(stringBob, notValidExpression), notValidExpressionError},
		{NewNotLikeOperator(stringBob, notValidExpression), notValidExpressionError},
	}

	tests.Run(t)
}

func TestCompareVerifyFormalNotation(t *testing.T) {

	numberSixty := NewLiteralNumber(60.0)
	stringBob := NewLiteralString("bob")

	tests := ExpressionVerifyFormalNotationTestSet{
		{NewGreaterThanOperator(numberSixty, numberSixty), nil, nil},
		{NewGreaterThanOrEqualOperator(numberSixty, numberSixty), nil, nil},
		{NewLessThanOperator(numberSixty, numberSixty), nil, nil},
		{NewLessThanOrEqualOperator(numberSixty, numberSixty), nil, nil},
		{NewEqualToOperator(numberSixty, numberSixty), nil, nil},
		{NewNotEqualToOperator(numberSixty, numberSixty), nil, nil},
		{NewLikeOperator(stringBob, stringBob), nil, nil},
		{NewNotLikeOperator(stringBob, stringBob), nil, nil},

		{NewIsNullOperator(stringBob), nil, nil},
		{NewIsNotNullOperator(stringBob), nil, nil},
		{NewIsMissingOperator(stringBob), nil, nil},
		{NewIsNotMissingOperator(stringBob), nil, nil},
		{NewIsValuedOperator(stringBob), nil, nil},
		{NewIsNotValuedOperator(stringBob), nil, nil},

		// first arg not formal
		{NewGreaterThanOperator(notFormalExpression, numberSixty), nil, notFormalExpressionError},
		{NewGreaterThanOrEqualOperator(notFormalExpression, numberSixty), nil, notFormalExpressionError},
		{NewLessThanOperator(notFormalExpression, numberSixty), nil, notFormalExpressionError},
		{NewLessThanOrEqualOperator(notFormalExpression, numberSixty), nil, notFormalExpressionError},
		{NewEqualToOperator(notFormalExpression, numberSixty), nil, notFormalExpressionError},
		{NewNotEqualToOperator(notFormalExpression, numberSixty), nil, notFormalExpressionError},
		{NewLikeOperator(notFormalExpression, stringBob), nil, notFormalExpressionError},
		{NewNotLikeOperator(notFormalExpression, stringBob), nil, notFormalExpressionError},

		{NewIsNullOperator(notFormalExpression), nil, notFormalExpressionError},
		{NewIsNotNullOperator(notFormalExpression), nil, notFormalExpressionError},
		{NewIsMissingOperator(notFormalExpression), nil, notFormalExpressionError},
		{NewIsNotMissingOperator(notFormalExpression), nil, notFormalExpressionError},
		{NewIsValuedOperator(notFormalExpression), nil, notFormalExpressionError},
		{NewIsNotValuedOperator(notFormalExpression), nil, notFormalExpressionError},

		// second arg not formal
		{NewGreaterThanOperator(numberSixty, notFormalExpression), nil, notFormalExpressionError},
		{NewGreaterThanOrEqualOperator(numberSixty, notFormalExpression), nil, notFormalExpressionError},
		{NewLessThanOperator(numberSixty, notFormalExpression), nil, notFormalExpressionError},
		{NewLessThanOrEqualOperator(numberSixty, notFormalExpression), nil, notFormalExpressionError},
		{NewEqualToOperator(numberSixty, notFormalExpression), nil, notFormalExpressionError},
		{NewNotEqualToOperator(numberSixty, notFormalExpression), nil, notFormalExpressionError},
		{NewLikeOperator(stringBob, notFormalExpression), nil, notFormalExpressionError},
		{NewNotLikeOperator(stringBob, notFormalExpression), nil, notFormalExpressionError},
	}

	tests.Run(t, []string{}, []string{"bucket", "child"}, "")

	// again with single bucket
	tests = ExpressionVerifyFormalNotationTestSet{

		// first arg not formal
		{NewGreaterThanOperator(notFormalExpression, numberSixty), nil, nil},
		{NewGreaterThanOrEqualOperator(notFormalExpression, numberSixty), nil, nil},
		{NewLessThanOperator(notFormalExpression, numberSixty), nil, nil},
		{NewLessThanOrEqualOperator(notFormalExpression, numberSixty), nil, nil},
		{NewEqualToOperator(notFormalExpression, numberSixty), nil, nil},
		{NewNotEqualToOperator(notFormalExpression, numberSixty), nil, nil},
		{NewLikeOperator(notFormalExpression, stringBob), nil, nil},
		{NewNotLikeOperator(notFormalExpression, stringBob), nil, nil},

		{NewIsNullOperator(notFormalExpression), nil, nil},
		{NewIsNotNullOperator(notFormalExpression), nil, nil},
		{NewIsMissingOperator(notFormalExpression), nil, nil},
		{NewIsNotMissingOperator(notFormalExpression), nil, nil},
		{NewIsValuedOperator(notFormalExpression), nil, nil},
		{NewIsNotValuedOperator(notFormalExpression), nil, nil},

		// second arg not formal
		{NewGreaterThanOperator(numberSixty, notFormalExpression), nil, nil},
		{NewGreaterThanOrEqualOperator(numberSixty, notFormalExpression), nil, nil},
		{NewLessThanOperator(numberSixty, notFormalExpression), nil, nil},
		{NewLessThanOrEqualOperator(numberSixty, notFormalExpression), nil, nil},
		{NewEqualToOperator(numberSixty, notFormalExpression), nil, nil},
		{NewNotEqualToOperator(numberSixty, notFormalExpression), nil, nil},
		{NewLikeOperator(stringBob, notFormalExpression), nil, nil},
		{NewNotLikeOperator(stringBob, notFormalExpression), nil, nil},
	}

	tests.Run(t, []string{}, []string{"bucket"}, "bucket")
}

func TestTypeMismatchError(t *testing.T) {
	err := TypeMismatch{4, 3}

	if err.Error() != "Types do not match, 4 3" {
		t.Errorf("Expected Types do not match 4,3, got %v", err.Error())
	}
}
