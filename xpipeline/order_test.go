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

func TestOrder(t *testing.T) {

	stubSource := NewStubSource(testData)

	order := NewOrder([]*ast.SortExpression{ast.NewSortExpression(ast.NewProperty("age"), true)})
	order.SetSource(stubSource)

	orderItemChannel := order.GetItemChannel()

	go order.Run()

	count := 0
	for item := range orderItemChannel {

		val, err := item.GetPath("name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count == 0 && val != "dustin" {
			t.Errorf("expected dustin, got %v", val)
		}
		if count == len(testData)-1 && val != "steve" {
			t.Errorf("expected steve, got %v", val)
		}
		count++
	}
}

func TestOrderDescending(t *testing.T) {

	stubSource := NewStubSource(testData)

	order := NewOrder([]*ast.SortExpression{ast.NewSortExpression(ast.NewProperty("age"), false)})
	order.SetSource(stubSource)

	orderItemChannel := order.GetItemChannel()

	go order.Run()

	count := 0
	for item := range orderItemChannel {

		val, err := item.GetPath("name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count == 0 && val != "steve" {
			t.Errorf("expected steve, got %v", val)
		}
		if count == len(testData)-1 && val != "dustin" {
			t.Errorf("expected dustin, got %v", val)
		}
		count++
	}
}
