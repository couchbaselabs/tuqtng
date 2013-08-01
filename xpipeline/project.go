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
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/mschoch/dparval"
)

type Project struct {
	Source         Operator
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
	Result         ast.ResultExpressionList
	projectEmpty   bool
	ok             bool
}

func NewProject(result ast.ResultExpressionList, projectEmpty bool) *Project {
	return &Project{
		Result:         result,
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
		projectEmpty:   projectEmpty,
	}
}

func (this *Project) SetSource(source Operator) {
	this.Source = source
}

func (this *Project) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *Project) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	go this.Source.Run()

	var item *dparval.Value
	var obj interface{}
	sourceItemChannel, supportChannel := this.Source.GetChannels()
	this.ok = true
	for this.ok {
		select {
		case item, this.ok = <-sourceItemChannel:
			if this.ok {
				this.processItem(item)
			}
		case obj, this.ok = <-supportChannel:
			if this.ok {
				switch obj := obj.(type) {
				case query.Error:
					this.supportChannel <- obj
					if obj.IsFatal() {
						return
					}
				default:
					this.supportChannel <- obj
				}
			}
		}
	}
}

func (this *Project) processItem(item *dparval.Value) {
	resultMap := map[string]interface{}{}
	for _, resultItem := range this.Result {
		if resultItem.Star {
			if resultItem.Expr != nil {
				// evaluate this expression first
				val, err := resultItem.Expr.Evaluate(item)
				if err != nil {
					switch err := err.(type) {
					case *dparval.Undefined:
						// undefined contributes nothing to the result map
						// but otherwise is NOT an error
						// FIXME review if this should be a warning
						continue
					default:
						this.supportChannel <- query.NewError(err, "unexpected error projecting dot star expression")
						this.ok = false
						return
					}
				}

				if val.Type() == dparval.OBJECT {
					valval := val.Value()
					switch valval := valval.(type) {
					case map[string]interface{}:
						// then if the result was an object
						// add its contents ot the result map
						for k, v := range valval {
							resultMap[k] = v
						}
					}
				}

			} else {
				if item.Type() == dparval.OBJECT {
					// just a star, get the value, if its a map project the key/value pairs
					val := item.Value()
					switch val := val.(type) {
					case map[string]interface{}:
						for k, v := range val {
							resultMap[k] = v
						}
					}
				}
			}
		} else if resultItem.Expr != nil {
			// evaluate the expression
			val, err := resultItem.Expr.Evaluate(item)
			if err != nil {
				switch err := err.(type) {
				case *dparval.Undefined:
					// undefined contributes nothing to the result map
					// but otherwise is NOT an error
					// FIXME review if this should be a warning
					continue
				default:
					this.supportChannel <- query.NewError(err, "unexpected error projecting expression")
					this.ok = false
					return
				}
			}
			resultMap[resultItem.As] = val

		}
	}

	if !this.projectEmpty && len(resultMap) == 0 {
		return
	}

	// create the actual result Item
	finalItem := dparval.NewValue(resultMap)
	itemMetaVal := item.Meta()
	if itemMetaVal != nil {
		itemMetaData, err := itemMetaVal.Path("meta")
		if err != nil {
			this.supportChannel <- query.NewError(err, "unable to find item metadata")
			this.ok = false
			return
		}
		finalItem.AddMeta("meta", itemMetaData)
	}

	// write this to the output
	this.itemChannel <- finalItem
}
