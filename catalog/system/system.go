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
)

const POOL_NAME = "sys_catalog"
const BUCKET_NAME_SITES = "sites"
const BUCKET_NAME_POOLS = "pools"
const BUCKET_NAME_BUCKETS = "buckets"
const BUCKET_NAME_INDEXES = "indexes"

type site struct {
	actualSite        catalog.Site
	systemCatalogPool *pool
}

func (s *site) URL() string {
	return s.actualSite.URL()
}

func (s *site) PoolNames() ([]string, query.Error) {
	poolNames, err := s.actualSite.PoolNames()
	if err != nil {
		return nil, err
	}
	poolNames = append(poolNames, s.systemCatalogPool.Name())
	return poolNames, err
}

func (s *site) Pool(name string) (catalog.Pool, query.Error) {
	if name == POOL_NAME {
		return s.systemCatalogPool, nil
	}
	return s.actualSite.Pool(name)
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
	name    string
	buckets map[string]catalog.Bucket
}

func (p *pool) Name() string {
	return p.name
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

func (p *pool) Bucket(name string) (catalog.Bucket, query.Error) {
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
	scanner catalog.Scanner
}

func (b *sitebucket) Release() {
}

func (b *sitebucket) Name() string {
	return b.name
}

func (b *sitebucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *sitebucket) ScannerNames() ([]string, query.Error) {
	return []string{b.scanner.Name()}, nil
}

func (b *sitebucket) Scanners() ([]catalog.Scanner, query.Error) {
	return []catalog.Scanner{b.scanner}, nil
}

func (b *sitebucket) Scanner(name string) (catalog.Scanner, query.Error) {
	return b.scanner, nil
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
	if id == b.pool.site.actualSite.URL() {
		doc := map[string]interface{}{
			"id":  b.pool.site.actualSite.URL(),
			"url": b.pool.site.actualSite.URL(),
		}
		return dparval.NewValue(doc), nil
	}
	return nil, query.NewError(nil, "Not Found")
}

func newSitesBucket(p *pool) (*sitebucket, query.Error) {
	b := new(sitebucket)
	b.pool = p
	b.name = BUCKET_NAME_SITES

	b.scanner = &siteScanner{name: "all", bucket: b}

	return b, nil
}

type siteScanner struct {
	name   string
	bucket *sitebucket
}

func (fs *siteScanner) Name() string {
	return fs.name
}

func (fs *siteScanner) ScanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	doc := dparval.NewValue(map[string]interface{}{})
	doc.SetAttachment("meta", map[string]interface{}{"id": fs.bucket.pool.site.actualSite.URL()})
	ch <- doc

}

// pools bucket
type poolbucket struct {
	pool    *pool
	name    string
	scanner catalog.Scanner
}

func (b *poolbucket) Release() {
}

func (b *poolbucket) Name() string {
	return b.name
}

func (b *poolbucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *poolbucket) ScannerNames() ([]string, query.Error) {
	return []string{b.scanner.Name()}, nil
}

func (b *poolbucket) Scanners() ([]catalog.Scanner, query.Error) {
	return []catalog.Scanner{b.scanner}, nil
}

func (b *poolbucket) Scanner(name string) (catalog.Scanner, query.Error) {
	return b.scanner, nil
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
	pool, err := b.pool.site.actualSite.Pool(id)
	if pool != nil {
		doc := map[string]interface{}{
			"id":      pool.Name(),
			"name":    pool.Name(),
			"site_id": b.pool.site.actualSite.URL(),
		}
		return dparval.NewValue(doc), nil
	}
	return nil, err
}

func newPoolsBucket(p *pool) (*poolbucket, query.Error) {
	b := new(poolbucket)
	b.pool = p
	b.name = BUCKET_NAME_POOLS

	b.scanner = &poolScanner{name: "all", bucket: b}

	return b, nil
}

type poolScanner struct {
	name   string
	bucket *poolbucket
}

func (ps *poolScanner) Name() string {
	return ps.name
}

func (ps *poolScanner) ScanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	poolNames, err := ps.bucket.pool.site.actualSite.PoolNames()
	if err == nil {
		for _, poolName := range poolNames {
			doc := dparval.NewValue(map[string]interface{}{})
			doc.SetAttachment("meta", map[string]interface{}{"id": poolName})
			ch <- doc
		}
	}
}

// buckets bucket
type bucketbucket struct {
	pool    *pool
	name    string
	scanner catalog.Scanner
}

func (b *bucketbucket) Release() {
}

func (b *bucketbucket) Name() string {
	return b.name
}

func (b *bucketbucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *bucketbucket) ScannerNames() ([]string, query.Error) {
	return []string{b.scanner.Name()}, nil
}

func (b *bucketbucket) Scanners() ([]catalog.Scanner, query.Error) {
	return []catalog.Scanner{b.scanner}, nil
}

func (b *bucketbucket) Scanner(name string) (catalog.Scanner, query.Error) {
	return b.scanner, nil
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
	names := strings.SplitN(id, "/", 2)

	pool, err := b.pool.site.actualSite.Pool(names[0])
	if pool != nil {
		bucket, _ := pool.Bucket(names[1])
		if bucket != nil {
			doc := map[string]interface{}{
				"id":      bucket.Name(),
				"name":    bucket.Name(),
				"pool_id": pool.Name(),
				"site_id": b.pool.site.actualSite.URL(),
			}
			return dparval.NewValue(doc), nil
		}
	}
	return nil, err
}

func newBucketsBucket(p *pool) (*bucketbucket, query.Error) {
	b := new(bucketbucket)
	b.pool = p
	b.name = BUCKET_NAME_BUCKETS

	b.scanner = &bucketScanner{name: "all", bucket: b}

	return b, nil
}

type bucketScanner struct {
	name   string
	bucket *bucketbucket
}

func (bs *bucketScanner) Name() string {
	return bs.name
}

func (bs *bucketScanner) ScanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	poolNames, err := bs.bucket.pool.site.actualSite.PoolNames()
	if err == nil {
		for _, poolName := range poolNames {
			pool, err := bs.bucket.pool.site.actualSite.Pool(poolName)
			if err == nil {
				bucketNames, err := pool.BucketNames()
				if err == nil {
					for _, bucketName := range bucketNames {
						doc := dparval.NewValue(map[string]interface{}{})
						doc.SetAttachment("meta", map[string]interface{}{"id": fmt.Sprintf("%s/%s", poolName, bucketName)})
						ch <- doc
					}
				}
			}
		}
	}
}
