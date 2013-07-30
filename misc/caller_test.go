//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package misc

import (
	"testing"
	"strings"
)

func TestCaller(t *testing.T) {
	c := CallerN(0)
	if !strings.HasPrefix(c, "caller_test") {
		t.Errorf("expected prefix: caller_test, got c: %v", c)
	}
	a := CallerN(0)
	b := CallerN(0)
	if a == b {
		t.Errorf("expected different linenums, got: %v vs %v", a, b)
	}
	c = Caller()
	if !strings.HasPrefix(c, "caller_test") {
		t.Errorf("expected prefix: caller_test, got c: %v", c)
	}
	a = Caller()
	b = Caller()
	if a == b {
		t.Errorf("expected different linenums, got: %v vs %v", a, b)
	}
	func() {
		x := Caller()
		if !strings.HasPrefix(c, "caller_test") {
			t.Errorf("expected caller_test, got: %v", x)
		}
	}()
}

func TestCallerLevel(t *testing.T) {
	testCallerLevel(t, 0, "caller_test")
	testCallerLevel(t, 1, "caller_test")
	testCallerLevel(t, 2, "testing")

	if CallerN(1000000) != "unknown:0" {
		t.Errorf("expected huge caller to be unknown:0")
	}
}

func testCallerLevel(t *testing.T, uplevel int, expect string) {
	c := CallerN(uplevel)
	if !strings.HasPrefix(c, expect + ":") {
		t.Errorf("expected prefix: %v, got c: %v", expect, c)
	}
}
