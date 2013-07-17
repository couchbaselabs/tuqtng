
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
const CASE = 57391
const WHEN = 57392
const THEN = 57393
const ELSE = 57394
const END = 57395
const MOD = 57396

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
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
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

const yyNprod = 98
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 438

var yyAct = []int{

	74, 125, 17, 86, 77, 117, 81, 19, 52, 53,
	54, 55, 57, 58, 154, 67, 60, 65, 63, 64,
	61, 62, 146, 75, 66, 52, 53, 54, 55, 57,
	73, 71, 67, 110, 56, 142, 123, 72, 67, 79,
	94, 94, 82, 14, 153, 93, 94, 88, 139, 89,
	19, 56, 136, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 68, 111,
	92, 141, 140, 122, 48, 120, 126, 132, 85, 52,
	53, 54, 55, 57, 58, 59, 67, 60, 65, 63,
	64, 61, 62, 91, 129, 66, 76, 143, 133, 124,
	131, 157, 128, 49, 68, 56, 130, 79, 127, 84,
	43, 137, 134, 135, 44, 52, 53, 54, 55, 57,
	58, 59, 67, 60, 65, 63, 64, 61, 62, 10,
	46, 66, 149, 148, 82, 124, 88, 151, 150, 12,
	90, 56, 50, 40, 122, 78, 120, 155, 156, 152,
	68, 54, 55, 57, 7, 8, 67, 32, 126, 158,
	20, 52, 53, 54, 55, 57, 58, 59, 67, 60,
	65, 63, 64, 61, 62, 56, 31, 66, 68, 138,
	30, 124, 69, 70, 147, 119, 118, 56, 145, 52,
	53, 54, 55, 57, 58, 59, 67, 60, 65, 63,
	64, 61, 62, 26, 24, 66, 68, 23, 83, 124,
	42, 87, 47, 16, 15, 56, 13, 52, 53, 54,
	55, 57, 58, 59, 67, 60, 65, 63, 64, 61,
	62, 6, 45, 66, 68, 11, 5, 144, 4, 41,
	9, 3, 2, 56, 1, 52, 53, 54, 55, 57,
	58, 59, 67, 60, 65, 63, 64, 61, 62, 0,
	112, 66, 0, 0, 0, 51, 52, 53, 54, 55,
	57, 56, 113, 67, 60, 65, 63, 64, 61, 62,
	0, 0, 66, 115, 114, 0, 0, 38, 0, 39,
	0, 0, 56, 33, 34, 35, 36, 37, 25, 29,
	0, 22, 121, 0, 0, 0, 0, 21, 0, 0,
	0, 0, 0, 0, 27, 116, 38, 0, 39, 0,
	0, 28, 33, 34, 35, 36, 37, 25, 29, 0,
	22, 121, 0, 0, 0, 0, 21, 0, 0, 0,
	0, 0, 0, 27, 0, 38, 0, 39, 0, 0,
	28, 33, 34, 35, 36, 37, 25, 29, 0, 22,
	18, 0, 0, 0, 0, 21, 0, 0, 0, 0,
	0, 0, 27, 0, 38, 0, 39, 80, 0, 28,
	33, 34, 35, 36, 37, 25, 29, 0, 22, 0,
	0, 0, 0, 0, 21, 0, 0, 0, 0, 0,
	0, 27, 0, 38, 0, 39, 0, 0, 28, 33,
	34, 35, 36, 37, 25, 29, 0, 22, 0, 0,
	0, 0, 0, 21, 0, 0, 0, 0, 0, 0,
	27, 0, 0, 0, 0, 0, 0, 28,
}
var yyPact = []int{

	150, -1000, -1000, 120, -1000, 132, 330, 138, -1000, 97,
	104, 122, 48, -1000, -1000, 84, -1000, 136, -1000, 217,
	-1000, 388, 388, -1000, -14, -5, -1000, 388, -27, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 80, 359,
	-1000, -1000, 95, 54, 388, -1000, 388, -1000, 134, 330,
	44, 15, 388, 388, 388, 388, 388, 388, 388, 388,
	388, 388, 388, 388, 388, 388, 388, -11, 388, -1000,
	-1000, 237, 272, -7, 87, 388, -1000, 92, 83, 74,
	-1000, 88, 81, -1000, 53, -1000, -1000, 79, 101, -1000,
	26, -1000, -1000, -1000, -1000, 121, 121, 3, 3, 3,
	3, 238, -20, -3, -3, -3, -3, -3, -3, -3,
	388, 161, -1000, 25, -1000, -1000, -1000, -8, 78, -1000,
	-1000, -1000, 189, -1000, 20, -30, 133, -1000, 12, 388,
	-1000, 388, -1000, 388, -1000, -1000, -1000, -3, -1000, -1000,
	-1000, -1000, -1000, 301, 14, -39, 388, 388, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 87, 51, 388, -1000,
}
var yyPgo = []int{

	0, 244, 242, 241, 240, 239, 238, 236, 235, 232,
	231, 216, 43, 214, 213, 2, 0, 212, 3, 211,
	210, 208, 160, 207, 204, 203, 1, 188, 5, 186,
	185, 180, 176, 157, 4, 145, 6,
}
var yyR1 = []int{

	0, 1, 2, 3, 6, 7, 10, 10, 11, 12,
	12, 13, 13, 13, 14, 14, 8, 8, 17, 17,
	9, 9, 4, 4, 18, 18, 19, 19, 19, 5,
	5, 5, 20, 21, 15, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 22, 22, 22, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 24, 24, 24,
	26, 26, 27, 27, 28, 28, 29, 29, 30, 30,
	25, 25, 25, 25, 25, 25, 25, 31, 31, 32,
	32, 34, 34, 35, 33, 33, 36, 36,
}
var yyR2 = []int{

	0, 1, 3, 1, 3, 2, 2, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 0, 2, 1, 3,
	0, 2, 0, 3, 1, 3, 1, 2, 2, 0,
	1, 2, 2, 2, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 1, 2, 2, 1, 1, 3, 4,
	3, 4, 3, 4, 1, 1, 3, 5, 3, 4,
	3, 5, 0, 2, 1, 3, 1, 1, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	3, 1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -6, -7, -10, 4, 5, -4,
	9, -8, 7, -11, -12, -13, -14, -15, 30, -16,
	-22, 35, 29, -23, -24, 26, -25, 42, 49, 27,
	-31, -32, -33, 21, 22, 23, 24, 25, 15, 17,
	5, -5, -20, 13, 10, -9, 8, -17, 26, 19,
	6, 48, 28, 29, 30, 31, 54, 32, 33, 34,
	36, 40, 41, 38, 39, 37, 44, 35, 17, -22,
	-22, 45, 42, -15, -16, 50, 16, -34, -35, 27,
	18, -36, -15, -21, 14, 24, -18, -19, -15, -15,
	6, -12, 26, 30, 26, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	44, -16, 23, 35, 47, 46, 43, -28, -29, -30,
	-15, 30, -16, 43, 48, -26, -16, 16, 19, 20,
	18, 19, 24, 19, 11, 12, 26, -16, 18, 23,
	47, 46, 43, 19, 48, -27, 52, 51, -34, -15,
	-36, -18, -28, 30, 53, -16, -16, 50, -26,
}
var yyDef = []int{

	0, -2, 1, 22, 3, 16, 0, 0, 7, 29,
	0, 20, 0, 5, 8, 9, 11, 12, 14, 34,
	53, 0, 0, 56, 57, 64, 65, 0, 0, 80,
	81, 82, 83, 84, 85, 86, 87, 88, 0, 0,
	6, 2, 30, 0, 0, 4, 0, 17, 18, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 54,
	55, 0, 0, 0, 34, 0, 89, 0, 91, 0,
	94, 0, 96, 31, 0, 32, 23, 24, 26, 21,
	0, 10, 13, 15, 51, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	0, 0, 58, 0, 60, 62, 68, 0, 74, 76,
	77, 78, 34, 66, 0, 72, 0, 90, 0, 0,
	95, 0, 33, 0, 27, 28, 19, 50, 52, 59,
	61, 63, 69, 0, 0, 0, 0, 0, 92, 93,
	97, 25, 75, 79, 67, 73, 70, 0, 71,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54,
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
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = &ast.From{Bucket: yyS[yypt-0].s}
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 18:
		//line unql.y:171
		{
		logDebugGrammar("FROM DATASOURCE")
	}
	case 19:
		//line unql.y:175
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
	}
	case 20:
		//line unql.y:181
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 21:
		//line unql.y:185
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
		//line unql.y:199
		{
		
	}
	case 24:
		//line unql.y:205
		{
		
	}
	case 25:
		//line unql.y:209
		{
		
	}
	case 26:
		//line unql.y:214
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
		//line unql.y:225
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
		//line unql.y:236
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
		//line unql.y:248
		{
		
	}
	case 30:
		//line unql.y:252
		{
		
	}
	case 31:
		//line unql.y:256
		{
		
	}
	case 32:
		//line unql.y:262
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
		//line unql.y:273
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
		//line unql.y:286
		{
		logDebugGrammar("EXPRESSION")
	}
	case 35:
		//line unql.y:291
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 36:
		//line unql.y:299
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 37:
		//line unql.y:307
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 38:
		//line unql.y:315
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 39:
		//line unql.y:323
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 40:
		//line unql.y:331
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 41:
		//line unql.y:339
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 42:
		//line unql.y:347
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 43:
		//line unql.y:355
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 44:
		//line unql.y:363
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 45:
		//line unql.y:371
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 46:
		//line unql.y:379
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 47:
		//line unql.y:387
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:395
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:403
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:411
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:419
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:427
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:435
		{
		
	}
	case 54:
		//line unql.y:441
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:448
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:455
		{
		
	}
	case 57:
		//line unql.y:460
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 58:
		//line unql.y:464
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:471
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:478
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:485
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:492
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:499
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:508
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)	
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 65:
		//line unql.y:514
		{
		logDebugGrammar("LITERAL")
	}
	case 66:
		//line unql.y:518
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 67:
		//line unql.y:522
		{
		logDebugGrammar("CASE WHEN THEN ELSE END")	
	}
	case 68:
		//line unql.y:526
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:532
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:541
		{
		logDebugGrammar("THEN_LIST - SINGLE")	
	}
	case 71:
		//line unql.y:545
		{
		logDebugGrammar("THEN_LIST - COMPOUND")
	}
	case 72:
		//line unql.y:551
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 73:
		//line unql.y:555
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 74:
		//line unql.y:561
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 75:
		//line unql.y:566
		{
		funarg_expr_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		// list items pushed onto the stack end up in reverse order
	// this prepends items in the list to restore order
	new_list := ast.FunctionArgExpressionList{funarg_expr}
		for _, v := range funarg_expr_list {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 76:
		//line unql.y:580
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 77:
		//line unql.y:584
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 78:
		//line unql.y:593
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 79:
		//line unql.y:599
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 80:
		//line unql.y:609
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line unql.y:615
		{
		logDebugGrammar("NUMBER")
	}
	case 82:
		//line unql.y:619
		{
		logDebugGrammar("OBJECT")
	}
	case 83:
		//line unql.y:623
		{
		logDebugGrammar("ARRAY")
	}
	case 84:
		//line unql.y:627
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line unql.y:633
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line unql.y:639
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 87:
		//line unql.y:647
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line unql.y:653
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line unql.y:661
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 90:
		//line unql.y:667
		{
		logDebugGrammar("OBJECT")
	}
	case 91:
		//line unql.y:673
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 92:
		//line unql.y:677
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 93:
		//line unql.y:689
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 94:
		//line unql.y:699
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray([]ast.Expression{})
		parsingStack.Push(thisExpression)
	}
	case 95:
		//line unql.y:705
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().([]ast.Expression)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line unql.y:714
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make([]ast.Expression, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 97:
		//line unql.y:721
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
