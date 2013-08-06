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

func TestAssignDefaultNames(t *testing.T) {
	tests := []struct {
		input  ResultExpressionList
		output ResultExpressionList
	}{
		// expressions without explicit alias get one assigned
		{
			ResultExpressionList{NewResultExpression(NewLiteralBool(true))},
			ResultExpressionList{NewResultExpressionWithAlias(NewLiteralBool(true), "$1")},
		},
		// simple path expressions get their own name as the alias
		{
			ResultExpressionList{NewResultExpression(NewProperty("path"))},
			ResultExpressionList{NewResultExpressionWithAlias(NewProperty("path"), "path")},
		},
		// explicitly named simple path expressions should get the name requested
		{
			ResultExpressionList{NewResultExpressionWithAlias(NewProperty("path"), "custom")},
			ResultExpressionList{NewResultExpressionWithAlias(NewProperty("path"), "custom")},
		},
		{
			ResultExpressionList{NewStarResultExpression(), NewResultExpression(NewLiteralBool(true))},
			ResultExpressionList{NewStarResultExpression(), NewResultExpressionWithAlias(NewLiteralBool(true), "$1")},
		},
	}

	for _, x := range tests {
		err := x.input.AssignDefaultNames([]string{})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(x.input, x.output) {
			t.Errorf("Expected %v, got %v", x.output, x.input)
		}
	}

}
