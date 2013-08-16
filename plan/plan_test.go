//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package plan

import (
	"encoding/json"
	"testing"

	"github.com/couchbaselabs/tuqtng/ast"
)

func TestPlanJSON(t *testing.T) {

	tests := []struct {
		input *Plan
	}{
		{&Plan{NewScan("bucket", "index")}},
		{&Plan{NewFetch(NewScan("bucket", "index"), "bucket", nil, "bucket")}},
		{&Plan{NewFilter(NewFetch(NewScan("bucket", "index"), "bucket", nil, "bucket"), ast.NewLiteralBool(true))}},
		{&Plan{NewFilter(NewFetch(NewScan("bucket", "index"), "bucket", nil, "bucket"), ast.NewPlusOperator(ast.NewLiteralNumber(1.0), ast.NewLiteralNumber(1.0)))}},
	}

	for _, x := range tests {
		_, err := json.MarshalIndent(x.input, "", "    ")
		if err != nil {
			t.Errorf("error serializing json: %v", err)
		}
		// FIXME test doesn't actually assert anything
	}

}
