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

func TestFunctionIsNum(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(-1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISNUM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
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
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString(""))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNull(),
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISSTR", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
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
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(0.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("1.5"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString(""))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISBOOL", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
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
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello")}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISATOM", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
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
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello")}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISARRAY", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
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
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNumber(1.0))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralString("hello"))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(true))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralBool(false))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralNull())}),
			nil,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(ExpressionList{}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralArray(
				ExpressionList{
					NewLiteralNumber(0),
					NewLiteralString("hello"),
				}))}),
			false,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(map[string]Expression{}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
				map[string]Expression{
					"zero": NewLiteralNumber(0),
				}))}),
			true,
			nil,
		},
		{
			NewFunctionCall("ISOBJ", FunctionArgExpressionList{NewFunctionArgExpression(NewLiteralObject(
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
