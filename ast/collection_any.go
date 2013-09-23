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
	"github.com/couchbaselabs/dparval"
)

type CollectionAnyOperator struct {
	Type string `json:"type"`
	CollectionOperator
}

func NewCollectionAnyOperator(condition Expression, over Expression, as string) *CollectionAnyOperator {
	if as == "" {
		prop := expressionEndsInProperty(over)
		if prop != nil {
			candidateName, uniq := propertyName(prop, []string{})
			if uniq {
				as = candidateName
			}
		}
	}
	return &CollectionAnyOperator{
		"any",
		CollectionOperator{
			operator:  "ANY",
			Condition: condition,
			Over:      over,
			As:        as,
		},
	}
}

func (this *CollectionAnyOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the over
	ov, err := this.Over.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// spec says return false
			return dparval.NewValue(false), nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	// see if we're dealing with an array
	if ov.Type() == dparval.ARRAY {
		// by accessing the array contents this way
		// we avoid parsing it
		ok := true
		index := 0
		for ok {
			inner, err := ov.Index(index)
			index = index + 1
			if err != nil {
				switch err := err.(type) {
				case *dparval.Undefined:
					ok = false
				default:
					return nil, err
				}
			} else {

				// create a new context with this object named as the alias
				innerContext := dparval.NewValue(map[string]interface{}{this.As: inner})
				// now evaluate the condition in this new context
				innerResult, err := this.Condition.Evaluate(innerContext)
				if err != nil {
					switch err := err.(type) {
					case *dparval.Undefined:
						// this is not true, keep trying
						continue
					default:
						// any other error should be returned to caller
						return nil, err
					}
				}
				innerResultVal := innerResult.Value()
				// check to see if this value is true
				innerBoolResult := ValueInBooleanContext(innerResultVal)
				if innerBoolResult == true {
					return dparval.NewValue(true), nil
				}
			}
		}
		return dparval.NewValue(false), nil
	}
	return dparval.NewValue(false), nil
}

func (this *CollectionAnyOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
