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

func init() {
	registerSystemFunction(&FunctionMeta{})
	registerSystemFunction(&FunctionValue{})
}

type FunctionMeta struct {
}

func (this *FunctionMeta) Name() string {
	return "META"
}

func (this *FunctionMeta) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {

	// FIXME the commented code below wont work until we fix how we store meta

	// av, err := arguments[0].Expr.Evaluate(item)
	// if err != nil {
	// 	return nil, err
	// }

	// meta := av.Meta()
	// if meta != nil {
	// 	metaData, err := meta.Path("meta")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return metaData, nil
	// }

	meta := item.GetAttachment("meta")
	metaValue := dparval.NewValue(meta)
	return metaValue, nil
}

func (this *FunctionMeta) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}

type FunctionValue struct {
}

func (this *FunctionValue) Name() string {
	return "VALUE"
}

func (this *FunctionValue) Evaluate(item *dparval.Value, arguments FunctionArgExpressionList) (*dparval.Value, error) {
	if len(arguments) > 0 {
		// first evaluate the argument
		av, err := arguments[0].Expr.Evaluate(item)
		if err != nil {
			return nil, err
		}
		return av, nil
	} else {
		// this mode is still relied up for projecting in the FROM clause
		// review for cleanup
		return item, nil
	}
}

func (this *FunctionValue) Validate(arguments FunctionArgExpressionList) error {
	err := ValidateArity(this, arguments, 1, 1)
	if err != nil {
		return err
	}
	return ValidateNoStars(this, arguments)
}
