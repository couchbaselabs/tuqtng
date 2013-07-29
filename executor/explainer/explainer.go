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
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/plan"
)

type ExplainerExecutor struct {
}

func NewExplainerExecutor() *ExplainerExecutor {
	return &ExplainerExecutor{}
}

func (this *ExplainerExecutor) Execute(optimalPlan *plan.Plan, query network.Query) {

	query.Response.SendResult(optimalPlan)
	query.Response.NoMoreResults()

	return
}
