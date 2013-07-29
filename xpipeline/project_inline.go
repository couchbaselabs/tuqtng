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

type ProjectInline struct {
	Source      Operator
	itemChannel query.ItemChannel
	errChannel  query.ErrorChannel
	warnChannel query.ErrorChannel
	Result      *ast.ResultExpression
}

func NewProjectInline(result *ast.ResultExpression) *ProjectInline {
	return &ProjectInline{
		Result:      result,
		itemChannel: make(query.ItemChannel),
		errChannel:  make(query.ErrorChannel),
		warnChannel: make(query.ErrorChannel),
	}
}

func (this *ProjectInline) SetSource(source Operator) {
	this.Source = source
}

func (this *ProjectInline) GetChannels() (query.ItemChannel, query.ErrorChannel, query.ErrorChannel) {
	return this.itemChannel, this.warnChannel, this.errChannel
}

func (this *ProjectInline) Run() {
	defer close(this.itemChannel)
	defer close(this.errChannel)
	defer close(this.warnChannel)

	go this.Source.Run()

	var item query.Item
	var warn query.Error
	var err query.Error
	sourceItemChannel, sourceWarnChannel, sourceErrorChannel := this.Source.GetChannels()
	ok := true
	for ok {
		select {
		case item, ok = <-sourceItemChannel:
			if ok {
				this.processItem(item)
			}
		case warn, ok = <-sourceWarnChannel:
			// propogate the warning
			if warn != nil {
				this.warnChannel <- warn
			}
		case err, ok = <-sourceErrorChannel:
			// propogate the error and return
			if err != nil {
				this.errChannel <- err
				return
			}
		}
	}
}

func (this *ProjectInline) processItem(item query.Item) {
	var res query.Value

	if this.Result.Star {
		if this.Result.Expr != nil {
			// evaluate this expression first
			val, err := this.Result.Expr.Evaluate(item)
			if err == nil {
				switch val := val.(type) {
				case map[string]query.Value:
					// then if the result was an object
					// then project it
					res = val
				}
			} else {
				switch err := err.(type) {
				case *query.Undefined:
					// undefined contributes nothing to the result
					// but otherwise is NOT an error
					// FIXME review if this should be a warning
					return
				default:
					log.Fatalf("unexpected error projecting dot star expression: %v", err)
				}
			}
		} else {
			// just a star, make item the result
			res = item.GetValue()
		}
	} else if this.Result.Expr != nil {
		// evaluate the expression
		val, err := this.Result.Expr.Evaluate(item)
		if err == nil {
			res = val
		} else {
			switch err := err.(type) {
			case *query.Undefined:
				// undefined contributes nothing to the result map
				// but otherwise is NOT an error
				// FIXME review if this should be a warning
				return
			default:
				log.Fatalf("unexpected error projecting expression: %v", err)
			}
		}
	}

	// create the actual result Item
	finalItem := query.NewParsedItem(res, item.GetMeta())

	// write this to the output
	this.itemChannel <- finalItem
}
