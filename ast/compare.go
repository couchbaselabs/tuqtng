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
	"log"
	"regexp"
	"strings"
)

type TypeMismatch struct {
	ltype int
	rtype int
}

func (this *TypeMismatch) Error() string {
	return fmt.Sprintf("Types do not match, %s %s", this.ltype, this.rtype)
}

type BinaryComparisonOperator struct {
	left  Expression
	right Expression
}

func (this *BinaryComparisonOperator) compare(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		// this could either be real error, or MISSING
		// if either side is MISSING, the result is MISSING
		return nil, err
	}
	// if either side is NULL, the result is NULL
	if lv == nil {
		return lv, nil
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		// this could either be real error, or MISSING
		// if either side is MISSING, the result is MISSING
		return nil, err
	}
	// if either side is NULL, the result is NULL
	if rv == nil {
		return rv, nil
	}

	// if we got this far, we evaluated both sides
	// there were no errors, and neither side was NULL or MISSING
	// now check types (types must be the same)
	ltype := collationType(lv)
	rtype := collationType(rv)
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

	return CollateJSON(lv, rv), nil
}

// ****************************************************************************
// Greater Than
// ****************************************************************************

type GreaterThanOperator struct {
	BinaryComparisonOperator
}

func NewGreaterThanOperator(left, right Expression) *GreaterThanOperator {
	return &GreaterThanOperator{
		BinaryComparisonOperator{
			left:  left,
			right: right,
		},
	}
}

func (this *GreaterThanOperator) Evaluate(item Item) (Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	switch compare := compare.(type) {
	case int:
		return compare > 0, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected result from comparison: %v", compare)
		return nil, nil
	}
}

// ****************************************************************************
// Greater Than Or Equal
// ****************************************************************************

type GreaterThanOrEqualOperator struct {
	BinaryComparisonOperator
}

func NewGreaterThanOrEqualOperator(left, right Expression) *GreaterThanOrEqualOperator {
	return &GreaterThanOrEqualOperator{
		BinaryComparisonOperator{
			left:  left,
			right: right,
		},
	}
}

func (this *GreaterThanOrEqualOperator) Evaluate(item Item) (Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	switch compare := compare.(type) {
	case int:
		return compare >= 0, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected result from comparison: %v", compare)
		return nil, nil
	}
}

// ****************************************************************************
// Less Than
// ****************************************************************************

type LessThanOperator struct {
	BinaryComparisonOperator
}

func NewLessThanOperator(left, right Expression) *LessThanOperator {
	return &LessThanOperator{
		BinaryComparisonOperator{
			left:  left,
			right: right,
		},
	}
}

func (this *LessThanOperator) Evaluate(item Item) (Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	switch compare := compare.(type) {
	case int:
		return compare < 0, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected result from comparison: %v", compare)
		return nil, nil
	}
}

// ****************************************************************************
// Less Than Or Equal
// ****************************************************************************

type LessThanOrEqualOperator struct {
	BinaryComparisonOperator
}

func NewLessThanOrEqualOperator(left, right Expression) *LessThanOrEqualOperator {
	return &LessThanOrEqualOperator{
		BinaryComparisonOperator{
			left:  left,
			right: right,
		},
	}
}

func (this *LessThanOrEqualOperator) Evaluate(item Item) (Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	switch compare := compare.(type) {
	case int:
		return compare <= 0, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected result from comparison: %v", compare)
		return nil, nil
	}
}

// ****************************************************************************
// Equal To
// ****************************************************************************

type EqualToOperator struct {
	BinaryComparisonOperator
}

func NewEqualToOperator(left, right Expression) *EqualToOperator {
	return &EqualToOperator{
		BinaryComparisonOperator{
			left:  left,
			right: right,
		},
	}
}

func (this *EqualToOperator) Evaluate(item Item) (Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	switch compare := compare.(type) {
	case int:
		return compare == 0, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected result from comparison: %v", compare)
		return nil, nil
	}
}

// ****************************************************************************
// Not Equal To
// ****************************************************************************

type NotEqualToOperator struct {
	BinaryComparisonOperator
}

func NewNotEqualToOperator(left, right Expression) *NotEqualToOperator {
	return &NotEqualToOperator{
		BinaryComparisonOperator{
			left:  left,
			right: right,
		},
	}
}

