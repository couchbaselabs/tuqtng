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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/server"
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
	go server.Server("dir:.", "json", qc)
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

func TestAllCaseFiles(t *testing.T) {
	matches, err := filepath.Glob("case_*.json")
	if err != nil {
		t.Errorf("glob failed: %v", err)
	}
	for _, m := range matches {
		testCaseFile(t, m)
	}
}

func testCaseFile(t *testing.T, fname string) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
		return
	}
	var cases []map[string]interface{}
	err = json.Unmarshal(b, &cases)
	if err != nil {
		t.Errorf("couldn't json unmarshal: %v, err: %v", string(b), err)
		return
	}
	for i, c := range cases {
		v, ok := c["statements"]
		if !ok || v == nil {
			t.Errorf("missing statements for case file: %v, index: %v", fname, i)
			return
		}
		statements := v.(string)

		qc := start()
		defer close(qc)
		resultsActual, err := run(qc, statements)
		if err != nil {
			v, ok := c["err"]
			if !ok || v == nil {
				t.Errorf("unexpected err: %v, statements: %v", err, statements)
				return
			}
			// TODO: Check that the err matches the expected err.
			continue
		}

		v, ok = c["results"]
		if !ok || v == nil {
			t.Errorf("missing results for case file: %v, index: %v", fname, i)
			return
		}
		results := v.([]interface{})
		if len(resultsActual) != len(results) {
			t.Errorf("results len don't match, %v vs %v, %v vs %v"+
				", for case file: %v, index: %v",
				len(resultsActual), len(results), resultsActual, results, fname, i)
			return
		}
		// TODO: Deep results vs resultsActual comparison.
	}
}
