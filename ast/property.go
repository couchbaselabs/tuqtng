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
	"log"

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

func (this *Property) Validate() error {
	return nil
}

// NOTE this should only be called on leading property references, not ones deeper in a chain (ie, for a.b.c, only a, not b or c)
func (this *Property) VerifyFormalNotation(forbiddenAliases []string, aliases []string, defaultAlias string) (Expression, error) {
	// check to see if any of the forbiddenAliases are mentioned
	if len(forbiddenAliases) > 0 {
		for _, forbiddenAlias := range forbiddenAliases {
			if this.Path == forbiddenAlias {
				log.Printf("forbiddenAliases were %v", forbiddenAliases)
				return nil, fmt.Errorf("Alias %s cannot be referenced", this.Path)
			}
		}
	}
	// this test is not needed when there are no aliases in the from clause (expression evaluation only)
	if len(aliases) > 0 {
		for _, alias := range aliases {
			if this.Path == alias {
				return nil, nil
			}
		}
		if defaultAlias != "" {
			return NewDotMemberOperator(NewProperty(defaultAlias), this), nil
		} else {
			return nil, fmt.Errorf("Property reference %s missing qualifier bucket/alias", this.Path)
		}
	}
	return nil, nil
}
