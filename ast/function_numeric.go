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

func init() {
	registerSystemFunction(&FunctionCeil{})
	registerSystemFunction(&FunctionFloor{})
	registerSystemFunction(&FunctionRound{})
	registerSystemFunction(&FunctionTrunc{})
}

type FunctionCeil struct{}

func (this *FunctionCeil) Name() string {
	return "CEIL"
}

func (this *FunctionCeil) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

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

func (this *FunctionCeil) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionCeil) IsAggregate() bool {
	return false
}

type FunctionFloor struct{}

func (this *FunctionFloor) Name() string {
	return "FLOOR"
}

func (this *FunctionFloor) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

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

func (this *FunctionFloor) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionFloor) IsAggregate() bool {
	return false
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

type FunctionRound struct{}

func (this *FunctionRound) Name() string {
	return "ROUND"
}

func (this *FunctionRound) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)
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
	if len(arguments) > 1 {
		// evaluate the second argument
		pv, err := arguments[1].Expr.Evaluate(item)
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

func (this *FunctionRound) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 2)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionRound) IsAggregate() bool {
	return false
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

type FunctionTrunc struct{}

func (this *FunctionTrunc) Name() string {
	return "TRUNC"
}

func (this *FunctionTrunc) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)
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
	if len(arguments) > 1 {
		// evaluate the second argument
		pv, err := arguments[1].Expr.Evaluate(item)

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

func (this *FunctionTrunc) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 2)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionTrunc) IsAggregate() bool {
	return false
}
