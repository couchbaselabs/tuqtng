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
	"github.com/couchbaselabs/dparval"
)

type FunctionCallLength struct {
	FunctionCall
}

func NewFunctionCallLength(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallLength{
		FunctionCall{
			Type:     "function",
			Name:     "LENGTH",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallLength) Copy() Expression {
	return &FunctionCallLength{
		FunctionCall{
			Type:     "function",
			Name:     "LENGTH",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallLength) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec does not define it to operate on missing, so return null
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	// return the length only if the operand is a string
	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			return dparval.NewValue(float64(len(avalue))), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallLength) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

//array_length()
type FunctionCallArrayLength struct {
	FunctionCall
}

func NewFunctionCallArrayLength(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallArrayLength{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_LENGTH",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallArrayLength) Copy() Expression {
	return &FunctionCallArrayLength{
		FunctionCall{
			Type:     "function",
			Name:     "ARRAY_LENGTH",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallArrayLength) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec does not define it to operate on missing, so return null
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	// return the length only if the operand is an array
	if av.Type() == dparval.ARRAY {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case []interface{}:
			return dparval.NewValue(float64(len(avalue))), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallArrayLength) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

//object_length()
type FunctionCallObjectLength struct {
	FunctionCall
}

func NewFunctionCallObjectLength(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallObjectLength{
		FunctionCall{
			Type:     "function",
			Name:     "OBJECT_LENGTH",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallObjectLength) Copy() Expression {
	return &FunctionCallObjectLength{
		FunctionCall{
			Type:     "function",
			Name:     "OBJECT_LENGTH",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallObjectLength) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec does not define it to operate on missing, so return null
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	// return the length only if the operand is an array
	if av.Type() == dparval.OBJECT {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case map[string]interface{}:
			return dparval.NewValue(float64(len(avalue))), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallObjectLength) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

//poly_length() operate on any type
type FunctionCallPolyLength struct {
	FunctionCall
}

func NewFunctionCallPolyLength(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallPolyLength{
		FunctionCall{
			Type:     "function",
			Name:     "POLY_LENGTH",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallPolyLength) Copy() Expression {
	return &FunctionCallPolyLength{
		FunctionCall{
			Type:     "function",
			Name:     "POLY_LENGTH",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallPolyLength) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec does not define it to operate on missing, so return null
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	// return the length only if the operand is an array
	if av.Type() == dparval.OBJECT || av.Type() == dparval.ARRAY || av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			return dparval.NewValue(float64(len(avalue))), nil
		case []interface{}:
			return dparval.NewValue(float64(len(avalue))), nil
		case map[string]interface{}:
			return dparval.NewValue(float64(len(avalue))), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallPolyLength) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
