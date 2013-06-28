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

var testData = ast.ItemCollection{
	ast.NewMapItem(map[string]ast.Value{"name": "mike", "age": 100.0}, map[string]ast.Value{"id": "1"}),
	ast.NewMapItem(map[string]ast.Value{"name": "dustin"}, map[string]ast.Value{"id": "1"}),
	ast.NewMapItem(map[string]ast.Value{"name": "bob", "age": nil}, map[string]ast.Value{"id": "1"}),
	ast.NewMapItem(map[string]ast.Value{"name": "marty", "age": 99.0}, map[string]ast.Value{"id": "1"}),
	ast.NewMapItem(map[string]ast.Value{"name": "steve", "age": 200.0}, map[string]ast.Value{"id": "2"}),
	ast.NewMapItem(map[string]ast.Value{"name": "gerald", "age": 175.0}, map[string]ast.Value{"id": "3"}),
	ast.NewMapItem(map[string]ast.Value{"name": "siri", "age": 74.0}, map[string]ast.Value{"id": "4"}),
	ast.NewMapItem(map[string]ast.Value{"name": "ali", "age": 100.0}, map[string]ast.Value{"id": "1"}),
}

func TestFilterTrue(t *testing.T) {

	stubSource := NewStubSource(testData)

	filter := NewFilter(ast.NewLiteralBool(true))
	filter.SetSource(stubSource)

	filterItemChannel := filter.GetItemChannel()

	go filter.Run()

	count := 0
	for _ = range filterItemChannel {
		count++
	}

	if count != len(testData) {
		t.Errorf("Expected %d items, got %d", len(testData), count)
	}

}

func TestFilterFalse(t *testing.T) {

	stubSource := NewStubSource(testData)

	filter := NewFilter(ast.NewLiteralBool(false))
	filter.SetSource(stubSource)

	filterItemChannel := filter.GetItemChannel()

	go filter.Run()

	count := 0
	for _ = range filterItemChannel {
		count++
	}

	if count != 0 {
		t.Errorf("Expected %d items, got %d", 0, count)
	}

}

func TestFilterSome(t *testing.T) {

	stubSource := NewStubSource(testData)

	filter := NewFilter(ast.NewGreaterThanOperator(ast.NewProperty("name"), ast.NewLiteralString("n")))
	filter.SetSource(stubSource)

	filterItemChannel := filter.GetItemChannel()

	go filter.Run()

	count := 0
	for _ = range filterItemChannel {
		count++
	}

	if count != 2 {
		t.Errorf("Expected %d items, got %d", 2, count)
	}

}
