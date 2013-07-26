
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

const yyNprod = 113
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 620

var yyAct = []int{

	34, 150, 141, 105, 146, 24, 71, 58, 101, 22,
	64, 186, 175, 97, 133, 89, 190, 142, 143, 183,
	63, 182, 189, 33, 61, 92, 31, 66, 166, 55,
	136, 56, 169, 66, 73, 50, 51, 52, 53, 54,
	40, 46, 137, 37, 147, 98, 99, 96, 65, 36,
	95, 168, 167, 139, 138, 149, 42, 140, 67, 66,
	100, 106, 134, 43, 67, 103, 185, 35, 44, 45,
	156, 103, 193, 192, 110, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 131, 132,
	67, 134, 135, 25, 160, 161, 148, 114, 151, 74,
	75, 76, 77, 79, 93, 94, 89, 134, 63, 76,
	77, 79, 61, 113, 89, 159, 92, 111, 112, 70,
	73, 172, 163, 158, 92, 155, 115, 107, 78, 162,
	157, 154, 7, 69, 164, 17, 78, 116, 117, 29,
	28, 14, 18, 148, 148, 170, 171, 16, 108, 3,
	32, 30, 11, 11, 10, 10, 11, 102, 177, 178,
	49, 180, 181, 106, 179, 20, 21, 91, 2, 48,
	47, 145, 12, 148, 144, 184, 187, 188, 74, 75,
	76, 77, 79, 80, 81, 89, 82, 87, 85, 86,
	83, 84, 151, 194, 88, 92, 174, 91, 90, 41,
	191, 39, 38, 68, 27, 72, 23, 78, 74, 75,
	76, 77, 79, 80, 81, 89, 82, 87, 85, 86,
	83, 84, 60, 59, 88, 92, 57, 19, 90, 91,
	165, 176, 9, 8, 15, 26, 13, 78, 6, 5,
	74, 75, 76, 77, 79, 80, 81, 89, 82, 87,
	85, 86, 83, 84, 4, 1, 88, 92, 0, 91,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	74, 75, 76, 77, 79, 80, 81, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 91,
	90, 0, 0, 0, 0, 0, 0, 0, 153, 78,
	74, 75, 76, 77, 79, 80, 81, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 91,
	90, 0, 0, 0, 0, 0, 0, 0, 152, 78,
	74, 75, 76, 77, 79, 80, 81, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 91,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	74, 75, 76, 77, 79, 80, 81, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 91,
	173, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	74, 75, 76, 77, 79, 80, 81, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 91,
	109, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	74, 75, 76, 77, 79, 80, 0, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 91,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	74, 75, 76, 77, 79, 0, 0, 89, 82, 87,
	85, 86, 83, 84, 0, 0, 88, 92, 0, 0,
	90, 0, 0, 0, 55, 0, 56, 0, 0, 78,
	50, 51, 52, 53, 54, 40, 46, 0, 37, 147,
	0, 0, 0, 0, 36, 0, 0, 0, 0, 0,
	0, 42, 0, 0, 0, 0, 0, 0, 43, 55,
	0, 56, 0, 44, 45, 50, 51, 52, 53, 54,
	40, 46, 0, 37, 62, 0, 0, 0, 0, 36,
	0, 0, 0, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 43, 55, 0, 56, 104, 44, 45,
	50, 51, 52, 53, 54, 40, 46, 0, 37, 0,
	0, 0, 0, 0, 36, 0, 0, 0, 0, 0,
	0, 42, 0, 0, 0, 0, 0, 0, 43, 55,
	0, 56, 0, 44, 45, 50, 51, 52, 53, 54,
	40, 46, 0, 37, 0, 0, 0, 0, 0, 36,
	0, 0, 0, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 43, 0, 0, 0, 0, 44, 45,
}
var yyPact = []int{

	145, -1000, -1000, 146, -1000, -1000, 130, 138, 132, 160,
	65, -1000, -1000, 125, 127, 132, 65, 149, 562, 492,
	-1000, -1000, -1000, -48, 40, -1000, -1000, 117, 93, 562,
	-1000, -1000, -1000, -1000, 300, -1000, 562, 562, -1000, -1000,
	6, -1000, 562, -39, 562, 562, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 42, 527, -1000, -1000, 106,
	-1000, 140, -1000, 360, 65, 89, 92, 85, -1000, 71,
	-1000, -1000, 105, 124, 562, 562, 562, 562, 562, 562,
	562, 562, 562, 562, 562, 562, 562, 562, 562, -32,
	79, 562, 5, -1000, -1000, 12, 10, 562, 270, 240,
	-1000, 113, 104, 48, -1000, 110, 102, 492, 66, 63,
	-1000, -1000, 109, -1000, -1000, 562, -1000, -1000, 77, 77,
	-22, -22, -22, -22, 420, 390, 69, 69, 69, 69,
	69, 69, 69, 562, -1000, 210, -1000, 3, -1000, -1000,
	-1000, -13, 457, 457, 100, -1000, -1000, -1000, 330, -1000,
	-42, 178, 65, 65, -1000, 36, 562, -1000, 562, -1000,
	-1000, -1000, -1000, -1000, 69, -1000, -1000, -1000, -1000, -1000,
	-24, -26, 457, 34, -44, 562, 562, 14, 8, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 300, 148, 45,
	44, 562, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 255, 168, 254, 239, 238, 236, 235, 132, 234,
	135, 233, 232, 227, 226, 7, 223, 222, 4, 0,
	9, 206, 5, 6, 205, 204, 203, 67, 202, 201,
	199, 1, 196, 2, 174, 171, 170, 169, 160, 8,
	157, 3,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 4, 5, 5, 8, 12,
	13, 13, 13, 14, 15, 15, 16, 16, 16, 17,
	17, 9, 9, 11, 20, 20, 21, 21, 10, 10,
	6, 6, 23, 23, 24, 24, 24, 7, 7, 7,
	25, 26, 18, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 27, 27,
	27, 28, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 31, 31, 32, 32, 22, 22, 22, 33,
	33, 34, 34, 35, 35, 30, 30, 30, 30, 30,
	30, 30, 36, 36, 37, 37, 39, 39, 40, 38,
	38, 41, 41,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 3, 3, 3, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 0, 2, 2, 1, 3, 1, 3, 0, 2,
	0, 3, 1, 3, 1, 2, 2, 0, 1, 2,
	2, 2, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 1, 2, 2,
	1, 1, 1, 1, 3, 5, 6, 6, 3, 4,
	5, 5, 3, 5, 0, 2, 1, 4, 3, 1,
	3, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 1, 3, 3, 2,
	3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -5, -8, -11, -12,
	9, 7, -2, -6, 11, -9, 9, -10, 10, -13,
	5, 6, -20, -21, -22, 28, -7, -25, 15, 12,
	-10, -20, -8, -18, -19, -27, 37, 31, -28, -29,
	28, -30, 44, 51, 56, 57, 29, -36, -37, -38,
	23, 24, 25, 26, 27, 17, 19, -14, -15, -16,
	-17, -18, 32, -19, 58, 8, 19, 50, -26, 16,
	26, -23, -24, -18, 30, 31, 32, 33, 59, 34,
	35, 36, 38, 42, 43, 40, 41, 39, 46, 37,
	50, 19, 47, -27, -27, 44, -18, 52, -19, -19,
	18, -39, -40, 29, 20, -41, -18, 21, 8, 50,
	-20, 28, 26, 28, 26, 21, 13, 14, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, 46, 28, -19, 25, 37, 49, 48,
	45, -33, 5, 6, -34, -35, -18, 32, -19, 45,
	-31, -19, 58, 58, 18, 21, 22, 20, 21, -15,
	28, 32, 20, -23, -19, 20, 25, 49, 48, 45,
	-33, -33, 21, 50, -32, 54, 53, -22, -22, -39,
	-18, -41, 45, 45, -33, 32, 55, -19, -19, 8,
	8, 52, 28, 28, -31,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 30, 21, 28, 10,
	0, 9, 2, 37, 0, 28, 0, 0, 0, 0,
	11, 12, 23, 24, 26, 86, 5, 38, 0, 0,
	6, 22, 7, 29, 42, 67, 0, 0, 70, 71,
	72, 73, 0, 0, 0, 0, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 0, 0, 8, 13, 14,
	16, 17, 19, 42, 0, 0, 0, 0, 39, 0,
	40, 31, 32, 34, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 68, 69, 0, 0, 0, 0, 0,
	104, 0, 106, 0, 109, 0, 111, 0, 0, 0,
	25, 27, 0, 88, 41, 0, 35, 36, 43, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 0, 59, 0, 61, 0, 63, 65,
	78, 0, 0, 0, 89, 91, 92, 93, 42, 74,
	84, 0, 0, 0, 105, 0, 0, 110, 0, 15,
	18, 20, 87, 33, 58, 60, 62, 64, 66, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 107,
	108, 112, 80, 81, 90, 94, 75, 85, 82, 0,
	0, 0, 76, 77, 83,
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
		//line unql.y:48
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line unql.y:52
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line unql.y:58
		{ 
		logDebugGrammar("STMT")
	}
	case 4:
		//line unql.y:65
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 5:
		//line unql.y:71
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 6:
		//line unql.y:78
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 7:
		//line unql.y:82
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 8:
		//line unql.y:88
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 9:
		//line unql.y:94
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 10:
		//line unql.y:100
		{
	}
	case 11:
		//line unql.y:103
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 12:
		//line unql.y:113
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 13:
		//line unql.y:125
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
	case 14:
		//line unql.y:139
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 15:
		//line unql.y:144
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
	case 16:
		//line unql.y:157
		{
		logDebugGrammar("RESULT STAR")
	}
	case 17:
		//line unql.y:161
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 18:
		//line unql.y:168
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 19:
		//line unql.y:177
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 20:
		//line unql.y:183
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 21:
		//line unql.y:191
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 22:
		//line unql.y:195
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
		froms := parsingStack.Pop().([]*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Froms = froms
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 23:
		//line unql.y:207
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
		froms := parsingStack.Pop().([]*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Froms = froms
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 24:
		//line unql.y:219
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
		from_list := make([]*ast.From, 0)
		from_list = append(from_list, parsingStack.Pop().(*ast.From))
		parsingStack.Push(from_list)
	}
	case 25:
		//line unql.y:226
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().([]*ast.From)
		last := parsingStack.Pop()
		new_list := make([]*ast.From, 0, len(rest) + 1)
		new_list = append(new_list, last.(*ast.From))
		for _, v := range rest {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 26:
		//line unql.y:240
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 27:
		//line unql.y:246
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 28:
		//line unql.y:254
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 29:
		//line unql.y:258
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
	case 31:
		//line unql.y:272
		{
		
	}
	case 32:
		//line unql.y:278
		{
		
	}
	case 33:
		//line unql.y:282
		{
		
	}
	case 34:
		//line unql.y:287
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
	case 35:
		//line unql.y:298
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
	case 36:
		//line unql.y:309
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
	case 37:
		//line unql.y:321
		{
		
	}
	case 38:
		//line unql.y:325
		{
		
	}
	case 39:
		//line unql.y:329
		{
		
	}
	case 40:
		//line unql.y:335
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 41:
		//line unql.y:346
		{ 
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 42:
		//line unql.y:359
		{
		logDebugGrammar("EXPRESSION")
	}
	case 43:
		//line unql.y:364
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 44:
		//line unql.y:372
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 45:
		//line unql.y:380
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 46:
		//line unql.y:388
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 47:
		//line unql.y:396
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:404
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:412
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:420
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator([]ast.Expression{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:428
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:436
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:444
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:452
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:460
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:468
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:476
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:484
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:492
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:500
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:508
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:515
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:522
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:529
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:536
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:543
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:550
		{
		
	}
	case 68:
		//line unql.y:556
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:563
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:570
		{
		
	}
	case 71:
		//line unql.y:575
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 72:
		//line unql.y:581
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 73:
		//line unql.y:587
		{
		logDebugGrammar("LITERAL")
	}
	case 74:
		//line unql.y:591
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 75:
		//line unql.y:595
		{
		logDebugGrammar("CASE WHEN THEN ELSE END")
		cwtee := ast.NewCaseOperator()
		topStack := parsingStack.Pop()
		switch topStack := topStack.(type) {
		case ast.Expression:
			cwtee.Else = topStack
			// now look for whenthens
		nextStack := parsingStack.Pop().([]*ast.WhenThen)
			cwtee.WhenThens = nextStack
		case []*ast.WhenThen:
			// no else
		cwtee.WhenThens = topStack
		}
		parsingStack.Push(cwtee)
	}
	case 76:
		//line unql.y:612
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 77:
		//line unql.y:620
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 78:
		//line unql.y:628
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line unql.y:634
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line unql.y:641
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line unql.y:648
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line unql.y:657
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 83:
		//line unql.y:665
		{
		logDebugGrammar("THEN_LIST - COMPOUND")
		rest := parsingStack.Pop().([]*ast.WhenThen)
		last := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		new_list := make([]*ast.WhenThen, 0, len(rest) + 1)
		new_list = append(new_list, &last)
		for _, v := range rest {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 84:
		//line unql.y:679
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 85:
		//line unql.y:683
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 86:
		//line unql.y:689
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 87:
		//line unql.y:695
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line unql.y:702
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line unql.y:713
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 90:
		//line unql.y:718
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
	case 91:
		//line unql.y:732
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 92:
		//line unql.y:736
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 93:
		//line unql.y:745
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 94:
		//line unql.y:751
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 95:
		//line unql.y:761
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line unql.y:767
		{
		logDebugGrammar("NUMBER")
	}
	case 97:
		//line unql.y:771
		{
		logDebugGrammar("OBJECT")
	}
	case 98:
		//line unql.y:775
		{
		logDebugGrammar("ARRAY")
	}
	case 99:
		//line unql.y:779
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:785
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line unql.y:791
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line unql.y:799
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line unql.y:805
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line unql.y:813
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 105:
		//line unql.y:819
		{
		logDebugGrammar("OBJECT")
	}
	case 106:
		//line unql.y:825
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 107:
		//line unql.y:829
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 108:
		//line unql.y:841
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 109:
		//line unql.y:851
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray([]ast.Expression{})
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line unql.y:857
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().([]ast.Expression)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 111:
		//line unql.y:866
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make([]ast.Expression, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 112:
		//line unql.y:873
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
