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

type AggregateFunctionCall struct {
	FunctionCall
}

// create a unique key where the current value for this
// aggregate function will be stored
func (this AggregateFunctionCall) Key() string {
	return fmt.Sprintf("%s-%t-%v-%t", this.FunctionCall.Name, this.FunctionCall.Operands[0].Star, this.FunctionCall.Operands[0].Expr, this.FunctionCall.Distinct)
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
