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

	"github.com/couchbaselabs/tuqtng/query"
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

func (this *PlusOperator) Evaluate(item query.Item) (query.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
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

func (this *SubtractOperator) Evaluate(item query.Item) (query.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
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

func (this *MultiplyOperator) Evaluate(item query.Item) (query.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
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

func (this *DivideOperator) Evaluate(item query.Item) (query.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
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

func (this *ModuloOperator) Evaluate(item query.Item) (query.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
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

func (this *ChangeSignOperator) Evaluate(item query.Item) (query.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch ov := ov.(type) {
	case float64:
		return -ov, nil
	default:
		return nil, nil
	}
}

func (this *ChangeSignOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}
