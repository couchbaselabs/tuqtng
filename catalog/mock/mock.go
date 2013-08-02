//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

/*

Package mock provides a fake, mock 100%-in-memory implementation of
the catalog package, which can be useful for testing.  Because it is
memory-oriented, performance testing of higher layers may be easier
with this mock catalog.

*/
package mock

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/dparval"
)

const (
	DEFAULT_NUM_POOLS   = 1
	DEFAULT_NUM_BUCKETS = 1
	DEFAULT_NUM_ITEMS   = 100000
)

// NewSite creates a new mock site for the given "path".  The path has
// prefix "mock:", with the rest of the path treated as a
// comma-separated key=value params.  For example:
//     mock:pools=2,buckets=5,items=50000
// The above means 2 pools.
// And, each pool has 5 buckets.
// And, each bucket with 50000 items.
// By default, you get...
//     mock:pools=1,buckets=1,items=100000
// Which is what you'd get by specifying a path of just...
//     mock:
func NewSite(path string) (catalog.Site, query.Error) {
	if strings.HasPrefix(path, "mock:") {
		path = path[5:]
	}
	params := map[string]int{}
	for _, kv := range strings.Split(path, ",") {
		if kv == "" {
			continue
		}
		pair := strings.Split(kv, "=")
		v, err := strconv.Atoi(pair[1])
		if err != nil {
			return nil, query.NewError(err,
				fmt.Sprintf("could not parse mock param key: %s, val: %s",
					pair[0], pair[1]))
		}
		params[pair[0]] = v
	}
	npools := paramVal(params, "pools", DEFAULT_NUM_POOLS)
	nbuckets := paramVal(params, "buckets", DEFAULT_NUM_BUCKETS)
	nitems := paramVal(params, "items", DEFAULT_NUM_ITEMS)
	s := &site{path: path, params: params, pools: map[string]*pool{}}
	for i := 0; i < npools; i++ {
		p := &pool{site: s, name: "p" + strconv.Itoa(i), buckets: map[string]*bucket{}}
		for j := 0; j < nbuckets; j++ {
			b := &bucket{pool: p, name: "b" + strconv.Itoa(j), nitems: nitems,
				scanners: map[string]catalog.Scanner{}}
			b.scanners["all_docs"] = &fullScanner{site: s, name: "all_docs", bucket: b}
			p.buckets[b.name] = b
		}
		s.pools[p.name] = p
	}
	return s, nil
}

func paramVal(params map[string]int, key string, defaultVal int) int {
	v, ok := params[key]
	if ok {
		return v
	}
	return defaultVal
}

type site struct {
	path   string
	params map[string]int
	pools  map[string]*pool
}

type pool struct {
	site    *site
	name    string
	buckets map[string]*bucket
}

type bucket struct {
	pool     *pool
	name     string
	nitems   int
	scanners map[string]catalog.Scanner
}

func (s *site) URL() string {
	return "mock:" + s.path
}

func (s *site) PoolNames() ([]string, query.Error) {
	names := make([]string, 0, len(s.pools))
	for name, _ := range s.pools {
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}

func (s *site) Pool(name string) (p catalog.Pool, e query.Error) {
	p, ok := s.pools[name]
	if !ok {
		return nil, query.NewError(nil, "Pool "+name+" not found.")
	}
	return p, nil
}

func (p *pool) Name() string {
	return p.name
}

func (p *pool) BucketNames() ([]string, query.Error) {
	names := make([]string, 0, len(p.buckets))
	for name, _ := range p.buckets {
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}

func (p *pool) Bucket(name string) (b catalog.Bucket, e query.Error) {
	b, ok := p.buckets[name]
	if !ok {
		return nil, query.NewError(nil, "Bucket "+name+" not found.")
	}
	return b, nil
}

func (b *bucket) Release() {
}

func (b *bucket) Name() string {
	return b.name
}

func (b *bucket) Count() (int64, query.Error) {
	return int64(b.nitems), nil
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
	for _, id := range ids {
		item, e := b.Fetch(id)
		if e != nil {
			return nil, e
		}
		rv[id] = item
	}
	return rv, nil
}

func (b *bucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, query.NewError(err,
			fmt.Sprintf("no mock item: %v", id))
	}
	return genItem(i, b.nitems)
}

type fullScanner struct {
	site   *site
	name   string
	bucket *bucket
}

func (fs *fullScanner) Name() string {
	return fs.name
}

func (fs *fullScanner) ScanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	go fs.scanAll(ch, warnch, errch)
}

func (fs *fullScanner) scanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	for i := 0; i < fs.bucket.nitems; i++ {
		doc := dparval.NewValue(map[string]interface{}{})
		doc.SetAttachment("meta", map[string]interface{}{"id": strconv.Itoa(i)})
		ch <- doc
	}
}

func genItem(i int, nitems int) (*dparval.Value, query.Error) {
	if i < 0 || i >= nitems {
		return nil, query.NewError(nil,
			fmt.Sprintf("item out of mock range: %v [0,%v)", i, nitems))
	}
	id := strconv.Itoa(i)
	doc := dparval.NewValue(map[string]interface{}{"id": id, "i": float64(i)})
	doc.SetAttachment("meta", map[string]interface{}{"id": id})
	return doc, nil
}
