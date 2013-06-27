//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

// Statement is the abstract representation of an UNQL statement
type Statement interface {
	GetResultExpressionList() ResultExpressionList
	GetFrom() From
	GetWhere() Expression
	GetOrderBy() []*SortExpression
	GetOffset() int
	GetLimit() int
	SetExplainOnly(bool)
	IsExplainOnly() bool
}

type SelectStatement struct {
	Select      ResultExpressionList `json:"select"`
	From        From                 `json:"from"`
	Where       Expression           `json:"where"`
	OrderBy     []*SortExpression    `json:"orderby"`
	Limit       int                  `json:"limit"`
	Offset      int                  `json:"offset"`
	ExplainOnly bool                 `json:"explain"`
}

func NewSelectStatement() *SelectStatement {
	return &SelectStatement{}
}

func (this *SelectStatement) GetFrom() From {
	return this.From
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

func (this *SelectStatement) GetResultExpressionList() ResultExpressionList {
	return this.Select
}

func (this *SelectStatement) SetExplainOnly(explainOnly bool) {
	this.ExplainOnly = explainOnly
}

func (this *SelectStatement) IsExplainOnly() bool {
	return this.ExplainOnly
}
