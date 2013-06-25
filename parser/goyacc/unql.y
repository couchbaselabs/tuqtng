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

%token LBRACE RBRACE LBRACKET RBRACKET
%token COMMA COLON DQUOTE
%token TRUE FALSE NULL
%token INT NUMBER IDENTIFIER STRING
%token PLUS MINUS MULT DIV
%token CONCAT
%token AND OR NOT
%token EQ NE GT GTE LT LTE
%token LPAREN RPAREN
%left OR
%left AND 
%left EQ LT LTE GT GTE NE
%left PLUS MINUS MULT DIV CONCAT
%right NOT

%%

input: 
expression { 
	logDebugGrammar("INPUT") 
}
;

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
prefix_expr {
	
}
;

prefix_expr: 
NOT prefix_expr {
	logDebugGrammar("EXPR - NOT")
}
|
suffix_expr {
	
};

suffix_expr: 
atom {
	logDebugGrammar("SUFFIX_EXPR")
};

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