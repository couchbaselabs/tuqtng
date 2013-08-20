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
	"reflect"
	"testing"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/misc"
)

func TestExpressionEvaluatorSource(t *testing.T) {
	source := NewExpressionEvaluatorSource()

	sourceItemChannel, _ := source.GetChannels()

	stopChannel := make(misc.StopChannel)
	go source.Run(stopChannel)

	count := 0
	for item := range sourceItemChannel {
		count++
		if !reflect.DeepEqual(item, dparval.NewValue(map[string]interface{}{})) {
			t.Errorf("Expected empty map item, got %v", item)
		}
	}

	if count != 1 {
		t.Errorf("expression evaluator should generate exactly 1 row, generated %d", count)
	}
}
