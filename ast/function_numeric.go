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
	"fmt"
	"math"
	"strconv"

	"github.com/couchbaselabs/tuqtng/query"
)

func init() {
	registerSystemFunction("CEIL", &FunctionCeil{})
	registerSystemFunction("FLOOR", &FunctionFloor{})
	registerSystemFunction("ROUND", &FunctionRound{})
}

type FunctionCeil struct{}

func (this *FunctionCeil) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *query.Undefined:
			// undefined returns null
			return nil, nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	switch av := av.(type) {
	case float64:
		return math.Ceil(av), nil
	default:
		return nil, nil
	}
}

func (this *FunctionCeil) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the CEIL() function expects a single argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the CEIL() function does not support *")
	}
	return nil
}

type FunctionFloor struct{}

func (this *FunctionFloor) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *query.Undefined:
			// undefined returns null
			return nil, nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	switch av := av.(type) {
	case float64:
		return math.Floor(av), nil
	default:
		return nil, nil
	}
}

func (this *FunctionFloor) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the FLOOR() function expects a single argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the FLOOR() function does not support *")
	}
	return nil
}

func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}

type FunctionRound struct{}

func (this *FunctionRound) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	precision := 0
	if len(arguments) > 1 {
		// evaluate the second argument
		pv, err := arguments[1].Expr.Evaluate(item)

		// we need precision to be an integer
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// undefined returns null
				return nil, nil
			default:
				// any other error return to caller
				return nil, err
			}
		}

		switch pv := pv.(type) {
		case float64:
			precision = int(pv)
		default:
			// FIXME log warning here?
			return nil, nil
		}
	}

	// the spec defines this functin to ONLY operate on numeric values
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *query.Undefined:
			// undefined returns null
			return nil, nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	switch av := av.(type) {
	case float64:
		return RoundFloat(av, precision), nil
	default:
		return nil, nil
	}
}

func (this *FunctionRound) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 || len(arguments) > 2 {
		return fmt.Errorf("the ROUND() function expects either one or two arguments")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the ROUND() function does not support *")
	}
	return nil
}
