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

	"github.com/couchbaselabs/dparval"
)

var notValidExpression = NewFunctionCall("LENGTH", FunctionArgExpressionList{})
var notValidExpressionError = fmt.Errorf("the LENGTH() function requires exactly 1 argument")
var notFormalExpression = NewProperty("property")
var _, notFormalExpressionError = notFormalExpression.VerifyFormalNotation([]string{"bucket"}, "bucket")

type ExpressionTest struct {
	input  Expression
	output interface{}
	err    error
}

type ExpressionTestSet []ExpressionTest

func (this ExpressionTestSet) Run(t *testing.T) {
	this.RunWithItem(t, nil)
}

func (this ExpressionTestSet) RunWithItem(t *testing.T, item *dparval.Value) {
	for _, x := range this {
		resval, err := x.input.Evaluate(item)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error: %v, got %v for %v", x.err, err, x.input)
		}
		var result interface{}
		if resval != nil {
			result = resval.Value()
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}
}

type ExpressionStringTest struct {
	input  fmt.Stringer
	output string
}

type ExpressionStringTestSet []ExpressionStringTest

func (this ExpressionStringTestSet) Run(t *testing.T) {
	for _, x := range this {
		result := x.input.String()
		if result != x.output {
			t.Errorf("Expected %v, got %v", x.output, result)
		}
	}
}

type ExpressionValidateTest struct {
	input  Expression
	output error
}

type ExpressionValidateTestSet []ExpressionValidateTest

func (this ExpressionValidateTestSet) Run(t *testing.T) {
	for _, x := range this {
		result := x.input.Validate()
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}
}

type ExpressionVerifyFormalNotationTest struct {
	input  Expression
	output Expression
	err    error
}

type ExpressionVerifyFormalNotationTestSet []ExpressionVerifyFormalNotationTest

func (this ExpressionVerifyFormalNotationTestSet) Run(t *testing.T, aliases []string, defaultAlias string) {
	for _, x := range this {
		result, err := x.input.VerifyFormalNotation(aliases, defaultAlias)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error: %v, got %v for %v", x.err, err, x.input)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}
}
