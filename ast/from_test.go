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
	"reflect"
	"testing"
)

func TestFromGenerateAlias(t *testing.T) {

	tests := []struct {
		input  *From
		output *From
	}{
		{
			&From{
				Bucket:     "bucket",
				Projection: nil,
			},
			&From{
				Bucket:     "bucket",
				Projection: nil,
				As:         "bucket",
			},
		},
		{
			&From{
				Bucket:     "bucket",
				Projection: NewProperty("b"),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewProperty("b"),
				As:         "b",
			},
		},
		{
			&From{
				Bucket:     "bucket",
				Projection: NewDotMemberOperator(NewProperty("b"), NewProperty("c")),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewDotMemberOperator(NewProperty("b"), NewProperty("c")),
				As:         "c",
			},
		},
		{
			&From{
				Bucket:     "bucket",
				Projection: NewBracketMemberOperator(NewProperty("b"), NewLiteralNumber(0.0)),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewBracketMemberOperator(NewProperty("b"), NewLiteralNumber(0.0)),
			},
		},
	}

	for _, x := range tests {
		x.input.GenerateAlias()
		if !reflect.DeepEqual(x.input, x.output) {
			t.Errorf("Expected %v got %v", x.output, x.input)
		}
	}
}

func TestGetAliases(t *testing.T) {
	tests := []struct {
		input  *From
		output []string
	}{
		{&From{Bucket: "bucket", Projection: NewProperty("b"), As: "b"}, []string{"b"}},
		{
			&From{
				Bucket:     "bucket",
				Projection: NewProperty("b"),
				As:         "b",
				Over: &From{
					Projection: NewProperty("c"),
					As:         "c",
				},
			},
			[]string{"b", "c"},
		},
	}

	for _, x := range tests {
		aliases := x.input.GetAliases()
		if !reflect.DeepEqual(aliases, x.output) {
			t.Errorf("Expected %v for %v, got %v", x.output, x.input, aliases)
		}
	}
}

func TestConvertToBucketFrom(t *testing.T) {
	tests := []struct {
		input  *From
		output *From
	}{
		// bucket should become "bucket"
		{
			&From{
				Projection: NewProperty("bucket"),
			},
			&From{
				Bucket:     "bucket",
				Projection: nil,
			},
		},
		// bucket.nest should become "bucket" and project "nest"
		{
			&From{
				Projection: NewDotMemberOperator(NewProperty("bucket"), NewProperty("nest")),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewProperty("nest"),
			},
		},
		// bucket.nest.nest2 "bucket" and project "nest.nest2"
		{
			&From{
				Projection: NewDotMemberOperator(NewDotMemberOperator(NewProperty("bucket"), NewProperty("nest")), NewProperty("nest2")),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewDotMemberOperator(NewProperty("nest"), NewProperty("nest2")),
			},
		},
		// bucket[0].nest should become "bucket" and project VALUE()[0].nest
		{
			&From{
				Projection: NewDotMemberOperator(NewBracketMemberOperator(NewProperty("bucket"), NewLiteralNumber(0.0)), NewProperty("nest")),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewDotMemberOperator(NewBracketMemberOperator(NewProperty("bucket"), NewLiteralNumber(0.0)), NewProperty("nest")),
			},
		},
		// bucket[0] should be come "bucket" and project VALUE()[0]
		{
			&From{
				Projection: NewBracketMemberOperator(NewProperty("bucket"), NewLiteralNumber(0.0)),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewBracketMemberOperator(NewProperty("bucket"), NewLiteralNumber(0.0)),
			},
		},
		// bucket[0][1] should be come "bucket" and project VALUE()[0][1]
		{
			&From{
				Projection: NewBracketMemberOperator(NewBracketMemberOperator(NewProperty("bucket"), NewLiteralNumber(0.0)), NewLiteralNumber(1.0)),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewBracketMemberOperator(NewBracketMemberOperator(NewProperty("bucket"), NewLiteralNumber(0.0)), NewLiteralNumber(1.0)),
			},
		},
		// bucket.nest[0] should become "bucket" and project "nest[0]"
		{
			&From{
				Projection: NewBracketMemberOperator(NewDotMemberOperator(NewProperty("bucket"), NewProperty("nest")), NewLiteralNumber(0.0)),
			},
			&From{
				Bucket:     "bucket",
				Projection: NewBracketMemberOperator(NewProperty("nest"), NewLiteralNumber(0.0)),
			},
		},
	}

	for _, x := range tests {
		x.input.ConvertToBucketFrom()
		if !reflect.DeepEqual(x.input, x.output) {
			t.Errorf("Expected %v got %v", x.output, x.input)
		}
	}
}
