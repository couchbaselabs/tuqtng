//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

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
		return value, nil
	}
	return nil, &Undefined{path}
}

func (this *MapItem) GetMeta() map[string]Value {
	return this.meta
}
