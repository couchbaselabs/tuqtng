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
type ExpressionNNF struct {
}

func NewExpressionNNF() *ExpressionNNF {
	return &ExpressionNNF{}
}

func (this *ExpressionNNF) Visit(e Expression) (Expression, error) {

	// first ensure the children are in NNF
	e, err := VisitChildren(this, e)
	if err != nil {
		return e, err
	}

	// now, if I'm a not, distribute myself
	switch e := e.(type) {
	case *NotOperator:
		return this.distributeNot(e.Operand), nil
	}

	return e, nil
}

func (this *ExpressionNNF) distributeNot(e Expression) Expression {
	switch e := e.(type) {
	case *NotOperator:
		// eliminiate double negation
		return e.Operand
	case *AndOperator:
		operands := make(ExpressionList, len(e.Operands))
		for i, operand := range e.Operands {
			operands[i] = this.distributeNot(operand)
		}
		return NewOrOperator(operands)
	case *OrOperator:
		operands := make(ExpressionList, len(e.Operands))
		for i, operand := range e.Operands {
			operands[i] = this.distributeNot(operand)
		}
		return NewAndOperator(operands)
	default:
		return NewNotOperator(e)
	}
}
