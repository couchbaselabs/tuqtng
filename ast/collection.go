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
)

type CollectionOperator struct {
	operator  string
	Condition Expression `json:"condition"`
	Over      Expression `json:"over"`
	As        string     `json:"as"`
}

func (this *CollectionOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Over, this.Condition}
	return rv
}

func (this *CollectionOperator) EquivalentTo(t Expression) bool {
	// another collection operator?
	that, ok := t.(CollectionOperatorExpression)
	if !ok {
		return false
	}

	// same type of operator?
	if this.operator != that.Operator() {
		return false
	}

	if this.As != that.GetAs() {
		return false
	}
	if !this.Over.EquivalentTo(that.GetOver()) {
		return false
	}
	if !this.Condition.EquivalentTo(that.GetCondition()) {
		return false
	}

	return true
}

func (this *CollectionOperator) String() string {
	return fmt.Sprintf("%s %v OVER %v AS %s", this.operator, this.Condition, this.Over, this.As)
}

func (this *CollectionOperator) Operator() string {
	return this.operator
}

func (this *CollectionOperator) GetOver() Expression {
	return this.Over
}

func (this *CollectionOperator) GetCondition() Expression {
	return this.Condition
}

func (this *CollectionOperator) GetAs() string {
	return this.As
}

func (this *CollectionOperator) SetOver(over Expression) {
	this.Over = over
}

func (this *CollectionOperator) SetCondition(condition Expression) {
	this.Condition = condition
}

func (this *CollectionOperator) SetAs(as string) {
	this.As = as
}
