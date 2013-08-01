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

	"github.com/mschoch/dparval"
)

func init() {
	registerSystemFunction(&FunctionGreatest{})
	registerSystemFunction(&FunctionLeast{})
	registerSystemFunction(&FunctionIfMissing{})
	registerSystemFunction(&FunctionIfNull{})
	registerSystemFunction(&FunctionIfMissingOrNull{})
	registerSystemFunction(&FunctionMissingIf{})
	registerSystemFunction(&FunctionNullIf{})
}

type FunctionGreatest struct{}

func (this *FunctionGreatest) Name() string {
	return "GREATEST"
}

func (this *FunctionGreatest) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {

	var rv interface{} = nil

	for _, arg := range arguments {
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

func (this *FunctionGreatest) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the GREATEST() function expects at least one argument")
	}
	return ValidateNoStars(this, arguments)
}

type FunctionLeast struct{}

func (this *FunctionLeast) Name() string {
	return "LEAST"
}

func (this *FunctionLeast) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {

	var rv interface{} = nil
	first := true

	for _, arg := range arguments {
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

func (this *FunctionLeast) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the LEAST() function expects at least one argument")
	}
	return ValidateNoStars(this, arguments)
}

type FunctionIfMissing struct{}

func (this *FunctionIfMissing) Name() string {
	return "IFMISSING"
}

func (this *FunctionIfMissing) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {

	for _, arg := range arguments {
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
	return dparval.NewNullValue(), nil
}

func (this *FunctionIfMissing) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the IFMISSING() function expects at least one argument")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionIfNull struct{}

func (this *FunctionIfNull) Name() string {
	return "IFNULL"
}

func (this *FunctionIfNull) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {

	for _, arg := range arguments {
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
	return dparval.NewNullValue(), nil
}

func (this *FunctionIfNull) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the IFNULL() function expects at least one argument")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionIfMissingOrNull struct{}

func (this *FunctionIfMissingOrNull) Name() string {
	return "IFMISSINGORNULL"
}

func (this *FunctionIfMissingOrNull) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {

	for _, arg := range arguments {
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
	return dparval.NewNullValue(), nil
}

func (this *FunctionIfMissingOrNull) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the IFMISSINGORNULL() function expects at least one argument")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionMissingIf struct{}

func (this *FunctionMissingIf) Name() string {
	return "MISSINGIF"
}

func (this *FunctionMissingIf) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	lav, lerr := arguments[0].Expr.Evaluate(item)
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
	rav, rerr := arguments[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *dparval.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, rerr
		}
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

func (this *FunctionMissingIf) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the MISSINGIF() function expects exactly two argument")
	}
	return ValidateNoStars(this, arguments)
}

type FunctionNullIf struct{}

func (this *FunctionNullIf) Name() string {
	return "NULLIF"
}

func (this *FunctionNullIf) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	lav, lerr := arguments[0].Expr.Evaluate(item)
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
	rav, rerr := arguments[1].Expr.Evaluate(item)
	if rerr != nil {
		switch rerr := rerr.(type) {
		case *dparval.Undefined:
			// do nothing yet
		default:
			// any other error return to caller
			return nil, rerr
		}
	}

	lavalue := lav.Value()
	ravalue := rav.Value()

	compres := CollateJSON(lavalue, ravalue)
	if compres == 0 {
		return dparval.NewNullValue(), nil
	}

	//otheriwse return the left arg
	return lav, nil
}

func (this *FunctionNullIf) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the NULLIF() function expects exactly two argument")
	}

	return ValidateNoStars(this, arguments)
}
