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
	"github.com/couchbaselabs/dparval"
)

type StubSource struct {
	data           dparval.ValueCollection
	itemChannel    dparval.ValueChannel
	supportChannel PipelineSupportChannel
}

func NewStubSource(data dparval.ValueCollection) *StubSource {
	return &StubSource{
		data:           data,
		itemChannel:    make(dparval.ValueChannel),
		supportChannel: make(PipelineSupportChannel),
	}
}

func (this *StubSource) SetSource(Operator) {}

func (this *StubSource) GetChannels() (dparval.ValueChannel, PipelineSupportChannel) {
	return this.itemChannel, this.supportChannel
}

func (this *StubSource) Run() {
	defer close(this.itemChannel)
	defer close(this.supportChannel)

	for _, item := range this.data {
		this.itemChannel <- item
	}
}

func (this *StubSource) processItem(item *dparval.Value) bool {
	return true
}

func (this *StubSource) afterItems() {}
