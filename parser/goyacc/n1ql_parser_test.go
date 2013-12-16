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
	`SELECT bob bill, bill bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT *, bob AS bill, bill AS bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT *, names.*, bob AS bill, bill AS bob FROM cat WHERE foo = bar and 3 > 4`,
	`SELECT *, names.*, bob AS bill, bill AS bob FROM cat WHERE foo = bar or 3 > 4`,
	`SELECT * FROM cat WHERE true ORDER BY name ASC`,
	`SELECT * FROM cat WHERE true ORDER BY name DESC`,
	`EXPLAIN SELECT * FROM cat WHERE true ORDER BY name DESC`,
	`SELECT Color(metal) FROM materials AS metal WHERE Conductivity(metal, 'heat') > 3.5`,
	`SELECT * FROM abucket WHERE CASE WHEN abv > 5 THEN true END`,
	`SELECT * FROM abucket WHERE CASE WHEN abv > 5 THEN true WHEN abv > 3 THEN false END`,
	`SELECT * FROM abucket WHERE CASE WHEN abv > 5 THEN true ELSE false END`,
	`SELECT * FROM abucket WHERE CASE WHEN abv > 5 THEN true WHEN abv > 3 THEN false ELSE 7 END`,
	`SELECT * FROM abucket WHERE CASE WHEN abv IS NOT NULL THEN abv ELSE ibu END > 0`,
	// collection expressions
	`SELECT * FROM abucket WHERE ANY child IN children SATISFIES child.age > 25 END`,
	`SELECT * FROM abucket WHERE ANY children SATISFIES children.age END`,
	`SELECT * FROM abucket WHERE EVERY child IN children SATISFIES child.age > 25 END`,
	`SELECT * FROM abucket WHERE EVERY children SATISFIES children.age > 25 END`,
	`SELECT * FROM abucket WHERE EVERY child IN children[0] SATISFIES child.age > 25 END`,
	`SELECT * FROM abucket WHERE EVERY child IN a.children[0] SATISFIES child.age > 25 END`,
	`SELECT * FROM abucket WHERE EVERY child IN b.a.children[0] SATISFIES child.age > 25 END`,
	`SELECT * FROM abucket WHERE EVERY child IN b.a.children SATISFIES child.age > 25 END`,
	`SELECT * FROM abucket WHERE EVERY child IN some.my.children SATISFIES child.age > 25 END`,
	"SELECT * FROM abucket WHERE EVERY child IN `all`.my.children SATISFIES child.age > 25 END",
	`SELECT * FROM abucket WHERE ANY   child IN children SATISFIES child.age > 65 END AND age > 65 OR color = "blue"`,

	// from clause
	`SELECT * FROM abucket AS buck`,
	`SELECT * FROM abucket buck`,
	`SELECT * FROM abucket AS buck OVER address IN buck.addresses`,
	`SELECT * FROM abucket AS buck OVER address IN buck.addresses OVER line IN address.lines`,
	`SELECT * FROM abucket AS buck OVER buck.addresses`,

	`SELECT DISTINCT name from abucket`,
	`SELECT UNIQUE name from abucket`,
	`SELECT COUNT(DISTINCT name) FROM abucket`,
	`SELECT COUNT(UNIQUE name) FROM abucket`,
	`SELECT * FROM contacts WHERE ANY child IN children SATISFIES child.age > 14 AND child.gender IS NOT NULL END`,
	`SELECT * FROM abucket WHERE FIRST child FOR child IN children WHEN child.age > 25 END`,
	`SELECT * FROM abucket WHERE FIRST children IN children WHEN children.age > 25 END`,
	`SELECT ARRAY child.name FOR child IN contacts.children WHEN child.age > 20 END FROM contacts`,
	`SELECT * FROM :apool.abucket`,
	`SELECT * FROM :apool.abucket.prop`,
	`SELECT * FROM :apool.abucket[3]`,
	`SELECT * FROM :apool.abucket.prop[4]`,

	// don's queries from the cheat sheet
	`SELECT * FROM organizations`,
	`SELECT * FROM tutorial WHERE name = 'Don'`,
	`SELECT children[0].name AS cname FROM tutorial WHERE name = 'Don'`,
	`SELECT name, age, age/7 AS age_pet_years FROM tutorial WHERE name = 'Don'`,
	`SELECT name, age FROM tutorial WHERE age >= 30`,
	`EXPLAIN SELECT * FROM organizations`,
	`SELECT * FROM tutorial WHERE CASE WHEN true THEN 7 ELSE 9 END`,
	`SELECT name, email FROM tutorial WHERE email LIKE '%@gmail.com'`,
	`SELECT name, email FROM tutorial WHERE email NOT LIKE '%@gmail.com'`,
	`SELECT name FROM tutorial ORDER BY name`,
	`SELECT name FROM tutorial ORDER BY name LIMIT 2`,
	`SELECT name FROM tutorial ORDER BY name LIMIT 2 OFFSET 2`,
	`SELECT DISTINCT organization.address FROM organizations WHERE address.city = "Mountain View"`,
	`SELECT LENGTH(some_array) FROM tutorial`,
	`SELECT some_array[0] FROM tutorial`,
	`SELECT * FROM organizations AS organization OVER employee IN organization.employees`,
	`SELECT * FROM tutorial WHERE ANY name IN collection SATISFIES predicate END`,
	`SELECT * FROM tutorial WHERE FIRST expression FOR name IN collection END`,
	`SELECT * FROM tutorial WHERE FIRST collection IN collection END`,
	`SELECT * FROM tutorial WHERE FIRST expression FOR name IN collection WHEN predicate END`,
	`SELECT * FROM tutorial WHERE ARRAY expression FOR name IN collection END`,
	`SELECT * FROM tutorial WHERE ARRAY expression FOR name IN collection WHEN predicate END`,
	`SELECT relation, COUNT(*) AS count FROM tutorial GROUP BY relation`,
	`SELECT relation, COUNT(*) AS count FROM tutorial GROUP BY relation HAVING COUNT(*) > 1`,
	`SELECT SUM(employee.salary) AS total FROM emp`,
	`SELECT AVG(employee.age) AS av_age FROM emp`,
	`SELECT MIN(employee.age) AS min_age FROM emp`,
	`SELECT MAX(employee.age) AS max_age FROM emp`,

	// index statements
	`CREATE INDEX abv_idx ON beer-sample(abv)`,
	`CREATE INDEX abv_idx ON beer-sample(abv) USING VIEW`,
	`CREATE INDEX abv_idx ON beer-sample(abv) USING magic`,
	`CREATE INDEX abv_idx ON :apool.beer-sample(abv)`,
	`CREATE INDEX abv_idx ON :apool.beer-sample(abv) USING VIEW`,
	`CREATE INDEX abv_idx ON :apool.beer-sample(abv) USING magic`,
	`CREATE INDEX abv_idx ON beer-sample(abv, ibu)`,
	`CREATE INDEX abv_idx ON beer-sample(abv, ibu) USING VIEW`,
	`CREATE INDEX abv_idx ON beer-sample(abv, ibu) USING magic`,
	`CREATE INDEX abv_idx ON :apool.beer-sample(abv, ibu)`,
	`CREATE INDEX abv_idx ON :apool.beer-sample(abv, ibu) USING VIEW`,
	`CREATE INDEX abv_idx ON :apool.beer-sample(abv, ibu) USING magic`,
	`CREATE INDEX name ON abucket(field) USING VIEW`,
	`CREATE INDEX name ON abucket(field1, field2) USING VIEW`,
	`CREATE INDEX name ON abucket(field)`,
	`CREATE PRIMARY INDEX ON beer-sample`,
	`CREATE PRIMARY INDEX ON beer-sample USING VIEW`,
	`CREATE PRIMARY INDEX ON beer-sample USING magic`,
	`CREATE PRIMARY INDEX ON :apool.beer-sample`,
	`CREATE PRIMARY INDEX ON :apool.beer-sample USING VIEW`,
	`CREATE PRIMARY INDEX ON :apool.beer-sample USING magic`,
	`DROP INDEX abucket.name`,
	`DROP INDEX :apool.abucket.name`,
}

