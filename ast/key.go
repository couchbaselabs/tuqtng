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

type KeyExpression struct {
	Expr Expression `json:"expr"`
	Type string     `json:"type"`
	Keys []string   `json:"keys"`
}

func NewKeyExpression(expr Expression, typeof string) *KeyExpression {
	return &KeyExpression{
		Expr: expr,
		Type: typeof,
	}
}

func (this *KeyExpression) Validate() error {
	var err error

	validator := NewExpressionValidatorNoAggregates()
	this.Expr, err = this.Expr.Accept(validator)
	if err != nil {
		return err
	}

	val, err := this.Expr.Evaluate(dparval.NewValue(map[string]interface{}{}))

	if val.Type() == dparval.ARRAY && this.Type == "KEY" {
		return fmt.Errorf("KEY expression used with multiple values")
	}

	if val.Type() == dparval.STRING && this.Type == "KEYS" {
		return fmt.Errorf("KEYS expression used with a single value")
	}

	// create the keylist here since we have evaluated the expression
	this.Keys = make([]string, 0)

	if val.Type() == dparval.ARRAY {
		keylist := val.Value()
		for _, key := range keylist.([]interface{}) {
			this.Keys = append(this.Keys, key.(string))
		}
	}

	if val.Type() == dparval.STRING {
		this.Keys = append(this.Keys, val.Value().(string))
	}

	return nil
}

func (this *KeyExpression) GetKeys() []string {
	return this.Keys
}

func (this *KeyExpression) String() string {
	return fmt.Sprintf("%v", this.Expr)
}
