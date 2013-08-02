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

func TestFromGenerateAlias(t *testing.T) {

	tests := []struct {
		input  *From
		output *From
	}{
		{&From{Bucket: "bucket", Projection: NewProperty("b")}, &From{Bucket: "bucket", Projection: NewProperty("b"), As: "b"}},
	}

	for _, x := range tests {
		x.input.GenerateAlias()
		if !reflect.DeepEqual(x.input, x.output) {
			t.Errorf("Expected %v got %v", x.output, x.input)
		}
	}
}

func TestGetAliases(t *testing.T) {
	tests := []struct {
		input  *From
		output []string
	}{
		{&From{Bucket: "bucket", Projection: NewProperty("b"), As: "b"}, []string{"b"}},
	}

	for _, x := range tests {
		aliases := x.input.GetAliases()
		if !reflect.DeepEqual(aliases, x.output) {
			t.Errorf("Expected %v for %v, got %v", x.output, x.input, aliases)
		}
	}
}
