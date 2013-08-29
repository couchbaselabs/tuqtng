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

func TestCollateJSON(t *testing.T) {

	tests := []struct {
		left   interface{}
		right  interface{}
		result int
	}{

		// scalars
		{true, false, 1},
		{false, true, -1},
		{nil, float64(17), -3},
		{float64(1), float64(1), 0},
		{float64(123), float64(1), 1},
		{float64(123), 0123.0, 0},
		{float64(123), "123", -1},
		{"1234", "123", 1},
		{"1234", "1235", -1},
		{"1234", "1234", 0},

		// verify n1ql collation
		{"A", "a", -1},
		{"a", "aa", -1},
		{"b", "aa", 1},

		// arrays
		{[]interface{}{}, "foo", 1},
		{[]interface{}{}, []interface{}{}, 0},
		{[]interface{}{true}, []interface{}{true}, 0},
		{[]interface{}{false}, []interface{}{nil}, 1},
		{[]interface{}{}, []interface{}{nil}, -1},
		{[]interface{}{float64(123)}, []interface{}{float64(45)}, 1},
		{[]interface{}{float64(123)}, []interface{}{float64(45), float64(67)}, 1},
		{[]interface{}{123.4, "wow"}, []interface{}{123.40, float64(789)}, 1},
		{[]interface{}{float64(5), "wow"}, []interface{}{float64(5), "wow"}, 0},
		{[]interface{}{float64(5), "wow"}, float64(1), 2},
		{[]interface{}{float64(1)}, []interface{}{float64(5), "wow"}, -1},

		// nested arrays
		{[]interface{}{[]interface{}{}}, []interface{}{}, 1},
		{[]interface{}{float64(1), []interface{}{float64(2), float64(3)}, float64(4)},
			[]interface{}{float64(1), []interface{}{float64(2), 3.1}, float64(4), float64(5), float64(6)}, -1},

		// unicode strings
		{"fréd", "fréd", 0},
		{"ømø", "omo", 1},
		{"\t", " ", -1},
		{"\001", " ", -1},

		// object
		{map[string]interface{}{}, "foo", 2},

		// actual object comparisons
		{map[string]interface{}{}, map[string]interface{}{}, 0},
		{map[string]interface{}{"key1": "val1"}, map[string]interface{}{"key1": "val1"}, 0},
		{map[string]interface{}{}, map[string]interface{}{"key1": "val1"}, -1},
		{map[string]interface{}{"key1": "val1"}, map[string]interface{}{}, 1},

		// bigger objects greater
		{map[string]interface{}{"key1": "val1"}, map[string]interface{}{"key1": "val1", "key2": "val2"}, -1},
		{map[string]interface{}{"key1": "val1", "altkey": "altval", "altkey2": "altval2"}, map[string]interface{}{"key1": "val1", "key2": "val2"}, 1},

		// objects with same number of keys but different values
		{map[string]interface{}{"key1": "val1", "key2": "val2a"}, map[string]interface{}{"key1": "val1", "key2": "val2"}, 1},

		// objects with same number of keys but one different key
		// "key2" sorts before "key3", obj1 has "missing" key2, therefore obj1 is less
		{map[string]interface{}{"key1": "val1", "key3": "val3"}, map[string]interface{}{"key1": "val1", "key2": "val2"}, -1},
	}

	for _, test := range tests {
		result := CollateJSON(test.left, test.right)
		if result != test.result {
			t.Errorf("Comparing %v with %v, expected %v, got %v", test.left, test.right, test.result, result)
		}
	}

}
