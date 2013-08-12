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

	"github.com/couchbaselabs/dparval"
)

// ****************************************************************************
// AND
// ****************************************************************************

type AndOperator struct {
	Type     string         `json:"type"`
	Operands ExpressionList `json:"operands"`
}

func NewAndOperator(operands ExpressionList) *AndOperator {
	return &AndOperator{
		Type:     "and",
		Operands: operands,
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

func (this *AndOperator) String() string {
	inside := ""
	for i, oper := range this.Operands {
		if i != 0 {
			inside = inside + " AND "
		}
		inside = inside + fmt.Sprintf("%v", oper)
	}
	return inside
}

func (this *AndOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*AndOperator)
	if !ok {
		return false
	}

	if len(this.Operands) != len(that.Operands) {
		return false
	}

	found := make([]bool, len(this.Operands))
	for _, thisOper := range this.Operands {
		for j, thatOper := range that.Operands {
			if thisOper.EquivalentTo(thatOper) {
				if !found[j] {
					// we cant match the same target item more than once
					found[j] = true
				}
			}
		}
	}

	// now walk the found list and see if anything wasnt found
	for _, fi := range found {
		if !fi {
			return false
		}
	}

	return false
}

func (this *AndOperator) Dependencies() ExpressionList {
	rv := ExpressionList{}

	for _, oper := range this.Operands {
		rv = append(rv, oper)
	}

	return rv
}

func (this *AndOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// OR
// ****************************************************************************

type OrOperator struct {
	Type     string         `json:"type"`
	Operands ExpressionList `json:"operands"`
}

func NewOrOperator(operands ExpressionList) *OrOperator {
	return &OrOperator{
		Type:     "or",
		Operands: operands,
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

func (this *OrOperator) String() string {
	inside := ""
	for i, oper := range this.Operands {
		if i != 0 {
			inside = inside + " OR "
		}
		inside = inside + fmt.Sprintf("%v", oper)
	}
	return inside
}

func (this *OrOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*OrOperator)
	if !ok {
		return false
	}

	if len(this.Operands) != len(that.Operands) {
		return false
	}

	found := make([]bool, len(this.Operands))
	for _, thisOper := range this.Operands {
		for j, thatOper := range that.Operands {
			if thisOper.EquivalentTo(thatOper) {
				if !found[j] {
					// we cant match the same target item more than once
					found[j] = true
				}
			}
		}
	}

	// now walk the found list and see if anything wasnt found
	for _, fi := range found {
		if !fi {
			return false
		}
	}

	return false
}

func (this *OrOperator) Dependencies() ExpressionList {
	rv := ExpressionList{}

	for _, oper := range this.Operands {
		rv = append(rv, oper)
	}

	return rv
}

func (this *OrOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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
		log.Fatalf("Unexpected type %T in NOT", operandBoolVal)
		return dparval.NewValue(nil), nil
	}
}

func (this *NotOperator) String() string {
	return fmt.Sprintf("NOT %v", this.Operand)
}

func (this *NotOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*NotOperator)
	if !ok {
		return false
	}

	if !this.Operand.EquivalentTo(that.Operand) {
		return false
	}

	return true
}

func (this *NotOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *NotOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
