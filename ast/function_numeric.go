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
	"math"

	"github.com/couchbaselabs/dparval"
)

type FunctionCallCeil struct {
	FunctionCall
}

func NewFunctionCallCeil(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallCeil{
		FunctionCall{
			Type:     "function",
			Name:     "CEIL",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallCeil) Copy() Expression {
	return &FunctionCallCeil{
		FunctionCall{
			Type:     "function",
			Name:     "CEIL",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallCeil) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
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

	if av.Type() == dparval.NUMBER {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case float64:
			return dparval.NewValue(math.Ceil(avalue)), nil
		}
	}

	return dparval.NewValue(nil), nil
}

func (this *FunctionCallCeil) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallFloor struct {
	FunctionCall
}

func NewFunctionCallFloor(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallFloor{
		FunctionCall{
			Type:     "function",
			Name:     "FLOOR",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallFloor) Copy() Expression {
	return &FunctionCallFloor{
		FunctionCall{
			Type:     "function",
			Name:     "FLOOR",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallFloor) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
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

	if av.Type() == dparval.NUMBER {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case float64:
			return dparval.NewValue(math.Floor(avalue)), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallFloor) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

func RoundFloat(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

type FunctionCallRound struct {
	FunctionCall
}

func NewFunctionCallRound(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallRound{
		FunctionCall{
			Type:     "function",
			Name:     "ROUND",
			Operands: operands,
			minArgs:  1,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallRound) Copy() Expression {
	return &FunctionCallRound{
		FunctionCall{
			Type:     "function",
			Name:     "ROUND",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallRound) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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

	precision := 0
	if len(this.Operands) > 1 {
		// evaluate the second argument
		pv, err := this.Operands[1].Expr.Evaluate(item)
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

		// we need precision to be an integer
		if pv.Type() == dparval.NUMBER {
			pvalue := pv.Value()
			switch pvalue := pvalue.(type) {
			case float64:
				precision = int(pvalue)
			}
		} else {
			// FIXME log warning here?
			return dparval.NewValue(nil), nil
		}
	}

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if av.Type() == dparval.NUMBER {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case float64:
			return dparval.NewValue(RoundFloat(avalue, precision)), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallRound) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

func TruncateFloat(x float64, prec int) float64 {

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	//_, frac := math.Modf(intermed)
	rounder = math.Floor(intermed)

	rv := rounder / pow
	return rv
}

type FunctionCallTrunc struct {
	FunctionCall
}

func NewFunctionCallTrunc(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallTrunc{
		FunctionCall{
			Type:     "function",
			Name:     "TRUNC",
			Operands: operands,
			minArgs:  1,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallTrunc) Copy() Expression {
	return &FunctionCallTrunc{
		FunctionCall{
			Type:     "function",
			Name:     "TRUNC",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallTrunc) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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

	precision := 0
	if len(this.Operands) > 1 {
		// evaluate the second argument
		pv, err := this.Operands[1].Expr.Evaluate(item)

		// we need precision to be an integer
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

		if pv.Type() == dparval.NUMBER {
			pvalue := pv.Value()
			switch pvalue := pvalue.(type) {
			case float64:
				precision = int(pvalue)
			}
		} else {
			// FIXME log warning here?
			return dparval.NewValue(nil), nil
		}

	}

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if av.Type() == dparval.NUMBER {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case float64:
			return dparval.NewValue(TruncateFloat(avalue, precision)), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallTrunc) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
