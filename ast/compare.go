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

func (this *BinaryComparisonOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
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

func (this *GreaterThanOperator) String() string {
	return fmt.Sprintf("%v > %v", this.Left, this.Right)
}

func (this *GreaterThanOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*GreaterThanOperator)
	if !ok {
		return false
	}

	if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Left) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Right) {
		return true
	}

	return false
}

func (this *GreaterThanOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *GreaterThanOrEqualOperator) String() string {
	return fmt.Sprintf("%v >= %v", this.Left, this.Right)
}

func (this *GreaterThanOrEqualOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*GreaterThanOrEqualOperator)
	if !ok {
		return false
	}

	if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Left) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Right) {
		return true
	}

	return false
}

func (this *GreaterThanOrEqualOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LessThanOperator) String() string {
	return fmt.Sprintf("%v < %v", this.Left, this.Right)
}

func (this *LessThanOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*LessThanOperator)
	if !ok {
		return false
	}

	if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Left) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Right) {
		return true
	}

	return false
}

func (this *LessThanOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LessThanOrEqualOperator) String() string {
	return fmt.Sprintf("%v <= %v", this.Left, this.Right)
}

func (this *LessThanOrEqualOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*LessThanOrEqualOperator)
	if !ok {
		return false
	}

	if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Left) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Right) {
		return true
	}

	return false
}

func (this *LessThanOrEqualOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *EqualToOperator) String() string {
	return fmt.Sprintf("%v = %v", this.Left, this.Right)
}

func (this *EqualToOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*EqualToOperator)
	if !ok {
		return false
	}

	// equals order doesnt matter
	if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Left) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Right) {
		return true
	} else if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Right) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Left) {
		return true
	}

	return false
}

func (this *EqualToOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *NotEqualToOperator) String() string {
	return fmt.Sprintf("%v != %v", this.Left, this.Right)
}

func (this *NotEqualToOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*NotEqualToOperator)
	if !ok {
		return false
	}

	// not equals order doesnt matter
	if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Left) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Right) {
		return true
	} else if this.BinaryComparisonOperator.Left.EquivalentTo(that.BinaryComparisonOperator.Right) &&
		this.BinaryComparisonOperator.Right.EquivalentTo(that.BinaryComparisonOperator.Left) {
		return true
	}

	return false
}

func (this *NotEqualToOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *LikeOperator) String() string {
	return fmt.Sprintf("%v LIKE %v", this.Left, this.Right)
}

func (this *LikeOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*LikeOperator)
	if !ok {
		return false
	}

	// not equals order doesnt matter
	if this.Left.EquivalentTo(that.Left) &&
		this.Right.EquivalentTo(that.Right) {
		return true
	}

	return false
}

func (this *LikeOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *LikeOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *NotLikeOperator) String() string {
	return fmt.Sprintf("%v NOT LIKE %v", this.Left, this.Right)
}

func (this *NotLikeOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*NotLikeOperator)
	if !ok {
		return false
	}

	// not equals order doesnt matter
	if this.Left.EquivalentTo(that.Left) &&
		this.Right.EquivalentTo(that.Right) {
		return true
	}

	return false
}

func (this *NotLikeOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *NotLikeOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *IsNullOperator) String() string {
	return fmt.Sprintf("%v IS NULL", this.Operand)
}

func (this *IsNullOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*IsNullOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}

	return false
}

func (this *IsNullOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *IsNullOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *IsNotNullOperator) String() string {
	return fmt.Sprintf("%v IS NOT NULL", this.Operand)
}

func (this *IsNotNullOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*IsNotNullOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}

	return false
}

func (this *IsNotNullOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *IsNotNullOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *IsMissingOperator) String() string {
	return fmt.Sprintf("%v IS MISSING", this.Operand)
}

func (this *IsMissingOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*IsMissingOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}

	return false
}

func (this *IsMissingOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *IsMissingOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *IsNotMissingOperator) String() string {
	return fmt.Sprintf("%v IS NOT MISSING", this.Operand)
}

func (this *IsNotMissingOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*IsNotMissingOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}

	return false
}

func (this *IsNotMissingOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *IsNotMissingOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *IsValuedOperator) String() string {
	return fmt.Sprintf("%v IS VALUED", this.Operand)
}

func (this *IsValuedOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*IsValuedOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}

	return false
}

func (this *IsValuedOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *IsValuedOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
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

func (this *IsNotValuedOperator) String() string {
	return fmt.Sprintf("%v IS NOT VALUED", this.Operand)
}

func (this *IsNotValuedOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*IsNotValuedOperator)
	if !ok {
		return false
	}

	if this.Operand.EquivalentTo(that.Operand) {
		return true
	}

	return false
}

func (this *IsNotValuedOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Operand}
	return rv
}

func (this *IsNotValuedOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
