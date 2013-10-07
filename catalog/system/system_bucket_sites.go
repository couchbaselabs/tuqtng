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
	return 1, nil
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

func (b *sitebucket) IndexesByPrimary() ([]catalog.PrimaryIndex, query.Error) {
	return []catalog.PrimaryIndex{b.primary}, nil
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

func (b *sitebucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {
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

func (pi *siteIndex) Type() catalog.IndexType {
	return catalog.UNSPECIFIED
}

func (pi *siteIndex) IsPrimary() bool {
	return true
}

func (pi *siteIndex) Key() catalog.IndexKey {
	// FIXME
	return nil
}

func (pi *siteIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *siteIndex) ScanBucket(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (pi *siteIndex) ScanEntries(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	entry := catalog.IndexEntry{PrimaryKey: pi.bucket.pool.site.actualSite.Id()}
	ch <- &entry
}
