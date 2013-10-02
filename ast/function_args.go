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
	if this.Expr != nil && !this.Expr.EquivalentTo(that.Expr) {
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
