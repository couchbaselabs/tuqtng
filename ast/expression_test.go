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
var notFormalExpressionAfter = NewDotMemberOperator(NewProperty("bucket"), NewProperty("property"))
var notFormalExpressionError = fmt.Errorf("Property reference property missing qualifier bucket/alias")

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
		validator := NewExpressionValidator()
		_, result := x.input.Accept(validator)
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

func (this ExpressionVerifyFormalNotationTestSet) Run(t *testing.T, forbiddenAliases []string, aliases []string, defaultAlias string) {
	for _, x := range this {
		formalNotation := NewExpressionFormalNotationConverter(forbiddenAliases, aliases, defaultAlias)
		result, err := x.input.Accept(formalNotation)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error: %v, got %v for %v", x.err, err, x.input)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}
}

// this type exists only so we can test that
// expressions properly return any internal errors
// that may occur within the evaluation process

func newInternalErrorExpression() *internalErrorExpression {
	return &internalErrorExpression{}
}

type internalErrorExpression struct {
}

var internalError = fmt.Errorf("Internal Error")

func (this *internalErrorExpression) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	return nil, fmt.Errorf("Internal Error")
}

func (this *internalErrorExpression) Validate() error {
	return nil
}

func (this *internalErrorExpression) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return nil, nil
}

func (this *internalErrorExpression) String() string {
	return fmt.Sprintf("NOT_A_REAL_EXPRESSION")
}

func (this *internalErrorExpression) EquivalentTo(t Expression) bool {
	_, ok := t.(*internalErrorExpression)
	return ok
}

func (this *internalErrorExpression) Dependencies() ExpressionList {
	rv := ExpressionList{}
	return rv
}

func (this *internalErrorExpression) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
