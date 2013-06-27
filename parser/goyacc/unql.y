%{
package goyacc
import "log"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
    if DebugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if DebugGrammar {
        log.Printf("DEBUG GRAMMAR " + format)
    }
}
%}

%union { 
s string 
n int
f float64}

%token SELECT AS FROM WHERE
%token ORDER BY ASC DESC
%token LIMIT OFFSET
%token LBRACE RBRACE LBRACKET RBRACKET
%token COMMA COLON DQUOTE
%token TRUE FALSE NULL
%token INT NUMBER IDENTIFIER STRING
%token PLUS MINUS MULT DIV
%token CONCAT
%token AND OR NOT
%token EQ NE GT GTE LT LTE
%token LPAREN RPAREN
%token LIKE IS VALUED MISSING
%token DOT
%left DOT LBRACKET
%left OR
%left AND 
%left EQ LT LTE GT GTE NE LIKE
%left PLUS MINUS MULT DIV MOD CONCAT
%right NOT

%%

input: 
select_stmt { 
	logDebugGrammar("INPUT") 
}
;

//STATEMENT
select_stmt:
select_compound select_order select_limit_offset {
	logDebugGrammar("SELECT_STMT")
}
;

select_compound:    
select_core { 
	// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
}
;

select_core:    
select_select select_from select_where { 
	logDebugGrammar("SELECT_CORE")
}
;

select_select:  
select_select_head select_select_tail {
	logDebugGrammar("SELECT_SELECT")
}
;

select_select_head:  
SELECT { 
	logDebugGrammar("SELECT_SELECT_HEAD")
	parsingStatement = ast.NewSelectStatement()
}
;

select_select_tail:		
result_list { 
	logDebugGrammar("SELECT SELECT TAIL - EXPR")
}
;

result_list:
result_single {
	
}
|
result_single COMMA result_list {
	
};

result_single:
dotted_path_star {
	logDebugGrammar("RESULT STAR")
}
|
expression { 
	logDebugGrammar("RESULT EXPR")
}
|
expression AS IDENTIFIER { 
	logDebugGrammar("SORT EXPR ASC")
}
;

dotted_path_star:
MULT {
	logDebugGrammar("STAR")
}
|
expr DOT MULT {
	logDebugGrammar("PATH DOT STAR")
}

select_from:
/* empty */ {
	logDebugGrammar("SELECT FROM - EMPTY")
}
|
FROM data_source {
	logDebugGrammar("SELECT FROM - DATASOURCE")
}

data_source:
IDENTIFIER {
	logDebugGrammar("FROM DATASOURCE")
}
|
IDENTIFIER AS IDENTIFIER {
    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
}

select_where:   
/* empty */ { 
	logDebugGrammar("SELECT WHERE - EMPTY")
}
|
WHERE expression {
	logDebugGrammar("SELECT WHERE - EXPR")
	where_part := parsingStack.Pop().(ast.Expression)
	switch parsingStatement := parsingStatement.(type) {
	case *ast.SelectStatement:
		parsingStatement.Where = where_part
	default:
		logDebugGrammar("This statement does not support WHERE")
	}
};

select_order:   
/* empty */
|
ORDER BY sorting_list {
	
}
;

sorting_list:
sorting_single {
	
}
|
sorting_single COMMA sorting_list {
	
};

sorting_single:
expression { 
	logDebugGrammar("SORT EXPR")
}
|
expression ASC { 
	logDebugGrammar("SORT EXPR ASC")
}
|
expression DESC { 
	logDebugGrammar("SORT EXPR DESC")
};

select_limit_offset:
/* empty */ {
	
}
|
select_limit {
	
}
|
select_limit select_offset {
	
}
;

select_limit:
LIMIT INT {
	logDebugGrammar("LIMIT")
};

select_offset:
OFFSET INT { 
	logDebugGrammar("OFFSET")
};

//EXPRESSION

expression:
expr {
	logDebugGrammar("EXPRESSION")
};

