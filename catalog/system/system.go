//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package system

import (
	"fmt"
	"strings"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/tuqtng/ast"
)

const POOL_ID   = "sys_catalog"
const POOL_NAME = "sys_catalog"
const BUCKET_NAME_SITES = "sites"
const BUCKET_NAME_POOLS = "pools"
const BUCKET_NAME_BUCKETS = "buckets"
const BUCKET_NAME_INDEXES = "indexes"

type site struct {
	actualSite        catalog.Site
	systemCatalogPool *pool
}

func (s *site) Id() string {
	return s.actualSite.Id()
}

func (s *site) URL() string {
	return s.actualSite.URL()
}

func (s *site) PoolIds() ([]string, query.Error) {
	poolIds, err := s.actualSite.PoolIds()
	if err != nil {
		return nil, err
	}
	poolIds = append(poolIds, s.systemCatalogPool.Id())
	return poolIds, err
}

func (s *site) PoolNames() ([]string, query.Error) {
	poolNames, err := s.actualSite.PoolNames()
	if err != nil {
		return nil, err
	}
	poolNames = append(poolNames, s.systemCatalogPool.Name())
	return poolNames, err
}

func (s *site) PoolById(id string) (catalog.Pool, query.Error) {
	if id == POOL_ID {
		return s.systemCatalogPool, nil
	}
	return s.actualSite.PoolById(id)
}

func (s *site) PoolByName(name string) (catalog.Pool, query.Error) {
	if name == POOL_NAME {
		return s.systemCatalogPool, nil
	}
	return s.actualSite.PoolByName(name)
}

func NewSite(actualSite catalog.Site) (catalog.Site, query.Error) {
	s := &site{actualSite: actualSite}

	e := s.loadPool()
	if e != nil {
		return nil, e
	}

	return s, e
}

func (s *site) loadPool() query.Error {
	p, e := newPool(s)
	if e != nil {
		return e
	}

	s.systemCatalogPool = p
	return nil
}

type pool struct {
	site    *site
	id      string
	name    string
	buckets map[string]catalog.Bucket
}

func (p *pool) SiteId() string {
	return p.site.Id()
}

func (p *pool) Id() string {
	return p.id
}

func (p *pool) Name() string {
	return p.name
}

func (p *pool) BucketIds() ([]string, query.Error) {
	return p.BucketNames()
}

func (p *pool) BucketNames() ([]string, query.Error) {
	rv := make([]string, len(p.buckets))
	i := 0
	for k, _ := range p.buckets {
		rv[i] = k
		i = i + 1
	}
	return rv, nil
}

func (p *pool) BucketById(id string) (catalog.Bucket, query.Error) {
	return p.BucketByName(id)
}

func (p *pool) BucketByName(name string) (catalog.Bucket, query.Error) {
	b, ok := p.buckets[name]
	if !ok {
		return nil, query.NewError(nil, "Bucket "+name+" not found.")
	}

	return b, nil
}

// newPool creates a new pool.
func newPool(s *site) (*pool, query.Error) {
	p := new(pool)
	p.site = s
	p.id = POOL_ID
	p.name = POOL_NAME
	p.buckets = make(map[string]catalog.Bucket)

	e := p.loadBuckets()
	if e != nil {
		return nil, e
	}
	return p, nil
}

func (p *pool) loadBuckets() (e query.Error) {

	sb, e := newSitesBucket(p)
	if e != nil {
		return e
	}
	p.buckets[sb.Name()] = sb

	pb, e := newPoolsBucket(p)
	if e != nil {
		return e
	}
	p.buckets[pb.Name()] = pb

	bb, e := newBucketsBucket(p)
	if e != nil {
		return e
	}
	p.buckets[bb.Name()] = bb

	return nil
}

// sites bucket
type sitebucket struct {
	pool    *pool
	name    string
	primary catalog.PrimaryIndex
}

func (b *sitebucket) Release() {
}

func (b *sitebucket) PoolId() string {
	return b.pool.Id()
}

func (b *sitebucket) Id() string {
	return b.Name()
}

func (b *sitebucket) Name() string {
	return b.name
}

func (b *sitebucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *sitebucket) IndexIds() ([]string, query.Error) {
	return []string{b.primary.Id()}, nil
}

func (b *sitebucket) IndexNames() ([]string, query.Error) {
	return []string{b.primary.Name()}, nil
}

func (b *sitebucket) IndexById(name string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *sitebucket) IndexByName(name string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *sitebucket) IndexByPrimary() (catalog.PrimaryIndex, query.Error) {
	return b.primary, nil
}

