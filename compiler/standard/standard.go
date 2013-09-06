//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package standard

import (

	// interfaces
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/optimizer"
	"github.com/couchbaselabs/tuqtng/parser"
	"github.com/couchbaselabs/tuqtng/plan"
	"github.com/couchbaselabs/tuqtng/planner"
	"github.com/couchbaselabs/tuqtng/query"

	// implementations
	simpleOptimizer "github.com/couchbaselabs/tuqtng/optimizer/simple"
	yaccParser "github.com/couchbaselabs/tuqtng/parser/goyacc"
	simplePlanner "github.com/couchbaselabs/tuqtng/planner/simple"
)

type StandardCompiler struct {
	site            catalog.Site
	defaultPoolName string
	parser          parser.Parser
	planner         planner.Planner
	optimizer       optimizer.Optimizer
}

func NewCompiler(site catalog.Site, defaultPoolName string) *StandardCompiler {
	return &StandardCompiler{
		site:            site,
		defaultPoolName: defaultPoolName,
		parser:          yaccParser.NewN1qlParser(),
		planner:         simplePlanner.NewSimplePlanner(site, defaultPoolName),
		optimizer:       simpleOptimizer.NewSimpleOptimizer(),
	}
}

func (this *StandardCompiler) Compile(queryString string) (*plan.Plan, query.Error) {

	ast, err := this.parser.Parse(queryString)
	if err != nil {
		return nil, query.NewParseError(err, "Parse Error")
	}

	// perform semantic verification
	err = ast.VerifySemantics()
	if err != nil {
		return nil, query.NewSemanticError(err, "Semantic Error")
	}

	// simplify the statement
	err = ast.Simplify()
	if err != nil {
		return nil, query.NewError(err, "Error Simplifying Expression")
	}

	planChannel, planErrChannel := this.planner.Plan(ast)

	optimalPlan, err := this.optimizer.Optimize(planChannel, planErrChannel)
	if err != nil {
		return nil, query.NewError(err, "Optimizer Error")
	}

	return optimalPlan, nil
}
