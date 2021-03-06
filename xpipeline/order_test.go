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
	"testing"

	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/misc"
)

func TestOrder(t *testing.T) {

	stubSource := NewStubSource(testData)

	order := NewOrder([]*ast.SortExpression{ast.NewSortExpression(ast.NewProperty("age"), true)}, []string{})
	order.SetSource(stubSource)

	orderItemChannel, _ := order.GetChannels()

	stopChannel := make(misc.StopChannel)
	go order.Run(stopChannel)

	count := 0
	for item := range orderItemChannel {

		val, err := item.Path("name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		value := val.Value()
		if count == 0 && value != "dustin" {
			t.Errorf("expected dustin, got %v", value)
		}
		if count == len(testData)-1 && value != "steve" {
			t.Errorf("expected steve, got %v", value)
		}
		count++
	}
}

func TestOrderDescending(t *testing.T) {

	stubSource := NewStubSource(testData)

	order := NewOrder([]*ast.SortExpression{ast.NewSortExpression(ast.NewProperty("age"), false)}, []string{})
	order.SetSource(stubSource)

	orderItemChannel, _ := order.GetChannels()

	stopChannel := make(misc.StopChannel)
	go order.Run(stopChannel)

	count := 0
	for item := range orderItemChannel {

		val, err := item.Path("name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		value := val.Value()
		if count == 0 && value != "steve" {
			t.Errorf("expected steve, got %v", value)
		}
		if count == len(testData)-1 && value != "dustin" {
			t.Errorf("expected dustin, got %v", value)
		}
		count++
	}
}
