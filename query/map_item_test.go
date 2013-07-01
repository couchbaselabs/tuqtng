//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package query

import (
	"reflect"
	"testing"
)

func TestNormalJSONDocItem(t *testing.T) {

	sampleContext := map[string]Value{
		"name": "will",
	}
	sampleMeta := map[string]Value{
		"id": "first",
	}

	tests := []struct {
		input  string
		output Value
		err    error
	}{
		{"name", "will", nil},
		{"dne", nil, &Undefined{"dne"}},
	}

	context := NewMapItem(sampleContext, sampleMeta)

	for _, x := range tests {
		value, err := context.GetPath(x.input)
		if !reflect.DeepEqual(err, x.err) {
			t.Errorf("Expected no error, got error: %v, %v", err, x.err)
		}
		if !reflect.DeepEqual(value, x.output) {
			t.Errorf("Expected value %v, got %v", x.output, value)
		}

	}

	meta := context.GetMeta()
	if meta["id"] != "first" {
		t.Errorf("Expected first, got %v", meta["id"])
	}

}

func TestTopLevelKeys(t *testing.T) {
	sampleContext := map[string]Value{
		"name": "will",
		"cat":  "dog",
	}

	context := NewMapItem(sampleContext, nil)

	if !reflect.DeepEqual(context.GetTopLevelKeys(), []string{"name", "cat"}) {
		t.Errorf("Expected name,cat, got %v", context.GetTopLevelKeys())
	}
}
