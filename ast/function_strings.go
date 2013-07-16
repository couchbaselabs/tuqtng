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
	"strings"

	"github.com/couchbaselabs/tuqtng/query"
)

func init() {
	registerSystemFunction("LOWER", &FunctionLower{})
	registerSystemFunction("UPPER", &FunctionUpper{})
	registerSystemFunction("LTRIM", &FunctionLTrim{})
	registerSystemFunction("RTRIM", &FunctionRTrim{})
	registerSystemFunction("TRIM", &FunctionTrim{})
}

type FunctionLower struct{}

func (this *FunctionLower) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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
	case string:
		return strings.ToLower(av), nil
	default:
		return nil, nil
	}
}

func (this *FunctionLower) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the LOWER() function expects a single argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the LOWER() function does not support *")
	}
	return nil
}

type FunctionUpper struct{}

func (this *FunctionUpper) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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
	case string:
		return strings.ToUpper(av), nil
	default:
		return nil, nil
	}
}

func (this *FunctionUpper) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the UPPER() function expects a single argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the UPPER() function does not support *")
	}
	return nil
}

type FunctionLTrim struct{}

func (this *FunctionLTrim) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	// evaluate the second argument
	cutlist, err := arguments[1].Expr.Evaluate(item)
	// the cut list MUST be a string, otherwise return null
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
	case string:
		switch cutlist := cutlist.(type) {
		case string:
			return strings.TrimLeft(av, cutlist), nil
		default:
			// FIXME warn that cutlist wasn't string?
			return nil, nil
		}
	default:
		return nil, nil
	}
}

func (this *FunctionLTrim) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the LTRIM() function expects two arguments")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the LTRIM() function does not support *")
	}
	return nil
}

type FunctionRTrim struct{}

func (this *FunctionRTrim) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	// evaluate the second argument
	cutlist, err := arguments[1].Expr.Evaluate(item)
	// the cut list MUST be a string, otherwise return null
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
	case string:
		switch cutlist := cutlist.(type) {
		case string:
			return strings.TrimRight(av, cutlist), nil
		default:
			// FIXME warn cutlist wasnt string?
			return nil, nil
		}
	default:
		return nil, nil
	}
}

func (this *FunctionRTrim) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the RTRIM() function expects two arguments")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the RTRIM() function does not support *")
	}
	return nil
}

type FunctionTrim struct{}

func (this *FunctionTrim) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	// evaluate the second argument
	cutlist, err := arguments[1].Expr.Evaluate(item)
	// the cut list MUST be a string, otherwise return null
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
	case string:
		switch cutlist := cutlist.(type) {
		case string:
			return strings.Trim(av, cutlist), nil
		default:
			// FIXME warn that cutlist wasnt string?
			return nil, nil
		}
	default:
		return nil, nil
	}
}

func (this *FunctionTrim) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the TRIM() function expects two arguments")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the TRIM() function does not support *")
	}
	return nil
}
