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

func (this *ExpressionValidator) Visit(e Expression) (Expression, error) {
	switch expr := e.(type) {
	case *FunctionCallUnknown:
		return e, fmt.Errorf("no system function named %s registered", expr.GetName())
	case AggregateFunctionCallExpression:
		if !this.allowAggregates {
			return e, fmt.Errorf("Aggregate function not allowed here")
		}
		if this.insideAggregate {
			return e, fmt.Errorf("Cannot use aggregate function inside another aggregate function")
		}
		this.insideAggregate = true
		err := this.ValidateFunctionCall(expr)
		this.insideAggregate = false
		return e, err
	case FunctionCallExpression:
		err := this.ValidateFunctionCall(expr)
		if err != nil {
			return e, err
		}
		return VisitChildren(this, e)
	default:
		return VisitChildren(this, e)
	}
}

func (this *ExpressionValidator) ValidateFunctionCall(expr FunctionCallExpression) error {
	err := expr.ValidateArity()
	if err != nil {
		return err
	}

	// an exception, COUNT allows *
	if expr.GetName() != "COUNT" {
		err = expr.ValidateStars()
		if err != nil {
			return err
		}
	}
	return nil
}
