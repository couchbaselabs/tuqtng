//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ast

type FunctionCallConstructor func(operands FunctionArgExpressionList) FunctionCallExpression

var SystemFunctionRegistry map[string]FunctionCallConstructor = map[string]FunctionCallConstructor{
	// utility functions
	"LENGTH":        NewFunctionCallLength,
	"ARRAY_LENGTH":  NewFunctionCallArrayLength,
	"OBJECT_LENGTH": NewFunctionCallObjectLength,
	"POLY_LENGTH":   NewFunctionCallPolyLength,

	//array utility functions
	"ARRAY_CONCAT":  NewFunctionCallArrayConcat,
	"ARRAY_APPEND":  NewFunctionCallArrayAppend,
	"ARRAY_PREPEND": NewFunctionCallArrayPrepend,
	"ARRAY_REMOVE":  NewFunctionCallArrayRemove,

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
	"DIV":   NewFunctionCallDiv,

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
	"SPLIT":  NewFunctionCallSplit,

	// date functions
	"DATE_PART_STR":    NewFunctionCallDatePartStr,
	"NOW_STR":          NewFunctionCallNowStr,
	"DATE_PART_MILLIS": NewFunctionCallDatePartMillis,
	"NOW_MILLIS":       NewFunctionCallNowMillis,
	"STR_TO_MILLIS":    NewFunctionCallStrToMillis,
	"MILLIS":           NewFunctionCallStrToMillis,
	"MILLIS_TO_STR":    NewFunctionCallMillisToStr,

	// typecast functions
	"TO_NUM":   NewFunctionCallToNum,
	"TO_STR":   NewFunctionCallToStr,
	"TO_BOOL":  NewFunctionCallToBool,
	"TO_ATOM":  NewFunctionCallToAtom,
	"TO_ARRAY": NewFunctionCallToArray,

	// typecheck functions
	"TYPE_NAME": NewFunctionCallTypeName,
	"IS_NUM":    NewFunctionCallIsNum,
	"IS_STR":    NewFunctionCallIsStr,
	"IS_BOOL":   NewFunctionCallIsBool,
	"IS_ATOM":   NewFunctionCallIsAtom,
	"IS_ARRAY":  NewFunctionCallIsArray,
	"IS_OBJ":    NewFunctionCallIsObj,
}
