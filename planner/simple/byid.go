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
		docId := isEqualsComparisonOnId(where)
		if docId != "" {
			return []string{docId}
		}
	case *ast.OrOperator:
		docIds := isOrOperatorComparisonOnIds(where)
		return docIds
	case *ast.CollectionAnyOperator:
		docIds := isAnyOperatorComparisonOnIds(where)
		return docIds
	}
	return nil
}

func isAnyOperatorComparisonOnIds(any *ast.CollectionAnyOperator) []string {
	// OVER must be a literal array
	docIds, ok := any.Over.(*ast.LiteralArray)
	if ok {
		// all items in the array must be a literal string
		rv := make([]string, len(docIds.Val))
		for i, docId := range docIds.Val {
			docId, ok := docId.(*ast.LiteralString)
			rv[i] = docId.Val
			if !ok {
				return nil
			}
		}

		// condition must be equals comparison
		cond, ok := any.Condition.(*ast.EqualToOperator)
		if ok {
			if isEqualsComparisonOnPropertyPath(cond, any.As) {
				return rv
			}
		}
	}
	return nil
}

func isOrOperatorComparisonOnIds(or *ast.OrOperator) []string {
	docIds := []string{}
	for _, operand := range or.Operands {
		switch operand := operand.(type) {
		case *ast.EqualToOperator:
			docId := isEqualsComparisonOnId(operand)
			if docId != "" {
				docIds = append(docIds, docId)
			} else {
				return nil
			}
		case *ast.OrOperator:
			orDocIds := isOrOperatorComparisonOnIds(operand)
			if orDocIds != nil {
				docIds = append(docIds, orDocIds...)
			} else {
				return nil
			}
		default:
			return nil
		}
	}
	return docIds
}

func isEqualsComparisonOnPropertyPath(equals *ast.EqualToOperator, path string) bool {
	prop, ok := equals.Right.(*ast.Property)
	if ok && prop.Path == path {
		// RHS is property matching any IN, need LHS to be dot member
		dotMember, ok := equals.Left.(*ast.DotMemberOperator)
		if ok {
			if isDotMemberMetaDotId(dotMember) {
				return true
			}
		}
	} else {
		prop, ok = equals.Left.(*ast.Property)
		if ok && prop.Path == path {
			// LHS is property matching any IN, need RHS to be dot member
			dotMember, ok := equals.Right.(*ast.DotMemberOperator)
			if ok {
				if isDotMemberMetaDotId(dotMember) {
					return true
				}
			}
		}
	}
	return false
}

func isEqualsComparisonOnId(equals *ast.EqualToOperator) string {
	docId, ok := equals.Right.(*ast.LiteralString)
	if ok {
		// RHS is a literal string, need LHS to be dot member
		dotMember, ok := equals.Left.(*ast.DotMemberOperator)
		if ok {
			if isDotMemberMetaDotId(dotMember) {
				return docId.Val
			}
		}
	} else {

		docId, ok = equals.Left.(*ast.LiteralString)
		if ok {
			// LHS is a literal string, need RHS to be dot member
			dotMember, ok := equals.Right.(*ast.DotMemberOperator)
			if ok {
				if isDotMemberMetaDotId(dotMember) {
					return docId.Val
				}
			}
		}
	}
	return ""
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
