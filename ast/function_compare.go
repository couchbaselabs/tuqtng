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
	registerSystemFunction("GREATEST", &FunctionGreatest{})
	registerSystemFunction("LEAST", &FunctionLeast{})
}

type FunctionGreatest struct{}

func (this *FunctionGreatest) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	var rv query.Value = nil

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// undefined doesn't change the result
			default:
				// any other error return to caller
				return nil, err
			}
		}

		// now compare this value with rv
		compres := CollateJSON(av, rv)
		if compres > 0 {
			rv = av
		}
	}

	return rv, nil
}

func (this *FunctionGreatest) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the GREATEST() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the GREATEST() function does not support *")
	}
	return nil
}

type FunctionLeast struct{}

func (this *FunctionLeast) Evaluate(item query.Item, arguments FunctionArgExpressionList) (query.Value, error) {

	var rv query.Value = nil
	first := true

	for _, arg := range arguments {
		av, err := arg.Expr.Evaluate(item)
		if err != nil {
			switch err := err.(type) {
			case *query.Undefined:
				// undefined doesn't change the result
			default:
				// any other error return to caller
				return nil, err
			}
		}

		// now compare this value with rv
		compres := CollateJSON(av, rv)
		if compres < 0 || first {
			first = false
			rv = av
		}
	}

	return rv, nil
}

func (this *FunctionLeast) Validate(arguments FunctionArgExpressionList) error {
	if len(arguments) < 1 {
		return fmt.Errorf("the LEAST() function expects at least one argument")
	}
	if arguments[0].Star == true {
		return fmt.Errorf("the LEAST() function does not support *")
	}
	return nil
}
