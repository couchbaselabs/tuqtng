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
engines, such as Couchbase server, file, mobile, CBGB, 3rd-party
storage engines, etc.

*/
package catalog

import (
	"github.com/couchbaselabs/tuqtng/query"
	"github.com/couchbaselabs/dparval"
)

// Site represents a cluster or single-node server.
type Site interface {
	URL() string
	PoolNames() ([]string, query.Error)
	Pool(name string) (Pool, query.Error)
}

// Pool represents a logical authentication, query, and resource
// allocation boundary, as well as a grouping of buckets.
type Pool interface {
	Name() string
	BucketNames() ([]string, query.Error)
	Bucket(name string) (Bucket, query.Error)
}

// Bucket is a collection of key-value entries (typically
// key-document, but not always).
type Bucket interface {
	Name() string
	Count() (int64, query.Error)
	Scanners() ([]Scanner, query.Error)
	ScannerNames() ([]string, query.Error)
	Scanner(name string) (Scanner, query.Error)
	Fetch(id string) (*dparval.Value, query.Error)
	BulkFetch([]string) (map[string]*dparval.Value, query.Error)
	Release()
}

// RangeStatistics captures statistics for a range index (view or
// declarative btree index).
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

// Scanner is the base type for all scanners.
type Scanner interface {
	Name() string
	ScanAll(ch dparval.ValueChannel, warnch, errch query.ErrorChannel)
}

// FullScanner performs full bucket scans.
type FullScanner interface {
	Scanner
}

// Direction represents ASC and DESC
// TODO: Is this needed?
type Direction int

const (
	ASC  Direction = 1
	DESC           = 2
)

// RangeScanner is the base type for view and declarative index
// scanners.
type RangeScanner interface {
	Scanner
	Key() []string
	Direction() Direction
	Statistics() (RangeStatistics, query.Error)
}

// ViewScanner represents Couchbase views.
type ViewScanner interface {
	RangeScanner
}

// IndexScanner represents declarative btree indexes.
type IndexScanner interface {
	RangeScanner
}

// SearchScanner represents full text search indexes.
type SearchScanner interface {
	Scanner
}
