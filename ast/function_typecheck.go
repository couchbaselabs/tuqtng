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

type FunctionCallTypeName struct {
	FunctionCall
}

func NewFunctionCallTypeName(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallTypeName{
		FunctionCall{
			Type:     "function",
			Name:     "TYPE_NAME",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallTypeName) Copy() Expression {
	return &FunctionCallTypeName{
		FunctionCall{
			Type:     "function",
			Name:     "TYPE_NAME",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallTypeName) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns "missing"
			return dparval.NewValue("missing"), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	switch av.Type() {
	case dparval.NUMBER:
		return dparval.NewValue("number"), nil
	case dparval.STRING:
		return dparval.NewValue("string"), nil
	case dparval.BOOLEAN:
		return dparval.NewValue("boolean"), nil
	case dparval.ARRAY:
		return dparval.NewValue("array"), nil
	case dparval.OBJECT:
		return dparval.NewValue("object"), nil
	case dparval.NULL:
		return dparval.NewValue("null"), nil
	default:
		return dparval.NewValue(nil), nil
	}
}

func (this *FunctionCallTypeName) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIsNum struct {
	FunctionCall
}

func NewFunctionCallIsNum(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIsNum{
		FunctionCall{
			Type:     "function",
			Name:     "IS_NUM",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsNum) Copy() Expression {
	return &FunctionCallIsNum{
		FunctionCall{
			Type:     "function",
			Name:     "IS_NUM",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsNum) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

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

	switch av.Type() {
	case dparval.NULL:
		return av, nil
	default:
		return dparval.NewValue(av.Type() == dparval.NUMBER), nil
	}
}

func (this *FunctionCallIsNum) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIsStr struct {
	FunctionCall
}

func NewFunctionCallIsStr(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIsStr{
		FunctionCall{
			Type:     "function",
			Name:     "IS_STR",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsStr) Copy() Expression {
	return &FunctionCallIsStr{
		FunctionCall{
			Type:     "function",
			Name:     "IS_STR",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsStr) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

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

	switch av.Type() {
	case dparval.NULL:
		return av, nil
	default:
		return dparval.NewValue(av.Type() == dparval.STRING), nil
	}
}

func (this *FunctionCallIsStr) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIsBool struct {
	FunctionCall
}

func NewFunctionCallIsBool(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIsBool{
		FunctionCall{
			Type:     "function",
			Name:     "IS_BOOL",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsBool) Copy() Expression {
	return &FunctionCallIsBool{
		FunctionCall{
			Type:     "function",
			Name:     "IS_BOOL",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsBool) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

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

	switch av.Type() {
	case dparval.NULL:
		return av, nil
	default:
		return dparval.NewValue(av.Type() == dparval.BOOLEAN), nil
	}
}

func (this *FunctionCallIsBool) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIsAtom struct {
	FunctionCall
}

func NewFunctionCallIsAtom(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIsAtom{
		FunctionCall{
			Type:     "function",
			Name:     "IS_ATOM",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsAtom) Copy() Expression {
	return &FunctionCallIsAtom{
		FunctionCall{
			Type:     "function",
			Name:     "IS_ATOM",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsAtom) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

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

	switch av.Type() {
	case dparval.NULL:
		return av, nil
	case dparval.NUMBER, dparval.STRING, dparval.BOOLEAN:
		return dparval.NewValue(true), nil
	default:
		return dparval.NewValue(false), nil
	}
}

func (this *FunctionCallIsAtom) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIsArray struct {
	FunctionCall
}

func NewFunctionCallIsArray(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIsArray{
		FunctionCall{
			Type:     "function",
			Name:     "IS_ARRAY",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsArray) Copy() Expression {
	return &FunctionCallIsArray{
		FunctionCall{
			Type:     "function",
			Name:     "IS_ARRAY",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsArray) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

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

	switch av.Type() {
	case dparval.NULL:
		return av, nil
	default:
		return dparval.NewValue(av.Type() == dparval.ARRAY), nil
	}
}

func (this *FunctionCallIsArray) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIsObj struct {
	FunctionCall
}

func NewFunctionCallIsObj(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIsObj{
		FunctionCall{
			Type:     "function",
			Name:     "IS_OBJ",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsObj) Copy() Expression {
	return &FunctionCallIsObj{
		FunctionCall{
			Type:     "function",
			Name:     "IS_OBJ",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallIsObj) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

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

	switch av.Type() {
	case dparval.NULL:
		return av, nil
	default:
		return dparval.NewValue(av.Type() == dparval.OBJECT), nil
	}
}

func (this *FunctionCallIsObj) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
