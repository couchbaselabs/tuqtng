//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import (
	"testing"

	"github.com/mschoch/dparval"
)

func TestString(t *testing.T) {

	dneProperty := NewProperty("foo")
	numberFive := NewLiteralNumber(5.0)
	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")

	tests := ExpressionTestSet{
		{NewStringConcatenateOperator(stringCouchbase, stringServer), "CouchbaseServer", nil},
		{NewStringConcatenateOperator(numberFive, stringServer), nil, nil},
		{NewStringConcatenateOperator(stringCouchbase, numberFive), nil, nil},
		{NewStringConcatenateOperator(dneProperty, stringServer), nil, &dparval.Undefined{"foo"}},
		{NewStringConcatenateOperator(stringCouchbase, dneProperty), nil, &dparval.Undefined{"foo"}},
	}

	tests.Run(t)

}
