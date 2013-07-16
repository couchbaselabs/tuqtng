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

	"github.com/couchbaselabs/tuqtng/query"
)

func init() {
	registerSystemFunction("GREATEST", &FunctionGreatest{})
	registerSystemFunction("LEAST", &FunctionLeast{})
	registerSystemFunction("IFMISSING", &FunctionIfMissing{})
	registerSystemFunction("IFNULL", &FunctionIfNull{})
	registerSystemFunction("IFMISSINGORNULL", &FunctionIfMissingOrNull{})
	registerSystemFunction("MISSINGIF", &FunctionMissingIf{})
	registerSystemFunction("NULLIF", &FunctionNullIf{})
}

type FunctionGreatest struct{}

func (this *FunctionGreatest) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	var rv query.Value = nil

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// undefined doesn't change the result
			default:
				// any other error return to caller
				return nil, err
			}
		}

		// now compare this value with rv
		compres := CollateJSON(av, rv)
		if compres > 0 {
			rv = av
		}
	}

	return rv, nil
}

func (this *FunctionGreatest) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the GREATEST() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the GREATEST() function does not support *")
	}
	return nil
}

type FunctionLeast struct{}

func (this *FunctionLeast) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	var rv query.Value = nil
	first := true

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// undefined doesn't change the result
			default:
				// any other error return to caller
				return nil, err
			}
		}

		// now compare this value with rv
		compres := CollateJSON(av, rv)
		if compres < 0 || first {
			first = false
			rv = av
		}
	}

	return rv, nil
}

func (this *FunctionLeast) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the LEAST() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the LEAST() function does not support *")
	}
	return nil
}

type FunctionIfMissing struct{}

func (this *FunctionIfMissing) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
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
	return nil, nil
}

func (this *FunctionIfMissing) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the IFMISSING() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the IFMISSING() function does not support *")
	}
	return nil
}

type FunctionIfNull struct{}

func (this *FunctionIfNull) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// missing is NOT null, so return it
				return nil, err
			default:
				// any other error return to caller
				return nil, err
			}
		}

		if av != nil {
			// wasn't NULL, return the value
			return av, nil
		}
	}

	// if all values were NULL return NULL
	return nil, nil
}

func (this *FunctionIfNull) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the IFNULL() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the IFNULL() function does not support *")
	}
	return nil
}

type FunctionIfMissingOrNull struct{}

func (this *FunctionIfMissingOrNull) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// do not return missing
				continue
			default:
				// any other error return to caller
				return nil, err
			}
		}

		if av != nil {
			// wasn't NULL, return the value
			return av, nil
		}
	}

	// if all values were NULL or MISSING return NULL
	return nil, nil
}

func (this *FunctionIfMissingOrNull) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the IFMISSINGORNULL() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the IFMISSINGNULL() function does not support *")
	}
	return nil
}

type FunctionMissingIf struct{}

func (this *FunctionMissingIf) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	lav, lerr := arguments[0].Expr.Evaluate(item)
	if lerr != nil {
		switch lerr := lerr.(type) {
		case *query.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, lerr
		}
	}

	// next evaluate the second argument
	rav, rerr := arguments[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *query.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, rerr
		}
	}

	compres := CollateJSON(lav, rav)
	if compres == 0 {
		return nil, &query.Undefined{}
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionMissingIf) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the MISSINGIF() function expects exactly two argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the MISSINGIF() function does not support *")
	}
	return nil
}

type FunctionNullIf struct{}

func (this *FunctionNullIf) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	lav, lerr := arguments[0].Expr.Evaluate(item)
	if lerr != nil {
		switch lerr := lerr.(type) {
		case *query.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, lerr
		}
	}

	// next evaluate the second argument
	rav, rerr := arguments[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *query.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, rerr
		}
	}

	compres := CollateJSON(lav, rav)
	if compres == 0 {
		return nil, nil
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionNullIf) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the NULLIF() function expects exactly two argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the NULLIF() function does not support *")
	}
	return nil
}
