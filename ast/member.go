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
	"encoding/json"
	"fmt"

	"github.com/couchbaselabs/dparval"
)

// ****************************************************************************
// DOT MEMBER
// ****************************************************************************

type DotMemberOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right *Property  `json:"right"`
}

func NewDotMemberOperator(left Expression, right *Property) *DotMemberOperator {
	return &DotMemberOperator{
		Type:  "dot_member",
		Left:  left,
		Right: right,
	}
}

func (this *DotMemberOperator) Copy() Expression {
	return &DotMemberOperator{
		Type:  "dot_member",
		Left:  this.Left.Copy(),
		Right: this.Right.Copy().(*Property),
	}
}

func (this *DotMemberOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == dparval.OBJECT {
		// now evaluate the property in this inner context
		return this.Right.Evaluate(lv)
	}

	return nil, &dparval.Undefined{this.Right.Path}
}

func (this *DotMemberOperator) String() string {
	return fmt.Sprintf("%v.%v", this.Left, this.Right)
}

func (this *DotMemberOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*DotMemberOperator)
	if !ok {
		return false
	}

	if !this.Left.EquivalentTo(that.Left) {
		return false
	}

	if !this.Right.EquivalentTo(that.Right) {
		return false
	}

	return true
}

// sliglty confusingly, a dot-member expression does not
// actually depend on its RHS
// a.b could be rewritten as a["b"] and since b is literal string
// it is not a dependency
// other way to look at it is that we're looking side what the LHS
// gave us, not the outer context, so we don't depend on the outside
// world at all
func (this *DotMemberOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left}
	return rv
}

func (this *DotMemberOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

func (this *DotMemberOperator) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string          `json:"type"`
		Left  json.RawMessage `json:"left"`
		Right json.RawMessage `json:"right"`
	}

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	if temp.Type != "dot_member" {
		return fmt.Errorf("Attempt to unmarshal type %v into dot-member", temp.Type)
	}

	this.Type = temp.Type
	this.Left, err = UnmarshalExpression(temp.Left)
	if err != nil {
		return err
	}
	rightExpr, err := UnmarshalExpression(temp.Right)
	if err != nil {
		return err
	}

	var ok bool
	this.Right, ok = rightExpr.(*Property)
	if !ok {
		return fmt.Errorf("Right-hand side of dot_member must be type Property")
	}

	return nil
}

// ****************************************************************************
// BRACKET MEMBER
// ****************************************************************************

type BracketMemberOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewBracketMemberOperator(left, right Expression) *BracketMemberOperator {
	return &BracketMemberOperator{
		Type:  "bracket_member",
		Left:  left,
		Right: right,
	}
}

func (this *BracketMemberOperator) Copy() Expression {
	return &BracketMemberOperator{
		Type:  "bracket_member",
		Left:  this.Left.Copy(),
		Right: this.Right.Copy(),
	}
}

func (this *BracketMemberOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// evaluting RHS first is more correct in case of side-effects
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == dparval.OBJECT {
		if rv.Type() == dparval.STRING {
			rval := rv.Value()
			switch rval := rval.(type) {
			case string:
				virtualProperty := NewProperty(rval)
				return virtualProperty.Evaluate(lv)
			}
		}
	} else if lv.Type() == dparval.ARRAY {
		if rv.Type() == dparval.NUMBER {
			rval := rv.Value()
			switch rval := rval.(type) {
			case float64:
				index := int(rval)
				return lv.Index(index)
			}

		}
	}

	return nil, &dparval.Undefined{}
}

func (this *BracketMemberOperator) String() string {
	return fmt.Sprintf("%v[%v]", this.Left, this.Right)
}

func (this *BracketMemberOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*BracketMemberOperator)
	if !ok {
		return false
	}

	if !this.Left.EquivalentTo(that.Left) {
		return false
	}

	if !this.Right.EquivalentTo(that.Right) {
		return false
	}

	return true
}

func (this *BracketMemberOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *BracketMemberOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

func (this *BracketMemberOperator) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type  string          `json:"type"`
		Left  json.RawMessage `json:"left"`
		Right json.RawMessage `json:"right"`
	}

	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	if temp.Type != "bracket_member" {
		return fmt.Errorf("Attempt to unmarshal type %v into bracket_member", temp.Type)
	}

	this.Type = temp.Type
	this.Left, err = UnmarshalExpression(temp.Left)
	if err != nil {
		return err
	}
	this.Right, err = UnmarshalExpression(temp.Right)
	if err != nil {
		return err
	}

	return nil
}
