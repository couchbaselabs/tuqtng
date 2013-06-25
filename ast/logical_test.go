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
)

func TestBooleanStringRepresentation(t *testing.T) {

	booleanTrue := NewLiteralBool(true)
	booleanFalse := NewLiteralBool(false)

	tests := []struct {
		input  fmt.Stringer
		output string
	}{
		{NewAndOperator([]Expression{booleanTrue, booleanTrue}), "AND [true true]"},
		{NewAndOperator([]Expression{booleanTrue, booleanFalse}), "AND [true false]"},
		{NewAndOperator([]Expression{booleanFalse, booleanTrue}), "AND [false true]"},
		{NewAndOperator([]Expression{booleanFalse, booleanFalse}), "AND [false false]"},

		{NewAndOperator([]Expression{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanTrue}), "AND [true true true true true]"},
		{NewAndOperator([]Expression{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanFalse}), "AND [true true true true false]"},

		{NewOrOperator([]Expression{booleanTrue, booleanTrue}), "OR [true true]"},
		{NewOrOperator([]Expression{booleanTrue, booleanFalse}), "OR [true false]"},
		{NewOrOperator([]Expression{booleanFalse, booleanTrue}), "OR [false true]"},
		{NewOrOperator([]Expression{booleanFalse, booleanFalse}), "OR [false false]"},

		{NewOrOperator([]Expression{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanTrue}), "OR [false false false false true]"},
		{NewOrOperator([]Expression{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanFalse}), "OR [false false false false false]"},

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
		output Value
		err    error
	}{
		{NewAndOperator([]Expression{booleanTrue, booleanTrue}), true, nil},
		{NewAndOperator([]Expression{booleanTrue, booleanFalse}), false, nil},
		{NewAndOperator([]Expression{booleanFalse, booleanTrue}), false, nil},
		{NewAndOperator([]Expression{booleanFalse, booleanFalse}), false, nil},

		{NewAndOperator([]Expression{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanTrue}), true, nil},
		{NewAndOperator([]Expression{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanFalse}), false, nil},

		{NewOrOperator([]Expression{booleanTrue, booleanTrue}), true, nil},
		{NewOrOperator([]Expression{booleanTrue, booleanFalse}), true, nil},
		{NewOrOperator([]Expression{booleanFalse, booleanTrue}), true, nil},
		{NewOrOperator([]Expression{booleanFalse, booleanFalse}), false, nil},

		{NewOrOperator([]Expression{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanTrue}), true, nil},
		{NewOrOperator([]Expression{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanFalse}), false, nil},

		{NewNotOperator(booleanTrue), false, nil},
		{NewNotOperator(booleanFalse), true, nil},

		// logical comparison test table from spec
		// AND
		{NewAndOperator([]Expression{booleanFalse, booleanFalse}), false, nil},
		{NewAndOperator([]Expression{booleanFalse, null}), false, nil},
		{NewAndOperator([]Expression{booleanFalse, missingProperty}), false, nil},
		{NewAndOperator([]Expression{booleanFalse, booleanTrue}), false, nil},

		{NewAndOperator([]Expression{null, booleanFalse}), false, nil},
		{NewAndOperator([]Expression{null, null}), nil, nil},
		{NewAndOperator([]Expression{null, missingProperty}), nil, &Undefined{"dne"}},
		{NewAndOperator([]Expression{null, booleanTrue}), nil, nil},

		{NewAndOperator([]Expression{missingProperty, booleanFalse}), false, nil},
		{NewAndOperator([]Expression{missingProperty, null}), nil, &Undefined{"dne"}},
		{NewAndOperator([]Expression{missingProperty, missingProperty}), nil, &Undefined{"dne"}},
		{NewAndOperator([]Expression{missingProperty, booleanTrue}), nil, &Undefined{"dne"}},

		{NewAndOperator([]Expression{booleanTrue, booleanFalse}), false, nil},
		{NewAndOperator([]Expression{booleanTrue, null}), nil, nil},
		{NewAndOperator([]Expression{booleanTrue, missingProperty}), nil, &Undefined{"dne"}},
		{NewAndOperator([]Expression{booleanTrue, booleanTrue}), true, nil},

		// OR
		{NewOrOperator([]Expression{booleanFalse, booleanFalse}), false, nil},
		{NewOrOperator([]Expression{booleanFalse, null}), nil, nil},
		{NewOrOperator([]Expression{booleanFalse, missingProperty}), nil, &Undefined{"dne"}},
		{NewOrOperator([]Expression{booleanFalse, booleanTrue}), true, nil},

		{NewOrOperator([]Expression{null, booleanFalse}), nil, nil},
		{NewOrOperator([]Expression{null, null}), nil, nil},
		{NewOrOperator([]Expression{null, missingProperty}), nil, &Undefined{"dne"}},
		{NewOrOperator([]Expression{null, booleanTrue}), true, nil},

		{NewOrOperator([]Expression{missingProperty, booleanFalse}), nil, &Undefined{"dne"}},
		{NewOrOperator([]Expression{missingProperty, null}), nil, &Undefined{"dne"}},
		{NewOrOperator([]Expression{missingProperty, missingProperty}), nil, &Undefined{"dne"}},
		{NewOrOperator([]Expression{missingProperty, booleanTrue}), true, nil},

		{NewOrOperator([]Expression{booleanTrue, booleanFalse}), true, nil},
		{NewOrOperator([]Expression{booleanTrue, null}), true, nil},
		{NewOrOperator([]Expression{booleanTrue, missingProperty}), true, nil},
		{NewOrOperator([]Expression{booleanTrue, booleanTrue}), true, nil},

		// NOT
		{NewNotOperator(booleanTrue), false, nil},
		{NewNotOperator(null), nil, nil},
		{NewNotOperator(missingProperty), nil, &Undefined{"dne"}},
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
