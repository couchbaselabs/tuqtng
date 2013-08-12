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

type CommutativeBinaryOperator struct {
	BinaryOperator
}

func (this *CommutativeBinaryOperator) EquivalentTo(t Expression) bool {
	// another binary operator?
	that, ok := t.(BinaryOperatorExpression)
	if !ok {
		return false
	}

	// same type of operator?
	if this.operator != that.Operator() {
		return false
	}

	// same operands?
	if this.Left.EquivalentTo(that.GetLeft()) &&
		this.Right.EquivalentTo(that.GetRight()) {
		return true
	} else if this.Left.EquivalentTo(that.GetRight()) &&
		this.Right.EquivalentTo(that.GetLeft()) {
		return true
	}

	return false
}
