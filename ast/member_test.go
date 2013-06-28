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
)

func TestDotMember(t *testing.T) {

	sampleContext := map[string]Value{
		"address": map[string]Value{
			"street": "1 recursive function",
		},
		"contact": map[string]Value{
			"name": map[string]Value{
				"first": "unql",
				"last":  "couchbase",
				"all": []Value{
					"unql",
					"couchbase",
				},
			},
		},
		"friends": []Value{
			"a",
			"b",
			"c",
		},
		"name": "bob",
	}
	sampleMeta := map[string]Value{
		"id": "first",
	}

	tests := []struct {
		input  Expression
		output Value
		err    error
	}{
		{NewDotMemberOperator(NewProperty("address"), NewProperty("street")), "1 recursive function", nil},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("first")), "unql", nil},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("last")), "couchbase", nil},

		{NewDotMemberOperator(NewProperty("address"), NewProperty("city")), nil, &Undefined{"city"}},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("middle")), nil, &Undefined{"middle"}},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("namez")), NewProperty("first")), nil, &Undefined{"namez"}},

		{NewDotMemberOperator(NewProperty("name"), NewProperty("city")), nil, &Undefined{"city"}},

		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(0.0)), "a", nil},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(1.0)), "b", nil},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(2.0)), "c", nil},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(-1.0)), nil, &Undefined{}},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(10.0)), nil, &Undefined{}},

		{NewBracketMemberOperator(NewProperty("foo"), NewLiteralNumber(10.0)), nil, &Undefined{"foo"}},
		{NewBracketMemberOperator(NewProperty("friends"), NewProperty("bar")), nil, &Undefined{"bar"}},

		//compound test
		{NewBracketMemberOperator(NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("all")), NewLiteralNumber(0.0)), "unql", nil},

		// test using bracket member on object instead of array
		{NewBracketMemberOperator(NewProperty("address"), NewLiteralString("street")), "1 recursive function", nil},
	}

	context := NewMapItem(sampleContext, sampleMeta)

	for _, x := range tests {
		value, err := x.input.Evaluate(context)
		if !reflect.DeepEqual(err, x.err) {
			t.Errorf("Expected no error, got error: %v, %v", err, x.err)
		}
		if !reflect.DeepEqual(value, x.output) {
			t.Errorf("Expected value %v, got %v", x.output, value)
		}

	}

}
