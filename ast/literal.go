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
	"fmt"

	"github.com/couchbaselabs/tuqtng/query"
)

// ****************************************************************************
// NULL
// ****************************************************************************

type LiteralNull struct {
	Type string `json:"type"`
}

func NewLiteralNull() *LiteralNull {
	return &LiteralNull{
		Type: "literal_null",
	}
}

func (this *LiteralNull) Evaluate(item query.Item) (query.Value, error) {
	return nil, nil
}

func (this *LiteralNull) String() string {
	return fmt.Sprintf("null")
}

func (this *LiteralNull) Validate() error {
	return nil
}

// ****************************************************************************
// Bool
// ****************************************************************************

type LiteralBool struct {
	Type string `json:"type"`
	Val  bool   `json:"value"`
}

func NewLiteralBool(val bool) *LiteralBool {
	return &LiteralBool{
		Type: "literal_bool",
		Val:  val,
	}
}

func (this *LiteralBool) Evaluate(item query.Item) (query.Value, error) {
	return this.Val, nil
}

func (this *LiteralBool) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralBool) Validate() error {
	return nil
}

// ****************************************************************************
// Number
// ****************************************************************************

type LiteralNumber struct {
	Type string  `json:"type"`
	Val  float64 `json:"value"`
}

func NewLiteralNumber(val float64) *LiteralNumber {
	return &LiteralNumber{
		Type: "literal_number",
		Val:  val,
	}
}

func (this *LiteralNumber) Evaluate(item query.Item) (query.Value, error) {
	return this.Val, nil
}

func (this *LiteralNumber) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralNumber) Validate() error {
	return nil
}

// ****************************************************************************
// String
// ****************************************************************************

type LiteralString struct {
	Type string `json:"type"`
	Val  string `json:"value"`
}

func NewLiteralString(val string) *LiteralString {
	return &LiteralString{
		Type: "literal_string",
		Val:  val,
	}
}

func (this *LiteralString) Evaluate(item query.Item) (query.Value, error) {
	return this.Val, nil
}

func (this *LiteralString) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralString) Validate() error {
	return nil
}

// ****************************************************************************
// Array
// ****************************************************************************

type LiteralArray struct {
	Type string       `json:"type"`
	Val  []Expression `json:"value"`
}

func NewLiteralArray(val []Expression) *LiteralArray {
	return &LiteralArray{
		Type: "literal_array",
		Val:  val,
	}
}

func (this *LiteralArray) Evaluate(item query.Item) (query.Value, error) {
	rv := make([]query.Value, 0, len(this.Val))
	for _, v := range this.Val {
		ev, err := v.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// ignore (missing values do not contribute to the array contents)
			default:
				// any other error should be returned to caller
				return nil, err
			}
		} else {
			rv = append(rv, ev)
		}
	}
	return rv, nil
}

func (this *LiteralArray) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralArray) Validate() error {
	for _, v := range this.Val {
		err := v.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// ****************************************************************************
// Object
// ****************************************************************************

type LiteralObject struct {
	Type string `json:"type"`
	Val  map[string]Expression
}

func NewLiteralObject(val map[string]Expression) *LiteralObject {
	return &LiteralObject{
		Type: "literal_object",
		Val:  val,
	}
}

func (this *LiteralObject) Evaluate(item query.Item) (query.Value, error) {
	rv := make(map[string]query.Value, len(this.Val))
	for k, v := range this.Val {
		ev, err := v.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// ignore (missing values do not contribute to the object contents)
			default:
				// any other error should be returned to caller
				return nil, err
			}
		} else {
			rv[k] = ev
		}
	}
	return rv, nil
}

func (this *LiteralObject) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralObject) Validate() error {
	for _, v := range this.Val {
		err := v.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
