//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

type ExpressionVisitor interface {
	Visit(Expression) (Expression, error)
}

// a utility function giving the default behavior
// of visiting a nodes children
func VisitChildren(v ExpressionVisitor, expr Expression) (Expression, error) {
	var err error
	switch expr := expr.(type) {
	//arithmetic
	case *PlusOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *SubtractOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *MultiplyOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *DivideOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *ModuloOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *ChangeSignOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
		}
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
	// collections
	case *CollectionAnyOperator:
		expr.Over, err = expr.Over.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Condition, err = expr.Condition.Accept(v)
		if err != nil {
			return expr, err
		}
	case *CollectionAllOperator:
		expr.Over, err = expr.Over.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Condition, err = expr.Condition.Accept(v)
		if err != nil {
			return expr, err
		}
	// comparison
	case *GreaterThanOperator:
		expr.BinaryComparisonOperator.Left, err = expr.BinaryComparisonOperator.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.BinaryComparisonOperator.Right, err = expr.BinaryComparisonOperator.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *GreaterThanOrEqualOperator:
		expr.BinaryComparisonOperator.Left, err = expr.BinaryComparisonOperator.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.BinaryComparisonOperator.Right, err = expr.BinaryComparisonOperator.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *LessThanOperator:
		expr.BinaryComparisonOperator.Left, err = expr.BinaryComparisonOperator.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.BinaryComparisonOperator.Right, err = expr.BinaryComparisonOperator.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *LessThanOrEqualOperator:
		expr.BinaryComparisonOperator.Left, err = expr.BinaryComparisonOperator.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.BinaryComparisonOperator.Right, err = expr.BinaryComparisonOperator.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *EqualToOperator:
		expr.BinaryComparisonOperator.Left, err = expr.BinaryComparisonOperator.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.BinaryComparisonOperator.Right, err = expr.BinaryComparisonOperator.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *NotEqualToOperator:
		expr.BinaryComparisonOperator.Left, err = expr.BinaryComparisonOperator.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.BinaryComparisonOperator.Right, err = expr.BinaryComparisonOperator.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *LikeOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *NotLikeOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	case *IsNullOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
		}
	case *IsNotNullOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
		}
	case *IsMissingOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
		}
	case *IsNotMissingOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
		}
	case *IsValuedOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
		}
	case *IsNotValuedOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
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
	// logical
	case *AndOperator:
		for i, oper := range expr.Operands {
			expr.Operands[i], err = oper.Accept(v)
			if err != nil {
				return expr, err
			}
		}
	case *OrOperator:
		for i, oper := range expr.Operands {
			expr.Operands[i], err = oper.Accept(v)
			if err != nil {
				return expr, err
			}
		}
	case *NotOperator:
		expr.Operand, err = expr.Operand.Accept(v)
		if err != nil {
			return expr, err
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
	// string
	case *StringConcatenateOperator:
		expr.Left, err = expr.Left.Accept(v)
		if err != nil {
			return expr, err
		}
		expr.Right, err = expr.Right.Accept(v)
		if err != nil {
			return expr, err
		}
	}
	return expr, nil
}
