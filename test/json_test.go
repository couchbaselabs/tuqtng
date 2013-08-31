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
)

func start() network.QueryChannel {
	return Start("dir:.", "json")
}

func TestMainClose(t *testing.T) {
	qc := start()
	close(qc)
}

func TestSyntaxErr(t *testing.T) {
	qc := start()
	defer close(qc)

	r, _, err := Run(qc, "this is a bad query")
	if err == nil || len(r) != 0 {
		t.Errorf("expected err")
	}
	r, _, err = Run(qc, "") // empty string query
	if err == nil || len(r) != 0 {
		t.Errorf("expected err")
	}
}

func TestSimpleSelect(t *testing.T) {
	qc := start()
	defer close(qc)

	r, _, err := Run(qc, "select * from orders")
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
	qc := start()
	defer close(qc)
	matches, err := filepath.Glob("json/cases/case_*.json")
	if err != nil {
		t.Errorf("glob failed: %v", err)
	}
	for _, m := range matches {
		testCaseFile(t, m, qc)
	}
}

func testCaseFile(t *testing.T, fname string, qc network.QueryChannel) {
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

		resultsActual, _, errActual := Run(qc, statements)

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
		if !reflect.DeepEqual(resultsActual, resultsExpected) {
			t.Errorf("results don't match, actual: %#v, expected: %#v"+
				", for case file: %v, index: %v",
				resultsActual, resultsExpected, fname, i)
		}
	}
}
