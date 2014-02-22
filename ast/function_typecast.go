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
	"bytes"
	"strconv"

	"github.com/couchbaselabs/dparval"
)

type FunctionCallToNum struct {
	FunctionCall
}

func NewFunctionCallToNum(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallToNum{
		FunctionCall{
			Type:     "function",
			Name:     "TO_NUM",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToNum) Copy() Expression {
	return &FunctionCallToNum{
		FunctionCall{
			Type:     "function",
			Name:     "TO_NUM",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToNum) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	case dparval.NUMBER, dparval.NULL:
		return av, nil
	case dparval.BOOLEAN, dparval.STRING:
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case bool:
			val := 0.0
			if avalue {
				val = 1.0
			}
			return dparval.NewValue(val), nil
		case string:
			val, err := strconv.ParseFloat(avalue, 64)
			if err == nil {
				return dparval.NewValue(val), nil
			}
		}
	}

	return dparval.NewValue(nil), nil
}

func (this *FunctionCallToNum) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallToStr struct {
	FunctionCall
}

func NewFunctionCallToStr(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallToStr{
		FunctionCall{
			Type:     "function",
			Name:     "TO_STR",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToStr) Copy() Expression {
	return &FunctionCallToStr{
		FunctionCall{
			Type:     "function",
			Name:     "TO_STR",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToStr) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	case dparval.STRING, dparval.NULL:
		return av, nil
	}

	val := av.Value()

	switch val := val.(type) {
	case float64:
		return dparval.NewValue(strconv.FormatFloat(val, 'g', -1, 64)), nil
	case bool:
		return dparval.NewValue(strconv.FormatBool(val)), nil
	case []interface{}, map[string]interface{}:
		var buf bytes.Buffer
		err = toStr(val, &buf)
		if err != nil {
			return nil, err
		}
		return dparval.NewValue(buf.String()), nil
	}

	return dparval.NewValue(nil), nil
}

func toStr(val interface{}, buf *bytes.Buffer) error {
	switch val := val.(type) {
	case string:
		buf.WriteRune('"')
		buf.WriteString(val)
		buf.WriteRune('"')
	case float64:
		buf.WriteString(strconv.FormatFloat(val, 'g', -1, 64))
	case bool:
		buf.WriteString(strconv.FormatBool(val))
	case nil:
		buf.WriteString("null")
	case []interface{}:
		buf.WriteRune('[')
		for i, v := range val {
			if i > 0 {
				buf.WriteString(", ")
			}
			err := toStr(v, buf)
			if err != nil {
				return err
			}
		}
		buf.WriteRune(']')
	case map[string]interface{}:
		buf.WriteRune('{')
		next := false
		for k, v := range val {
			if next {
				buf.WriteString(", ")
			}
			next = true
			buf.WriteRune('"')
			buf.WriteString(k)
			buf.WriteString("\": ")
			err := toStr(v, buf)
			if err != nil {
				return err
			}
		}
		buf.WriteRune('}')
	}

	return nil
}

func (this *FunctionCallToStr) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallToBool struct {
	FunctionCall
}

func NewFunctionCallToBool(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallToBool{
		FunctionCall{
			Type:     "function",
			Name:     "TO_BOOL",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToBool) Copy() Expression {
	return &FunctionCallToBool{
		FunctionCall{
			Type:     "function",
			Name:     "TO_BOOL",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToBool) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	case dparval.BOOLEAN, dparval.NULL:
		return av, nil
	case dparval.NUMBER, dparval.STRING, dparval.ARRAY, dparval.OBJECT:
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case float64:
			return dparval.NewValue(avalue != 0.0), nil
		case string:
			return dparval.NewValue(len(avalue) > 0), nil
		case []interface{}:
			return dparval.NewValue(len(avalue) > 0), nil
		case map[string]interface{}:
			return dparval.NewValue(len(avalue) > 0), nil
		}
	}

	return dparval.NewValue(nil), nil
}

func (this *FunctionCallToBool) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallToAtom struct {
	FunctionCall
}

func NewFunctionCallToAtom(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallToAtom{
		FunctionCall{
			Type:     "function",
			Name:     "TO_ATOM",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToAtom) Copy() Expression {
	return &FunctionCallToAtom{
		FunctionCall{
			Type:     "function",
			Name:     "TO_ATOM",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToAtom) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	case dparval.STRING, dparval.BOOLEAN, dparval.NUMBER, dparval.NULL:
		return av, nil
	case dparval.ARRAY, dparval.OBJECT:
		return dparval.NewValue(toAtom(av.Value())), nil
	}

	return dparval.NewValue(nil), nil
}

func toAtom(val interface{}) interface{} {
	switch val := val.(type) {
	case []interface{}:
		if len(val) != 1 {
			return nil
		}
		return toAtom(val[0])
	case map[string]interface{}:
		if len(val) != 1 {
			return nil
		}
		for _, v := range val {
			return toAtom(v)
		}
	}

	return val
}

func (this *FunctionCallToAtom) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallToArray struct {
	FunctionCall
}

func NewFunctionCallToArray(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallToArray{
		FunctionCall{
			Type:     "function",
			Name:     "TO_ARRAY",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToArray) Copy() Expression {
	return &FunctionCallToArray{
		FunctionCall{
			Type:     "function",
			Name:     "TO_ARRAY",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallToArray) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
	case dparval.ARRAY, dparval.NULL:
		return av, nil
	case dparval.NUMBER, dparval.STRING, dparval.BOOLEAN, dparval.OBJECT:
		return dparval.NewValue([]interface{}{av.Value()}), nil
	}

	return dparval.NewValue(nil), nil
}

func (this *FunctionCallToArray) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
