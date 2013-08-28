//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

// Plan is a description of a sequence of steps to produce a correct
// result for a query.

package simple

import (
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/plan"
)

func CanIUseThisIndexForThisWhereClause(index catalog.RangeIndex, where ast.Expression, bucket string) (bool, plan.ScanRanges, ast.Expression, error) {

	// convert the index key to formal notation
	indexKeyFormal, err := IndexKeyInFormalNotation(index.Key(), bucket)
	if err != nil {
		return false, nil, nil, err
	}

	// see if the where clause expression is sargable with respect to the index key
	es := NewExpressionSargable(indexKeyFormal[0])
	where.Accept(es)
	if es.IsSargable() {
		return true, es.ScanRanges(), nil, nil
	}

	// cannot use this index
	return false, nil, nil, nil
}

func IndexKeyInFormalNotation(key catalog.IndexKey, bucket string) (catalog.IndexKey, error) {
	fkey := make(catalog.IndexKey, len(key))
	fnot := ast.NewExpressionFormalNotationConverter([]string{}, []string{bucket}, bucket)
	for i, kp := range key {
		nkey, err := kp.Accept(fnot)
		if err != nil {
			return nil, err
		}
		fkey[i] = nkey
	}
	return fkey, nil
}
