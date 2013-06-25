//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package goyacc

import (
	"testing"
)

var validQueries = []string{
	`null`,
	`NULL`,
	`NuLl`,
	`true`,
	`TRUE`,
	`tRuE`,
	`false`,
	`FALSE`,
	`fAlSe`,
	`1`,
	`-3`,
	`1.5`,
	`-3.14`,
	`1e6`,
	`1.3e23`,
	`-4e-4`,
	`[]`,
	`[null, false, true, 7, 3.14, "bob"]`,
	`{}`,
	`{"bob": "wood"}`,
	`{"bob": 1}`,
	`{"null": null, "bool": true, "number": 7, "array": [2, 3, "cat"], "object": {"nested": 99}}`,
	"`bob wood`",
	`3 > 7`,
	`1 < 2`,
	`1 == 3`,
	`2 = 4`,
	`3 <> cat`,
	`dog != cat`,
	`wow >= what`,
	`[] <= 7`,
	`a + b`,
	`d - c`,
	`x * y`,
	`z / 2`,
	`str1 || str2`,
	`(3 + c) * 7`,
}

var invalidQueries = []string{}

func TestParser(t *testing.T) {
	DebugTokens = true
	DebugGrammar = true
	unqlParser := NewUnqlParser()

	for _, v := range validQueries {
		_, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
	}

	for _, v := range invalidQueries {
		_, err := unqlParser.Parse(v)
		if err == nil {
			t.Errorf("Invalid Query Parsed Successfully: %v - %v", v, err)
		}
	}

}
