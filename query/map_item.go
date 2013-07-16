//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package query

//  MapItem is an implementation of Item using a map structure
//  (typically this is from parsed JSON)
type MapItem struct {
	contents map[string]Value
	meta     map[string]Value
}

func NewMapItem(contents, meta map[string]Value) *MapItem {
	return &MapItem{
		contents: contents,
		meta:     meta,
	}
}

func (this *MapItem) GetPath(path string) (Value, error) {
	value, ok := this.contents[path]
	if ok {
		// FIXME review this to possibly avoid rebuilding maps (custom JSON parser up front?)
		// check type of value
		// map[string]interface{} should be converted to map[string]Value
		// []interface{} should be converted to []Value
		switch value := value.(type) {
		case map[string]interface{}:
			rv := make(map[string]Value)
			for k, v := range value {
				rv[k] = v
			}
			return rv, nil
		case []interface{}:
			rv := make([]Value, 0)
			for _, v := range value {
				rv = append(rv, v)
			}
			return rv, nil
		default:
			return value, nil
		}

	}
	return nil, &Undefined{path}
}

func (this *MapItem) GetTopLevelKeys() []string {
	rv := make([]string, 0, len(this.contents))
	for k, _ := range this.contents {
		rv = append(rv, k)
	}
	return rv
}

func (this *MapItem) GetValue() Value {
	return this.contents
}

func (this *MapItem) GetMeta() map[string]Value {
	return this.meta
}
