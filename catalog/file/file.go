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
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
	"path/filepath"
	"strings"
)

// site is the root for the file-based Site.
type site struct {
	path      string
	pools     map[string]*pool
	poolNames []string
}

func (s *site) URL() string {
	return "file://" + s.path
}

func (s *site) PoolNames() ([]string, query.Error) {
	return s.poolNames, nil
}

func (s *site) Pool(name string) (p Pool, e query.Error) {
	p, ok := s.pools[strings.ToUpper(name)]
	if !ok {
		e = query.NewError("Pool " + name + " not found.")
	}
}

// NewSite creates a new file-based site for the given filepath.
func NewSite(path string) (s Site, e query.Error) {
	s = new(site)
	s.path = path

	e = s.loadPools()
	if e != nil {
		s = nil
	}
}

func (s *site) loadPools() (e query.Error) {
	f, err := os.Open(s.path)
	if err != nil {
		if f != nil {
			f.Close()
		}

		return query.NewError(err)
	}

	dirs, err := f.Readdirnames(0)
	f.Close()

	if err != nil {
		return query.NewError(err)
	}

	s.pools = make(map[string]*pool, len(dirs))
	s.poolNames = dirs

	for _, dir := range dirs {
		diru := strings.ToUpper(dir)
		if _, ok := s.pools[diru]; ok {
			return query.NewError("Duplicate pool name " + dir)
		}

		pool, e := newPool(s, dir)
		if e != nil {
			return
		}

		s.pools[diru] = pool
	}
}

// pool represents a file-based Pool.
type pool struct {
	site        *site
	name        string
	buckets     map[string]*bucket
	bucketNames []string
}

func (p *pool) Name() string {
	return p.name
}

func (p *pool) BucketNames() ([]string, query.Error) {
	return p.bucketNames, nil
}

func (p *pool) Bucket(name string) (b Bucket, e query.Error) {
	b, ok := p.buckets[strings.ToUpper(name)]
	if !ok {
		e = query.NewError("Bucket " + name + " not found.")
	}
}

func (p *pool) path() string {
	return filepath.Join(p.site.path, p.name)
}

// newPool creates a new pool.
func newPool(s *site, dir string) (p pool, e query.Error) {
	p = new(pool)
	p.site = s
	p.name = dir

	e = p.loadBuckets()
}

func (p *pool) loadBuckets() (e query.Error) {
	f, err := os.Open(p.path())
	if err != nil {
		if f != nil {
			f.Close()
		}

		return query.NewError(err)
	}

	dirs, err := f.Readdirnames(0)
	f.Close()

	if err != nil {
		return query.NewError(err)
	}

	p.buckets = make(map[string]*pool, len(dirs))
	s.bucketNames = dirs

	for _, dir := range dirs {
		diru := strings.ToUpper(dir)
		if _, ok := p.buckets[diru]; ok {
			return query.NewError("Duplicate bucket name " + dir)
		}

		bucket, e := newBucket(p, dir)
		if e != nil {
			return
		}

		p.buckets[diru] = bucket
	}
}

// bucket is a file-based bucket.
type bucket struct {
	pool      *pool
	name      string
	filenames []string
	scanners  []Scanner
}

func (b *bucket) Name() string {
	return b.name
}

func (b *bucket) Count() (int64, query.Error) {
	return len(b.filenames), nil
}

func (b *bucket) Scanners() ([]Scanner, query.Error) {
	return b.scanners, nil
}

func (b *bucket) Fetch(id string) (item query.Item, e query.Error) {
	path := filepath.Join(b.path(), id)
	fi, err := os.Stat(path)
	if err != nil {
		return nil, query.NewError(err)
	}

	f, err := os.Open(path)
	if err != nil {
		if f != nil {
			f.Close()
		}

		return nil, query.NewError(err)
	}

	bytes := make([]byte, fi.Size())
	n, err := f.Read(bytes)
	f.Close()
	if err != nil {
		return nil, query.NewError(err)
	}
}

func (b *bucket) path() string {
	return filepath.Join(b.pool.path(), b.name)
}

// newBucket creates a new bucket.
func newBucket(p *pool, dir string) (b bucket, e query.Error) {
	b = new(bucket)
	b.pool = p
	b.name = dir

	f, err := os.Open(b.path())
	if err != nil {
		if f != nil {
			f.Close()
		}

		return nil, query.NewError(err)
	}

	b.filenames, err = f.Readdirnames()
	f.Close()
	if err != nil {
		return nil, query.NewError(err)
	}

	b.scanners = make([]Scanner, 1)
	fs := new(fullScanner)
	fs.bucket = b
	b.scanners.append(fs)
}

// fullScanner performs full bucket scans.
type fullScanner struct {
	bucket *bucket
}

func (fs *fullScanner) ScanAll(ch ItemChannel) query.Error {
}
