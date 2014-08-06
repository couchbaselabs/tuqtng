package couchbase

import (
	"bytes"
	"fmt"
	"github.com/couchbaselabs/clog"
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/catalog"
	"github.com/couchbaselabs/tuqtng/query"
)

type lsmIndex struct {
	name   string
	uuid   string
	using  catalog.IndexType
	on     catalog.IndexKey
	bucket catalog.Bucket
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
	//FIXME This should use EXISTS. But that's possible only once Check method return
	//bool rather than a channel
	li.executeScanQuery(LOOKUP, value, value, catalog.Both, 0, ch, warnch, errch)
}

func (li *lsmIndex) Lookup(value catalog.LookupValue, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	li.executeScanQuery(LOOKUP, value, value, catalog.Both, 0, ch, warnch, errch)
}

func (li *lsmIndex) Statistics() (catalog.RangeStatistics, query.Error) {
	return nil, query.NewError(nil, "statistics not implemented")
}

func (li *lsmIndex) Direction() catalog.Direction {
	return catalog.ASC
}

func (li *lsmIndex) ScanRange(low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion, limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {

	li.executeScanQuery(RANGESCAN, low, high, inclusion, limit, ch, warnch, errch)
}

func (li *lsmIndex) executeScanQuery(scanType ScanType, low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion, limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {

	defer close(ch)
	defer close(warnch)
	defer close(errch)

	var q QueryParams

	//FIXME convert to a function
	lv := make([][]byte, len(low))
	for i, v := range low {
		lv[i] = v.Bytes()
	}

	hv := make([][]byte, len(high))
	for i, v := range high {
		hv[i] = v.Bytes()
	}

	q.Low = lv
	q.High = hv

	q.Inclusion = inclusion
	q.Limit = limit
	q.ScanType = scanType

	//	nodes, _ := client.Nodes()
	//FIXME choose a random node
	//indexerClient := NewRestClient(nodes[0].IndexerURL)
	clog.To(catalog.CHANNEL, "Requesting Scan with Params %v", q)
	//FIXME Send multiple scan requests in batches
	rows, _, err := client.Scan(li.Id(), q)

	if err != nil {
		clog.To(catalog.CHANNEL, "Error in Scan response %v", err)
		errch <- query.NewError(err, "Error in Scan response")
		return
	}

	for _, row := range rows {
		clog.To(catalog.CHANNEL, "Received Row %v", row)

		entry := catalog.IndexEntry{PrimaryKey: row.Value}
		for _, key := range row.Key {
			val := dparval.NewValueFromBytes(key)
			entry.EntryKey = append(entry.EntryKey, val)
		}
		ch <- &entry
	}
}

func (li *lsmIndex) executeCountQuery(scanType ScanType, low catalog.LookupValue, high catalog.LookupValue, inclusion catalog.RangeInclusion, limit int64) (uint64, query.Error) {

	var q QueryParams

	//FIXME convert to a function
	lv := make([][]byte, len(low))
	for i, v := range low {
		lv[i] = v.Bytes()
	}

	hv := make([][]byte, len(high))
	for i, v := range high {
		hv[i] = v.Bytes()
	}

	q.Low = lv
	q.High = hv

	q.Inclusion = inclusion
	q.Limit = limit
	q.ScanType = scanType

	nodes, _ := client.Nodes()
	//FIXME choose a random node
	indexerClient := NewRestClient(nodes[0].IndexerURL)

	clog.To(catalog.CHANNEL, "Requesting Count with Params %v", q)
	//FIXME Send multiple scan requests in batches
	_, totalRows, err := indexerClient.Scan(li.Id(), q)
	return totalRows, query.NewError(err, "Error In Scan Response")
}

func (li *lsmIndex) ScanEntries(limit int64, ch catalog.EntryChannel, warnch, errch query.ErrorChannel) {
	li.executeScanQuery(FULLSCAN, nil, nil, catalog.Both, limit, ch, warnch, errch)
}

func (li *lsmIndex) ValueCount() (int64, query.Error) {
	totalRows, err := li.executeCountQuery(COUNT, nil, nil, catalog.Both, 0)
	return int64(totalRows), err
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
