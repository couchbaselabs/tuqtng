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
	"github.com/couchbaselabs/dparval"
)

type Expression interface {
	fmt.Stringer

	Evaluate(item *dparval.Value) (*dparval.Value, error)

	// Is this Expresion equivalent to that Expression?
	EquivalentTo(Expression) bool

	// A list of other Expressions up on which this depends
	Dependencies() ExpressionList

	// Vistor Pattern
	Accept(ExpressionVisitor) (Expression, error)
}

type OperatorExpression interface {
	Operator() string
}

type UnaryOperatorExpression interface {
	OperatorExpression
	GetOperand() Expression
	SetOperand(Expression)
}

type BinaryOperatorExpression interface {
	OperatorExpression
	GetLeft() Expression
	GetRight() Expression
	SetLeft(Expression)
	SetRight(Expression)
}

type NaryOperatorExpression interface {
	OperatorExpression
	GetOperands() ExpressionList
	SetOperands(ExpressionList)
	SetOperand(int, Expression)
}

type CollectionOperatorExpression interface {
	OperatorExpression
	GetOver() Expression
	GetCondition() Expression
	GetAs() string
	SetOver(Expression)
	SetCondition(Expression)
	SetAs(string)
}

type FunctionCallExpression interface {
	GetName() string
	GetOperands() FunctionArgExpressionList
	SetOperands(FunctionArgExpressionList)
	ValidateArity() error
	ValidateStars() error
}

type AggregateFunctionCallExpression interface {
	FunctionCallExpression
	UpdateAggregate(group *dparval.Value, item *dparval.Value) error
	DefaultAggregate(group *dparval.Value) error
	Key() string
}
