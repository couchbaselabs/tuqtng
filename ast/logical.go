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
	var rv Value
	var re error
	rv = true
	re = nil
	for _, operand := range this.operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *Undefined:
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
	var rv Value
	var re error
	rv = false
	re = nil
	for _, operand := range this.operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *Undefined:
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
	return fmt.Sprintf("NOT %v", this.operand)
}
