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
	"net/http"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

const (
	PRIMARY_INDEX = "#primary"
	ALLDOCS_INDEX = "#alldocs"
)

type viewIndex struct {
	name   string
	using  catalog.IndexType
	on     catalog.IndexKey
	ddoc   *designdoc
	bucket *bucket
}

type primaryIndex struct {
	viewIndex
}

type designdoc struct {
	name     string
	viewname string
	mapfn    string
	reducefn string
}

func (vi *viewIndex) BucketId() string {
	return vi.bucket.Id()
}

func (vi *viewIndex) Id() string {
	return vi.name
}

func (vi *viewIndex) Name() string {
	return vi.name
}

func (vi *viewIndex) Type() catalog.IndexType {
	return vi.using
}

func (vi *viewIndex) IsPrimary() bool {
	return false
}

func (vi *viewIndex) Key() catalog.IndexKey {
	return vi.on
}

func (idx *viewIndex) DDocName() string {
	return idx.ddoc.name
}

func (idx *viewIndex) ViewName() string {
	return idx.ddoc.viewname
}

func (vi *viewIndex) Drop() query.Error {
	bucket := vi.bucket
	if vi.IsPrimary() {
		return query.NewError(nil, "Primary index cannot be dropped.")
	}
	err := vi.DropViewIndex()
	if err != nil {
		return query.NewError(err, fmt.Sprintf("Cannot drop index %s", vi.Name()))
	}
	delete(bucket.indexes, vi.name)
	return nil
}

func (vi *viewIndex) ScanEntries(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	vi.ScanRange(nil, nil, catalog.Both, limit, ch, warnch, errch)
}

func (vi *viewIndex) ValueCount() (int64, query.Error) {
	indexItemChannel := make(catalog.EntryChannel)
	indexWarnChannel := make(query.ErrorChannel)
	indexErrorChannel := make(query.ErrorChannel)

	go vi.ScanRange(catalog.LookupValue{dparval.NewValue(nil)}, catalog.LookupValue{dparval.NewValue(nil)}, catalog.Both, 0, indexItemChannel, indexWarnChannel, indexErrorChannel)

	var err query.Error
	nullCount := int64(0)
	ok := true
	for ok {
		select {
		case _, ok = <-indexItemChannel:
			if ok {
				nullCount += 1
			}
		case _, ok = <-indexWarnChannel:
			// ignore warnings here
		case err, ok = <-indexErrorChannel:
			if err != nil {
				return 0, err
			}
		}
	}

	totalCount, err := ViewTotalRows(vi.bucket.cbbucket, vi.DDocName(), vi.ViewName(), map[string]interface{}{})
	if err != nil {
		return 0, err
	}

	return totalCount - nullCount, nil

}

func (vi *viewIndex) ScanRange(low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion, limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {

	viewOptions := generateViewOptions(low, high, inclusion)

	defer close(ch)
	defer close(warnch)
	defer close(errch)

	viewRowChannel := make(chan cb.ViewRow)
	viewErrChannel := make(query.ErrorChannel)
	go WalkViewInBatches(viewRowChannel, viewErrChannel, vi.bucket.cbbucket, vi.DDocName(), vi.ViewName(), viewOptions, 1000, limit)

	var viewRow cb.ViewRow
	var err query.Error
	sentRows := false
	ok := true
	for ok {
		select {
		case viewRow, ok = <-viewRowChannel:
			if ok {
				entry := catalog.IndexEntry{PrimaryKey: viewRow.ID}

				// try to add the view row key as the entry key (unless this is _all_docs)
				if vi.DDocName() != "" {
					lookupValue, err := convertCouchbaseViewKeyToLookupValue(viewRow.Key)
					if err == nil {
						entry.EntryKey = lookupValue
					} else {
						clog.To(catalog.CHANNEL, "unable to convert index key to lookup value:%v", err)
					}
				}

				ch <- &entry
				sentRows = true
			}
		case err, ok = <-viewErrChannel:
			if err != nil {
				clog.Error(err)
				// check to possibly detect a bucket that was already deleted
				if !sentRows {
					clog.Printf("Checking bucket URI: %v", vi.bucket.cbbucket.URI)
					_, err := http.Get(vi.bucket.cbbucket.URI)
					if err != nil {
						clog.Error(err)
						// remove this specific bucket from the pool cache
						delete(vi.bucket.pool.bucketCache, vi.bucket.Name())
						// close this bucket
						vi.bucket.Release()
						// ask the pool to refresh
						vi.bucket.pool.refresh()
						// bucket doesnt exist any more
						errch <- query.NewError(nil, fmt.Sprintf("Bucket %v not found.", vi.bucket.Name()))
						return
					}

				}

				errch <- err
				return
			}
		}
	}
}

func (pi *primaryIndex) BucketId() string {
	return pi.viewIndex.bucket.Id()
}

func (pi *primaryIndex) IsPrimary() bool {
	return true
}

func (pi *primaryIndex) ScanBucket(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (vi *viewIndex) Check(value catalog.LookupValue, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	vi.ScanRange(value, value, catalog.Both, 0, ch, warnch, errch)
}

func (vi *viewIndex) Lookup(value catalog.LookupValue, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	vi.ScanRange(value, value, catalog.Both, 0, ch, warnch, errch)
}

func (vi *viewIndex) Statistics() (catalog.RangeStatistics, query.Error) {
	return nil, query.NewError(nil, "statistics not implemented")
}

func (vi *viewIndex) Direction() catalog.Direction {
	return catalog.ASC
}
