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

type ExpressionVisitor interface {
	Visit(Expression) (Expression, error)
}

// a utility function giving the default behavior
// of visiting a nodes children
func VisitChildren(v ExpressionVisitor, e Expression) (Expression, error) {
	var err error

	switch expr := e.(type) {
	//case
	case *CaseOperator:
		for _, wt := range expr.WhenThens {
			wt.When, err = wt.When.Accept(v)
			if err != nil {
				return expr, err
			}
			wt.Then, err = wt.Then.Accept(v)
			if err != nil {
				return expr, err
			}
		}
		if expr.Else != nil {
			expr.Else, err = expr.Else.Accept(v)
			if err != nil {
				return expr, err
			}
		}
	// handle all collection operators
	case CollectionOperatorExpression:
		newOver, err := expr.GetOver().Accept(v)
		if err != nil {
			return e, err
		}
		expr.SetOver(newOver)
		newCondition, err := expr.GetCondition().Accept(v)
		if err != nil {
			return e, err
		}
		expr.SetCondition(newCondition)

	// handle all binary operators
	case BinaryOperatorExpression:
		newLeft, err := expr.GetLeft().Accept(v)
		if err != nil {
			return e, err
		}
		expr.SetLeft(newLeft)
		newRight, err := expr.GetRight().Accept(v)
		if err != nil {
			return e, err
		}
		expr.SetRight(newRight)

	// handle all unary operators
	case UnaryOperatorExpression:
		newOper, err := expr.GetOperand().Accept(v)
		if err != nil {
			return e, err
		}
		expr.SetOperand(newOper)

	// handle all nary operators
	case NaryOperatorExpression:
		for i, oper := range expr.GetOperands() {
			newoper, err := oper.Accept(v)
			if err != nil {
				return e, err
			}
			expr.SetOperand(i, newoper)
		}

	// function
	case *FunctionCall:
		for _, arg := range expr.Operands {
			if arg.Expr != nil {
				arg.Expr, err = arg.Expr.Accept(v)
				if err != nil {
					return expr, err
				}
			}
		}
	// literal
	case *LiteralNull:
	case *LiteralBool:
	case *LiteralNumber:
	case *LiteralString:
	case *LiteralArray:
		for i, val := range expr.Val {
			expr.Val[i], err = val.Accept(v)
			if err != nil {
				return expr, err
			}
		}
	case *LiteralObject:
		for k, val := range expr.Val {
			expr.Val[k], err = val.Accept(v)
			if err != nil {
				return expr, err
			}
		}
	// member
	case *DotMemberOperator:
		var newProp Expression
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		newProp, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
		newProperty, ok := newProp.(*Property)
		if ok {
			expr.Right = newProperty
		}
	case *BracketMemberOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	// property
	case *Property:
	}
	return e, nil
}
