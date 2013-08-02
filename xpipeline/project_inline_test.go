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
)

func TestProjectInline(t *testing.T) {

	stubSource := NewStubSource(testData)

	project := NewProjectInline(ast.NewResultExpression(ast.NewProperty("boss")))
	project.SetSource(stubSource)

	projectItemChannel, _ := project.GetChannels()

	go project.Run()

	for item := range projectItemChannel {
		_, err := item.Path("fname")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
