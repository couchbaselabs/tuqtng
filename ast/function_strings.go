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

	"github.com/mschoch/dparval"
)

func init() {
	registerSystemFunction(&FunctionLower{})
	registerSystemFunction(&FunctionUpper{})
	registerSystemFunction(&FunctionLTrim{})
	registerSystemFunction(&FunctionRTrim{})
	registerSystemFunction(&FunctionTrim{})
	registerSystemFunction(&FunctionSubStr{})
}

type FunctionLower struct{}

func (this *FunctionLower) Name() string {
	return "LOWER"
}

func (this *FunctionLower) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			return dparval.NewStringValue(strings.ToLower(avalue)), nil
		}
	}
	return dparval.NewNullValue(), nil
}

func (this *FunctionLower) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the LOWER() function expects a single argument")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionUpper struct{}

func (this *FunctionUpper) Name() string {
	return "UPPER"
}

func (this *FunctionUpper) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			return dparval.NewStringValue(strings.ToUpper(avalue)), nil
		}
	}
	return dparval.NewNullValue(), nil
}

func (this *FunctionUpper) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the UPPER() function expects a single argument")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionLTrim struct{}

func (this *FunctionLTrim) Name() string {
	return "LTRIM"
}

func (this *FunctionLTrim) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
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
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if av.Type() == dparval.STRING {
		if cutlist.Type() == dparval.STRING {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				cutvlal := cutlist.Value()
				switch cutvlal := cutvlal.(type) {
				case string:
					return dparval.NewStringValue(strings.TrimLeft(avalue, cutvlal)), nil
				}
			}
		}

	}
	// FIXME warn if cutlist wasnt string?
	return dparval.NewNullValue(), nil
}

func (this *FunctionLTrim) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the LTRIM() function expects two arguments")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionRTrim struct{}

func (this *FunctionRTrim) Name() string {
	return "RTRIM"
}

func (this *FunctionRTrim) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
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
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if av.Type() == dparval.STRING {
		if cutlist.Type() == dparval.STRING {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				cutval := cutlist.Value()
				switch cutval := cutval.(type) {
				case string:
					return dparval.NewStringValue(strings.TrimRight(avalue, cutval)), nil
				}
			}
		}
	}
	// FIXME warn if cutlist wasnt string?
	return dparval.NewNullValue(), nil
}

func (this *FunctionRTrim) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the RTRIM() function expects two arguments")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionTrim struct{}

func (this *FunctionTrim) Name() string {
	return "TRIM"
}

func (this *FunctionTrim) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
	// all other types result in NULL
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
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
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if av.Type() == dparval.STRING {
		if cutlist.Type() == dparval.STRING {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				cutval := cutlist.Value()
				switch cutval := cutval.(type) {
				case string:
					return dparval.NewStringValue(strings.Trim(avalue, cutval)), nil
				}
			}
		}
	}
	// FIXME warn if that cutlist wasnt string?
	return dparval.NewNullValue(), nil
}

func (this *FunctionTrim) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 2 {
		return fmt.Errorf("the TRIM() function expects two arguments")
	}

	return ValidateNoStars(this, arguments)
}

type FunctionSubStr struct{}

func (this *FunctionSubStr) Name() string {
	return "SUBSTR"
}

func (this *FunctionSubStr) Evaluate(item dparval.Value, arguments FunctionArgExpressionList) (dparval.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	position, err := arguments[1].Expr.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewNullValue(), nil
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
			case *dparval.Undefined:
				// undefined returns null
				return dparval.NewNullValue(), nil
			default:
				// any other error return to caller
				return nil, err
			}
		}

		if lenarg.Type() == dparval.NUMBER {
			lenval := lenarg.Value()
			switch lenval := lenval.(type) {
			case float64:
				maxLen = int(lenval)
				// FIXME add checks for negative values here?
			}
		} else {
			return dparval.NewNullValue(), nil
		}

	}

	// ensure that arg1 is a string
	if av.Type() == dparval.STRING {
		if position.Type() == dparval.NUMBER {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				// ensure that arg2 is a number
				posval := position.Value()
				switch posval := posval.(type) {
				case float64:
					pos := int(posval)

					//validate that pos is valid
					if pos < 0 || pos >= len(avalue) {
						// FIXME add warning for invalid pos?
						return nil, nil
					}

					if maxLen < 0 {
						// no end limit
						return dparval.NewStringValue(avalue[pos:]), nil
					} else {
						// validate that maxLen is valid
						endPos := pos + maxLen + 1
						if endPos < pos || endPos >= len(avalue) {
							// FIXME add warning for invalid max len?
							return nil, nil
						}
						return dparval.NewStringValue(avalue[pos:endPos]), nil
					}
				}
			}
		}
	}
	// FIXME warn if arguments were wrong type?
	return dparval.NewNullValue(), nil
}

func (this *FunctionSubStr) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 2 || len(arguments) > 4 {
		return fmt.Errorf("the SUBSTR() function expects two or three arguments")
	}

	return ValidateNoStars(this, arguments)
}
