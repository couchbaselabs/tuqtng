//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

// +build couchbase

package test

import (
	"path/filepath"
	"testing"

	"github.com/couchbaselabs/tuqtng/network"
)

func startCouchbase() network.QueryChannel {
	return Start("http://localhost:8091", "default")
}

func TestAllCaseFilesCouchbase(t *testing.T) {
	qc := startCouchbase()
	defer close(qc)
	matches, err := filepath.Glob("json/cases/case_*.json")
	if err != nil {
		t.Errorf("glob failed: %v", err)
	}
	for _, m := range matches {
		testCaseFile(t, m, qc)
	}

}
