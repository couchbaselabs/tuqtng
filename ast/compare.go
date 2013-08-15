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

// ****************************************************************************
// Greater Than
// ****************************************************************************

type GreaterThanOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewGreaterThanOperator(left, right Expression) *GreaterThanOperator {
	return &GreaterThanOperator{
		"greater_than",
		BinaryOperator{
			operator: ">",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *GreaterThanOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.compare(item)
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

func (this *GreaterThanOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Greater Than Or Equal
// ****************************************************************************

type GreaterThanOrEqualOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewGreaterThanOrEqualOperator(left, right Expression) *GreaterThanOrEqualOperator {
	return &GreaterThanOrEqualOperator{
		"greater_than_or_equal",
		BinaryOperator{
			operator: ">=",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *GreaterThanOrEqualOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.compare(item)
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

func (this *GreaterThanOrEqualOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Less Than
// ****************************************************************************

type LessThanOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewLessThanOperator(left, right Expression) *LessThanOperator {
	return &LessThanOperator{
		"less_than",
		BinaryOperator{
			operator: "<",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *LessThanOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.compare(item)
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

func (this *LessThanOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Less Than Or Equal
// ****************************************************************************

type LessThanOrEqualOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewLessThanOrEqualOperator(left, right Expression) *LessThanOrEqualOperator {
	return &LessThanOrEqualOperator{
		"less_than_or_equal",
		BinaryOperator{
			operator: "<=",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *LessThanOrEqualOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.compare(item)
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

func (this *LessThanOrEqualOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Equal To
// ****************************************************************************

type EqualToOperator struct {
	Type string `json:"type"`
	CommutativeBinaryOperator
}

func NewEqualToOperator(left, right Expression) *EqualToOperator {
	return &EqualToOperator{
		"equals",
		CommutativeBinaryOperator{
			BinaryOperator{
				operator: "=",
				Left:     left,
				Right:    right,
			},
		},
	}
}

func (this *EqualToOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.compare(item)
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

func (this *EqualToOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Not Equal To
// ****************************************************************************

type NotEqualToOperator struct {
	Type string `json:"type"`
	CommutativeBinaryOperator
}

func NewNotEqualToOperator(left, right Expression) *NotEqualToOperator {
	return &NotEqualToOperator{
		"not_equals",
		CommutativeBinaryOperator{
			BinaryOperator{
				operator: "!=",
				Left:     left,
				Right:    right,
			},
		},
	}
}

func (this *NotEqualToOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	compare, err := this.compare(item)
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

func (this *NotEqualToOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Like
// ****************************************************************************
// FIXME - optimize case where RHS is string literal, only compile
//         the regular expression once
type LikeOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewLikeOperator(left, right Expression) *LikeOperator {
	return &LikeOperator{
		"like",
		BinaryOperator{
			operator: "LIKE",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *LikeOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothString, err := this.EvaluateBothRequireString(context)
	if err != nil {
		return nil, err
	}
	if bothString {
		pattern := strings.Replace(rv, "%", "(.*)", -1)
		pattern = strings.Replace(pattern, "_", "(.)", -1)
		pattern = "^" + pattern + "$"
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		result = re.MatchString(lv)
	}
	return dparval.NewValue(result), nil
}

func (this *LikeOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// Not Like
// ****************************************************************************
// FIXME - consolidate common code with LIKE
type NotLikeOperator struct {
	Type string `json:"type"`
	BinaryOperator
}

func NewNotLikeOperator(left, right Expression) *NotLikeOperator {
	return &NotLikeOperator{
		"not_like",
		BinaryOperator{
			operator: "NOT LIKE",
			Left:     left,
			Right:    right,
		},
	}
}

func (this *NotLikeOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = nil
	lv, rv, bothString, err := this.EvaluateBothRequireString(context)
	if err != nil {
		return nil, err
	}
	if bothString {
		pattern := strings.Replace(rv, "%", "(.*)", -1)
		pattern = strings.Replace(pattern, "_", "(.)", -1)
		pattern = "^" + pattern + "$"
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		result = !re.MatchString(lv)
	}
	return dparval.NewValue(result), nil
}

func (this *NotLikeOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// IS NULL
// ****************************************************************************

type IsNullOperator struct {
	Type string `json:"type"`
	UnaryOperator
}

func NewIsNullOperator(operand Expression) *IsNullOperator {
	return &IsNullOperator{
		"is_null",
		UnaryOperator{
			operator: "IS NULL",
			Operand:  operand,
		},
	}
}

func (this *IsNullOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = false
	ov, isMissing, err := this.EvaluateFlagMissing(context)
	if err != nil && !isMissing {
		return nil, err
	}

	if ov != nil && ov.Type() == dparval.NULL {
		result = true
	}
	return dparval.NewValue(result), nil
}

func (this *IsNullOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// IS NOT NULL
// ****************************************************************************

type IsNotNullOperator struct {
	Type string `json:"type"`
	UnaryOperator
}

func NewIsNotNullOperator(operand Expression) *IsNotNullOperator {
	return &IsNotNullOperator{
		"is_not_null",
		UnaryOperator{
			operator: "IS NOT NULL",
			Operand:  operand,
		},
	}
}

func (this *IsNotNullOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = true
	ov, isMissing, err := this.EvaluateFlagMissing(context)
	if err != nil && !isMissing {
		return nil, err
	}

	if isMissing || ov.Type() == dparval.NULL {
		result = false
	}
	return dparval.NewValue(result), nil
}

func (this *IsNotNullOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// IS MISSING
// ****************************************************************************

type IsMissingOperator struct {
	Type string `json:"type"`
	UnaryOperator
}

func NewIsMissingOperator(operand Expression) *IsMissingOperator {
	return &IsMissingOperator{
		"is_missing",
		UnaryOperator{
			operator: "IS MISSING",
			Operand:  operand,
		},
	}
}

func (this *IsMissingOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = false
	_, isMissing, err := this.EvaluateFlagMissing(context)
	if err != nil && !isMissing {
		return nil, err
	}

	if isMissing {
		result = true
	}
	return dparval.NewValue(result), nil
}

func (this *IsMissingOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// IS NOT MISSING
// ****************************************************************************

type IsNotMissingOperator struct {
	Type string `json:"type"`
	UnaryOperator
}

func NewIsNotMissingOperator(operand Expression) *IsNotMissingOperator {
	return &IsNotMissingOperator{
		"is_not_missing",
		UnaryOperator{
			operator: "IS NOT MISSING",
			Operand:  operand,
		},
	}
}

func (this *IsNotMissingOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = true
	_, isMissing, err := this.EvaluateFlagMissing(context)
	if err != nil && !isMissing {
		return nil, err
	}

	if isMissing {
		result = false
	}
	return dparval.NewValue(result), nil
}

func (this *IsNotMissingOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// IS VALUED
// ****************************************************************************

type IsValuedOperator struct {
	Type string `json:"type"`
	UnaryOperator
}

func NewIsValuedOperator(operand Expression) *IsValuedOperator {
	return &IsValuedOperator{
		"is_valued",
		UnaryOperator{
			operator: "IS VALUED",
			Operand:  operand,
		},
	}
}

func (this *IsValuedOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = true
	ov, isMissing, err := this.EvaluateFlagMissing(context)
	if err != nil && !isMissing {
		return nil, err
	}

	if isMissing || ov.Type() == dparval.NULL {
		result = false
	}
	return dparval.NewValue(result), nil
}

func (this *IsValuedOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

// ****************************************************************************
// IS NOT VALUED
// ****************************************************************************

type IsNotValuedOperator struct {
	Type string `json:"type"`
	UnaryOperator
}

func NewIsNotValuedOperator(operand Expression) *IsNotValuedOperator {
	return &IsNotValuedOperator{
		"is_not_valued",
		UnaryOperator{
			operator: "IS NOT VALUED",
			Operand:  operand,
		},
	}
}

func (this *IsNotValuedOperator) Evaluate(context *dparval.Value) (*dparval.Value, error) {
	var result interface{} = false
	ov, isMissing, err := this.EvaluateFlagMissing(context)
	if err != nil && !isMissing {
		return nil, err
	}

	if ov != nil && ov.Type() == dparval.NULL {
		result = true
	}
	return dparval.NewValue(result), nil
}

func (this *IsNotValuedOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
