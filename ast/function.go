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
	"strings"

	"github.com/couchbaselabs/dparval"
)

type FunctionCall struct {
	Type     string                    `json:"type"`
	Name     string                    `json:"name"`
	Operands FunctionArgExpressionList `json:"operands"`
}

func NewFunctionCall(name string, operands FunctionArgExpressionList) *FunctionCall {
	return &FunctionCall{
		Type:     "function",
		Name:     strings.ToUpper(name),
		Operands: operands,
	}
}

func (this *FunctionCall) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	functionImpl := SystemFunctionRegistry[this.Name]
	if functionImpl != nil {
		return functionImpl.Evaluate(item, this.Operands)
	}

	// FIXME should never happen once we have semantic validation
	log.Printf("no system function named %v registered", this.Name)
	return dparval.NewValue(nil), nil
}

func (this *FunctionCall) EquivalentTo(t Expression) bool {
	that, ok := t.(*FunctionCall)
	if !ok {
		return false
	}

	if this.Name != that.Name {
		return false
	}

	if len(this.Operands) != len(that.Operands) {
		return false
	}

	for i, thisOperand := range this.Operands {
		thatOperand := that.Operands[i]
		if !thisOperand.EquivalentTo(thatOperand) {
			return false
		}
	}

	return true
}

func (this *FunctionCall) IsAggregate() bool {
	functionImpl := SystemFunctionRegistry[this.Name]
	if functionImpl != nil {
		return functionImpl.IsAggregate()
	}
	return false
}

func (this *FunctionCall) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	functionImpl := SystemFunctionRegistry[this.Name]
	switch functionImpl := functionImpl.(type) {
	case AggregateFunction:
		return functionImpl.UpdateAggregate(group, item, this.Operands)
	}
	return fmt.Errorf("Not an aggregate function %T", functionImpl)
}

func (this *FunctionCall) DefaultAggregate(group *dparval.Value) error {
	functionImpl := SystemFunctionRegistry[this.Name]
	switch functionImpl := functionImpl.(type) {
	case AggregateFunction:
		return functionImpl.DefaultAggregate(group, this.Operands)
	}
	return fmt.Errorf("Not an aggregate function %T", functionImpl)
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

func (this *FunctionCall) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// function arguments

type FunctionArgExpressionList []*FunctionArgExpression

func (this FunctionArgExpressionList) String() string {
	inside := ""
	for i, arg := range this {
		if i != 0 {
			inside = inside + ", "
		}
		inside = inside + fmt.Sprintf("%v", arg)
	}
	return inside
}

type FunctionArgExpression struct {
	Star bool       `json:"star"`
	Expr Expression `json:"expr"`
}

func (this *FunctionArgExpression) String() string {
	inside := ""
	if this.Expr != nil {
		inside = fmt.Sprintf("%v", this.Expr)
	}
	if this.Star {
		if inside != "" {
			inside = inside + ".*"
		} else {
			inside = "*"
		}
	}
	return inside
}

func (this *FunctionArgExpression) EquivalentTo(that *FunctionArgExpression) bool {
	if this.Star != that.Star {
		return false
	}
	if !this.Expr.EquivalentTo(that.Expr) {
		return false
	}
	return true
}

func NewStarFunctionArgExpression() *FunctionArgExpression {
	return &FunctionArgExpression{
		Star: true,
	}
}

func NewDotStarFunctionArgExpression(expr Expression) *FunctionArgExpression {
	return &FunctionArgExpression{
		Star: true,
		Expr: expr,
	}
}

func NewFunctionArgExpression(expr Expression) *FunctionArgExpression {
	return &FunctionArgExpression{
		Star: false,
		Expr: expr,
	}
}

func ValidateNoStars(function SystemFunction, arguments FunctionArgExpressionList) error {
	for _, arg := range arguments {
		if arg.Star == true {
			return fmt.Errorf("the %s() function does not support *", function.Name())
		}
	}

	return nil
}

func ValidateArity(function SystemFunction, arguments FunctionArgExpressionList, min, max int) error {
	if min > 0 && max > 0 && min == max {
		// check for an exact number of arguments
		argMessage := "argument"
		if min > 1 {
			argMessage = "arguments"
		}
		if len(arguments) != min {
			return fmt.Errorf("the %s() function requires exactly %d %s", function.Name(), min, argMessage)
		}
		return nil
	}

	if min > 0 && len(arguments) < min {
		argMessage := "argument"
		if min > 1 {
			argMessage = "arguments"
		}
		return fmt.Errorf("the %s() function requires at least %d %s", function.Name(), min, argMessage)
	}

	if max > 0 && len(arguments) > max {
		argMessage := "argument"
		if max > 1 {
			argMessage = "arguments"
		}
		return fmt.Errorf("the %s() function requires no more than %d %s", function.Name(), max, argMessage)
	}

	return nil
}
