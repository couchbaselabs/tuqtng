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
	"math"

	"github.com/couchbaselabs/dparval"
)

// ****************************************************************************
// PLUS
// ****************************************************************************

type PlusOperator struct {
	Type string `json:"type"`
	CommutativeBinaryOperator
}

func NewPlusOperator(left, right Expression) *PlusOperator {
	return &PlusOperator{
		"plus",
		CommutativeBinaryOperator{
			BinaryOperator{
				operator: "+",
				Left:     left,
				Right:    right,
			},
		},
	}
}

func (this *PlusOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	lv, rv, err := this.EvaluateBoth(context)
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

func (this *PlusOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// MINUS
// ****************************************************************************

type SubtractOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewSubtractOperator(left, right Expression) *SubtractOperator {
	return &SubtractOperator{
		"minus",
		BinaryOperator{
			operator: "-",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *SubtractOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	lv, rv, err := this.EvaluateBoth(context)
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

func (this *SubtractOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// MULT
// ****************************************************************************

type MultiplyOperator struct {
	Type string `json:"type"`
	CommutativeBinaryOperator
}

func NewMultiplyOperator(left, right Expression) *MultiplyOperator {
	return &MultiplyOperator{
		"multiply",
		CommutativeBinaryOperator{
			BinaryOperator{
				operator: "*",
				Left:     left,
				Right:    right,
			},
		},
	}
}

func (this *MultiplyOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	lv, rv, err := this.EvaluateBoth(context)
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

func (this *MultiplyOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// DIV
// ****************************************************************************

type DivideOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewDivideOperator(left, right Expression) *DivideOperator {
	return &DivideOperator{
		"divide",
		BinaryOperator{
			operator: "/",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *DivideOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	lv, rv, err := this.EvaluateBoth(context)
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

func (this *DivideOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// MOD
// ****************************************************************************

type ModuloOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewModuloOperator(left, right Expression) *ModuloOperator {
	return &ModuloOperator{
		"modulo",
		BinaryOperator{
			operator: "%",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *ModuloOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	lv, rv, err := this.EvaluateBoth(context)
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

func (this *ModuloOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Change Sign
// ****************************************************************************

type ChangeSignOperator struct {
	Type string `json:"type"`
	PrefixUnaryOperator
}

func NewChangeSignOperator(operand Expression) *ChangeSignOperator {
	return &ChangeSignOperator{
		"changesign",
		PrefixUnaryOperator{
			UnaryOperator{
				operator: "-",
				Operand:  operand,
			},
		},
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

func (this *ChangeSignOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
