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

	"github.com/couchbaselabs/tuqtng/query"
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

		// verify unicode collation
		{"a", "A", -1},
		{"A", "aa", -1},
		{"B", "aa", 1},

		// arrays
		{[]query.Value{}, "foo", 1},
		{[]query.Value{}, []query.Value{}, 0},
		{[]query.Value{true}, []query.Value{true}, 0},
		{[]query.Value{false}, []query.Value{nil}, 1},
		{[]query.Value{}, []query.Value{nil}, -1},
		{[]query.Value{float64(123)}, []query.Value{float64(45)}, 1},
		{[]query.Value{float64(123)}, []query.Value{float64(45), float64(67)}, 1},
		{[]query.Value{123.4, "wow"}, []query.Value{123.40, float64(789)}, 1},
		{[]query.Value{float64(5), "wow"}, []query.Value{float64(5), "wow"}, 0},
		{[]query.Value{float64(5), "wow"}, float64(1), 2},
		{[]query.Value{float64(1)}, []query.Value{float64(5), "wow"}, -1},

		// nested arrays
		{[]query.Value{[]query.Value{}}, []query.Value{}, 1},
		{[]query.Value{float64(1), []query.Value{float64(2), float64(3)}, float64(4)},
			[]query.Value{float64(1), []query.Value{float64(2), 3.1}, float64(4), float64(5), float64(6)}, -1},

		// unicode strings
		{"fréd", "fréd", 0},
		{"ømø", "omo", 1},
		{"\t", " ", -1},
		{"\001", " ", -1},

		// object
		{map[string]query.Value{}, "foo", 2},

		// actual object comparisons
		{map[string]query.Value{}, map[string]query.Value{}, 0},
		{map[string]query.Value{"key1": "val1"}, map[string]query.Value{"key1": "val1"}, 0},
		{map[string]query.Value{}, map[string]query.Value{"key1": "val1"}, -1},
		{map[string]query.Value{"key1": "val1"}, map[string]query.Value{}, 1},

		// bigger objects greater
		{map[string]query.Value{"key1": "val1"}, map[string]query.Value{"key1": "val1", "key2": "val2"}, -1},
		{map[string]query.Value{"key1": "val1", "altkey": "altval", "altkey2": "altval2"}, map[string]query.Value{"key1": "val1", "key2": "val2"}, 1},

		// objects with same number of keys but different values
		{map[string]query.Value{"key1": "val1", "key2": "val2a"}, map[string]query.Value{"key1": "val1", "key2": "val2"}, 1},

		// objects with same number of keys but one different key
		// "key2" sorts before "key3", obj1 has "missing" key2, therefore obj1 is less
		{map[string]query.Value{"key1": "val1", "key3": "val3"}, map[string]query.Value{"key1": "val1", "key2": "val2"}, -1},
	}

	for _, test := range tests {
		result := CollateJSON(test.left, test.right)
		if result != test.result {
			t.Errorf("Comparing %v with %v, expected %v, got %v", test.left, test.right, test.result, result)
		}
	}

}
