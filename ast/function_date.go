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
	"fmt"
	"math"
	"time"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/network"
)

var supportedDateFormats = []string{
	"2006-01-02T15:04:05.999Z07:00", // time.RFC3339Milli
	"2006-01-02T15:04:05Z07:00",     // time.RFC3339
	"2006-01-02T15:04:05.999",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05.999Z07:00",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05.999",
	"2006-01-02 15:04:05",
	"2006-01-02",
	"15:04:05.999Z07:00",
	"15:04:05Z07:00",
	"15:04:05.999",
	"15:04:05",
}

type FunctionCallDatePartStr struct {
	FunctionCall
}

func NewFunctionCallDatePartStr(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallDatePartStr{
		FunctionCall{
			Type:     "function",
			Name:     "DATE_PART_STR",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallDatePartStr) Copy() Expression {
	return &FunctionCallDatePartStr{
		FunctionCall{
			Type:     "function",
			Name:     "DATE_PART_STR",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallDatePartStr) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument (part)
	dv, err := this.Operands[0].Expr.Evaluate(item)

	// the part must be a string, undefined results in null
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if dv.Type() == dparval.STRING {
		date := dv.Value()
		switch date := date.(type) {
		case string:
			// now look at the second argument
			av, err := this.Operands[1].Expr.Evaluate(item)

			// *currently* the date must be an rfc3339 string
			// the part must be a string, undefined results in null
			if err != nil {
				switch err := err.(type) {
				case *dparval.Undefined:
					// undefined returns null
					return dparval.NewValue(nil), nil
				default:
					// any other error return to caller
					return nil, err
				}
			}

			if av.Type() == dparval.STRING {
				part := av.Value()
				switch part := part.(type) {
				case string:
					var t time.Time
					for _, sf := range supportedDateFormats {
						t, err = time.Parse(sf, date)
						if err == nil {
							// found a matching format
							break
						}
					}

					if err != nil {
						return nil, fmt.Errorf("Date not in a recognized format. Error: %v", err)
					}

					return datePart(t, part)
				}
			}

		}
	}

	return dparval.NewValue(nil), nil
}

func (this *FunctionCallDatePartStr) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallDatePartMillis struct {
	FunctionCall
}

func NewFunctionCallDatePartMillis(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallDatePartMillis{
		FunctionCall{
			Type:     "function",
			Name:     "DATE_PART_MILLIS",
			Operands: operands,
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallDatePartMillis) Copy() Expression {
	return &FunctionCallDatePartMillis{
		FunctionCall{
			Type:     "function",
			Name:     "DATE_PART_MILLIS",
			Operands: this.Operands.Copy(),
			minArgs:  2,
			maxArgs:  2,
		},
	}
}

func (this *FunctionCallDatePartMillis) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first evaluate the argument (part)
	dv, err := this.Operands[0].Expr.Evaluate(item)

	// the part must be a string, undefined results in ull
	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	if dv.Type() == dparval.NUMBER {
		millis := dv.Value()
		switch millis := millis.(type) {
		case float64:
			// now look at the second argument
			av, err := this.Operands[1].Expr.Evaluate(item)

			// the part must be a string, undefined results in null
			if err != nil {
				switch err := err.(type) {
				case *dparval.Undefined:
					// undefined returns null
					return dparval.NewValue(nil), nil
				default:
					// any other error return to caller
					return nil, err
				}
			}

			if av.Type() == dparval.STRING {
				part := av.Value()
				switch part := part.(type) {
				case string:
					t := time.Unix(0, int64(millis)*1000000)
					return datePart(t, part)
				}
			}

		}
	}

	return dparval.NewValue(nil), nil
}

func (this *FunctionCallDatePartMillis) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

func datePart(t time.Time, part string) (*dparval.Value, error) {
	// now look for the requested part
	switch part {
	case "century":
		cen := float64(t.Year() / 100.0)
		if cen > 0 {
			cen = cen + 1 // there is no century 0
		}
		return dparval.NewValue(cen), nil
	case "day":
		return dparval.NewValue(float64(t.Day())), nil
	case "decade":
		return dparval.NewValue(float64(t.Year() / 10.0)), nil
	case "dow":
		return dparval.NewValue(float64(t.Weekday())), nil
	case "doy":
		return dparval.NewValue(float64(t.YearDay())), nil
	case "epoch":
		return dparval.NewValue(float64(t.Unix())), nil
	case "hour":
		return dparval.NewValue(float64(t.Hour())), nil
	case "isodow":
		isodow := float64(t.Weekday())
		if isodow == 0.0 {
			isodow = 7.0
		}
		return dparval.NewValue(isodow), nil
	case "isoyear":
		y, _ := t.ISOWeek()
		return dparval.NewValue(float64(y)), nil
	case "microseconds":
		us := float64(t.Second() * 1000000)
		us = us + float64(t.Nanosecond()/int(time.Microsecond))
		return dparval.NewValue(us), nil
	case "millennium":
		mil := float64(t.Year() / 1000.0)
		if mil > 0 {
			mil = mil + 1 // there is no millennium 0
		}
		return dparval.NewValue(mil), nil
	case "milliseconds":
		ms := float64(t.Second() * 1000)
		ms = ms + float64(t.Nanosecond()/int(time.Millisecond))
		return dparval.NewValue(ms), nil
	case "minute":
		return dparval.NewValue(float64(t.Minute())), nil
	case "month":
		return dparval.NewValue(float64(t.Month())), nil
	case "quarter":
		return dparval.NewValue(math.Ceil(float64(t.Month()) / 3.0)), nil
	case "second":
		return dparval.NewValue(float64(t.Second())), nil
	case "timezone":
		_, z := t.Zone()
		return dparval.NewValue(float64(z)), nil
	case "timezone_hour":
		_, z := t.Zone()
		zh := int64(z / (60 * 60))
		return dparval.NewValue(float64(zh)), nil
	case "timezone_minute":
		_, z := t.Zone()
		zh := int(z / (60 * 60))
		z = z - (zh * (60 * 60))
		zm := int64(z / 60)
		return dparval.NewValue(float64(zm)), nil
	case "week":
		_, w := t.ISOWeek()
		return dparval.NewValue(float64(w)), nil
	case "year":
		return dparval.NewValue(float64(t.Year())), nil
	default:
		return nil, fmt.Errorf("Unknown date part %s", part)
	}
}

type FunctionCallNowStr struct {
	FunctionCall
}

func NewFunctionCallNowStr(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallNowStr{
		FunctionCall{
			Type:     "function",
			Name:     "NOW_STR",
			Operands: operands,
			minArgs:  0,
			maxArgs:  0,
		},
	}
}

func (this *FunctionCallNowStr) Copy() Expression {
	return &FunctionCallNowStr{
		FunctionCall{
			Type:     "function",
			Name:     "NOW_STR",
			Operands: this.Operands.Copy(),
			minArgs:  0,
			maxArgs:  0,
		},
	}
}

func (this *FunctionCallNowStr) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first retrieve the query object
	q := item.GetAttachment("query")

	query, ok := q.(network.Query)
	if !ok {
		return dparval.NewValue(nil), nil
	}

	return dparval.NewValue(timeToStr(query.StartTime())), nil
}

func (this *FunctionCallNowStr) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallNowMillis struct {
	FunctionCall
}

func NewFunctionCallNowMillis(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallNowMillis{
		FunctionCall{
			Type:     "function",
			Name:     "NOW_MILLIS",
			Operands: operands,
			minArgs:  0,
			maxArgs:  0,
		},
	}
}

func (this *FunctionCallNowMillis) Copy() Expression {
	return &FunctionCallNowMillis{
		FunctionCall{
			Type:     "function",
			Name:     "NOW_MILLIS",
			Operands: this.Operands.Copy(),
			minArgs:  0,
			maxArgs:  0,
		},
	}
}

func (this *FunctionCallNowMillis) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	// first retrieve the query object
	q := item.GetAttachment("query")

	query, ok := q.(network.Query)
	if !ok {
		return dparval.NewValue(nil), nil
	}

	return dparval.NewValue(timeToMillis(query.StartTime())), nil
}

func (this *FunctionCallNowMillis) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallStrToMillis struct {
	FunctionCall
}

func NewFunctionCallStrToMillis(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallStrToMillis{
		FunctionCall{
			Type:     "function",
			Name:     "STR_TO_MILLIS",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallStrToMillis) Copy() Expression {
	return &FunctionCallStrToMillis{
		FunctionCall{
			Type:     "function",
			Name:     "STR_TO_MILLIS",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallStrToMillis) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	av, err := this.Operands[0].Expr.Evaluate(item)

	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	val := av.Value()

	switch val := val.(type) {
	case string:
		t, err := strToTime(val)
		if err != nil {
			return nil, fmt.Errorf("Date not in a recognized format.")
		}
		return dparval.NewValue(timeToMillis(t)), nil
	default:
		return dparval.NewValue(nil), nil
	}
}

func (this *FunctionCallStrToMillis) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

type FunctionCallMillisToStr struct {
	FunctionCall
}

func NewFunctionCallMillisToStr(operands FunctionArgExpressionList) FunctionCallExpression {
	return &FunctionCallMillisToStr{
		FunctionCall{
			Type:     "function",
			Name:     "MILLIS_TO_STR",
			Operands: operands,
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallMillisToStr) Copy() Expression {
	return &FunctionCallMillisToStr{
		FunctionCall{
			Type:     "function",
			Name:     "MILLIS_TO_STR",
			Operands: this.Operands.Copy(),
			minArgs:  1,
			maxArgs:  1,
		},
	}
}

func (this *FunctionCallMillisToStr) Evaluate(item *dparval.Value) (*dparval.Value, error) {
	av, err := this.Operands[0].Expr.Evaluate(item)

	if err != nil {
		switch err := err.(type) {
		case *dparval.Undefined:
			// undefined returns null
			return dparval.NewValue(nil), nil
		default:
			// any other error return to caller
			return nil, err
		}
	}

	val := av.Value()

	switch val := val.(type) {
	case float64:
		t := millisToTime(val)
		return dparval.NewValue(timeToStr(t)), nil
	default:
		return dparval.NewValue(nil), nil
	}
}

func (this *FunctionCallMillisToStr) Accept(ev ExpressionVisitor) (Expression, error) {
	return ev.Visit(this)
}

func strToTime(s string) (time.Time, error) {
	var t time.Time
	var err error

	for _, sf := range supportedDateFormats {
		t, err = time.Parse(sf, s)
		if err == nil {
			// found a matching format
			return t, nil
		}
	}

	return t, err
}

func timeToStr(t time.Time) string {
	return t.Format(supportedDateFormats[0])
}

func millisToTime(millis float64) time.Time {
	return time.Unix(0, int64(millis)*1000000)
}

func timeToMillis(t time.Time) float64 {
	return float64(t.UnixNano()) / 1000000
}
