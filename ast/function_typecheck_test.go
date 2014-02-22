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

func TestFunctionTypeName(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			"number",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			"number",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(-1.0))}),
			"number",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			"string",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			"string",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			"boolean",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			"boolean",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			"null",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			"array",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			"array",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			"object",
			nil,
		},
		{
			NewFunctionCall("TYPE_NAME", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			"object",
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionIsNum(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(-1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_NUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionIsStr(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString(""))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNull(),
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_STR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionIsBool(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString(""))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_BOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionIsAtom(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello")}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionIsArray(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello")}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_ARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
	}

	tests.Run(t)
}

func TestFunctionIsObj(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("IS_OBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero":  NewLiteralNumber(0),
					"hello": NewLiteralString("hello"),
				}))}),
			true,
			nil,
		},
	}

	tests.Run(t)
}
