
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
const DISTINCT = 57347
const UNIQUE = 57348
const SELECT = 57349
const AS = 57350
const FROM = 57351
const WHERE = 57352
const ORDER = 57353
const BY = 57354
const ASC = 57355
const DESC = 57356
const LIMIT = 57357
const OFFSET = 57358
const LBRACE = 57359
const RBRACE = 57360
const LBRACKET = 57361
const RBRACKET = 57362
const COMMA = 57363
const COLON = 57364
const TRUE = 57365
const FALSE = 57366
const NULL = 57367
const INT = 57368
const NUMBER = 57369
const IDENTIFIER = 57370
const STRING = 57371
const PLUS = 57372
const MINUS = 57373
const MULT = 57374
const DIV = 57375
const CONCAT = 57376
const AND = 57377
const OR = 57378
const NOT = 57379
const EQ = 57380
const NE = 57381
const GT = 57382
const GTE = 57383
const LT = 57384
const LTE = 57385
const LPAREN = 57386
const RPAREN = 57387
const LIKE = 57388
const IS = 57389
const VALUED = 57390
const MISSING = 57391
const DOT = 57392
const CASE = 57393
const WHEN = 57394
const THEN = 57395
const ELSE = 57396
const END = 57397
const ANY = 57398
const ALL = 57399
const OVER = 57400
const MOD = 57401

