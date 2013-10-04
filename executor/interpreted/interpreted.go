//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package interpreted

import (
	"time"

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

type InterpretedExecutor struct {
	xpipelinebuilder xpipelinebuilder.ExecutablePipelineBuilder
	site             catalog.Site
}

func NewExecutor(site catalog.Site, defaultPool string) *InterpretedExecutor {
	return &InterpretedExecutor{
		xpipelinebuilder: simpleBuilder.NewSimpleExecutablePipelineBuilder(site, defaultPool),
		site:             site,
	}
}

func (this *InterpretedExecutor) Execute(optimalPlan *plan.Plan, q network.Query, timeout *time.Duration) {
	stopChannel := make(misc.StopChannel)
	if timeout.Nanoseconds() < 0 {
		this.executeInternal(optimalPlan, q, stopChannel)
	} else {
		c := make(chan error, 1)
		go func() {
			this.executeInternal(optimalPlan, q, stopChannel)
			c <- nil
		}()
		select {
		case <-c:
			return
		case <-time.After(*timeout):
			clog.To(executor.CHANNEL, "simple executor timeout trigger")
			close(stopChannel)
			clog.To(executor.CHANNEL, "stop channel closed")
		}
		<-c
		q.Response().SendError(query.NewTimeoutError(timeout))
	}
}

func (this *InterpretedExecutor) executeInternal(optimalPlan *plan.Plan, q network.Query, timeoutStopChannel misc.StopChannel) {

	clog.To(executor.CHANNEL, "simple executor started")

	// first make the plan excutable
	executablePipeline, berr := this.xpipelinebuilder.Build(optimalPlan)
	if berr != nil {
		q.Response().SendError(query.NewError(berr, ""))
		return
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
		case _, ok = <-timeoutStopChannel:
			clog.To(executor.CHANNEL, "simple execution aborted, timeout")
			return
		}
	}

	q.Response().NoMoreResults()
	clog.To(executor.CHANNEL, "simple executor finished")
}

func (this *InterpretedExecutor) processItem(q network.Query, item *dparval.Value) bool {
	projection := item.GetAttachment("projection")
	switch projection := projection.(type) {
	case *dparval.Value:
		result := projection.Value()
		q.Response().SendResult(result)
	}
	return true
}
