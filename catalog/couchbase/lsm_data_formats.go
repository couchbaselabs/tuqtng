package couchbase

import (
	"github.com/couchbaselabs/tuqtng/catalog"
)

//REQUEST DATA FORMATS

type RequestType string

const (
	CREATE RequestType = "create"
	DROP   RequestType = "drop"
	LIST   RequestType = "list"
	NOTIFY RequestType = "notify"
	NODES  RequestType = "nodes"
	SCAN   RequestType = "scan"
	STATS  RequestType = "stats"
)

type IndexRequest struct {
	Type       RequestType `json:"type,omitempty"`
	Index      IndexInfo   `json:"index,omitempty"`
	ServerUuid string      `json:"serverUuid,omitempty"`
	Params     QueryParams `json:"params,omitempty"`
}

type IndexInfo struct {
	Name       string            `json:"name,omitempty"`
	Uuid       string            `json:"uuid, omitempty"`
	Using      catalog.IndexType `json:"using,omitempty"`
	OnExprList []string          `json:"onExprList,omitempty"`
	Bucket     string            `json:"bucket,omitempty"`
	IsPrimary  bool              `json:"isPrimary,omitempty"`
}

type QueryParams struct {
	ScanType  ScanType               `json:"scanType,omitempty"`
	Low       [][]byte               `json:"low,omitempty"`
	High      [][]byte               `json:"high,omitempty"`
	Inclusion catalog.RangeInclusion `json:"inclusion,omitempty"`
	Limit     int64                  `json:"limit,omitempty"`
}

type ScanType string

const (
	COUNT      ScanType = "count"
	EXISTS     ScanType = "exists"
	LOOKUP     ScanType = "lookup"
	RANGESCAN  ScanType = "rangeScan"
	FULLSCAN   ScanType = "fullScan"
	RANGECOUNT ScanType = "rangeCount"
)

//RESPONSE DATA FORMATS

type ResponseStatus string

const (
	SUCCESS       ResponseStatus = "success"
	ERROR         ResponseStatus = "error"
	INVALID_CACHE ResponseStatus = "invalid_cache"
)

type IndexScanResponse struct {
	Status    ResponseStatus `json:"status,omitempty"`
	TotalRows uint64         `json:"totalrows,omitempty"`
	Rows      []IndexRow     `json:"rows,omitempty"`
	Errors    []IndexError   `json:"errors,omitempty"`
}

type IndexMetaResponse struct {
	Status     ResponseStatus `json:"status,omitempty"`
	Indexes    []IndexInfo    `json:"indexes,omitempty"`
	ServerUuid string         `json:"serverUuid,omitempty"`
	Nodes      []NodeInfo     `json:"nodes,omitempty"`
	Errors     []IndexError   `json:"errors,omitempty"`
}

type IndexStatsResponse struct {
	Status ResponseStatus `json:"status,omitempty"`
	Stats  IndexStats     `json:"stats,omitempty"`
	Errors []IndexError   `json:"errors,omitempty"`
}

type IndexRow struct {
	Key   [][]byte `json:"key,omitempty"`
	Value string   `json:"value,omitempty"`
}

type IndexError struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type NodeInfo struct {
	IndexerURL string `json:"indexerURL,omitempty"`
}

type IndexStats struct {
	Count         int64      `json:"count,omitempty"`
	Min           string     `json:"min,omitempty"`
	Max           string     `json:"max,omitempty"`
	DistinctCount int64      `json:"distinctCount,omitempty"`
	Bins          []IndexBin `json:"bins,omitempty"`
}

type IndexBin struct {
	Count         int64  `json:"count,omitempty"`
	Min           string `json:"min,omitempty"`
	Max           string `json:"max,omitempty"`
	DistinctCount int64  `json:"distinctCount,omitempty"`
}