var yyToknames = []string{
	"EXPLAIN",
	"DISTINCT",
	"UNIQUE",
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

const yyNprod = 110
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 621

var yyAct = []int{

	62, 147, 138, 101, 97, 59, 23, 64, 180, 169,
	93, 90, 131, 177, 33, 87, 176, 163, 28, 66,
	146, 91, 25, 34, 86, 187, 71, 72, 73, 74,
	76, 77, 78, 86, 79, 84, 82, 83, 80, 81,
	133, 99, 85, 160, 94, 95, 107, 143, 185, 115,
	67, 96, 134, 179, 115, 75, 186, 26, 114, 88,
	89, 31, 99, 136, 135, 115, 162, 161, 61, 33,
	63, 108, 116, 117, 118, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 128, 129, 130, 112, 132, 92,
	111, 113, 145, 109, 148, 110, 103, 58, 153, 166,
	155, 152, 104, 102, 68, 157, 87, 159, 154, 151,
	156, 57, 105, 106, 10, 19, 31, 71, 72, 73,
	74, 76, 77, 78, 86, 79, 84, 82, 83, 80,
	81, 20, 158, 85, 73, 74, 76, 107, 22, 86,
	145, 145, 164, 165, 12, 184, 75, 71, 72, 73,
	74, 76, 61, 183, 86, 69, 65, 173, 16, 175,
	7, 75, 98, 8, 14, 15, 48, 145, 47, 178,
	181, 182, 171, 172, 46, 142, 75, 141, 168, 40,
	38, 37, 56, 18, 60, 24, 148, 188, 30, 29,
	27, 13, 6, 87, 21, 11, 5, 4, 17, 9,
	3, 174, 2, 102, 71, 72, 73, 74, 76, 77,
	78, 86, 79, 84, 82, 83, 80, 81, 1, 0,
	85, 0, 0, 87, 107, 0, 0, 170, 0, 0,
	0, 0, 0, 75, 71, 72, 73, 74, 76, 77,
	78, 86, 79, 84, 82, 83, 80, 81, 0, 0,
	85, 0, 0, 87, 107, 0, 0, 0, 0, 0,
	0, 0, 150, 75, 71, 72, 73, 74, 76, 77,
	78, 86, 79, 84, 82, 83, 80, 81, 0, 0,
	85, 0, 0, 87, 107, 0, 0, 0, 0, 0,
	0, 0, 149, 75, 71, 72, 73, 74, 76, 77,
	78, 86, 79, 84, 82, 83, 80, 81, 0, 0,
	85, 0, 0, 87, 107, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 71, 72, 73, 74, 76, 77,
	78, 86, 79, 84, 82, 83, 80, 81, 0, 0,
	85, 0, 0, 87, 167, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 71, 72, 73, 74, 76, 77,
	78, 86, 79, 84, 82, 83, 80, 81, 139, 140,
	85, 0, 0, 0, 70, 0, 0, 0, 0, 0,
	54, 0, 55, 75, 0, 0, 49, 50, 51, 52,
	53, 39, 45, 0, 36, 144, 0, 0, 0, 0,
	35, 0, 0, 0, 0, 0, 0, 41, 137, 0,
	0, 0, 0, 0, 42, 0, 0, 0, 0, 43,
	44, 71, 72, 73, 74, 76, 77, 0, 86, 79,
	84, 82, 83, 80, 81, 0, 0, 85, 0, 0,
	0, 0, 0, 0, 0, 54, 0, 55, 0, 0,
	75, 49, 50, 51, 52, 53, 39, 45, 0, 36,
	144, 0, 0, 0, 0, 35, 0, 0, 0, 0,
	0, 0, 41, 0, 0, 0, 0, 0, 0, 42,
	0, 0, 0, 0, 43, 44, 71, 72, 73, 74,
	76, 0, 0, 86, 79, 84, 82, 83, 80, 81,
	0, 0, 85, 0, 0, 0, 0, 0, 0, 0,
	54, 0, 55, 0, 0, 75, 49, 50, 51, 52,
	53, 39, 45, 0, 36, 32, 0, 0, 0, 0,
	35, 0, 0, 0, 0, 0, 0, 41, 0, 0,
	0, 0, 0, 0, 42, 54, 0, 55, 100, 43,
	44, 49, 50, 51, 52, 53, 39, 45, 0, 36,
	0, 0, 0, 0, 0, 35, 0, 0, 0, 0,
	0, 0, 41, 0, 0, 0, 0, 0, 0, 42,
	54, 0, 55, 0, 43, 44, 49, 50, 51, 52,
	53, 39, 45, 0, 36, 0, 0, 0, 0, 0,
	35, 0, 0, 0, 0, 0, 0, 41, 0, 0,
	0, 0, 0, 0, 42, 0, 0, 0, 0, 43,
	44,
}
var yyPact = []int{

	156, -1000, -1000, 103, -1000, 135, 159, 151, -1000, 100,
	119, 128, 29, 493, -1000, -1000, -1000, -1000, 95, 71,
	563, -1000, 563, -1000, -51, 148, 0, -1000, -1000, 83,
	-1000, 147, -1000, 324, -1000, 563, 563, -1000, -36, -23,
	-1000, 563, -42, 563, 563, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 33, 528, -1000, 70, -1000, -1000,
	81, 99, 264, -1000, 29, 65, 69, 29, 493, 63,
	26, 563, 563, 563, 563, 563, 563, 563, 563, 563,
	563, 563, 563, 563, 563, 563, -34, 563, -1000, -1000,
	15, 363, -25, 563, 234, 204, -1000, 91, 80, 76,
	-1000, 88, 79, -1000, 563, -1000, -1000, 37, -1000, -1000,
	85, -1000, -1000, -1000, -1000, -1000, 102, 102, -13, -13,
	-13, -13, 456, 391, 117, 117, 117, 117, 117, 117,
	117, 563, 87, -1000, 18, -1000, -1000, -1000, -28, 428,
	428, 78, -1000, -1000, -1000, 294, -1000, -45, 174, 29,
	29, -1000, 12, 563, -1000, 563, -1000, -1000, 117, -1000,
	-1000, -1000, -1000, -1000, -29, -32, 428, 21, -47, 563,
	563, 145, 137, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 264, -4, 28, -3, 563, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 218, 202, 200, 199, 198, 197, 196, 195, 194,
	192, 191, 190, 18, 189, 188, 47, 0, 6, 185,
	22, 5, 184, 183, 182, 23, 181, 180, 179, 1,
	178, 2, 177, 175, 174, 168, 166, 4, 162, 3,
}
var yyR1 = []int{

	0, 1, 2, 3, 6, 7, 10, 10, 11, 11,
	11, 12, 13, 13, 14, 14, 14, 15, 15, 8,
	8, 18, 18, 19, 19, 9, 9, 4, 4, 21,
	21, 22, 22, 22, 5, 5, 5, 23, 24, 16,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 25,
	25, 25, 26, 26, 26, 26, 26, 26, 26, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 29,
	29, 30, 30, 20, 20, 20, 31, 31, 32, 32,
	33, 33, 28, 28, 28, 28, 28, 28, 28, 34,
	34, 35, 35, 37, 37, 38, 36, 36, 39, 39,
}
var yyR2 = []int{

	0, 1, 3, 1, 3, 3, 2, 1, 0, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 0,
	2, 1, 3, 1, 3, 0, 2, 0, 3, 1,
	3, 1, 2, 2, 0, 1, 2, 2, 2, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 1, 2,
	2, 1, 1, 3, 4, 3, 4, 3, 4, 1,
	1, 3, 5, 6, 6, 3, 4, 5, 5, 3,
	5, 0, 2, 1, 4, 3, 1, 3, 1, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 3, 1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -6, -7, -10, 4, 7, -4,
	11, -8, 9, -11, 5, 6, 7, -5, -23, 15,
	12, -9, 10, -18, -19, -20, 28, -12, -13, -14,
	-15, -16, 32, -17, -25, 37, 31, -26, -27, 28,
	-28, 44, 51, 56, 57, 29, -34, -35, -36, 23,
	24, 25, 26, 27, 17, 19, -24, 16, 26, -21,
	-22, -16, -17, -16, 58, 8, 19, 50, 21, 8,
	50, 30, 31, 32, 33, 59, 34, 35, 36, 38,
	42, 43, 40, 41, 39, 46, 37, 19, -25, -25,
	47, 44, -16, 52, -17, -17, 18, -37, -38, 29,
	20, -39, -16, 26, 21, 13, 14, 50, -18, 28,
	26, -20, -13, 28, 32, 28, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, 46, -17, 25, 37, 49, 48, 45, -31, 5,
	6, -32, -33, -16, 32, -17, 45, -29, -17, 58,
	58, 18, 21, 22, 20, 21, -21, 20, -17, 20,
	25, 49, 48, 45, -31, -31, 21, 50, -30, 54,
	53, -20, -20, -37, -16, -39, 45, 45, -31, 32,
	55, -17, -17, 8, 8, 52, 28, 28, -29,
}
var yyDef = []int{

	0, -2, 1, 27, 3, 19, 8, 0, 7, 34,
	0, 25, 0, 0, 9, 10, 6, 2, 35, 0,
	0, 4, 0, 20, 21, 23, 83, 5, 11, 12,
	14, 15, 17, 39, 58, 0, 0, 61, 62, 69,
	70, 0, 0, 0, 0, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 0, 0, 36, 0, 37, 28,
	29, 31, 39, 26, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 59, 60,
	0, 0, 0, 0, 0, 0, 101, 0, 103, 0,
	106, 0, 108, 38, 0, 32, 33, 0, 22, 24,
	0, 85, 13, 16, 18, 56, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 0, 0, 63, 0, 65, 67, 75, 0, 0,
	0, 86, 88, 89, 90, 39, 71, 81, 0, 0,
	0, 102, 0, 0, 107, 0, 30, 84, 55, 57,
	64, 66, 68, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 104, 105, 109, 77, 78, 87, 91,
	72, 82, 79, 0, 0, 0, 73, 74, 80,
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
	52, 53, 54, 55, 56, 57, 58, 59,
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
		//line unql.y:47
		{ 
		logDebugGrammar("INPUT") 
	}
	case 2:
		//line unql.y:54
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 3:
		//line unql.y:60
		{ 
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 4:
		//line unql.y:67
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 5:
		//line unql.y:73
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 6:
		//line unql.y:79
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
		parsingStatement.SetExplainOnly(true)
	}
	case 7:
		//line unql.y:85
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
		parsingStatement = ast.NewSelectStatement()
	}
	case 8:
		//line unql.y:92
		{
	}
	case 9:
		//line unql.y:95
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
	}
	case 10:
		//line unql.y:99
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
	}
	case 11:
		//line unql.y:105
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
	case 12:
		//line unql.y:119
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 13:
		//line unql.y:124
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
	case 14:
		//line unql.y:137
		{
		logDebugGrammar("RESULT STAR")
	}
	case 15:
		//line unql.y:141
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 16:
		//line unql.y:148
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 17:
		//line unql.y:157
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 18:
		//line unql.y:163
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 19:
		//line unql.y:171
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 20:
		//line unql.y:175
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
		from := parsingStack.Pop().(*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 21:
		//line unql.y:187
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 22:
		//line unql.y:191
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
	}
	case 23:
		//line unql.y:197
		{
		logDebugGrammar("FROM DATASOURCE")
		parsingStack.Push(&ast.From{Bucket: yyS[yypt-0].s})
	}
	case 24:
		//line unql.y:202
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		parsingStack.Push(&ast.From{Bucket: yyS[yypt-2].s})
	}
	case 25:
		//line unql.y:209
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 26:
		//line unql.y:213
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
	case 28:
		//line unql.y:227
		{
		
	}
	case 29:
		//line unql.y:233
		{
		
	}
	case 30:
		//line unql.y:237
		{
		
	}
	case 31:
		//line unql.y:242
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
	case 32:
		//line unql.y:253
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
	case 33:
		//line unql.y:264
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
	case 34:
		//line unql.y:276
		{
		
	}
	case 35:
		//line unql.y:280
		{
		
	}
	case 36:
		//line unql.y:284
		{
		
	}
	case 37:
		//line unql.y:290
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 38:
		//line unql.y:301
		{ 
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 39:
		//line unql.y:314
		{
		logDebugGrammar("EXPRESSION")
	}
	case 40:
		//line unql.y:319
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 41:
		//line unql.y:327
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 42:
		//line unql.y:335
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 43:
		//line unql.y:343
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 44:
		//line unql.y:351
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 45:
		//line unql.y:359
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 46:
		//line unql.y:367
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 47:
		//line unql.y:375
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:383
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:391
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:399
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:407
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:415
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:423
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:431
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:439
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:447
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:455
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:463
		{
		
	}
	case 59:
		//line unql.y:469
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:476
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:483
		{
		
	}
	case 62:
		//line unql.y:488
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 63:
		//line unql.y:492
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:499
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:506
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:513
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:520
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:527
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:536
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 70:
		//line unql.y:542
		{
		logDebugGrammar("LITERAL")
	}
	case 71:
		//line unql.y:546
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 72:
		//line unql.y:550
		{
		logDebugGrammar("CASE WHEN THEN ELSE END")
	}
	case 73:
		//line unql.y:554
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
	}
	case 74:
		//line unql.y:558
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
	}
	case 75:
		//line unql.y:562
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:568
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line unql.y:575
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line unql.y:582
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line unql.y:591
		{
		logDebugGrammar("THEN_LIST - SINGLE")
	}
	case 80:
		//line unql.y:595
		{
		logDebugGrammar("THEN_LIST - COMPOUND")
	}
	case 81:
		//line unql.y:601
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 82:
		//line unql.y:605
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 83:
		//line unql.y:611
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
	}
	case 84:
		//line unql.y:615
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
	}
	case 85:
		//line unql.y:619
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
	}
	case 86:
		//line unql.y:626
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 87:
		//line unql.y:631
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
	case 88:
		//line unql.y:645
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 89:
		//line unql.y:649
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 90:
		//line unql.y:658
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 91:
		//line unql.y:664
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 92:
		//line unql.y:674
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line unql.y:680
		{
		logDebugGrammar("NUMBER")
	}
	case 94:
		//line unql.y:684
		{
		logDebugGrammar("OBJECT")
	}
	case 95:
		//line unql.y:688
		{
		logDebugGrammar("ARRAY")
	}
	case 96:
		//line unql.y:692
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line unql.y:698
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line unql.y:704
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line unql.y:712
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:718
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line unql.y:726
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 102:
		//line unql.y:732
		{
		logDebugGrammar("OBJECT")
	}
	case 103:
		//line unql.y:738
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 104:
		//line unql.y:742
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 105:
		//line unql.y:754
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 106:
		//line unql.y:764
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray([]ast.Expression{})
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line unql.y:770
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().([]ast.Expression)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 108:
		//line unql.y:779
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make([]ast.Expression, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 109:
		//line unql.y:786
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
