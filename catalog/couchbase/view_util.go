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
	"log"
	"strings"

	cb "github.com/couchbaselabs/go-couchbase"
)

var MIN_KEY map[string]interface{}
var MAX_KEY = map[string]interface{}{} //FIXME better approximation of max
// a key in Couchbase is at most 250 bytes
// in a view, items with the same key are sorted by id (using basic memcmp comparison)
// so 251 bytes of 0xff is higher than any valid key in couchbase
var MAX_ID = cb.DocId(strings.Repeat(string([]byte{0xff}), 251))

func WalkViewInBatches(result chan cb.ViewRow, bucket *cb.Bucket,
	ddoc string, view string, options map[string]interface{}, batchSize int) {

	defer close(result)

	options["limit"] = batchSize + 1
	logURL, err := bucket.ViewURL(ddoc, view, options)
	if err == nil {
		log.Printf("Request View: %v", logURL)
	}
	vres, err := bucket.View(ddoc, view, options)

	if err != nil {
		log.Printf("Error accessing view: %v", err)
		return
	}

	for i, row := range vres.Rows {
		if i < batchSize {
			// dont process the last row, its just used to see if we
			// need to continue processing
			result <- row
		}
	}

	// as long as we continue to get batchSize + 1 results back we have to keep going
	for len(vres.Rows) > batchSize {
		skey := vres.Rows[batchSize].Key
		skeydocid := vres.Rows[batchSize].ID
		options["startkey"] = skey
		options["startkey_docid"] = cb.DocId(skeydocid)

		logURL, err := bucket.ViewURL(ddoc, view, options)
		if err == nil {
			log.Printf("Request View: %v", logURL)
		}
		vres, err = bucket.View(ddoc, view, options)
		if err != nil {
			log.Printf("Error accessing view: %v", err)
			return
		}

		for i, row := range vres.Rows {
			if i < batchSize {
				// dont process the last row, its just used to see if we
				// need to continue processing
				result <- row
			}
		}
	}
}
