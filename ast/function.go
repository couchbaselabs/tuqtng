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

type FunctionCall struct {
	Type     string                    `json:"type"`
	Name     string                    `json:"name"`
	Operands FunctionArgExpressionList `json:"operands"`
}

func NewFunctionCall(name string, operands FunctionArgExpressionList) *FunctionCall {
	return &FunctionCall{
		Type:     "function",
		Name:     name,
		Operands: operands,
	}
}

func (this *FunctionCall) Evaluate(item query.Item) (query.Value, error) {
	functionImpl := SystemFunctionRegistry[this.Name]
	if functionImpl != nil {
		return functionImpl.Evaluate(item, this.Operands)
	}

	// FIXME should never happen once we have semantic validation
	log.Printf("no system function named %v registered", this.Name)
	return nil, nil
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

// function arguments

type FunctionArgExpressionList []*FunctionArgExpression

type FunctionArgExpression struct {
	Star bool       `json:"star"`
	Expr Expression `json:"expr"`
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
