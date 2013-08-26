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
	"time"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type site struct {
	client cb.Client
}

func (s *site) Id() string {
	return s.URL()
}

func (s *site) URL() string {
	return s.client.BaseURL.String()
}

func (s *site) PoolIds() ([]string, query.Error) {
	// FIXME discover couchbase pools
	return s.PoolNames()
}

func (s *site) PoolNames() ([]string, query.Error) {
	// FIXME discover couchbase pools
	return []string{"default"}, nil
}

func (s *site) PoolById(id string) (p catalog.Pool, e query.Error) {
	return s.PoolByName(id)
}

func (s *site) PoolByName(name string) (p catalog.Pool, e query.Error) {
	return newPool(s, name)
}

// NewSite creates a new Couchbase site for the given url.
func NewSite(url string) (catalog.Site, query.Error) {

	client, err := cb.Connect(url)

	if err != nil {
		return nil, query.NewError(err, "")
	}

	return &site{client: client}, nil
}

type pool struct {
	site        *site
	name        string
	cbpool      cb.Pool
	bucketCache map[string]catalog.Bucket
}

func (p *pool) SiteId() string {
	return p.site.Id()
}

func (p *pool) Id() string {
	return p.Name()
}

func (p *pool) Name() string {
	return p.name
}

func (p *pool) BucketIds() ([]string, query.Error) {
        return p.BucketNames()
}

func (p *pool) BucketNames() ([]string, query.Error) {
	rv := make([]string, 0, len(p.cbpool.BucketMap))
	for name, _ := range p.cbpool.BucketMap {
		rv = append(rv, name)
	}
	return rv, nil
}

func (p *pool) BucketById(id string) (catalog.Bucket, query.Error) {
	return p.BucketByName(id)
}

func (p *pool) BucketByName(name string) (catalog.Bucket, query.Error) {
	bucket, ok := p.bucketCache[name]
	if !ok {
		var err query.Error
		bucket, err = newBucket(p, name)
		if err != nil {
			return nil, err
		}
		p.bucketCache[name] = bucket
	}
	return bucket, nil
}

func (p *pool) refresh() {
	// trigger refresh of this pool
	newpool, err := p.site.client.GetPool(p.name)
	if err != nil {
		clog.Warnf("Error updating pool: %v", err)
		return
	}
	p.cbpool = newpool
}

func newPool(s *site, name string) (*pool, query.Error) {
	cbpool, err := s.client.GetPool(name)
	if err != nil {
		return nil, query.NewError(nil, fmt.Sprintf("Pool %v not found.", name))
	}
	rv := pool{
		site:        s,
		name:        name,
		cbpool:      cbpool,
		bucketCache: make(map[string]catalog.Bucket),
	}
	go keepPoolFresh(&rv)
	return &rv, nil
}

func keepPoolFresh(p *pool) {

	tickChan := time.Tick(5 * time.Minute)

	for _ = range tickChan {
		p.refresh()
	}
}

type bucket struct {
	pool     *pool
	name     string
	indexes  map[string]catalog.Index
	primary  catalog.PrimaryIndex
	cbbucket *cb.Bucket
}

func (b *bucket) Release() {
	b.cbbucket.Close()
}

func (b *bucket) PoolId() string {
	return b.pool.Id()
}

func (b *bucket) Id() string {
	return b.Name()
}

func (b *bucket) Name() string {
	return b.name
}

func (b *bucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, fmt.Sprintf("Count not supported on Couchbase at this time"))
}

func (b *bucket) IndexIds() ([]string, query.Error) {
	return b.IndexNames()
}

func (b *bucket) IndexNames() ([]string, query.Error) {
	rv := make([]string, 0, len(b.indexes))
	for name, _ := range b.indexes {
		rv = append(rv, name)
	}
	return rv, nil
}

func (b *bucket) IndexById(id string) (catalog.Index, query.Error) {
	return b.IndexByName(id)
}

func (b *bucket) IndexByName(name string) (catalog.Index, query.Error) {
	index, ok := b.indexes[name]
	if !ok {
		return nil, query.NewError(nil, fmt.Sprintf("Index %v not found.", name))
	}
	return index, nil
}

func (b *bucket) IndexByPrimary() (catalog.PrimaryIndex, query.Error) {
	return b.primary, nil
}

