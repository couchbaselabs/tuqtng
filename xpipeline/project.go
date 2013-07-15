//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"log"

	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

type Project struct {
	Source      Operator
	itemChannel query.ItemChannel
	Result      ast.ResultExpressionList
}

func NewProject(result ast.ResultExpressionList) *Project {
	return &Project{
		Result:      result,
		itemChannel: make(query.ItemChannel),
	}
}

func (this *Project) SetSource(source Operator) {
	this.Source = source
}

func (this *Project) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *Project) Run() {
	defer close(this.itemChannel)

	// start the source
	go this.Source.Run()
	for item := range this.Source.GetItemChannel() {

		resultMap := map[string]query.Value{}
		for _, resultItem := range this.Result {
			if resultItem.Star {
				if resultItem.Expr != nil {
					// evaluate this expression first
					val, err := resultItem.Expr.Evaluate(item)
					if err == nil {
						switch val := val.(type) {
						case map[string]query.Value:
							// then if the result was an object
							// add its contents ot the result map
							for k, v := range val {
								resultMap[k] = v
							}
						}
					} else {
						switch err := err.(type) {
						case *query.Undefined:
							// undefined contributes nothing to the result map
							// but otherwise is NOT an error
							// FIXME review if this should be a warning
							continue
						default:
							log.Fatalf("unexpected error projecting dot star expression: %v", err)
						}
					}
				} else {
					// just a star, take all the contents of the source item
					// and add them to the result item
					topLevelKeys := item.GetTopLevelKeys()
					for _, key := range topLevelKeys {
						val, err := item.GetPath(key)
						if err == nil {
							resultMap[key] = val
						} else {
							log.Fatalf("unexpected error projecting star expression: %v", err)
						}
					}
				}
			} else if resultItem.Expr != nil {
				// evaluate the expression
				val, err := resultItem.Expr.Evaluate(item)
				if err == nil {
					resultMap[resultItem.As] = val
				} else {
					switch err := err.(type) {
					case *query.Undefined:
						// undefined contributes nothing to the result map
						// but otherwise is NOT an error
						// FIXME review if this should be a warning
						continue
					default:
						log.Fatalf("unexpected error projecting expression: %v", err)
					}
				}
			}
		}

		// create the actual result Item
		finalItem := query.NewMapItem(resultMap, item.GetMeta())

		// write this to the output
		this.itemChannel <- finalItem
	}
}
