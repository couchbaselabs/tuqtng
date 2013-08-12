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

// this ExpressionVisitor searches the expression
// for any aggregate function calls embedded inside
// so that the appropriate calculations can be
// made during the grouping phase
type ExpressionAggregateFinder struct {
	aggregates ExpressionList
}

func NewExpressionAggregateFinder() *ExpressionAggregateFinder {
	return &ExpressionAggregateFinder{
		aggregates: make(ExpressionList, 0),
	}
}

func (this *ExpressionAggregateFinder) GetAggregates() ExpressionList {
	return this.aggregates
}

func (this *ExpressionAggregateFinder) Visit(expr Expression) (Expression, error) {
	switch expr := expr.(type) {
	case *FunctionCall:
		if expr.IsAggregate() {
			this.aggregates = append(this.aggregates, expr)
		}
		return expr, nil
	default:
		return VisitChildren(this, expr)
	}
}
