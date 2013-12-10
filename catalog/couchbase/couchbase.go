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
	"time"

	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type site struct {
	client    cb.Client
	poolCache map[string]catalog.Pool
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
	pool, ok := s.poolCache[name]
	if !ok {
		var err query.Error
		pool, err = newPool(s, name)
		if err != nil {
			return nil, err
		}
		s.poolCache[name] = pool
	}
	return pool, nil
}

// NewSite creates a new Couchbase site for the given url.
func NewSite(url string) (catalog.Site, query.Error) {
	clog.To(catalog.CHANNEL, "Created New Site %s", url)
	client, err := cb.Connect(url)

	if err != nil {
		return nil, query.NewError(err, "")
	}

	return &site{
		client:    client,
		poolCache: make(map[string]catalog.Pool),
	}, nil
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
	clog.To(catalog.CHANNEL, "Refreshing Pool %s", p.name)
	newpool, err := p.site.client.GetPool(p.name)
	if err != nil {
		clog.Warnf("Error updating pool: %v", err)
		return
	}
	p.cbpool = newpool
}

func newPool(s *site, name string) (*pool, query.Error) {
	clog.To(catalog.CHANNEL, "Created New Pool %s", name)
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
	return ViewTotalRows(b.cbbucket, "", "_all_docs", map[string]interface{}{})
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
	idx, ok := b.indexes[PRIMARY_INDEX]
	if ok {
		primary := idx.(catalog.PrimaryIndex)
		return primary, nil
	}
	all, ok := b.indexes[ALLDOCS_INDEX]
	if ok {
		primary := all.(catalog.PrimaryIndex)
		return primary, nil
	}
	return nil, query.NewError(nil, "no primary index found!")
}

func (b *bucket) IndexesByPrimary() ([]catalog.PrimaryIndex, query.Error) {
	rv := make([]catalog.PrimaryIndex, 0, 2)
	for _, index := range b.indexes {
		switch pi := index.(type) {
		case catalog.PrimaryIndex:
			rv = append(rv, pi)
		}
	}
	return rv, nil
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

	bulkResponse, err := b.cbbucket.GetBulk(ids)
	if err != nil {
		return nil, query.NewError(err, "Error doing bulk get")
	}
	for k, v := range bulkResponse {

		doc := dparval.NewValueFromBytes(v.Body)
		meta_flags := (v.Extras[0]&0xff)<<24 | (v.Extras[1]&0xff)<<16 | (v.Extras[2]&0xff)<<8 | (v.Extras[3] & 0xff)
		meta_type := "json"
		if doc.Type() == dparval.NOT_JSON {
			meta_type = "base64"
		}
		doc.SetAttachment("meta", map[string]interface{}{
			"id":    k,
			"cas":   float64(v.Cas),
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

func (b *bucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {

	if _, exists := b.indexes[PRIMARY_INDEX]; exists {
		return nil, query.NewError(nil, "Primary index already exists")
	}

	idx, err := newPrimaryIndex(b)
	if err != nil {
		return nil, query.NewError(err, "Error creating primary index")
	}
	b.indexes[idx.Name()] = idx
	return idx, nil
}

func (b *bucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {

	if using == "" {
		// current default is VIEW
		using = catalog.VIEW
	}

	//index names are unique across IndexType
	if _, exists := b.indexes[name]; exists {
		return nil, query.NewError(nil, fmt.Sprintf("Index already exists: %s", name))
	}

	switch using {
	case catalog.VIEW:
		idx, err := newViewIndex(name, key, b)
		if err != nil {
			return nil, query.NewError(err, fmt.Sprintf("Error creating index: %s", name))
		}
		b.indexes[idx.Name()] = idx
		return idx, nil

	case catalog.LSM:
		idx, err := newLsmIndex(name, key, b)
		if err != nil {
			return nil, query.NewError(err, fmt.Sprintf("Error creating index: %s", name))
		}
		b.indexes[idx.Name()] = idx
		return idx, nil

	default:
		return nil, query.NewError(nil, "Not yet implemented.")
	}
}

func (b *bucket) loadIndexes() query.Error {
	// #alldocs implicitly exists
	pi := newAllDocsIndex(b)
	b.indexes[pi.name] = pi

	// and recreate remaining from ddocs
	indexes, err := loadViewIndexes(b)
	if err != nil {
		return query.NewError(err, "Error loading view indexes")
	}

	for _, index := range indexes {
		name := (*index).Name()
		b.indexes[name] = *index
	}

	// and recreate remaining from LSM indexes
	indexes, err = loadLsmIndexesForBucket(b)
	if err != nil {
		return query.NewError(err, "Error loading lsm indexes")
	}

	for _, index := range indexes {
		name := (*index).Name()
		b.indexes[name] = *index
	}

	return nil
}

func newBucket(p *pool, name string) (*bucket, query.Error) {
	clog.To(catalog.CHANNEL, "Created New Bucket %s", name)
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
		indexes:  make(map[string]catalog.Index),
	}

	ierr := rv.loadIndexes()
	if err != nil {
		return nil, ierr
	}

	return rv, nil
}
