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

type SortExpression struct {
	Expr      Expression `json:"expr"`
	Ascending bool       `json:"asc"`
}

func NewSortExpression(expr Expression, asc bool) *SortExpression {
	return &SortExpression{
		Expr:      expr,
		Ascending: asc,
	}
}

type SortExpressionList []*SortExpression

func (this SortExpressionList) Validate() error {
	var err error
	validator := NewExpressionValidator()
	for _, orderExpr := range this {
		if orderExpr.Expr != nil {
			orderExpr.Expr, err = orderExpr.Expr.Accept(validator)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this SortExpressionList) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) error {
	var err error
	formalNotation := NewExpressionFormalNotationConverter(forbiddenAliases, aliases, defaultAlias)
	for _, orderExpr := range this {
		if orderExpr.Expr != nil {
			orderExpr.Expr, err = orderExpr.Expr.Accept(formalNotation)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this SortExpressionList) VerifyAllAggregateFunctionsOrInThisList(groupBy ExpressionList) error {
	for _, orderExpr := range this {
		if orderExpr.Expr != nil {
			fdc := NewExpressionFunctionalDependencyChecker(groupBy)
			_, err := orderExpr.Expr.Accept(fdc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this SortExpressionList) findAggregateFunctionReferences() ExpressionList {
	af := NewExpressionAggregateFinder()
	for _, orderExpr := range this {
		orderExpr.Expr.Accept(af)
	}
	return af.GetAggregates()
}

func (this SortExpressionList) String() string {
	rv := ""
	for i, expr := range this {
		if i != 0 {
			rv = rv + ", "
		}
		rv = rv + fmt.Sprintf("%v", expr)
	}
	return rv
}
