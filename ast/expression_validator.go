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

type ExpressionValidator struct {
	insideAggregate bool
	allowAggregates bool
}

func NewExpressionValidator() *ExpressionValidator {
	return &ExpressionValidator{
		allowAggregates: true,
	}
}

func NewExpressionValidatorNoAggregates() *ExpressionValidator {
	return &ExpressionValidator{
		allowAggregates: false,
	}
}

func (this *ExpressionValidator) Visit(expr Expression) (Expression, error) {
	switch expr := expr.(type) {
	case *FunctionCall:
		if expr.IsAggregate() {
			if !this.allowAggregates {
				return expr, fmt.Errorf("Aggregate function not allowed here")
			}
			if this.insideAggregate {
				return expr, fmt.Errorf("Cannot use aggregate function inside another aggregate function")
			}
			this.insideAggregate = true
		}
		return this.ValidateFunctionCall(expr)
	default:
		return VisitChildren(this, expr)
	}
}

func (this *ExpressionValidator) ValidateFunctionCall(expr *FunctionCall) (Expression, error) {
	// validate the function name is valid
	functionImpl := SystemFunctionRegistry[expr.Name]
	if functionImpl == nil {
		return expr, fmt.Errorf("no system function named %v registered", expr.Name)
	}

	// ask the implementation to validate the arguments
	operError := functionImpl.Validate(expr.Operands)
	if operError != nil {
		return expr, operError
	}

	e, err := VisitChildren(this, expr)
	this.insideAggregate = false
	return e, err
}
