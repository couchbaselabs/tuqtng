//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package couchbase

import (
	"reflect"
	"testing"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
)

func TestViewOptions(t *testing.T) {

	tests := []struct {
		low         catalog.LookupValue
		high        catalog.LookupValue
		inclusion   catalog.RangeInclusion
		viewOptions map[string]interface{}
	}{
		// no start or end, include both, should have no view options
		{nil, nil, catalog.Both, map[string]interface{}{}},

		// start at number 7, inclusive
		{
			catalog.LookupValue{dparval.NewValue(7.0)},
			nil,
			catalog.Both,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
			},
		},

		// start at number 7, exclusive
		{
			catalog.LookupValue{dparval.NewValue(7.0)},
			nil,
			catalog.High,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
				"startkey_docid": MAX_ID,
			},
		},

		// end at number 7, inclusive
		{
			nil,
			catalog.LookupValue{dparval.NewValue(7.0)},
			catalog.Both,
			map[string]interface{}{
				"endkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
			},
		},

		// end at number 7, exclusive
		{
			nil,
			catalog.LookupValue{dparval.NewValue(7.0)},
			catalog.Low,
			map[string]interface{}{
				"endkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
				"endkey_docid": MIN_ID,
			},
		},

		// start at 5 end at number 7, inclusive both
		{
			catalog.LookupValue{dparval.NewValue(5.0)},
			catalog.LookupValue{dparval.NewValue(7.0)},
			catalog.Both,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 5.0,
					},
				},
				"endkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
			},
		},

		// start at 5 end at number 7, inclusive low
		{
			catalog.LookupValue{dparval.NewValue(5.0)},
			catalog.LookupValue{dparval.NewValue(7.0)},
			catalog.Low,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 5.0,
					},
				},
				"endkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
				"endkey_docid": MIN_ID,
			},
		},

		// start at 5 end at number 7, inclusive high
		{
			catalog.LookupValue{dparval.NewValue(5.0)},
			catalog.LookupValue{dparval.NewValue(7.0)},
			catalog.High,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 5.0,
					},
				},
				"startkey_docid": MAX_ID,
				"endkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
			},
		},

		// start end at number 7, exclusive both
		{
			catalog.LookupValue{dparval.NewValue(5.0)},
			catalog.LookupValue{dparval.NewValue(7.0)},
			catalog.Neither,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 5.0,
					},
				},
				"startkey_docid": MAX_ID,
				"endkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 7.0,
					},
				},
				"endkey_docid": MIN_ID,
			},
		},

		// start/end different types
		{
			catalog.LookupValue{dparval.NewValue(5.0)},
			catalog.LookupValue{dparval.NewValue("m")},
			catalog.Both,
			map[string]interface{}{
				"startkey": []interface{}{
					[]interface{}{
						TYPE_NUMBER, 5.0,
					},
				},
				"endkey": []interface{}{
					[]interface{}{
						TYPE_STRING, []float64{109.0},
					},
				},
			},
		},
	}

	for _, test := range tests {
		options := generateViewOptions(test.low, test.high, test.inclusion)
		if !reflect.DeepEqual(options, test.viewOptions) {
			t.Errorf("Expected %v, got %v, for range %v - %v %v", test.viewOptions, options, test.low, test.high, test.inclusion)
		}
	}
}
