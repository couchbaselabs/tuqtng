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
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
)

// ****************************************************************************
// AND
// ****************************************************************************

type AndOperator struct {
	Type string `json:"type"`
	NaryOperator
}

func NewAndOperator(operands ExpressionList) *AndOperator {
	return &AndOperator{
		"and",
		NaryOperator{
			operator: "AND",
			Operands: operands,
		},
	}
}

func (this *AndOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	var rv interface{}
	var re error
	rv = true
	re = nil
	for _, operand := range this.Operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				rv = nil
				re = err
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}
		// now interpret the evaluated value in a boolean context
		operandValVal := operandVal.Value()
		operandBoolVal := ValueInBooleanContext(operandValVal)
		if operandBoolVal == false {
			return dparval.NewValue(false), nil
		} else if operandBoolVal == nil && rv == true {
			rv = operandBoolVal
			re = nil
		}
		// if operandBoolVal is true, do nothing
		// rv starts as true, and should never change back to true
	}
	// return missing correclty with value nil
	if re != nil {
		return nil, re
	}
	return dparval.NewValue(rv), re
}

func (this *AndOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// OR
// ****************************************************************************

type OrOperator struct {
	Type string `json:"type"`
	NaryOperator
}

func NewOrOperator(operands ExpressionList) *OrOperator {
	return &OrOperator{
		"or",
		NaryOperator{
			operator: "OR",
			Operands: operands,
		},
	}
}

func (this *OrOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	var rv interface{}
	var re error
	rv = false
	re = nil
	for _, operand := range this.Operands {
		operandVal, err := operand.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				rv = nil
				re = err
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}
		// now interpret the evaluated value in a boolean context
		operandValVal := operandVal.Value()
		operandBoolVal := ValueInBooleanContext(operandValVal)
		if operandBoolVal == true {
			return dparval.NewValue(true), nil
		} else if operandBoolVal == nil && rv == false {
			rv = operandBoolVal
			re = nil
		}
		// if operandBoolVal is true, do nothing
		// rv starts as true, and should never change back to true
	}
	if re != nil {
		return nil, re
	}
	return dparval.NewValue(rv), re
}

func (this *OrOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// NOT
// ****************************************************************************

type NotOperator struct {
	Type string `json:"type"`
	PrefixUnaryOperator
}

func NewNotOperator(operand Expression) *NotOperator {
	return &NotOperator{
		"not",
		PrefixUnaryOperator{
			UnaryOperator{
				operator: "NOT ",
				Operand:  operand,
			},
		},
	}
}

func (this *NotOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		return nil, err
	}

	oval := ov.Value()
	operandBoolVal := ValueInBooleanContext(oval)
	switch operandBoolVal := operandBoolVal.(type) {
	case bool:
		return dparval.NewValue(!operandBoolVal), nil
	case nil:
		return dparval.NewValue(nil), nil
	default:
		clog.Fatal("Unexpected type %T in NOT", operandBoolVal)
		return dparval.NewValue(nil), nil
	}
}

func (this *NotOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
