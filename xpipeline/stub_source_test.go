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
)

func TestStubSource(t *testing.T) {

	testData := dparval.ValueCollection{}

	doc := dparval.NewValue(map[string]interface{}{"name": "marty", "mentor": map[string]interface{}{"fname": "bob"}})
	doc.SetAttachment("meta", map[string]interface{}{"id": "1"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "steve", "mentor": map[string]interface{}{"fname": "ross"}})
	doc.SetAttachment("meta", map[string]interface{}{"id": "2"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "gerald", "mentor": map[string]interface{}{"fname": "wendy"}})
	doc.SetAttachment("meta", map[string]interface{}{"id": "3"})
	testData = append(testData, doc)

	doc = dparval.NewValue(map[string]interface{}{"name": "siri", "mentor": map[string]interface{}{"fname": "marley"}})
	doc.SetAttachment("meta", map[string]interface{}{"id": "4"})
	testData = append(testData, doc)

	stubSource := NewStubSource(testData)

	stubItemChannel, _ := stubSource.GetChannels()

	go stubSource.Run()

	count := 0
	for item := range stubItemChannel {
		count++

		val, err := item.Path("name")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		value := val.Value()
		if count == 0 && value != "marty" {
			t.Errorf("expected marty, got %v", value)
		}
	}

	if count != len(testData) {
		t.Errorf("Expected %d items, got %d", len(testData), count)
	}
}
