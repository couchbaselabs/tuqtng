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
	"sort"
	"strings"

	"github.com/couchbaselabs/clog"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

const NETWORK_CHANNEL = "NETWORK"

const TYPE_NULL = 64
const TYPE_BOOLEAN = 96
const TYPE_NUMBER = 128
const TYPE_STRING = 160
const TYPE_ARRAY = 192
const TYPE_OBJECT = 224

var MIN_ID = cb.DocId("")
var MAX_ID = cb.DocId(strings.Repeat(string([]byte{0xff}), 251))

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
			clog.To(NETWORK_CHANNEL, "Request View: %v", logURL)
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

func generateViewOptions(low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion) map[string]interface{} {
	viewOptions := map[string]interface{}{}

	if low != nil {
		viewOptions["startkey"] = encodeValueAsMapKey(low)
		if inclusion == catalog.Neither || inclusion == catalog.High {
			viewOptions["startkey_docid"] = MAX_ID
		}
	}

	if high != nil {
		viewOptions["endkey"] = encodeValueAsMapKey(high)
		if inclusion == catalog.Neither || inclusion == catalog.Low {
			viewOptions["endkey_docid"] = MIN_ID
		}
	}

	return viewOptions
}

func encodeValueAsMapKey(keys catalog.LookupValue) interface{} {
	rv := make([]interface{}, len(keys))
	for i, lv := range keys {
		val := lv.Value()
		rv[i] = encodeValue(val)
	}
	return rv
}

func encodeValue(val interface{}) interface{} {
	switch val := val.(type) {
	case nil:
		return []interface{}{TYPE_NULL}
	case bool:
		return []interface{}{TYPE_BOOLEAN, val}
	case float64:
		return []interface{}{TYPE_NUMBER, val}
	case string:
		return []interface{}{TYPE_STRING, encodeStringAsNumericArray(val)}
	case []interface{}:
		return []interface{}{TYPE_ARRAY, val}
	case map[string]interface{}:
		return []interface{}{TYPE_OBJECT, encodeObjectAsCompoundArray(val)}
	default:
		panic(fmt.Sprintf("Unable to encode type %T to map key", val))
	}
}

func encodeStringAsNumericArray(str string) []float64 {
	rv := make([]float64, len(str))
	for i, rune := range str {
		rv[i] = float64(rune)
	}
	return rv
}

func encodeObjectAsCompoundArray(obj map[string]interface{}) []interface{} {
	keys := make([]string, len(obj))
	counter := 0
	for k, _ := range obj {
		keys[counter] = k
		counter++
	}
	sort.Strings(keys)
	vals := make([]interface{}, len(obj))
	for i, key := range keys {
		vals[i] = encodeValue(obj[key])
	}
	return []interface{}{keys, vals}
}
