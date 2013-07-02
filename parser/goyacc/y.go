
//line unql.y:2
package goyacc
import __yyfmt__ "fmt"
//line unql.y:2
		import "log"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
	if DebugGrammar {
    	log.Printf("DEBUG GRAMMAR " + format, v...)
    }
}

//line unql.y:13
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
const TRUE = 57363
const FALSE = 57364
const NULL = 57365
const INT = 57366
const NUMBER = 57367
const IDENTIFIER = 57368
const STRING = 57369
const PLUS = 57370
const MINUS = 57371
const MULT = 57372
const DIV = 57373
const CONCAT = 57374
const AND = 57375
const OR = 57376
const NOT = 57377
const EQ = 57378
const NE = 57379
const GT = 57380
const GTE = 57381
const LT = 57382
const LTE = 57383
const LPAREN = 57384
const RPAREN = 57385
const LIKE = 57386
const IS = 57387
const VALUED = 57388
const MISSING = 57389
const DOT = 57390
const MOD = 57391

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
	76, 91, 124, 89, 85, 47, 86, 112, 111, 127,
	19, 120, 55, 92, 93, 94, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 105, 106, 82, 108,
	67, 126, 129, 128, 117, 88, 121, 119, 116, 48,
	118, 51, 52, 53, 54, 56, 57, 58, 66, 59,
	64, 62, 63, 60, 61, 115, 81, 65, 53, 54,
	56, 114, 55, 66, 122, 123, 43, 42, 10, 125,
	45, 12, 87, 49, 67, 7, 8, 55, 131, 39,
	79, 130, 85, 132, 133, 51, 52, 53, 54, 56,
	57, 58, 66, 59, 64, 62, 63, 60, 61, 67,
	75, 65, 20, 31, 30, 114, 55, 29, 26, 24,
	51, 52, 53, 54, 56, 57, 58, 66, 59, 64,
	62, 63, 60, 61, 68, 69, 65, 23, 80, 41,
	50, 55, 51, 52, 53, 54, 56, 84, 46, 66,
	59, 64, 62, 63, 60, 61, 16, 15, 65, 37,
	13, 38, 6, 55, 44, 32, 33, 34, 35, 36,
	25, 28, 11, 22, 18, 5, 4, 40, 9, 21,
	3, 37, 2, 38, 77, 1, 27, 32, 33, 34,
	35, 36, 25, 28, 0, 22, 0, 0, 0, 0,
	0, 21, 0, 37, 0, 38, 0, 0, 27, 32,
	33, 34, 35, 36, 25, 28, 0, 22, 0, 0,
	0, 0, 0, 21, 0, 0, 0, 0, 0, 0,
	27,
}
var yyPact = []int{

	111, -1000, -1000, 99, -1000, 104, 174, 114, -1000, 94,
	96, 102, 19, -1000, -1000, 60, -1000, 107, -1000, 122,
	-1000, 218, 218, -1000, -40, -1000, -1000, 218, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 13, 196, -1000,
	-1000, 82, 44, 218, -1000, 218, -1000, 106, 174, 17,
	-3, 218, 218, 218, 218, 218, 218, 218, 218, 218,
	218, 218, 218, 218, 218, 218, -38, 218, -1000, -1000,
	1, -36, 97, -1000, 79, 59, 54, -1000, 62, 58,
	-1000, 27, -1000, -1000, 57, 93, -1000, 16, -1000, -1000,
	-1000, -1000, 68, 68, -20, -20, -20, -20, 144, -19,
	3, 3, 3, 3, 3, 3, 3, 218, 53, -1000,
	26, -1000, -1000, -1000, 15, -1000, 10, 218, -1000, 218,
	-1000, 218, -1000, -1000, -1000, 3, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 215, 212, 210, 208, 207, 206, 205, 202, 194,
	192, 190, 26, 187, 186, 0, 1, 178, 2, 177,
	169, 168, 142, 167, 149, 148, 147, 144, 143, 4,
	140, 3,
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
	9, -8, 7, -11, -12, -13, -14, -15, 30, -16,
	-22, 35, 29, -23, -24, 26, -25, 42, 27, -26,
	-27, -28, 21, 22, 23, 24, 25, 15, 17, 5,
	-5, -20, 13, 10, -9, 8, -17, 26, 19, 6,
	48, 28, 29, 30, 31, 49, 32, 33, 34, 36,
	40, 41, 38, 39, 37, 44, 35, 17, -22, -22,
	45, -15, -16, 16, -29, -30, 27, 18, -31, -15,
	-21, 14, 24, -18, -19, -15, -15, 6, -12, 26,
	30, 26, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, 44, -16, 23,
	35, 47, 46, 43, 48, 16, 19, 20, 18, 19,
	24, 19, 11, 12, 26, -16, 18, 23, 47, 46,
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
	42, 43, 44, 45, 46, 47, 48, 49,
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
		//line unql.y:44
		{ 
		logDebugGrammar("INPUT") 
	}
	case 2:
		//line unql.y:51
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 3:
		//line unql.y:57
		{ 
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 4:
		//line unql.y:64
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 5:
		//line unql.y:70
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 6:
		//line unql.y:76
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
		parsingStatement.SetExplainOnly(true)
	}
	case 7:
		//line unql.y:82
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
	}
	case 8:
		//line unql.y:89
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
		//line unql.y:103
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 10:
		//line unql.y:108
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
		//line unql.y:121
		{
		logDebugGrammar("RESULT STAR")
	}
	case 12:
		//line unql.y:125
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 13:
		//line unql.y:132
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 14:
		//line unql.y:141
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 15:
		//line unql.y:147
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 16:
		//line unql.y:155
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 17:
		//line unql.y:159
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = &ast.From{Bucket: yyS[yypt-0].s}
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 18:
		//line unql.y:170
		{
		logDebugGrammar("FROM DATASOURCE")
	}
	case 19:
		//line unql.y:174
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
	}
	case 20:
		//line unql.y:180
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 21:
		//line unql.y:184
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
		//line unql.y:198
		{
		
	}
	case 24:
		//line unql.y:204
		{
		
	}
	case 25:
		//line unql.y:208
		{
		
	}
	case 26:
		//line unql.y:213
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
		//line unql.y:224
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
		//line unql.y:235
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
		//line unql.y:247
		{
		
	}
	case 30:
		//line unql.y:251
		{
		
	}
	case 31:
		//line unql.y:255
		{
		
	}
	case 32:
		//line unql.y:261
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
		//line unql.y:272
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
		//line unql.y:285
		{
		logDebugGrammar("EXPRESSION")
	}
	case 35:
		//line unql.y:290
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 36:
		//line unql.y:298
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 37:
		//line unql.y:306
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 38:
		//line unql.y:314
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 39:
		//line unql.y:322
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 40:
		//line unql.y:330
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 41:
		//line unql.y:338
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 42:
		//line unql.y:346
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 43:
		//line unql.y:354
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 44:
		//line unql.y:362
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 45:
		//line unql.y:370
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 46:
		//line unql.y:378
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 47:
		//line unql.y:386
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:394
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:402
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:410
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:418
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:426
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:434
		{
		
	}
	case 54:
		//line unql.y:440
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:447
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:454
		{
		
	}
	case 57:
		//line unql.y:459
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 58:
		//line unql.y:463
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:470
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:477
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:484
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:491
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:498
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:507
		{
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 65:
		//line unql.y:512
		{
		logDebugGrammar("LITERAL")
	}
	case 66:
		//line unql.y:516
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 67:
		//line unql.y:523
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:529
		{
		logDebugGrammar("NUMBER")
	}
	case 69:
		//line unql.y:533
		{
		logDebugGrammar("OBJECT")
	}
	case 70:
		//line unql.y:537
		{
		logDebugGrammar("ARRAY")
	}
	case 71:
		//line unql.y:541
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:547
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:553
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:561
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line unql.y:567
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:575
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 77:
		//line unql.y:581
		{
		logDebugGrammar("OBJECT")
	}
	case 78:
		//line unql.y:587
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 79:
		//line unql.y:591
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
		//line unql.y:603
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 81:
		//line unql.y:613
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray([]ast.Expression{})
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line unql.y:619
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().([]ast.Expression)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line unql.y:628
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make([]ast.Expression, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 84:
		//line unql.y:635
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
