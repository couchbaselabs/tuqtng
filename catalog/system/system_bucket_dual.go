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

type dualbucket struct {
	pool    *pool
	name    string
	primary catalog.PrimaryIndex
}

func (b *dualbucket) Release() {
}

func (b *dualbucket) PoolId() string {
	return b.pool.Id()
}

func (b *dualbucket) Id() string {
	return b.Name()
}

func (b *dualbucket) Name() string {
	return b.name
}

func (b *dualbucket) Count() (int64, query.Error) {
	return 0, query.NewError(nil, "Not Supported")
}

func (b *dualbucket) IndexIds() ([]string, query.Error) {
	return []string{b.primary.Id()}, nil
}

func (b *dualbucket) IndexNames() ([]string, query.Error) {
	return []string{b.primary.Name()}, nil
}

func (b *dualbucket) IndexById(id string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *dualbucket) IndexByName(name string) (catalog.Index, query.Error) {
	return b.primary, nil
}

func (b *dualbucket) IndexByPrimary() (catalog.PrimaryIndex, query.Error) {
	return b.primary, nil
}

func (b *dualbucket) Indexes() ([]catalog.Index, query.Error) {
	return []catalog.Index{b.primary}, nil
}

func (b *dualbucket) BulkFetch(ids []string) (map[string]*dparval.Value, query.Error) {
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

func (b *dualbucket) Fetch(id string) (item *dparval.Value, e query.Error) {
	doc := map[string]interface{}{}
	return dparval.NewValue(doc), nil
}

func (b *dualbucket) CreatePrimaryIndex() (catalog.PrimaryIndex, query.Error) {
	if b.primary != nil {
		return b.primary, nil
	}

	return nil, query.NewError(nil, "Not supported.")
}

func (b *dualbucket) CreateIndex(name string, key catalog.IndexKey, using catalog.IndexType) (catalog.Index, query.Error) {
	return nil, query.NewError(nil, "Not supported.")
}

func newDualBucket(p *pool) (*dualbucket, query.Error) {
	b := new(dualbucket)
	b.pool = p
	b.name = BUCKET_NAME_DUAL

	b.primary = &dualIndex{name: "primary", bucket: b}

	return b, nil
}

type dualIndex struct {
	name   string
	bucket *dualbucket
}

func (pi *dualIndex) BucketId() string {
	return pi.bucket.Id()
}

func (pi *dualIndex) Id() string {
	return pi.Name()
}

func (pi *dualIndex) Name() string {
	return pi.name
}

func (pi *dualIndex) Type() catalog.IndexType {
	return catalog.UNSPECIFIED
}

func (pi *dualIndex) IsPrimary() bool {
	return true
}

func (pi *dualIndex) Key() catalog.IndexKey {
	// FIXME
	return nil
}

func (pi *dualIndex) Drop() query.Error {
	return query.NewError(nil, "Primary index cannot be dropped.")
}

func (pi *dualIndex) ScanBucket(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(ch, warnch, errch)
}

func (pi *dualIndex) ScanEntries(ch dparval.ValueChannel, warnch, errch query.ErrorChannel) {
	defer close(ch)
	defer close(warnch)
	defer close(errch)

	doc := dparval.NewValue(map[string]interface{}{})
	doc.SetAttachment("meta", map[string]interface{}{"id": "expression_evaluator"})
	ch <- doc
}
