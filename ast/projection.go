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

type DuplicateAlias struct {
	alias string
}

func (this *DuplicateAlias) Error() string {
	return fmt.Sprintf("alias %s is defined more than once", this.alias)
}

type ResultExpressionList []*ResultExpression

func (this ResultExpressionList) Validate() error {
	for _, resultExpr := range this {
		if resultExpr.Expr != nil {
			err := resultExpr.Expr.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// this function should be called before assigning default names
// it should check to see if any explicitly named aliases are duplicated
// if so, this is an error
func (this ResultExpressionList) CheckForDuplicateAliases() error {
	namesInUse := map[string]bool{}

	for _, resultExpr := range this {
		if resultExpr.As != "" {
			_, ok := namesInUse[resultExpr.As]
			if !ok {
				namesInUse[resultExpr.As] = true
			} else {
				// name is already in use
				return &DuplicateAlias{resultExpr.As}
			}
		}
	}

	return nil
}

// this function should walk through the result expression list
// assign names to any unnamed expressions
func (this ResultExpressionList) AssignDefaultNames() {

	namesInUse := make([]string, 0)

	// first collect the names
	for _, resultExpr := range this {
		if resultExpr.As != "" {
			namesInUse = append(namesInUse, resultExpr.As)
		}
	}

	// now try to assign a name if expression is a property
	for _, resultExpr := range this {
		prop, ok := (resultExpr.Expr).(*Property)
		if resultExpr.As == "" && ok {
			// assign the new name
			if candidateName, uniq := propertyName(prop, namesInUse); uniq {
				// assign the new name
				resultExpr.As = candidateName
				// record that this name is used
				namesInUse = append(namesInUse, candidateName)
			}
		}
	}

	// now assign default names for anything remaining
	for _, resultExpr := range this {
		if resultExpr.Star == false && resultExpr.As == "" {
			// assign the new name
			resultExpr.As = defaultExpressionName(resultExpr.Expr, namesInUse)
			// record that this name is used
			namesInUse = append(namesInUse, resultExpr.As)
		}
	}

}

type ResultExpression struct {
	Star bool       `json:"star"`
	Expr Expression `json:"expr"`
	As   string     `json:"as"`
}

func NewStarResultExpression() *ResultExpression {
	return &ResultExpression{
		Star: true,
	}
}

func NewDotStarResultExpression(expr Expression) *ResultExpression {
	return &ResultExpression{
		Star: true,
		Expr: expr,
	}
}

func NewResultExpression(expr Expression) *ResultExpression {
	return &ResultExpression{
		Star: false,
		Expr: expr,
	}
}

func NewResultExpressionWithAlias(expr Expression, as string) *ResultExpression {
	return &ResultExpression{
		Star: false,
		Expr: expr,
		As:   as,
	}
}

// this function returns (true, property path) unless the property
// path in question has already been used as a name
func propertyName(prop *Property, namesInUse []string) (string, bool) {
	candidateName := prop.Path
	for _, name := range namesInUse {
		if candidateName == name {
			return "", false
		}
	}
	return candidateName, true
}

// this function is responsible for generating default names for expressions
// that do not otherwise have a logical name. also used to resolve name clashes.
func defaultExpressionName(expr Expression, namesInUse []string) string {
	// starts with $1
	// if already in use, increment to $2
	// repeat until not in use
OUTER:
	for counter := 1; ; counter++ {
		candidateName := fmt.Sprintf("$%d", counter)
		for _, name := range namesInUse {
			if candidateName == name {
				continue OUTER
			}
		}
		return candidateName
	}
}
