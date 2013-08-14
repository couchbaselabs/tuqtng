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
type ExpressionFunctionalDependencyChecker struct {
	Dependencies ExpressionList
}

func NewExpressionFunctionalDependencyChecker(deps ExpressionList) *ExpressionFunctionalDependencyChecker {
	return &ExpressionFunctionalDependencyChecker{
		Dependencies: deps,
	}
}

func (this *ExpressionFunctionalDependencyChecker) Visit(expr Expression) (Expression, error) {
	// first let scalar literals pass, they do not depend on anything
	switch expr.(type) {
	case *LiteralNull:
		return nil, nil
	case *LiteralBool:
		return nil, nil
	case *LiteralNumber:
		return nil, nil
	case *LiteralString:
		return nil, nil
	case AggregateFunctionCallExpression:
		return nil, nil
	}

	// next see if this expression is directly equivalent
	// to one of the provided dependencies
	e := this.Dependencies.EquivalentTo(expr)
	if e != nil {
		return nil, nil
	}

	// if this expression has no dependencies
	// then this expression was unsatisfied
	deps := expr.Dependencies()
	if len(deps) == 0 {
		return nil, fmt.Errorf("The expression %v is not satisfied by these dependencies", expr)
	}

	// otherwise, we need to satisfy ALL of its dependencies
	// for ANY/ALL we need a custom checker
	// for all others we continue to use *this
	checker := this.getAppropriateCheckerForDeps(expr)

	for _, dep := range deps {
		_, err := dep.Accept(checker)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (this *ExpressionFunctionalDependencyChecker) getAppropriateCheckerForDeps(expr Expression) *ExpressionFunctionalDependencyChecker {
	switch expr := expr.(type) {
	case CollectionOperatorExpression:
		// add the collection AS value
		// as a new satisifed dependency before checking
		newDeps := make(ExpressionList, len(this.Dependencies)+1)
		for i, dep := range this.Dependencies {
			newDeps[i] = dep
		}
		newDeps[len(this.Dependencies)] = NewProperty(expr.GetAs())
		colDepChecker := NewExpressionFunctionalDependencyChecker(newDeps)
		return colDepChecker
	default:
		return this
	}
}
