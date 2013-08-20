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
	"encoding/json"
	"reflect"
	"testing"

	"github.com/couchbaselabs/tuqtng/ast"
)

var validQueries = []string{
	// this section is all basic expression testing
	`SELECT null`,
	`SELECT NULL`,
	`SELECT NuLl`,
	`SELECT true`,
	`SELECT TRUE`,
	`SELECT tRuE`,
	`SELECT false`,
	`SELECT FALSE`,
	`SELECT fAlSe`,
	`SELECT 1`,
	`SELECT -3`,
	`SELECT 1.5`,
	`SELECT -3.14`,
	`SELECT 1e6`,
	`SELECT 1.3e23`,
	`SELECT -4e-4`,
	`SELECT []`,
	`SELECT [null, false, true, 7, 3.14, "bob"]`,
	`SELECT {}`,
	`SELECT {"bob": "wood"}`,
	`SELECT {"bob": 1}`,
	`SELECT {"null": null, "bool": true, "number": 7, "array": [2, 3, "cat"], "object": {"nested": 99}}`,
	"SELECT `bob wood`",
	`SELECT 3 > 7`,
	`SELECT 1 < 2`,
	`SELECT 1 == 3`,
	`SELECT 2 = 4`,
	`SELECT 3 <> cat`,
	`SELECT dog != cat`,
	`SELECT wow >= what`,
	`SELECT [] <= 7`,
	`SELECT a + b`,
	`SELECT d - c`,
	`SELECT x * y`,
	`SELECT z / 2`,
	`SELECT minutes % 60`,
	`SELECT str1 || str2`,
	`SELECT (3 + c) * 7`,
	`SELECT cat IS NULL`,
	`SELECT dog IS NOT NULL`,
	`SELECT steve IS MISSING`,
	`SELECT gerald IS NOT MISSING`,
	`SELECT siri IS VALUED`,
	`SELECT marty IS NOT VALUED`,
	`SELECT noone LIKE them`,
	`SELECT someone NOT LIKE me`,
	`SELECT -abv`,
	`SELECT contact.name`,
	`SELECT contact.name.firstn`,
	`SELECT {"bob": "wood"}.wood`,
	`SELECT family["father"]`,
	`SELECT [a,b,c][0]`,
	`SELECT [0,[1,2,[3,4,5]]][1][2][1]`,
	`SELECT cat.a+b`,  // DOT has higher precedance than PLUS so this is valid
	`SELECT cat[a+b]`, // if you wanted to evaluate a+b and use that as property name
	// more complicated query testing
	`SELECT bob FROM cat WHERE foo = bar ORDER BY this LIMIT 10 OFFSET 20`,
	`SELECT bob FROM cat WHERE foo = bar LIMIT 10 OFFSET 20`,
	`SELECT bob FROM cat WHERE foo = bar LIMIT 10`,
	`SELECT bob FROM cat WHERE foo = bar`,
	`SELECT bob FROM cat`,
	`SELECT bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT bob, bill FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT bob AS bill, bill AS bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT *, bob AS bill, bill AS bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT *, names.*, bob AS bill, bill AS bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT *, names.*, bob AS bill, bill AS bob FROM cat WHERE foo = bar or 3 > 4`,
	`SELECT * FROM cat WHERE true ORDER BY name ASC`,
	`SELECT * FROM cat WHERE true ORDER BY name DESC`,
	`EXPLAIN SELECT * FROM cat WHERE true ORDER BY name DESC`,
	`SELECT Color(metal) FROM materials AS metal WHERE Conductivity(metal, 'heat') > 3.5`,
	`SELECT * FROM bucket WHERE CASE WHEN abv > 5 THEN true END`,
	`SELECT * FROM bucket WHERE CASE WHEN abv > 5 THEN true WHEN abv > 3 THEN false END`,
	`SELECT * FROM bucket WHERE CASE WHEN abv > 5 THEN true ELSE false END`,
	`SELECT * FROM bucket WHERE CASE WHEN abv > 5 THEN true WHEN abv > 3 THEN false ELSE 7 END`,
	`SELECT * FROM bucket WHERE CASE WHEN abv IS NOT NULL THEN abv ELSE ibu END > 0`,
	`SELECT * FROM bucket WHERE ANY child.age > 25 OVER children AS child`,
	`SELECT * FROM bucket WHERE ALL child.age > 25 OVER children AS child`,
	`SELECT * FROM bucket WHERE ALL child.age > 25 OVER children[0] AS child`,
	`SELECT * FROM bucket WHERE ALL child.age > 25 OVER a.children[0] AS child`,
	`SELECT * FROM bucket WHERE ALL child.age > 25 OVER b.a.children[0] AS child`,
	`SELECT * FROM bucket WHERE ALL child.age > 25 OVER b.a.children AS child`,
	`SELECT * FROM bucket WHERE ALL child.age > 25 OVER some.my.children AS child`,
	"SELECT * FROM bucket WHERE ALL child.age > 25 OVER `all`.my.children AS child",
	`SELECT * FROM bucket WHERE ANY child.age > 25 OVER children AS child AND age > 65 OR color = "blue"`,
	`SELECT * FROM bucket AS buck`,
	`SELECT * FROM bucket AS buck OVER buck.addresses`,
	`SELECT * FROM bucket AS buck OVER buck.addresses AS address`,
	`SELECT * FROM bucket AS buck OVER buck.addresses AS address OVER address.lines`,
	`SELECT * FROM bucket AS buck OVER buck.addresses AS address OVER address.lines AS line`,
	`SELECT DISTINCT name from bucket`,
	`SELECT UNIQUE name from bucket`,
	`SELECT COUNT(DISTINCT name) FROM bucket`,
	`SELECT COUNT(UNIQUE name) FROM bucket`,
	`SELECT * FROM contacts WHERE ANY child.age > 14 AND child.gender IS NOT NULL OVER children AS child`,
	`SELECT * FROM bucket WHERE FIRST child.age > 25 OVER children AS child`,
	`CREATE INDEX abv_idx ON beer-sample(abv)`,
	`CREATE INDEX abv_idx ON beer-sample(abv) USING VIEW`,
	`CREATE INDEX abv_idx ON beer-sample(abv) USING magic`,
	`CREATE INDEX abv_idx ON beer-sample(abv, ibu)`,
	`CREATE INDEX abv_idx ON beer-sample(abv, ibu) USING VIEW`,
	`CREATE INDEX abv_idx ON beer-sample(abv, ibu) USING magic`,
	`SELECT ARRAY child.name WHEN child.age > 20 OVER contacts.children AS child FROM contacts`,
}

