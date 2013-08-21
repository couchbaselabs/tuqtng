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

// this ExpressionVisitor will determine if the
// given expression can be satisfied by the
// specified dependencies, or if it instead
// has other dependencies
type ExpressionEquivalenceChecker struct {
	Expressions ExpressionList
}

func NewExpressionEquivalenceChecker(exprs ExpressionList) *ExpressionEquivalenceChecker {
	return &ExpressionEquivalenceChecker{
		Expressions: exprs,
	}
}

func (this *ExpressionEquivalenceChecker) Visit(expr Expression) (Expression, error) {

	// first if this expression is directly equivalent
	// to one of the provided expressions
	e := this.Expressions.EquivalentTo(expr)
	if e != nil {
		return e, nil
	}

	return nil, fmt.Errorf("Expression %v is not in the list", expr)
}
