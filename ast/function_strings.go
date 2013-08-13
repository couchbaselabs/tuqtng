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
	"strings"

	"github.com/couchbaselabs/dparval"
)

type FunctionCallLower struct {
	FunctionCall
}

func NewFunctionCallLower(operands FunctionArgExpressionList) Expression {
	return &FunctionCallLower{
		FunctionCall{
			Type:     "function",
			Name:     "LOWER",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallLower) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			return dparval.NewValue(strings.ToLower(avalue)), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallLower) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallUpper struct {
	FunctionCall
}

func NewFunctionCallUpper(operands FunctionArgExpressionList) Expression {
	return &FunctionCallUpper{
		FunctionCall{
			Type:     "function",
			Name:     "UPPER",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallUpper) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	if av.Type() == dparval.STRING {
		avalue := av.Value()
		switch avalue := avalue.(type) {
		case string:
			return dparval.NewValue(strings.ToUpper(avalue)), nil
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallUpper) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallLTrim struct {
	FunctionCall
}

func NewFunctionCallLTrim(operands FunctionArgExpressionList) Expression {
	return &FunctionCallLTrim{
		FunctionCall{
			Type:     "function",
			Name:     "LTRIM",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallLTrim) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	// evaluate the second argument
	cutlist, err := this.Operands[1].Expr.Evaluate(item)
	// the cut list MUST be a string, otherwise return null
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

	if av.Type() == dparval.STRING {
		if cutlist.Type() == dparval.STRING {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				cutvlal := cutlist.Value()
				switch cutvlal := cutvlal.(type) {
				case string:
					return dparval.NewValue(strings.TrimLeft(avalue, cutvlal)), nil
				}
			}
		}

	}
	// FIXME warn if cutlist wasnt string?
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallLTrim) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallRTrim struct {
	FunctionCall
}

func NewFunctionCallRTrim(operands FunctionArgExpressionList) Expression {
	return &FunctionCallRTrim{
		FunctionCall{
			Type:     "function",
			Name:     "RTRIM",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallRTrim) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	// evaluate the second argument
	cutlist, err := this.Operands[1].Expr.Evaluate(item)
	// the cut list MUST be a string, otherwise return null
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

	if av.Type() == dparval.STRING {
		if cutlist.Type() == dparval.STRING {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				cutval := cutlist.Value()
				switch cutval := cutval.(type) {
				case string:
					return dparval.NewValue(strings.TrimRight(avalue, cutval)), nil
				}
			}
		}
	}
	// FIXME warn if cutlist wasnt string?
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallRTrim) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallTrim struct {
	FunctionCall
}

func NewFunctionCallTrim(operands FunctionArgExpressionList) Expression {
	return &FunctionCallTrim{
		FunctionCall{
			Type:     "function",
			Name:     "TRIM",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallTrim) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument
	av, err := this.Operands[0].Expr.Evaluate(item)

	// the spec defines this functin to ONLY operate on strings
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

	// evaluate the second argument
	cutlist, err := this.Operands[1].Expr.Evaluate(item)
	// the cut list MUST be a string, otherwise return null
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

	if av.Type() == dparval.STRING {
		if cutlist.Type() == dparval.STRING {
			avalue := av.Value()
			switch avalue := avalue.(type) {
			case string:
				cutval := cutlist.Value()
				switch cutval := cutval.(type) {
				case string:
					return dparval.NewValue(strings.Trim(avalue, cutval)), nil
				}
			}
		}
	}
	// FIXME warn if that cutlist wasnt string?
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallTrim) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallSubStr struct {
	FunctionCall
}

func NewFunctionCallSubStr(operands FunctionArgExpressionList) Expression {
	return &FunctionCallSubStr{
		FunctionCall{
			Type:     "function",
			Name:     "SUBSTR",
			Operands: operands,
			minArgs:  2,
			maxArgs:  3,
		},
	}
}

func (this *FunctionCallSubStr) Evaluate(item *dparval.Value) (*dparval.Value, error) {
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

	position, err := this.Operands[1].Expr.Evaluate(item)
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

	var maxLen int = -1

	if len(this.Operands) == 3 {
		lenarg, err := this.Operands[2].Expr.Evaluate(item)
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

		if lenarg.Type() == dparval.NUMBER {
			lenval := lenarg.Value()
			switch lenval := lenval.(type) {
			case float64:
				maxLen = int(lenval)
				// FIXME add checks for negative values here?
			}
		} else {
			return dparval.NewValue(nil), nil
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
						return dparval.NewValue(nil), nil
					}

					if maxLen < 0 {
						// no end limit
						return dparval.NewValue(avalue[pos:]), nil
					} else {
						// validate that maxLen is valid
						endPos := pos + maxLen
						if endPos < pos || endPos >= len(avalue) {
							// FIXME add warning for invalid max len?
							return dparval.NewValue(nil), nil
						}
						return dparval.NewValue(avalue[pos:endPos]), nil
					}
				}
			}
		}
	}
	// FIXME warn if arguments were wrong type?
	return dparval.NewValue(nil), nil
}

func (this *FunctionCallSubStr) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
