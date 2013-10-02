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

type FunctionCallCount struct {
	AggregateFunctionCall
}

func NewFunctionCallCount(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallCount{
		AggregateFunctionCall{
			FunctionCall{
				Type:     "function",
				Name:     "COUNT",
				Operands: operands,
				minArgs:  1,
				maxArgs:  1,
			},
		},
	}
}

func (this *FunctionCallCount) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	aggregate_key := this.Key()
	val, err := aggregateValue(item, aggregate_key)
	return val, err
}

func (this *FunctionCallCount) DefaultAggregate(group *dparval.Value) error {
	aggregate_key := this.Key()
	if this.Distinct {
		aggregate_unique_key := aggregate_key + "_unique"
		uniqueness_map := dparval.NewValue(map[string]interface{}{})
		setAggregateValue(group, aggregate_unique_key, uniqueness_map)
	}
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(0.0)
		// store this, so that even if all values are eliminated we return 0
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionCallCount) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}
	currentValVal := currentVal.Value()
	currentFloat, ok := currentValVal.(float64)
	if !ok {
		return fmt.Errorf("count value not a number")
	}

	if this.Operands[0].Star && this.Operands[0].Expr == nil {
		// pure star
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentFloat+1))
	} else if this.Operands[0].Star && this.Operands[0].Expr != nil {
		// dot star
		val, err := this.Operands[0].Expr.Evaluate(item)
		val, err = eliminateNonObject(val, err)
		if err != nil {
			return err
		}
		if val != nil {

			if this.Distinct {
				uniqueness_map, err := aggregateValue(group, aggregate_key+"_unique")
				if err != nil {
					return fmt.Errorf("group uniqueness defaults not set correctly")
				}
				// check to see if we already have this value
				valkey := string(val.Bytes())
				exists, _ := uniqueness_map.Path(valkey)
				if exists != nil {
					return nil
				} else {
					uniqueness_map.SetPath(valkey, true)
					setAggregateValue(group, aggregate_key+"_unique", uniqueness_map)
					// and allow it to continue
				}
			}

			setAggregateValue(group, aggregate_key, dparval.NewValue(currentFloat+1))
		}
	} else if !this.Operands[0].Star && this.Operands[0].Expr != nil {
		// no star
		val, err := this.Operands[0].Expr.Evaluate(item)
		val, err = eliminateNullMissing(val, err)
		if err != nil {
			return err
		}
		if val != nil {

			if this.Distinct {
				uniqueness_map, err := aggregateValue(group, aggregate_key+"_unique")
				if err != nil {
					return fmt.Errorf("group uniqueness defaults not set correctly")
				}
				// check to see if we already have this value
				valkey := string(val.Bytes())
				exists, _ := uniqueness_map.Path(valkey)
				if exists != nil {
					return nil
				} else {
					uniqueness_map.SetPath(valkey, true)
					setAggregateValue(group, aggregate_key+"_unique", uniqueness_map)
					// and allow it to continue
				}
			}

			setAggregateValue(group, aggregate_key, dparval.NewValue(currentFloat+1))
		}
	}
	return nil
}

func (this *FunctionCallCount) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallSum struct {
	AggregateFunctionCall
}

func NewFunctionCallSum(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallSum{
		AggregateFunctionCall{
			FunctionCall{
				Type:     "function",
				Name:     "SUM",
				Operands: operands,
				minArgs:  1,
				maxArgs:  1,
			},
		},
	}
}

