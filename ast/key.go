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
	"strings"
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

	this.Keys = this.GetKeys()
	if this.Keys == nil {
		//expr := this.Expr.Evaluate(this.Expr.value())
	}

	if this.Keys != nil && this.Type == "KEY" && len(this.Keys) > 1 {
		//array specifed for a single key
		return fmt.Errorf("KEY expression used with multiple values")
	}
	if this.Keys != nil && this.Type == "KEYS" && len(this.Keys) < 2 {
		return fmt.Errorf("KEYS expression used with a single value")
	}

	return nil
}

func (this *KeyExpression) String() string {
	return fmt.Sprintf("%v", this.Expr)
}

func (this *KeyExpression) GetKeys() []string {
	keylist := this.String()
	if strings.Contains(keylist, "\"") {
		// expr is either a string or array literal
		keylist = strings.Trim(keylist, "[")
		keylist = strings.Trim(keylist, "]")
		this.Keys = strings.Split(keylist, ",")
		for i, key := range this.Keys {
			key = strings.TrimSpace(key)
			this.Keys[i] = strings.Trim(key, "\"")
		}
		return this.Keys
	}
	return nil
}
