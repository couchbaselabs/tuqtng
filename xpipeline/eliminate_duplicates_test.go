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
	"github.com/couchbaselabs/tuqtng/misc"
)

var duplicateTestData = dparval.ValueCollection{}

func init() {
	doc := dparval.NewValue(map[string]interface{}{
		"name": "mike",
	})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	duplicateTestData = append(duplicateTestData, doc)

	doc = dparval.NewValue(map[string]interface{}{
		"name": "mike",
	})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	duplicateTestData = append(duplicateTestData, doc)
}

func TestEliminateDuplicates(t *testing.T) {

	stubSource := NewStubSource(duplicateTestData)

	ed := NewEliminateDuplicates()
	ed.SetSource(stubSource)

	edItemChannel, _ := ed.GetChannels()

	stopChannel := make(misc.StopChannel)
	go ed.Run(stopChannel)

	count := 0
	for _ = range edItemChannel {
		count++
	}

	if count != 1 {
		t.Errorf("Expected %d items, got %d", 1, count)
	}

}
