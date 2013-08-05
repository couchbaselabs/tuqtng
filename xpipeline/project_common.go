//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
)

func projectedValueOfResultExpression(item *dparval.Value, resultExpr *ast.ResultExpression) (*dparval.Value, error) {

	if resultExpr.Star {
		if resultExpr.Expr != nil {
			// evaluate this expression first
			val, err := resultExpr.Expr.Evaluate(item)
			if err != nil {
				return nil, err
			}
			if val.Type() == dparval.OBJECT {
				return val, nil
			}
		} else {
			return item, nil
		}
	} else if resultExpr.Expr != nil {
		// evaluate the expression
		val, err := resultExpr.Expr.Evaluate(item)
		if err != nil {
			return nil, err
		}
		return val, err
	}

	return nil, nil
}
