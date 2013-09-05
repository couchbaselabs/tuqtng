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

type FunctionCallIfNaN struct {
	FunctionCall
}

func NewFunctionCallIfNaN(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfNaN{
		FunctionCall{
			Type:     "function",
			Name:     "IFNAN",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfNaN) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			// missing is not NaN, so return it
			return nil, err
		}

		if av.Type() == dparval.NUMBER {
			num := av.Value().(float64)
			if !math.IsNaN(num) {
				return av, nil
			} else {
				continue
			}
		} else {
			return av, nil
		}
	}

	// if all values were NaN return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfNaN) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfPosInf struct {
	FunctionCall
}

func NewFunctionCallIfPosInf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfPosInf{
		FunctionCall{
			Type:     "function",
			Name:     "IFPOSINF",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfPosInf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			// missing is not NaN, so return it
			return nil, err
		}

		if av.Type() == dparval.NUMBER {
			num := av.Value().(float64)
			if !math.IsInf(num, 1) {
				return av, nil
			} else {
				continue
			}
		} else {
			return av, nil
		}
	}

	// if all values were +Infinity return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfPosInf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfNegInf struct {
	FunctionCall
}

func NewFunctionCallIfNegInf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfNegInf{
		FunctionCall{
			Type:     "function",
			Name:     "IFNEGINF",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfNegInf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			// missing is not NaN, so return it
			return nil, err
		}

		if av.Type() == dparval.NUMBER {
			num := av.Value().(float64)
			if !math.IsInf(num, -1) {
				return av, nil
			} else {
				continue
			}
		} else {
			return av, nil
		}
	}

	// if all values were -Infinity return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfNegInf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfInf struct {
	FunctionCall
}

func NewFunctionCallIfInf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfInf{
		FunctionCall{
			Type:     "function",
			Name:     "IFINF",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfInf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			// missing is not NaN, so return it
			return nil, err
		}

		if av.Type() == dparval.NUMBER {
			num := av.Value().(float64)
			if !math.IsInf(num, 1) && !math.IsInf(num, -1) {
				return av, nil
			} else {
				continue
			}
		} else {
			return av, nil
		}
	}

	// if all values were +/-Infinity return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfInf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfNaNOrInf struct {
	FunctionCall
}

func NewFunctionCallIfNaNOrInf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfNaNOrInf{
		FunctionCall{
			Type:     "function",
			Name:     "IFNANORINF",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfNaNOrInf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			// missing is not NaN, so return it
			return nil, err
		}

		if av.Type() == dparval.NUMBER {
			num := av.Value().(float64)
			if !math.IsNaN(num) && !math.IsInf(num, 1) && !math.IsInf(num, -1) {
				return av, nil
			} else {
				continue
			}
		} else {
			return av, nil
		}
	}

	// if all values were +/-Infinity or NaN return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfNaNOrInf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallFirstNum struct {
	FunctionCall
}

func NewFunctionCallFirstNum(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallFirstNum{
		FunctionCall{
			Type:     "function",
			Name:     "FIRSTNUM",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallFirstNum) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// do NOT return missing
				continue
			default:
				// any other error return to caller
				return nil, err
			}
		}

		if av.Type() == dparval.NUMBER {
			num := av.Value().(float64)
			if !math.IsNaN(num) && !math.IsInf(num, 1) && !math.IsInf(num, -1) {
				return av, nil
			} else {
				continue
			}
		} else {
			continue
		}
	}

	// if all values were +/-Infinity or NaN return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallFirstNum) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallNaNIf struct {
	FunctionCall
}

func NewFunctionCallNaNIf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallNaNIf{
		FunctionCall{
			Type:     "function",
			Name:     "NANIF",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallNaNIf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	lav, lerr := this.Operands[0].Expr.Evaluate(item)
	if lerr != nil {
		switch lerr := lerr.(type) {
		case *dparval.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, lerr
		}
	}

	// next evaluate the second argument
	rav, rerr := this.Operands[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *dparval.Undefined:
			if lerr != nil {
				// if lerr isnt null it must also be
				// undefined (all others returned)
				return dparval.NewValue(math.NaN()), nil
			}
		default:
			// any other error return to caller
			return nil, rerr
		}
	}

	if lerr != nil {
		// lav was MISSING, but rav wasn't
		// so return lav
		return lav, lerr
	}

	lavalue := lav.Value()
	ravalue := rav.Value()

	compres := CollateJSON(lavalue, ravalue)
	if compres == 0 {
		return dparval.NewValue(math.NaN()), nil
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionCallNaNIf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallPosInfIf struct {
	FunctionCall
}

func NewFunctionCallPosInfIf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallPosInfIf{
		FunctionCall{
			Type:     "function",
			Name:     "POSINFIF",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallPosInfIf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	lav, lerr := this.Operands[0].Expr.Evaluate(item)
	if lerr != nil {
		switch lerr := lerr.(type) {
		case *dparval.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, lerr
		}
	}

	// next evaluate the second argument
	rav, rerr := this.Operands[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *dparval.Undefined:
			if lerr != nil {
				// if lerr isnt null it must also be
				// undefined (all others returned)
				return dparval.NewValue(math.Inf(1)), nil
			}
		default:
			// any other error return to caller
			return nil, rerr
		}
	}

	if lerr != nil {
		// lav was MISSING, but rav wasn't
		// so return lav
		return lav, lerr
	}

	lavalue := lav.Value()
	ravalue := rav.Value()

	compres := CollateJSON(lavalue, ravalue)
	if compres == 0 {
		return dparval.NewValue(math.Inf(1)), nil
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionCallPosInfIf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallNegInfIf struct {
	FunctionCall
}

func NewFunctionCallNegInfIf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallNegInfIf{
		FunctionCall{
			Type:     "function",
			Name:     "NEGINFIF",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallNegInfIf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	lav, lerr := this.Operands[0].Expr.Evaluate(item)
	if lerr != nil {
		switch lerr := lerr.(type) {
		case *dparval.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, lerr
		}
	}

	// next evaluate the second argument
	rav, rerr := this.Operands[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *dparval.Undefined:
			if lerr != nil {
				// if lerr isnt null it must also be
				// undefined (all others returned)
				return dparval.NewValue(math.Inf(-1)), nil
			}
		default:
			// any other error return to caller
			return nil, rerr
		}
	}

	if lerr != nil {
		// lav was MISSING, but rav wasn't
		// so return lav
		return lav, lerr
	}

	lavalue := lav.Value()
	ravalue := rav.Value()

	compres := CollateJSON(lavalue, ravalue)
	if compres == 0 {
		return dparval.NewValue(math.Inf(-1)), nil
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionCallNegInfIf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
