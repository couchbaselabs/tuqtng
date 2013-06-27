
//line unql.y:2
package goyacc
import __yyfmt__ "fmt"
//line unql.y:2
		import "log"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
    if DebugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if DebugGrammar {
        log.Printf("DEBUG GRAMMAR " + format)
    }
}

//line unql.y:15
type yySymType struct {
	yys int 
s string 
n int
f float64}

const EXPLAIN = 57346
const SELECT = 57347
const AS = 57348
const FROM = 57349
const WHERE = 57350
const ORDER = 57351
const BY = 57352
const ASC = 57353
const DESC = 57354
const LIMIT = 57355
const OFFSET = 57356
const LBRACE = 57357
const RBRACE = 57358
const LBRACKET = 57359
const RBRACKET = 57360
const COMMA = 57361
const COLON = 57362
const DQUOTE = 57363
const TRUE = 57364
const FALSE = 57365
const NULL = 57366
const INT = 57367
const NUMBER = 57368
const IDENTIFIER = 57369
const STRING = 57370
const PLUS = 57371
const MINUS = 57372
const MULT = 57373
const DIV = 57374
const CONCAT = 57375
const AND = 57376
const OR = 57377
const NOT = 57378
const EQ = 57379
const NE = 57380
const GT = 57381
const GTE = 57382
const LT = 57383
const LTE = 57384
const LPAREN = 57385
const RPAREN = 57386
const LIKE = 57387
const IS = 57388
const VALUED = 57389
const MISSING = 57390
const DOT = 57391
const MOD = 57392

