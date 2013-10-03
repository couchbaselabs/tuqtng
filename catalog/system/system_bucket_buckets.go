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
	count := int64(0)
	poolIds, err := b.pool.site.actualSite.PoolIds()
	if err == nil {
		for _, poolId := range poolIds {
			pool, err := b.pool.site.actualSite.PoolById(poolId)
			if err == nil {
				bucketIds, err := pool.BucketIds()
				if err == nil {
					count += int64(len(bucketIds))
				} else {
					return 0, query.NewError(err, "")
				}
			} else {
				return 0, query.NewError(err, "")
			}
		}
		return count, nil
	}
	return 0, query.NewError(err, "")
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

func (b *bucketbucket) IndexesByPrimary() ([]catalog.PrimaryIndex, query.Error) {
	return []catalog.PrimaryIndex{b.primary}, nil
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

func (b *bucketbucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {
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

func (pi *bucketIndex) Type() catalog.IndexType {
	return catalog.UNSPECIFIED
}

func (pi *bucketIndex) IsPrimary() bool {
	return true
}

func (pi *bucketIndex) Key() catalog.IndexKey {
	// FIXME
	return nil
}

func (pi *bucketIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *bucketIndex) ScanBucket(limit int64, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (pi *bucketIndex) ScanEntries(limit int64, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
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
					for i, bucketId := range bucketIds {
						if limit > 0 && int64(i) > limit {
							break
						}
						doc := dparval.NewValue(map[string]interface{}{})
						doc.SetAttachment("meta", map[string]interface{}{"id": fmt.Sprintf("%s/%s", poolId, bucketId)})
						ch <- doc
					}
				}
			}
		}
	}
}
