//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package misc

import (
	"math"
)

// a function to replace unrepresentable numeric values with string representations
func SanitizeUnrepresentableJSON(obj interface{}) interface{} {
	switch obj := obj.(type) {
	case map[string]interface{}:
		for k, v := range obj {
			obj[k] = SanitizeUnrepresentableJSON(v)
		}
	case []interface{}:
		for i, v := range obj {
			obj[i] = SanitizeUnrepresentableJSON(v)
		}
	case float64:
		if math.IsNaN(obj) {
			return "NaN"
		} else if math.IsInf(obj, 1) {
			return "+Infinity"
		} else if math.IsInf(obj, -1) {
			return "-Infinity"
		}
	}
	return obj
}
