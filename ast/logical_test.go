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

	booleanTrue := NewLiteralBool(true)
	booleanFalse := NewLiteralBool(false)

	tests := []struct {
		input  Expression
		output Value
	}{
		{NewAndOperator([]Expression{booleanTrue, booleanTrue}), true},
		{NewAndOperator([]Expression{booleanTrue, booleanFalse}), false},
		{NewAndOperator([]Expression{booleanFalse, booleanTrue}), false},
		{NewAndOperator([]Expression{booleanFalse, booleanFalse}), false},

		{NewAndOperator([]Expression{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanTrue}), true},
		{NewAndOperator([]Expression{booleanTrue, booleanTrue, booleanTrue, booleanTrue, booleanFalse}), false},

		{NewOrOperator([]Expression{booleanTrue, booleanTrue}), true},
		{NewOrOperator([]Expression{booleanTrue, booleanFalse}), true},
		{NewOrOperator([]Expression{booleanFalse, booleanTrue}), true},
		{NewOrOperator([]Expression{booleanFalse, booleanFalse}), false},

		{NewOrOperator([]Expression{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanTrue}), true},
		{NewOrOperator([]Expression{booleanFalse, booleanFalse, booleanFalse, booleanFalse, booleanFalse}), false},

		{NewNotOperator(booleanTrue), false},
		{NewNotOperator(booleanFalse), true},
	}

	for _, x := range tests {
		result, err := x.input.Evaluate(nil)
		if err != nil {
			t.Fatalf("Error evaluating expression: %v", err)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %t %v, got %t %v", x.output, x.output, result, result)
		}
	}

}
