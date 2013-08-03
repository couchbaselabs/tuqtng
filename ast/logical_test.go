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

func TestBooleanStringRepresentation(t *testing.T) {

	booleanTrue := NewLiteralBool(true)
	booleanFalse := NewLiteralBool(false)

	tests := ExpressionStringTestSet{
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue}), "true AND true"},
		{NewAndOperator(ExpressionList{booleanTrue, booleanFalse}), "true AND false"},
		{NewAndOperator(ExpressionList{booleanFalse, booleanTrue}), "false AND true"},
		{NewAndOperator(ExpressionList{booleanFalse, booleanFalse}), "false AND false"},

		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanTrue}), "true AND true AND true AND true AND true"},
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanFalse}), "true AND true AND true AND true AND false"},

		{NewOrOperator(ExpressionList{booleanTrue, booleanTrue}), "true OR true"},
		{NewOrOperator(ExpressionList{booleanTrue, booleanFalse}), "true OR false"},
		{NewOrOperator(ExpressionList{booleanFalse, booleanTrue}), "false OR true"},
		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse}), "false OR false"},

		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanTrue}), "false OR false OR false OR false OR true"},
		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanFalse}), "false OR false OR false OR false OR false"},

		{NewNotOperator(booleanTrue), "NOT true"},
		{NewNotOperator(booleanFalse), "NOT false"},
	}

	tests.Run(t)

}

func TestBooleanValidate(t *testing.T) {

	booleanTrue := NewLiteralBool(true)

	tests := ExpressionValidateTestSet{
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue}), nil},
		{NewOrOperator(ExpressionList{booleanTrue, booleanTrue}), nil},
		{NewNotOperator(booleanTrue), nil},
		// first arg not valid
		{NewAndOperator(ExpressionList{notValidExpression, booleanTrue}), notValidExpressionError},
		{NewOrOperator(ExpressionList{notValidExpression, booleanTrue}), notValidExpressionError},
		{NewNotOperator(notValidExpression), notValidExpressionError},
		// second arg not valid
		{NewAndOperator(ExpressionList{booleanTrue, notValidExpression}), notValidExpressionError},
		{NewOrOperator(ExpressionList{booleanTrue, notValidExpression}), notValidExpressionError},
	}

	tests.Run(t)
}

func TestBooleanVerifyFormalNotation(t *testing.T) {

	booleanTrue := NewLiteralBool(true)

	tests := ExpressionVerifyFormalNotationTestSet{
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue}), nil, nil},
		{NewOrOperator(ExpressionList{booleanTrue, booleanTrue}), nil, nil},
		{NewNotOperator(booleanTrue), nil, nil},
	}

	tests.Run(t, []string{"bucket"}, "bucket")
}

func TestBoolean(t *testing.T) {

	missingProperty := NewProperty("dne")
	null := NewLiteralNull()
	booleanTrue := NewLiteralBool(true)
	booleanFalse := NewLiteralBool(false)

	tests := ExpressionTestSet{
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue}), true, nil},
		{NewAndOperator(ExpressionList{booleanTrue, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{booleanFalse, booleanTrue}), false, nil},
		{NewAndOperator(ExpressionList{booleanFalse, booleanFalse}), false, nil},

		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanTrue}), true, nil},
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanFalse}), false, nil},

		{NewOrOperator(ExpressionList{booleanTrue, booleanTrue}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, booleanFalse}), true, nil},
		{NewOrOperator(ExpressionList{booleanFalse, booleanTrue}), true, nil},
		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse}), false, nil},

		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanTrue}), true, nil},
		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanFalse}), false, nil},

		{NewNotOperator(booleanTrue), false, nil},
		{NewNotOperator(booleanFalse), true, nil},

		// logical comparison test table from spec
		// AND
		{NewAndOperator(ExpressionList{booleanFalse, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{booleanFalse, null}), false, nil},
		{NewAndOperator(ExpressionList{booleanFalse, missingProperty}), false, nil},
		{NewAndOperator(ExpressionList{booleanFalse, booleanTrue}), false, nil},

		{NewAndOperator(ExpressionList{null, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{null, null}), nil, nil},
		{NewAndOperator(ExpressionList{null, missingProperty}), nil, &dparval.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{null, booleanTrue}), nil, nil},

		{NewAndOperator(ExpressionList{missingProperty, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{missingProperty, null}), nil, &dparval.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{missingProperty, missingProperty}), nil, &dparval.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{missingProperty, booleanTrue}), nil, &dparval.Undefined{"dne"}},

		{NewAndOperator(ExpressionList{booleanTrue, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{booleanTrue, null}), nil, nil},
		{NewAndOperator(ExpressionList{booleanTrue, missingProperty}), nil, &dparval.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue}), true, nil},

		// OR
		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse}), false, nil},
		{NewOrOperator(ExpressionList{booleanFalse, null}), nil, nil},
		{NewOrOperator(ExpressionList{booleanFalse, missingProperty}), nil, &dparval.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{booleanFalse, booleanTrue}), true, nil},

		{NewOrOperator(ExpressionList{null, booleanFalse}), nil, nil},
		{NewOrOperator(ExpressionList{null, null}), nil, nil},
		{NewOrOperator(ExpressionList{null, missingProperty}), nil, &dparval.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{null, booleanTrue}), true, nil},

		{NewOrOperator(ExpressionList{missingProperty, booleanFalse}), nil, &dparval.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{missingProperty, null}), nil, &dparval.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{missingProperty, missingProperty}), nil, &dparval.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{missingProperty, booleanTrue}), true, nil},

		{NewOrOperator(ExpressionList{booleanTrue, booleanFalse}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, null}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, missingProperty}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, booleanTrue}), true, nil},

		// NOT
		{NewNotOperator(booleanTrue), false, nil},
		{NewNotOperator(null), nil, nil},
		{NewNotOperator(missingProperty), nil, &dparval.Undefined{"dne"}},
		{NewNotOperator(booleanFalse), true, nil},
	}

	tests.Run(t)

}
