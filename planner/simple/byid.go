//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package simple

import (
	"github.com/couchbaselabs/tuqtng/ast"
)

func WhereClauseFindById(where ast.Expression) []string {
	switch where := where.(type) {
	case *ast.EqualToOperator:
		docId, ok := where.Right.(*ast.LiteralString)
		if ok {
			// RHS is a literal string, need LHS to be dot member
			dotMember, ok := where.Left.(*ast.DotMemberOperator)
			if ok {
				if isDotMemberMetaDotId(dotMember) {
					return []string{docId.Val}
				}
			}
		} else {

			docId, ok = where.Left.(*ast.LiteralString)
			if ok {
				// LHS is a literal string, need RHS to be dot member
				dotMember, ok := where.Right.(*ast.DotMemberOperator)
				if ok {
					if isDotMemberMetaDotId(dotMember) {
						return []string{docId.Val}
					}
				}
			}
		}
	}
	return nil
}

func isDotMemberMetaDotId(dotMember *ast.DotMemberOperator) bool {
	if dotMember.Right.Path == "id" {
		// need LHS to function call META()
		_, ok := dotMember.Left.(*ast.FunctionCallMeta)
		if ok {
			// every thing is in order, this where clause is just matching a single ID
			return true
		}
	}
	return false
}
