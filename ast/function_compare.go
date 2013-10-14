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

type FunctionCallGreatest struct {
	FunctionCall
}

func NewFunctionCallGreatest(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallGreatest{
		FunctionCall{
			Type:     "function",
			Name:     "GREATEST",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallGreatest) Copy() Expression {
	return &FunctionCallGreatest{
		FunctionCall{
			Type:     "function",
			Name:     "GREATEST",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallGreatest) Evaluate(item *dparval.Value) (*dparval.Value, error) {

	var rv interface{} = nil

	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// undefined doesn't change the result
			default:
				// any other error return to caller
				return nil, err
			}
		}

		avalue := av.Value()

		// now compare this value with rv
		compres := CollateJSON(avalue, rv)
		if compres > 0 {
			rv = avalue
		}
	}

	return dparval.NewValue(rv), nil
}

func (this *FunctionCallGreatest) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallLeast struct {
	FunctionCall
}

func NewFunctionCallLeast(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallLeast{
		FunctionCall{
			Type:     "function",
			Name:     "LEAST",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallLeast) Copy() Expression {
	return &FunctionCallLeast{
		FunctionCall{
			Type:     "function",
			Name:     "LEAST",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallLeast) Evaluate(item *dparval.Value) (*dparval.Value, error) {

	var rv interface{} = nil
	first := true

	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// undefined doesn't change the result
			default:
				// any other error return to caller
				return nil, err
			}
		}

		avalue := av.Value()

		// now compare this value with rv
		compres := CollateJSON(avalue, rv)
		if compres < 0 || first {
			first = false
			rv = avalue
		}
	}

	return dparval.NewValue(rv), nil
}

func (this *FunctionCallLeast) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfMissing struct {
	FunctionCall
}

func NewFunctionCallIfMissing(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfMissing{
		FunctionCall{
			Type:     "function",
			Name:     "IFMISSING",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfMissing) Copy() Expression {
	return &FunctionCallIfMissing{
		FunctionCall{
			Type:     "function",
			Name:     "IFMISSING",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfMissing) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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

		// wasn't missing, return the value
		return av, nil
	}

	// if all values were MISSING return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfMissing) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfNull struct {
	FunctionCall
}

func NewFunctionCallIfNull(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfNull{
		FunctionCall{
			Type:     "function",
			Name:     "IFNULL",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfNull) Copy() Expression {
	return &FunctionCallIfNull{
		FunctionCall{
			Type:     "function",
			Name:     "IFNULL",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfNull) Evaluate(item *dparval.Value) (*dparval.Value, error) {

	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// missing is NOT null, so return it
				return nil, err // weird, but nil, err IS the value
			default:
				// any other error return to caller
				return nil, err
			}
		}

		if av.Type() != dparval.NULL {
			// wasn't NULL, return the value
			return av, nil
		}
	}

	// if all values were NULL return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfNull) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallIfMissingOrNull struct {
	FunctionCall
}

func NewFunctionCallIfMissingOrNull(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallIfMissingOrNull{
		FunctionCall{
			Type:     "function",
			Name:     "IFMISSINGORNULL",
			Operands: operands,
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfMissingOrNull) Copy() Expression {
	return &FunctionCallIfMissingOrNull{
		FunctionCall{
			Type:     "function",
			Name:     "IFMISSINGORNULL",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  -1,
		},
	}
}

func (this *FunctionCallIfMissingOrNull) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	for _, arg := range this.Operands {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// do not return missing
				continue
			default:
				// any other error return to caller
				return nil, err
			}
		}

		if av.Type() != dparval.NULL {
			// wasn't NULL, return the value
			return av, nil
		}
	}

	// if all values were NULL or MISSING return NULL
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallIfMissingOrNull) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallMissingIf struct {
	FunctionCall
}

func NewFunctionCallMissingIf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallMissingIf{
		FunctionCall{
			Type:     "function",
			Name:     "MISSINGIF",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallMissingIf) Copy() Expression {
	return &FunctionCallMissingIf{
		FunctionCall{
			Type:     "function",
			Name:     "MISSINGIF",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallMissingIf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
				return nil, &dparval.Undefined{}
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
		return nil, &dparval.Undefined{}
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionCallMissingIf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallNullIf struct {
	FunctionCall
}

func NewFunctionCallNullIf(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallNullIf{
		FunctionCall{
			Type:     "function",
			Name:     "NULLIF",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallNullIf) Copy() Expression {
	return &FunctionCallNullIf{
		FunctionCall{
			Type:     "function",
			Name:     "NULLIF",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallNullIf) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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
				return dparval.NewValue(nil), nil
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
		return dparval.NewValue(nil), nil
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionCallNullIf) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
