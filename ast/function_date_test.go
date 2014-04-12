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
)

func TestFunctionDatePartStr(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-02T00:00:00.555-07:00")),
				NewFunctionArgExpression(NewLiteralString("year")),
			}),
			2014.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T00:00:00+07:00")),
				NewFunctionArgExpression(NewLiteralString("month")),
			}),
			4.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T00:00:00.555")),
				NewFunctionArgExpression(NewLiteralString("day")),
			}),
			11.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T03:00:00")),
				NewFunctionArgExpression(NewLiteralString("hour")),
			}),
			3.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-02 00:00:00.555-07:00")),
				NewFunctionArgExpression(NewLiteralString("year")),
			}),
			2014.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 00:00:00+07:00")),
				NewFunctionArgExpression(NewLiteralString("month")),
			}),
			4.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 00:00:00.555")),
				NewFunctionArgExpression(NewLiteralString("day")),
			}),
			11.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 03:00:00")),
				NewFunctionArgExpression(NewLiteralString("hour")),
			}),
			3.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-02")),
				NewFunctionArgExpression(NewLiteralString("year")),
			}),
			2014.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("01:00:00.555-07:00")),
				NewFunctionArgExpression(NewLiteralString("hour")),
			}),
			1.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("00:02:00+07:00")),
				NewFunctionArgExpression(NewLiteralString("minute")),
			}),
			2.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("00:00:04.555")),
				NewFunctionArgExpression(NewLiteralString("second")),
			}),
			4.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_STR", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("03:00:00")),
				NewFunctionArgExpression(NewLiteralString("hour")),
			}),
			3.0,
			nil,
		},
	}

	tests.Run(t)
}

// The UNIX timestamp 1397203323000 represents the datetime
// "2014-04-11 01:02:03"
func TestFunctionDatePartMillis(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("year")),
			}),
			2014.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("month")),
			}),
			4.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("day")),
			}),
			11.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("hour")),
			}),
			1.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("minute")),
			}),
			2.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("second")),
			}),
			3.0,
			nil,
		},
		{
			NewFunctionCall("DATE_PART_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralNumber(1397203323000)),
				NewFunctionArgExpression(NewLiteralString("dow")),
			}),
			5.0,
			nil,
		},
	}

	tests.Run(t)
}

// The UNIX timestamp 1397203323000 represents the datetime
// "2014-04-11 01:02:03"
func TestFunctionStrToMillis(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("STR_TO_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T01:02:03.000-07:00")),
			}),
			1397203323000.0,
			nil,
		},
		{
			NewFunctionCall("STR_TO_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T01:02:03-07:00")),
			}),
			1397203323000.0,
			nil,
		},
		{
			NewFunctionCall("STR_TO_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 01:02:03.000-07:00")),
			}),
			1397203323000.0,
			nil,
		},
		{
			NewFunctionCall("STR_TO_MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 01:02:03-07:00")),
			}),
			1397203323000.0,
			nil,
		},
	}

	tests.Run(t)
}

// The UNIX timestamp 1397203323000 represents the datetime
// "2014-04-11 01:02:03"
func TestFunctionMillis(t *testing.T) {
	tests := ExpressionTestSet{
		{
			NewFunctionCall("MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T01:02:03.000-07:00")),
			}),
			1397203323000.0,
			nil,
		},
		{
			NewFunctionCall("MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11T01:02:03-07:00")),
			}),
			1397203323000.0,
			nil,
		},
		{
			NewFunctionCall("MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 01:02:03.000-07:00")),
			}),
			1397203323000.0,
			nil,
		},
		{
			NewFunctionCall("MILLIS", FunctionArgExpressionList{
				NewFunctionArgExpression(NewLiteralString("2014-04-11 01:02:03-07:00")),
			}),
			1397203323000.0,
			nil,
		},
	}

	tests.Run(t)
}
