//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

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
	for _, orderExpr := range this {
		if orderExpr.Expr != nil {
			err := orderExpr.Expr.Validate()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this SortExpressionList) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) error {
	for _, orderExpr := range this {
		if orderExpr.Expr != nil {
			neworderexpr, err := orderExpr.Expr.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
			if err != nil {
				return err
			}
			if neworderexpr != nil {
				orderExpr.Expr = neworderexpr
			}
		}
	}
	return nil
}
