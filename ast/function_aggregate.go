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

func init() {
	registerSystemFunction(&FunctionCount{})
	registerSystemFunction(&FunctionSum{})
	registerSystemFunction(&FunctionAvg{})
	registerSystemFunction(&FunctionMin{})
	registerSystemFunction(&FunctionMax{})
}

// create a unique key where the current value for this
// aggregate function will be stored
func aggregateKey(f SystemFunction, arg *FunctionArgExpression) string {
	return fmt.Sprintf("%s-%t-%v", f.Name(), arg.Star, arg.Expr)
}

// lookup the current value of the aggregate stored
// at the specified key
func aggregateValue(item *dparval.Value, key string) (*dparval.Value, error) {
	aggregates := item.GetAttachment("aggregates")
	aggregatesMap, ok := aggregates.(map[string]interface{})
	if ok {
		value := aggregatesMap[key]
		valuedparval, ok := value.(*dparval.Value)
		if ok {
			return valuedparval, nil
		}
	}
	return nil, fmt.Errorf("Unable to find aggregate")
}

func setAggregateValue(item *dparval.Value, key string, val *dparval.Value) {
	aggregates := item.GetAttachment("aggregates")
	aggregatesMap, ok := aggregates.(map[string]interface{})
	if !ok {
		// create a new aggregates map
		aggregatesMap = map[string]interface{}{}
		item.SetAttachment("aggregates", aggregatesMap)
	}
	aggregatesMap[key] = val
}

func eliminateNullMissing(val *dparval.Value, err error) (*dparval.Value, error) {
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// missing elimination
			return nil, nil
		default:
			return nil, err
		}
	}
	if val.Type() == dparval.NULL {
		// null elimination
		return nil, nil
	}
	return val, nil
}

func eliminateNonNumber(val *dparval.Value, err error) (*dparval.Value, error) {
	val, err = eliminateNullMissing(val, err)
	if err != nil {
		return nil, err
	}
	if val != nil {
		if val.Type() == dparval.NUMBER {
			return val, nil
		}
	}
	return nil, nil
}

func eliminateNonObject(val *dparval.Value, err error) (*dparval.Value, error) {
	val, err = eliminateNullMissing(val, err)
	if err != nil {
		return nil, err
	}
	if val != nil {
		if val.Type() == dparval.OBJECT {
			return val, nil
		}
	}
	return nil, nil
}

type FunctionCount struct {
}

func (this *FunctionCount) Name() string {
	return "COUNT"
}

func (this *FunctionCount) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	aggregate_key := aggregateKey(this, arguments[0])
	val, err := aggregateValue(item, aggregate_key)
	return val, err
}

func (this *FunctionCount) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return nil
}

func (this *FunctionCount) IsAggregate() bool {
	return true
}

func (this *FunctionCount) DefaultAggregate(group *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(0.0)
		// store this, so that even if all values are eliminated we return 0
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionCount) UpdateAggregate(group *dparval.Value, item *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}
	currentValVal := currentVal.Value()
	currentFloat, ok := currentValVal.(float64)
	if !ok {
		return fmt.Errorf("count value not a number")
	}
	if arguments[0].Star && arguments[0].Expr == nil {
		// pure star
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentFloat+1))
	} else if arguments[0].Star && arguments[0].Expr != nil {
		// dot star
		val, err := arguments[0].Expr.Evaluate(item)
		val, err = eliminateNonObject(val, err)
		if err != nil {
			return err
		}
		if val != nil {
			setAggregateValue(group, aggregate_key, dparval.NewValue(currentFloat+1))
		}
	} else if !arguments[0].Star && arguments[0].Expr != nil {
		// no star
		val, err := arguments[0].Expr.Evaluate(item)
		val, err = eliminateNullMissing(val, err)
		if err != nil {
			return err
		}
		if val != nil {
			setAggregateValue(group, aggregate_key, dparval.NewValue(currentFloat+1))
		}
	}
	return nil
}

type FunctionSum struct {
}

func (this *FunctionSum) Name() string {
	return "SUM"
}

func (this *FunctionSum) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	aggregate_key := aggregateKey(this, arguments[0])
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionSum) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionSum) IsAggregate() bool {
	return true
}

func (this *FunctionSum) DefaultAggregate(group *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionSum) UpdateAggregate(group *dparval.Value, item *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if arguments[0].Expr != nil {
		val, err := arguments[0].Expr.Evaluate(item)
		val, err = eliminateNonNumber(val, err)
		if err != nil {
			return err
		}
		if val != nil {
			nextVal := val.Value()
			if currentVal.Type() == dparval.NUMBER {
				currentValVal := currentVal.Value()
				currentFloat, ok := currentValVal.(float64)
				if !ok {
					return fmt.Errorf("count value not a number")
				}
				nextVal = nextVal.(float64) + currentFloat
			}
			setAggregateValue(group, aggregate_key, dparval.NewValue(nextVal))
		}
	}
	return nil
}

