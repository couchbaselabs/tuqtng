//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package file

import (
	"log"
	"testing"

	"github.com/couchbaselabs/tuqtng/query"
)

func TestFile(t *testing.T) {
	site, err := NewSite("../../test")
	if err != nil {
		t.Errorf("failed to create site: %v", err)
	}

	poolNames, err := site.PoolNames()
	if err != nil {
		t.Errorf("failed to get pool names: %v", err)
	}

	if len(poolNames) != 1 || poolNames[0] != "json" {
		t.Errorf("expected 1 pool named json")
	}

	pool, err := site.Pool("json")
	if err != nil {
		t.Errorf("failed to get pool: %v", err)
	}

	_, err = pool.BucketNames()
	if err != nil {
		t.Errorf("failed to get bucket names: %v", err)
	}

	bucket, err := pool.Bucket("contacts")
	if err != nil {
		t.Errorf("failed to get bucket contacts")
	}

	scanners, err := bucket.Scanners()
	if err != nil {
		t.Errorf("failed ot get scanners")
	}

	if len(scanners) < 1 {
		t.Errorf("Expected at least 1 scanner for bucket")
	}

	scanner := scanners[0]
	switch scanner := scanner.(type) {
	case *fullScanner:
		itemChannel := make(query.ItemChannel)
		warnChannel := make(query.ErrorChannel)
		errorChannel := make(query.ErrorChannel)
		go scanner.ScanAll(itemChannel, warnChannel, errorChannel)

		var item query.Item
		var err query.Error
		ok := true
		for ok {
			select {
			case item, ok = <-itemChannel:
				log.Printf("got item %v", item)
			case err, ok = <-errorChannel:
				if err != nil {
					t.Errorf("got error while scanning: %v", err)
				}
			}
		}
	default:
		log.Printf("not full scanner %T", scanner)
	}

	doc, err := bucket.Fetch("fred")
	if err != nil {
		t.Errorf("failed to fetch fred: %v", err)
	}

	log.Printf("%v", doc)

}
