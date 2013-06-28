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
)

func TestOffset(t *testing.T) {

	stubSource := NewStubSource(testData)

	offset := NewOffset(3)
	offset.SetSource(stubSource)

	offsetItemChannel := offset.GetItemChannel()

	go offset.Run()

	count := 0
	for _ = range offsetItemChannel {
		count++
	}

	if count != len(testData)-3 {
		t.Errorf("Expected %d items, got %d", len(testData)-3, count)
	}
}
