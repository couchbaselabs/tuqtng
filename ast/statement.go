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

// Statement is the abstract representation of an UNQL statement
type Statement interface {
	IsDistinct() bool
	GetResultExpressionList() ResultExpressionList
	GetFroms() []*From
	GetWhere() Expression
	GetOrderBy() []*SortExpression
	GetOffset() int
	GetLimit() int
	SetExplainOnly(bool)
	IsExplainOnly() bool
	VerifySemantics() error
	GetFromAliases() []string
}

type SelectStatement struct {
	Distinct    bool                 `json:"distinct"`
	Select      ResultExpressionList `json:"select"`
	Froms       []*From              `json:"froms"`
	Where       Expression           `json:"where"`
	OrderBy     []*SortExpression    `json:"orderby"`
	Limit       int                  `json:"limit"`
	Offset      int                  `json:"offset"`
	ExplainOnly bool                 `json:"explain"`
}

func NewSelectStatement() *SelectStatement {
	return &SelectStatement{
		Limit: -1,
	}
}

func (this *SelectStatement) GetFromAliases() []string {
	rv := make([]string, 0)
	for _, f := range this.Froms {
		rv = append(rv, f.As)
	}
	return rv
}

func (this *SelectStatement) GetFroms() []*From {
	return this.Froms
}

func (this *SelectStatement) GetWhere() Expression {
	return this.Where
}

func (this *SelectStatement) GetOrderBy() []*SortExpression {
	return this.OrderBy
}

func (this *SelectStatement) GetLimit() int {
	return this.Limit
}

func (this *SelectStatement) GetOffset() int {
	return this.Offset
}

func (this *SelectStatement) IsDistinct() bool {
	return this.Distinct
}

func (this *SelectStatement) GetResultExpressionList() ResultExpressionList {
	return this.Select
}

func (this *SelectStatement) SetExplainOnly(explainOnly bool) {
	this.ExplainOnly = explainOnly
}

func (this *SelectStatement) IsExplainOnly() bool {
	return this.ExplainOnly
}

func (this *SelectStatement) VerifySemantics() error {
	//check for duplicate aliases
	err := this.GetResultExpressionList().CheckForDuplicateAliases()
	if err != nil {
		return err
	}

	// now apply default naming function
	this.GetResultExpressionList().AssignDefaultNames()

	// fix up all the froms (FIXME need to refactor this into FROM/OVER clenaup)
	if this.Froms != nil && len(this.Froms) > 0 {
		this.Froms[0].ConvertToBucketFrom()
		for _, from := range this.Froms {
			from.GenerateAlias()
		}
	}

	// gather the list of aliases
	aliases := this.GetFromAliases()
	defaultAlias := ""
	if len(aliases) == 1 {
		defaultAlias = aliases[0]
	}

	// verify formal notations
	err = this.Select.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return err
	}
	if this.Where != nil {
		newwhere, err := this.Where.VerifyFormalNotation(aliases, defaultAlias)
		if err != nil {
			return err
		}
		if newwhere != nil {
			this.Where = newwhere
		}
	}
	// verify the order by
	for _, orderExpr := range this.OrderBy {
		if orderExpr.Expr != nil {
			neworderexpr, err := orderExpr.Expr.VerifyFormalNotation(aliases, defaultAlias)
			if err != nil {
				return err
			}
			if neworderexpr != nil {
				orderExpr.Expr = neworderexpr
			}
		}
	}

	// verify the projection
	err = this.Select.Validate()
	if err != nil {
		return err
	}

	// verify the where
	if this.Where != nil {
		err = this.Where.Validate()
		if err != nil {
			return err
		}
	}

	// verify the order by
	for _, orderExpr := range this.OrderBy {
		if orderExpr.Expr != nil {
			err := orderExpr.Expr.Validate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
