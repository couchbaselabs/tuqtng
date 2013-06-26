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
)

// ****************************************************************************
// PLUS
// ****************************************************************************

type PlusOperator struct {
	left  Expression
	right Expression
}

func NewPlusOperator(left, right Expression) *PlusOperator {
	return &PlusOperator{
		left:  left,
		right: right,
	}
}

func (this *PlusOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case float64:
		switch rv := rv.(type) {
		case float64:
			// if both values are numeric add
			return lv + rv, nil
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// MINUS
// ****************************************************************************

type SubtractOperator struct {
	left  Expression
	right Expression
}

func NewSubtractOperator(left, right Expression) *SubtractOperator {
	return &SubtractOperator{
		left:  left,
		right: right,
	}
}

func (this *SubtractOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case float64:
		switch rv := rv.(type) {
		case float64:
			// if both values are numeric subtract
			return lv - rv, nil
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// MULT
// ****************************************************************************

type MultiplyOperator struct {
	left  Expression
	right Expression
}

func NewMultiplyOperator(left, right Expression) *MultiplyOperator {
	return &MultiplyOperator{
		left:  left,
		right: right,
	}
}

func (this *MultiplyOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case float64:
		switch rv := rv.(type) {
		case float64:
			// if both values are numeric multiply
			return lv * rv, nil
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// DIV
// ****************************************************************************

type DivideOperator struct {
	left  Expression
	right Expression
}

func NewDivideOperator(left, right Expression) *DivideOperator {
	return &DivideOperator{
		left:  left,
		right: right,
	}
}

func (this *DivideOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case float64:
		switch rv := rv.(type) {
		case float64:
			// if both values are numeric divide
			return lv / rv, nil
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// MOD
// ****************************************************************************

type ModuloOperator struct {
	left  Expression
	right Expression
}

func NewModuloOperator(left, right Expression) *ModuloOperator {
	return &ModuloOperator{
		left:  left,
		right: right,
	}
}

func (this *ModuloOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case float64:
		switch rv := rv.(type) {
		case float64:
			// if both values are numeric divide
			return math.Mod(lv, rv), nil
		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// Change Sign
// ****************************************************************************

type ChangeSignOperator struct {
	operand Expression
}

func NewChangeSignOperator(operand Expression) *ChangeSignOperator {
	return &ChangeSignOperator{
		operand: operand,
	}
}

func (this *ChangeSignOperator) Evaluate(item Item) (Value, error) {
	ov, err := this.operand.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch ov := ov.(type) {
	case float64:
		return -ov, nil
	default:
		return nil, nil
	}

	return nil, nil
}
