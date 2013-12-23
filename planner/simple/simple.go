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
	"strings"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/catalog/system"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/planner"
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

func validateKeys(keylist string) []string {
	// remove the brackets and then split the string
	keylist = strings.Trim(keylist, "[")
	keylist = strings.Trim(keylist, "]")
	keys := strings.Split(keylist, ",")
	for i, key := range keys {
		key = strings.TrimSpace(key)
		keys[i] = strings.Trim(key, "\"")
	}
	return keys
}

func (this *SimplePlanner) buildSelectStatementPlans(stmt *ast.SelectStatement, pc plan.PlanChannel, ec query.ErrorChannel) {

	var planHeads []plan.PlanElement

	from := stmt.GetFrom()
	if from == nil {
		// point to :system.dual
		from = &ast.From{Pool: system.POOL_NAME, Bucket: system.BUCKET_NAME_DUAL}
	}

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
		ec <- query.NewError(err, fmt.Sprintf("No indexes found for bucket %v", from.Bucket))
		return
	}

	keylist := stmt.Keys.GetKeys()

	clog.To(planner.CHANNEL, "Indexes in bucket %v", indexes)

	if keylist == nil {
		for _, index := range indexes {
			var lastStep plan.PlanElement
			switch index := index.(type) {
			case catalog.PrimaryIndex:
				clog.To(planner.CHANNEL, "See primary index %v", index.Name())
				// if from.Over == nil && stmt.Where == nil && stmt.GroupBy != nil && len(stmt.GroupBy) == 0 && CanFastCountBucket(stmt.Select) {
				// 	lastStep = plan.NewFastCount(pool.Name(), bucket.Name(), "", nil, nil)
				// } else {
				lastStep = plan.NewScan(pool.Name(), bucket.Name(), index.Name(), nil)
				// }
			case catalog.RangeIndex:
				// see if this index can be used
				clog.To(planner.CHANNEL, "See index %v", index.Name())
				clog.To(planner.CHANNEL, "with Key %v", index.Key())
				if stmt.Where != nil && from.Projection == nil {
					possible, ranges, _, err := CanIUseThisIndexForThisWhereClause(index, stmt.Where, stmt.From.As)
					if err != nil {
						clog.Error(err)
						continue
					}
					clog.To(planner.CHANNEL, "Can I use it1: %v", possible)
					if possible {

						// great, but lets check for a min optimizatin too
						if stmt.GroupBy != nil && len(stmt.GroupBy) == 0 {
							possible, minranges, _, _ := CanIUseThisIndexForThisProjectionNoWhereNoGroupClause(index, stmt.Select, stmt.From.As)
							if possible {
								for _, r := range ranges {
									r.Limit = minranges[0].Limit
								}
							}
						}

						scan := plan.NewScan(pool.Name(), bucket.Name(), index.Name(), ranges)
						// see if this index covers the query
						if DoesIndexCoverStatement(index, stmt) {
							scan.Cover = true
							scan.As = from.As
						}
						lastStep = scan
					} else {
						continue
					}
				} else if from.Projection == nil {

					// try to do a fast count if its possible
					doingFastCount := false
					// countIndex, isCountIndex := index.(catalog.CountIndex)
					// if isCountIndex {
					// 	fastCountIndexOnExpr := CanFastCountIndex(countIndex, stmt.From.As, stmt.Select)
					// 	if fastCountIndexOnExpr != nil && from.Over == nil && stmt.Where == nil && stmt.GroupBy != nil && len(stmt.GroupBy) == 0 {
					// 		lastStep = plan.NewFastCount(pool.Name(), bucket.Name(), countIndex.Name(), fastCountIndexOnExpr, nil)
					// 		doingFastCount = true
					// 	}

					// }

					// this works for aggregates on the whole bucket
					if !doingFastCount && stmt.GroupBy != nil && len(stmt.GroupBy) == 0 {
						possible, ranges, _, err := CanIUseThisIndexForThisProjectionNoWhereNoGroupClause(index, stmt.Select, stmt.From.As)
						if err != nil {
							clog.Error(err)
							continue
						}
						clog.To(planner.CHANNEL, "Can I use it2: %v", possible)
						if possible {
							lastStep = plan.NewScan(pool.Name(), bucket.Name(), index.Name(), ranges)
						} else {
							continue
						}
					} else if !doingFastCount {
						continue
					}
				}

			default:
				clog.To(planner.CHANNEL, "Unsupported type of index %T", index)
				continue
			}
			scanOp, lastStepWasScan := lastStep.(*plan.Scan)
			if lastStepWasScan {
				if !scanOp.Cover {
					lastStep = plan.NewFetch(lastStep, pool.Name(), bucket.Name(), from.Projection, from.As)
					nextFrom := from.Over
					for nextFrom != nil {
						// add document joins
						lastStep = plan.NewDocumentJoin(lastStep, nextFrom.Projection, nextFrom.As)
						nextFrom = nextFrom.Over
					}
				}
			}
			planHeads = append(planHeads, lastStep)
		}
	} else if keylist != nil {
		var lastStep plan.PlanElement
		lastStep = plan.NewKeyScan(keylist)
		lastStep = plan.NewFetch(lastStep, pool.Name(), bucket.Name(), from.Projection, from.As)
		planHeads = append(planHeads, lastStep)
	}

	if len(planHeads) == 0 {
		ec <- query.NewError(nil, fmt.Sprintf("No usable indexes found for bucket %v", from.Bucket))
		return
	}

	// now for all the plan heads, create a full plan
	for _, lastStep := range planHeads {

		if stmt.GetWhere() != nil {
			ids := WhereClauseFindById(stmt.GetWhere())
			fetch, ok := lastStep.(*plan.Fetch)
			if ids != nil && ok {
				fetch.ConvertToIds(ids)
			} else {
				lastStep = plan.NewFilter(lastStep, stmt.GetWhere())
			}
		}

		if stmt.GetGroupBy() != nil {
			_, isFastCount := lastStep.(*plan.FastCount)
			if !isFastCount {
				lastStep = plan.NewGroup(lastStep, stmt.GetGroupBy(), stmt.GetAggregateReferences())
			}
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

}

func (this *SimplePlanner) buildCreateIndexStatementPlans(stmt *ast.CreateIndexStatement, pc plan.PlanChannel, ec query.ErrorChannel) {

	poolName := stmt.Pool
	if poolName == "" {
		poolName = this.defaultPool
	}

	pool, err := this.site.PoolByName(poolName)
	if err != nil {
		ec <- query.NewPoolDoesNotExist(this.defaultPool)
		return
	}

	bucket, err := pool.BucketByName(stmt.Bucket)
	if err != nil {
		ec <- query.NewBucketDoesNotExist(stmt.Bucket)
		return
	}

	var lastStep plan.PlanElement

	lastStep = plan.NewCreateIndex(pool.Name(), bucket.Name(), stmt.Name, stmt.Method, stmt.Primary, stmt.On)
	if stmt.ExplainOnly {
		lastStep = plan.NewExplain(lastStep)
	}

	pc <- plan.Plan{Root: lastStep}
	return
}

func (this *SimplePlanner) buildDropIndexStatementPlans(stmt *ast.DropIndexStatement, pc plan.PlanChannel, ec query.ErrorChannel) {

	poolName := stmt.Pool
	if poolName == "" {
		poolName = this.defaultPool
	}

	pool, err := this.site.PoolByName(poolName)
	if err != nil {
		ec <- query.NewPoolDoesNotExist(this.defaultPool)
		return
	}

	bucket, err := pool.BucketByName(stmt.Bucket)
	if err != nil {
		ec <- query.NewBucketDoesNotExist(stmt.Bucket)
		return
	}

	var lastStep plan.PlanElement

	lastStep = plan.NewDropIndex(pool.Name(), bucket.Name(), stmt.Name)
	if stmt.ExplainOnly {
		lastStep = plan.NewExplain(lastStep)
	}

	pc <- plan.Plan{Root: lastStep}
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
	case *ast.DropIndexStatement:
		this.buildDropIndexStatementPlans(stmt, pc, ec)
	}

}
