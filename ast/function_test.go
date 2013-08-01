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

	"github.com/mschoch/dparval"
)

func TestFunction(t *testing.T) {

	sampleContext := map[string]interface{}{
		"bucket": map[string]interface{}{
			"name": "will",
		},
	}
	sampleMeta := map[string]interface{}{
		"id": "first",
	}

	context := dparval.NewValue(sampleContext)
	context.AddMeta("meta", sampleMeta)

	tests := ExpressionTestSet{
		// meta/value functions
		{
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			map[string]interface{}{
				"id": "first",
			},
			nil,
		},

		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			map[string]interface{}{
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
			&dparval.Undefined{},
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
			5.0,
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("hello"),
						NewLiteralString("hello")}))}),
			2.0,
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralObject(map[string]Expression{
						"val1": NewLiteralString("hello"),
						"val2": NewLiteralString("hello"),
						"val3": NewLiteralString("hello")}))}),
			3.0,
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

	tests.RunWithItem(t, context)

}
