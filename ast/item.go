//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

// this is the primary object passed through the execution pipeline
// still open questions about how aliases will be handled

type Item interface {
	// path should be a simple path (not containing . or [])
	// those must be implemented as separate sub-expressions
	GetPath(path string) (Value, error)
	GetMeta() map[string]Value
}

type ItemCollection []Item
