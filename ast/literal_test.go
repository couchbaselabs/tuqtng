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
	"testing"

	"github.com/mschoch/dparval"
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
	tests := ExpressionTestSet{
		{NewLiteralNull(), nil, nil},
		{NewLiteralBool(true), true, nil},
		{NewLiteralBool(false), false, nil},
		{NewLiteralNumber(1.0), 1.0, nil},
		{NewLiteralNumber(3.14), 3.14, nil},
		{NewLiteralString("couchbase"), "couchbase", nil},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0)}), []interface{}{1.0}, nil},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewLiteralBool(false)}), []interface{}{1.0, false}, nil},
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewLiteralBool(false), NewLiteralString("bob")}), []interface{}{1.0, false, "bob"}, nil},
		{NewLiteralObject(map[string]Expression{"name": NewLiteralString("bob")}), map[string]interface{}{"name": "bob"}, nil},
		{NewLiteralObject(map[string]Expression{"user": NewLiteralString("test"), "age": NewLiteralNumber(27.0)}), map[string]interface{}{"age": 27.0, "user": "test"}, nil},
	}

	tests.Run(t)

}

func TestEvaluateComplexLiteralContainingMissing(t *testing.T) {

	sampleDocument := map[string]interface{}{}
	item := dparval.NewValue(sampleDocument)

	tests := ExpressionTestSet{
		{NewLiteralArray(ExpressionList{NewLiteralNumber(1.0), NewProperty("bob")}), []interface{}{1.0}, nil},
		{NewLiteralObject(map[string]Expression{"name": NewLiteralString("bob"), "cat": NewProperty("bob")}), map[string]interface{}{"name": "bob"}, nil},
	}

	tests.RunWithItem(t, item)

}
