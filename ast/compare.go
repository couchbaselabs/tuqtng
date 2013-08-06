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
	"regexp"
	"strings"

	"github.com/couchbaselabs/dparval"
)

type TypeMismatch struct {
	ltype int
	rtype int
}

func (this *TypeMismatch) Error() string {
	return fmt.Sprintf("Types do not match, %d %d", this.ltype, this.rtype)
}

type BinaryComparisonOperator struct {
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func (this *BinaryComparisonOperator) compare(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		// this could either be real error, or MISSING
		// if either side is MISSING, the result is MISSING
		return nil, err
	}
	// if either side is NULL, the result is NULL
	if lv.Type() == dparval.NULL {
		return nil, nil
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		// this could either be real error, or MISSING
		// if either side is MISSING, the result is MISSING
		return nil, err
	}
	// if either side is NULL, the result is NULL
	if rv.Type() == dparval.NULL {
		return nil, nil
	}

	lvalue := lv.Value()
	rvalue := rv.Value()

	// if we got this far, we evaluated both sides
	// there were no errors, and neither side was NULL or MISSING
	// now check types (types must be the same)
	ltype := collationType(lvalue)
	rtype := collationType(rvalue)
	// ugly fixups for boolean (returns different values for true/false)
	if ltype == 2 {
		// fixup for boolean type
		ltype = 1
	}
	if rtype == 2 {
		rtype = 1
	}

	if ltype != rtype {
		return nil, &TypeMismatch{ltype, rtype}
	}

	return dparval.NewValue(float64(CollateJSON(lvalue, rvalue))), nil
}

func (this *BinaryComparisonOperator) validate() error {
	err := this.Left.Validate()
	if err != nil {
		return err
	}
	err = this.Right.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *BinaryComparisonOperator) verifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

// ****************************************************************************
// Greater Than
// ****************************************************************************

type GreaterThanOperator struct {
	Type string `json:"type"`
	BinaryComparisonOperator
}

func NewGreaterThanOperator(left, right Expression) *GreaterThanOperator {
	return &GreaterThanOperator{
		"greater_than",
		BinaryComparisonOperator{
			Left:  left,
			Right: right,
		},
	}
}

func (this *GreaterThanOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	if compare == nil {
		return dparval.NewValue(nil), nil
	}
	compareValue := compare.Value()
	switch compareValue := compareValue.(type) {
	case float64:
		return dparval.NewValue(compareValue > 0), nil
	default:
		panic(fmt.Sprintf("Unexpected result from comparison: %v for %v, %v", compareValue, this.Left, this.Right))
	}
}

func (this *GreaterThanOperator) Validate() error {
	return this.BinaryComparisonOperator.validate()
}

func (this *GreaterThanOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return this.BinaryComparisonOperator.verifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
}

func (this *GreaterThanOperator) String() string {
	return fmt.Sprintf("%v > %v", this.Left, this.Right)
}

// ****************************************************************************
// Greater Than Or Equal
// ****************************************************************************

type GreaterThanOrEqualOperator struct {
	Type string `json:"type"`
	BinaryComparisonOperator
}

func NewGreaterThanOrEqualOperator(left, right Expression) *GreaterThanOrEqualOperator {
	return &GreaterThanOrEqualOperator{
		"greater_than_or_equal",
		BinaryComparisonOperator{
			Left:  left,
			Right: right,
		},
	}
}

func (this *GreaterThanOrEqualOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	if compare == nil {
		return dparval.NewValue(nil), nil
	}
	compareValue := compare.Value()
	switch compareValue := compareValue.(type) {
	case float64:
		return dparval.NewValue(compareValue >= 0), nil
	default:
		panic(fmt.Sprintf("Unexpected result from comparison: %v for %v, %v", compareValue, this.Left, this.Right))
	}
}

func (this *GreaterThanOrEqualOperator) Validate() error {
	return this.BinaryComparisonOperator.validate()
}

func (this *GreaterThanOrEqualOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return this.BinaryComparisonOperator.verifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
}

func (this *GreaterThanOrEqualOperator) String() string {
	return fmt.Sprintf("%v >= %v", this.Left, this.Right)
}

// ****************************************************************************
// Less Than
// ****************************************************************************

type LessThanOperator struct {
	Type string `json:"type"`
	BinaryComparisonOperator
}

func NewLessThanOperator(left, right Expression) *LessThanOperator {
	return &LessThanOperator{
		"less_than",
		BinaryComparisonOperator{
			Left:  left,
			Right: right,
		},
	}
}

func (this *LessThanOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	if compare == nil {
		return dparval.NewValue(nil), nil
	}
	compareValue := compare.Value()
	switch compareValue := compareValue.(type) {
	case float64:
		return dparval.NewValue(compareValue < 0), nil
	default:
		panic(fmt.Sprintf("Unexpected result from comparison: %v for %v, %v", compareValue, this.Left, this.Right))
	}
}

func (this *LessThanOperator) Validate() error {
	return this.BinaryComparisonOperator.validate()
}

