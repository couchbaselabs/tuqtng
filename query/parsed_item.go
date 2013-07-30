//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package query

import ()

//  ParsedItem is an implementation of Item using an already parsed Value
//  (typically this is from parsed JSON)
type ParsedItem struct {
	contents Value
	meta     map[string]Value
}

func NewParsedItem(contents Value, meta map[string]Value) *ParsedItem {
	return &ParsedItem{
		contents: convertEmptyInterfaceToValue(contents),
		meta:     meta,
	}
}

func convertEmptyInterfaceToValue(input Value) Value {
	switch input := input.(type) {
	case map[string]Value:
		for k, v := range input {
			input[k] = convertEmptyInterfaceToValue(v)
		}
		return input
	case map[string]interface{}:
		rv := map[string]Value{}
		for k, v := range input {
			rv[k] = convertEmptyInterfaceToValue(v)
		}
		return rv
	case []Value:
		for i, v := range input {
			input[i] = convertEmptyInterfaceToValue(v)
		}
		return input
	case []interface{}:
		rv := make([]Value, 0)
		for _, v := range input {
			rv = append(rv, convertEmptyInterfaceToValue(v))
		}
		return rv
	default:
		return input
	}
}

func (this *ParsedItem) GetPath(path string) (Value, error) {
	contents := this.contents
	switch contents := contents.(type) {
	case map[string]Value:
		value, ok := contents[path]
		if ok {
			return value, nil
		}
	}

	return nil, &Undefined{path}
}

func (this *ParsedItem) GetValue() Value {
	return this.contents
}

func (this *ParsedItem) GetMeta() map[string]Value {
	return this.meta
}
