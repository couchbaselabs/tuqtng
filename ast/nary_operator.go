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

type NaryOperator struct {
	operator string
	Operands ExpressionList `json:"operands"`
}

func (this *NaryOperator) Dependencies() ExpressionList {
	rv := ExpressionList{}

	for _, oper := range this.Operands {
		rv = append(rv, oper)
	}

	return rv
}

func (this *NaryOperator) Operator() string {
	return this.operator
}

func (this *NaryOperator) GetOperands() ExpressionList {
	return this.Operands
}

func (this *NaryOperator) SetOperands(operands ExpressionList) {
	this.Operands = operands
}

func (this *NaryOperator) SetOperand(i int, operand Expression) {
	this.Operands[i] = operand
}

func (this *NaryOperator) String() string {
	inside := ""
	for i, oper := range this.Operands {
		if i != 0 {
			inside = inside + fmt.Sprintf(" %s ", this.operator)
		}
		inside = inside + fmt.Sprintf("%v", oper)
	}
	return inside
}

func (this *NaryOperator) EquivalentTo(t Expression) bool {
	// another nary operator?
	that, ok := t.(NaryOperatorExpression)
	if !ok {
		return false
	}

	// same type of operator?
	if this.operator != that.Operator() {
		return false
	}

	// same number of operands?
	if len(this.Operands) != len(that.GetOperands()) {
		return false
	}

	found := make([]bool, len(this.Operands))
	for _, thisOper := range this.Operands {
		for j, thatOper := range that.GetOperands() {
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
