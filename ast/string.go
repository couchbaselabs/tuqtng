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

// ****************************************************************************
// String Concatenation
// ****************************************************************************

type StringConcatenateOperator struct {
	Type  string     `json:"type"`
	Left  Expression `json:"left"`
	Right Expression `json:"right"`
}

func NewStringConcatenateOperator(left, right Expression) *StringConcatenateOperator {
	return &StringConcatenateOperator{
		Type:  "string_concat",
		Left:  left,
		Right: right,
	}
}

func (this *StringConcatenateOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	lv, err := this.Left.Evaluate(item)
	if err != nil {
		return nil, err
	}
	rv, err := this.Right.Evaluate(item)
	if err != nil {
		return nil, err
	}

	if lv.Type() == dparval.STRING {
		if rv.Type() == dparval.STRING {
			lval := lv.Value()
			switch lval := lval.(type) {
			case string:
				rval := rv.Value()
				switch rval := rval.(type) {
				case string:
					// if both values are string concatenate them
					return dparval.NewValue(lval + rval), nil
				}
			}
		}
	}
	return dparval.NewValue(nil), nil
}

func (this *StringConcatenateOperator) String() string {
	return fmt.Sprintf("%v || %v", this.Left, this.Right)
}

func (this *StringConcatenateOperator) EquivalentTo(t Expression) bool {
	that, ok := t.(*StringConcatenateOperator)
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

func (this *StringConcatenateOperator) Dependencies() ExpressionList {
	rv := ExpressionList{this.Left, this.Right}
	return rv
}

func (this *StringConcatenateOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