var invalidQueries = []string{
	`bob`,         // must have select
	`SELECT 01`,   // numbers cannot start with leading zeros
	`SELECT 3dog`, // unescaped identifiers cannot start with number
	`SELECT bob FROM cat WHERE foo = bar ORDER BY this OFFSET 20`,                  // offset requires limit
	`SELECT * AS all, bob AS bill, bill AS bob FROM cat WHERE foo = bar and 3 > 4`, // cannot alias *
	`SELECT * WHERE true AND`,
	`SELECT * WHERE true ORDER BY DESC`,
	`SELECT "a`,
	`CREATE *`,
	`CREATE INDEX abv`,
	`CREATE INDEX abv ON beer-sample`,
	`CREATE INDEX abv ON beer-sample()`,
	`CREATE INDEX abv USING VIEW`,
	`CREATE INDEX abv ON beer-sample USING VIEW`,
	`CREATE INDEX abv ON beer-sample() USING VIEW`,
	`CREATE INDEX abv USING magic`,
	`CREATE INDEX abv ON beer-sample USING magic`,
	`CREATE INDEX abv ON beer-sample() USING magic`,
	// these are me trying to understand code coverage in the parser
	`\`,
	`\r`,
	`\u`,
	`"\`,
	`"t`,
	`"r`,
	`"u`,
	`""`,
	`"/`,
	`"//`,
	`"/n`,
	`"/f`,
	`"/\`,
	`"\/`,
	`"\n`,
	`"\t`,
	`"\r`,
	`"\u`,
}

