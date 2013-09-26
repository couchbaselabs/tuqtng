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
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

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

func (b *poolbucket) IndexesByPrimary() ([]catalog.PrimaryIndex, query.Error) {
	return []catalog.PrimaryIndex{b.primary}, nil
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

func (b *poolbucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {
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

func (pi *poolIndex) Type() catalog.IndexType {
	return catalog.UNSPECIFIED
}

func (pi *poolIndex) IsPrimary() bool {
	return true
}

func (pi *poolIndex) Key() catalog.IndexKey {
	// FIXME
	return nil
}

func (pi *poolIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *poolIndex) ScanBucket(limit int64, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (pi *poolIndex) ScanEntries(limit int64, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	poolIds, err := pi.bucket.pool.site.actualSite.PoolIds()
	if err == nil {
		for i, poolId := range poolIds {
			if limit > 0 && int64(i) > limit {
				break
			}
			doc := dparval.NewValue(map[string]interface{}{})
			doc.SetAttachment("meta", map[string]interface{}{"id": poolId})
			ch <- doc
		}
	}
}
