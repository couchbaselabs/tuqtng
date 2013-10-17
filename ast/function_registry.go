//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

import ()

type FunctionCallConstructor func(operands FunctionArgExpressionList) FunctionCallExpression

var SystemFunctionRegistry map[string]FunctionCallConstructor = map[string]FunctionCallConstructor{
	// utility functions
	"LENGTH": NewFunctionCallLength,

	// aggregate functions
	"COUNT":     NewFunctionCallCount,
	"SUM":       NewFunctionCallSum,
	"AVG":       NewFunctionCallAvg,
	"MIN":       NewFunctionCallMin,
	"MAX":       NewFunctionCallMax,
	"ARRAY_AGG": NewFunctionCallArrayAgg,

	// comparison functions
	"GREATEST":        NewFunctionCallGreatest,
	"LEAST":           NewFunctionCallLeast,
	"IFMISSING":       NewFunctionCallIfMissing,
	"IFNULL":          NewFunctionCallIfNull,
	"IFMISSINGORNULL": NewFunctionCallIfMissingOrNull,
	"MISSINGIF":       NewFunctionCallMissingIf,
	"NULLIF":          NewFunctionCallNullIf,

	// meta/value functions
	"META":         NewFunctionCallMeta,
	"VALUE":        NewFunctionCallValue,
	"BASE64_VALUE": NewFunctionCallBase64Value,

	// numeric functions
	"CEIL":  NewFunctionCallCeil,
	"FLOOR": NewFunctionCallFloor,
	"ROUND": NewFunctionCallRound,
	"TRUNC": NewFunctionCallTrunc,

	// numeric infinite functions
	"IFNAN":      NewFunctionCallIfNaN,
	"IFPOSINF":   NewFunctionCallIfPosInf,
	"IFNEGINF":   NewFunctionCallIfNegInf,
	"IFINF":      NewFunctionCallIfInf,
	"IFNANORINF": NewFunctionCallIfNaNOrInf,
	"FIRSTNUM":   NewFunctionCallFirstNum,
	"NANIF":      NewFunctionCallNaNIf,
	"POSINFIF":   NewFunctionCallPosInfIf,
	"NEGINFIF":   NewFunctionCallNegInfIf,

	// string functions
	"LOWER":  NewFunctionCallLower,
	"UPPER":  NewFunctionCallUpper,
	"TRIM":   NewFunctionCallTrim,
	"RTRIM":  NewFunctionCallRTrim,
	"LTRIM":  NewFunctionCallLTrim,
	"SUBSTR": NewFunctionCallSubStr,

	// date functions
	"DATE_PART": NewFunctionCallDatePart,
	"NOW_STR":   NewFunctionCallNowStr,
}