func (this *LessThanOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return this.BinaryComparisonOperator.verifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
}

func (this *LessThanOperator) String() string {
	return fmt.Sprintf("%v < %v", this.Left, this.Right)
}

// ****************************************************************************
// Less Than Or Equal
// ****************************************************************************

type LessThanOrEqualOperator struct {
	Type string `json:"type"`
	BinaryComparisonOperator
}

func NewLessThanOrEqualOperator(left, right Expression) *LessThanOrEqualOperator {
	return &LessThanOrEqualOperator{
		"less_than_or_equal",
		BinaryComparisonOperator{
			Left:  left,
			Right: right,
		},
	}
}

func (this *LessThanOrEqualOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	if compare == nil {
		return dparval.NewValue(nil), nil
	}
	compareValue := compare.Value()
	switch compareValue := compareValue.(type) {
	case float64:
		return dparval.NewValue(compareValue <= 0), nil
	default:
		panic(fmt.Sprintf("Unexpected result from comparison: %v for %v, %v", compareValue, this.Left, this.Right))
	}
}

func (this *LessThanOrEqualOperator) Validate() error {
	return this.BinaryComparisonOperator.validate()
}

func (this *LessThanOrEqualOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return this.BinaryComparisonOperator.verifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
}

func (this *LessThanOrEqualOperator) String() string {
	return fmt.Sprintf("%v <= %v", this.Left, this.Right)
}

// ****************************************************************************
// Equal To
// ****************************************************************************

type EqualToOperator struct {
	Type string `json:"type"`
	BinaryComparisonOperator
}

func NewEqualToOperator(left, right Expression) *EqualToOperator {
	return &EqualToOperator{
		"equals",
		BinaryComparisonOperator{
			Left:  left,
			Right: right,
		},
	}
}

func (this *EqualToOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	if compare == nil {
		return dparval.NewValue(nil), nil
	}
	compareValue := compare.Value()
	switch compareValue := compareValue.(type) {
	case float64:
		return dparval.NewValue(compareValue == 0), nil
	default:
		panic(fmt.Sprintf("Unexpected result from comparison: %v for %v, %v", compareValue, this.Left, this.Right))
	}
}

func (this *EqualToOperator) Validate() error {
	return this.BinaryComparisonOperator.validate()
}

func (this *EqualToOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return this.BinaryComparisonOperator.verifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
}

func (this *EqualToOperator) String() string {
	return fmt.Sprintf("%v = %v", this.Left, this.Right)
}

// ****************************************************************************
// Not Equal To
// ****************************************************************************

type NotEqualToOperator struct {
	Type string `json:"type"`
	BinaryComparisonOperator
}

func NewNotEqualToOperator(left, right Expression) *NotEqualToOperator {
	return &NotEqualToOperator{
		"not_equals",
		BinaryComparisonOperator{
			Left:  left,
			Right: right,
		},
	}
}

func (this *NotEqualToOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	if compare == nil {
		return dparval.NewValue(nil), nil
	}
	compareValue := compare.Value()
	switch compareValue := compareValue.(type) {
	case float64:
		return dparval.NewValue(compareValue != 0), nil
	default:
		panic(fmt.Sprintf("Unexpected result from comparison: %v for %v, %v", compareValue, this.Left, this.Right))
	}
}

func (this *NotEqualToOperator) Validate() error {
	return this.BinaryComparisonOperator.validate()
}

func (this *NotEqualToOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	return this.BinaryComparisonOperator.verifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
}

func (this *NotEqualToOperator) String() string {
	return fmt.Sprintf("%v != %v", this.Left, this.Right)
}

// ****************************************************************************
// Like
// ****************************************************************************
// FIXME - optimize case where RHS is string literal, only compile
//         the regular expression once
type LikeOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewLikeOperator(left, right Expression) *LikeOperator {
	return &LikeOperator{
		Type:  "like",
		Left:  left,
		Right: right,
	}
}

