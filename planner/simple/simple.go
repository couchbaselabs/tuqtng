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
	site        catalog.Site
	defaultPool string
}

func NewSimplePlanner(site catalog.Site, defaultPool string) *SimplePlanner {
	return &SimplePlanner{
		site:        site,
		defaultPool: defaultPool,
	}
}

func (this *SimplePlanner) Plan(stmt ast.Statement) (plan.PlanChannel, query.ErrorChannel) {
	pc := make(plan.PlanChannel)
	ec := make(query.ErrorChannel)
	go this.buildPlans(stmt, pc, ec)
	return pc, ec
}

func (this *SimplePlanner) buildSelectStatementPlans(stmt *ast.SelectStatement, pc plan.PlanChannel, ec query.ErrorChannel) {
	from := stmt.GetFrom()

	var lastStep plan.PlanElement

	if from == nil {
		// simple expression evaluation
		lastStep = plan.NewExpressionEvaluator()

	} else {
		// get the pool
		poolName := from.Pool
		if poolName == "" {
			poolName = this.defaultPool
		}

		pool, err := this.site.PoolByName(poolName)
		if err != nil {
			ec <- query.NewPoolDoesNotExist(poolName)
			return
		}

		bucket, err := pool.BucketByName(from.Bucket)
		if err != nil {
			ec <- query.NewBucketDoesNotExist(from.Bucket)
			return
		}

		// find all docs index
		indexes, err := bucket.Indexes()
		if err != nil {
			ec <- query.NewError(err, fmt.Sprintf("No usable index found for bucket %v", from.Bucket))
			return
		}

		foundUsableIndex := false
		for _, index := range indexes {
			switch index.(type) {
			case catalog.PrimaryIndex:
				lastStep = plan.NewScan(pool.Name(), bucket.Name(), index.Name())
				lastStep = plan.NewFetch(lastStep, pool.Name(), bucket.Name(), from.Projection, from.As)
				nextFrom := from.Over
				for nextFrom != nil {
					// add document joins
					lastStep = plan.NewDocumentJoin(lastStep, nextFrom.Projection, nextFrom.As)
					nextFrom = nextFrom.Over
				}
				foundUsableIndex = true
				break
			}
		}

		if !foundUsableIndex {
			ec <- query.NewError(nil, fmt.Sprintf("No usable index found for bucket %v", from.Bucket))
			return
		}

	}

	if stmt.GetWhere() != nil {
		lastStep = plan.NewFilter(lastStep, stmt.GetWhere())
	}

	if stmt.GetGroupBy() != nil {
		lastStep = plan.NewGroup(lastStep, stmt.GetGroupBy(), stmt.GetAggregateReferences())
	}

	if stmt.GetHaving() != nil {
		lastStep = plan.NewFilter(lastStep, stmt.GetHaving())
	}

	lastStep = plan.NewProjector(lastStep, stmt.GetResultExpressionList(), true)

	if stmt.IsDistinct() {
		lastStep = plan.NewEliminateDuplicates(lastStep)
	}

	if stmt.GetOrderBy() != nil {
		explicitAliases := stmt.GetExplicitProjectionAliases()
		lastStep = plan.NewOrder(lastStep, stmt.GetOrderBy(), explicitAliases)
	}

	if stmt.GetOffset() != 0 {
		lastStep = plan.NewOffset(lastStep, stmt.GetOffset())
	}

	if stmt.GetLimit() >= 0 {
		lastStep = plan.NewLimit(lastStep, stmt.GetLimit())
	}

	if stmt.ExplainOnly {
		lastStep = plan.NewExplain(lastStep)
	}

	pc <- plan.Plan{Root: lastStep}
}

func (this *SimplePlanner) buildCreateIndexStatementPlans(stmt *ast.CreateIndexStatement, pc plan.PlanChannel, ec query.ErrorChannel) {
	ec <- query.NewError(nil, fmt.Sprintf("I don't know how to create indexes yet"))
	return
}

func (this *SimplePlanner) buildPlans(stmt ast.Statement, pc plan.PlanChannel, ec query.ErrorChannel) {
	defer close(pc)
	defer close(ec)
	switch stmt := stmt.(type) {
	case *ast.SelectStatement:
		this.buildSelectStatementPlans(stmt, pc, ec)
	case *ast.CreateIndexStatement:
		this.buildCreateIndexStatementPlans(stmt, pc, ec)
	}

}
