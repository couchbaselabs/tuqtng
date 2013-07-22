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
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/xpipelinebuilder"
	simpleBuilder "github.com/couchbaselabs/tuqtng/xpipelinebuilder/simple"
)

type SimpleExecutor struct {
	xpipelinebuilder xpipelinebuilder.ExecutablePipelineBuilder
	pool             catalog.Pool
}

func NewSimpleExecutor(pool catalog.Pool) *SimpleExecutor {
	return &SimpleExecutor{
		xpipelinebuilder: simpleBuilder.NewSimpleExecutablePipelineBuilder(pool),
		pool:             pool,
	}
}

func (this *SimpleExecutor) Execute(optimalPlan *plan.Plan, q network.Query) error {

	// first make the plan excutable
	executablePipeline, err := this.xpipelinebuilder.Build(optimalPlan)
	if err != nil {
		return err
	}

	// now execute it
	root := executablePipeline.Root
	itemChannel := root.GetItemChannel()
	go root.Run()
	for item := range itemChannel {
		result := item.GetValue()
		q.Response.SendResult(result)
	}
	q.Response.NoMoreResults()

	return nil
}
