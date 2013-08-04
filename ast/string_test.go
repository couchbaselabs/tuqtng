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

	"github.com/couchbaselabs/dparval"
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

func TestStringStringRepresentation(t *testing.T) {
	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")

	tests := ExpressionStringTestSet{
		{NewStringConcatenateOperator(stringCouchbase, stringServer), `"Couchbase" || "Server"`},
	}

	tests.Run(t)
}

func TestStringValidate(t *testing.T) {

	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")

	tests := ExpressionValidateTestSet{
		{NewStringConcatenateOperator(stringCouchbase, stringServer), nil},
		// first arg invalid
		{NewStringConcatenateOperator(notValidExpression, stringServer), notValidExpressionError},
		// second arg invalid
		{NewStringConcatenateOperator(stringCouchbase, notValidExpression), notValidExpressionError},
	}

	tests.Run(t)
}

func TestStringVerifyFormalNotation(t *testing.T) {

	stringCouchbase := NewLiteralString("Couchbase")
	stringServer := NewLiteralString("Server")

	tests := ExpressionVerifyFormalNotationTestSet{
		{NewStringConcatenateOperator(stringCouchbase, stringServer), nil, nil},
		// first arg not formal
		{NewStringConcatenateOperator(notFormalExpression, stringServer), nil, notFormalExpressionError},
		// second arg not formal
		{NewStringConcatenateOperator(stringCouchbase, notFormalExpression), nil, notFormalExpressionError},
	}

	tests.Run(t, []string{"bucket", "child"}, "")

	// again with single bucket
	tests = ExpressionVerifyFormalNotationTestSet{
		// first arg not formal
		{NewStringConcatenateOperator(notFormalExpression, stringServer), nil, nil},
		// second arg not formal
		{NewStringConcatenateOperator(stringCouchbase, notFormalExpression), nil, nil},
	}

	tests.Run(t, []string{"bucket"}, "bucket")
}