func (this *NotEqualToOperator) Evaluate(item Item) (Value, error) {
	compare, err := this.BinaryComparisonOperator.compare(item)
	if err != nil {
		switch err := err.(type) {
		case *TypeMismatch:
			// type mismatch is false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}
	switch compare := compare.(type) {
	case int:
		return compare != 0, nil
	case nil:
		return nil, nil
	default:
		log.Fatalf("Unexpected result from comparison: %v", compare)
		return nil, nil
	}
}

// ****************************************************************************
// Like
// ****************************************************************************
// FIXME - optimize case where RHS is string literal, only compile
//         the regular expression once
type LikeOperator struct {
	left  Expression
	right Expression
}

func NewLikeOperator(left, right Expression) *LikeOperator {
	return &LikeOperator{
		left:  left,
		right: right,
	}
}

func (this *LikeOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case string:
		switch rv := rv.(type) {
		case string:
			// if both values are string we can proceed
			pattern := strings.Replace(rv, "%", "(.*)", -1)
			pattern = strings.Replace(pattern, "_", "(.)", -1)
			re, err := regexp.Compile(pattern)
			if err != nil {
				return err, nil
			}
			return re.MatchString(lv), nil

		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// Not Like
// ****************************************************************************
// FIXME - consolidate common code with LIKE
type NotLikeOperator struct {
	left  Expression
	right Expression
}

func NewNotLikeOperator(left, right Expression) *NotLikeOperator {
	return &NotLikeOperator{
		left:  left,
		right: right,
	}
}

func (this *NotLikeOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case string:
		switch rv := rv.(type) {
		case string:
			// if both values are string we can proceed
			pattern := strings.Replace(rv, "%", "(.*)", -1)
			pattern = strings.Replace(pattern, "_", "(.)", -1)
			re, err := regexp.Compile(pattern)
			if err != nil {
				return err, nil
			}
			return !re.MatchString(lv), nil

		default:
			return nil, nil
		}
	default:
		return nil, nil
	}

	return nil, nil
}

// ****************************************************************************
// IS NULL
// ****************************************************************************

type IsNullOperator struct {
	operand Expression
}

func NewIsNullOperator(operand Expression) *IsNullOperator {
	return &IsNullOperator{
		operand: operand,
	}
}

func (this *IsNullOperator) Evaluate(item Item) (Value, error) {
	ov, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov == nil {
		return true, nil
	}

	return false, nil
}

// ****************************************************************************
// IS NOT NULL
// ****************************************************************************

type IsNotNullOperator struct {
	operand Expression
}

func NewIsNotNullOperator(operand Expression) *IsNotNullOperator {
	return &IsNotNullOperator{
		operand: operand,
	}
}

func (this *IsNotNullOperator) Evaluate(item Item) (Value, error) {
	ov, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov == nil {
		return false, nil
	}

	return true, nil
}

// ****************************************************************************
// IS MISSING
// ****************************************************************************

type IsMissingOperator struct {
	operand Expression
}

func NewIsMissingOperator(operand Expression) *IsMissingOperator {
	return &IsMissingOperator{
		operand: operand,
	}
}

func (this *IsMissingOperator) Evaluate(item Item) (Value, error) {
	_, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return true, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	return false, nil
}

// ****************************************************************************
// IS NOT MISSING
// ****************************************************************************

type IsNotMissingOperator struct {
	operand Expression
}

func NewIsNotMissingOperator(operand Expression) *IsNotMissingOperator {
	return &IsNotMissingOperator{
		operand: operand,
	}
}

func (this *IsNotMissingOperator) Evaluate(item Item) (Value, error) {
	_, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	return true, nil
}

// ****************************************************************************
// IS VALUED
// ****************************************************************************

type IsValuedOperator struct {
	operand Expression
}

func NewIsValuedOperator(operand Expression) *IsValuedOperator {
	return &IsValuedOperator{
		operand: operand,
	}
}

func (this *IsValuedOperator) Evaluate(item Item) (Value, error) {
	ov, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov == nil {
		return false, nil
	}

	return true, nil
}

// ****************************************************************************
// IS NOT VALUED
// ****************************************************************************

type IsNotValuedOperator struct {
	operand Expression
}

func NewIsNotValuedOperator(operand Expression) *IsNotValuedOperator {
	return &IsNotValuedOperator{
		operand: operand,
	}
}

func (this *IsNotValuedOperator) Evaluate(item Item) (Value, error) {
	ov, err := this.operand.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *Undefined:
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	if ov == nil {
		return true, nil
	}

	return false, nil
}
