
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
const GROUP = 57359
const HAVING = 57360
const LBRACE = 57361
const RBRACE = 57362
const LBRACKET = 57363
const RBRACKET = 57364
const COMMA = 57365
const COLON = 57366
const TRUE = 57367
const FALSE = 57368
const NULL = 57369
const INT = 57370
const NUMBER = 57371
const IDENTIFIER = 57372
const STRING = 57373
const PLUS = 57374
const MINUS = 57375
const MULT = 57376
const DIV = 57377
const CONCAT = 57378
const AND = 57379
const OR = 57380
const NOT = 57381
const EQ = 57382
const NE = 57383
const GT = 57384
const GTE = 57385
const LT = 57386
const LTE = 57387
const LPAREN = 57388
const RPAREN = 57389
const LIKE = 57390
const IS = 57391
const VALUED = 57392
const MISSING = 57393
const DOT = 57394
const CASE = 57395
const WHEN = 57396
const THEN = 57397
const ELSE = 57398
const END = 57399
const ANY = 57400
const ALL = 57401
const OVER = 57402
const MOD = 57403

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
	"GROUP",
	"HAVING",
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

const yyNprod = 117
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 635

var yyAct = []int{

	35, 155, 146, 105, 24, 72, 59, 65, 194, 182,
	101, 95, 138, 191, 190, 176, 154, 99, 107, 201,
	64, 109, 78, 79, 80, 81, 83, 84, 85, 93,
	86, 91, 89, 90, 87, 88, 141, 22, 92, 96,
	173, 7, 94, 118, 199, 93, 102, 103, 142, 200,
	104, 82, 139, 198, 31, 96, 193, 197, 139, 144,
	143, 107, 166, 175, 174, 25, 67, 32, 165, 139,
	67, 117, 115, 167, 76, 151, 116, 36, 71, 123,
	124, 125, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 136, 137, 34, 62, 140, 68, 75, 122,
	153, 68, 156, 114, 161, 74, 179, 78, 79, 80,
	81, 83, 64, 163, 93, 97, 98, 160, 164, 100,
	56, 119, 57, 111, 96, 168, 51, 52, 53, 54,
	55, 41, 47, 110, 38, 152, 82, 162, 159, 171,
	37, 80, 81, 83, 66, 170, 93, 43, 153, 153,
	177, 178, 33, 110, 44, 70, 96, 67, 17, 45,
	46, 28, 184, 185, 186, 120, 121, 77, 82, 29,
	14, 18, 3, 16, 30, 11, 11, 10, 10, 112,
	153, 11, 192, 195, 196, 188, 106, 62, 68, 20,
	21, 95, 2, 50, 49, 74, 12, 48, 150, 149,
	156, 202, 78, 79, 80, 81, 83, 84, 85, 93,
	86, 91, 89, 90, 87, 88, 181, 42, 92, 96,
	40, 39, 94, 69, 27, 183, 73, 23, 61, 60,
	58, 82, 19, 9, 169, 8, 15, 187, 26, 110,
	95, 172, 13, 6, 5, 4, 189, 1, 0, 0,
	0, 78, 79, 80, 81, 83, 84, 85, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 94, 0, 0, 0, 0, 0, 0, 0, 0,
	82, 78, 79, 80, 81, 83, 84, 85, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 94, 0, 0, 0, 0, 0, 0, 0, 158,
	82, 78, 79, 80, 81, 83, 84, 85, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 94, 0, 0, 0, 0, 0, 0, 0, 157,
	82, 78, 79, 80, 81, 83, 84, 85, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 94, 0, 0, 0, 0, 0, 0, 0, 0,
	82, 78, 79, 80, 81, 83, 84, 85, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 180, 0, 0, 0, 0, 0, 0, 0, 0,
	82, 78, 79, 80, 81, 83, 84, 85, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 113, 0, 0, 0, 0, 0, 0, 0, 0,
	82, 78, 79, 80, 81, 83, 84, 0, 93, 86,
	91, 89, 90, 87, 88, 0, 0, 92, 96, 0,
	95, 94, 0, 0, 0, 0, 0, 0, 0, 0,
	82, 78, 79, 80, 81, 83, 0, 0, 93, 86,
	91, 89, 90, 87, 88, 147, 148, 92, 96, 0,
	0, 94, 0, 0, 0, 0, 0, 0, 0, 56,
	82, 57, 0, 0, 0, 51, 52, 53, 54, 55,
	41, 47, 0, 38, 152, 0, 0, 0, 0, 37,
	0, 0, 0, 0, 0, 0, 43, 145, 0, 0,
	0, 0, 0, 44, 56, 0, 57, 0, 45, 46,
	51, 52, 53, 54, 55, 41, 47, 0, 38, 63,
	0, 0, 0, 0, 37, 0, 0, 0, 0, 0,
	0, 43, 0, 0, 0, 0, 0, 0, 44, 56,
	0, 57, 108, 45, 46, 51, 52, 53, 54, 55,
	41, 47, 0, 38, 0, 0, 0, 0, 0, 37,
	0, 0, 0, 0, 0, 0, 43, 0, 0, 0,
	0, 0, 0, 44, 56, 0, 57, 0, 45, 46,
	51, 52, 53, 54, 55, 41, 47, 0, 38, 0,
	0, 0, 0, 0, 37, 0, 0, 0, 0, 0,
	0, 43, 0, 0, 0, 0, 0, 0, 44, 0,
	0, 0, 0, 45, 46,
}
var yyPact = []int{

	168, -1000, -1000, 169, -1000, -1000, 159, 164, 161, 184,
	35, -1000, -1000, 146, 157, 161, 35, 135, 575, 505,
	-1000, -1000, -1000, -53, 136, -1000, -1000, 139, 50, 575,
	135, -1000, 174, 155, -1000, 309, -1000, 575, 575, -1000,
	-1000, -29, -1000, 575, -44, 575, 575, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 30, 540, -1000, -1000,
	100, -1000, 171, -1000, 369, 35, 42, 48, 41, -1000,
	15, -1000, -1000, 98, 152, -1000, -1000, 575, 575, 575,
	575, 575, 575, 575, 575, 575, 575, 575, 575, 575,
	575, 575, 575, -36, 39, 575, 9, -1000, -1000, 470,
	-31, 575, 279, 249, -1000, 118, 94, 80, -1000, 115,
	90, 505, 38, 28, -1000, -1000, 51, -1000, -1000, 575,
	-1000, -1000, 127, 107, 107, 6, 6, 6, 6, 429,
	399, 75, 75, 75, 75, 75, 75, 75, 575, -1000,
	219, -1000, 13, -1000, -1000, -1000, -32, 101, 101, 83,
	-1000, -1000, -1000, 339, -1000, -47, 170, 35, 35, -1000,
	-13, 575, -1000, 575, -1000, -1000, -1000, -1000, -1000, -1000,
	575, 75, -1000, -1000, -1000, -1000, -1000, -33, -34, 101,
	22, -49, 575, 575, 49, 45, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 309, -10, 19, -11, 575,
	-1000, -1000, -1000,
}
var yyPgo = []int{

	0, 247, 192, 245, 244, 243, 242, 238, 41, 236,
	158, 67, 235, 21, 234, 75, 233, 232, 230, 6,
	229, 228, 0, 37, 227, 4, 5, 226, 224, 223,
	77, 221, 220, 217, 1, 216, 2, 199, 198, 197,
	194, 193, 3, 186,
}
var yyR1 = []int{

	0, 1, 1, 2, 3, 4, 5, 5, 11, 11,
	14, 14, 8, 16, 17, 17, 17, 18, 19, 19,
	20, 20, 20, 21, 21, 9, 9, 12, 23, 23,
	24, 24, 10, 10, 6, 6, 26, 26, 27, 27,
	27, 7, 7, 7, 28, 29, 15, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 30, 30, 30, 31, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 32, 34, 34, 35, 35,
	25, 25, 25, 36, 36, 37, 37, 38, 38, 33,
	33, 33, 33, 33, 33, 33, 39, 39, 40, 40,
	42, 42, 43, 41, 41, 13, 13,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 3, 4, 4, 0, 4,
	0, 2, 3, 1, 0, 1, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 0, 2, 2, 1, 3,
	1, 3, 0, 2, 0, 3, 1, 3, 1, 2,
	2, 0, 1, 2, 2, 2, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 1, 2, 2, 1, 1, 1, 1, 3, 5,
	6, 6, 3, 4, 5, 5, 3, 5, 0, 2,
	1, 4, 3, 1, 3, 1, 1, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -5, -8, -12, -16,
	9, 7, -2, -6, 11, -9, 9, -10, 10, -17,
	5, 6, -23, -24, -25, 30, -7, -28, 15, 12,
	-10, -23, -11, 17, -15, -22, -30, 39, 33, -31,
	-32, 30, -33, 46, 53, 58, 59, 31, -39, -40,
	-41, 25, 26, 27, 28, 29, 19, 21, -18, -19,
	-20, -21, -15, 34, -22, 60, 8, 21, 52, -29,
	16, 28, -26, -27, -15, -11, -8, 12, 32, 33,
	34, 35, 61, 36, 37, 38, 40, 44, 45, 42,
	43, 41, 48, 39, 52, 21, 49, -30, -30, 46,
	-15, 54, -22, -22, 20, -42, -43, 31, 22, -13,
	-15, 23, 8, 52, -23, 30, 28, 30, 28, 23,
	13, 14, -13, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, 48, 30,
	-22, 27, 39, 51, 50, 47, -36, 5, 6, -37,
	-38, -15, 34, -22, 47, -34, -22, 60, 60, 20,
	23, 24, 22, 23, -19, 30, 34, 22, -26, -14,
	18, -22, 22, 27, 51, 50, 47, -36, -36, 23,
	52, -35, 56, 55, -25, -25, -42, -15, -13, -15,
	47, 47, -36, 34, 57, -22, -22, 8, 8, 54,
	30, 30, -34,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 34, 25, 32, 14,
	0, 13, 2, 41, 0, 32, 0, 8, 0, 0,
	15, 16, 27, 28, 30, 90, 5, 42, 0, 0,
	8, 26, 0, 0, 33, 46, 71, 0, 0, 74,
	75, 76, 77, 0, 0, 0, 0, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 0, 0, 12, 17,
	18, 20, 21, 23, 46, 0, 0, 0, 0, 43,
	0, 44, 35, 36, 38, 6, 7, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 72, 73, 0,
	0, 0, 0, 0, 108, 0, 110, 0, 113, 0,
	115, 0, 0, 0, 29, 31, 0, 92, 45, 0,
	39, 40, 10, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 0, 63,
	0, 65, 0, 67, 69, 82, 0, 0, 0, 93,
	95, 96, 97, 46, 78, 88, 0, 0, 0, 109,
	0, 0, 114, 0, 19, 22, 24, 91, 37, 9,
	0, 62, 64, 66, 68, 70, 83, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 111, 112, 116, 11,
	84, 85, 94, 98, 79, 89, 86, 0, 0, 0,
	80, 81, 87,
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
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
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
		//line unql.y:49
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line unql.y:53
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line unql.y:59
		{ 
		logDebugGrammar("STMT")
	}
	case 4:
		//line unql.y:66
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 5:
		//line unql.y:72
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 6:
		//line unql.y:79
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 7:
		//line unql.y:83
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 8:
		//line unql.y:89
		{
	}
	case 9:
		//line unql.y:92
		{
		
	}
	case 10:
		//line unql.y:98
		{
	}
	case 11:
		//line unql.y:101
		{
		logDebugGrammar("SELECT HAVING - EXPR")
		having_part := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Having = having_part
		default:
			logDebugGrammar("This statement does not support HAVING")
		}
	}
	case 12:
		//line unql.y:114
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 13:
		//line unql.y:120
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 14:
		//line unql.y:126
		{
	}
	case 15:
		//line unql.y:129
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 16:
		//line unql.y:139
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 17:
		//line unql.y:151
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
	case 18:
		//line unql.y:165
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 19:
		//line unql.y:170
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
	case 20:
		//line unql.y:183
		{
		logDebugGrammar("RESULT STAR")
	}
	case 21:
		//line unql.y:187
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 22:
		//line unql.y:194
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 23:
		//line unql.y:203
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 24:
		//line unql.y:209
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 25:
		//line unql.y:217
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 26:
		//line unql.y:221
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
	case 27:
		//line unql.y:233
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
	case 28:
		//line unql.y:245
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 29:
		//line unql.y:249
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 30:
		//line unql.y:259
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 31:
		//line unql.y:265
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 32:
		//line unql.y:273
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 33:
		//line unql.y:277
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
	case 35:
		//line unql.y:291
		{
		
	}
	case 36:
		//line unql.y:297
		{
		
	}
	case 37:
		//line unql.y:301
		{
		
	}
	case 38:
		//line unql.y:306
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
	case 39:
		//line unql.y:317
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
	case 40:
		//line unql.y:328
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
	case 41:
		//line unql.y:340
		{
		
	}
	case 42:
		//line unql.y:344
		{
		
	}
	case 43:
		//line unql.y:348
		{
		
	}
	case 44:
		//line unql.y:354
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 45:
		//line unql.y:365
		{ 
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 46:
		//line unql.y:378
		{
		logDebugGrammar("EXPRESSION")
	}
	case 47:
		//line unql.y:383
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 48:
		//line unql.y:391
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 49:
		//line unql.y:399
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 50:
		//line unql.y:407
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:415
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:423
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:431
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:439
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)}) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:447
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:455
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:463
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:471
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:479
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:487
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:495
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:503
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:511
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:519
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:527
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:534
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:541
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:548
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:555
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:562
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:569
		{
		
	}
	case 72:
		//line unql.y:575
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:582
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:589
		{
		
	}
	case 75:
		//line unql.y:594
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 76:
		//line unql.y:600
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 77:
		//line unql.y:606
		{
		logDebugGrammar("LITERAL")
	}
	case 78:
		//line unql.y:610
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 79:
		//line unql.y:614
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
	case 80:
		//line unql.y:631
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 81:
		//line unql.y:639
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 82:
		//line unql.y:647
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line unql.y:653
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line unql.y:660
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line unql.y:667
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line unql.y:676
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 87:
		//line unql.y:684
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
	case 88:
		//line unql.y:698
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 89:
		//line unql.y:702
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 90:
		//line unql.y:708
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 91:
		//line unql.y:714
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line unql.y:721
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line unql.y:732
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 94:
		//line unql.y:737
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
	case 95:
		//line unql.y:751
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 96:
		//line unql.y:755
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 97:
		//line unql.y:764
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 98:
		//line unql.y:770
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 99:
		//line unql.y:780
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:786
		{
		logDebugGrammar("NUMBER")
	}
	case 101:
		//line unql.y:790
		{
		logDebugGrammar("OBJECT")
	}
	case 102:
		//line unql.y:794
		{
		logDebugGrammar("ARRAY")
	}
	case 103:
		//line unql.y:798
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line unql.y:804
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 105:
		//line unql.y:810
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line unql.y:818
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line unql.y:824
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 108:
		//line unql.y:832
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 109:
		//line unql.y:838
		{
		logDebugGrammar("OBJECT")
	}
	case 110:
		//line unql.y:844
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 111:
		//line unql.y:848
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 112:
		//line unql.y:860
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 113:
		//line unql.y:870
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line unql.y:876
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line unql.y:885
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 116:
		//line unql.y:892
		{ 
		logDebugGrammar("EXPRESSION LIST COMPOUND")
		rest := parsingStack.Pop().(ast.ExpressionList)
		last := parsingStack.Pop()
		new_list := make(ast.ExpressionList, 0, len(rest) + 1)
		new_list = append(new_list, last.(ast.Expression))
		for _, v := range rest {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	}
	goto yystack /* stack new state and value */
}
