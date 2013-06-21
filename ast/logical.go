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
// AND
// ****************************************************************************

type AndOperator struct {
	operands []Expression
}

func NewAndOperator(operands []Expression) *AndOperator {
	return &AndOperator{
		operands: operands,
	}
}

func (this *AndOperator) Evaluate(item Item) (Value, error) {
	for _, operand := range this.operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *Undefined:
				// undefined is false
				return false, nil
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}
		// now interpret the evaluated value in a boolean context
		operandBoolVal := ValueInBooleanContext(operandVal)
		if !operandBoolVal {
			return false, nil
		}
	}
	return true, nil
}

func (this *AndOperator) String() string {
	return fmt.Sprintf("AND %v", this.operands)
}

// ****************************************************************************
// OR
// ****************************************************************************

type OrOperator struct {
	operands []Expression
}

func NewOrOperator(operands []Expression) *OrOperator {
	return &OrOperator{
		operands: operands,
	}
}

func (this *OrOperator) Evaluate(item Item) (Value, error) {
	for _, operand := range this.operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *Undefined:
				// do nothing, undefined is false, need to keep looking for true value
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}

		// now interpret the evaluated value in a boolean context
		operandBoolVal := ValueInBooleanContext(operandVal)
		if operandBoolVal {
			return true, nil
		}
	}
	return false, nil
}

func (this *OrOperator) String() string {
	return fmt.Sprintf("OR %v", this.operands)
}

// ****************************************************************************
// NOT
// ****************************************************************************

type NotOperator struct {
	operand Expression
}

func NewNotOperator(operand Expression) *NotOperator {
	return &NotOperator{
		operand: operand,
	}
}

func (this *NotOperator) Evaluate(item Item) (Value, error) {
	ov, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return true, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	operandBoolVal := ValueInBooleanContext(ov)

	return !operandBoolVal, nil
}

func (this *NotOperator) String() string {
	return fmt.Sprintf("NOT %v", this.operand)
}