func (b *sitebucket) Indexes() ([]catalog.Index, query.Error) {
	return []catalog.Index{b.primary}, nil
}

func (b *sitebucket) BulkFetch(ids []string) (map[string]*dparval.Value, query.Error) {
	rv := make(map[string]*dparval.Value, 0)
	for _, id := range ids {
		item, e := b.Fetch(id)
		if e != nil {
			return nil, e
		}
		rv[id] = item
	}
	return rv, nil
}

func (b *sitebucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	if id == b.pool.site.actualSite.Id() {
		doc := map[string]interface{}{
			"id":  b.pool.site.actualSite.Id(),
			"url": b.pool.site.actualSite.URL(),
		}
		return dparval.NewValue(doc), nil
	}
	return nil, query.NewError(nil, "Not Found")
}

func (b *sitebucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
        if b.primary != nil {
	        return b.primary, nil
	}

        return nil, query.NewError(nil, "Not supported.")
}

func (b *sitebucket) CreateIndex(name string, key []ast.Expression, using string) (catalog.Index, query.Error) {
        return nil, query.NewError(nil, "Not supported.")
}

func newSitesBucket(p *pool) (*sitebucket, query.Error) {
	b := new(sitebucket)
	b.pool = p
	b.name = BUCKET_NAME_SITES

	b.primary = &siteIndex{name: "primary", bucket: b}

	return b, nil
}

type siteIndex struct {
	name   string
	bucket *sitebucket
}

func (pi *siteIndex) BucketId() string {
	return pi.name
}

func (pi *siteIndex) Id() string {
	return pi.Name()
}

func (pi *siteIndex) Name() string {
	return pi.name
}

func (pi *siteIndex) Type() string {
	return "primary"
}

func (pi *siteIndex) Key() catalog.IndexKey {
	return catalog.IndexKey{"meta().id"}
}

func (pi *siteIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *siteIndex) ScanEntries(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	doc := dparval.NewValue(map[string]interface{}{})
	doc.SetAttachment("meta", map[string]interface{}{"id": pi.bucket.pool.site.actualSite.Id()})
	ch <- doc

}

// pools bucket
type poolbucket struct {
	pool    *pool
	name    string
	primary catalog.PrimaryIndex
}

func (b *poolbucket) Release() {
}

func (b *poolbucket) PoolId() string {
	return b.pool.Id()
}

func (b *poolbucket) Id() string {
	return b.Name()
}

func (b *poolbucket) Name() string {
	return b.name
}

func (b *poolbucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *poolbucket) IndexIds() ([]string, query.Error) {
	return []string{b.primary.Id()}, nil
}

func (b *poolbucket) IndexNames() ([]string, query.Error) {
	return []string{b.primary.Name()}, nil
}

func (b *poolbucket) IndexById(id string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *poolbucket) IndexByName(name string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *poolbucket) IndexByPrimary() (catalog.PrimaryIndex, query.Error) {
	return b.primary, nil
}

func (b *poolbucket) Indexes() ([]catalog.Index, query.Error) {
	return []catalog.Index{b.primary}, nil
}

func (b *poolbucket) BulkFetch(ids []string) (map[string]*dparval.Value, query.Error) {
	rv := make(map[string]*dparval.Value, 0)
	for _, id := range ids {
		item, e := b.Fetch(id)
		if e != nil {
			return nil, e
		}
		rv[id] = item
	}
	return rv, nil
}

func (b *poolbucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	pool, err := b.pool.site.actualSite.PoolById(id)
	if pool != nil {
		doc := map[string]interface{}{
			"id":      pool.Id(),
			"name":    pool.Name(),
			"site_id": b.pool.site.actualSite.Id(),
		}
		return dparval.NewValue(doc), nil
	}
	return nil, err
}

func (b *poolbucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
        if b.primary != nil {
	        return b.primary, nil
	}

        return nil, query.NewError(nil, "Not supported.")
}

func (b *poolbucket) CreateIndex(name string, key []ast.Expression, using string) (catalog.Index, query.Error) {
        return nil, query.NewError(nil, "Not supported.")
}

func newPoolsBucket(p *pool) (*poolbucket, query.Error) {
	b := new(poolbucket)
	b.pool = p
	b.name = BUCKET_NAME_POOLS

	b.primary = &poolIndex{name: "primary", bucket: b}

	return b, nil
}

