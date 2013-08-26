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

type ExpressionSimplifier struct {
}

func NewExpressionSimplifier() *ExpressionSimplifier {
	return &ExpressionSimplifier{}
}

func (this *ExpressionSimplifier) Visit(e Expression) (Expression, error) {

	// see if the expression depends on anything
	dc := NewExpressionFunctionalDependencyCheckerFull(ExpressionList{})

	_, err := e.Accept(dc)
	if err != nil {
		// it depends on something, try the children
		return VisitChildren(this, e)
	} else {
		// it doesn't depend on anything, so lets evaluate it with the empty object
		val, err := e.Evaluate(dparval.NewValue(map[string]interface{}{}))
		if err != nil {
			switch err := err.(type) {
			case *dparval.Undefined:
				// we have no literal missing, so leave this unchanged
				return e, nil
			default:
				return nil, err
			}
		}
		return this.newLiteralFromValue(val.Value()), nil
	}
}

func (this *ExpressionSimplifier) newLiteralFromValue(value interface{}) Expression {
	switch value := value.(type) {
	case nil:
		return NewLiteralNull()
	case bool:
		return NewLiteralBool(value)
	case float64:
		return NewLiteralNumber(value)
	case string:
		return NewLiteralString(value)
	case []interface{}:
		contents := make(ExpressionList, len(value))
		for i, element := range value {
			contents[i] = this.newLiteralFromValue(element)
		}
		return NewLiteralArray(contents)
	case map[string]interface{}:
		contents := make(map[string]Expression, len(value))
		for k, element := range value {
			contents[k] = this.newLiteralFromValue(element)
		}
		return NewLiteralObject(contents)
	default:
		panic(fmt.Sprintf("Unable to create literal for %v", value))
	}
}
