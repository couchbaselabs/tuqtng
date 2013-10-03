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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type indexbucket struct {
	pool    *pool
	name    string
	primary catalog.PrimaryIndex
}

func (b *indexbucket) Release() {
}

func (b *indexbucket) PoolId() string {
	return b.pool.Id()
}

func (b *indexbucket) Id() string {
	return b.Name()
}

func (b *indexbucket) Name() string {
	return b.name
}

func (b *indexbucket) Count() (int64, query.Error) {
	count := int64(0)
	poolIds, err := b.pool.site.actualSite.PoolIds()
	if err == nil {
		for _, poolId := range poolIds {
			pool, err := b.pool.site.actualSite.PoolById(poolId)
			if err == nil {
				bucketIds, err := pool.BucketIds()
				if err == nil {
					for _, bucketId := range bucketIds {
						bucket, err := pool.BucketById(bucketId)
						if err == nil {
							indexIds, err := bucket.IndexIds()
							if err == nil {
								count += int64(len(indexIds))
							} else {
								return 0, query.NewError(err, "")
							}
						} else {
							return 0, query.NewError(err, "")
						}
					}
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

func (b *indexbucket) IndexIds() ([]string, query.Error) {
	return []string{b.primary.Id()}, nil
}

func (b *indexbucket) IndexNames() ([]string, query.Error) {
	return []string{b.primary.Name()}, nil
}

func (b *indexbucket) IndexById(id string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *indexbucket) IndexByName(name string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *indexbucket) IndexByPrimary() (catalog.PrimaryIndex, query.Error) {
	return b.primary, nil
}

func (b *indexbucket) IndexesByPrimary() ([]catalog.PrimaryIndex, query.Error) {
	return []catalog.PrimaryIndex{b.primary}, nil
}

func (b *indexbucket) Indexes() ([]catalog.Index, query.Error) {
	return []catalog.Index{b.primary}, nil
}

func (b *indexbucket) BulkFetch(ids []string) (map[string]*dparval.Value, query.Error) {
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

func (b *indexbucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	ids := strings.SplitN(id, "/", 3)

	pool, err := b.pool.site.actualSite.PoolById(ids[0])
	if pool != nil {
		bucket, _ := pool.BucketById(ids[1])
		if bucket != nil {
			index, _ := bucket.IndexById(ids[2])
			if index != nil {
				doc := map[string]interface{}{
					"id":         index.Id(),
					"name":       index.Name(),
					"bucket_id":  bucket.Id(),
					"pool_id":    pool.Id(),
					"site_id":    b.pool.site.actualSite.Id(),
					"index_key":  catalogObjectToJSONSafe(indexKeyToIndexKeyStringArray(index.Key())),
					"index_type": catalogObjectToJSONSafe(index.Type()),
				}
				return dparval.NewValue(doc), nil
			}
		}
	}
	return nil, err
}

func indexKeyToIndexKeyStringArray(key catalog.IndexKey) []string {
	rv := make([]string, len(key))
	for i, kp := range key {
		rv[i] = kp.String()
	}
	return rv
}

func catalogObjectToJSONSafe(catobj interface{}) interface{} {
	var rv interface{}
	bytes, err := json.Marshal(catobj)
	if err == nil {
		json.Unmarshal(bytes, &rv)
	}
	return rv
}

func (b *indexbucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
	if b.primary != nil {
		return b.primary, nil
	}

	return nil, query.NewError(nil, "Not supported.")
}

func (b *indexbucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {
	return nil, query.NewError(nil, "Not supported.")
}

func newIndexesBucket(p *pool) (*indexbucket, query.Error) {
	b := new(indexbucket)
	b.pool = p
	b.name = BUCKET_NAME_INDEXES

	b.primary = &indexIndex{name: "primary", bucket: b}

	return b, nil
}

type indexIndex struct {
	name   string
	bucket *indexbucket
}

func (pi *indexIndex) BucketId() string {
	return pi.bucket.Id()
}

func (pi *indexIndex) Id() string {
	return pi.Name()
}

func (pi *indexIndex) Name() string {
	return pi.name
}

func (pi *indexIndex) Type() catalog.IndexType {
	return catalog.UNSPECIFIED
}

func (pi *indexIndex) IsPrimary() bool {
	return true
}

func (pi *indexIndex) Key() catalog.IndexKey {
	// FIXME
	return nil
}

func (pi *indexIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *indexIndex) ScanBucket(limit int64, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (pi *indexIndex) ScanEntries(limit int64, ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
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
						bucket, err := pool.BucketById(bucketId)
						if err == nil {
							indexIds, err := bucket.IndexIds()
							if err == nil {
								for i, indexId := range indexIds {
									if limit > 0 && int64(i) > limit {
										break
									}
									doc := dparval.NewValue(map[string]interface{}{})
									doc.SetAttachment("meta", map[string]interface{}{"id": fmt.Sprintf("%s/%s/%s", poolId, bucketId, indexId)})
									ch <- doc
								}
							}
						}
					}
				}
			}
		}
	}
}
