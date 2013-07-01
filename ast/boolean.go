//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"log"
	"math"

	"github.com/couchbaselabs/tuqtng/query"
)

// this function is repsonible for determining if a value
// should be considered true or false in a boolean context
// NOTE: this is my first attempt at making this behave consisten with javascript
// SEE http://ecma-international.org/ecma-262/5.1/#sec-9.2
func ValueInBooleanContext(val query.Value) query.Value {

	switch val := val.(type) {
	case nil:
		return nil
	case bool:
		return val
	case float64:
		if val == 0 || math.IsNaN(val) {
			return false
		}
		return true
	case string:
		if len(val) == 0 {
			return false
		}
		return true
	case []query.Value:
		return true
	case map[string]query.Value:
		return true
	default:
		// conservative - throw up on any type we didn't code for
		log.Fatalf("Unexpected type %T", val)
		return false
	}
}
