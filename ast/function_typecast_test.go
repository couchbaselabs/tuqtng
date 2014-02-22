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
)

func TestFunctionToNum(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			1.0,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			1.5,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			1.0,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			0.0,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			nil,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionToStr(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			"1",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			"1.5",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			"true",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			"false",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			"[]",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNull(),
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			"[null, 0, \"hello\"]",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			"{}",
			nil,
		},
		{
			NewFunctionCall("TO_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			"{\"zero\": 0, \"hello\": \"hello\"}",
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionToBool(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString(""))}),
			false,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			true,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("TO_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			true,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionToAtom(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			1.0,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			"hello",
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			true,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			0.0,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralObject(
						map[string]Expression{
							"zero": NewLiteralNumber(0),
						})}))}),
			0.0,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			0.0,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralArray(
						ExpressionList{
							NewLiteralNumber(0),
						})}))}),
			0.0,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionToArray(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			[]interface{}{1.0},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			[]interface{}{"hello"},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			[]interface{}{true},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			[]interface{}{false},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			[]interface{}{},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			[]interface{}{0.0},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			[]interface{}{0.0, "hello"},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			[]interface{}{map[string]interface{}{}},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			[]interface{}{map[string]interface{}{"zero": 0.0}},
			nil,
		},
		{
			NewFunctionCall("TO_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			[]interface{}{map[string]interface{}{"zero": 0.0, "hello": "hello"}},
			nil,
		},
	}

	tests.Run(t)
}
