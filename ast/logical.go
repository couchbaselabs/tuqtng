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
	"log"

	"github.com/couchbaselabs/tuqtng/query"
)

// ****************************************************************************
// AND
// ****************************************************************************

type AndOperator struct {
	Type     string       `json:"type"`
	Operands []Expression `json:"operands"`
}

func NewAndOperator(operands []Expression) *AndOperator {
	return &AndOperator{
		Type:     "and",
		Operands: operands,
	}
}

func (this *AndOperator) Evaluate(item query.Item) (query.Value, error) {
	var rv query.Value
	var re error
	rv = true
	re = nil
	for _, operand := range this.Operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				rv = nil
				re = err
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}
		// now interpret the evaluated value in a boolean context
		operandBoolVal := ValueInBooleanContext(operandVal)
		if operandBoolVal == false {
			return false, nil
		} else if operandBoolVal == nil && rv == true {
			rv = operandBoolVal
			re = nil
		}
		// if operandBoolVal is true, do nothing
		// rv starts as true, and should never change back to true
	}
	return rv, re
}

func (this *AndOperator) String() string {
	return fmt.Sprintf("AND %v", this.Operands)
}

func (this *AndOperator) Validate() error {
	for _, o := range this.Operands {
		err := o.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// ****************************************************************************
// OR
// ****************************************************************************

type OrOperator struct {
	Type     string       `json:"type"`
	Operands []Expression `json:"operands"`
}

func NewOrOperator(operands []Expression) *OrOperator {
	return &OrOperator{
		Type:     "or",
		Operands: operands,
	}
}

func (this *OrOperator) Evaluate(item query.Item) (query.Value, error) {
	var rv query.Value
	var re error
	rv = false
	re = nil
	for _, operand := range this.Operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				rv = nil
				re = err
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}
		// now interpret the evaluated value in a boolean context
		operandBoolVal := ValueInBooleanContext(operandVal)
		if operandBoolVal == true {
			return true, nil
		} else if operandBoolVal == nil && rv == false {
			rv = operandBoolVal
			re = nil
		}
		// if operandBoolVal is true, do nothing
		// rv starts as true, and should never change back to true
	}
	return rv, re
}

func (this *OrOperator) String() string {
	return fmt.Sprintf("OR %v", this.Operands)
}

func (this *OrOperator) Validate() error {
	for _, o := range this.Operands {
		err := o.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// ****************************************************************************
// NOT
// ****************************************************************************

type NotOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewNotOperator(operand Expression) *NotOperator {
	return &NotOperator{
		Type:    "not",
		Operand: operand,
	}
}

func (this *NotOperator) Evaluate(item query.Item) (query.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		return nil, err
	}

	operandBoolVal := ValueInBooleanContext(ov)
	switch operandBoolVal := operandBoolVal.(type) {
	case bool:
		return !operandBoolVal, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected type %T in NOT", operandBoolVal)
		return nil, nil
	}
}

func (this *NotOperator) String() string {
	return fmt.Sprintf("NOT %v", this.Operand)
}

func (this *NotOperator) Validate() error {
	err := this.Operand.Validate()
	return err
}
