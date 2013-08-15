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

type FunctionCallUnknown struct {
	FunctionCall
}

func NewFunctionCallUnknown(operands FunctionArgExpressionList, name string) FunctionCallExpression {
	return &FunctionCallUnknown{
		FunctionCall{
			Type:     "function",
			Name:     name,
			Operands: operands,
		},
	}
}

func (this *FunctionCallUnknown) Arity() (int, int) {
	return -1, -1
}

func (this *FunctionCallUnknown) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	return nil, fmt.Errorf("Unknown Function %s", this.Name)
}

func (this *FunctionCallUnknown) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
