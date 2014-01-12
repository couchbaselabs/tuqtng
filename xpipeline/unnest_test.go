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

var joinerTestData = dparval.ValueCollection{}

func init() {
	doc := dparval.NewValue(map[string]interface{}{
		"name": "mike",
		"children": []interface{}{
			map[string]interface{}{
				"name": "bob",
			},
			map[string]interface{}{
				"name": "dan",
			},
		},
	})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	joinerTestData = append(joinerTestData, doc)

	doc = dparval.NewValue(map[string]interface{}{
		"name": "dustin",
		"children": []interface{}{
			map[string]interface{}{
				"name": "mary",
			},
			map[string]interface{}{
				"name": "jane",
			},
		},
	})
	doc.SetAttachment("meta", map[string]interface{}{"id": "2"})
	joinerTestData = append(joinerTestData, doc)
}

func TestUnnest(t *testing.T) {

	stubSource := NewStubSource(joinerTestData)

	joiner := NewUnnest(ast.NewProperty("children"), "INNER", "child")
	joiner.SetSource(stubSource)

	joinerItemChannel, _ := joiner.GetChannels()

	stopChannel := make(misc.StopChannel)
	go joiner.Run(stopChannel)

	count := 0
	for item := range joinerItemChannel {
		count++
		_, err := item.Path("name")
		if err != nil {
			t.Errorf("Expected to find field name")
		}
		_, err = item.Path("child")
		if err != nil {
			t.Errorf("Expected to find field child")
		}
	}

	if count != 4 {
		t.Errorf("Expected %d items, got %d", 4, count)
	}

}
