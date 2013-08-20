
//line unql.y:2
package goyacc
import __yyfmt__ "fmt"
//line unql.y:2
		import "github.com/couchbaselabs/clog"
import "github.com/couchbaselabs/tuqtng/parser"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
    clog.To(parser.PARSER_CHANNEL, format, v...)
}

//line unql.y:12
type yySymType struct {
	yys int 
s string 
n int
f float64}

const EXPLAIN = 57346
const CREATE = 57347
const VIEW = 57348
const INDEX = 57349
const ON = 57350
const USING = 57351
const DISTINCT = 57352
const UNIQUE = 57353
const SELECT = 57354
const AS = 57355
const FROM = 57356
const WHERE = 57357
const ORDER = 57358
const BY = 57359
const ASC = 57360
const DESC = 57361
const LIMIT = 57362
const OFFSET = 57363
const GROUP = 57364
const HAVING = 57365
const LBRACE = 57366
const RBRACE = 57367
const LBRACKET = 57368
const RBRACKET = 57369
const COMMA = 57370
const COLON = 57371
const TRUE = 57372
const FALSE = 57373
const NULL = 57374
const INT = 57375
const NUMBER = 57376
const IDENTIFIER = 57377
const STRING = 57378
const PLUS = 57379
const MINUS = 57380
const MULT = 57381
const DIV = 57382
const CONCAT = 57383
const AND = 57384
const OR = 57385
const NOT = 57386
const EQ = 57387
const NE = 57388
const GT = 57389
const GTE = 57390
const LT = 57391
const LTE = 57392
const LPAREN = 57393
const RPAREN = 57394
const LIKE = 57395
const IS = 57396
const VALUED = 57397
const MISSING = 57398
const DOT = 57399
const CASE = 57400
const WHEN = 57401
const THEN = 57402
const ELSE = 57403
const END = 57404
const ANY = 57405
const ALL = 57406
const OVER = 57407
const FIRST = 57408
const ARRAY = 57409
const MOD = 57410

