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

	"github.com/couchbaselabs/tuqtng/query"
)

func TestFunction(t *testing.T) {

	sampleContext := map[string]query.Value{
		"bucket": map[string]query.Value{
			"name": "will",
		},
	}
	sampleMeta := map[string]query.Value{
		"id": "first",
	}

	context := query.NewParsedItem(sampleContext, sampleMeta)

	tests := []struct {
		input  Expression
		output query.Value
		err    error
	}{
		// meta/value functions
		{
			NewFunctionCall("META", FunctionArgExpressionList{}),
			map[string]query.Value{
				"id": "first",
			},
			nil,
		},

		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{}),
			map[string]query.Value{
				"name": "will",
			},
			nil,
		},

		// string functions
		{
			NewFunctionCall("LOWER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("HELLO"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("UPPER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			"HELLO",
			nil,
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			"hello     ",
			nil,
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			"     hello",
			nil,
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			"ello",
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(1.0)), NewFunctionArgExpression(NewLiteralNumber(2.0))}),
			"el",
			nil,
		},

		// comparison functions
		{
			NewFunctionCall("GREATEST", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("GREATEST", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("LEAST", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("LEAST", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFMISSING", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFMISSINGORNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},

		{
			NewFunctionCall("MISSINGIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("MISSINGIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
			&query.Undefined{},
		},
		{
			NewFunctionCall("NULLIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("NULLIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
			nil,
		},

		// util functions
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			5,
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray([]Expression{
						NewLiteralString("hello"),
						NewLiteralString("hello")}))}),
			2,
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralObject(map[string]Expression{
						"val1": NewLiteralString("hello"),
						"val2": NewLiteralString("hello"),
						"val3": NewLiteralString("hello")}))}),
			3,
			nil,
		},

		// numeric functions
		{
			NewFunctionCall("CEIL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			6.0,
			nil,
		},
		{
			NewFunctionCall("FLOOR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			6.0,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36))}),
			55.0,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(0))}),
			55.0,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(1))}),
			55.4,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(2))}),
			55.36,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(3))}),
			55.36,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(4))}),
			55.36,
			nil,
		},

		//trunc
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36))}),
			55.0,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(0))}),
			55.0,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(1))}),
			55.3,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(2))}),
			55.36,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(3))}),
			55.36,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralNumber(4))}),
			55.36,
			nil,
		},
	}

	for _, x := range tests {
		x.input.VerifyFormalNotation([]string{"bucket"}, "bucket")
		x.input.Validate()
		result, err := x.input.Evaluate(context)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error %v, got %v for %v", x.err, err, x.input)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}

}
