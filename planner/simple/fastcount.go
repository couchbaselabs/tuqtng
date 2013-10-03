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

func CanFastCountBucket(resultExprList ast.ResultExpressionList) bool {

	// start looking at the projection
	for _, resultExpr := range resultExprList {

		// cannot be *
		if resultExpr.Star {
			return false
		}

		switch resultExpr := resultExpr.Expr.(type) {
		case *ast.FunctionCallCount:
			// aggregates all take 1 operand
			operands := resultExpr.GetOperands()
			if len(operands) < 1 {
				return false
			}
			aggOperand := operands[0]
			// must be *
			if !aggOperand.Star {
				return false
			}
		default:
			return false
		}
	}

	// if we made it this far, can do fast count on bucket
	return true
}