func TestParser(t *testing.T) {
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

func TestParserASTOutput(t *testing.T) {

	tests := []struct {
		input  string
		output ast.Statement
	}{
		{"SELECT * FROM test WHERE true",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewStarResultExpression(),
				},
				Distinct: false,
				From:     &ast.From{Projection: ast.NewProperty("test")},
				Where:    ast.NewLiteralBool(true),
				Limit:    -1,
			},
		},
		{"SELECT * FROM test ORDER BY foo",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewStarResultExpression(),
				},
				Distinct: false,
				From:     &ast.From{Projection: ast.NewProperty("test")},
				Where:    nil,
				OrderBy: []*ast.SortExpression{
					ast.NewSortExpression(ast.NewProperty("foo"), true),
				},
				Limit: -1,
			},
		},
		{"SELECT * FROM test LIMIT 10 OFFSET 3",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewStarResultExpression(),
				},
				Distinct: false,
				From:     &ast.From{Projection: ast.NewProperty("test")},
				Where:    nil,
				Limit:    10,
				Offset:   3,
			},
		},
		{"SELECT a FROM test",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewResultExpression(ast.NewProperty("a")),
				},
				Distinct: false,
				From:     &ast.From{Projection: ast.NewProperty("test")},
				Where:    nil,
				Limit:    -1,
			},
		},
		{"SELECT 1+1*30 as steve",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewResultExpressionWithAlias(ast.NewPlusOperator(ast.NewLiteralNumber(1.0), ast.NewMultiplyOperator(ast.NewLiteralNumber(1.0), ast.NewLiteralNumber(30.0))), "steve"),
				},
				Distinct: false,
				From:     nil,
				Where:    nil,
				Limit:    -1,
			},
		},
		{"SELECT DISTINCT 1+1*30 as steve",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewResultExpressionWithAlias(ast.NewPlusOperator(ast.NewLiteralNumber(1.0), ast.NewMultiplyOperator(ast.NewLiteralNumber(1.0), ast.NewLiteralNumber(30.0))), "steve"),
				},
				Distinct: true,
				From:     nil,
				Where:    nil,
				Limit:    -1,
			},
		},
		{"CREATE INDEX abv_idx ON beer-sample(abv) USING VIEW",
			&ast.CreateIndexStatement{
				Method: "view",
				Bucket: "beer-sample",
				Name:   "abv_idx",
				On:     ast.ExpressionList{ast.NewProperty("abv")},
			},
		},
		{"CREATE INDEX abv_idx ON beer-sample(abv) USING magic",
			&ast.CreateIndexStatement{
				Method: "magic",
				Bucket: "beer-sample",
				Name:   "abv_idx",
				On:     ast.ExpressionList{ast.NewProperty("abv")},
			},
		},
		{"CREATE INDEX abv_idx ON beer-sample(abv)",
			&ast.CreateIndexStatement{
				Method: "",
				Bucket: "beer-sample",
				Name:   "abv_idx",
				On:     ast.ExpressionList{ast.NewProperty("abv")},
			},
		},
	}

	unqlParser := NewUnqlParser()

	for _, v := range tests {
		query, err := unqlParser.Parse(v.input)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}

		if !reflect.DeepEqual(query, v.output) {
			t.Errorf("Expected %v, got %v", v.output, query)

			js, err := json.MarshalIndent(v.output, "", "  ")
			if err == nil {
				t.Logf("Expected %v", string(js))
			}

			js, err = json.MarshalIndent(query, "", "  ")
			if err == nil {
				t.Logf("Got %v", string(js))
			}

		}
	}

}
