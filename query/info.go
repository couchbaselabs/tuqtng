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

import ()

// Info will eventually include code, message key, and internal error
// object (cause) and message
type Info interface {
	error
	Code() int32
	TranslationKey() string
	Cause() error
}

type InfoChannel chan Info

func NewInfo(e error, internalMsg string) Error {
	return &err{cause: e, internalMsg: internalMsg}
}

type info struct {
	code        int32
	key         string
	cause       error
	internalMsg string
}

func (i *info) Error() string {
	switch {
	default:
		return "Unspecified error."
	case i.internalMsg != "" && i.cause != nil:
		return i.internalMsg + " - cause: " + i.cause.Error()
	case i.internalMsg != "":
		return i.internalMsg
	case i.cause != nil:
		return i.cause.Error()
	}
}

func (i *info) Code() int32 {
	return i.code
}

func (i *info) TranslationKey() string {
	return i.key
}

func (i *info) Cause() error {
	return i.cause
}
