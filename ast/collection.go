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
	"github.com/couchbaselabs/tuqtng/query"
)

type CollectionAnyOperator struct {
	Condition Expression
	Over      Expression
	As        string
}

func NewCollectionAnyOperator(condition Expression, over Expression, as string) *CollectionAnyOperator {
	return &CollectionAnyOperator{
		Condition: condition,
		Over:      over,
		As:        as,
	}
}

func (this *CollectionAnyOperator) Evaluate(item query.Item) (query.Value, error) {
	// first evaluate the over
	ov, err := this.Over.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *query.Undefined:
			// spec says return false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	// see if we're dealing with an array
	switch ov := ov.(type) {
	case []query.Value:
		// yes, now we need to walk through the elements
		for _, inner := range ov {
			// create a new context with this object named as the alias
			innerContext := query.NewParsedItem(map[string]query.Value{this.As: inner}, nil)
			// now evaluate the condition in this new context
			innerResult, err := this.Condition.Evaluate(innerContext)
			if err != nil {
				switch err := err.(type) {
				case *query.Undefined:
					// this is not true, keep trying
					continue
				default:
					// any other error should be returned to caller
					return nil, err
				}
			}
			// check to see if this value is true
			innerBoolResult := ValueInBooleanContext(innerResult)
			if innerBoolResult == true {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, nil
	}
}

func (this *CollectionAnyOperator) Validate() error {
	err := this.Condition.Validate()
	if err != nil {
		return err
	}
	err = this.Over.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *CollectionAnyOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newover, err := this.Over.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newover != nil {
		this.Over = newover
	}

	updatedAliases := append(aliases, this.As)
	newcond, err := this.Condition.VerifyFormalNotation(updatedAliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newcond != nil {
		this.Condition = newcond
	}
	return nil, nil
}

type CollectionAllOperator struct {
	Condition Expression
	Over      Expression
	As        string
}

func NewCollectionAllOperator(condition Expression, over Expression, as string) *CollectionAllOperator {
	return &CollectionAllOperator{
		Condition: condition,
		Over:      over,
		As:        as,
	}
}

func (this *CollectionAllOperator) Evaluate(item query.Item) (query.Value, error) {
	// first evaluate the over
	ov, err := this.Over.Evaluate(item)
	if err != nil {
		switch err := err.(type) {
		case *query.Undefined:
			// spec says return false
			return false, nil
		default:
			// any other error should be returned to caller
			return nil, err
		}
	}

	// see if we're dealing with an array
	switch ov := ov.(type) {
	case []query.Value:
		// yes, now we need to walk through the elements
		for _, inner := range ov {
			// create a new context with this object named as the alias
			innerContext := query.NewParsedItem(map[string]query.Value{this.As: inner}, nil)
			// now evaluate the condition in this new context
			innerResult, err := this.Condition.Evaluate(innerContext)
			if err != nil {
				switch err := err.(type) {
				case *query.Undefined:
					// not true, stop looking
					return false, nil
				default:
					// any other error should be returned to caller
					return nil, err
				}
			}
			// check to see if this value is true
			innerBoolResult := ValueInBooleanContext(innerResult)
			if innerBoolResult == false {
				return false, nil
			}
		}
		return true, nil
	default:
		return false, nil
	}
}

func (this *CollectionAllOperator) Validate() error {
	err := this.Condition.Validate()
	if err != nil {
		return err
	}
	err = this.Over.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *CollectionAllOperator) VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error) {
	newover, err := this.Over.VerifyFormalNotation(aliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newover != nil {
		this.Over = newover
	}

	updatedAliases := append(aliases, this.As)
	newcond, err := this.Condition.VerifyFormalNotation(updatedAliases, defaultAlias)
	if err != nil {
		return nil, err
	}
	if newcond != nil {
		this.Condition = newcond
	}
	return nil, nil
}
