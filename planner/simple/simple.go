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
	"fmt"

	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
)

type SimplePlanner struct {
	pool catalog.Pool
}

func NewSimplePlanner(pool catalog.Pool) *SimplePlanner {
	return &SimplePlanner{
		pool: pool,
	}
}

func (this *SimplePlanner) Plan(stmt ast.Statement) (plan.PlanChannel, query.ErrorChannel) {
	pc := make(plan.PlanChannel)
	ec := make(query.ErrorChannel)
	go this.buildPlans(stmt, pc, ec)
	return pc, ec
}

func (this *SimplePlanner) buildPlans(stmt ast.Statement, pc plan.PlanChannel, ec query.ErrorChannel) {
	defer close(pc)
	defer close(ec)

	from := stmt.GetFrom()

	var lastStep plan.PlanElement

	if from == nil {
		// simple expression evaluation
		lastStep = plan.NewExpressionEvaluator()

	} else {

		// see if the bucket exists
		if this.pool != nil {
			bucket, err := this.pool.Bucket(from.Bucket)
			if err != nil {
				ec <- query.NewBucketDoesNotExist(from.Bucket)
				return
			}

			// find all docs scanner
			scanners, err := bucket.Scanners()
			if err != nil {
				ec <- query.NewError(err, fmt.Sprintf("No usable scanner found for bucket %v", from.Bucket))
				return
			}

			foundUsableScanner := false
			for _, scanner := range scanners {
				switch scanner.(type) {
				case catalog.FullScanner:
					lastStep = plan.NewScan(bucket.Name(), scanner.Name())
					lastStep = plan.NewFetch(lastStep, bucket.Name(), from.Projection, from.As)
					nextFrom := from.Over
					for nextFrom != nil {
						// add document joins
						lastStep = plan.NewDocumentJoin(lastStep, nextFrom.Projection, nextFrom.As)
						nextFrom = nextFrom.Over
					}
					foundUsableScanner = true
					break
				}
			}

			if !foundUsableScanner {
				ec <- query.NewError(nil, fmt.Sprintf("No usable scanner found for bucket %v", from.Bucket))
				return
			}
		} else {
			ec <- query.NewPoolDoesNotExist(this.pool.Name())
			return
		}
	}

	if stmt.GetWhere() != nil {
		lastStep = plan.NewFilter(lastStep, stmt.GetWhere())
	}

	lastStep = plan.NewProjector(lastStep, stmt.GetResultExpressionList(), true)

	if stmt.IsDistinct() {
		lastStep = plan.NewEliminateDuplicates(lastStep)
	}

	if stmt.GetOrderBy() != nil {
		lastStep = plan.NewOrder(lastStep, stmt.GetOrderBy())
	}

	if stmt.GetOffset() != 0 {
		lastStep = plan.NewOffset(lastStep, stmt.GetOffset())
	}

	if stmt.GetLimit() >= 0 {
		lastStep = plan.NewLimit(lastStep, stmt.GetLimit())
	}

	pc <- plan.Plan{Root: lastStep}
}
