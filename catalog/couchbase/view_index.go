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

	"github.com/couchbaselabs/dparval"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
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

func (b *bucket) createViewIndex(name string, key catalog.IndexKey) (catalog.Index, query.Error) {
	idx, err := newViewIndex(name, key, b)
	if err != nil {
		return nil, query.NewError(err, fmt.Sprintf("Cannot create index %s", name))
	}
	return idx, nil
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

func (b *bucket) LoadViewIndexes() query.Error {
	// put in indicative entry for primary index
	pi, err := newPrimaryIndex(b, "", "_all_docs")
	if err != nil {
		return query.NewError(err, "")
	}
	b.indexes[pi.Name()] = pi
	b.primary = pi

	return nil
}

func (vi *viewIndex) ScanEntries(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	vi.ScanRange(nil, nil, catalog.Both, ch, warnch, errch)
}

func (vi *viewIndex) ScanRange(low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {

	viewOptions := generateViewOptions(low, high, inclusion)

	defer close(ch)
	defer close(warnch)
	defer close(errch)

	viewRowChannel := make(chan cb.ViewRow)
	viewErrChannel := make(query.ErrorChannel)
	go WalkViewInBatches(viewRowChannel, viewErrChannel, vi.bucket.cbbucket, vi.DDocName(), vi.ViewName(), viewOptions, 1000)

	var viewRow cb.ViewRow
	var err query.Error
	sentRows := false
	ok := true
	for ok {
		select {
		case viewRow, ok = <-viewRowChannel:
			if ok {
				doc := dparval.NewValue(map[string]interface{}{})
				doc.SetAttachment("meta", map[string]interface{}{"id": viewRow.ID})
				ch <- doc
				sentRows = true
			}
		case err, ok = <-viewErrChannel:
			if err != nil {
				// check to possibly detect a bucket that was already deleted
				if !sentRows {
					_, err := http.Get(vi.bucket.cbbucket.URI)
					if err != nil {
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

func (pi *primaryIndex) IsPrimary() bool {
	return true
}

func (pi *primaryIndex) ScanBucket(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(ch, warnch, errch)
}

func (vi *viewIndex) Check(value catalog.LookupValue, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	vi.ScanRange(value, value, catalog.Both, ch, warnch, errch)
}

func (vi *viewIndex) Lookup(value catalog.LookupValue, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	vi.ScanRange(value, value, catalog.Both, ch, warnch, errch)
}

func (vi *viewIndex) Statistics() (catalog.RangeStatistics, query.Error) {
	return nil, query.NewError(nil, "statistics not implemented")
}

func (vi *viewIndex) Direction() catalog.Direction {
	return catalog.ASC
}
