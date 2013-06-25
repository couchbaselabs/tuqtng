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

func TestValueInBooleanContext(t *testing.T) {
	tests := []struct {
		input  Value
		output Value
	}{
		{nil, nil},
		{true, true},
		{false, false},
		{0.0, false},
		{3.14, true},
		{"", false},
		{"couchbase", true},
		{[]Value{1.0}, true},
		{map[string]Value{"name": "bob"}, true},
	}

	for _, x := range tests {
		result := ValueInBooleanContext(x.input)
		if result != x.output {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}
}
