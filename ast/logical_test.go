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
	"fmt"
	"reflect"
	"testing"

	"github.com/couchbaselabs/tuqtng/query"
)

func TestBooleanStringRepresentation(t *testing.T) {

	booleanTrue := NewLiteralBool(true)
	booleanFalse := NewLiteralBool(false)

	tests := []struct {
		input  fmt.Stringer
		output string
	}{
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

	for _, x := range tests {
		result := x.input.String()
		if result != x.output {
			t.Errorf("Expected %v, got %v", x.output, result)
		}
	}

}

func TestBoolean(t *testing.T) {

	missingProperty := NewProperty("dne")
	null := NewLiteralNull()
	booleanTrue := NewLiteralBool(true)
	booleanFalse := NewLiteralBool(false)

	tests := []struct {
		input  Expression
		output query.Value
		err    error
	}{
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
		{NewAndOperator(ExpressionList{null, missingProperty}), nil, &query.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{null, booleanTrue}), nil, nil},

		{NewAndOperator(ExpressionList{missingProperty, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{missingProperty, null}), nil, &query.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{missingProperty, missingProperty}), nil, &query.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{missingProperty, booleanTrue}), nil, &query.Undefined{"dne"}},

		{NewAndOperator(ExpressionList{booleanTrue, booleanFalse}), false, nil},
		{NewAndOperator(ExpressionList{booleanTrue, null}), nil, nil},
		{NewAndOperator(ExpressionList{booleanTrue, missingProperty}), nil, &query.Undefined{"dne"}},
		{NewAndOperator(ExpressionList{booleanTrue, booleanTrue}), true, nil},

		// OR
		{NewOrOperator(ExpressionList{booleanFalse, booleanFalse}), false, nil},
		{NewOrOperator(ExpressionList{booleanFalse, null}), nil, nil},
		{NewOrOperator(ExpressionList{booleanFalse, missingProperty}), nil, &query.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{booleanFalse, booleanTrue}), true, nil},

		{NewOrOperator(ExpressionList{null, booleanFalse}), nil, nil},
		{NewOrOperator(ExpressionList{null, null}), nil, nil},
		{NewOrOperator(ExpressionList{null, missingProperty}), nil, &query.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{null, booleanTrue}), true, nil},

		{NewOrOperator(ExpressionList{missingProperty, booleanFalse}), nil, &query.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{missingProperty, null}), nil, &query.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{missingProperty, missingProperty}), nil, &query.Undefined{"dne"}},
		{NewOrOperator(ExpressionList{missingProperty, booleanTrue}), true, nil},

		{NewOrOperator(ExpressionList{booleanTrue, booleanFalse}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, null}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, missingProperty}), true, nil},
		{NewOrOperator(ExpressionList{booleanTrue, booleanTrue}), true, nil},

		// NOT
		{NewNotOperator(booleanTrue), false, nil},
		{NewNotOperator(null), nil, nil},
		{NewNotOperator(missingProperty), nil, &query.Undefined{"dne"}},
		{NewNotOperator(booleanFalse), true, nil},
	}

	for _, x := range tests {
		result, err := x.input.Evaluate(nil)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error %v, got %v for %v", x.err, err, x.input)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}

}
