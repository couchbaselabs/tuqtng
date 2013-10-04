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
	"encoding/json"
	"io"
	"net/http"
	"net/http/pprof"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/gorilla/mux"
)

const CHANNEL = "HTTP"

type HttpEndpoint struct {
	queryChannel network.QueryChannel
}

func NewHttpEndpoint(address string, includeProfileHandlers bool, staticPath string) *HttpEndpoint {
	rv := &HttpEndpoint{}

	r := mux.NewRouter()

	r.Handle("/query", rv).Methods("GET", "POST")

	if includeProfileHandlers {
		clog.To(CHANNEL, "Enabling HTTP Profiling Handlers")
		r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticPath)))

	go func() {
		err := http.ListenAndServe(address, r)
		if err != nil {
			clog.Fatal("ListenAndServe: ", err)
		}
	}()

	return rv
}

func (this *HttpEndpoint) SendQueriesTo(queryChannel network.QueryChannel) {
	this.queryChannel = queryChannel
}

func (this *HttpEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	clog.To(CHANNEL, "request received")
	q := NewHttpQuery(w, r)
	if q != nil {
		this.queryChannel <- q
		q.Process()
	}
}

func mustEncode(w io.Writer, i interface{}) {
	if headered, ok := w.(http.ResponseWriter); ok {
		headered.Header().Set("Cache-Control", "no-cache")
		headered.Header().Set("Content-type", "application/json")
	}

	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		panic(err)
	}
}

func showError(w http.ResponseWriter, msg string, code int) {
	clog.To(CHANNEL, "reporting error %v/%v", code, msg)
	http.Error(w, msg, code)
}
