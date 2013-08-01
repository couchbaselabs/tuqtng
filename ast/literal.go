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

	"github.com/mschoch/dparval"
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

func (this *LiteralNull) Evaluate(item dparval.Value) (dparval.Value, error) {
	return dparval.NewNullValue(), nil
}

func (this *LiteralNull) String() string {
	return fmt.Sprintf("null")
}

func (this *LiteralNull) Validate() error {
	return nil
}

func (this *LiteralNull) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	return nil, nil
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

func (this *LiteralBool) Evaluate(item dparval.Value) (dparval.Value, error) {
	return dparval.NewBooleanValue(this.Val), nil
}

func (this *LiteralBool) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralBool) Validate() error {
	return nil
}

func (this *LiteralBool) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	return nil, nil
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

func (this *LiteralNumber) Evaluate(item dparval.Value) (dparval.Value, error) {
	return dparval.NewNumberValue(this.Val), nil
}

func (this *LiteralNumber) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralNumber) Validate() error {
	return nil
}

func (this *LiteralNumber) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	return nil, nil
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

func (this *LiteralString) Evaluate(item dparval.Value) (dparval.Value, error) {
	return dparval.NewStringValue(this.Val), nil
}

func (this *LiteralString) String() string {
	return fmt.Sprintf("\"%s\"", this.Val)
}

func (this *LiteralString) Validate() error {
	return nil
}

func (this *LiteralString) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	return nil, nil
}

// ****************************************************************************
// Array
// ****************************************************************************

type LiteralArray struct {
	Type string         `json:"type"`
	Val  ExpressionList `json:"value"`
}

func NewLiteralArray(val ExpressionList) *LiteralArray {
	return &LiteralArray{
		Type: "literal_array",
		Val:  val,
	}
}

func (this *LiteralArray) Evaluate(item dparval.Value) (dparval.Value, error) {
	rv := make([]interface{}, 0, len(this.Val))
	for _, v := range this.Val {
		ev, err := v.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// ignore (missing values do not contribute to the array contents)
			default:
				// any other error should be returned to caller
				return nil, err
			}
		} else {
			rv = append(rv, ev)
		}
	}
	return dparval.NewArrayValue(rv), nil
}

func (this *LiteralArray) String() string {
	inside := "["
	for i, a := range this.Val {
		if i != 0 {
			inside = inside + ", "
		}
		inside = inside + fmt.Sprintf("%v", a)
	}
	inside = inside + "]"
	return inside
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

func (this *LiteralArray) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	for i, v := range this.Val {
		newi, err := v.VerifyFormalNotation(aliases, defaultAlias)
		if err != nil {
			return nil, err
		}
		if newi != nil {
			this.Val[i] = newi
		}
	}
	return nil, nil
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

func (this *LiteralObject) Evaluate(item dparval.Value) (dparval.Value, error) {
	rv := make(map[string]interface{}, len(this.Val))
	for k, v := range this.Val {
		ev, err := v.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// ignore (missing values do not contribute to the object contents)
			default:
				// any other error should be returned to caller
				return nil, err
			}
		} else {
			rv[k] = ev
		}
	}
	return dparval.NewObjectValue(rv), nil
}

func (this *LiteralObject) String() string {
	inside := "{"
	count := 0
	for k, v := range this.Val {
		if count != 0 {
			inside = inside + ", "
		}
		inside = inside + fmt.Sprintf("\"%s\": %v", k, v)
		count++
	}
	inside = inside + "}"
	return inside
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

func (this *LiteralObject) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	for k, v := range this.Val {
		newv, err := v.VerifyFormalNotation(aliases, defaultAlias)
		if err != nil {
			return nil, err
		}
		if newv != nil {
			this.Val[k] = newv
		}
	}
	return nil, nil
}
