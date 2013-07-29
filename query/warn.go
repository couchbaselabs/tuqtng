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

// Warning will eventually include code, message key, and internal error
// object (cause) and message
type Warning interface {
	error
	Code() int32
	TranslationKey() string
	Cause() error
}

type WarningChannel chan Warning

func NewWarning(e error, internalMsg string) Warning {
	return &warning{cause: e, internalMsg: internalMsg}
}

type warning struct {
	code        int32
	key         string
	cause       error
	internalMsg string
}

func (w *warning) Error() string {
	switch {
	default:
		return "Unspecified error."
	case w.internalMsg != "" && w.cause != nil:
		return w.internalMsg + " - cause: " + w.cause.Error()
	case w.internalMsg != "":
		return w.internalMsg
	case w.cause != nil:
		return w.cause.Error()
	}
}

func (w *warning) Code() int32 {
	return w.code
}

func (w *warning) TranslationKey() string {
	return w.key
}

func (w *warning) Cause() error {
	return w.cause
}
