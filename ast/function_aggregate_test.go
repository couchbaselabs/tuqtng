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
	"reflect"
	"testing"

	"github.com/couchbaselabs/dparval"
)

type AggregateTest struct {
	function Expression
	result   *dparval.Value
}

type AggregateTestSet []AggregateTest

func (this AggregateTestSet) Run(t *testing.T, dataset dparval.ValueCollection) {
	for _, x := range this {
		af, ok := x.function.(AggregateFunctionCallExpression)
		if !ok {
			t.Errorf("Require an aggregate function call expression to run this test")
		}
		group := dparval.NewValue(map[string]interface{}{})
		af.DefaultAggregate(group)
		for _, d := range dataset {
			af.UpdateAggregate(group, d)
		}
		groupMeta := group.GetAttachment("aggregates").(map[string]interface{})
		final := groupMeta[af.Key()]

		if !reflect.DeepEqual(final, x.result) {
			t.Errorf("Expected %v, got %v, for %v", x.result, final, af)
		}
	}
}

func TestCount(t *testing.T) {

	dataset := dparval.ValueCollection{
		dparval.NewValue(map[string]interface{}{
			"name":   "marty",
			"status": "alive",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":   "gerald",
			"status": nil,
		}),
		dparval.NewValue(map[string]interface{}{
			"name": "steve",
		}),
	}

	tests := AggregateTestSet{
		// test * (counts all rows)
		{
			NewFunctionCall("COUNT", FunctionArgExpressionList{NewStarFunctionArgExpression()}),
			dparval.NewValue(3.0),
		},
		// test expression (eliminiates null and missing)
		{
			NewFunctionCall("COUNT", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("status"))}),
			dparval.NewValue(1.0),
		},
	}

	tests.Run(t, dataset)

}

// sum and avg have same elimination rules, test them together
func TestSumAndAvg(t *testing.T) {

	dataset := dparval.ValueCollection{
		dparval.NewValue(map[string]interface{}{
			"name":  "marty",
			"score": 20.0,
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "gerald",
			"score": nil,
		}),
		dparval.NewValue(map[string]interface{}{
			"name": "steve",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "siri",
			"score": "thirty",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "deep",
			"score": 10.0,
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "ketaki",
			"score": "false",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "pratap",
			"score": []interface{}{5.5},
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "karen",
			"score": map[string]interface{}{"score": 5.5},
		}),
	}

	tests := AggregateTestSet{
		// test expression (eliminiates null and missing)
		{
			NewFunctionCall("SUM", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("score"))}),
			dparval.NewValue(30.0),
		},
		{
			NewFunctionCall("AVG", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("score"))}),
			dparval.NewValue(15.0),
		},
	}

	tests.Run(t, dataset)

}

// sum and avg have same elimination rules, test them together
func TestMinAndMax(t *testing.T) {

	dataset := dparval.ValueCollection{
		dparval.NewValue(map[string]interface{}{
			"name":  "marty",
			"score": 20.0,
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "gerald",
			"score": nil,
		}),
		dparval.NewValue(map[string]interface{}{
			"name": "steve",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "siri",
			"score": "thirty",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "deep",
			"score": 10.0,
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "ketaki",
			"score": "false",
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "pratap",
			"score": []interface{}{5.5},
		}),
		dparval.NewValue(map[string]interface{}{
			"name":  "karen",
			"score": map[string]interface{}{"score": 5.5},
		}),
	}

	tests := AggregateTestSet{
		// test expression (eliminiates null and missing)
		{
			NewFunctionCall("MIN", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("score"))}),
			dparval.NewValue(10.0),
		},
		{
			NewFunctionCall("MAX", FunctionArgExpressionList{NewFunctionArgExpression(NewProperty("score"))}),
			dparval.NewValue(map[string]interface{}{"score": 5.5}),
		},
	}

	tests.Run(t, dataset)

}
