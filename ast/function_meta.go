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

	"github.com/couchbaselabs/tuqtng/query"
)

func init() {
	registerSystemFunction("META", &FunctionMeta{})
	registerSystemFunction("VALUE", &FunctionValue{})
}

type FunctionMeta struct {
}

func (this *FunctionMeta) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	return item.GetMeta(), nil
}

func (this *FunctionMeta) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) > 0 {
		return fmt.Errorf("the META() function takes no arguments")
	}
	return nil
}

type FunctionValue struct {
}

func (this *FunctionValue) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	return item.GetValue(), nil
}

func (this *FunctionValue) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) > 0 {
		return fmt.Errorf("the VALUE() function takes no arguments")
	}
	return nil
}
