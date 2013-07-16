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
	registerSystemFunction("SUBSTR", &FunctionSubStr{})
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

type FunctionSubStr struct{}

func (this *FunctionSubStr) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

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

	position, err := arguments[1].Expr.Evaluate(item)
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

	var maxLen int = -1

	if len(arguments) == 3 {
		lenarg, err := arguments[1].Expr.Evaluate(item)
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

		switch lenarg := lenarg.(type) {
		case float64:
			maxLen = int(lenarg)
			// FIXME add checks for negative values here?
		default:
			return nil, nil
		}
	}

	// ensure that arg1 is a string
	switch av := av.(type) {
	case string:
		// ensure that arg2 is a number
		switch position := position.(type) {
		case float64:
			pos := int(position)

			//validate that pos is valid
			if pos < 0 || pos >= len(av) {
				// FIXME add warning for invalid pos?
				return nil, nil
			}

			if maxLen < 0 {
				// no end limit
				return av[pos:], nil
			} else {
				// validate that maxLen is valid
				endPos := pos + maxLen + 1
				if endPos < pos || endPos >= len(av) {
					// FIXME add warning for invalid max len?
					return nil, nil
				}
				return av[pos:endPos], nil
			}
		default:
			// FIXME warn that position wasn't int?
			return nil, nil
		}
	default:
		// FIXME warn that arg1 isnt a string?
		return nil, nil
	}

}

func (this *FunctionSubStr) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 2 || len(arguments) > 4 {
		return fmt.Errorf("the SUBSTR() function expects two or three arguments")
	}
	if arguments[0].Star == true || arguments[1].Star == true || (len(arguments) == 3 && arguments[2].Star == true) {
		return fmt.Errorf("the SUBSTR() function does not support *")
	}
	return nil
}
