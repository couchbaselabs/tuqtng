//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package mock

import (
	"strconv"
	"testing"
)

func TestMock(t *testing.T) {
	s, err := NewSite("mock:")
	if err != nil {
		t.Errorf("failed to create site: %v", err)
	}
	if s.URL() != "mock:" {
		t.Errorf("expected site URL to be same")
	}
	n, err := s.PoolNames()
	if err != nil || len(n) != DEFAULT_NUM_POOLS {
		t.Errorf("expected num pools to be same")
	}
	p, err := s.Pool("not-a-pool")
	if err == nil || p != nil {
		t.Errorf("expected not-a-pool")
	}
	p, err = s.Pool("p0")
	if err != nil || p == nil {
		t.Errorf("expected pool p0")
	}
	if p.Name() != "p0" {
		t.Errorf("expected p0 name")
	}
	n, err = p.BucketNames()
	if err != nil || len(n) != DEFAULT_NUM_BUCKETS {
		t.Errorf("expected num buckets to be same")
	}
	b, err := p.Bucket("not-a-bucket")
	if err == nil || b != nil {
		t.Errorf("expected not-a-bucket")
	}
	b, err = p.Bucket("b0")
	if err != nil || b == nil {
		t.Errorf("expected bucket b0")
	}
	if b.Name() != "b0" {
		t.Errorf("expected b0 name")
	}
	c, err := b.Count()
	if err != nil || c != int64(DEFAULT_NUM_ITEMS) {
		t.Errorf("expected num items")
	}
	v, err := b.Fetch("123")
	if err != nil || v == nil {
		t.Errorf("expected item 123")
	}
	x, gerr := v.GetPath("id")
	if gerr != nil || x == nil {
		t.Errorf("expected item.id")
	}
	x, gerr = v.GetPath("i")
	if gerr != nil || x == nil {
		t.Errorf("expected item.i")
	}
	x, gerr = v.GetPath("not-a-valid-path")
	if gerr == nil || x != nil {
		t.Errorf("expected not-a-valid-path to err")
	}
	v, err = b.Fetch("not-an-item")
	if err == nil || v != nil {
		t.Errorf("expected not-an-item")
	}
	v, err = b.Fetch(strconv.Itoa(DEFAULT_NUM_ITEMS))
	if err == nil || v != nil {
		t.Errorf("expected not-an-item")
	}
}
