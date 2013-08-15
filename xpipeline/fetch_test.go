//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package xpipeline

import (
	"testing"

	"github.com/couchbaselabs/tuqtng/catalog/mock"
)

func TestFetch(t *testing.T) {

	mocksite, err := mock.NewSite("mock:items=10")
	if err != nil {
		t.Errorf("Error creating mock site")
	}
	pool, err := mocksite.Pool("p0")
	if err != nil {
		t.Errorf("Error accessing pool p0")
	}
	bucket, err := pool.Bucket("b0")
	if err != nil {
		t.Errorf("Error accessing bucket b0")
	}
	scanner, err := bucket.Scanner("all_docs")
	if err != nil {
		t.Errorf("Error accessing scanner all_docs")
	}

	scan := NewScan(bucket, scanner)
	fetch := NewFetch(bucket, nil, "bucket")
	fetch.SetSource(scan)

	fetchItemChannel, _ := fetch.GetChannels()

	go fetch.Run()

	count := 0
	for item := range fetchItemChannel {
		bucketValue, err := item.Path("bucket")
		if err != nil {
			t.Errorf("Expected item to contain value at path bucket")
		}
		iValue, err := bucketValue.Path("i")
		if err != nil {
			t.Errorf("Expected item to contain value at path i")
		}
		i := iValue.Value()
		_, ok := i.(float64)
		if !ok {
			t.Errorf("Expected i to be a number")
		}
		count++
	}

	if count != 10 {
		t.Errorf("Expected %d items, got %d", 10, count)
	}

}