func (this *LikeOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.STRING {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case string:
			switch rvalue := rvalue.(type) {
			case string:
				// if both values are string we can proceed
				pattern := strings.Replace(rvalue, "%", "(.*)", -1)
				pattern = strings.Replace(pattern, "_", "(.)", -1)
				pattern = "^" + pattern + "$"
				re, err := regexp.Compile(pattern)
				if err != nil {
					return nil, err
				}
				return dparval.NewValue(re.MatchString(lvalue)), nil
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *LikeOperator) Validate() error {
	err := this.Left.Validate()
	if err != nil {
		return err
	}
	err = this.Right.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *LikeOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *LikeOperator) String() string {
	return fmt.Sprintf("%v LIKE %v", this.Left, this.Right)
}

// ****************************************************************************
// Not Like
// ****************************************************************************
// FIXME - consolidate common code with LIKE
type NotLikeOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewNotLikeOperator(left, right Expression) *NotLikeOperator {
	return &NotLikeOperator{
		Type:  "not_like",
		Left:  left,
		Right: right,
	}
}

func (this *NotLikeOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == rv.Type() && rv.Type() == dparval.STRING {
		lvalue := lv.Value()
		rvalue := rv.Value()
		switch lvalue := lvalue.(type) {
		case string:
			switch rvalue := rvalue.(type) {
			case string:
				// if both values are string we can proceed
				pattern := strings.Replace(rvalue, "%", "(.*)", -1)
				pattern = strings.Replace(pattern, "_", "(.)", -1)
				pattern = "^" + pattern + "$"
				re, err := regexp.Compile(pattern)
				if err != nil {
					return nil, err
				}
				return dparval.NewValue(!re.MatchString(lvalue)), nil
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *NotLikeOperator) Validate() error {
	err := this.Left.Validate()
	if err != nil {
		return err
	}
	err = this.Right.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *NotLikeOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newleft, err := this.Left.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newleft != nil {
		this.Left = newleft
	}
	newright, err := this.Right.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newright != nil {
		this.Right = newright
	}
	return nil, nil
}

func (this *NotLikeOperator) String() string {
	return fmt.Sprintf("%v NOT LIKE %v", this.Left, this.Right)
}

// ****************************************************************************
// IS NULL
// ****************************************************************************

type IsNullOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewIsNullOperator(operand Expression) *IsNullOperator {
	return &IsNullOperator{
		Type:    "is_null",
		Operand: operand,
	}
}

func (this *IsNullOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov.Type() == dparval.NULL {
		return dparval.NewValue(true), nil
	}

	return dparval.NewValue(false), nil
}

func (this *IsNullOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *IsNullOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *IsNullOperator) String() string {
	return fmt.Sprintf("%v IS NULL", this.Operand)
}

// ****************************************************************************
// IS NOT NULL
// ****************************************************************************

type IsNotNullOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewIsNotNullOperator(operand Expression) *IsNotNullOperator {
	return &IsNotNullOperator{
		Type:    "is_not_null",
		Operand: operand,
	}
}

func (this *IsNotNullOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov.Type() == dparval.NULL {
		return dparval.NewValue(false), nil
	}

	return dparval.NewValue(true), nil
}

func (this *IsNotNullOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *IsNotNullOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *IsNotNullOperator) String() string {
	return fmt.Sprintf("%v IS NOT NULL", this.Operand)
}

// ****************************************************************************
// IS MISSING
// ****************************************************************************

type IsMissingOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewIsMissingOperator(operand Expression) *IsMissingOperator {
	return &IsMissingOperator{
		Type:    "is_missing",
		Operand: operand,
	}
}

func (this *IsMissingOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	_, err := this.Operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return dparval.NewValue(true), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	return dparval.NewValue(false), nil
}

func (this *IsMissingOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *IsMissingOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *IsMissingOperator) String() string {
	return fmt.Sprintf("%v IS MISSING", this.Operand)
}

// ****************************************************************************
// IS NOT MISSING
// ****************************************************************************

type IsNotMissingOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewIsNotMissingOperator(operand Expression) *IsNotMissingOperator {
	return &IsNotMissingOperator{
		Type:    "is_not_missing",
		Operand: operand,
	}
}

func (this *IsNotMissingOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	_, err := this.Operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	return dparval.NewValue(true), nil
}

func (this *IsNotMissingOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *IsNotMissingOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *IsNotMissingOperator) String() string {
	return fmt.Sprintf("%v IS NOT MISSING", this.Operand)
}

// ****************************************************************************
// IS VALUED
// ****************************************************************************

type IsValuedOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewIsValuedOperator(operand Expression) *IsValuedOperator {
	return &IsValuedOperator{
		Type:    "is_valued",
		Operand: operand,
	}
}

func (this *IsValuedOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov.Type() == dparval.NULL {
		return dparval.NewValue(false), nil
	}

	return dparval.NewValue(true), nil
}

func (this *IsValuedOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *IsValuedOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *IsValuedOperator) String() string {
	return fmt.Sprintf("%v IS VALUED", this.Operand)
}

// ****************************************************************************
// IS NOT VALUED
// ****************************************************************************

type IsNotValuedOperator struct {
	Type    string     `json:"type"`
	Operand Expression `json:"operand"`
}

func NewIsNotValuedOperator(operand Expression) *IsNotValuedOperator {
	return &IsNotValuedOperator{
		Type:    "is_not_valued",
		Operand: operand,
	}
}

func (this *IsNotValuedOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	ov, err := this.Operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov.Type() == dparval.NULL {
		return dparval.NewValue(true), nil
	}

	return dparval.NewValue(false), nil
}

func (this *IsNotValuedOperator) Validate() error {
	err := this.Operand.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *IsNotValuedOperator) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	newoper, err := this.Operand.VerifyFormalNotation(forbiddenAliases, aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newoper != nil {
		this.Operand = newoper
	}
	return nil, nil
}

func (this *IsNotValuedOperator) String() string {
	return fmt.Sprintf("%v IS NOT VALUED", this.Operand)
}
