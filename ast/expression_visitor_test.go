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
)

func TestExpressionVisitor(t *testing.T) {
	validator := &ExpressionValidator{}
	plus := NewPlusOperator(NewLiteralNumber(3.5), NewLiteralNumber(3.5))
	plus.Accept(validator)
}

func TestFunctionalDependencies(t *testing.T) {
	testNumberA := NewLiteralNumber(5.5)
	testPropertyA := NewProperty("test")
	testPropertyB := NewProperty("dog")
	testPlus := NewPlusOperator(testPropertyA, testPropertyB)

	childrenProperty := NewProperty("children")
	childProperty := NewProperty("child")
	ageProperty := NewProperty("age")
	childDotAge := NewDotMemberOperator(childProperty, ageProperty)
	condition := NewGreaterThanOperator(childDotAge, testNumberA)
	anyChildOver5Point5 := NewCollectionAnyOperator(condition, childrenProperty, "child")
	tests := []struct {
		input Expression
		deps  ExpressionList
		err   error
	}{
		// literals have no dependencies
		{testNumberA, ExpressionList{}, nil},
		// property has itself as dependency
		{testPropertyA, ExpressionList{testPropertyA}, nil},
		// property with nothing
		{testPropertyA, ExpressionList{}, fmt.Errorf("The expression test is not satisfied by these dependencies")},
		// addition can be satisfied by itself
		{testPlus, ExpressionList{testPlus}, nil},
		// or by both its arguments
		{testPlus, ExpressionList{testPropertyA, testPropertyB}, nil},
		// but NOT by only one of its arguments
		{testPlus, ExpressionList{testPropertyA}, fmt.Errorf("The expression dog is not satisfied by these dependencies")},
		// this should pass but will fail
		{anyChildOver5Point5, ExpressionList{childrenProperty}, nil},
	}

	for _, test := range tests {
		fdc := NewExpressionFunctionalDependencyChecker(test.deps)
		_, err := test.input.Accept(fdc)
		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("Expected error %v, got %v", test.err, err)
		}
	}
}