func (this *FunctionCallSum) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	aggregate_key := this.Key()
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionCallSum) DefaultAggregate(group *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionCallSum) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if this.Operands[0].Expr != nil {
		val, err := this.Operands[0].Expr.Evaluate(item)
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

func (this *FunctionCallSum) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallAvg struct {
	AggregateFunctionCall
}

func NewFunctionCallAvg(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallAvg{
		AggregateFunctionCall{
			FunctionCall{
				Type:     "function",
				Name:     "AVG",
				Operands: operands,
				minArgs:  1,
				maxArgs:  1,
			},
		},
	}
}

func (this *FunctionCallAvg) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	aggregate_key := this.Key()
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionCallAvg) DefaultAggregate(group *dparval.Value) error {
	aggregate_key := this.Key()
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

func (this *FunctionCallAvg) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	aggregate_key := this.Key()
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

	if this.Operands[0].Expr != nil {
		val, err := this.Operands[0].Expr.Evaluate(item)
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

func (this *FunctionCallAvg) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallMin struct {
	AggregateFunctionCall
}

func NewFunctionCallMin(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallMin{
		AggregateFunctionCall{
			FunctionCall{
				Type:     "function",
				Name:     "MIN",
				Operands: operands,
				minArgs:  1,
				maxArgs:  1,
			},
		},
	}
}

func (this *FunctionCallMin) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	aggregate_key := this.Key()
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionCallMin) DefaultAggregate(group *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionCallMin) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if this.Operands[0].Expr != nil {
		val, err := this.Operands[0].Expr.Evaluate(item)
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

func (this *FunctionCallMin) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallMax struct {
	AggregateFunctionCall
}

func NewFunctionCallMax(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallMax{
		AggregateFunctionCall{
			FunctionCall{
				Type:     "function",
				Name:     "MAX",
				Operands: operands,
				minArgs:  1,
				maxArgs:  1,
			},
		},
	}
}

func (this *FunctionCallMax) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	aggregate_key := this.Key()
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionCallMax) DefaultAggregate(group *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue(nil)
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionCallMax) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if this.Operands[0].Expr != nil {
		val, err := this.Operands[0].Expr.Evaluate(item)
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

func (this *FunctionCallMax) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallArrayAgg struct {
	AggregateFunctionCall
}

func NewFunctionCallArrayAgg(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallArrayAgg{
		AggregateFunctionCall{
			FunctionCall{
				Type:     "function",
				Name:     "ARRAY_AGG",
				Operands: operands,
				minArgs:  1,
				maxArgs:  1,
			},
		},
	}
}

func (this *FunctionCallArrayAgg) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	aggregate_key := this.Key()
	return aggregateValue(item, aggregate_key)
}

func (this *FunctionCallArrayAgg) DefaultAggregate(group *dparval.Value) error {
	aggregate_key := this.Key()
	if this.Distinct {
		aggregate_unique_key := aggregate_key + "_unique"
		uniqueness_map := dparval.NewValue(map[string]interface{}{})
		setAggregateValue(group, aggregate_unique_key, uniqueness_map)
	}
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		currentVal = dparval.NewValue([]interface{}{})
		// store this, so that even if all values are eliminated we return null
		setAggregateValue(group, aggregate_key, dparval.NewValue(currentVal))
	}
	return nil
}

func (this *FunctionCallArrayAgg) UpdateAggregate(group *dparval.Value, item *dparval.Value) error {
	aggregate_key := this.Key()
	currentVal, err := aggregateValue(group, aggregate_key)
	if err != nil {
		return fmt.Errorf("group defaults not set correctly")
	}

	if this.Operands[0].Expr != nil {
		val, err := this.Operands[0].Expr.Evaluate(item)
		if err != nil {
			// only eliminate missing
			return err
		}
		if val != nil {

			nextVal := val.Value()

			if this.Distinct {
				uniqueness_map, err := aggregateValue(group, aggregate_key+"_unique")
				if err != nil {
					return fmt.Errorf("group uniqueness defaults not set correctly")
				}
				// check to see if we already have this value
				valkey := string(val.Bytes())
				exists, _ := uniqueness_map.Path(valkey)
				if exists != nil {
					return nil
				} else {
					uniqueness_map.SetPath(valkey, true)
					setAggregateValue(group, aggregate_key+"_unique", uniqueness_map)
					// and allow it to continue
				}
			}

			currVal := currentVal.Value()

			currValArray, ok := currVal.([]interface{})
			if ok {
				currValArray = append(currValArray, nextVal)
			}
			setAggregateValue(group, aggregate_key, dparval.NewValue(currValArray))

		}
	}
	return nil
}

func (this *FunctionCallArrayAgg) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
