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

type CollectionArrayOperator struct {
	Type string `json:"type"`
	CollectionOperator
}

func NewCollectionArrayOperator(condition Expression, over Expression, as string, output Expression) *CollectionArrayOperator {
	if as == "" {
		prop := expressionEndsInProperty(over)
		if prop != nil {
			candidateName, uniq := propertyName(prop, []string{})
			if uniq {
				as = candidateName
			}
		}
	}
	return &CollectionArrayOperator{
		"array",
		CollectionOperator{
			operator:  "ARRAY",
			Condition: condition,
			Over:      over,
			As:        as,
			Output:    output,
		},
	}
}

func (this *CollectionArrayOperator) Copy() Expression {
	rv := CollectionArrayOperator{
		"array",
		CollectionOperator{
			operator: "ARRAY",
			Over:     this.Over.Copy(),
			As:       this.As,
			Output:   this.Output.Copy(),
		},
	}

	if this.Condition != nil {
		rv.Condition = this.Condition.Copy()
	}

	return &rv
}

func (this *CollectionArrayOperator) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the over
	ov, err := this.Over.Evaluate(item)
	if err != nil {
		return nil, err
	}

	// see if we're dealing with an array
	if ov.Type() == dparval.ARRAY {
		// by accessing the array contents this way
		// we avoid parsing it
		ok := true
		index := 0
		rv := make([]interface{}, 0)
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
				// duplicate the existing context
				innerContext := item.Duplicate()
				// add this object named as the alias
				innerContext.SetPath(this.As, inner)
				if this.Condition != nil {
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
						// now we have to evaluate the output expression
						outputResult, err := this.Output.Evaluate(innerContext)
						if err != nil {
							switch err := err.(type) {
							case *dparval.Undefined:
								// contribute nothing to the result array
							default:
								// any other error should be returned to caller
								return nil, err
							}
						} else {
							rv = append(rv, outputResult)
						}
					}
				} else {
					// now we have to evaluate the output expression
					outputResult, err := this.Output.Evaluate(innerContext)
					if err != nil {
						switch err := err.(type) {
						case *dparval.Undefined:
							// contribute nothing to the result array
						default:
							// any other error should be returned to caller
							return nil, err
						}
					} else {
						rv = append(rv, outputResult)
					}
				}
			}
		}
		return dparval.NewValue(rv), nil
	}
	return nil, &dparval.Undefined{}
}

func (this *CollectionArrayOperator) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
