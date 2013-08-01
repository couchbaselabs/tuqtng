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
	"log"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/xpipelinebuilder"
	simpleBuilder "github.com/couchbaselabs/tuqtng/xpipelinebuilder/simple"
	"github.com/mschoch/dparval"
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

func (this *SimpleExecutor) Execute(optimalPlan *plan.Plan, q network.Query) {

	// first make the plan excutable
	executablePipeline, berr := this.xpipelinebuilder.Build(optimalPlan)
	if berr != nil {
		q.Response.SendError(query.NewError(berr, ""))
	}

	root := executablePipeline.Root
	go root.Run()

	// now execute it
	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := root.GetChannels()
	ok := true
	for ok {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				this.processItem(q, item)
			}
		case obj, ok = <-supportChannel:
			if ok {
				switch obj := obj.(type) {
				case query.Error:
					log.Printf("Sending client error: %v", obj)
					q.Response.SendError(obj)
					if obj.IsFatal() {
						return
					}
				default:
					log.Printf("Unexpected object tyep on the support channel %T", obj)
				}
			}
		}
	}

	q.Response.NoMoreResults()
}

func (this *SimpleExecutor) processItem(q network.Query, item *dparval.Value) {
	result := item.Value()
	q.Response.SendResult(result)
}
