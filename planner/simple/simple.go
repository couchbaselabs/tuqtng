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
	"log"

	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/plan"
)

type SimplePlanner struct {
	pool catalog.Pool
}

func NewSimplePlanner(pool catalog.Pool) *SimplePlanner {
	return &SimplePlanner{
		pool: pool,
	}
}

func (this *SimplePlanner) Plan(stmt ast.Statement) plan.PlanChannel {
	rv := make(plan.PlanChannel)
	go this.buildPlans(stmt, rv)
	return rv
}

func (this *SimplePlanner) buildPlans(stmt ast.Statement, pc plan.PlanChannel) {
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
				log.Printf("no bucket named %v", from.Bucket)
				// no bucket, no plan
				close(pc)
				return
			}

			// find all docs scanner
			scanners, err := bucket.Scanners()
			if err != nil {
				// no scanner, no plan
				close(pc)
				return
			}

			foundUsableScanner := false
			for _, scanner := range scanners {
				switch scanner.(type) {
				case catalog.FullScanner:
					lastStep = plan.NewScan(bucket.Name(), "_all_docs") // FIXME need scanner names
					lastStep = plan.NewFetch(lastStep, bucket.Name())
					foundUsableScanner = true
					break
				}
			}

			if !foundUsableScanner {
				// no useable scanner
				close(pc)
				return
			}
		} else {
			// there is no pool
			close(pc)
			return
		}
	}

	if stmt.GetWhere() != nil {
		lastStep = plan.NewFilter(lastStep, stmt.GetWhere())
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

	lastStep = plan.NewProjector(lastStep, stmt.GetResultExpressionList())

	pc <- plan.Plan{Root: lastStep}

	close(pc)
}
