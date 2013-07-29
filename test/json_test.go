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
	"reflect"
	"testing"

	"github.com/couchbaselabs/tuqtng/network"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/server"
)

type MockResponse struct {
	err      query.Error
	results  []query.Value
	warnings []query.Error
	done     chan bool
}

func (this *MockResponse) SendError(err query.Error) {
	this.err = err
	if err.IsFatal() {
		close(this.done)
	}
}

func (this *MockResponse) SendResult(val query.Value) {
	this.results = append(this.results, val)
}

func (this *MockResponse) NoMoreResults() {
	close(this.done)
}

func run(qc network.QueryChannel, q string) ([]query.Value, []query.Error, query.Error) {
	mr := &MockResponse{
		results: []query.Value{}, warnings: []query.Error{}, done: make(chan bool),
	}
	query := network.Query{
		Request:  network.UNQLStringQueryRequest{QueryString: q},
		Response: mr,
	}
	qc <- query
	<-mr.done
	return mr.results, mr.warnings, mr.err
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

	r, _, err := run(qc, "this is a bad query")
	if err == nil || len(r) != 0 {
		t.Errorf("expected err")
	}
	r, _, err = run(qc, "") // empty string query
	if err == nil || len(r) != 0 {
		t.Errorf("expected err")
	}
}

func TestSimpleSelect(t *testing.T) {
	qc := start()
	defer close(qc)

	r, _, err := run(qc, "select * from orders")
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
	matches, err := filepath.Glob("json/cases/case_*.json")
	if err != nil {
		t.Errorf("glob failed: %v", err)
	}
	for _, m := range matches {
		testCaseFile(t, m)
	}
}

func testCaseFile(t *testing.T, fname string) {
	t.Logf("testCaseFile: %v\n", fname)
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
		t.Logf("  %d: %v\n", i, statements)

		qc := start()
		defer close(qc)
		resultsActual, _, errActual := run(qc, statements)

		errExpected := ""
		v, ok = c["error"]
		if ok {
			errExpected = v.(string)
		}
		if errActual != nil {
			if errExpected == "" {
				t.Errorf("unexpected err: %v, statements: %v"+
					", for case file: %v, index: %v", errActual, statements, fname, i)
				return
			}
			// TODO: Check that the actual err matches the expected err.
			continue
		}
		if errExpected != "" {
			t.Errorf("did not see the expected err: %v, statements: %v"+
				", for case file: %v, index: %v", errActual, statements, fname, i)
			return
		}

		v, ok = c["results"]
		if !ok || v == nil {
			t.Errorf("missing results for case file: %v, index: %v", fname, i)
			return
		}
		resultsExpected := v.([]interface{})
		if len(resultsActual) != len(resultsExpected) {
			t.Errorf("results len don't match, %v vs %v, %v vs %v"+
				", for case file: %v, index: %v",
				len(resultsActual), len(resultsExpected),
				resultsActual, resultsExpected, fname, i)
			return
		}
		// Extra marshal/unmarshal hop is to get from query.Value to
		// interface{} so that reflect.DeepEqual() works.
		j, err := json.Marshal(resultsActual)
		if err != nil {
			t.Errorf("json marshal failed: %v", err)
			return
		}
		var ra interface{}
		err = json.Unmarshal(j, &ra)
		if err != nil {
			t.Errorf("json unmarshal failed: %v", err)
			return
		}
		if !reflect.DeepEqual(ra, resultsExpected) {
			t.Errorf("results don't match, actual: %#v, expected: %#v"+
				", for case file: %v, index: %v",
				ra, resultsExpected, fname, i)
		}
	}
}
