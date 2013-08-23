//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package static

import (
	"fmt"

	"github.com/couchbaselabs/clog"

	// interfaces
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/executor"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/optimizer"
	"github.com/couchbaselabs/tuqtng/parser"
	"github.com/couchbaselabs/tuqtng/planner"
	"github.com/couchbaselabs/tuqtng/query"

	// implementations
	interpretedExecutor "github.com/couchbaselabs/tuqtng/executor/interpreted"
	simpleOptimizer "github.com/couchbaselabs/tuqtng/optimizer/simple"
	yaccParser "github.com/couchbaselabs/tuqtng/parser/goyacc"
	simplePlanner "github.com/couchbaselabs/tuqtng/planner/simple"
)

type StaticPipeline struct {
	site            catalog.Site
	defaultPoolName string
	parser          parser.Parser
	planner         planner.Planner
	optimizer       optimizer.Optimizer
	executor        executor.Executor
}

func NewStaticPipeline(site catalog.Site, defaultPoolName string) *StaticPipeline {
	return &StaticPipeline{
		site:            site,
		defaultPoolName: defaultPoolName,
		parser:          yaccParser.NewUnqlParser(),
		planner:         simplePlanner.NewSimplePlanner(site, defaultPoolName),
		optimizer:       simpleOptimizer.NewSimpleOptimizer(),
		executor:        interpretedExecutor.NewExecutor(site, defaultPoolName),
	}
}

// this construct allows to support alternate types of queries
// that enter the pipeline at different places
// ie, an UNQLASTQueryRequest could bypass the parser
// or an UNQLPlanRequest could bypass parser, planner, and optimizer
func (this *StaticPipeline) DispatchQuery(q network.Query) {
	request := q.Request()
	response := q.Response()

	switch request := request.(type) {
	case network.StringQueryRequest:
		ast, err := this.parser.Parse(request.QueryString)
		if err != nil {
			response.SendError(query.NewParseError(err, "Parse Error"))
			return
		}

		// perform semantic verification
		err = ast.VerifySemantics()
		if err != nil {
			response.SendError(query.NewSemanticError(err, "Semantic Error"))
			return
		}

		planChannel, planErrChannel := this.planner.Plan(ast)

		optimalPlan, err := this.optimizer.Optimize(planChannel, planErrChannel)
		if err != nil {
			response.SendError(query.NewError(err, ""))
			return
		}

		this.executor.Execute(optimalPlan, q)

	default:
		clog.Error(fmt.Errorf("Unsupported Request Type %T", request))
	}
}
