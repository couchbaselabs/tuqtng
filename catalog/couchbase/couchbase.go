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
	"encoding/json"
	"fmt"

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
	site   *site
	name   string
	cbpool cb.Pool
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
	return newBucket(p, name)
}

func newPool(s *site, name string) (*pool, query.Error) {
	cbpool, err := s.client.GetPool(name)
	if err != nil {
		return nil, query.NewError(nil, fmt.Sprintf("Pool %v not found.", name))
	}
	return &pool{
		site:   s,
		name:   name,
		cbpool: cbpool,
	}, nil
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

func (b *bucket) BulkFetch(ids []string) (map[string]query.Item, query.Error) {
	rv := make(map[string]query.Item, 0)

	bulkResponse := b.cbbucket.GetBulk(ids)
	for k, v := range bulkResponse {
		var doc query.Value
		err := json.Unmarshal(v.Body, &doc)
		if err != nil {
			return nil, query.NewError(err, "")
		}

		meta := map[string]query.Value{
			"id": k,
		}

		item := query.NewParsedItem(doc, meta)

		rv[k] = item
	}

	return rv, nil
}

func (b *bucket) Fetch(id string) (query.Item, query.Error) {
	var doc query.Value
	err := b.cbbucket.Get(id, &doc)

	if err != nil {
		return nil, query.NewError(err, "")
	}

	meta := map[string]query.Value{
		"id": id,
	}

	item := query.NewParsedItem(doc, meta)

	return item, nil
}

func newBucket(p *pool, name string) (*bucket, query.Error) {
	cbbucket, err := p.cbpool.GetBucket(name)
	if err != nil {
		return nil, query.NewError(nil, fmt.Sprintf("Bucket %v not found.", name))
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

func (vs *viewScanner) ScanAll(ch query.ItemChannel, warnch, errch query.ErrorChannel) {
	go vs.scanAll(ch, warnch, errch)
}

func (vs *viewScanner) scanAll(ch query.ItemChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	viewRowChannel := make(chan cb.ViewRow)
	go WalkViewInBatches(viewRowChannel, vs.bucket.cbbucket, vs.ddoc, vs.view, map[string]interface{}{}, 1000)
	for viewRow := range viewRowChannel {
		doc := map[string]query.Value{}
		meta := map[string]query.Value{"id": viewRow.ID}
		ch <- query.NewParsedItem(doc, meta)
	}
}

func newViewScanner(b *bucket, ddoc string, view string) (*viewScanner, query.Error) {
	return &viewScanner{
		bucket: b,
		view:   view,
		ddoc:   ddoc,
	}, nil
}
