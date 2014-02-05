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
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/couchbaselabs/tuqtng/query"
)

type tuqError struct {
	Code    float64 `json:"code,omitempty"`
	Key     string  `json:"key,omitempty"`
	Message string  `json:"message,omitempty"`
	Cause   string  `json:"cause,omitempty"`
	Caller  string  `json:"caller,omitempty"`
}

type tuqResponse struct {
	Resultset []interface{} `json:"resultset,omitempty"`
	Info      []tuqError    `json:"info,omitempty"`
	Warnings  []tuqError    `json:"warnings,omitempty"`
	Error     *tuqError     `json:"error,omitempty"`
}

func TestHttpResponseNoResults(t *testing.T) {

	// NOTE the query isn't actually used, its just to allow the constructor to work correctly
	req, err := http.NewRequest("POST", "http://localhost:8093/query", strings.NewReader("SELECT * FROM bucket"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resrec := httptest.NewRecorder()
	q := NewHttpQuery(resrec, req, true)

	res := q.Response()
	go func() {
		res.NoMoreResults()
	}()
	q.Process()

	var tuqRes tuqResponse
	err = json.Unmarshal(resrec.Body.Bytes(), &tuqRes)
	if err != nil {
		t.Logf("`%s`", resrec.Body.String())
		t.Errorf("tuq response didn't parse as json: %v", err)
	}

	if len(tuqRes.Resultset) != 0 {
		t.Errorf("expected 0 rows, got %d", len(tuqRes.Resultset))
	}
	if tuqRes.Info != nil {
		found := false
		for _, info := range tuqRes.Info {
			if info.Key == "total_rows" && info.Message == "0" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected total_rows = 0")
		}
	} else {
		t.Errorf("exptected info section, was nil")
	}
}

func TestHttpResponseSomeResults(t *testing.T) {

	// NOTE the query isn't actually used, its just to allow the constructor to work correctly
	req, err := http.NewRequest("POST", "http://localhost:8093/query", strings.NewReader("SELECT * FROM bucket"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resrec := httptest.NewRecorder()
	q := NewHttpQuery(resrec, req, true)

	res := q.Response()
	go func() {
		res.SendResult(map[string]interface{}{"name": "marty"})
		res.SendResult(map[string]interface{}{"name": "steve"})
		res.SendResult(map[string]interface{}{"name": "gerald"})
		res.SendResult(map[string]interface{}{"name": "siri"})
		res.NoMoreResults()
	}()
	q.Process()

	var tuqRes tuqResponse
	err = json.Unmarshal(resrec.Body.Bytes(), &tuqRes)
	if err != nil {
		t.Logf("`%s`", resrec.Body.String())
		t.Errorf("tuq response didn't parse as json: %v", err)
	}

	if len(tuqRes.Resultset) != 4 {
		t.Errorf("expected 4 rows, got %d", len(tuqRes.Resultset))
	}
	if tuqRes.Info != nil {
		found := false
		for _, info := range tuqRes.Info {
			if info.Key == "total_rows" && info.Message == "4" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected total_rows = 4")
		}
	} else {
		t.Errorf("exptected info section, was nil")
	}
}

func TestHttpErrorResponseNoResults(t *testing.T) {

	// NOTE the query isn't actually used, its just to allow the constructor to work correctly
	req, err := http.NewRequest("POST", "http://localhost:8093/query", strings.NewReader("SELECT * FROM bucket"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resrec := httptest.NewRecorder()
	q := NewHttpQuery(resrec, req, true)

	res := q.Response()
	go func() {
		res.SendError(query.NewError(fmt.Errorf("test error"), "internal msg"))
	}()
	q.Process()

	var tuqRes tuqResponse
	err = json.Unmarshal(resrec.Body.Bytes(), &tuqRes)
	if err != nil {
		t.Logf("`%s`", resrec.Body.String())
		t.Errorf("tuq response didn't parse as json: %v", err)
	}

	if tuqRes.Resultset != nil {
		t.Errorf("expected nil resultset")
	}
	if tuqRes.Error != nil {
		if tuqRes.Error.Key != "Internal Error" {
			t.Errorf("exptected key `Internal Error`, got %s", tuqRes.Error.Key)
		}
		if tuqRes.Error.Message != "internal msg" {
			t.Errorf("expected message `internal msg`, got %s", tuqRes.Error.Message)
		}
		if tuqRes.Error.Cause != "test error" {
			t.Errorf("expected cause `test error`, got %s", tuqRes.Error.Cause)
		}
	} else {
		t.Errorf("exptected error section, was nil")
	}
}

func TestHttpErrorResponseSomeResults(t *testing.T) {

	// NOTE the query isn't actually used, its just to allow the constructor to work correctly
	req, err := http.NewRequest("POST", "http://localhost:8093/query", strings.NewReader("SELECT * FROM bucket"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resrec := httptest.NewRecorder()
	q := NewHttpQuery(resrec, req, true)

	res := q.Response()
	go func() {
		res.SendResult(map[string]interface{}{"name": "marty"})
		res.SendResult(map[string]interface{}{"name": "steve"})
		res.SendError(query.NewError(fmt.Errorf("test error"), "internal msg"))
	}()
	q.Process()

	var tuqRes tuqResponse
	err = json.Unmarshal(resrec.Body.Bytes(), &tuqRes)
	if err != nil {
		t.Logf("`%s`", resrec.Body.String())
		t.Errorf("tuq response didn't parse as json: %v", err)
	}

	if len(tuqRes.Resultset) != 2 {
		t.Errorf("expected 4 rows, got %d", len(tuqRes.Resultset))
	}

	if tuqRes.Error != nil {
		if tuqRes.Error.Key != "Internal Error" {
			t.Errorf("exptected key `Internal Error`, got %s", tuqRes.Error.Key)
		}
		if tuqRes.Error.Message != "internal msg" {
			t.Errorf("expected message `internal msg`, got %s", tuqRes.Error.Message)
		}
		if tuqRes.Error.Cause != "test error" {
			t.Errorf("expected cause `test error`, got %s", tuqRes.Error.Cause)
		}
	} else {
		t.Errorf("exptected error section, was nil")
	}
}
