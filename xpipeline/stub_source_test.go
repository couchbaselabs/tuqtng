//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"github.com/couchbaselabs/tuqtng/ast"
	"testing"
)

func TestStubSource(t *testing.T) {

	testData := ast.ItemCollection{
		ast.NewMapItem(map[string]ast.Value{"name": "marty"}, map[string]ast.Value{"id": "1"}),
		ast.NewMapItem(map[string]ast.Value{"name": "steve"}, map[string]ast.Value{"id": "2"}),
		ast.NewMapItem(map[string]ast.Value{"name": "gerald"}, map[string]ast.Value{"id": "3"}),
		ast.NewMapItem(map[string]ast.Value{"name": "siri"}, map[string]ast.Value{"id": "4"}),
	}

	stubSource := NewStubSource(testData)

	stubItemChannel := stubSource.GetItemChannel()

	go stubSource.Run()

	count := 0
	for item := range stubItemChannel {
		count++

		val, err := item.GetPath("name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count == 0 && val != "marty" {
			t.Errorf("expected marty, got %v", val)
		}
	}

	if count != len(testData) {
		t.Errorf("Expected %d items, got %d", len(testData), count)
	}
}
