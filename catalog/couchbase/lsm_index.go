package couchbase

import (
	"bytes"
	"fmt"
	//	"github.com/couchbaselabs/clog"
	//	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type lsmIndex struct {
	name   string
	uuid   string
	using  catalog.IndexType
	on     catalog.IndexKey
	bucket catalog.Bucket
	//	createStmt string
}

type primaryLsmIndex struct {
	lsmIndex
}

func (li *lsmIndex) BucketId() string {
	return li.bucket.Id()
}

func (li *lsmIndex) Id() string {
	return li.uuid
}

func (li *lsmIndex) Name() string {
	return li.name
}

func (li *lsmIndex) Type() catalog.IndexType {
	return li.using
}

func (li *lsmIndex) IsPrimary() bool {
	return false
}

func (li *lsmIndex) Key() catalog.IndexKey {
	return li.on
}

func (li *lsmIndex) Drop() query.Error {
	bucket := li.bucket.(*bucket)
	if li.IsPrimary() {
		return query.NewError(nil, "Primary index cannot be dropped.")
	}
	err := li.DropLsmIndex()
	if err != nil {
		return query.NewError(err, fmt.Sprintf("Cannot drop index %s", li.Name()))
	}
	delete(bucket.indexes, li.name)
	return nil
}

func (li *lsmIndex) Check(value catalog.LookupValue, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	li.ScanRange(value, value, catalog.Both, 0, ch, warnch, errch)
}

func (li *lsmIndex) Lookup(value catalog.LookupValue, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	li.ScanRange(value, value, catalog.Both, 0, ch, warnch, errch)
}

func (li *lsmIndex) Statistics() (catalog.RangeStatistics, query.Error) {
	return nil, query.NewError(nil, "statistics not implemented")
}

func (li *lsmIndex) Direction() catalog.Direction {
	return catalog.ASC
}

func (li *lsmIndex) ScanRange(low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion, limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	//	TODO
}

func (li *lsmIndex) ScanEntries(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	li.ScanRange(nil, nil, catalog.Both, limit, ch, warnch, errch)
}

func (li *lsmIndex) ValueCount() (int64, query.Error) {
	//	TODO
	return 0, nil
}

func (pi *primaryLsmIndex) IsPrimary() bool {
	return true
}

func (pi *primaryLsmIndex) ScanBucket(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	pi.ScanEntries(limit, ch, warnch, errch)
}

func (li *lsmIndex) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("name: %v ", li.name))
	buf.WriteString(fmt.Sprintf("uuid: %v ", li.uuid))
	buf.WriteString(fmt.Sprintf("using: %v ", li.using))
	buf.WriteString(fmt.Sprintf("on: %v ", li.on))
	buf.WriteString(fmt.Sprintf("bucket: %v", li.bucket.Name()))
	return buf.String()
}
