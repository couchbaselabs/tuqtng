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

	"github.com/couchbaselabs/dparval"
)

type ExpressionTest struct {
	input  Expression
	output interface{}
	err    error
}

type ExpressionTestSet []ExpressionTest

func (this ExpressionTestSet) Run(t *testing.T) {
	this.RunWithItem(t, nil)
}

func (this ExpressionTestSet) RunWithItem(t *testing.T, item *dparval.Value) {
	for _, x := range this {
		resval, err := x.input.Evaluate(item)
		if !reflect.DeepEqual(err, x.err) {
			t.Fatalf("Expected error: %v, got %v for %v", x.err, err, x.input)
		}
		var result interface{}
		if resval != nil {
			result = resval.Value()
		}
		if !reflect.DeepEqual(result, x.output) {
			t.Errorf("Expected %v, got %v for %v", x.output, result, x.input)
		}
	}
}
