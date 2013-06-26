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

func TestArithmetic(t *testing.T) {

	numberSix := NewLiteralNumber(6.0)
	numberSeven := NewLiteralNumber(7.0)
	numberNegativeSeven := NewLiteralNumber(-7.0)
	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")

	tests := []struct {
		input  Expression
		output interface{}
	}{
		{NewPlusOperator(numberSeven, numberSeven), 14.0},
		{NewPlusOperator(stringCouchbase, stringServer), nil}, // no longer support string concatenation, uses different operator
		{NewPlusOperator(numberSeven, stringCouchbase), nil},
		{NewSubtractOperator(numberSeven, numberSeven), 0.0},
		{NewSubtractOperator(numberSeven, stringCouchbase), nil},
		{NewMultiplyOperator(numberSeven, numberSeven), 49.0},
		{NewMultiplyOperator(numberSeven, stringCouchbase), nil},
		{NewDivideOperator(numberSeven, numberSeven), 1.0},
		{NewDivideOperator(numberSeven, stringCouchbase), nil},
		{NewModuloOperator(numberSeven, numberSix), 1.0},
		{NewChangeSignOperator(numberSeven), -7.0},
		{NewChangeSignOperator(numberNegativeSeven), 7.0},
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
