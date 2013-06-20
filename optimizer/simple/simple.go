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
	"fmt"

	"github.com/couchbaselabs/tuqtng/plan"
)

type SimpleOptimizer struct {
}

func NewSimpleOptimizer() *SimpleOptimizer {
	return &SimpleOptimizer{}
}

// simplest possible implementation
// 1.  read all plans off plan channel
// 2.  return first channel
func (this *SimpleOptimizer) Optimize(planChannel plan.PlanChannel) (plan.Plan, error) {
	plans := make([]plan.Plan, 0)
	for plan := range planChannel {
		plans = append(plans, plan)
	}

	if len(plans) > 0 {
		return plans[0], nil
	}

	return nil, fmt.Errorf("No Plans to Choose From")
}
