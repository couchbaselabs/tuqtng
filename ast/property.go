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

type Property struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func NewProperty(path string) *Property {
	return &Property{
		Type: "property",
		Path: path,
	}
}

func (this *Property) Copy() Expression {
	return &Property{
		Type: "property",
		Path: this.Path,
	}
}

func (this *Property) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	if item == nil {
		return nil, &dparval.Undefined{this.Path}
	}
	pv, err := item.Path(this.Path)
	if err != nil {
		return nil, err
	}
	return pv, nil
}

func (this *Property) String() string {
	return fmt.Sprintf("%s", this.Path)
}

func (this *Property) EquivalentTo(t Expression) bool {
	that, ok := t.(*Property)
	if !ok {
		return false
	}

	if this.Path != that.Path {
		return false
	}

	return true
}

func (this *Property) Dependencies() ExpressionList {
	rv := ExpressionList{}
	return rv
}

func (this *Property) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}
