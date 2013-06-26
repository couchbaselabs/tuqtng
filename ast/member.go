//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import ()

// ****************************************************************************
// DOT MEMBER
// ****************************************************************************

type DotMemberOperator struct {
	left  Expression
	right *Property
}

func NewDotMemberOperator(left Expression, right *Property) *DotMemberOperator {
	return &DotMemberOperator{
		left:  left,
		right: right,
	}
}

func (this *DotMemberOperator) Evaluate(item Item) (Value, error) {
	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case map[string]Value:
		innerContext := NewMapItem(lv, nil)
		// now evaluate the property in this inner context
		return this.right.Evaluate(innerContext)
	}

	return nil, &Undefined{this.right.Path}
}

// ****************************************************************************
// BRACKET MEMBER
// ****************************************************************************

type BracketMemberOperator struct {
	left  Expression
	right Expression
}

func NewBracketMemberOperator(left, right Expression) *BracketMemberOperator {
	return &BracketMemberOperator{
		left:  left,
		right: right,
	}
}

func (this *BracketMemberOperator) Evaluate(item Item) (Value, error) {
	// evaluting RHS first is more correct in case of side-effects
	rv, err := this.right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	lv, err := this.left.Evaluate(item)
	if err != nil {
		return nil, err
	}

	switch lv := lv.(type) {
	case map[string]Value:
		switch rv := rv.(type) {
		case string:
			innerContext := NewMapItem(lv, nil)
			virtualProperty := NewProperty(rv)
			return virtualProperty.Evaluate(innerContext)
		}
	case []Value:
		switch rv := rv.(type) {
		case float64:
			index := int(rv)
			if index >= 0 && index < len(lv) {
				return lv[index], nil
			}
		}
	}

	return nil, &Undefined{}
}
