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

func (this *PlusOperator) Validate() error {
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

func (this *PlusOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *PlusOperator) String() string {
	return fmt.Sprintf("%v + %v", this.Left, this.Right)
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

func (this *SubtractOperator) Validate() error {
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

func (this *SubtractOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *SubtractOperator) String() string {
	return fmt.Sprintf("%v - %v", this.Left, this.Right)
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

func (this *MultiplyOperator) Validate() error {
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

func (this *MultiplyOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *MultiplyOperator) String() string {
	return fmt.Sprintf("%v * %v", this.Left, this.Right)
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

func (this *DivideOperator) Validate() error {
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

func (this *DivideOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *DivideOperator) String() string {
	return fmt.Sprintf("%v / %v", this.Left, this.Right)
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

func (this *ModuloOperator) Validate() error {
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

func (this *ModuloOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *ModuloOperator) String() string {
	return fmt.Sprintf("%v %% %v", this.Left, this.Right)
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

func (this *ChangeSignOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *ChangeSignOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *ChangeSignOperator) String() string {
	return fmt.Sprintf("-%v", this.Operand)
}