expr:
expr PLUS expr {
	logDebugGrammar("EXPR - PLUS")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr MINUS expr {
	logDebugGrammar("EXPR - MINUS")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr MULT expr {
	logDebugGrammar("EXPR - MULT")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr DIV expr {
	logDebugGrammar("EXPR - DIV")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr MOD expr {
	logDebugGrammar("EXPR - MOD")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr CONCAT expr {
	logDebugGrammar("EXPR - CONCAT")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr AND expr {
	logDebugGrammar("EXPR - AND")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
	parsingStack.Push(thisExpression)
}
|
expr OR expr {
	logDebugGrammar("EXPR - OR")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
	parsingStack.Push(thisExpression)
}
|
expr EQ expr {
	logDebugGrammar("EXPR - EQ")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr LT expr {
	logDebugGrammar("EXPR - LT")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr LTE expr {
	logDebugGrammar("EXPR - LTE")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr GT expr {
	logDebugGrammar("EXPR - GT")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr GTE expr {
	logDebugGrammar("EXPR - GTE")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr NE expr {
	logDebugGrammar("EXPR - NE")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr LIKE expr {
	logDebugGrammar("EXPR - LIKE")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr NOT LIKE expr {
	logDebugGrammar("EXPR - NOT LIKE")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
expr DOT IDENTIFIER {
	logDebugGrammar("EXPR DOT MEMBER")
	right := ast.NewProperty($3.s) 
	left := parsingStack.Pop()
	thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
	parsingStack.Push(thisExpression)
}
|
expr LBRACKET expr RBRACKET {
	logDebugGrammar("EXPR BRACKET MEMBER")
	right := parsingStack.Pop()
	left := parsingStack.Pop()
	thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
prefix_expr {
	
}
;

prefix_expr: 
NOT prefix_expr {
	logDebugGrammar("EXPR - NOT")
	operand := parsingStack.Pop()
	thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
MINUS prefix_expr {
	logDebugGrammar("EXPR - CHANGE SIGN")
	operand := parsingStack.Pop()
	thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
suffix_expr {
	
};

suffix_expr: 
atom {
	logDebugGrammar("SUFFIX_EXPR")
}
|
atom IS NULL {
	logDebugGrammar("SUFFIX_EXPR IS NULL")
	operand := parsingStack.Pop()
	thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
atom IS NOT NULL {
	logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
	operand := parsingStack.Pop()
	thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
atom IS MISSING {
	logDebugGrammar("SUFFIX_EXPR IS MISSING")
	operand := parsingStack.Pop()
	thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
atom IS NOT MISSING {
	logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
	operand := parsingStack.Pop()
	thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
atom IS VALUED {
	logDebugGrammar("SUFFIX_EXPR IS VALUED")
	operand := parsingStack.Pop()
	thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
|
atom IS NOT VALUED {
	logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
	operand := parsingStack.Pop()
	thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
	parsingStack.Push(thisExpression)
}
;

atom:
IDENTIFIER {
	thisExpression := ast.NewProperty($1.s) 
	parsingStack.Push(thisExpression) 
}
|
literal_value {
	logDebugGrammar("LITERAL")
}
|
LPAREN expression RPAREN {
	logDebugGrammar("NESTED EXPR")
}

//JSON

literal_value:
STRING {
	logDebugGrammar("STRING %s", $1.s)
	thisExpression := ast.NewLiteralString($1.s) 
	parsingStack.Push(thisExpression)
}
|
number {
	logDebugGrammar("NUMBER %f", $1.f)
	thisExpression := ast.NewLiteralNumber($1.f)
	parsingStack.Push(thisExpression)
}
|
object {
	logDebugGrammar("OBJECT")
}
|
array {
	logDebugGrammar("ARRAY")
}
|
TRUE {
	logDebugGrammar("TRUE")
	thisExpression := ast.NewLiteralBool(true) 
	parsingStack.Push(thisExpression)
}
|
FALSE {
	logDebugGrammar("FALSE")
	thisExpression := ast.NewLiteralBool(false) 
	parsingStack.Push(thisExpression)
}
|
NULL {
	logDebugGrammar("NULL")
	thisExpression := ast.NewLiteralNull()
	parsingStack.Push(thisExpression)
}
;

number:
INT {
	logDebugGrammar("NUMBER %d", $1.n)
}
|
NUMBER {
	logDebugGrammar("NUMBER %f", $1.f)
}
;

object:
LBRACE RBRACE {
	logDebugGrammar("EMPTY OBJECT")
}
|
LBRACE named_expression_list RBRACE {
	logDebugGrammar("OBJECT")
}
;

named_expression_list:
named_expression_single {
	logDebugGrammar("NAMED EXPR LIST SINGLE")
}
|
named_expression_single COMMA named_expression_list {
	logDebugGrammar("NAMED EXPR LIST COMPOUND")
	last := parsingStack.Pop().(*ast.LiteralObject)
	rest := parsingStack.Pop().(*ast.LiteralObject)
	for k,v := range last.Val {
		rest.Val[k] = v
	}
	parsingStack.Push(rest)
}
;

named_expression_single:   
STRING COLON expression {  
	logDebugGrammar("NAMED EXPR SINGLE")
	thisKey := $1.s
	thisValue := parsingStack.Pop().(ast.Expression)
	thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
	parsingStack.Push(thisExpression) 
}
;

array:
LBRACKET RBRACKET {
	logDebugGrammar("EMPTY ARRAY")
	thisExpression := ast.NewLiteralArray([]ast.Expression{})
	parsingStack.Push(thisExpression)
}
|
LBRACKET expression_list RBRACKET {
	logDebugGrammar("ARRAY")
	exp_list := parsingStack.Pop().([]ast.Expression)
	thisExpression := ast.NewLiteralArray(exp_list)
	parsingStack.Push(thisExpression)
}
;

expression_list:
expression {
	logDebugGrammar("EXPRESSION LIST SINGLE")
	exp_list := make([]ast.Expression, 0)
	exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
	parsingStack.Push(exp_list)
}
|
expression COMMA expression_list { 
	logDebugGrammar("EXPRESSION LIST COMPOUND")
	rest := parsingStack.Pop().([]ast.Expression)
	last := parsingStack.Pop()
	new_list := make([]ast.Expression, 0, len(rest) + 1)
	new_list = append(new_list, last.(ast.Expression))
	for _, v := range rest {
		new_list = append(new_list, v)
	}
	parsingStack.Push(new_list)
}
;