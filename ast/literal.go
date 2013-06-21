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
)

// ****************************************************************************
// NULL
// ****************************************************************************

type LiteralNull struct {
}

func NewLiteralNull() *LiteralNull {
	return &LiteralNull{}
}

func (this *LiteralNull) Evaluate(item Item) (Value, error) {
	return nil, nil
}

func (this *LiteralNull) String() string {
	return fmt.Sprintf("null")
}

// ****************************************************************************
// Bool
// ****************************************************************************

type LiteralBool struct {
	val bool
}

func NewLiteralBool(val bool) *LiteralBool {
	return &LiteralBool{
		val: val,
	}
}

func (this *LiteralBool) Evaluate(item Item) (Value, error) {
	return this.val, nil
}

func (this *LiteralBool) String() string {
	return fmt.Sprintf("%v", this.val)
}

// ****************************************************************************
// Number
// ****************************************************************************

type LiteralNumber struct {
	val float64
}

func NewLiteralNumber(val float64) *LiteralNumber {
	return &LiteralNumber{
		val: val,
	}
}

func (this *LiteralNumber) Evaluate(item Item) (Value, error) {
	return this.val, nil
}

func (this *LiteralNumber) String() string {
	return fmt.Sprintf("%v", this.val)
}

// ****************************************************************************
// String
// ****************************************************************************

type LiteralString struct {
	val string
}

func NewLiteralString(val string) *LiteralString {
	return &LiteralString{
		val: val,
	}
}

func (this *LiteralString) Evaluate(item Item) (Value, error) {
	return this.val, nil
}

func (this *LiteralString) String() string {
	return fmt.Sprintf("%v", this.val)
}

// ****************************************************************************
// Array
// ****************************************************************************

type LiteralArray struct {
	val []Expression
}

func NewLiteralArray(val []Expression) *LiteralArray {
	return &LiteralArray{
		val: val,
	}
}

func (this *LiteralArray) Evaluate(item Item) (Value, error) {
	rv := make([]Value, 0, len(this.val))
	for _, v := range this.val {
		ev, err := v.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *Undefined:
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
	return fmt.Sprintf("%v", this.val)
}

// ****************************************************************************
// Object
// ****************************************************************************

type LiteralObject struct {
	val map[string]Expression
}

func NewLiteralObject(val map[string]Expression) *LiteralObject {
	return &LiteralObject{
		val: val,
	}
}

func (this *LiteralObject) Evaluate(item Item) (Value, error) {
	rv := make(map[string]Value, len(this.val))
	for k, v := range this.val {
		ev, err := v.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *Undefined:
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
	return fmt.Sprintf("%v", this.val)
}
