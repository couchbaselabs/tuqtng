//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package test

import (
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/server"
)

type MockQuery struct {
	request     network.QueryRequest
	response    *MockResponse
	stopChannel misc.StopChannel
}

func (this *MockQuery) Request() network.QueryRequest {
	return this.request
}

func (this *MockQuery) Response() network.QueryResponse {
	return this.response
}

func (this *MockQuery) SetStopChannel(stopChannel misc.StopChannel) {
	this.stopChannel = stopChannel
}

type MockResponse struct {
	err      query.Error
	results  []interface{}
	warnings []query.Error
	done     chan bool
}

func (this *MockResponse) SendError(err query.Error) {
	this.err = err
	if err.IsFatal() {
		close(this.done)
	}
}

func (this *MockResponse) SendResult(val interface{}) {
	// sanitize the value the same way the HTTP endpoint does
	sanitizedValue := misc.SanitizeUnrepresentableJSON(val)
	this.results = append(this.results, sanitizedValue)
}

func (this *MockResponse) NoMoreResults() {
	close(this.done)
}

func Run(qc network.QueryChannel, q string) ([]interface{}, []query.Error, query.Error) {
	mr := &MockResponse{
		results: []interface{}{}, warnings: []query.Error{}, done: make(chan bool),
	}

	query := MockQuery{
		request:  network.StringQueryRequest{QueryString: q},
		response: mr,
	}
	qc <- &query
	<-mr.done
	return mr.results, mr.warnings, mr.err
}

func Start(site, pool string) network.QueryChannel {
	qc := make(network.QueryChannel)
	go server.Server("TEST", site, pool, qc)
	return qc
}
