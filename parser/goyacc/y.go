
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
const ANY = 57396
const ALL = 57397
const OVER = 57398
const MOD = 57399

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
	"ANY",
	"ALL",
	"OVER",
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

const yyNprod = 103
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 579

var yyAct = []int{

	76, 129, 121, 90, 154, 81, 85, 19, 54, 55,
	56, 57, 59, 60, 163, 69, 62, 67, 65, 66,
	63, 64, 152, 77, 68, 54, 55, 56, 57, 59,
	78, 79, 69, 14, 73, 114, 148, 58, 127, 69,
	74, 83, 17, 167, 98, 56, 57, 59, 162, 174,
	69, 155, 19, 145, 58, 99, 100, 101, 102, 103,
	104, 105, 106, 107, 108, 109, 110, 111, 112, 113,
	75, 115, 58, 171, 168, 126, 147, 146, 130, 98,
	98, 70, 142, 97, 86, 95, 96, 50, 135, 92,
	172, 93, 54, 55, 56, 57, 59, 60, 61, 69,
	62, 67, 65, 66, 63, 64, 138, 80, 68, 89,
	176, 70, 128, 136, 170, 143, 149, 124, 83, 139,
	137, 58, 54, 55, 56, 57, 59, 60, 61, 69,
	62, 67, 65, 66, 63, 64, 116, 156, 68, 20,
	157, 134, 128, 160, 159, 153, 51, 133, 117, 88,
	126, 58, 161, 164, 165, 140, 141, 46, 45, 119,
	118, 71, 72, 10, 48, 12, 169, 166, 94, 52,
	42, 130, 175, 173, 7, 8, 82, 34, 158, 33,
	86, 32, 92, 70, 144, 123, 122, 151, 26, 24,
	23, 87, 124, 44, 54, 55, 56, 57, 59, 60,
	61, 69, 62, 67, 65, 66, 63, 64, 91, 49,
	68, 16, 15, 70, 128, 13, 6, 47, 11, 5,
	4, 43, 9, 58, 54, 55, 56, 57, 59, 60,
	61, 69, 62, 67, 65, 66, 63, 64, 3, 2,
	68, 1, 0, 70, 128, 0, 0, 0, 0, 0,
	0, 0, 132, 58, 54, 55, 56, 57, 59, 60,
	61, 69, 62, 67, 65, 66, 63, 64, 0, 0,
	68, 0, 0, 70, 128, 0, 0, 0, 0, 0,
	0, 0, 131, 58, 54, 55, 56, 57, 59, 60,
	61, 69, 62, 67, 65, 66, 63, 64, 0, 0,
	68, 0, 0, 70, 128, 0, 0, 0, 0, 0,
	0, 0, 0, 58, 54, 55, 56, 57, 59, 60,
	61, 69, 62, 67, 65, 66, 63, 64, 0, 0,
	68, 0, 0, 70, 150, 0, 0, 0, 0, 0,
	0, 0, 0, 58, 54, 55, 56, 57, 59, 60,
	61, 69, 62, 67, 65, 66, 63, 64, 0, 0,
	68, 0, 0, 0, 53, 0, 0, 0, 40, 0,
	41, 0, 0, 58, 35, 36, 37, 38, 39, 25,
	31, 0, 22, 125, 0, 0, 0, 0, 21, 0,
	0, 0, 0, 0, 0, 27, 120, 0, 0, 0,
	0, 0, 28, 40, 0, 41, 0, 29, 30, 35,
	36, 37, 38, 39, 25, 31, 0, 22, 125, 0,
	0, 0, 0, 21, 0, 0, 0, 0, 0, 0,
	27, 0, 0, 0, 0, 0, 0, 28, 0, 0,
	0, 0, 29, 30, 54, 55, 56, 57, 59, 0,
	0, 69, 62, 67, 65, 66, 63, 64, 0, 0,
	68, 0, 0, 0, 0, 0, 0, 0, 40, 0,
	41, 0, 0, 58, 35, 36, 37, 38, 39, 25,
	31, 0, 22, 18, 0, 0, 0, 0, 21, 0,
	0, 0, 0, 0, 0, 27, 0, 0, 0, 0,
	0, 0, 28, 40, 0, 41, 84, 29, 30, 35,
	36, 37, 38, 39, 25, 31, 0, 22, 0, 0,
	0, 0, 0, 21, 0, 0, 0, 0, 0, 0,
	27, 0, 0, 0, 0, 0, 0, 28, 40, 0,
	41, 0, 29, 30, 35, 36, 37, 38, 39, 25,
	31, 0, 22, 0, 0, 0, 0, 0, 21, 0,
	0, 0, 0, 0, 0, 27, 0, 0, 0, 0,
	0, 0, 28, 0, 0, 0, 0, 29, 30,
}
var yyPact = []int{

	170, -1000, -1000, 154, -1000, 158, 453, 165, -1000, 145,
	147, 156, 61, -1000, -1000, 127, -1000, 163, -1000, 316,
	-1000, 523, 523, -1000, -11, -2, -1000, 523, -27, 523,
	523, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	91, 488, -1000, -1000, 135, 85, 523, -1000, 523, -1000,
	162, 453, 60, 53, 523, 523, 523, 523, 523, 523,
	523, 523, 523, 523, 523, 523, 523, 523, 523, -9,
	523, -1000, -1000, 113, 353, -5, 256, 523, 226, 196,
	-1000, 131, 122, 68, -1000, 95, 101, -1000, 82, -1000,
	-1000, 100, 144, -1000, 56, -1000, -1000, -1000, -1000, 15,
	15, 4, 4, 4, 4, 416, -20, -3, -3, -3,
	-3, -3, -3, -3, 523, 166, -1000, 30, -1000, -1000,
	-1000, -7, 97, -1000, -1000, -1000, 286, -1000, 54, -30,
	94, 25, 25, -1000, 14, 523, -1000, 523, -1000, 523,
	-1000, -1000, -1000, -3, -1000, -1000, -1000, -1000, -1000, 388,
	18, -39, 523, 523, 161, 26, 160, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 256, 64, 47, 66, 25, 23,
	523, -1000, 92, -1000, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 241, 239, 238, 222, 221, 220, 219, 218, 217,
	216, 215, 33, 212, 211, 42, 0, 209, 3, 208,
	193, 191, 139, 190, 189, 188, 1, 187, 4, 2,
	186, 185, 181, 179, 177, 5, 176, 6,
}
var yyR1 = []int{

	0, 1, 2, 3, 6, 7, 10, 10, 11, 12,
	12, 13, 13, 13, 14, 14, 8, 8, 17, 17,
	9, 9, 4, 4, 18, 18, 19, 19, 19, 5,
	5, 5, 20, 21, 15, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 22, 22, 22, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 24, 24, 24,
	24, 24, 26, 26, 27, 27, 28, 28, 28, 29,
	29, 30, 30, 31, 31, 25, 25, 25, 25, 25,
	25, 25, 32, 32, 33, 33, 35, 35, 36, 34,
	34, 37, 37,
}
var yyR2 = []int{

	0, 1, 3, 1, 3, 2, 2, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 0, 2, 1, 3,
	0, 2, 0, 3, 1, 3, 1, 2, 2, 0,
	1, 2, 2, 2, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 1, 2, 2, 1, 1, 3, 4,
	3, 4, 3, 4, 1, 1, 3, 5, 6, 6,
	3, 4, 3, 5, 0, 2, 1, 4, 3, 1,
	3, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 1, 3, 3, 2,
	3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -6, -7, -10, 4, 5, -4,
	9, -8, 7, -11, -12, -13, -14, -15, 30, -16,
	-22, 35, 29, -23, -24, 26, -25, 42, 49, 54,
	55, 27, -32, -33, -34, 21, 22, 23, 24, 25,
	15, 17, 5, -5, -20, 13, 10, -9, 8, -17,
	26, 19, 6, 48, 28, 29, 30, 31, 57, 32,
	33, 34, 36, 40, 41, 38, 39, 37, 44, 35,
	17, -22, -22, 45, 42, -15, -16, 50, -16, -16,
	16, -35, -36, 27, 18, -37, -15, -21, 14, 24,
	-18, -19, -15, -15, 6, -12, 26, 30, 26, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, 44, -16, 23, 35, 47, 46,
	43, -29, -30, -31, -15, 30, -16, 43, 48, -26,
	-16, 56, 56, 16, 19, 20, 18, 19, 24, 19,
	11, 12, 26, -16, 18, 23, 47, 46, 43, 19,
	48, -27, 52, 51, -28, 26, -28, -35, -15, -37,
	-18, -29, 30, 53, -16, -16, 6, 17, 48, 6,
	50, 26, 24, -28, 26, -26, 18,
}
var yyDef = []int{

	0, -2, 1, 22, 3, 16, 0, 0, 7, 29,
	0, 20, 0, 5, 8, 9, 11, 12, 14, 34,
	53, 0, 0, 56, 57, 64, 65, 0, 0, 0,
	0, 85, 86, 87, 88, 89, 90, 91, 92, 93,
	0, 0, 6, 2, 30, 0, 0, 4, 0, 17,
	18, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 54, 55, 0, 0, 0, 34, 0, 0, 0,
	94, 0, 96, 0, 99, 0, 101, 31, 0, 32,
	23, 24, 26, 21, 0, 10, 13, 15, 51, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 0, 0, 58, 0, 60, 62,
	70, 0, 79, 81, 82, 83, 34, 66, 0, 74,
	0, 0, 0, 95, 0, 0, 100, 0, 33, 0,
	27, 28, 19, 50, 52, 59, 61, 63, 71, 0,
	0, 0, 0, 0, 0, 76, 0, 97, 98, 102,
	25, 80, 84, 67, 75, 72, 0, 0, 0, 0,
	0, 68, 0, 78, 69, 73, 77,
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
	52, 53, 54, 55, 56, 57,
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
		//line unql.y:46
		{ 
		logDebugGrammar("INPUT") 
	}
	case 2:
		//line unql.y:53
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 3:
		//line unql.y:59
		{ 
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 4:
		//line unql.y:66
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 5:
		//line unql.y:72
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 6:
		//line unql.y:78
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
		parsingStatement.SetExplainOnly(true)
	}
	case 7:
		//line unql.y:84
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
	}
	case 8:
		//line unql.y:91
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
		//line unql.y:105
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 10:
		//line unql.y:110
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
		//line unql.y:123
		{
		logDebugGrammar("RESULT STAR")
	}
	case 12:
		//line unql.y:127
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 13:
		//line unql.y:134
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 14:
		//line unql.y:143
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 15:
		//line unql.y:149
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 16:
		//line unql.y:157
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 17:
		//line unql.y:161
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
		//line unql.y:172
		{
		logDebugGrammar("FROM DATASOURCE")
	}
	case 19:
		//line unql.y:176
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
	}
	case 20:
		//line unql.y:182
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 21:
		//line unql.y:186
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
		//line unql.y:200
		{
		
	}
	case 24:
		//line unql.y:206
		{
		
	}
	case 25:
		//line unql.y:210
		{
		
	}
	case 26:
		//line unql.y:215
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
		//line unql.y:226
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
		//line unql.y:237
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
		//line unql.y:249
		{
		
	}
	case 30:
		//line unql.y:253
		{
		
	}
	case 31:
		//line unql.y:257
		{
		
	}
	case 32:
		//line unql.y:263
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
		//line unql.y:274
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
		//line unql.y:287
		{
		logDebugGrammar("EXPRESSION")
	}
	case 35:
		//line unql.y:292
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 36:
		//line unql.y:300
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 37:
		//line unql.y:308
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 38:
		//line unql.y:316
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 39:
		//line unql.y:324
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 40:
		//line unql.y:332
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 41:
		//line unql.y:340
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 42:
		//line unql.y:348
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 43:
		//line unql.y:356
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 44:
		//line unql.y:364
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 45:
		//line unql.y:372
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 46:
		//line unql.y:380
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 47:
		//line unql.y:388
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:396
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:404
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:412
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:420
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:428
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:436
		{
		
	}
	case 54:
		//line unql.y:442
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:449
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:456
		{
		
	}
	case 57:
		//line unql.y:461
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 58:
		//line unql.y:465
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:472
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:479
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:486
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:493
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:500
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:509
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 65:
		//line unql.y:515
		{
		logDebugGrammar("LITERAL")
	}
	case 66:
		//line unql.y:519
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 67:
		//line unql.y:523
		{
		logDebugGrammar("CASE WHEN THEN ELSE END")
	}
	case 68:
		//line unql.y:527
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
	}
	case 69:
		//line unql.y:531
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
	}
	case 70:
		//line unql.y:535
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:541
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:550
		{
		logDebugGrammar("THEN_LIST - SINGLE")
	}
	case 73:
		//line unql.y:554
		{
		logDebugGrammar("THEN_LIST - COMPOUND")
	}
	case 74:
		//line unql.y:560
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 75:
		//line unql.y:564
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 76:
		//line unql.y:570
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
	}
	case 77:
		//line unql.y:574
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
	}
	case 78:
		//line unql.y:578
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
	}
	case 79:
		//line unql.y:585
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 80:
		//line unql.y:590
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
	case 81:
		//line unql.y:604
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 82:
		//line unql.y:608
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 83:
		//line unql.y:617
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 84:
		//line unql.y:623
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 85:
		//line unql.y:633
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line unql.y:639
		{
		logDebugGrammar("NUMBER")
	}
	case 87:
		//line unql.y:643
		{
		logDebugGrammar("OBJECT")
	}
	case 88:
		//line unql.y:647
		{
		logDebugGrammar("ARRAY")
	}
	case 89:
		//line unql.y:651
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line unql.y:657
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line unql.y:663
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line unql.y:671
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line unql.y:677
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line unql.y:685
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 95:
		//line unql.y:691
		{
		logDebugGrammar("OBJECT")
	}
	case 96:
		//line unql.y:697
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 97:
		//line unql.y:701
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 98:
		//line unql.y:713
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 99:
		//line unql.y:723
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray([]ast.Expression{})
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:729
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().([]ast.Expression)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line unql.y:738
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make([]ast.Expression, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 102:
		//line unql.y:745
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
