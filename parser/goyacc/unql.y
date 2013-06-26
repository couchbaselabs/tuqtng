%{
package goyacc
import "log"

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
}
|
expr MINUS expr {
	logDebugGrammar("EXPR - MINUS")
}
|
expr MULT expr {
	logDebugGrammar("EXPR - MULT")
}
|
expr DIV expr {
	logDebugGrammar("EXPR - DIV")
}
|
expr MOD expr {
	logDebugGrammar("EXPR - MOD")
}
|
expr CONCAT expr {
	logDebugGrammar("EXPR - DIV")
}
|
expr AND expr {
	logDebugGrammar("EXPR - AND")
}
|
expr OR expr {
	logDebugGrammar("EXPR - OR")
}
|
expr EQ expr {
	logDebugGrammar("EXPR - EQ")
}
|
expr LT expr {
	logDebugGrammar("EXPR - LT")
}
|
expr LTE expr {
	logDebugGrammar("EXPR - LTE")
}
|
expr GT expr {
	logDebugGrammar("EXPR - GT")
}
|
expr GTE expr {
	logDebugGrammar("EXPR - GTE")
}
|
expr NE expr {
	logDebugGrammar("EXPR - NE")
}
|
expr LIKE expr {
	logDebugGrammar("EXPR - LIKE")
}
|
expr NOT LIKE expr {
	logDebugGrammar("EXPR - NOT LIKE")
}
|
expr DOT IDENTIFIER {
	logDebugGrammar("EXPR DOT MEMBER")
}
|
expr LBRACKET expr RBRACKET {
	logDebugGrammar("EXPR BRACKET MEMBER")
}
|
prefix_expr {
	
}
;

prefix_expr: 
NOT prefix_expr {
	logDebugGrammar("EXPR - NOT")
}
|
MINUS prefix_expr {
	logDebugGrammar("EXPR - CHANGE SIGN")
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
}
|
atom IS NOT NULL {
	logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
}
|
atom IS MISSING {
	logDebugGrammar("SUFFIX_EXPR IS MISSING")
}
|
atom IS NOT MISSING {
	logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
}
|
atom IS VALUED {
	logDebugGrammar("SUFFIX_EXPR IS VALUED")
}
|
atom IS NOT VALUED {
	logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
}
;

atom:
IDENTIFIER
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
}
|
number {
	logDebugGrammar("NUMBER")
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
}
|
FALSE {
	logDebugGrammar("FALSE")
}
|
NULL {
	logDebugGrammar("NULL")
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
}
;

named_expression_single:   
STRING COLON expression {  
	logDebugGrammar("NAMED EXPR SINGLE")
}
;

array:
LBRACKET RBRACKET {
	logDebugGrammar("EMPTY ARRAY")
}
|
LBRACKET expression_list RBRACKET {
	logDebugGrammar("ARRAY")
}
;

expression_list:
expression {
	logDebugGrammar("EXPRESSION LIST SINGLE")
}
|
expression COMMA expression_list { 
	logDebugGrammar("EXPRESSION LIST COMPOUND")
}
;