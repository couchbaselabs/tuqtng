//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import ()

type ExpressionList []Expression

func (this ExpressionList) Simplify() error {
	var err error
	es := NewExpressionSimplifier()
	for _, expr := range this {
		if expr != nil {
			expr, err = expr.Accept(es)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this ExpressionList) Validate() error {
	var err error
	validator := NewExpressionValidatorNoAggregates()
	for i, resultExpr := range this {
		if resultExpr != nil {
			this[i], err = resultExpr.Accept(validator)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this ExpressionList) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) error {
	var err error
	formalNotation := NewExpressionFormalNotationConverter(forbiddenAliases, aliases, defaultAlias)
	for i, resultExpr := range this {
		if resultExpr != nil {
			this[i], err = resultExpr.Accept(formalNotation)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this ExpressionList) EquivalentTo(e Expression) Expression {
	for _, expr := range this {
		if expr.EquivalentTo(e) {
			return expr
		}
	}
	return nil
}
