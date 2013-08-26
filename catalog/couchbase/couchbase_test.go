//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package couchbase

import (
	"testing"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/catalog"
)

// not named like a proper test function
// because requires couchbase server
func notTestCouchbase(t *testing.T) {
	site, err := NewSite("http://localhost:8091")
	if err != nil {
		t.Errorf("failed to create site: %v", err)
	}

	poolIds, err := site.PoolIds()
	if err != nil {
		t.Errorf("failed to get pool ids: %v", err)
	}

	if len(poolIds) != 1 || poolIds[0] != "default" {
		t.Errorf("expected 1 pool id'd default")
	}

	pool, err := site.PoolById("default")
	if err != nil {
		t.Errorf("failed to get pool by id: %v", err)
	}

	poolNames, err := site.PoolNames()
	if err != nil {
		t.Errorf("failed to get pool names: %v", err)
	}

	if len(poolNames) != 1 || poolNames[0] != "default" {
		t.Errorf("expected 1 pool named default")
	}

	pool, err = site.PoolByName("default")
	if err != nil {
		t.Errorf("failed to get pool by name: %v", err)
	}

	_, err = pool.BucketIds()
	if err != nil {
		t.Errorf("failed to get bucket ids: %v", err)
	}

	bucket, err := pool.BucketById("beer-sample")
	if err != nil {
		t.Errorf("failed to get bucket by id: beer-sample")
	}

	_, err = pool.BucketNames()
	if err != nil {
		t.Errorf("failed to get bucket names: %v", err)
	}

	bucket, err = pool.BucketByName("beer-sample")
	if err != nil {
		t.Errorf("failed to get bucket by name: beer-sample")
	}

	indexes, err := bucket.Indexes()
	if err != nil {
		t.Errorf("failed ot get indexes")
	}

	if len(indexes) < 1 {
		t.Errorf("Expected at least 1 index for bucket")
	}

	index := indexes[0]
	switch index := index.(type) {
	case catalog.ScanIndex:
	        si := index.(catalog.ScanIndex)
		itemChannel := make(dparval.ValueChannel)
		warnChannel := make(query.ErrorChannel)
		errorChannel := make(query.ErrorChannel)
		go si.ScanEntries(itemChannel, warnChannel, errorChannel)

		var err query.Error
		ok := true
		for ok {
			select {
			case _, ok = <-itemChannel:
			case err, ok = <-errorChannel:
				if err != nil {
					t.Errorf("got error while scanning: %v", err)
				}
			}
		}
	}

	_, err = bucket.Fetch("21st_amendment_brewery_cafe-watermelon_wheat")
	if err != nil {
		t.Errorf("failed to fetch 21st_amendment_brewery_cafe-watermelon_wheat: %v", err)
	}
}
