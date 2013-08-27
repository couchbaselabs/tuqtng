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

// convert an expresion to Negation Normal Form
// http://en.wikipedia.org/wiki/Negation_normal_form
type ExpressionCNF struct {
}

func NewExpressionCNF() *ExpressionCNF {
	return &ExpressionCNF{}
}

func (this *ExpressionCNF) Visit(e Expression) (Expression, error) {

	// first ensure the children are in CNF
	e, err := VisitChildren(this, e)
	if err != nil {
		return e, err
	}

	// now some special handling for our logical connectors
	switch e := e.(type) {
	case *AndOperator:
		// new operands will be at least as long as before
		operands := make(ExpressionList, 0, len(e.Operands))
		// if any of our children are AND operators, roll them up into us
		for _, operand := range e.Operands {
			switch operand := operand.(type) {
			case *AndOperator:
				for _, inner := range operand.Operands {
					operands = append(operands, inner)
				}
			default:
				operands = append(operands, operand)
			}
		}
		return NewAndOperator(operands), nil
	case *OrOperator:

		// first we need to example all the operands
		// divide them into 2 groups
		// the AND operands, and everything else
		// also, if we encounter an OR operands, we roll them into us
		andOperands := make([]*AndOperator, 0)
		otherOperands := make(ExpressionList, 0)
		for _, operand := range e.Operands {
			switch operand := operand.(type) {
			case *AndOperator:
				andOperands = append(andOperands, operand)
			case *OrOperator:
				for _, inner := range operand.Operands {
					otherOperands = append(otherOperands, inner)
				}
			default:
				otherOperands = append(otherOperands, operand)
			}
		}

		if len(andOperands) == 0 {
			// if at this point there were no ANDs, bail out early
			return NewOrOperator(otherOperands), nil
		}

		// now we have to distribute the contents of the ANDs across the ORs
		// its important to note that as we go, we're building longer ORs
		// and so there is increasingly more work to do as we distribute
		resultPieces := []*OrOperator{NewOrOperator(otherOperands)}
		for _, andOperand := range andOperands {
			newResultPieces := []*OrOperator{}
			for _, innerAnd := range andOperand.Operands {
				for _, resultPiece := range resultPieces {
					newResultOperands := make(ExpressionList, len(resultPiece.Operands))
					copy(newResultOperands, resultPiece.Operands)
					newResultOperands = append(newResultOperands, innerAnd)
					newResultPieces = append(newResultPieces, NewOrOperator(newResultOperands))
				}
			}
			// this is a key line, we're updating an element we iterate over
			// inside a more inner loop.  this is by design because as we
			// distribute the elements, there is more for the subsequent elements
			// to distribute over
			resultPieces = newResultPieces
		}

		resultOperands := make(ExpressionList, len(resultPieces))
		for i, v := range resultPieces {
			resultOperands[i] = v
		}
		return NewAndOperator(resultOperands), nil

	default:
		return e, nil
	}

}