var yyToknames = []string{
	"EXPLAIN",
	"SELECT",
	"AS",
	"FROM",
	"WHERE",
	"ORDER",
	"BY",
	"ASC",
	"DESC",
	"LIMIT",
	"OFFSET",
	"LBRACE",
	"RBRACE",
	"LBRACKET",
	"RBRACKET",
	"COMMA",
	"COLON",
	"DQUOTE",
	"TRUE",
	"FALSE",
	"NULL",
	"INT",
	"NUMBER",
	"IDENTIFIER",
	"STRING",
	"PLUS",
	"MINUS",
	"MULT",
	"DIV",
	"CONCAT",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"NE",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"LPAREN",
	"RPAREN",
	"LIKE",
	"IS",
	"VALUED",
	"MISSING",
	"DOT",
	"MOD",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 85
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 261

var yyAct = []int{

	17, 72, 83, 78, 74, 70, 107, 113, 19, 51,
	52, 53, 54, 56, 57, 66, 66, 59, 64, 62,
	63, 60, 61, 91, 109, 65, 14, 90, 71, 73,
	55, 51, 52, 53, 54, 56, 110, 76, 66, 79,
	117, 76, 91, 124, 85, 89, 86, 112, 111, 127,
	19, 47, 55, 92, 93, 94, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 105, 106, 120, 108,
	67, 126, 129, 128, 82, 88, 121, 20, 119, 116,
	48, 118, 51, 52, 53, 54, 56, 57, 58, 66,
	59, 64, 62, 63, 60, 61, 115, 81, 65, 68,
	69, 42, 114, 55, 122, 123, 45, 43, 10, 125,
	12, 87, 49, 67, 7, 8, 39, 75, 131, 31,
	79, 130, 85, 132, 133, 51, 52, 53, 54, 56,
	57, 58, 66, 59, 64, 62, 63, 60, 61, 67,
	30, 65, 29, 26, 24, 114, 55, 23, 80, 41,
	84, 51, 52, 53, 54, 56, 57, 58, 66, 59,
	64, 62, 63, 60, 61, 46, 16, 65, 15, 13,
	6, 50, 55, 51, 52, 53, 54, 56, 44, 11,
	66, 59, 64, 62, 63, 60, 61, 5, 37, 65,
	38, 4, 40, 9, 55, 32, 33, 34, 35, 36,
	25, 28, 3, 22, 18, 2, 1, 0, 0, 21,
	37, 0, 38, 77, 0, 0, 27, 32, 33, 34,
	35, 36, 25, 28, 0, 22, 0, 0, 0, 0,
	0, 21, 37, 0, 38, 0, 0, 0, 27, 32,
	33, 34, 35, 36, 25, 28, 0, 22, 0, 0,
	0, 0, 0, 21, 0, 0, 0, 0, 0, 0,
	27,
}
var yyPact = []int{

	110, -1000, -1000, 99, -1000, 103, 173, 111, -1000, 88,
	97, 98, 24, -1000, -1000, 61, -1000, 106, -1000, 122,
	-1000, 217, 217, -1000, -41, -1000, -1000, 217, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 13, 195, -1000,
	-1000, 83, 49, 217, -1000, 217, -1000, 105, 173, 18,
	-4, 217, 217, 217, 217, 217, 217, 217, 217, 217,
	217, 217, 217, 217, 217, 217, -39, 217, -1000, -1000,
	0, -37, 96, -1000, 80, 60, 20, -1000, 63, 59,
	-1000, 43, -1000, -1000, 57, 93, -1000, 16, -1000, -1000,
	-1000, -1000, -21, -21, -21, -21, -21, -21, 144, -20,
	2, 2, 2, 2, 2, 2, 2, 217, 53, -1000,
	25, -1000, -1000, -1000, 15, -1000, 9, 217, -1000, 217,
	-1000, 217, -1000, -1000, -1000, 2, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 206, 205, 202, 193, 192, 191, 187, 179, 178,
	170, 169, 26, 168, 166, 0, 1, 165, 2, 150,
	149, 148, 77, 147, 144, 143, 142, 140, 119, 4,
	117, 3,
}
var yyR1 = []int{

	0, 1, 2, 3, 6, 7, 10, 10, 11, 12,
	12, 13, 13, 13, 14, 14, 8, 8, 17, 17,
	9, 9, 4, 4, 18, 18, 19, 19, 19, 5,
	5, 5, 20, 21, 15, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 22, 22, 22, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 25, 25, 25,
	25, 25, 25, 25, 26, 26, 27, 27, 29, 29,
	30, 28, 28, 31, 31,
}
var yyR2 = []int{

	0, 1, 3, 1, 3, 2, 2, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 0, 2, 1, 3,
	0, 2, 0, 3, 1, 3, 1, 2, 2, 0,
	1, 2, 2, 2, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 1, 2, 2, 1, 1, 3, 4,
	3, 4, 3, 4, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 3, 1, 3,
	3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -6, -7, -10, 4, 5, -4,
	9, -8, 7, -11, -12, -13, -14, -15, 31, -16,
	-22, 36, 30, -23, -24, 27, -25, 43, 28, -26,
	-27, -28, 22, 23, 24, 25, 26, 15, 17, 5,
	-5, -20, 13, 10, -9, 8, -17, 27, 19, 6,
	49, 29, 30, 31, 32, 50, 33, 34, 35, 37,
	41, 42, 39, 40, 38, 45, 36, 17, -22, -22,
	46, -15, -16, 16, -29, -30, 28, 18, -31, -15,
	-21, 14, 25, -18, -19, -15, -15, 6, -12, 27,
	31, 27, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 45, -16, 24,
	36, 48, 47, 44, 49, 16, 19, 20, 18, 19,
	25, 19, 11, 12, 27, -16, 18, 24, 48, 47,
	-29, -15, -31, -18,
}
var yyDef = []int{

	0, -2, 1, 22, 3, 16, 0, 0, 7, 29,
	0, 20, 0, 5, 8, 9, 11, 12, 14, 34,
	53, 0, 0, 56, 57, 64, 65, 0, 67, 68,
	69, 70, 71, 72, 73, 74, 75, 0, 0, 6,
	2, 30, 0, 0, 4, 0, 17, 18, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 54, 55,
	0, 0, 34, 76, 0, 78, 0, 81, 0, 83,
	31, 0, 32, 23, 24, 26, 21, 0, 10, 13,
	15, 51, 35, 36, 37, 38, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 0, 0, 58,
	0, 60, 62, 66, 0, 77, 0, 0, 82, 0,
	33, 0, 27, 28, 19, 50, 52, 59, 61, 63,
	79, 80, 84, 25,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line unql.y:45
		{ 
		logDebugGrammar("INPUT") 
	}
	case 2:
		//line unql.y:52
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 3:
		//line unql.y:58
		{ 
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 4:
		//line unql.y:65
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 5:
		//line unql.y:71
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 6:
		//line unql.y:77
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
		parsingStatement.SetExplainOnly(true)
	}
	case 7:
		//line unql.y:83
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
	}
	case 8:
		//line unql.y:90
		{ 
		logDebugGrammar("SELECT SELECT TAIL - EXPR")
		result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Select = result_expr_list
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	
	}
	case 9:
		//line unql.y:104
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 10:
		//line unql.y:109
		{
		result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		// list items pushed onto the stack end up in reverse order
	// this prepends items in the list to restore order
	new_list := ast.ResultExpressionList{result_expr}
		for _, v := range result_expr_list {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 11:
		//line unql.y:122
		{
		logDebugGrammar("RESULT STAR")
	}
	case 12:
		//line unql.y:126
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 13:
		//line unql.y:133
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 14:
		//line unql.y:142
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 15:
		//line unql.y:148
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 16:
		//line unql.y:156
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 17:
		//line unql.y:160
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
	}
	case 18:
		//line unql.y:165
		{
		logDebugGrammar("FROM DATASOURCE")
	}
	case 19:
		//line unql.y:169
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
	}
	case 20:
		//line unql.y:175
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 21:
		//line unql.y:179
		{
		logDebugGrammar("SELECT WHERE - EXPR")
		where_part := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Where = where_part
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 23:
		//line unql.y:193
		{
		
	}
	case 24:
		//line unql.y:199
		{
		
	}
	case 25:
		//line unql.y:203
		{
		
	}
	case 26:
		//line unql.y:208
		{ 
		logDebugGrammar("SORT EXPR")
		expr := parsingStack.Pop()
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), true))
		default:
			logDebugGrammar("This statement does not support ORDER BY")
		}
	}
	case 27:
		//line unql.y:219
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr := parsingStack.Pop()
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), true))
		default:
			logDebugGrammar("This statement does not support ORDER BY")
		}
	}
	case 28:
		//line unql.y:230
		{ 
		logDebugGrammar("SORT EXPR DESC")
		expr := parsingStack.Pop()
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), false))
		default:
			logDebugGrammar("This statement does not support ORDER BY")
		}
	}
	case 29:
		//line unql.y:242
		{
		
	}
	case 30:
		//line unql.y:246
		{
		
	}
	case 31:
		//line unql.y:250
		{
		
	}
	case 32:
		//line unql.y:256
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 33:
		//line unql.y:267
		{ 
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 34:
		//line unql.y:280
		{
		logDebugGrammar("EXPRESSION")
	}
	case 35:
		//line unql.y:285
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 36:
		//line unql.y:293
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 37:
		//line unql.y:301
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 38:
		//line unql.y:309
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 39:
		//line unql.y:317
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 40:
		//line unql.y:325
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 41:
		//line unql.y:333
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 42:
		//line unql.y:341
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 43:
		//line unql.y:349
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 44:
		//line unql.y:357
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 45:
		//line unql.y:365
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 46:
		//line unql.y:373
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 47:
		//line unql.y:381
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:389
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:397
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:405
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:413
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:421
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:429
		{
		
	}
	case 54:
		//line unql.y:435
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:442
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:449
		{
		
	}
	case 57:
		//line unql.y:454
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 58:
		//line unql.y:458
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:465
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:472
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:479
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:486
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:493
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:502
		{
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 65:
		//line unql.y:507
		{
		logDebugGrammar("LITERAL")
	}
	case 66:
		//line unql.y:511
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 67:
		//line unql.y:518
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:524
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:530
		{
		logDebugGrammar("OBJECT")
	}
	case 70:
		//line unql.y:534
		{
		logDebugGrammar("ARRAY")
	}
	case 71:
		//line unql.y:538
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:544
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:550
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:558
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
	}
	case 75:
		//line unql.y:562
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
	}
	case 76:
		//line unql.y:568
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 77:
		//line unql.y:574
		{
		logDebugGrammar("OBJECT")
	}
	case 78:
		//line unql.y:580
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 79:
		//line unql.y:584
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 80:
		//line unql.y:596
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 81:
		//line unql.y:606
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray([]ast.Expression{})
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line unql.y:612
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().([]ast.Expression)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line unql.y:621
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make([]ast.Expression, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 84:
		//line unql.y:628
		{ 
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
	}
	goto yystack /* stack new state and value */
}
