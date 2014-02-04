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
	"math"
	"testing"

	"github.com/couchbaselabs/dparval"
)

func TestFunction(t *testing.T) {

	sampleDocument := map[string]interface{}{
		"name": "will",
	}
	documentValue := dparval.NewValue(sampleDocument)

	sampleMeta := map[string]interface{}{
		"id": "first",
	}
	documentValue.SetAttachment("meta", sampleMeta)

	sampleContext := map[string]interface{}{
		"bucket": documentValue,
	}

	context := dparval.NewValue(sampleContext)

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
		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},

		// string functions
		{
			NewFunctionCall("LOWER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("HELLO"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("LOWER", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("LOWER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("UPPER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			"HELLO",
			nil,
		},
		{
			NewFunctionCall("UPPER", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("UPPER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			"hello     ",
			nil,
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString(" "))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			"     hello",
			nil,
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString(" "))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString(" "))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
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
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(0.0)), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(0.0)), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(27.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(0.0)), NewFunctionArgExpression(NewLiteralNumber(27.0))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.5)), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
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
			NewFunctionCall("IFMISSING", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("IFMISSINGORNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFMISSINGORNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNull())}),
			nil,
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
			NewFunctionCall("MISSINGIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{},
		},
		{
			NewFunctionCall("MISSINGIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			&dparval.Undefined{"dne"},
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
		{
			NewFunctionCall("NULLIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("NULLIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			&dparval.Undefined{"dne"},
		},

		// util functions
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("ARRAY_LENGTH", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("hello"),
						NewLiteralString("hello")}))}),
			2.0,
			nil,
		},
		{
			NewFunctionCall("OBJECT_LENGTH", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralObject(map[string]Expression{
						"val1": NewLiteralString("hello"),
						"val2": NewLiteralString("hello"),
						"val3": NewLiteralString("hello")}))}),
			3.0,
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
			nil,
		},
		// array utility functions
		{
			NewFunctionCall("ARRAY_CONCAT", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Hello world"),
						NewLiteralString("Live free or die")})),
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Enter light"),
						NewLiteralString("Exit Light")}))}),
			[]interface{}{"Hello world", "Live free or die", "Enter light", "Exit Light"},
			nil,
		},

		{
			NewFunctionCall("ARRAY_APPEND", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralString("Hello world")),
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Enter light"),
						NewLiteralString("Exit Light")}))}),
			nil,
			fmt.Errorf("the ARRAY_APPEND() function requires that first operand be of type ARRAY"),
		},

		{
			NewFunctionCall("ARRAY_APPEND", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Hello world"),
						NewLiteralString("Live free or die")})),
				NewFunctionArgExpression(
					NewLiteralString("Exit Light"))}),
			[]interface{}{"Hello world", "Live free or die", "Exit Light"},
			nil,
		},

		{
			NewFunctionCall("ARRAY_PREPEND", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Hello world"),
						NewLiteralString("Live free or die")})),
				NewFunctionArgExpression(
					NewLiteralString("Exit Light"))}),
			nil,
			fmt.Errorf("the ARRAY_PREPEND() function requires that the second operand be of type ARRAY"),
		},

		{
			NewFunctionCall("ARRAY_PREPEND", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralString("Live free or die")),
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Enter light"),
						NewLiteralString("Exit Light")}))}),
			[]interface{}{"Live free or die", "Enter light", "Exit Light"},
			nil,
		},

		// append numerics to arrays
		{
			NewFunctionCall("ARRAY_APPEND", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Hello world"),
						NewLiteralString("Live free or die")})),
				NewFunctionArgExpression(
					NewLiteralNumber(7.0))}),
			[]interface{}{"Hello world", "Live free or die", 7.0},
			nil,
		},

		{
			NewFunctionCall("ARRAY_PREPEND", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralNumber(7.0)),
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Enter light"),
						NewLiteralString("Exit Light")}))}),
			[]interface{}{7.0, "Enter light", "Exit Light"},
			nil,
		},
		// array remove functions
		{
			NewFunctionCall("ARRAY_REMOVE", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Hello world"),
						NewLiteralString("Live free or die")})),
				NewFunctionArgExpression(
					NewLiteralString("Hello world"))}),
			[]interface{}{"Live free or die"},
			nil,
		},
		{
			NewFunctionCall("ARRAY_REMOVE", FunctionArgExpressionList{
				NewFunctionArgExpression(
					NewLiteralArray(ExpressionList{
						NewLiteralString("Hello world"),
						NewLiteralString("Live free or die"),
						NewLiteralString("Hello world")})),
				NewFunctionArgExpression(
					NewLiteralString("Hello world"))}),
			[]interface{}{"Live free or die"},
			nil,
		},

		// numeric functions
		{
			NewFunctionCall("CEIL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			6.0,
			nil,
		},
		{
			NewFunctionCall("CEIL", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("CEIL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("FLOOR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("FLOOR", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("FLOOR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
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
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
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
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNumber(4))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(55.36)), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("bob")), NewFunctionArgExpression(NewLiteralNumber(4))}),
			nil,
			nil,
		},

		// numeric infinite functions
		{
			NewFunctionCall("IFNAN", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.NaN())), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNAN", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.NaN()))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNAN", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("IFPOSINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFPOSINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFPOSINF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("IFNEGINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNEGINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNEGINF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("IFINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFINF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.NaN())), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.NaN()))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IFNANORINF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.NaN())), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1))), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			5.0,
			nil,
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.NaN()))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(math.Inf(-1)))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("FIRSTNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne"))}),
			nil,
			nil,
		},

		{
			NewFunctionCall("NANIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			"hello",
			nil,
		},
		// these are commented out because NaN != NaN
		// and I haven't decided how to cope with that yet
		// {
		// 	NewFunctionCall("NANIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello"))}),
		// 	math.NaN(),
		// 	nil,
		// },
		// {
		// 	NewFunctionCall("NANIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewProperty("dne"))}),
		// 	math.NaN(),
		// 	nil,
		// },
		{
			NewFunctionCall("NANIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			&dparval.Undefined{"dne"},
		},

		{
			NewFunctionCall("POSINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("POSINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello"))}),
			math.Inf(1),
			nil,
		},
		{
			NewFunctionCall("POSINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewProperty("dne"))}),
			math.Inf(1),
			nil,
		},
		{
			NewFunctionCall("POSINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
		{
			NewFunctionCall("NEGINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("NEGINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello"))}),
			math.Inf(-1),
			nil,
		},
		{
			NewFunctionCall("NEGINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewProperty("dne"))}),
			math.Inf(-1),
			nil,
		},
		{
			NewFunctionCall("NEGINFIF", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralString("bob"))}),
			nil,
			&dparval.Undefined{"dne"},
		},
	}

	tests.RunWithItem(t, context)

}

func TestFunctionStringRepresentation(t *testing.T) {

	tests := ExpressionStringTestSet{
		{NewFunctionCall("LOWER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("HELLO"))}),
			`LOWER("HELLO")`,
		},
	}

	tests.Run(t)
}

func TestFunctionArgExpressionListStringRepresentation(t *testing.T) {
	arglist := FunctionArgExpressionList{
		NewFunctionArgExpression(NewProperty("a")),
		NewFunctionArgExpression(NewProperty("b")),
		NewStarFunctionArgExpression(),
		NewDotStarFunctionArgExpression(NewProperty("c")),
	}

	arglistString := arglist.String()
	if arglistString != "a, b, *, c.*" {
		t.Errorf("Expected arg list `a, b, *, c.*`, got: %v", arglistString)
	}
}

func TestFunctionValidate(t *testing.T) {

	tests := ExpressionValidateTestSet{
		// meta/value functions
		{
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			nil,
		},
		{
			NewFunctionCall("META", FunctionArgExpressionList{}),
			fmt.Errorf("the META() function requires exactly 1 argument"),
		},

		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			nil,
		},
		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{}),
			fmt.Errorf("the VALUE() function requires exactly 1 argument"),
		},

		// string functions
		{
			NewFunctionCall("LOWER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("HELLO"))}),
			nil,
		},
		{
			NewFunctionCall("LOWER", FunctionArgExpressionList{}),
			fmt.Errorf("the LOWER() function requires exactly 1 argument"),
		},
		{
			NewFunctionCall("UPPER", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("UPPER", FunctionArgExpressionList{}),
			fmt.Errorf("the UPPER() function requires exactly 1 argument"),
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			nil,
		},
		{
			NewFunctionCall("LTRIM", FunctionArgExpressionList{}),
			fmt.Errorf("the LTRIM() function requires exactly 2 arguments"),
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			nil,
		},
		{
			NewFunctionCall("RTRIM", FunctionArgExpressionList{}),
			fmt.Errorf("the RTRIM() function requires exactly 2 arguments"),
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("     hello     ")), NewFunctionArgExpression(NewLiteralString(" "))}),
			nil,
		},
		{
			NewFunctionCall("TRIM", FunctionArgExpressionList{}),
			fmt.Errorf("the TRIM() function requires exactly 2 arguments"),
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			nil,
		},
		{
			NewFunctionCall("SUBSTR", FunctionArgExpressionList{}),
			fmt.Errorf("the SUBSTR() function requires at least 2 arguments"),
		},

		// comparison functions
		{
			NewFunctionCall("GREATEST", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("GREATEST", FunctionArgExpressionList{}),
			fmt.Errorf("the GREATEST() function requires at least 1 argument"),
		},
		{
			NewFunctionCall("LEAST", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("LEAST", FunctionArgExpressionList{}),
			fmt.Errorf("the LEAST() function requires at least 1 argument"),
		},
		{
			NewFunctionCall("IFMISSING", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("IFMISSING", FunctionArgExpressionList{}),
			fmt.Errorf("the IFMISSING() function requires at least 1 argument"),
		},
		{
			NewFunctionCall("IFNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("IFNULL", FunctionArgExpressionList{}),
			fmt.Errorf("the IFNULL() function requires at least 1 argument"),
		},
		{
			NewFunctionCall("IFMISSINGORNULL", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("dne")), NewFunctionArgExpression(NewLiteralNull()), NewFunctionArgExpression(NewLiteralNumber(5.0)), NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("IFMISSINGORNULL", FunctionArgExpressionList{}),
			fmt.Errorf("the IFMISSINGORNULL() function requires at least 1 argument"),
		},

		{
			NewFunctionCall("MISSINGIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			nil,
		},
		{
			NewFunctionCall("MISSINGIF", FunctionArgExpressionList{}),
			fmt.Errorf("the MISSINGIF() function requires exactly 2 arguments"),
		},
		{
			NewFunctionCall("NULLIF", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello")), NewFunctionArgExpression(NewLiteralString("hello2"))}),
			nil,
		},
		{
			NewFunctionCall("NULLIF", FunctionArgExpressionList{}),
			fmt.Errorf("the NULLIF() function requires exactly 2 arguments"),
		},

		// util functions
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{}),
			fmt.Errorf("the LENGTH() function requires exactly 1 argument"),
		},

		// numeric functions
		{
			NewFunctionCall("CEIL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			nil,
		},
		{
			NewFunctionCall("CEIL", FunctionArgExpressionList{}),
			fmt.Errorf("the CEIL() function requires exactly 1 argument"),
		},
		{
			NewFunctionCall("FLOOR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			nil,
		},
		{
			NewFunctionCall("FLOOR", FunctionArgExpressionList{}),
			fmt.Errorf("the FLOOR() function requires exactly 1 argument"),
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			nil,
		},
		{
			NewFunctionCall("ROUND", FunctionArgExpressionList{}),
			fmt.Errorf("the ROUND() function requires at least 1 argument"),
		},

		//trunc
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.8))}),
			nil,
		},
		{
			NewFunctionCall("TRUNC", FunctionArgExpressionList{}),
			fmt.Errorf("the TRUNC() function requires at least 1 argument"),
		},

		// non-existant function
		{
			NewFunctionCall("DOESNOTEXIST", FunctionArgExpressionList{}),
			fmt.Errorf("no system function named DOESNOTEXIST registered"),
		},
	}

	tests.Run(t)
}

func TestFunctionVerifyFormalNotation(t *testing.T) {

	tests := ExpressionVerifyFormalNotationTestSet{
		{
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			nil,
		},
		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			NewFunctionCall("VALUE", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(notFormalExpression)}),
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(notFormalExpression)}),
			notFormalExpressionError,
		},
	}

	tests.Run(t, []string{}, []string{"bucket", "child"}, "")

	// single alias tests

	tests = ExpressionVerifyFormalNotationTestSet{
		{
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("notbucket"))}),
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewDotMemberOperator(NewProperty("bucket"), NewProperty("notbucket")))}),
			fmt.Errorf("invalid argument to META() function, must be bucket name/alias"),
		},
		{
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.4))}),
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(5.4))}),
			fmt.Errorf("invalid argument to META() function, must be bucket name/alias"),
		},
		{
			NewFunctionCall("META", FunctionArgExpressionList{}),
			NewFunctionCall("META", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			nil,
		},
		{
			NewFunctionCall("VALUE", FunctionArgExpressionList{}),
			NewFunctionCall("VALUE", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bucket"))}),
			nil,
		},
		{
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("bob"))}),
			NewFunctionCall("LENGTH", FunctionArgExpressionList{NewFunctionArgExpression(NewDotMemberOperator(NewProperty("bucket"), NewProperty("bob")))}),
			nil,
		},
	}

	tests.Run(t, []string{}, []string{"bucket"}, "bucket")

}
