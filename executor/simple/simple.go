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
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/executor"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/xpipelinebuilder"
	simpleBuilder "github.com/couchbaselabs/tuqtng/xpipelinebuilder/simple"
)

type SimpleExecutor struct {
	xpipelinebuilder xpipelinebuilder.ExecutablePipelineBuilder
	site             catalog.Site
}

func NewSimpleExecutor(site catalog.Site, defaultPool string) *SimpleExecutor {
	return &SimpleExecutor{
		xpipelinebuilder: simpleBuilder.NewSimpleExecutablePipelineBuilder(site, defaultPool),
		site:             site,
	}
}

func (this *SimpleExecutor) Execute(optimalPlan *plan.Plan, q network.Query) {

	clog.To(executor.CHANNEL, "simple executor started")

	// first make the plan excutable
	executablePipeline, berr := this.xpipelinebuilder.Build(optimalPlan)
	if berr != nil {
		q.Response().SendError(query.NewError(berr, ""))
	}

	root := executablePipeline.Root

	// create a stop channel
	stopChannel := make(misc.StopChannel)
	// set it on the query object, so HTTP layer can
	// stop us if the client goes away
	q.SetStopChannel(stopChannel)
	go root.Run(stopChannel)

	// now execute it
	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := root.GetChannels()
	ok := true
	for ok {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				ok = this.processItem(q, item)
				clog.To(executor.CHANNEL, "simple executor sent client item: %v", item)
			}
		case obj, ok = <-supportChannel:
			if ok {
				switch obj := obj.(type) {
				case query.Error:
					q.Response().SendError(obj)
					clog.To(executor.CHANNEL, "simple executor sent client error: %v", obj)
					if obj.IsFatal() {
						return
					}
				}
			}
		}
	}

	q.Response().NoMoreResults()
	clog.To(executor.CHANNEL, "simple executor finished")
}

func (this *SimpleExecutor) processItem(q network.Query, item *dparval.Value) bool {
	projection := item.GetAttachment("projection")
	switch projection := projection.(type) {
	case *dparval.Value:
		result := projection.Value()
		q.Response().SendResult(result)
	default:
		q.Response().SendResult(projection)
	}
	return true
}