type FunctionAvg struct {
}

func (this *FunctionAvg) Name() string {
	return "AVG"
}

func (this *FunctionAvg) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	aggregate_key := aggregateKey(this, arguments[0])
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionAvg) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionAvg) IsAggregate() bool {
	return true
}

func (this *FunctionAvg) DefaultAggregate(group *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	// avg needs to track sum and count to produce its value
	count_key := aggregate_key + "_count"
	sum_key := aggregate_key + "_sum"
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}

	currentCount, err := aggregateValue(group, count_key)
	if err != nil {
		currentCount = dparval.NewValue(0.0)
		setAggregateValue(group, count_key, dparval.NewValue(currentCount))
	}

	currentSum, err := aggregateValue(group, sum_key)
	if err != nil {
		currentSum = dparval.NewValue(0.0)
		setAggregateValue(group, sum_key, dparval.NewValue(currentSum))
	}

	return nil
}

func (this *FunctionAvg) UpdateAggregate(group *dparval.Value, item *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	// avg needs to track sum and count to produce its value
	count_key := aggregate_key + "_count"
	sum_key := aggregate_key + "_sum"
	_, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	currentCount, err := aggregateValue(group, count_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	currentSum, err := aggregateValue(group, sum_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if arguments[0].Expr != nil {
		val, err := arguments[0].Expr.Evaluate(item)
		val, err = eliminateNonNumber(val, err)
		if err != nil {
			return err
		}
		if val != nil {
			nextVal := val.Value()
			nextValFloat := nextVal.(float64)

			if currentCount.Type() == dparval.NUMBER {
				currentCountVal := currentCount.Value()
				currentCountFloat, ok := currentCountVal.(float64)
				if !ok {
					return fmt.Errorf("count value not a number")
				}
				if currentSum.Type() == dparval.NUMBER {
					currentSumVal := currentSum.Value()
					currentSumFloat, ok := currentSumVal.(float64)
					if !ok {
						return fmt.Errorf("sum value not a number")
					}
					nextCountFloat := currentCountFloat + 1
					nextSumFloat := currentSumFloat + nextValFloat
					nextVal := nextSumFloat / nextCountFloat

					setAggregateValue(group, count_key, dparval.NewValue(nextCountFloat))
					setAggregateValue(group, sum_key, dparval.NewValue(nextSumFloat))
					setAggregateValue(group, aggregate_key, dparval.NewValue(nextVal))
				}
			}

		}
	}
	return nil
}

type FunctionMin struct {
}

func (this *FunctionMin) Name() string {
	return "MIN"
}

func (this *FunctionMin) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	aggregate_key := aggregateKey(this, arguments[0])
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionMin) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionMin) IsAggregate() bool {
	return true
}

func (this *FunctionMin) DefaultAggregate(group *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionMin) UpdateAggregate(group *dparval.Value, item *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if arguments[0].Expr != nil {
		val, err := arguments[0].Expr.Evaluate(item)
		val, err = eliminateNullMissing(val, err)
		if err != nil {
			return err
		}
		if val != nil {

			nextVal := val.Value()
			currVal := currentVal.Value()

			if currVal == nil {
				// any value is greater than nil (we eliminated null/mising already)
				setAggregateValue(group, aggregate_key, dparval.NewValue(nextVal))
			} else {
				// check to see
				comp := CollateJSON(nextVal, currVal)
				if comp < 0 {
					setAggregateValue(group, aggregate_key, dparval.NewValue(nextVal))
				}
			}
		}
	}
	return nil
}

type FunctionMax struct {
}

func (this *FunctionMax) Name() string {
	return "MAX"
}

func (this *FunctionMax) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	aggregate_key := aggregateKey(this, arguments[0])
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionMax) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

func (this *FunctionMax) IsAggregate() bool {
	return true
}

func (this *FunctionMax) DefaultAggregate(group *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionMax) UpdateAggregate(group *dparval.Value, item *dparval.Value, arguments FunctionArgExpressionList) error {
	aggregate_key := aggregateKey(this, arguments[0])
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if arguments[0].Expr != nil {
		val, err := arguments[0].Expr.Evaluate(item)
		val, err = eliminateNullMissing(val, err)
		if err != nil {
			return err
		}
		if val != nil {

			nextVal := val.Value()
			currVal := currentVal.Value()

			// check to see
			comp := CollateJSON(nextVal, currVal)
			if comp > 0 {
				setAggregateValue(group, aggregate_key, dparval.NewValue(nextVal))
			}

		}
	}
	return nil
}
