//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package catalog

// Error will be moved to a shared package.
type Error error

// Item will be moved to a shared package.
type Item interface {
}

// Site represents a cluster or single-node server.
type Site interface {
	URL() string
	PoolNames() ([]string, Error)
	Pool(name string) (Pool, Error)
}

// Pool represents a logical authentication, query, and resource
// allocation boundary, as well as a grouping of buckets.
type Pool interface {
	Name() string
	BucketNames() ([]string, Error)
	Bucket(name string) (Bucket, Error)
}

// Bucket is a collection of key-value entries (typically
// key-document, but not always).
type Bucket interface {
	Name() string
	Count() (int64, Error)
	Scanners() ([]Scanner, Error)
	Lookup(id string) (Item, Error)
}

// RangeStatistics captures statistics for a range index (view or
// declarative btree index).
type RangeStatistics interface {
	Count() (int64, Error)
	Min() (Item, Error)
	Max() (Item, Error)
	DistinctCount(int64, Error)
	Bins() ([]Bin, Error)
}

// Bin represents a range bin within IndexStatistics.
type Bin interface {
	Count() (int64, Error)
	Min() (Item, Error)
	Max() (Item, Error)
	DistinctCount(int64, Error)
}

type ItemChannel chan *Item

// Scanner is the base type for all scanners.
type Scanner interface {
	Channel() (ItemChannel, Error)
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
	Name() string
	Key() []string
	Direction() Direction
	Statistics() (RangeStatistics, Error)
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
