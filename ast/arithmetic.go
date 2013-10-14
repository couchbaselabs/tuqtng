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

func (this *PlusOperator) Copy() Expression {
	return &PlusOperator{
		"plus",
		CommutativeBinaryOperator{
			BinaryOperator{
				operator: "+",
				Left:     this.Left.Copy(),
				Right:    this.Right.Copy(),
			},
		},
	}
}

func (this *PlusOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothNumeric, err := this.EvaluateBothRequireNumber(context)
	if err != nil {
		return nil, err
	}
	if bothNumeric {
		result = lv + rv
	}
	return dparval.NewValue(result), nil
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

func (this *SubtractOperator) Copy() Expression {
	return &SubtractOperator{
		"minus",
		BinaryOperator{
			operator: "-",
			Left:     this.Left.Copy(),
			Right:    this.Right.Copy(),
		},
	}
}

func (this *SubtractOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothNumeric, err := this.EvaluateBothRequireNumber(context)
	if err != nil {
		return nil, err
	}
	if bothNumeric {
		result = lv - rv
	}
	return dparval.NewValue(result), nil
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

func (this *MultiplyOperator) Copy() Expression {
	return &MultiplyOperator{
		"multiply",
		CommutativeBinaryOperator{
			BinaryOperator{
				operator: "*",
				Left:     this.Left.Copy(),
				Right:    this.Right.Copy(),
			},
		},
	}
}

func (this *MultiplyOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothNumeric, err := this.EvaluateBothRequireNumber(context)
	if err != nil {
		return nil, err
	}
	if bothNumeric {
		result = lv * rv
	}
	return dparval.NewValue(result), nil
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

func (this *DivideOperator) Copy() Expression {
	return &DivideOperator{
		"divide",
		BinaryOperator{
			operator: "/",
			Left:     this.Left.Copy(),
			Right:    this.Right.Copy(),
		},
	}
}

func (this *DivideOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothNumeric, err := this.EvaluateBothRequireNumber(context)
	if err != nil {
		return nil, err
	}
	if bothNumeric {
		result = lv / rv
	}
	return dparval.NewValue(result), nil
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

func (this *ModuloOperator) Copy() Expression {
	return &ModuloOperator{
		"modulo",
		BinaryOperator{
			operator: "%",
			Left:     this.Left.Copy(),
			Right:    this.Right.Copy(),
		},
	}
}

func (this *ModuloOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothNumeric, err := this.EvaluateBothRequireNumber(context)
	if err != nil {
		return nil, err
	}
	if bothNumeric {
		result = math.Mod(lv, rv)
	}
	return dparval.NewValue(result), nil
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

func (this *ChangeSignOperator) Copy() Expression {
	return &ChangeSignOperator{
		"changesign",
		PrefixUnaryOperator{
			UnaryOperator{
				operator: "-",
				Operand:  this.Operand.Copy(),
			},
		},
	}
}

func (this *ChangeSignOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	ov, isNumber, err := this.EvaluateRequireNumber(context)
	if err != nil {
		return nil, err
	}

	if isNumber {
		result = -ov
	}
	return dparval.NewValue(result), nil
}

func (this *ChangeSignOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
