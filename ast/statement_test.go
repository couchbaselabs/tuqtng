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

func TestSelectStatement(t *testing.T) {
	stmt := NewSelectStatement()

	stmt.Distinct = false
	stmt.ExplainOnly = true
	stmt.Select = ResultExpressionList{NewStarResultExpression(), NewDotStarResultExpression(NewProperty("foo")), NewResultExpression(NewProperty("bar"))}
	stmt.From = &From{Projection: NewProperty("test")}
	stmt.Where = NewLiteralBool(true)
	stmt.GroupBy = ExpressionList{NewProperty("bar")}
	stmt.Having = NewLiteralBool(true)
	stmt.OrderBy = []*SortExpression{NewSortExpression(NewProperty("foo"), true)}
	stmt.Limit = 10
	stmt.Offset = 5

	stmt.From.ConvertToBucketFrom()

	if !stmt.IsExplainOnly() {
		t.Errorf("Expected query to be explain only")
	}
	if stmt.IsDistinct() {
		t.Errorf("Expected query to not be distinct")
	}
	if !reflect.DeepEqual(stmt.GetResultExpressionList(), ResultExpressionList{NewStarResultExpression(), NewDotStarResultExpression(NewProperty("foo")), NewResultExpression(NewProperty("bar"))}) {
		t.Errorf("Expected star result expression")
	}
	if stmt.GetFrom().Bucket != "test" {
		t.Errorf("Expected from test")
	}
	if !reflect.DeepEqual(stmt.GetWhere(), NewLiteralBool(true)) {
		t.Errorf("Expected where true")
	}
	if !reflect.DeepEqual(stmt.GetOrderBy(), SortExpressionList{NewSortExpression(NewProperty("foo"), true)}) {
		t.Errorf("Expected order by foo ascending")
	}
	if !reflect.DeepEqual(stmt.GetGroupBy(), ExpressionList{NewProperty("bar")}) {
		t.Errorf("Expected group by bar")
	}
	if !reflect.DeepEqual(stmt.GetHaving(), NewLiteralBool(true)) {
		t.Errorf("Expected having true")
	}
	if stmt.GetLimit() != 10 {
		t.Errorf("Expected limit 10")
	}
	if stmt.GetOffset() != 5 {
		t.Errorf("Expected offset 5")
	}

	stmt.SetExplainOnly(false)

	if stmt.IsExplainOnly() {
		t.Errorf("Expected query to not be explain only")
	}
}

func TestSelectStatementWithDuplicateAlias(t *testing.T) {
	stmt := NewSelectStatement()

	stmt.Select = ResultExpressionList{NewResultExpressionWithAlias(NewProperty("foo"), "foo"), NewResultExpressionWithAlias(NewProperty("foo"), "bar")}

	err := stmt.VerifySemantics()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	stmt.Select = ResultExpressionList{NewResultExpressionWithAlias(NewProperty("foo"), "foo"), NewResultExpressionWithAlias(NewProperty("foo"), "foo")}

	err = stmt.VerifySemantics()
	if err == nil || err.Error() != "alias foo is defined more than once" {
		t.Errorf("expected error: alias foo is defined more than once")
	}
}

func TestSelectStatementDefaultNaming(t *testing.T) {
	stmt := NewSelectStatement()

	stmt.Select = ResultExpressionList{NewResultExpression(NewProperty("foo")), NewResultExpression(NewProperty("foo"))}

	err := stmt.VerifySemantics()
	if err != nil {
		t.Errorf("exepect no errors verifying semantics, got %v", err)
	}

	if stmt.Select[0].As != "foo" {
		t.Errorf("Expected alias ot be foo, got %v", stmt.Select[0].As)
	}

	if stmt.Select[1].As != "$1" {
		t.Errorf("Expected alias to be $1, got %v", stmt.Select[1].As)
	}
}

func TestSelectStatementVerifyFormalNotation(t *testing.T) {

	goodSelectStmt := NewSelectStatement()
	goodSelectStmt.Select = ResultExpressionList{NewResultExpression(NewBracketMemberOperator(NewProperty("bucket"), NewProperty("name")))}
	goodSelectStmt.From = &From{Bucket: "bucket", As: "bucket"}
	goodSelectStmt.Where = NewBracketMemberOperator(NewProperty("bucket"), NewProperty("name"))

	tests := []struct {
		input  *SelectStatement
		output error
	}{
		{goodSelectStmt, nil},
	}

	for _, test := range tests {
		err := test.input.verifyFormalNotation()
		if !reflect.DeepEqual(err, test.output) {
			t.Errorf("Expected error %v, got %v", test.output, err)
		}
	}

}
