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
	"github.com/couchbaselabs/tuqtng/query"
)

type Expression interface {
	Evaluate(item query.Item) (query.Value, error)
	Validate() error

	// this method takes a list of valid aliases
	// if there is more than 1 alias in the list
	// all property references MUST start with one of these aliases
	// if not, an appropriate error is returned
	// if there is only 1 alias, and the reference can be converted
	// a new expression with the proper reference is returned
	// it is up to the caller to update any references it may have
	VerifyFormalNotation(aliases []string, defaultAlias string) (Expression, error)
}

type ExpressionList []Expression
