//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

/*

Package file provides a file-based implementation of the catalog
package.

*/
package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

// site is the root for the file-based Site.
type site struct {
	path      string
	pools     map[string]*pool
	poolNames []string
}

func (s *site) Id() string {
	return s.path
}

func (s *site) URL() string {
	return "file://" + s.path
}

func (s *site) PoolIds() ([]string, query.Error) {
	return s.PoolNames()
}

func (s *site) PoolNames() ([]string, query.Error) {
	return s.poolNames, nil
}

func (s *site) PoolById(id string) (p catalog.Pool, e query.Error) {
	return s.PoolByName(id)
}

func (s *site) PoolByName(name string) (p catalog.Pool, e query.Error) {
	p, ok := s.pools[strings.ToUpper(name)]
	if !ok {
		e = query.NewError(nil, "Pool "+name+" not found.")
	}

	return
}

// NewSite creates a new file-based site for the given filepath.
func NewSite(path string) (s catalog.Site, e query.Error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, query.NewError(err, "")
	}

	fs := &site{path: path}

	e = fs.loadPools()
	if e != nil {
		return
	}

	s = fs
	return
}

func (s *site) loadPools() (e query.Error) {
	dirEntries, err := ioutil.ReadDir(s.path)
	if err != nil {
		return query.NewError(err, "")
	}

	s.pools = make(map[string]*pool)
	s.poolNames = make([]string, 0)

	var p *pool
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			s.poolNames = append(s.poolNames, dirEntry.Name())
			diru := strings.ToUpper(dirEntry.Name())
			if _, ok := s.pools[diru]; ok {
				return query.NewError(nil, "Duplicate pool name "+dirEntry.Name())
			}

			p, e = newPool(s, dirEntry.Name())
			if e != nil {
				return
			}

			s.pools[diru] = p
		}
	}

	return
}

// pool represents a file-based Pool.
type pool struct {
	site        *site
	name        string
	buckets     map[string]*bucket
	bucketNames []string
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
	return p.bucketNames, nil
}

func (p *pool) BucketById(id string) (b catalog.Bucket, e query.Error) {
	return p.BucketByName(id)
}

func (p *pool) BucketByName(name string) (b catalog.Bucket, e query.Error) {
	b, ok := p.buckets[strings.ToUpper(name)]
	if !ok {
		e = query.NewError(nil, "Bucket "+name+" not found.")
	}

	return
}

func (p *pool) path() string {
	return filepath.Join(p.site.path, p.name)
}

// newPool creates a new pool.
func newPool(s *site, dir string) (p *pool, e query.Error) {
	p = new(pool)
	p.site = s
	p.name = dir

	e = p.loadBuckets()
	return
}

func (p *pool) loadBuckets() (e query.Error) {
	dirEntries, err := ioutil.ReadDir(p.path())
	if err != nil {
		return query.NewError(err, "")
	}

	p.buckets = make(map[string]*bucket)
	p.bucketNames = make([]string, 0)

	var b *bucket
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			diru := strings.ToUpper(dirEntry.Name())
			if _, ok := p.buckets[diru]; ok {
				return query.NewError(nil, "Duplicate bucket name "+dirEntry.Name())
			}

			b, e = newBucket(p, dirEntry.Name())
			if e != nil {
				return
			}

			p.buckets[diru] = b
			p.bucketNames = append(p.bucketNames, b.Name())
		}
	}

	return
}

// bucket is a file-based bucket.
type bucket struct {
	pool    *pool
	name    string
	indexes map[string]catalog.Index
	primary catalog.PrimaryIndex
}

func (b *bucket) Release() {
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
	dirEntries, err := ioutil.ReadDir(b.path())
	if err != nil {
		return 0, query.NewError(err, "")
	}
	return int64(len(dirEntries)), nil
}

func (b *bucket) IndexIds() ([]string, query.Error) {
	rv := make([]string, 0, len(b.indexes))
	for name, _ := range b.indexes {
		rv = append(rv, name)
	}
	return rv, nil
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

func (b *bucket) IndexesByPrimary() ([]catalog.PrimaryIndex, query.Error) {
	return []catalog.PrimaryIndex{b.primary}, nil
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
	for _, id := range ids {
		item, e := b.Fetch(id)
		if e != nil {
			return nil, e
		}
		if item != nil {
			rv[id] = item
		}
	}
	return rv, nil
}

func (b *bucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	path := filepath.Join(b.path(), id+".json")
	item, e = fetch(path)
	if e != nil {
		item = nil
	}

	return
}

func (b *bucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
	if b.primary != nil {
		return b.primary, nil
	}

	return nil, query.NewError(nil, "Not supported.")
}

func (b *bucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {
	return nil, query.NewError(nil, "Not supported.")
}

func (b *bucket) path() string {
	return filepath.Join(b.pool.path(), b.name)
}

// newBucket creates a new bucket.
func newBucket(p *pool, dir string) (b *bucket, e query.Error) {
	b = new(bucket)
	b.pool = p
	b.name = dir

	f, err := os.Open(b.path())
	if err != nil {
		if f != nil {
			f.Close()
		}

		return nil, query.NewError(err, "")
	}

	b.indexes = make(map[string]catalog.Index, 1)
	pi := new(primaryIndex)
	b.primary = pi
	pi.bucket = b
	pi.name = "all_docs"
	b.indexes[pi.name] = pi

	return
}

// primaryIndex performs full bucket scans.
type primaryIndex struct {
	name   string
	bucket *bucket
}

func (pi *primaryIndex) BucketId() string {
	return pi.bucket.Id()
}

func (pi *primaryIndex) Id() string {
	return pi.Name()
}

func (pi *primaryIndex) Name() string {
	return pi.name
}

func (pi *primaryIndex) Type() catalog.IndexType {
	return catalog.UNSPECIFIED
}

func (pi *primaryIndex) IsPrimary() bool {
	return true
}

func (pi *primaryIndex) Key() catalog.IndexKey {
	// FIXME
	return nil
}

func (pi *primaryIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *primaryIndex) ScanBucket(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (pi *primaryIndex) ScanEntries(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	dirEntries, err := ioutil.ReadDir(pi.bucket.path())
	if err != nil {
		errch <- query.NewError(err, "")
		return
	}

	for i, dirEntry := range dirEntries {
		if limit > 0 && int64(i) > limit {
			break
		}
		if !dirEntry.IsDir() {
			entry := catalog.IndexEntry{PrimaryKey: documentPathToId(dirEntry.Name())}
			ch <- &entry
		}
	}
}

func fetch(path string) (item *dparval.Value, e query.Error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// file doesn't exist should simply return nil,nil
			return
		}
		return nil, query.NewError(err, "")
	}

	doc := dparval.NewValueFromBytes(bytes)
	doc.SetAttachment("meta", map[string]interface{}{"id": documentPathToId(path)})
	item = doc

	return
}

func documentPathToId(p string) string {
	_, file := filepath.Split(p)
	ext := filepath.Ext(file)
	return file[0 : len(file)-len(ext)]
}
