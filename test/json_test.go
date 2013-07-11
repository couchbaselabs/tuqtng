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
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/couchbaselabs/tuqtng"
	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
)

type MockResponse struct {
	err     error
	results []query.Value
	done    chan bool
}

func (this *MockResponse) SendError(err error) {
	this.err = err
	close(this.done)
}

func (this *MockResponse) SendResult(val query.Value) {
	this.results = append(this.results, val)
}

func (this *MockResponse) NoMoreResults() {
	close(this.done)
}

func run(qc network.QueryChannel, q string) ([]query.Value, error) {
	mr := &MockResponse{results: []query.Value{}, done: make(chan bool)}
	query := network.Query{
		Request:  network.UNQLStringQueryRequest{QueryString: q},
		Response: mr,
	}
	qc <- query
	<-mr.done
	return mr.results, mr.err
}

func start() network.QueryChannel {
	qc := make(network.QueryChannel)
	go main.Main("dir:.", "json", qc)
	return qc
}

func TestMainClose(t *testing.T) {
	qc := start()
	close(qc)
}

func TestSyntaxErr(t *testing.T) {
	qc := start()
	defer close(qc)

	r, err := run(qc, "this is a bad query")
	if err == nil || len(r) != 0 {
		t.Errorf("expected err")
	}
	r, err = run(qc, "") // empty string query
	if err == nil || len(r) != 0 {
		t.Errorf("expected err")
	}
}

func TestSimpleSelect(t *testing.T) {
	qc := start()
	defer close(qc)

	r, err := run(qc, "select * from orders")
	if err != nil || len(r) == 0 {
		t.Errorf("did not expect err")
	}

	fileInfos, _ := ioutil.ReadDir("./json/orders")
	if len(r) != len(fileInfos) {
		fmt.Printf("r: %#v, fileInfos: %#v\n", r, fileInfos)
		t.Errorf("expected # of results to match directory listing")
	}
}
