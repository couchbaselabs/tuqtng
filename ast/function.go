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

func (this *FunctionCall) Validate() error {
	// validate the function name is valid
	functionImpl := SystemFunctionRegistry[this.Name]
	if functionImpl == nil {
		return fmt.Errorf("no system function named %v registered", this.Name)
	}

	// ask the implementation to validate the arguments
	return functionImpl.Validate(this.Operands)
}

func (this *FunctionCall) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {

	// two specific checks need to be made here for special functions
	if this.Name == "VALUE" {
		// VALUE() with 0 args is converted to VALUE(defaultAlias) when there is one
		if len(this.Operands) == 0 && defaultAlias != "" {
			this.Operands = append(this.Operands, NewFunctionArgExpression(NewProperty(defaultAlias)))
		}
	}
	if this.Name == "META" {
		// META() with 0 args is converted to META(defaultAlias) when there is one
		if len(this.Operands) == 0 && defaultAlias != "" {
			this.Operands = append(this.Operands, NewFunctionArgExpression(NewProperty(defaultAlias)))
		} else if len(this.Operands) > 0 {
			// check to see that the correct bucket is referenced (currently always aliases[0])
			switch operexpr := this.Operands[0].Expr.(type) {
			case *Property:
				if operexpr.Path != aliases[0] {
					return nil, fmt.Errorf("invalid argument to META() function, must be bucket name/alias")
				}
			default:
				return nil, fmt.Errorf("invalid argument to META() function, must be bucket name/alias")
			}
		}
	}

	for _, operand := range this.Operands {
		if operand.Expr != nil {
			newoper, err := operand.Expr.VerifyFormalNotation(aliases, defaultAlias)
			if err != nil {
				return nil, err
			}
			if newoper != nil {
				operand.Expr = newoper
			}
		}
	}

	return nil, nil
}

func (this *FunctionCall) String() string {
	return fmt.Sprintf("%s(%v)", this.Name, this.Operands)
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
