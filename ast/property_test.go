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

	"github.com/couchbaselabs/dparval"
)

func TestPropertyStringRepresentation(t *testing.T) {
	tests := ExpressionStringTestSet{
		{NewProperty("name"), "name"},
	}

	tests.Run(t)
}

func TestEvaluateProperty(t *testing.T) {
	sampleDocument := map[string]interface{}{
		"name": "will",
	}

	tests := ExpressionTestSet{
		{NewProperty("name"), "will", nil},
	}

	item := dparval.NewValue(sampleDocument)

	tests.RunWithItem(t, item)

}

func TestPropertyVerifyFormalNotation(t *testing.T) {
	prop := NewProperty("path")
	newop, err := prop.VerifyFormalNotation([]string{"bucket"}, "bucket")
	if err != nil {
		t.Errorf("Expected no error")
	}
	_, ok := newop.(*DotMemberOperator)
	if !ok {
		t.Errorf("Expected to be converted to DotMemberOperator")
	}
}
