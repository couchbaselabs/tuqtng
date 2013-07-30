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

func TestLiteralStringRepresentation(t *testing.T) {
	tests := []struct {
		input  fmt.Stringer
		output string
	}{
		{NewLiteralNull(), "null"},
		{NewLiteralBool(true), "true"},
		{NewLiteralBool(false), "false"},
		{NewLiteralNumber(1.0), "1"},
		{NewLiteralNumber(3.14), "3.14"},
		{NewLiteralString("couchbase"), "\"couchbase\""},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0)}), "[1]"},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewLiteralBool(false)}), "[1, false]"},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewLiteralBool(false), NewLiteralString("bob")}), "[1, false, \"bob\"]"},
		{NewLiteralObject(map[string]Expression{"name": NewLiteralString("bob")}), "{\"name\": \"bob\"}"},
		{NewLiteralObject(map[string]Expression{"user": NewLiteralString("test"), "age": NewLiteralNumber(27.0)}), "{\"user\": \"test\", \"age\": 27}"},
	}

	for _, x := range tests {
		result := x.input.String()
		if result != x.output {
			t.Errorf("Expected %v, got %v", x.output, result)
		}
	}

}

func TestEvaluateLiteral(t *testing.T) {
	tests := []struct {
		input  Expression
		output query.Value
	}{
		{NewLiteralNull(), nil},
		{NewLiteralBool(true), true},
		{NewLiteralBool(false), false},
		{NewLiteralNumber(1.0), 1.0},
		{NewLiteralNumber(3.14), 3.14},
		{NewLiteralString("couchbase"), "couchbase"},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0)}), []query.Value{1.0}},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewLiteralBool(false)}), []query.Value{1.0, false}},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewLiteralBool(false), NewLiteralString("bob")}), []query.Value{1.0, false, "bob"}},
		{NewLiteralObject(map[string]Expression{"name": NewLiteralString("bob")}), map[string]query.Value{"name": "bob"}},
		{NewLiteralObject(map[string]Expression{"user": NewLiteralString("test"), "age": NewLiteralNumber(27.0)}), map[string]query.Value{"age": 27.0, "user": "test"}},
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

func TestEvaluateComplexLiteralContainingMissing(t *testing.T) {

	sampleDocument := map[string]query.Value{}
	item := query.NewParsedItem(sampleDocument, nil)

	tests := []struct {
		input  Expression
		output query.Value
	}{
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewProperty("bob")}), []query.Value{1.0}},
		{NewLiteralObject(map[string]Expression{"name": NewLiteralString("bob"), "cat": NewProperty("bob")}), map[string]query.Value{"name": "bob"}},
	}

	for _, x := range tests {
		result, err := x.input.Evaluate(item)
		if err != nil {
			t.Fatalf("Error evaluating expression: %v", err)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %t %v, got %t %v", x.output, x.output, result, result)
		}
	}

}
