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
	on     catalog.IndexKey
	ddoc   *designdoc
	bucket *bucket
}

type designdoc struct {
	mapfn    string
	reducefn string
}

type ddocJSON struct {
	cb.DDocJSON
	indexOn       string  `json:"indexOn"`
	indexChecksum int     `json:"indexChecksum"`
}

func (vi *viewIndex) BucketId() string {
	return vi.bucket.Id()
}

func (vi *viewIndex) Id() string {
	return vi.Name()
}

func (vi *viewIndex) Name() string {
	return fmt.Sprintf("%s/%s", vi.bucket.name, vi.name)
}

func (vi *viewIndex) Type() catalog.IndexType {
	return catalog.VIEW
}

func (vi *viewIndex) Key() catalog.IndexKey {
	return vi.on
}

func (b *bucket) createViewIndex(name string, key catalog.IndexKey) (catalog.Index, query.Error) {
	idx, err := newViewIndex(name, key, b)
	if (idx != nil) {
		return nil, query.NewError(err, fmt.Sprintf("Cannot create index %s", name))
	}
	return idx, nil
}

func (vi *viewIndex) Drop() query.Error {
	err := vi.Drop()
	if err != nil {
		return query.NewError(err, fmt.Sprintf("Cannot drop index %s", vi.Name()))
	}
	return nil
}

func (idx *viewIndex) DDocName() string {
	return "query_" + idx.Name()
}

func (idx *viewIndex) ViewName() string {
	return "autogen"
}

func (b *bucket) LoadViewIndexes() query.Error {
	b.indexes = make(map[string]catalog.Index, 1)

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

type primaryIndex struct {
	viewIndex
}

func (pi *primaryIndex) Type() catalog.IndexType {
	return catalog.PRIMARY
}

func (pi *primaryIndex) Key() catalog.IndexKey {
	return pi.on
}

func (pi *primaryIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *primaryIndex) DDocName() string {
	return ""
}

func (pi *primaryIndex) ViewName() string {
	return "_all_docs"
}

func newPrimaryIndex(b *bucket, ddoc string, view string) (*primaryIndex, query.Error) {
	// FIXME
	return &primaryIndex{
		viewIndex{
			name:   "primary",
			on:     make(catalog.IndexKey, 0),
			ddoc:   nil,
			bucket: b,
		},
	}, nil
}
