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
	"strings"

	"github.com/couchbaselabs/dparval"
)

type FunctionCall struct {
	Type     string                    `json:"type"`
	Name     string                    `json:"name"`
	Operands FunctionArgExpressionList `json:"operands"`
	Distinct bool                      `json:"distinct"`
	minArgs  int
	maxArgs  int
}

func NewFunctionCall(name string, operands FunctionArgExpressionList) FunctionCallExpression {
	functionConstructor := SystemFunctionRegistry[strings.ToUpper(name)]
	if functionConstructor != nil {
		return functionConstructor(operands)
	} else {
		return NewFunctionCallUnknown(operands, name)
	}
}

func (this *FunctionCall) GetName() string {
	return this.Name
}

func (this *FunctionCall) GetOperands() FunctionArgExpressionList {
	return this.Operands
}

func (this *FunctionCall) SetOperands(operands FunctionArgExpressionList) {
	this.Operands = operands
}

func (this *FunctionCall) IsDistinct() bool {
	return this.Distinct
}
func (this *FunctionCall) SetDistinct(distinct bool) {
	this.Distinct = distinct
}

func (this *FunctionCall) EquivalentTo(t Expression) bool {
	// another function call expression?
	that, ok := t.(FunctionCallExpression)
	if !ok {
		return false
	}

	// same name?
	if this.Name != that.GetName() {
		return false
	}

	// same distinctness?
	if this.Distinct != that.IsDistinct() {
		return false
	}

	// same number of operands?
	if len(this.Operands) != len(that.GetOperands()) {
		return false
	}

	// operands are equivalent?
	for i, thisOperand := range this.Operands {
		thatOperand := that.GetOperands()[i]
		if !thisOperand.EquivalentTo(thatOperand) {
			return false
		}
	}

	return true
}

func (this *FunctionCall) String() string {
	return fmt.Sprintf("%s(%v)", this.Name, this.Operands)
}

func (this *FunctionCall) Dependencies() ExpressionList {
	rv := ExpressionList{}

	for _, operand := range this.Operands {
		if operand.Expr != nil {
			rv = append(rv, operand.Expr)
		}
	}

	return rv
}

func (this *FunctionCall) ValidateStars() error {
	for _, arg := range this.Operands {
		if arg.Star == true {
			return fmt.Errorf("the %s() function does not support *", this.Name)
		}
	}

	return nil
}

func (this *FunctionCall) ValidateDistinct() error {
	if this.Distinct {
		return fmt.Errorf("the %s() function does not support DISTINCT", this.Name)
	}
	return nil
}

func (this *FunctionCall) ValidateArity() error {
	if this.minArgs > 0 && this.maxArgs > 0 && this.minArgs == this.maxArgs {
		// check for an exact number of arguments
		argMessage := "argument"
		if this.minArgs > 1 {
			argMessage = "arguments"
		}
		if len(this.Operands) != this.minArgs {
			return fmt.Errorf("the %s() function requires exactly %d %s", this.Name, this.minArgs, argMessage)
		}
		return nil
	}

	if this.minArgs > 0 && len(this.Operands) < this.minArgs {
		argMessage := "argument"
		if this.minArgs > 1 {
			argMessage = "arguments"
		}
		return fmt.Errorf("the %s() function requires at least %d %s", this.Name, this.minArgs, argMessage)
	}

	if this.maxArgs > 0 && len(this.Operands) > this.maxArgs {
		argMessage := "argument"
		if this.maxArgs > 1 {
			argMessage = "arguments"
		}
		return fmt.Errorf("the %s() function requires no more than %d %s", this.Name, this.maxArgs, argMessage)
	}

	return nil
}

func (this *FunctionCall) EvaluateBoth(context *dparval.Value) (*dparval.Value, *dparval.Value, error) {
	lv, err := this.Operands[0].Expr.Evaluate(context)
	if err != nil {
		return nil, nil, err
	}
	rv, err := this.Operands[1].Expr.Evaluate(context)
	if err != nil {
		return nil, nil, err
	}
	return lv, rv, nil
}

func (this *FunctionCall) EvaluateBothRequireArray(context *dparval.Value) ([]interface{}, []interface{}, bool, error) {
	lv, rv, err := this.EvaluateBoth(context)
	if err != nil {
		return nil, nil, false, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.ARRAY {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case []interface{}:
			switch rvalue := rvalue.(type) {
			case []interface{}:
				return lvalue, rvalue, true, nil
			}
		}
	}

	return nil, nil, false, fmt.Errorf("the %s() function requires operands to be of type ARRAY", this.Name)
}

func (this *FunctionCall) EvaluateOperandsForArrayAppend(context *dparval.Value) ([]interface{}, error) {

	lv, err := this.Operands[0].Expr.Evaluate(context)
	if err != nil {
		return nil, err
	}

	rv, err := this.Operands[1].Expr.Evaluate(context)
	if err != nil {
		return nil, err
	}

	if lv.Type() == dparval.ARRAY {
		lvalue := lv.Value()
		rvalue := rv.Value()

		switch lvalue := lvalue.(type) {
		case []interface{}:
			switch rvalue := rvalue.(type) {
			case []interface{}:
				return append(lvalue, rvalue...), nil
			default:
				return append(lvalue, rvalue), nil
			}
		}

	}
	return nil, fmt.Errorf("the %s() function requires that first operand be of type ARRAY", this.Name)
}

func (this *FunctionCall) EvaluateOperandsForArrayPrepend(context *dparval.Value) ([]interface{}, error) {

	lv, err := this.Operands[0].Expr.Evaluate(context)
	if err != nil {
		return nil, err
	}

	rv, err := this.Operands[1].Expr.Evaluate(context)
	if err != nil {
		return nil, err
	}

	if rv.Type() == dparval.ARRAY {
		lvalue := lv.Value()
		rvalue := rv.Value()

		switch rvalue := rvalue.(type) {
		case []interface{}:
			switch lvalue := lvalue.(type) {
			case []interface{}:
				return append(lvalue, rvalue...), nil
			default:
				result := make([]interface{}, 1)
				result[0] = lv
				result[0] = lvalue
				return append(result, rvalue...), nil
			}
		}

	}

	return nil, fmt.Errorf("the %s() function requires that the second operand be of type ARRAY", this.Name)
}
