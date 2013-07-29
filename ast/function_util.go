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
	registerSystemFunction(&FunctionLength{})
}

type FunctionLength struct {
}

func (this *FunctionLength) Name() string {
	return "LENGTH"
}

func (this *FunctionLength) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {
	// first evaluate the argument
	av, err := arguments[0].Expr.Evaluate(item)

	// the spec does not define it to operate on missing, so return null
	if err != nil {
		switch err := err.(type) {
		case *query.Undefined:
			// undefined returns null
			return nil, nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	switch av := av.(type) {
	case map[string]query.Value:
		return float64(len(av)), nil
	case []query.Value:
		return float64(len(av)), nil
	case string:
		return float64(len(av)), nil
	default:
		return nil, nil
	}
}

func (this *FunctionLength) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) != 1 {
		return fmt.Errorf("the LENGTH() function takes exactly one argument")
	}

	return ValidateNoStars(this, arguments)
}
