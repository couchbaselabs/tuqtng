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

func TestExpressionCNF(t *testing.T) {

	pa := NewProperty("a")
	pb := NewProperty("b")
	pc := NewProperty("c")

	tests := []struct {
		input  Expression
		output Expression
	}{
		// reduce complex AND into simple AND
		{
			NewAndOperator(ExpressionList{NewAndOperator(ExpressionList{pa, pb}), pc}),
			NewAndOperator(ExpressionList{pa, pb, pc}),
		},
		// reduce complex OR into simple OR
		{
			NewOrOperator(ExpressionList{NewOrOperator(ExpressionList{pa, pb}), pc}),
			NewOrOperator(ExpressionList{pa, pb, pc}),
		},
		// distribute OR across nested ANDs
		{
			NewOrOperator(ExpressionList{NewAndOperator(ExpressionList{pa, pb}), pc}),
			NewAndOperator(ExpressionList{NewOrOperator(ExpressionList{pc, pa}), NewOrOperator(ExpressionList{pc, pb})}),
		},
	}

	ecnf := NewExpressionCNF()
	for _, test := range tests {
		actual, err := test.input.Accept(ecnf)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(test.output, actual) {
			t.Errorf("Expected %v, got %v, for %v", test.output, actual, test.input)
		}
	}

}
