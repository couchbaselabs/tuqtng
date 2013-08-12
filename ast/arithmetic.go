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
	"math"

	"github.com/couchbaselabs/dparval"
)

// ****************************************************************************
// PLUS
// ****************************************************************************

type PlusOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewPlusOperator(left, right Expression) *PlusOperator {
	return &PlusOperator{
		Type:  "plus",
		Left:  left,
		Right: right,
	}
}

func (this *PlusOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.NUMBER {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case float64:
			switch rvalue := rvalue.(type) {
			case float64:
				// if both values are numeric add
				return dparval.NewValue(lvalue + rvalue), nil
			}
		}
	}

	return dparval.NewValue(nil), nil
}

func (this *PlusOperator) String() string {
	return fmt.Sprintf("%v + %v", this.Left, this.Right)
}

func (this *PlusOperator) EquivalentTo(that Expression) bool {
	thatPlus, ok := that.(*PlusOperator)
	if !ok {
		return false
	}

	// for plus order doesnt matter
	if this.Left.EquivalentTo(thatPlus.Left) && this.Right.EquivalentTo(thatPlus.Right) {
		return true
	} else if this.Left.EquivalentTo(thatPlus.Right) && this.Right.EquivalentTo(thatPlus.Left) {
		return true
	}
	return false
}

func (this *PlusOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *PlusOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// MINUS
// ****************************************************************************

type SubtractOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewSubtractOperator(left, right Expression) *SubtractOperator {
	return &SubtractOperator{
		Type:  "minus",
		Left:  left,
		Right: right,
	}
}

func (this *SubtractOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.NUMBER {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case float64:
			switch rvalue := rvalue.(type) {
			case float64:
				// if both values are numeric subtract
				return dparval.NewValue(lvalue - rvalue), nil
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *SubtractOperator) String() string {
	return fmt.Sprintf("%v - %v", this.Left, this.Right)
}

func (this *SubtractOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*SubtractOperator)
	if !ok {
		return false
	}

	// for subtraction order does matter
	if this.Left.EquivalentTo(that.Left) && this.Right.EquivalentTo(that.Right) {
		return true
	}
	return false
}

func (this *SubtractOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *SubtractOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// MULT
// ****************************************************************************

type MultiplyOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewMultiplyOperator(left, right Expression) *MultiplyOperator {
	return &MultiplyOperator{
		Type:  "multiply",
		Left:  left,
		Right: right,
	}
}

func (this *MultiplyOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.NUMBER {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case float64:
			switch rvalue := rvalue.(type) {
			case float64:
				// if both values are numeric multiply
				return dparval.NewValue(lvalue * rvalue), nil
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *MultiplyOperator) String() string {
	return fmt.Sprintf("%v * %v", this.Left, this.Right)
}

func (this *MultiplyOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*MultiplyOperator)
	if !ok {
		return false
	}

	// for multiplication order doesnt matter
	if this.Left.EquivalentTo(that.Left) && this.Right.EquivalentTo(that.Right) {
		return true
	} else if this.Left.EquivalentTo(that.Right) && this.Right.EquivalentTo(that.Left) {
		return true
	}
	return false
}

func (this *MultiplyOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *MultiplyOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// DIV
// ****************************************************************************

type DivideOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewDivideOperator(left, right Expression) *DivideOperator {
	return &DivideOperator{
		Type:  "divide",
		Left:  left,
		Right: right,
	}
}

func (this *DivideOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.NUMBER {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case float64:
			switch rvalue := rvalue.(type) {
			case float64:
				// if both values are numeric divide
				return dparval.NewValue(lvalue / rvalue), nil
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *DivideOperator) String() string {
	return fmt.Sprintf("%v / %v", this.Left, this.Right)
}

func (this *DivideOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*DivideOperator)
	if !ok {
		return false
	}

	// for division order does matter
	if this.Left.EquivalentTo(that.Left) && this.Right.EquivalentTo(that.Right) {
		return true
	}
	return false
}

func (this *DivideOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *DivideOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// MOD
// ****************************************************************************

type ModuloOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewModuloOperator(left, right Expression) *ModuloOperator {
	return &ModuloOperator{
		Type:  "modulo",
		Left:  left,
		Right: right,
	}
}

func (this *ModuloOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.NUMBER {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case float64:
			switch rvalue := rvalue.(type) {
			case float64:
				// if both values are numeric divide
				return dparval.NewValue(math.Mod(lvalue, rvalue)), nil
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *ModuloOperator) String() string {
	return fmt.Sprintf("%v %% %v", this.Left, this.Right)
}

func (this *ModuloOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*ModuloOperator)
	if !ok {
		return false
	}

	// for modulo order does matter
	if this.Left.EquivalentTo(that.Left) && this.Right.EquivalentTo(that.Right) {
		return true
	}
	return false
}

func (this *ModuloOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *ModuloOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Change Sign
// ****************************************************************************

type ChangeSignOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewChangeSignOperator(operand Expression) *ChangeSignOperator {
	return &ChangeSignOperator{
		Type:    "changesign",
		Operand: operand,
	}
}

func (this *ChangeSignOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if ov.Type() == dparval.NUMBER {
		ovalue := ov.Value()
		switch ovalue := ovalue.(type) {
		case float64:
			return dparval.NewValue(-ovalue), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *ChangeSignOperator) String() string {
	return fmt.Sprintf("-%v", this.Operand)
}

func (this *ChangeSignOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*ChangeSignOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}
	return false
}

func (this *ChangeSignOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *ChangeSignOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
