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

	"github.com/couchbaselabs/dparval"
)

type WhenThen struct {
	When Expression `json:"when"`
	Then Expression `json:"then"`
}

type CaseOperator struct {
	Type      string      `json:"type"`
	WhenThens []*WhenThen `json:"whenthens"`
	Else      Expression  `json:"else"`
}

func NewCaseOperator() *CaseOperator {
	return &CaseOperator{
		Type: "case",
	}
}

func (this *CaseOperator) Copy() Expression {
	rv := CaseOperator{
		Type:      "case",
		WhenThens: make([]*WhenThen, len(this.WhenThens)),
	}

	for i, wt := range this.WhenThens {
		rv.WhenThens[i] = &WhenThen{When: wt.When.Copy(), Then: wt.Then.Copy()}
	}

	if this.Else != nil {
		rv.Else = this.Else.Copy()
	}

	return &rv
}

func (this *CaseOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// walk through the WhenThens in order
	for _, WhenThen := range this.WhenThens {
		// evalute the when
		whenVal, err := WhenThen.When.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// MISSING is not true, keep looking
				continue
			default:
				// any other error should be returned to caller
				return nil, err
			}
		}

		whenValVal := whenVal.Value()

		whenBoolVal := ValueInBooleanContext(whenValVal)
		if whenBoolVal == true {
			// evalauate the then
			thenVal, err := WhenThen.Then.Evaluate(item)
			if err != nil {
				return nil, err
			}
			return thenVal, nil
		}
	}

	// if we got this point, none of the WhenThen pairs were satisfied
	// if there was an else, evaluate it and return
	if this.Else != nil {
		elseVal, err := this.Else.Evaluate(item)
		if err != nil {
			return nil, err
		}
		return elseVal, nil
	}
	// if there was no else, return null
	return dparval.NewValue(nil), nil
}

func (this *CaseOperator) String() string {
	inside := ""
	for i, wt := range this.WhenThens {
		if i != 0 {
			inside = inside + " "
		}
		inside = inside + fmt.Sprintf("WHEN %v THEN %v", wt.When, wt.Then)
	}
	if this.Else != nil {
		inside = inside + fmt.Sprintf(" ELSE %v", this.Else)
	}
	return fmt.Sprintf("CASE %v END", inside)
}

func (this *CaseOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*CaseOperator)
	if !ok {
		return false
	}

	// order of the when/the conditions does matter
	// for now, insist on the exact same order
	if len(this.WhenThens) != len(that.WhenThens) {
		return false
	}
	for i, thisWT := range this.WhenThens {
		thatWT := that.WhenThens[i]
		if !thisWT.When.EquivalentTo(thatWT.When) {
			return false
		}
		if !thatWT.Then.EquivalentTo(thatWT.Then) {
			return false
		}
	}
	if this.Else != nil && that.Else != nil {
		if !this.Else.EquivalentTo(that.Else) {
			return false
		}
	}

	return true
}

func (this *CaseOperator) Dependencies() ExpressionList {
	rv := ExpressionList{}

	for _, WhenThen := range this.WhenThens {
		rv = append(rv, WhenThen.When)
		rv = append(rv, WhenThen.Then)
	}
	if this.Else != nil {
		rv = append(rv, this.Else)
	}

	return rv
}

func (this *CaseOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
