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
	"github.com/couchbaselabs/tuqtng/query"
)

type ExpressionEvaluatorSource struct {
	itemChannel query.ItemChannel
}

func NewExpressionEvaluatorSource() *ExpressionEvaluatorSource {
	return &ExpressionEvaluatorSource{
		itemChannel: make(query.ItemChannel),
	}
}

func (this *ExpressionEvaluatorSource) SetSource(source Operator) {
	panic("Cannot set source for a datasource")
}

func (this *ExpressionEvaluatorSource) GetItemChannel() query.ItemChannel {
	return this.itemChannel
}

func (this *ExpressionEvaluatorSource) Run() {
	defer close(this.itemChannel)
	item := query.NewMapItem(map[string]query.Value{}, nil)
	this.itemChannel <- item
}