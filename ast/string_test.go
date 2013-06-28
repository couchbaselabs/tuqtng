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

func TestString(t *testing.T) {

	dneProperty := NewProperty("foo")
	numberFive := NewLiteralNumber(5.0)
	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")

	tests := []struct {
		input  Expression
		output interface{}
		err    error
	}{
		{NewStringConcatenateOperator(stringCouchbase, stringServer), "CouchbaseServer", nil},
		{NewStringConcatenateOperator(numberFive, stringServer), nil, nil},
		{NewStringConcatenateOperator(stringCouchbase, numberFive), nil, nil},
		{NewStringConcatenateOperator(dneProperty, stringServer), nil, &Undefined{"foo"}},
		{NewStringConcatenateOperator(stringCouchbase, dneProperty), nil, &Undefined{"foo"}},
	}

	for _, x := range tests {
		result, err := x.input.Evaluate(nil)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error: %v, got %v", x.err, err)
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %t %v, got %t %v", x.output, x.output, result, result)
		}
	}

}
