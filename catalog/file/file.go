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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
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

func (s *site) Pool(name string) (p catalog.Pool, e query.Error) {
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

func (p *pool) Name() string {
	return p.name
}

func (p *pool) BucketNames() ([]string, query.Error) {
	return p.bucketNames, nil
}

func (p *pool) Bucket(name string) (b catalog.Bucket, e query.Error) {
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
		}
	}

	return
}

// bucket is a file-based bucket.
type bucket struct {
	pool     *pool
	name     string
	scanners map[string]catalog.Scanner
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

func (b *bucket) Fetch(id string) (item query.Item, e query.Error) {
	path := filepath.Join(b.path(), id+".json")
	item, e = fetch(path)
	if e != nil {
		item = nil
	}

	return
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

	b.scanners = make(map[string]catalog.Scanner, 1)
	fs := new(fullScanner)
	fs.bucket = b
	fs.name = "all_docs"
	b.scanners[fs.name] = fs

	return
}

// fullScanner performs full bucket scans.
type fullScanner struct {
	name   string
	bucket *bucket
}

func (fs *fullScanner) Name() string {
	return fs.name
}

func (fs *fullScanner) ScanAll(ch query.ItemChannel, warnch, errch query.ErrorChannel) {
	go fs.scanAll(ch, warnch, errch)
}

func (fs *fullScanner) scanAll(ch query.ItemChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	dirEntries, err := ioutil.ReadDir(fs.bucket.path())
	if err != nil {
		errch <- query.NewError(err, "")
		return
	}

	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			doc := map[string]query.Value{}
			meta := map[string]query.Value{"id": documentPathToId(dirEntry.Name())}
			ch <- query.NewMapItem(doc, meta)
		}
	}
}

func fetch(path string) (item query.Item, e query.Error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, query.NewError(err, "")
	}

	// convert file bytes to json
	doc := map[string]query.Value{}
	err = json.Unmarshal(bytes, &doc)
	if err != nil {
		return nil, query.NewError(err, "")
	}

	meta := map[string]query.Value{
		"id": documentPathToId(path),
	}

	item = query.NewMapItem(doc, meta)

	return
}

func documentPathToId(p string) string {
	_, file := path.Split(p)
	ext := path.Ext(file)
	return file[0 : len(file)-len(ext)]
}
