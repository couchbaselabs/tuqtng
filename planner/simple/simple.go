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
	"github.com/couchbaselabs/tuqtng/plan"
)

type SimplePlanner struct {
}

func NewSimplePlanner() *SimplePlanner {
	return &SimplePlanner{}
}

func (this *SimplePlanner) Plan(ast.Statement) plan.PlanChannel {
	rv := make(plan.PlanChannel)
	// FIXME not implemented
	close(rv)
	return rv
}
