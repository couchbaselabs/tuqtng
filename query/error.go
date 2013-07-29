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
}

type ErrorChannel chan Error

func NewError(e error, internalMsg string) Error {
	return &err{level: EXCEPTION, cause: e, internalMsg: internalMsg}
}

func NewWarning(e error, internalMsg string) Error {
	return &err{level: WARNING, cause: e, internalMsg: internalMsg}
}

func NewNotice(e error, internalMsg string) Error {
	return &err{level: NOTICE, cause: e, internalMsg: internalMsg}
}

func NewInfo(e error, internalMsg string) Error {
	return &err{level: INFO, cause: e, internalMsg: internalMsg}
}

func NewLog(e error, internalMsg string) Error {
	return &err{level: LOG, cause: e, internalMsg: internalMsg}
}

func NewDebug(e error, internalMsg string) Error {
	return &err{level: DEBUG, cause: e, internalMsg: internalMsg}
}

type err struct {
	code        int32
	key         string
	cause       error
	internalMsg string
	level       int
}

func (e *err) Error() string {
	switch {
	default:
		return "Unspecified error."
	case e.internalMsg != "" && e.cause != nil:
		return e.internalMsg + " - cause: " + e.cause.Error()
	case e.internalMsg != "":
		return e.internalMsg
	case e.cause != nil:
		return e.cause.Error()
	}
}

func (e *err) Level() int {
	return e.level
}

func (e *err) Code() int32 {
	return e.code
}

func (e *err) TranslationKey() string {
	return e.key
}

func (e *err) Cause() error {
	return e.cause
}

func NewBucketDoesNotExist(bucket string) Error {
	return &err{level: EXCEPTION, internalMsg: fmt.Sprintf("Bucket %s does not exist", bucket)}
}

func NewPoolDoesNotExist(pool string) Error {
	return &err{level: EXCEPTION, internalMsg: fmt.Sprintf("Pool %s does not exist", pool)}
}
