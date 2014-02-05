//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package http

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/misc"
	"github.com/couchbaselabs/tuqtng/network"
)

type HttpQuery struct {
	request   network.QueryRequest
	response  *HttpResponse
	stop      misc.StopChannel
	startTime time.Time
	info      bool
}

func NewHttpQuery(w http.ResponseWriter, r *http.Request, info bool) *HttpQuery {
	q := HttpQuery{startTime: time.Now(), info: info}

	queryString := findQueryStringInRequest(r)

	if queryString == "" {
		showError(w, "Missing required query string", 500)
		return nil
	} else {
		clog.To(CHANNEL, "query string: %v", queryString)
	}

	q.request = network.StringQueryRequest{QueryString: queryString}
	httpResponse := &HttpResponse{query: &q, w: w, results: make(chan interface{}), returnInfo: info}
	q.response = httpResponse

	return &q
}

func (this *HttpQuery) StartTime() time.Time {
	return this.startTime
}

func (this *HttpQuery) Duration() time.Duration {
	return time.Since(this.startTime)
}

func (this *HttpQuery) StopProcessing() {
	if this.stop != nil {
		close(this.stop)
		this.stop = nil
	}
}

func (this *HttpQuery) Request() network.QueryRequest {
	return this.request
}

func (this *HttpQuery) Response() network.QueryResponse {
	return this.response
}

func (this *HttpQuery) SetStopChannel(stop misc.StopChannel) {
	this.stop = stop
}

func (this *HttpQuery) Process() {
	err := this.response.Process()
	if err != nil {
		clog.To(CHANNEL, "error writing to client, aborting query")
		this.StopProcessing()
	} else {
		clog.To(CHANNEL, "response complete")
	}
}

func findQueryStringInRequest(r *http.Request) string {
	queryString := r.FormValue("q")
	if queryString == "" && r.Method == "POST" {
		queryStringBytes, err := ioutil.ReadAll(r.Body)
		if err == nil {
			queryString = string(queryStringBytes)
		}
	}
	return queryString
}
