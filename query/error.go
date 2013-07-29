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

// Error will eventually include code, message key, and internal error
// object (cause) and message
type Error interface {
	error
	Code() int32
	TranslationKey() string
	Cause() error
}

type ErrorChannel chan Error

func NewError(e error, internalMsg string) Error {
	return &err{cause: e, internalMsg: internalMsg}
}

type err struct {
	code        int32
	key         string
	cause       error
	internalMsg string
}

func (e *err) Error() string {
	switch {
	default:
		return "Unspecified error."
	case e.internalMsg != "":
		return e.internalMsg
	case e.cause != nil:
		return e.cause.Error()
	}
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
	return &err{internalMsg: fmt.Sprintf("Bucket %s does not exist", bucket)}
}

func NewPoolDoesNotExist(pool string) Error {
	return &err{internalMsg: fmt.Sprintf("Pool %s does not exist", pool)}
}