type poolIndex struct {
	name   string
	bucket *poolbucket
}

func (pi *poolIndex) BucketId() string {
	return pi.bucket.Id()
}

func (pi *poolIndex) Id() string {
	return pi.Name()
}

func (pi *poolIndex) Name() string {
	return pi.name
}

func (pi *poolIndex) Type() string {
	return "primary"
}

func (pi *poolIndex) Key() catalog.IndexKey {
	return catalog.IndexKey{"meta().id"}
}

func (pi *poolIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *poolIndex) ScanEntries(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	poolIds, err := pi.bucket.pool.site.actualSite.PoolIds()
	if err == nil {
		for _, poolId := range poolIds {
			doc := dparval.NewValue(map[string]interface{}{})
			doc.SetAttachment("meta", map[string]interface{}{"id": poolId})
			ch <- doc
		}
	}
}

// buckets bucket
type bucketbucket struct {
	pool    *pool
	name    string
	primary catalog.PrimaryIndex
}

func (b *bucketbucket) Release() {
}

func (b *bucketbucket) PoolId() string {
	return b.pool.Id()
}

func (b *bucketbucket) Id() string {
	return b.Name()
}

func (b *bucketbucket) Name() string {
	return b.name
}

func (b *bucketbucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *bucketbucket) IndexIds() ([]string, query.Error) {
	return []string{b.primary.Id()}, nil
}

func (b *bucketbucket) IndexNames() ([]string, query.Error) {
	return []string{b.primary.Name()}, nil
}

func (b *bucketbucket) IndexById(id string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *bucketbucket) IndexByName(name string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *bucketbucket) IndexByPrimary() (catalog.PrimaryIndex, query.Error) {
	return b.primary, nil
}

func (b *bucketbucket) Indexes() ([]catalog.Index, query.Error) {
	return []catalog.Index{b.primary}, nil
}

func (b *bucketbucket) BulkFetch(ids []string) (map[string]*dparval.Value, query.Error) {
	rv := make(map[string]*dparval.Value, 0)
	for _, id := range ids {
		item, e := b.Fetch(id)
		if e != nil {
			return nil, e
		}
		rv[id] = item
	}
	return rv, nil
}

func (b *bucketbucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	ids := strings.SplitN(id, "/", 2)

	pool, err := b.pool.site.actualSite.PoolById(ids[0])
	if pool != nil {
		bucket, _ := pool.BucketById(ids[1])
		if bucket != nil {
			doc := map[string]interface{}{
				"id":      bucket.Id(),
				"name":    bucket.Name(),
				"pool_id": pool.Id(),
				"site_id": b.pool.site.actualSite.Id(),
			}
			return dparval.NewValue(doc), nil
		}
	}
	return nil, err
}

func (b *bucketbucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
        if b.primary != nil {
	        return b.primary, nil
	}

        return nil, query.NewError(nil, "Not supported.")
}

func (b *bucketbucket) CreateIndex(name string, key []ast.Expression, using string) (catalog.Index, query.Error) {
        return nil, query.NewError(nil, "Not supported.")
}

func newBucketsBucket(p *pool) (*bucketbucket, query.Error) {
	b := new(bucketbucket)
	b.pool = p
	b.name = BUCKET_NAME_BUCKETS

	b.primary = &bucketIndex{name: "primary", bucket: b}

	return b, nil
}

type bucketIndex struct {
	name   string
	bucket *bucketbucket
}

func (pi *bucketIndex) BucketId() string {
	return pi.bucket.Id()
}

func (pi *bucketIndex) Id() string {
	return pi.Name()
}

func (pi *bucketIndex) Name() string {
	return pi.name
}

func (pi *bucketIndex) Type() string {
	return "primary"
}

func (pi *bucketIndex) Key() catalog.IndexKey {
	return catalog.IndexKey{"meta().id"}
}

func (pi *bucketIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *bucketIndex) ScanEntries(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	poolIds, err := pi.bucket.pool.site.actualSite.PoolIds()
	if err == nil {
		for _, poolId := range poolIds {
			pool, err := pi.bucket.pool.site.actualSite.PoolById(poolId)
			if err == nil {
				bucketIds, err := pool.BucketIds()
				if err == nil {
					for _, bucketId := range bucketIds {
						doc := dparval.NewValue(map[string]interface{}{})
						doc.SetAttachment("meta", map[string]interface{}{"id": fmt.Sprintf("%s/%s", poolId, bucketId)})
						ch <- doc
					}
				}
			}
		}
	}
}
