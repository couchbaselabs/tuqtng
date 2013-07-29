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

type StubSource struct {
	data           query.ItemCollection
	itemChannel    query.ItemChannel
	supportChannel PipelineSupportChannel
}

func NewStubSource(data query.ItemCollection) *StubSource {
	return &StubSource{
		data:           data,
		itemChannel:    make(query.ItemChannel),
		supportChannel: make(PipelineSupportChannel),
	}
}

func (this *StubSource) SetSource(Operator) {
	panic("stub source does not have a source")
}

func (this *StubSource) GetChannels() (query.ItemChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *StubSource) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	for _, item := range this.data {
		this.itemChannel <- item
	}
}
