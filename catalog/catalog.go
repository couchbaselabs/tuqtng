//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

/*

Package catalog provides a common catalog abstraction over all storage
engines, such as Couchbase server, cloud, mobile, file, 3rd-party
databases and storage engines, etc.

*/
package catalog

import (
	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/query"
)

// log channel for the catalog lifecycle
const CHANNEL = "CATALOG"

// Site represents a cluster or single-node server.
type Site interface {
	Id() string
	URL() string
	PoolIds() ([]string, query.Error)
	PoolNames() ([]string, query.Error)
	PoolById(id string) (Pool, query.Error)
	PoolByName(name string) (Pool, query.Error)
}

// Pool represents a logical authentication, query, and resource
// allocation boundary, as well as a grouping of buckets.
type Pool interface {
	SiteId() string
	Id() string
	Name() string
	BucketIds() ([]string, query.Error)
	BucketNames() ([]string, query.Error)
	BucketById(name string) (Bucket, query.Error)
	BucketByName(name string) (Bucket, query.Error)
}

// Bucket is a collection of key-value entries (typically
// key-document, but not always).
type Bucket interface {
	PoolId() string
	Id() string
	Name() string
	Count() (int64, query.Error)
	IndexIds() ([]string, query.Error)
	IndexNames() ([]string, query.Error)
	IndexById(id string) (Index, query.Error)
	IndexByName(name string) (Index, query.Error)
	IndexByPrimary() (PrimaryIndex, query.Error)     // Returns the server-recommended primary index
	IndexesByPrimary() ([]PrimaryIndex, query.Error) // Returns all available primary indexes
	Indexes() ([]Index, query.Error)
	Fetch(id string) (*dparval.Value, query.Error)
	BulkFetch([]string) (map[string]*dparval.Value, query.Error)
	Release()
	CreatePrimaryIndex() (PrimaryIndex, query.Error)
	CreateIndex(name string, key IndexKey, using IndexType) (Index, query.Error)
}

type IndexType string

const (
	UNSPECIFIED IndexType = "unspecified" // used by non-view primary_indexes
	VIEW        IndexType = "view"
	LSM         IndexType = "llrb"
)

type IndexKey []ast.Expression

// Index is the base type for all indexes.
type Index interface {
	BucketId() string
	Id() string
	Name() string
	Type() IndexType
	IsPrimary() bool
	Key() IndexKey
	Drop() query.Error // PrimaryIndexes cannot be dropped
}

type LookupValue []*dparval.Value

type IndexEntry struct {
	EntryKey   LookupValue
	PrimaryKey string
}

type EntryChannel chan *IndexEntry

// ScanIndex represents scanning indexes.
type ScanIndex interface {
	ScanEntries(limit int64, ch EntryChannel, warnch, errch query.ErrorChannel)
}

// CountIndex represents indexes that can efficiently produce entry counts.
type CountIndex interface {
	Index
	ValueCount() (int64, query.Error)
}

// PrimaryIndex represents primary key indexes.
type PrimaryIndex interface {
	Index
	ScanIndex
	ScanBucket(limit int64, ch EntryChannel, warnch, errch query.ErrorChannel)
}

// ExistenceIndex represents existence indexes.
type ExistenceIndex interface {
	Check(value LookupValue, ch EntryChannel, warnch, errch query.ErrorChannel)
}

// LookupIndex represents lookup indexes.
type LookupIndex interface {
	ExistenceIndex
	Lookup(value LookupValue, ch EntryChannel, warnch, errch query.ErrorChannel)
}

// Direction represents ASC and DESC
// TODO: Is this needed?
type Direction int

const (
	ASC  Direction = 1
	DESC           = 2
)

// Inclusion controls how the boundaries values of a range are treated
type RangeInclusion int

const (
	Neither RangeInclusion = iota
	Low
	High
	Both
)

// RangeIndex represents range scan indexes.
type RangeIndex interface {
	Index
	ScanIndex
	LookupIndex
	Direction() Direction
	Statistics() (RangeStatistics, query.Error)
	ScanRange(low LookupValue, high LookupValue, inclusion RangeInclusion, limit int64, ch EntryChannel, warnch, errch query.ErrorChannel)
}

// SearchIndex represents full text search indexes.
type SearchIndex interface {
	Index
	SearchValue(value string, limit int64, ch EntryChannel, warnch, errch query.ErrorChannel)
}

// RangeStatistics captures statistics for a range index.
type RangeStatistics interface {
	Count() (int64, query.Error)
	Min() (dparval.Value, query.Error)
	Max() (dparval.Value, query.Error)
	DistinctCount(int64, query.Error)
	Bins() ([]Bin, query.Error)
}

// Bin represents a range bin within IndexStatistics.
type Bin interface {
	Count() (int64, query.Error)
	Min() (dparval.Value, query.Error)
	Max() (dparval.Value, query.Error)
	DistinctCount(int64, query.Error)
}
