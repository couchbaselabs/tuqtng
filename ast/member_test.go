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

func TestDotMember(t *testing.T) {

	sampleContext := map[string]interface{}{
		"address": map[string]interface{}{
			"street": "1 recursive function",
		},
		"contact": map[string]interface{}{
			"name": map[string]interface{}{
				"first": "unql",
				"last":  "couchbase",
				"all": []interface{}{
					"unql",
					"couchbase",
				},
			},
		},
		"friends": []interface{}{
			"a",
			"b",
			"c",
		},
		"name": "bob",
	}
	sampleMeta := map[string]interface{}{
		"id": "first",
	}

	tests := ExpressionTestSet{
		{NewDotMemberOperator(NewProperty("address"), NewProperty("street")), "1 recursive function", nil},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("first")), "unql", nil},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("last")), "couchbase", nil},

		{NewDotMemberOperator(NewProperty("address"), NewProperty("city")), nil, &dparval.Undefined{"city"}},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("middle")), nil, &dparval.Undefined{"middle"}},
		{NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("namez")), NewProperty("first")), nil, &dparval.Undefined{"namez"}},

		{NewDotMemberOperator(NewProperty("name"), NewProperty("city")), nil, &dparval.Undefined{"city"}},

		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(0.0)), "a", nil},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(1.0)), "b", nil},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(2.0)), "c", nil},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(-1.0)), nil, &dparval.Undefined{}},
		{NewBracketMemberOperator(NewProperty("friends"), NewLiteralNumber(10.0)), nil, &dparval.Undefined{}},

		{NewBracketMemberOperator(NewProperty("foo"), NewLiteralNumber(10.0)), nil, &dparval.Undefined{"foo"}},
		{NewBracketMemberOperator(NewProperty("friends"), NewProperty("bar")), nil, &dparval.Undefined{"bar"}},

		//compound test
		{NewBracketMemberOperator(NewDotMemberOperator(NewDotMemberOperator(NewProperty("contact"), NewProperty("name")), NewProperty("all")), NewLiteralNumber(0.0)), "unql", nil},

		// test using bracket member on object instead of array
		{NewBracketMemberOperator(NewProperty("address"), NewLiteralString("street")), "1 recursive function", nil},
	}

	context := dparval.NewValue(sampleContext)
	context.AddMeta("meta", sampleMeta)

	tests.RunWithItem(t, context)

}
