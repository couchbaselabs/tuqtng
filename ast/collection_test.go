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

	"github.com/couchbaselabs/dparval"
)

func TestCollection(t *testing.T) {

	sampleDocument := map[string]interface{}{
		"children": []interface{}{
			map[string]interface{}{
				"age": 35.0,
			},
			map[string]interface{}{
				"age": 18.0,
			},
		},
		"emptyArray": []interface{}{},
		"notAnArray": "bob",
	}

	propDNE := NewProperty("dne")
	propEmpty := NewProperty("emptyArray")
	propNot := NewProperty("notAnArray")
	propChild := NewProperty("child")
	propChildren := NewProperty("children")
	propAge := NewProperty("age")
	numberThirty := NewLiteralNumber(30.0)
	pathChildDotAge := NewDotMemberOperator(propChild, propAge)
	overThirty := NewGreaterThanOperator(pathChildDotAge, numberThirty)
	anyChild := NewCollectionAnyOperator(overThirty, propChildren, "child")
	allChild := NewCollectionAllOperator(overThirty, propChildren, "child")
	anyEmpty := NewCollectionAnyOperator(overThirty, propEmpty, "child")
	allEmpty := NewCollectionAllOperator(overThirty, propEmpty, "child")
	anyNot := NewCollectionAnyOperator(overThirty, propNot, "child")
	allNot := NewCollectionAllOperator(overThirty, propNot, "child")
	anyDNE := NewCollectionAnyOperator(overThirty, propDNE, "child")
	allDNE := NewCollectionAllOperator(overThirty, propDNE, "child")
	anyMissingInside := NewCollectionAnyOperator(propDNE, propChildren, "child")
	allMissingInside := NewCollectionAllOperator(propDNE, propChildren, "child")

	tests := ExpressionTestSet{
		{anyChild, true, nil},
		{allChild, false, nil},
		{anyEmpty, false, nil},
		{allEmpty, true, nil},
		{anyNot, false, nil},
		{allNot, false, nil},
		{anyDNE, false, nil},
		{allDNE, false, nil},
		{anyMissingInside, false, nil},
		{allMissingInside, false, nil},
	}

	item := dparval.NewValue(sampleDocument)
	tests.RunWithItem(t, item)
}

func TestCollectionStringRepresentation(t *testing.T) {

	propChild := NewProperty("child")
	propChildren := NewProperty("children")
	propAge := NewProperty("age")
	numberThirty := NewLiteralNumber(30.0)
	pathChildDotAge := NewDotMemberOperator(propChild, propAge)
	overThirty := NewGreaterThanOperator(pathChildDotAge, numberThirty)
	anyChild := NewCollectionAnyOperator(overThirty, propChildren, "child")
	allChild := NewCollectionAllOperator(overThirty, propChildren, "child")

	tests := ExpressionStringTestSet{
		{anyChild, `ANY child.age > 30 OVER children AS child`},
		{allChild, `ALL child.age > 30 OVER children AS child`},
	}

	tests.Run(t)
}

func TestCollectionValidate(t *testing.T) {

	propChild := NewProperty("child")
	propChildren := NewProperty("children")
	propAge := NewProperty("age")
	numberThirty := NewLiteralNumber(30.0)
	pathChildDotAge := NewDotMemberOperator(propChild, propAge)
	overThirty := NewGreaterThanOperator(pathChildDotAge, numberThirty)
	anyChild := NewCollectionAnyOperator(overThirty, propChildren, "child")
	allChild := NewCollectionAllOperator(overThirty, propChildren, "child")

	tests := ExpressionValidateTestSet{
		{anyChild, nil},
		{allChild, nil},
		// condition invalid
		{NewCollectionAnyOperator(notValidExpression, propChild, "child"), notValidExpressionError},
		{NewCollectionAllOperator(notValidExpression, propChild, "child"), notValidExpressionError},
		// over invalid
		{NewCollectionAnyOperator(overThirty, notValidExpression, "child"), notValidExpressionError},
		{NewCollectionAllOperator(overThirty, notValidExpression, "child"), notValidExpressionError},
	}

	tests.Run(t)
}

func TestCollectionVerifyFormalNotation(t *testing.T) {

	propChild := NewProperty("child")
	propChildren := NewProperty("children")
	propAge := NewProperty("age")
	numberThirty := NewLiteralNumber(30.0)
	pathChildDotAge := NewDotMemberOperator(propChild, propAge)
	overThirty := NewGreaterThanOperator(pathChildDotAge, numberThirty)
	anyChild := NewCollectionAnyOperator(overThirty, propChildren, "child")
	allChild := NewCollectionAllOperator(overThirty, propChildren, "child")

	tests := ExpressionVerifyFormalNotationTestSet{
		{anyChild, anyChild, nil},
		{allChild, allChild, nil},
	}

	tests.Run(t, []string{}, []string{"bucket"}, "bucket")

}
