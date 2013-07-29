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
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
)

type SimpleOptimizer struct {
}

func NewSimpleOptimizer() *SimpleOptimizer {
	return &SimpleOptimizer{}
}

// simplest possible implementation
// 1.  read all plans off plan channel
// 2.  return first channel
func (this *SimpleOptimizer) Optimize(planChannel plan.PlanChannel, errChannel query.ErrorChannel) (*plan.Plan, query.Error) {

	plans := make([]plan.Plan, 0)

	var p plan.Plan
	var err query.Error
	ok := true
	for ok {
		select {
		case p, ok = <-planChannel:
			plans = append(plans, p)
		case err, ok = <-errChannel:
			if err != nil {
				return nil, err
			}
		}
	}

	if len(plans) > 0 {
		return &plans[0], nil
	}

	return nil, query.NewError(nil, "No plans produced for optimizer to choose from")
}
