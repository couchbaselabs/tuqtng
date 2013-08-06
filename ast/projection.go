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

func (this ResultExpressionList) VerifyFormalNotation(aliases []string, defaultAlias string) error {
	for _, resultExpr := range this {
		if resultExpr.Expr != nil {
			newres, err := resultExpr.Expr.VerifyFormalNotation(aliases, defaultAlias)
			if err != nil {
				return err
			}
			if newres != nil {
				resultExpr.Expr = newres
			}
		}
		if resultExpr.Star && resultExpr.Expr == nil {
			// only a star, need to fixup
			if defaultAlias != "" {
				resultExpr.Expr = NewProperty(defaultAlias)
			}
			// else {
			// 	return fmt.Errorf("* is missing qualifier bucket/alias")
			// }
		}
	}
	return nil
}

// this function should be called before assigning default names
// it should check to see if any explicitly named aliases are duplicated
// if so, this is an error
func (this ResultExpressionList) CheckForDuplicateAliases() ([]string, error) {
	rv := []string{}
	namesInUse := map[string]bool{}

	for _, resultExpr := range this {
		if resultExpr.As != "" {
			_, ok := namesInUse[resultExpr.As]
			if !ok {
				namesInUse[resultExpr.As] = true
				rv = append(rv, resultExpr.As)
			} else {
				// name is already in use
				return nil, &DuplicateAlias{resultExpr.As}
			}
		}
	}

	return rv, nil
}

// this function should walk through the result expression list
// assign names to any unnamed expressions
func (this ResultExpressionList) AssignDefaultNames(namesInUse []string) error {

	// now try to assign a name if expression is a property
	for _, resultExpr := range this {

		// here we're only concerned with expressions not explicitly named
		if resultExpr.As == "" {
			// we're looking for expressions ending in a property
			prop := expressionEndsInProperty(resultExpr.Expr)
			if prop != nil {
				// assign the new name
				if candidateName, uniq := propertyName(prop, namesInUse); uniq {
					// assign the new name
					resultExpr.As = candidateName
					// record that this name is used
					namesInUse = append(namesInUse, candidateName)
				} else {
					return &DuplicateAlias{prop.Path}
				}
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

	return nil
}

func (this ResultExpressionList) String() string {
	rv := ""
	for i, expr := range this {
		if i != 0 {
			rv = rv + ", "
		}
		rv = rv + fmt.Sprintf("%v", expr)
	}
	return rv
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

// if this expression ends in a Property
// return that property
// otherwise nil
// this allows Expressions of the form
// a to return the Property a
// and
// a.b.c to return the Property c
func expressionEndsInProperty(expr Expression) *Property {
	switch expr := expr.(type) {
	case *Property:
		return expr
	case *DotMemberOperator:
		return expr.Right
	}
	return nil
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
