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

	whenThenWhenMissing := &WhenThen{
		When: NewProperty("dne"),
		Then: numberSeven,
	}

	whenThenThenMissing := &WhenThen{
		When: boolTrue,
		Then: NewProperty("dne"),
	}

	caseOne := NewCaseOperator()
	caseOne.WhenThens = []*WhenThen{whenThenTrue}
	caseOne.Else = numberNine

	caseTwo := NewCaseOperator()
	caseTwo.WhenThens = []*WhenThen{whenThenFalse}
	caseTwo.Else = numberNine

	caseThree := NewCaseOperator()
	caseThree.WhenThens = []*WhenThen{whenThenWhenMissing}
	caseThree.Else = numberNine

	caseFour := NewCaseOperator()
	caseFour.WhenThens = []*WhenThen{whenThenThenMissing}
	caseFour.Else = numberNine

	caseFive := NewCaseOperator()
	caseFive.WhenThens = []*WhenThen{whenThenFalse}
	caseFive.Else = NewProperty("dne")

	tests := ExpressionTestSet{
		{caseOne, 7.5, nil},
		{caseTwo, 9.3, nil},
		{caseThree, 9.3, nil},
		{caseFour, nil, &dparval.Undefined{"dne"}},
		{caseFive, nil, &dparval.Undefined{"dne"}},
	}

	tests.Run(t)
}

func TestCaseStringRepresentation(t *testing.T) {

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

	caseTwo := NewCaseOperator()
	caseTwo.WhenThens = []*WhenThen{whenThenTrue, whenThenFalse}
	caseTwo.Else = numberNine

	tests := ExpressionStringTestSet{
		{caseTwo, `CASE WHEN true THEN 7.5 WHEN false THEN 7.5 ELSE 9.3 END`},
	}

	tests.Run(t)
}

func TestCaseValidate(t *testing.T) {

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

	WhenThenWhenInvalid := &WhenThen{
		When: notValidExpression,
		Then: numberSeven,
	}

	WhenThenThenInvalid := &WhenThen{
		When: boolFalse,
		Then: notValidExpression,
	}

	caseTwo := NewCaseOperator()
	caseTwo.WhenThens = []*WhenThen{whenThenTrue, whenThenFalse}
	caseTwo.Else = numberNine

	caseThree := NewCaseOperator()
	caseThree.WhenThens = []*WhenThen{WhenThenWhenInvalid, whenThenFalse}
	caseThree.Else = numberNine

	caseFour := NewCaseOperator()
	caseFour.WhenThens = []*WhenThen{WhenThenThenInvalid, whenThenFalse}
	caseFour.Else = numberNine

	caseFive := NewCaseOperator()
	caseFive.WhenThens = []*WhenThen{whenThenTrue, whenThenFalse}
	caseFive.Else = notValidExpression

	tests := ExpressionValidateTestSet{
		{caseTwo, nil},
		{caseThree, notValidExpressionError},
		{caseFour, notValidExpressionError},
		{caseFive, notValidExpressionError},
	}

	tests.Run(t)
}

func TestCaseVerifyFormalNotation(t *testing.T) {

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

	WhenThenWhenNotFormal := &WhenThen{
		When: notFormalExpression,
		Then: numberSeven,
	}

	WhenThenThenNotFormal := &WhenThen{
		When: boolFalse,
		Then: notFormalExpression,
	}

	caseTwo := NewCaseOperator()
	caseTwo.WhenThens = []*WhenThen{whenThenTrue, whenThenFalse}
	caseTwo.Else = numberNine

	caseThree := NewCaseOperator()
	caseThree.WhenThens = []*WhenThen{WhenThenWhenNotFormal, whenThenFalse}
	caseThree.Else = numberNine

	caseFour := NewCaseOperator()
	caseFour.WhenThens = []*WhenThen{WhenThenThenNotFormal, whenThenFalse}
	caseFour.Else = numberNine

	caseFive := NewCaseOperator()
	caseFive.WhenThens = []*WhenThen{whenThenTrue, whenThenFalse}
	caseFive.Else = notFormalExpression

	tests := ExpressionVerifyFormalNotationTestSet{
		{caseTwo, nil, nil},
		{caseThree, nil, notFormalExpressionError},
		{caseFour, nil, notFormalExpressionError},
		{caseFive, nil, notFormalExpressionError},
	}

	tests.Run(t, []string{"bucket"}, "bucket")
}
