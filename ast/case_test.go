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

	//"github.com/couchbaselabs/dparval"
)

func TestCase(t *testing.T) {

	boolTrue := NewLiteralBool(true)
	boolFalse := NewLiteralBool(false)
	numberSeven := NewLiteralNumber(7.5)
	numberNine := NewLiteralNumber(9.3)

	whenThenTrue := &WhenThen{
		When: boolTrue,
		Then: numberSeven,
	}

	whenThenFalse := &WhenThen{
		When: boolFalse,
		Then: numberSeven,
	}

	caseOne := NewCaseOperator()
	caseOne.WhenThens = []*WhenThen{whenThenTrue}
	caseOne.Else = numberNine

	caseTwo := NewCaseOperator()
	caseTwo.WhenThens = []*WhenThen{whenThenFalse}
	caseTwo.Else = numberNine

	tests := ExpressionTestSet{
		{caseOne, 7.5, nil},
		{caseTwo, 9.3, nil},
	}

	tests.Run(t)
}

func TestCaseStringRepresentation(t *testing.T) {

	boolFalse := NewLiteralBool(false)
	numberSeven := NewLiteralNumber(7.5)
	numberNine := NewLiteralNumber(9.3)

	whenThenFalse := &WhenThen{
		When: boolFalse,
		Then: numberSeven,
	}

	caseTwo := NewCaseOperator()
	caseTwo.WhenThens = []*WhenThen{whenThenFalse}
	caseTwo.Else = numberNine

	tests := ExpressionStringTestSet{
		{caseTwo, `CASE WHEN false THEN 7.5 ELSE 9.3 END`},
	}

	tests.Run(t)
}
