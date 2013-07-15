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
	"github.com/couchbaselabs/tuqtng/query"
)

// ****************************************************************************
// String Concatenation
// ****************************************************************************

type StringConcatenateOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewStringConcatenateOperator(left, right Expression) *StringConcatenateOperator {
	return &StringConcatenateOperator{
		Type:  "string_concat",
		Left:  left,
		Right: right,
	}
}

func (this *StringConcatenateOperator) Evaluate(item query.Item) (query.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case string:
		switch rv := rv.(type) {
		case string:
			// if both values are string concatenate them
			return lv + rv, nil
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}
}

func (this *StringConcatenateOperator) Validate() error {
	err := this.Left.Validate()
	if err != nil {
		return err
	}
	err = this.Right.Validate()
	if err != nil {
		return err
	}
	return nil
}
