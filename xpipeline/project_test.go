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

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/misc"
)

func TestProject(t *testing.T) {

	stubSource := NewStubSource(testData)

	project := NewProject(ast.ResultExpressionList{ast.NewResultExpressionWithAlias(ast.NewProperty("name"), "f_name")}, true)
	project.SetSource(stubSource)

	projectItemChannel, _ := project.GetChannels()

	stopChannel := make(misc.StopChannel)
	go project.Run(stopChannel)

	for item := range projectItemChannel {

		projection := item.GetAttachment("projection")
		projectionValue, ok := projection.(*dparval.Value)
		if !ok {
			t.Errorf("Expected item projection to be type Value")
		}
		_, err := projectionValue.Path("f_name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
