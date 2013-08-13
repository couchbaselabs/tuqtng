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

	"github.com/couchbaselabs/dparval"
)

type UnaryOperator struct {
	operator string
	Operand  Expression `json:"operand"`
}

func (this *UnaryOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *UnaryOperator) GetOperand() Expression {
	return this.Operand
}

func (this *UnaryOperator) SetOperand(operand Expression) {
	this.Operand = operand
}

func (this *UnaryOperator) Operator() string {
	return this.operator
}

func (this *UnaryOperator) String() string {
	return fmt.Sprintf("%v %s", this.Operand, this.operator)
}

func (this *UnaryOperator) EquivalentTo(t Expression) bool {
	// another binary operator?
	that, ok := t.(UnaryOperatorExpression)
	if !ok {
		return false
	}

	// same type of operator?
	if this.operator != that.Operator() {
		return false
	}

	// same operands?
	if this.Operand.EquivalentTo(that.GetOperand()) {
		return true
	}

	return false
}

func (this *UnaryOperator) EvaluateFlagMissing(context *dparval.Value) (*dparval.Value, bool, error) {
	ov, err := this.Operand.Evaluate(context)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return nil, true, err
		default:
			// any other error should be returned to caller
			return nil, false, err
		}
	}

	return ov, false, nil
}

func (this *UnaryOperator) EvaluateRequireNumber(context *dparval.Value) (float64, bool, error) {
	ov, err := this.Operand.Evaluate(context)
	if err != nil {
		return 0.0, true, err
	}

	if ov.Type() == dparval.NUMBER {
		ovalue := ov.Value()
		switch ovalue := ovalue.(type) {
		case float64:
			return ovalue, true, nil
		}
	}

	return 0.0, false, nil
}