var invalidQueries = []string{
	`bob`,                                                                          // must have select
	`SELECT 01`,                                                                    // numbers cannot start with leading zeros
	`SELECT 3dog AS y`,                                                             // unescaped identifiers cannot start with number
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
	`DROP INDEX name`,
	`DROP INDEX bucket.*`,
	`CREATE PRIMARY INDEX abv_idx ON beer-sample(abv)`,
	`CREATE PRIMARY INDEX ON beer-sample(abv) USING VIEW`,
	`DROP PRIMARY INDEX`,

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
	n1qlParser := NewN1qlParser()

	for _, v := range validQueries {
		_, err := n1qlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
	}

	for _, v := range invalidQueries {
		_, err := n1qlParser.Parse(v)
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
		{"SELECT * FROM :apool.test WHERE true",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewStarResultExpression(),
				},
				Distinct: false,
				From:     &ast.From{Pool: "apool", Projection: ast.NewProperty("test")},
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
		{"SELECT a FROM test t2",
			&ast.SelectStatement{
				Select: ast.ResultExpressionList{
					ast.NewResultExpression(ast.NewProperty("a")),
				},
				Distinct: false,
				From:     &ast.From{Projection: ast.NewProperty("test"), As: "t2"},
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
		{"SELECT 1+1*30 steve",
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
		{"CREATE INDEX abv_idx ON :apool.beer-sample(abv) USING VIEW",
			&ast.CreateIndexStatement{
				Method: "view",
				Pool:   "apool",
				Bucket: "beer-sample",
				Name:   "abv_idx",
				On:     ast.ExpressionList{ast.NewProperty("abv")},
			},
		},
		{"CREATE INDEX abv_idx ON :apool.beer-sample(abv) USING magic",
			&ast.CreateIndexStatement{
				Method: "magic",
				Pool:   "apool",
				Bucket: "beer-sample",
				Name:   "abv_idx",
				On:     ast.ExpressionList{ast.NewProperty("abv")},
			},
		},
		{"CREATE INDEX abv_idx ON :apool.beer-sample(abv)",
			&ast.CreateIndexStatement{
				Method: "",
				Pool:   "apool",
				Bucket: "beer-sample",
				Name:   "abv_idx",
				On:     ast.ExpressionList{ast.NewProperty("abv")},
			},
		},
		{"DROP INDEX beer-sample.abv",
			&ast.DropIndexStatement{
				Bucket: "beer-sample",
				Name:   "abv",
			},
		},
		{"DROP INDEX :apool.beer-sample.abv",
			&ast.DropIndexStatement{
				Pool:   "apool",
				Bucket: "beer-sample",
				Name:   "abv",
			},
		},
	}

	n1qlParser := NewN1qlParser()

	for _, v := range tests {
		query, err := n1qlParser.Parse(v.input)
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
