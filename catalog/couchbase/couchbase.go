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
	"log"
	"net/http"
	"time"

	"github.com/couchbaselabs/dparval"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type site struct {
	client cb.Client
}

func (s *site) URL() string {
	return s.client.BaseURL.String()
}

func (s *site) PoolNames() ([]string, query.Error) {
	// FIXME discover couchbase pools
	return []string{"default"}, nil
}

func (s *site) Pool(name string) (p catalog.Pool, e query.Error) {
	return newPool(s, name)
}

// NewSite creates a new file-based site for the given filepath.
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

func (p *pool) Name() string {
	return p.name
}

func (p *pool) BucketNames() ([]string, query.Error) {
	rv := make([]string, 0, len(p.cbpool.BucketMap))
	for name, _ := range p.cbpool.BucketMap {
		rv = append(rv, name)
	}
	return rv, nil
}

func (p *pool) Bucket(name string) (catalog.Bucket, query.Error) {
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
		log.Printf("Error updating pool: %v", err)
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
	scanners map[string]catalog.Scanner
	cbbucket *cb.Bucket
}

func (b *bucket) Release() {
	b.cbbucket.Close()
}

func (b *bucket) Name() string {
	return b.name
}

func (b *bucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, fmt.Sprintf("Count not supported on Couchbase at this time"))
}

func (b *bucket) ScannerNames() ([]string, query.Error) {
	rv := make([]string, 0, len(b.scanners))
	for name, _ := range b.scanners {
		rv = append(rv, name)
	}
	return rv, nil
}

func (b *bucket) Scanners() ([]catalog.Scanner, query.Error) {
	rv := make([]catalog.Scanner, 0, len(b.scanners))
	for _, scanner := range b.scanners {
		rv = append(rv, scanner)
	}
	return rv, nil
}

func (b *bucket) Scanner(name string) (catalog.Scanner, query.Error) {
	scanner, ok := b.scanners[name]
	if !ok {
		return nil, query.NewError(nil, fmt.Sprintf("Scanner %v not found.", name))
	}
	return scanner, nil
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

	rv.scanners = make(map[string]catalog.Scanner, 1)
	// build scanner
	allDocsScanner, err := newViewScanner(rv, "", "_all_docs")
	if err != nil {
		return nil, query.NewError(err, "")
	}
	rv.scanners[allDocsScanner.Name()] = allDocsScanner

	return rv, nil
}

type viewScanner struct {
	ddoc   string
	view   string
	bucket *bucket
}

func (vs *viewScanner) Name() string {
	return fmt.Sprintf("_design/%s/_view/%s", vs.ddoc, vs.view)
}

func (vs *viewScanner) ScanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	go vs.scanAll(ch, warnch, errch)
}

func (vs *viewScanner) scanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	viewRowChannel := make(chan cb.ViewRow)
	viewErrChannel := make(query.ErrorChannel)
	go WalkViewInBatches(viewRowChannel, viewErrChannel, vs.bucket.cbbucket, vs.ddoc, vs.view, map[string]interface{}{}, 1000)

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
					_, err := http.Get(vs.bucket.cbbucket.URI)
					if err != nil {
						// remove this specific bucket from the pool cache
						delete(vs.bucket.pool.bucketCache, vs.bucket.Name())
						// close this bucket
						vs.bucket.Release()
						// ask the pool to refresh
						vs.bucket.pool.refresh()
						// bucket doesnt exist any more
						errch <- query.NewError(nil, fmt.Sprintf("Bucket %v not found.", vs.bucket.Name()))
						return
					}

				}

				errch <- err
				return
			}
		}
	}

}

func newViewScanner(b *bucket, ddoc string, view string) (*viewScanner, query.Error) {
	return &viewScanner{
		bucket: b,
		view:   view,
		ddoc:   ddoc,
	}, nil
}
