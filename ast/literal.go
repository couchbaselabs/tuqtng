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

	"github.com/couchbaselabs/dparval"
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

func (this *LiteralNull) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	return dparval.NewValue(nil), nil
}

func (this *LiteralNull) String() string {
	return fmt.Sprintf("null")
}

func (this *LiteralNull) EquivalentTo(t Expression) bool {
	_, ok := t.(*LiteralNull)
	return ok
}

func (this *LiteralNull) Dependencies() ExpressionList {
	return ExpressionList{}
}

func (this *LiteralNull) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LiteralBool) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	return dparval.NewValue(this.Val), nil
}

func (this *LiteralBool) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralBool) EquivalentTo(t Expression) bool {
	that, ok := t.(*LiteralBool)
	if !ok {
		return false
	}

	if this.Val == that.Val {
		return true
	}

	return false
}

func (this *LiteralBool) Dependencies() ExpressionList {
	return ExpressionList{}
}

func (this *LiteralBool) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LiteralNumber) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	return dparval.NewValue(this.Val), nil
}

func (this *LiteralNumber) String() string {
	return fmt.Sprintf("%v", this.Val)
}

func (this *LiteralNumber) EquivalentTo(t Expression) bool {
	that, ok := t.(*LiteralNumber)
	if !ok {
		return false
	}

	if this.Val == that.Val {
		return true
	}

	return false
}

func (this *LiteralNumber) Dependencies() ExpressionList {
	return ExpressionList{}
}

func (this *LiteralNumber) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LiteralString) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	return dparval.NewValue(this.Val), nil
}

func (this *LiteralString) String() string {
	return fmt.Sprintf("\"%s\"", this.Val)
}

func (this *LiteralString) EquivalentTo(t Expression) bool {
	that, ok := t.(*LiteralString)
	if !ok {
		return false
	}

	if this.Val == that.Val {
		return true
	}

	return false
}

func (this *LiteralString) Dependencies() ExpressionList {
	return ExpressionList{}
}

func (this *LiteralString) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LiteralArray) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	return dparval.NewValue(rv), nil
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

func (this *LiteralArray) EquivalentTo(t Expression) bool {
	that, ok := t.(*LiteralArray)
	if !ok {
		return false
	}

	if len(this.Val) != len(that.Val) {
		return false
	}
	for i, thisVal := range this.Val {
		thatVal := that.Val[i]
		if !thisVal.EquivalentTo(thatVal) {
			return false
		}
	}

	return true
}

func (this *LiteralArray) Dependencies() ExpressionList {
	rv := ExpressionList{}
	for _, thisVal := range this.Val {
		rv = append(rv, thisVal)
	}
	return rv
}

func (this *LiteralArray) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LiteralObject) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	return dparval.NewValue(rv), nil
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

func (this *LiteralObject) EquivalentTo(t Expression) bool {
	that, ok := t.(*LiteralObject)
	if !ok {
		return false
	}

	if len(this.Val) != len(that.Val) {
		return false
	}
	for key, thisVal := range this.Val {
		thatVal, ok := that.Val[key]
		if !ok {
			return false
		}
		if !thisVal.EquivalentTo(thatVal) {
			return false
		}
	}

	return true
}

func (this *LiteralObject) Dependencies() ExpressionList {
	rv := ExpressionList{}
	for _, thisVal := range this.Val {
		rv = append(rv, thisVal)
	}
	return rv
}

func (this *LiteralObject) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
