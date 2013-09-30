//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package plan

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/couchbaselabs/dparval"
	"github.com/couchbaselabs/tuqtng/ast"
	"github.com/couchbaselabs/tuqtng/catalog"
)

func TestPlanJSON(t *testing.T) {

	tests := []struct {
		input *Plan
	}{
		{&Plan{NewScan("pool", "bucket", "index", nil)}},
		{&Plan{NewFetch(NewScan("pool", "bucket", "index", nil), "pool", "bucket", nil, "bucket")}},
		{&Plan{NewFilter(NewFetch(NewScan("pool", "bucket", "index", nil), "pool", "bucket", nil, "bucket"), ast.NewLiteralBool(true))}},
		{&Plan{NewFilter(NewFetch(NewScan("pool", "bucket", "index", nil), "pool", "bucket", nil, "bucket"), ast.NewPlusOperator(ast.NewLiteralNumber(1.0), ast.NewLiteralNumber(1.0)))}},
	}

	for _, x := range tests {
		_, err := json.MarshalIndent(x.input, "", "    ")
		if err != nil {
			t.Errorf("error serializing json: %v", err)
		}
		// FIXME test doesn't actually assert anything
	}

}

func TestFurtherness(t *testing.T) {
	res := compareHigh(catalog.LookupValue{dparval.NewValue(5.0)}, catalog.LookupValue{dparval.NewValue(8.0)})
	if res != -1 {
		t.Errorf("Expected 8 > 5, got 5 > 8")
	}

	res = compareHigh(catalog.LookupValue{dparval.NewValue(5.0)}, nil)
	if res != -1 {
		t.Errorf("Expected nil > 5, got 5 > nil")
	}

	res = compareLow(catalog.LookupValue{dparval.NewValue(5.0)}, catalog.LookupValue{dparval.NewValue(8.0)})
	if res != -1 {
		t.Errorf("Expected 8 > 5, got 5 > 8")
	}

	res = compareLow(catalog.LookupValue{dparval.NewValue(5.0)}, nil)
	if res != 1 {
		t.Errorf("Expected 8 > 5, got 5 > 8")
	}
}

func TestOverlap(t *testing.T) {
	tests := []struct {
		a       *ScanRange
		b       *ScanRange
		overlap *ScanRange
	}{
		// the following test overlap of x < 5 and x > 3, with various combinations of endpoing inclusion
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Low},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: nil, Inclusion: catalog.High},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Neither},
		},
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: nil, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Both},
		},
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Low},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: nil, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Low},
		},
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: nil, Inclusion: catalog.High},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.High},
		},

		// these ranges do not overlap x < 3 and x > 5 (no matter what the inclusion)
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Low},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(5.0)}, High: nil, Inclusion: catalog.High},
			nil,
		},
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(5.0)}, High: nil, Inclusion: catalog.Both},
			nil,
		},
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Low},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(5.0)}, High: nil, Inclusion: catalog.Both},
			nil,
		},
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(5.0)}, High: nil, Inclusion: catalog.High},
			nil,
		},

		// x <= 3 and x >= 3
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: nil, Inclusion: catalog.Both},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Both},
		},

		// x < 3 and x > 3 do not really overlap
		{
			&ScanRange{Low: nil, High: catalog.LookupValue{dparval.NewValue(3.0)}, Inclusion: catalog.Low},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: nil, Inclusion: catalog.High},
			nil,
		},

		// 3 < x <= 5, 4 <= x < 6, should become 4 <= x <= 5
		{
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.High},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(4.0)}, High: catalog.LookupValue{dparval.NewValue(6.0)}, Inclusion: catalog.Low},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(4.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Both},
		},

		// 3 < x < 5, 5 <= x < 6, do not really overlap
		{
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Neither},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(5.0)}, High: catalog.LookupValue{dparval.NewValue(6.0)}, Inclusion: catalog.Low},
			nil,
		},

		// 3 < x < 5 and 3 < x < 6, become 3 < x < 5
		{
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Neither},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(6.0)}, Inclusion: catalog.Neither},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(5.0)}, Inclusion: catalog.Neither},
		},

		// 3 < x < 9 and 1 < x < 6, become 3 < x < 6
		{
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(9.0)}, Inclusion: catalog.Neither},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(1.0)}, High: catalog.LookupValue{dparval.NewValue(6.0)}, Inclusion: catalog.Neither},
			&ScanRange{Low: catalog.LookupValue{dparval.NewValue(3.0)}, High: catalog.LookupValue{dparval.NewValue(6.0)}, Inclusion: catalog.Neither},
		},
	}

	for i, x := range tests {
		actual := x.a.Overlap(x.b)
		if !reflect.DeepEqual(actual, x.overlap) {
			t.Errorf("Expected %v, got %v for index %d", x.overlap, actual, i)
		}
	}
}
