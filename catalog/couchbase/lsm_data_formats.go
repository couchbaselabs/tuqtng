package couchbase

import (
	"github.com/couchbaselabs/tuqtng/catalog"
)

//REQUEST DATA FORMATS

type RequestType int

const (
	CREATE RequestType = iota
	DROP   RequestType = iota
	LIST   RequestType = iota
	NOTIFY RequestType = iota
	NODES  RequestType = iota
	SCAN   RequestType = iota
	STATS  RequestType = iota
)

type IndexRequest struct {
	Type       RequestType
	Index      IndexInfo
	ServerUuid string
	Params     QueryParams
}

type IndexInfo struct {
	Name       string
	Uuid       string
	Using      catalog.IndexType
	OnExprList []string
	Bucket     string
	IsPrimary  bool
}

type QueryParams struct {
	Low       string
	High      string
	Inclusion catalog.RangeInclusion
	Limit     int64
}

//RESPONSE DATA FORMATS

type ResponseStatus int

const (
	SUCCESS       ResponseStatus = iota
	ERROR         ResponseStatus = iota
	INVALID_CACHE ResponseStatus = iota
)

type IndexScanResponse struct {
	Status    ResponseStatus
	TotalRows int64
	Rows      []IndexRow
	Errors    []IndexError
}

type IndexMetaResponse struct {
	Status     ResponseStatus
	Indexes    []IndexInfo
	ServerUuid string
	Nodes      []NodeInfo
	Errors     []IndexError
}

type IndexStatsResponse struct {
	Status ResponseStatus
	Stats  IndexStats
	Errors []IndexError
}

type IndexRow struct {
	Key   interface{}
	Value interface{}
}

type IndexError struct {
	Code string
	Msg  string
}

type NodeInfo struct {
	IndexerURL string
}

type IndexStats struct {
	Count         int64
	Min           string
	Max           string
	DistinctCount int64
	Bins          []IndexBin
}

type IndexBin struct {
	Count         int64
	Min           string
	Max           string
	DistinctCount int64
}