var yyToknames = []string{
	"EXPLAIN",
	"CREATE",
	"VIEW",
	"INDEX",
	"ON",
	"USING",
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
	"FIRST",
	"ARRAY",
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

const yyNprod = 124
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 753

var yyAct = []int{

	39, 114, 165, 79, 27, 65, 119, 25, 156, 71,
	87, 88, 90, 212, 196, 100, 108, 229, 148, 118,
	219, 217, 100, 70, 220, 103, 209, 35, 38, 68,
	73, 208, 103, 73, 73, 62, 190, 63, 117, 89,
	81, 57, 58, 59, 60, 61, 45, 53, 164, 42,
	109, 110, 111, 112, 107, 41, 181, 216, 215, 116,
	151, 74, 47, 106, 74, 74, 232, 72, 226, 48,
	73, 73, 152, 28, 49, 50, 224, 51, 52, 123,
	73, 149, 187, 154, 153, 211, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 74, 74, 150, 132, 189, 188, 163, 62, 166,
	63, 74, 223, 161, 57, 58, 59, 60, 61, 45,
	53, 70, 42, 162, 9, 149, 177, 68, 41, 179,
	222, 230, 178, 182, 113, 47, 81, 85, 86, 87,
	88, 90, 48, 40, 100, 116, 149, 49, 50, 185,
	51, 52, 127, 126, 103, 124, 29, 36, 163, 163,
	231, 83, 128, 125, 161, 161, 191, 192, 89, 78,
	174, 201, 198, 199, 200, 203, 202, 193, 176, 173,
	129, 204, 120, 180, 175, 104, 105, 172, 184, 37,
	77, 207, 82, 32, 163, 84, 205, 213, 214, 20,
	161, 206, 210, 130, 131, 21, 33, 3, 7, 17,
	19, 121, 7, 13, 102, 13, 75, 12, 34, 13,
	227, 12, 166, 225, 228, 85, 86, 87, 88, 90,
	91, 92, 100, 93, 98, 96, 97, 94, 95, 23,
	24, 99, 103, 15, 2, 101, 102, 170, 14, 115,
	56, 55, 54, 171, 160, 159, 89, 85, 86, 87,
	88, 90, 91, 92, 100, 93, 98, 96, 97, 94,
	95, 195, 46, 99, 103, 44, 43, 101, 102, 221,
	76, 31, 80, 26, 67, 66, 64, 22, 89, 85,
	86, 87, 88, 90, 91, 92, 100, 93, 98, 96,
	97, 94, 95, 11, 183, 99, 103, 10, 18, 101,
	102, 30, 16, 8, 6, 5, 4, 218, 1, 0,
	89, 85, 86, 87, 88, 90, 91, 92, 100, 93,
	98, 96, 97, 94, 95, 0, 0, 99, 103, 0,
	0, 101, 102, 186, 197, 0, 0, 0, 0, 0,
	0, 0, 89, 85, 86, 87, 88, 90, 91, 92,
	100, 93, 98, 96, 97, 94, 95, 0, 0, 99,
	103, 0, 0, 101, 102, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 89, 85, 86, 87, 88, 90,
	91, 92, 100, 93, 98, 96, 97, 94, 95, 0,
	0, 99, 103, 0, 0, 101, 102, 0, 0, 0,
	0, 0, 0, 169, 0, 0, 89, 85, 86, 87,
	88, 90, 91, 92, 100, 93, 98, 96, 97, 94,
	95, 0, 0, 99, 103, 0, 0, 101, 102, 0,
	0, 0, 0, 0, 0, 168, 0, 0, 89, 85,
	86, 87, 88, 90, 91, 92, 100, 93, 98, 96,
	97, 94, 95, 0, 0, 99, 103, 157, 158, 101,
	0, 0, 0, 0, 0, 0, 0, 167, 0, 0,
	89, 62, 0, 63, 0, 0, 0, 57, 58, 59,
	60, 61, 45, 53, 0, 42, 162, 0, 0, 0,
	0, 41, 0, 0, 0, 0, 0, 0, 47, 155,
	0, 0, 0, 0, 102, 48, 0, 0, 0, 0,
	49, 50, 0, 51, 52, 85, 86, 87, 88, 90,
	91, 92, 100, 93, 98, 96, 97, 94, 95, 0,
	0, 99, 103, 0, 0, 101, 102, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 89, 85, 86, 87,
	88, 90, 91, 92, 100, 93, 98, 96, 97, 94,
	95, 0, 0, 99, 103, 0, 0, 194, 102, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 89, 85,
	86, 87, 88, 90, 91, 92, 100, 93, 98, 96,
	97, 94, 95, 102, 0, 99, 103, 0, 0, 122,
	0, 0, 0, 0, 85, 86, 87, 88, 90, 91,
	89, 100, 93, 98, 96, 97, 94, 95, 102, 0,
	99, 103, 0, 0, 101, 0, 0, 0, 0, 85,
	86, 87, 88, 90, 0, 89, 100, 93, 98, 96,
	97, 94, 95, 0, 0, 99, 103, 0, 0, 101,
	0, 0, 0, 0, 0, 62, 0, 63, 0, 0,
	89, 57, 58, 59, 60, 61, 45, 53, 0, 42,
	69, 0, 0, 0, 0, 41, 0, 0, 0, 0,
	0, 0, 47, 0, 0, 0, 0, 0, 0, 48,
	0, 0, 0, 0, 49, 50, 0, 51, 52, 62,
	0, 63, 0, 0, 0, 57, 58, 59, 60, 61,
	45, 53, 0, 42, 0, 0, 0, 0, 0, 41,
	0, 0, 0, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 0, 48, 0, 0, 0, 0, 49, 50,
	0, 51, 52,
}
var yyPact = []int{

	203, -1000, -1000, 207, -1000, -1000, -1000, 236, 193, 196,
	190, 229, 38, -1000, -1000, 121, 173, 189, 190, 38,
	167, 685, 641, -1000, -1000, -1000, -56, 54, -1000, 208,
	-1000, 169, 136, 685, 167, -1000, 201, 178, -1000, 488,
	-1000, 685, 685, -1000, -1000, 12, -1000, 685, -43, 685,
	685, 685, 685, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 109, 11, -1000, -1000, 154, -1000, 198, -1000,
	552, 38, 120, 130, 118, 117, -1000, 129, -1000, -1000,
	152, 185, -1000, -1000, 685, 685, 685, 685, 685, 685,
	685, 685, 685, 685, 685, 685, 685, 685, 685, 685,
	-35, 111, 685, 28, -1000, -1000, 457, -4, 685, 412,
	380, 348, 188, -1000, 162, 151, 141, -1000, 157, 150,
	641, 97, 90, -1000, -1000, 156, -1000, 5, -1000, 685,
	-1000, -1000, 165, -29, -29, -22, -22, -22, -22, 602,
	577, 100, 100, 100, 100, 100, 100, 100, 685, -1000,
	316, -1000, 50, -1000, -1000, -1000, -16, 84, 84, 149,
	-1000, -1000, -1000, 520, -1000, -47, 284, 38, 38, 38,
	685, 38, -1000, 23, 685, -1000, 685, -1000, -1000, -1000,
	-1000, 685, -1000, -1000, 685, 100, -1000, -1000, -1000, -1000,
	-1000, -21, -26, 84, 46, -49, 685, 685, 45, 44,
	8, 252, 7, -1000, -1000, -1000, -28, -1000, -1000, -1000,
	-1000, -1000, -1000, 488, 220, 95, 77, 41, 38, 33,
	211, 685, -1000, -1000, -1000, 4, -1000, 125, -1000, 31,
	-1000, -1000, -1000,
}
var yyPgo = []int{

	0, 318, 244, 316, 315, 19, 314, 313, 312, 311,
	124, 308, 199, 157, 307, 304, 6, 303, 287, 286,
	5, 285, 284, 0, 7, 283, 4, 3, 282, 281,
	280, 143, 276, 275, 272, 2, 271, 8, 255, 254,
	252, 251, 250, 1, 249,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 4, 4, 4, 3, 6,
	7, 7, 13, 13, 15, 15, 10, 17, 18, 18,
	18, 19, 20, 20, 21, 21, 21, 22, 22, 11,
	11, 14, 24, 24, 25, 25, 12, 12, 8, 8,
	27, 27, 28, 28, 28, 9, 9, 9, 29, 30,
	16, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 31, 31, 31, 32,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 35, 35, 36, 36, 26, 26, 26,
	37, 37, 38, 38, 39, 39, 34, 34, 34, 34,
	34, 34, 34, 40, 40, 41, 41, 43, 43, 44,
	42, 42, 5, 5,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 8, 10, 10, 1, 3,
	4, 4, 0, 4, 0, 2, 3, 1, 0, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 0,
	2, 2, 1, 3, 1, 3, 0, 2, 0, 3,
	1, 3, 1, 2, 2, 0, 1, 2, 2, 2,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 1, 2, 2, 1, 1,
	1, 1, 3, 5, 6, 6, 6, 8, 6, 3,
	4, 5, 5, 3, 5, 0, 2, 1, 4, 3,
	1, 3, 1, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 1, 3, 3,
	2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -6, 5, -7, -10,
	-14, -17, 14, 12, -2, 7, -8, 16, -11, 14,
	-12, 15, -18, 10, 11, -24, -25, -26, 35, 35,
	-9, -29, 20, 17, -12, -24, -13, 22, -16, -23,
	-31, 44, 38, -32, -33, 35, -34, 51, 58, 63,
	64, 66, 67, 36, -40, -41, -42, 30, 31, 32,
	33, 34, 24, 26, -19, -20, -21, -22, -16, 39,
	-23, 65, 13, 26, 57, 8, -30, 21, 33, -27,
	-28, -16, -13, -10, 17, 37, 38, 39, 40, 68,
	41, 42, 43, 45, 49, 50, 47, 48, 46, 53,
	44, 57, 26, 54, -31, -31, 51, -16, 59, -23,
	-23, -23, -23, 25, -43, -44, 36, 27, -5, -16,
	28, 13, 57, -24, 35, 33, 35, 35, 33, 28,
	18, 19, -5, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, 53, 35,
	-23, 32, 44, 56, 55, 52, -37, 10, 11, -38,
	-39, -16, 39, -23, 52, -35, -23, 65, 65, 65,
	59, 65, 25, 28, 29, 27, 28, -20, 35, 39,
	27, 51, -27, -15, 23, -23, 27, 32, 56, 55,
	52, -37, -37, 28, 57, -36, 61, 60, -26, -26,
	-26, -23, -26, -43, -16, -5, -5, -16, 52, 52,
	-37, 39, 62, -23, -23, 13, 13, 13, 65, 13,
	52, 59, 35, 35, 35, -26, 35, 9, -35, 13,
	6, 35, 35,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 8, 0, 38, 29,
	36, 18, 0, 17, 2, 0, 45, 0, 36, 0,
	12, 0, 0, 19, 20, 31, 32, 34, 97, 0,
	9, 46, 0, 0, 12, 30, 0, 0, 37, 50,
	75, 0, 0, 78, 79, 80, 81, 0, 0, 0,
	0, 0, 0, 106, 107, 108, 109, 110, 111, 112,
	113, 114, 0, 0, 16, 21, 22, 24, 25, 27,
	50, 0, 0, 0, 0, 0, 47, 0, 48, 39,
	40, 42, 10, 11, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 76, 77, 0, 0, 0, 0,
	0, 0, 0, 115, 0, 117, 0, 120, 0, 122,
	0, 0, 0, 33, 35, 0, 99, 0, 49, 0,
	43, 44, 14, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 0, 67,
	0, 69, 0, 71, 73, 89, 0, 0, 0, 100,
	102, 103, 104, 50, 82, 95, 0, 0, 0, 0,
	0, 0, 116, 0, 0, 121, 0, 23, 26, 28,
	98, 0, 41, 13, 0, 66, 68, 70, 72, 74,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 118, 119, 123, 0, 15, 91, 92,
	101, 105, 83, 96, 93, 0, 0, 0, 0, 0,
	5, 0, 84, 85, 86, 0, 88, 0, 94, 0,
	6, 7, 87,
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
	62, 63, 64, 65, 66, 67, 68,
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
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line unql.y:63
		{
		logDebugGrammar("STMT - CREATE INDEX")
	}
	case 5:
		//line unql.y:70
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-3].s
		name := yyS[yypt-5].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		parsingStatement = createIndexStmt
	}
	case 6:
		//line unql.y:81
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-5].s
		name := yyS[yypt-7].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = "view"
		parsingStatement = createIndexStmt
	}
	case 7:
		//line unql.y:93
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-5].s
		name := yyS[yypt-7].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = yyS[yypt-0].s
		parsingStatement = createIndexStmt
	}
	case 8:
		//line unql.y:109
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 9:
		//line unql.y:115
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 10:
		//line unql.y:122
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 11:
		//line unql.y:126
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 12:
		//line unql.y:132
		{
	}
	case 13:
		//line unql.y:135
		{
		group_by := parsingStack.Pop().(ast.ExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.GroupBy = group_by
		default:
			logDebugGrammar("This statement does not support GROUP BY")
		}
	}
	case 14:
		//line unql.y:147
		{
	}
	case 15:
		//line unql.y:150
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
	case 16:
		//line unql.y:163
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 17:
		//line unql.y:169
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 18:
		//line unql.y:175
		{
	}
	case 19:
		//line unql.y:178
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 20:
		//line unql.y:188
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 21:
		//line unql.y:200
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
	case 22:
		//line unql.y:214
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 23:
		//line unql.y:219
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
	case 24:
		//line unql.y:232
		{
		logDebugGrammar("RESULT STAR")
	}
	case 25:
		//line unql.y:236
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 26:
		//line unql.y:243
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 27:
		//line unql.y:252
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 28:
		//line unql.y:258
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 29:
		//line unql.y:266
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 30:
		//line unql.y:270
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
	case 31:
		//line unql.y:282
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
	case 32:
		//line unql.y:294
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 33:
		//line unql.y:298
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 34:
		//line unql.y:308
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 35:
		//line unql.y:314
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 36:
		//line unql.y:322
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 37:
		//line unql.y:326
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
	case 39:
		//line unql.y:340
		{
		
	}
	case 40:
		//line unql.y:346
		{
		
	}
	case 41:
		//line unql.y:350
		{
		
	}
	case 42:
		//line unql.y:355
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
	case 43:
		//line unql.y:366
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
	case 44:
		//line unql.y:377
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
	case 45:
		//line unql.y:389
		{
		
	}
	case 46:
		//line unql.y:393
		{
		
	}
	case 47:
		//line unql.y:397
		{
		
	}
	case 48:
		//line unql.y:403
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		if yyS[yypt-0].n < 0 {
			panic("LIMIT cannot be negative")
		}
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 49:
		//line unql.y:417
		{ 
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		if yyS[yypt-0].n < 0 {
			panic("OFFSET cannot be negative")
		}
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 50:
		//line unql.y:433
		{
		logDebugGrammar("EXPRESSION")
	}
	case 51:
		//line unql.y:438
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:446
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:454
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:462
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:470
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:478
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:486
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:494
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:502
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:510
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:518
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:526
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:534
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:542
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:550
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:558
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:566
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:574
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:582
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:589
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:596
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:603
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:610
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:617
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line unql.y:624
		{
		
	}
	case 76:
		//line unql.y:630
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line unql.y:637
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line unql.y:644
		{
		
	}
	case 79:
		//line unql.y:649
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 80:
		//line unql.y:655
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 81:
		//line unql.y:661
		{
		logDebugGrammar("LITERAL")
	}
	case 82:
		//line unql.y:665
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 83:
		//line unql.y:669
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
	case 84:
		//line unql.y:686
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 85:
		//line unql.y:694
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 86:
		//line unql.y:702
		{
		logDebugGrammar("FIRST OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionFirst)
	}
	case 87:
		//line unql.y:710
		{
		logDebugGrammar("ARRAY WHEN OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, path, yyS[yypt-0].s, output)
		parsingStack.Push(collectionArray)
	}
	case 88:
		//line unql.y:719
		{
		logDebugGrammar("ARRAY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, path, yyS[yypt-0].s, output)
		parsingStack.Push(collectionArray)
	}
	case 89:
		//line unql.y:727
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line unql.y:733
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line unql.y:740
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 92:
		//line unql.y:748
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line unql.y:757
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 94:
		//line unql.y:765
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
	case 95:
		//line unql.y:779
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 96:
		//line unql.y:783
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 97:
		//line unql.y:789
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 98:
		//line unql.y:795
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line unql.y:802
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:813
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 101:
		//line unql.y:818
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
	case 102:
		//line unql.y:832
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 103:
		//line unql.y:836
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 104:
		//line unql.y:845
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 105:
		//line unql.y:851
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 106:
		//line unql.y:861
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line unql.y:867
		{
		logDebugGrammar("NUMBER")
	}
	case 108:
		//line unql.y:871
		{
		logDebugGrammar("OBJECT")
	}
	case 109:
		//line unql.y:875
		{
		logDebugGrammar("ARRAY")
	}
	case 110:
		//line unql.y:879
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 111:
		//line unql.y:885
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 112:
		//line unql.y:891
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 113:
		//line unql.y:899
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line unql.y:905
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line unql.y:913
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 116:
		//line unql.y:919
		{
		logDebugGrammar("OBJECT")
	}
	case 117:
		//line unql.y:925
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 118:
		//line unql.y:929
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 119:
		//line unql.y:941
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 120:
		//line unql.y:951
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 121:
		//line unql.y:957
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 122:
		//line unql.y:966
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 123:
		//line unql.y:973
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