func (b *bucket) Indexes() ([]catalog.Index, query.Error) {
	rv := make([]catalog.Index, 0, len(b.indexes))
	for _, index := range b.indexes {
		rv = append(rv, index)
	}
	return rv, nil
}

func (b *bucket) BulkFetch(ids []string) (map[string]*dparval.Value, query.Error) {
	rv := make(map[string]*dparval.Value, 0)

	bulkResponse := b.cbbucket.GetBulk(ids)
	for k, v := range bulkResponse {

		doc := dparval.NewValueFromBytes(v.Body)
		meta_flags := (v.Extras[0]&0xff)<<24 | (v.Extras[1]&0xff)<<16 | (v.Extras[2]&0xff)<<8 | (v.Extras[3] & 0xff)
		meta_type := "json"
		if doc.Type() == dparval.NOT_JSON {
			meta_type = "base64"
		}
		doc.SetAttachment("meta", map[string]interface{}{
			"id":    k,
			"case":  float64(v.Cas),
			"type":  meta_type,
			"flags": float64(meta_flags),
		})

		rv[k] = doc
	}

	return rv, nil
}

func (b *bucket) Fetch(id string) (*dparval.Value, query.Error) {
	// use bulk get of single key
	values, err := b.BulkFetch([]string{id})
	if err != nil {
		return nil, err
	}
	return values[id], nil
}

// FIXME
func (b *bucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
        if b.primary != nil {
	        return b.primary, nil
	}

        return nil, query.NewError(nil, "Not yet implemented.")
}

// FIXME
func (b *bucket) CreateIndex(name string, key []string, using string) (catalog.Index, query.Error) {
        return nil, query.NewError(nil, "Not yet implemented.")
}

func newBucket(p *pool, name string) (*bucket, query.Error) {
	cbbucket, err := p.cbpool.GetBucket(name)
	if err != nil {
		// go-couchbase caches the buckets
		// to be sure no such bucket exists right now
		// we trigger a refresh
		p.refresh()
		// and then check one more time
		cbbucket, err = p.cbpool.GetBucket(name)
		if err != nil {
			// really no such bucket exists
			return nil, query.NewError(nil, fmt.Sprintf("Bucket %v not found.", name))
		}
	}

	rv := &bucket{
		pool:     p,
		name:     name,
		cbbucket: cbbucket,
	}

	rv.indexes = make(map[string]catalog.Index, 1)
	// build index
	pi, err := newPrimaryIndex(rv, "", "_all_docs")
	if err != nil {
		return nil, query.NewError(err, "")
	}
	rv.indexes[pi.Name()] = pi
	rv.primary = pi

	return rv, nil
}

type viewIndex struct {
	ddoc   string
	view   string
	bucket *bucket
}

func (vi *viewIndex) BucketId() string {
	return vi.bucket.Id()
}

func (vi *viewIndex) Id() string {
	return vi.Name()
}

func (vi *viewIndex) Name() string {
	return fmt.Sprintf("_design/%s/_view/%s", vi.ddoc, vi.view)
}

func (vi *viewIndex) Type() string {
	return ("view")
}

func (vi *viewIndex) Key() catalog.IndexKey {
        // FIXME
	return nil
}

func (vi *viewIndex) Drop() query.Error {
        // FIXME
	return query.NewError(nil, "Not yet implemented.")
}

func (vi *viewIndex) ScanEntries(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	viewRowChannel := make(chan cb.ViewRow)
	viewErrChannel := make(query.ErrorChannel)
	go WalkViewInBatches(viewRowChannel, viewErrChannel, vi.bucket.cbbucket, vi.ddoc, vi.view, map[string]interface{}{}, 1000)

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

func newViewIndex(b *bucket, ddoc string, view string) (*viewIndex, query.Error) {
	return &viewIndex{
		bucket: b,
		view:   view,
		ddoc:   ddoc,
	}, nil
}

type primaryIndex struct {
        viewIndex
}

func (pi *primaryIndex) Type() string {
	return ("primary")
}

func (pi *primaryIndex) Key() catalog.IndexKey {
	return catalog.IndexKey{"meta().id"}
}

func (pi *primaryIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func newPrimaryIndex(b *bucket, ddoc string, view string) (*primaryIndex, query.Error) {
	return &primaryIndex{
	        viewIndex {
			bucket: b,
			view:   view,
			ddoc:   ddoc,
		},
	}, nil
}
