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
	"fmt"
	"runtime/debug"

	"github.com/couchbaselabs/clog"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/query"
)

const CHANNEL = "NETWORK"

func WalkViewInBatches(result chan cb.ViewRow, errs query.ErrorChannel, bucket *cb.Bucket,
	ddoc string, view string, options map[string]interface{}, batchSize int) {

	defer close(result)
	defer close(errs)

	defer func() {
		r := recover()
		if r != nil {
			clog.Error(fmt.Errorf("View Walking Panic: %v\n%s", r, debug.Stack()))
			errs <- query.NewError(nil, "Panic In View Walking")
		}
	}()

	options["limit"] = batchSize + 1

	ok := true
	for ok {

		logURL, err := bucket.ViewURL(ddoc, view, options)
		if err == nil {
			clog.To(CHANNEL, "Request View: %v", logURL)
		}
		vres, err := bucket.View(ddoc, view, options)
		if err != nil {
			errs <- query.NewError(err, "Unable to access view")
			return
		}

		for i, row := range vres.Rows {
			if i < batchSize {
				// dont process the last row, its just used to see if we
				// need to continue processing
				result <- row
			}
		}

		if len(vres.Rows) > batchSize {
			// prepare for next run
			skey := vres.Rows[batchSize].Key
			skeydocid := vres.Rows[batchSize].ID
			options["startkey"] = skey
			options["startkey_docid"] = cb.DocId(skeydocid)
		} else {
			// stop
			ok = false
		}
	}
}
