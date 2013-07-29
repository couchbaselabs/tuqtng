//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

/*

Package query contains artifacts shared across other packages in query
processing.

*/
package query

import (
	"encoding/json"
	"fmt"
)

const (
	EXCEPTION = iota
	WARNING
	NOTICE
	INFO
	LOG
	DEBUG
)

// Error will eventually include code, message key, and internal error
// object (cause) and message
type Error interface {
	error
	Code() int32
	TranslationKey() string
	Cause() error
	Level() int
	IsFatal() bool
}

type ErrorChannel chan Error

func NewError(e error, internalMsg string) Error {
	return &err{level: EXCEPTION, ICode: 5000, IKey: "Internal Error", ICause: e, InternalMsg: internalMsg}
}

func NewWarning(internalMsg string) Error {
	return &err{level: WARNING, InternalMsg: internalMsg}
}

func NewNotice(internalMsg string) Error {
	return &err{level: NOTICE, InternalMsg: internalMsg}
}

func NewInfo(internalMsg string) Error {
	return &err{level: INFO, InternalMsg: internalMsg}
}

func NewLog(internalMsg string) Error {
	return &err{level: LOG, InternalMsg: internalMsg}
}

func NewDebug(internalMsg string) Error {
	return &err{level: DEBUG, InternalMsg: internalMsg}
}

type err struct {
	ICode       int32
	IKey        string
	ICause      error
	InternalMsg string
	level       int
}

func (e *err) Error() string {
	switch {
	default:
		return "Unspecified error."
	case e.InternalMsg != "" && e.ICause != nil:
		return e.InternalMsg + " - cause: " + e.ICause.Error()
	case e.InternalMsg != "":
		return e.InternalMsg
	case e.ICause != nil:
		return e.ICause.Error()
	}
}

func (e *err) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"code":    e.ICode,
		"key":     e.IKey,
		"message": e.InternalMsg,
	}
	if e.ICause != nil {
		m["cause"] = e.ICause.Error()
	}
	return json.Marshal(m)
}

func (e *err) Level() int {
	return e.level
}

func (e *err) IsFatal() bool {
	if e.level == EXCEPTION {
		return true
	}
	return false
}

func (e *err) Code() int32 {
	return e.ICode
}

func (e *err) TranslationKey() string {
	return e.IKey
}

func (e *err) Cause() error {
	return e.ICause
}

func NewParseError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 4100, IKey: "parse_error", ICause: e, InternalMsg: msg}
}

func NewSemanticError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 4200, IKey: "semantic_error", ICause: e, InternalMsg: msg}
}

func NewBucketDoesNotExist(bucket string) Error {
	return &err{level: EXCEPTION, ICode: 4040, IKey: "bucket_not_found", InternalMsg: fmt.Sprintf("Bucket %s does not exist", bucket)}
}

func NewPoolDoesNotExist(pool string) Error {
	return &err{level: EXCEPTION, ICode: 4041, IKey: "pool_not_found", InternalMsg: fmt.Sprintf("Pool %s does not exist", pool)}
}

func NewTotalRowsInfo(rows int) Error {
	return &err{level: INFO, ICode: 100, IKey: "total_rows", InternalMsg: fmt.Sprintf("%d", rows)}
}

func NewTotalElapsedTimeInfo(time string) Error {
	return &err{level: INFO, ICode: 101, IKey: "total_elapsed_time", InternalMsg: fmt.Sprintf("%s", time)}
}
