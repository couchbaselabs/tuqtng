
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

const yyNprod = 125
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 814

var yyAct = []int{

	40, 116, 168, 80, 27, 25, 121, 66, 159, 72,
	87, 88, 89, 90, 92, 216, 200, 102, 110, 134,
	120, 151, 102, 71, 224, 35, 233, 105, 39, 69,
	223, 221, 105, 220, 213, 219, 73, 212, 194, 74,
	82, 91, 167, 74, 74, 104, 74, 184, 74, 74,
	108, 111, 112, 113, 114, 109, 87, 88, 89, 90,
	92, 93, 94, 102, 95, 100, 98, 99, 96, 97,
	75, 118, 101, 105, 75, 75, 103, 75, 125, 75,
	75, 41, 152, 9, 222, 236, 215, 91, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 230, 28, 153, 228, 135, 154, 166,
	191, 169, 89, 90, 92, 164, 234, 102, 152, 227,
	155, 85, 182, 71, 106, 107, 115, 105, 226, 69,
	180, 157, 156, 193, 192, 185, 36, 118, 82, 181,
	186, 91, 28, 152, 129, 235, 128, 126, 84, 29,
	37, 130, 189, 127, 79, 177, 197, 179, 176, 131,
	122, 166, 166, 183, 178, 175, 188, 164, 164, 195,
	196, 38, 78, 32, 205, 202, 203, 204, 207, 206,
	132, 133, 86, 33, 208, 83, 3, 7, 20, 17,
	7, 21, 19, 123, 13, 211, 12, 13, 166, 12,
	209, 217, 218, 13, 164, 210, 214, 34, 23, 24,
	231, 76, 15, 2, 117, 57, 56, 14, 104, 55,
	163, 162, 199, 47, 45, 44, 169, 229, 232, 87,
	88, 89, 90, 92, 93, 94, 102, 95, 100, 98,
	99, 96, 97, 77, 31, 101, 105, 81, 26, 103,
	104, 173, 68, 67, 65, 22, 11, 174, 187, 10,
	91, 87, 88, 89, 90, 92, 93, 94, 102, 95,
	100, 98, 99, 96, 97, 18, 30, 101, 105, 16,
	8, 103, 104, 225, 6, 5, 4, 1, 0, 0,
	0, 0, 91, 87, 88, 89, 90, 92, 93, 94,
	102, 95, 100, 98, 99, 96, 97, 0, 0, 101,
	105, 0, 0, 103, 104, 190, 201, 0, 0, 0,
	0, 0, 0, 0, 91, 87, 88, 89, 90, 92,
	93, 94, 102, 95, 100, 98, 99, 96, 97, 0,
	0, 101, 105, 0, 0, 103, 104, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 91, 87, 88, 89,
	90, 92, 93, 94, 102, 95, 100, 98, 99, 96,
	97, 0, 0, 101, 105, 0, 0, 103, 104, 0,
	0, 0, 0, 0, 0, 172, 0, 0, 91, 87,
	88, 89, 90, 92, 93, 94, 102, 95, 100, 98,
	99, 96, 97, 0, 0, 101, 105, 0, 0, 103,
	104, 0, 0, 0, 0, 0, 0, 171, 0, 0,
	91, 87, 88, 89, 90, 92, 93, 94, 102, 95,
	100, 98, 99, 96, 97, 0, 0, 101, 105, 160,
	161, 103, 0, 0, 0, 0, 0, 0, 0, 170,
	0, 0, 91, 63, 0, 64, 0, 0, 0, 58,
	59, 60, 61, 62, 46, 54, 0, 43, 165, 0,
	0, 0, 0, 42, 0, 0, 0, 0, 0, 0,
	48, 158, 0, 0, 0, 0, 104, 49, 0, 0,
	0, 0, 50, 51, 0, 52, 53, 87, 88, 89,
	90, 92, 93, 94, 102, 95, 100, 98, 99, 96,
	97, 0, 0, 101, 105, 0, 0, 103, 104, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 91, 87,
	88, 89, 90, 92, 93, 94, 102, 95, 100, 98,
	99, 96, 97, 0, 0, 101, 105, 0, 0, 198,
	104, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	91, 87, 88, 89, 90, 92, 93, 94, 102, 95,
	100, 98, 99, 96, 97, 104, 0, 101, 105, 0,
	0, 124, 0, 0, 0, 0, 87, 88, 89, 90,
	92, 93, 91, 102, 95, 100, 98, 99, 96, 97,
	0, 0, 101, 105, 0, 0, 103, 0, 0, 0,
	0, 0, 63, 0, 64, 0, 0, 91, 58, 59,
	60, 61, 62, 46, 54, 0, 43, 165, 0, 0,
	0, 0, 42, 0, 0, 0, 0, 0, 0, 48,
	0, 0, 0, 0, 0, 104, 49, 0, 0, 0,
	0, 50, 51, 0, 52, 53, 87, 88, 89, 90,
	92, 0, 0, 102, 95, 100, 98, 99, 96, 97,
	0, 0, 101, 105, 0, 0, 103, 0, 0, 0,
	0, 0, 63, 0, 64, 0, 0, 91, 58, 59,
	60, 61, 62, 46, 54, 0, 43, 70, 0, 0,
	0, 0, 42, 0, 0, 0, 0, 0, 0, 48,
	0, 0, 0, 0, 0, 0, 49, 0, 0, 0,
	0, 50, 51, 0, 52, 53, 63, 0, 64, 119,
	0, 0, 58, 59, 60, 61, 62, 46, 54, 0,
	43, 0, 0, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 48, 0, 0, 0, 0, 0, 0,
	49, 0, 0, 0, 0, 50, 51, 0, 52, 53,
	63, 0, 64, 0, 0, 0, 58, 59, 60, 61,
	62, 46, 54, 0, 43, 0, 0, 0, 0, 0,
	42, 0, 0, 0, 0, 0, 0, 48, 0, 0,
	0, 0, 0, 0, 49, 0, 0, 0, 0, 50,
	51, 0, 52, 53,
}
var yyPact = []int{

	182, -1000, -1000, 185, -1000, -1000, -1000, 205, 173, 178,
	176, 198, 69, -1000, -1000, 114, 153, 166, 176, 107,
	149, 746, 658, -1000, -1000, -1000, -56, 23, -1000, 203,
	-1000, 151, 121, 746, 149, -1000, 113, 191, 165, -1000,
	460, -1000, 746, 746, -1000, -1000, -1, -1000, 746, -41,
	746, 746, 746, 746, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 101, 702, -1000, -1000, 132, -1000, 180,
	-1000, 524, 69, 112, 120, 111, 109, -1000, 118, -1000,
	-1000, 131, 162, -1000, -38, -1000, 746, 746, 746, 746,
	746, 746, 746, 746, 746, 746, 746, 746, 746, 746,
	746, 746, -32, 108, 746, 76, -1000, -1000, 429, -10,
	746, 384, 352, 320, 192, -1000, 140, 130, 126, -1000,
	137, 129, 658, 104, 83, -1000, -1000, 136, -1000, -4,
	-1000, 746, -1000, -1000, 69, 143, 73, 73, -22, -22,
	-22, -22, 619, 549, -27, -27, -27, -27, -27, -27,
	-27, 746, -1000, 288, -1000, 78, -1000, -1000, -1000, -14,
	588, 588, 128, -1000, -1000, -1000, 492, -1000, -45, 256,
	69, 69, 69, 746, 69, -1000, 35, 746, -1000, 746,
	-1000, -1000, -1000, -1000, 746, -1000, -1000, -1000, 746, -27,
	-1000, -1000, -1000, -1000, -1000, -15, -18, 588, 47, -47,
	746, 746, 22, 20, 18, 19, 17, -1000, -1000, -1000,
	-28, -1000, -1000, -1000, -1000, -1000, -1000, 460, 224, 93,
	84, 71, 69, 68, 201, 746, -1000, -1000, -1000, 13,
	-1000, 110, -1000, 50, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 287, 213, 286, 285, 20, 284, 280, 279, 276,
	83, 275, 188, 150, 259, 258, 6, 256, 255, 254,
	7, 253, 252, 0, 5, 248, 4, 3, 247, 244,
	243, 81, 225, 224, 223, 2, 222, 8, 221, 220,
	219, 216, 215, 1, 214,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 4, 4, 4, 3, 6,
	7, 7, 13, 13, 15, 15, 10, 17, 18, 18,
	18, 19, 20, 20, 21, 21, 21, 22, 22, 11,
	11, 11, 14, 24, 24, 25, 25, 12, 12, 8,
	8, 27, 27, 28, 28, 28, 9, 9, 9, 29,
	30, 16, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 31, 31, 31,
	32, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 35, 35, 36, 36, 26, 26,
	26, 37, 37, 38, 38, 39, 39, 34, 34, 34,
	34, 34, 34, 34, 40, 40, 41, 41, 43, 43,
	44, 42, 42, 5, 5,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 8, 10, 10, 1, 3,
	4, 4, 0, 4, 0, 2, 3, 1, 0, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 0,
	2, 5, 2, 1, 3, 1, 3, 0, 2, 0,
	3, 1, 3, 1, 2, 2, 0, 1, 2, 2,
	2, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 1, 2, 2, 1,
	1, 1, 1, 3, 5, 6, 6, 6, 8, 6,
	3, 4, 5, 5, 3, 5, 0, 2, 1, 4,
	3, 1, 3, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 3, 1, 3,
	3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -6, 5, -7, -10,
	-14, -17, 14, 12, -2, 7, -8, 16, -11, 14,
	-12, 15, -18, 10, 11, -24, -25, -26, 35, 35,
	-9, -29, 20, 17, -12, -24, 29, -13, 22, -16,
	-23, -31, 44, 38, -32, -33, 35, -34, 51, 58,
	63, 64, 66, 67, 36, -40, -41, -42, 30, 31,
	32, 33, 34, 24, 26, -19, -20, -21, -22, -16,
	39, -23, 65, 13, 26, 57, 8, -30, 21, 33,
	-27, -28, -16, -13, 35, -10, 17, 37, 38, 39,
	40, 68, 41, 42, 43, 45, 49, 50, 47, 48,
	46, 53, 44, 57, 26, 54, -31, -31, 51, -16,
	59, -23, -23, -23, -23, 25, -43, -44, 36, 27,
	-5, -16, 28, 13, 57, -24, 35, 33, 35, 35,
	33, 28, 18, 19, 57, -5, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, 53, 35, -23, 32, 44, 56, 55, 52, -37,
	10, 11, -38, -39, -16, 39, -23, 52, -35, -23,
	65, 65, 65, 59, 65, 25, 28, 29, 27, 28,
	-20, 35, 39, 27, 51, -27, -24, -15, 23, -23,
	27, 32, 56, 55, 52, -37, -37, 28, 57, -36,
	61, 60, -26, -26, -26, -23, -26, -43, -16, -5,
	-5, -16, 52, 52, -37, 39, 62, -23, -23, 13,
	13, 13, 65, 13, 52, 59, 35, 35, 35, -26,
	35, 9, -35, 13, 6, 35, 35,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 8, 0, 39, 29,
	37, 18, 0, 17, 2, 0, 46, 0, 37, 0,
	12, 0, 0, 19, 20, 32, 33, 35, 98, 0,
	9, 47, 0, 0, 12, 30, 0, 0, 0, 38,
	51, 76, 0, 0, 79, 80, 81, 82, 0, 0,
	0, 0, 0, 0, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 0, 0, 16, 21, 22, 24, 25,
	27, 51, 0, 0, 0, 0, 0, 48, 0, 49,
	40, 41, 43, 10, 0, 11, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 78, 0, 0,
	0, 0, 0, 0, 0, 116, 0, 118, 0, 121,
	0, 123, 0, 0, 0, 34, 36, 0, 100, 0,
	50, 0, 44, 45, 0, 14, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	66, 0, 68, 0, 70, 0, 72, 74, 90, 0,
	0, 0, 101, 103, 104, 105, 51, 83, 96, 0,
	0, 0, 0, 0, 0, 117, 0, 0, 122, 0,
	23, 26, 28, 99, 0, 42, 31, 13, 0, 67,
	69, 71, 73, 75, 91, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 119, 120, 124,
	0, 15, 92, 93, 102, 106, 84, 97, 94, 0,
	0, 0, 0, 0, 5, 0, 85, 86, 87, 0,
	89, 0, 95, 0, 6, 7, 88,
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
		//line unql.y:267
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 30:
		//line unql.y:271
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
		from.Pool = yyS[yypt-2].s
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 32:
		//line unql.y:296
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
	case 33:
		//line unql.y:309
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 34:
		//line unql.y:313
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 35:
		//line unql.y:323
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 36:
		//line unql.y:329
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 37:
		//line unql.y:337
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 38:
		//line unql.y:341
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
	case 40:
		//line unql.y:355
		{
		
	}
	case 41:
		//line unql.y:361
		{
		
	}
	case 42:
		//line unql.y:365
		{
		
	}
	case 43:
		//line unql.y:370
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
	case 44:
		//line unql.y:381
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
	case 45:
		//line unql.y:392
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
	case 46:
		//line unql.y:404
		{
		
	}
	case 47:
		//line unql.y:408
		{
		
	}
	case 48:
		//line unql.y:412
		{
		
	}
	case 49:
		//line unql.y:418
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
	case 50:
		//line unql.y:432
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
	case 51:
		//line unql.y:448
		{
		logDebugGrammar("EXPRESSION")
	}
	case 52:
		//line unql.y:453
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:461
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:469
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:477
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:485
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:493
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:501
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:509
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:517
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:525
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:533
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:541
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:549
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:557
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:565
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:573
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:581
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:589
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:597
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:604
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:611
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:618
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:625
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line unql.y:632
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:639
		{
		
	}
	case 77:
		//line unql.y:645
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line unql.y:652
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line unql.y:659
		{
		
	}
	case 80:
		//line unql.y:664
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 81:
		//line unql.y:670
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 82:
		//line unql.y:676
		{
		logDebugGrammar("LITERAL")
	}
	case 83:
		//line unql.y:680
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 84:
		//line unql.y:684
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
	case 85:
		//line unql.y:701
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 86:
		//line unql.y:709
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 87:
		//line unql.y:717
		{
		logDebugGrammar("FIRST OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionFirst)
	}
	case 88:
		//line unql.y:725
		{
		logDebugGrammar("ARRAY WHEN OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, path, yyS[yypt-0].s, output)
		parsingStack.Push(collectionArray)
	}
	case 89:
		//line unql.y:734
		{
		logDebugGrammar("ARRAY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, path, yyS[yypt-0].s, output)
		parsingStack.Push(collectionArray)
	}
	case 90:
		//line unql.y:742
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line unql.y:748
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line unql.y:755
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 93:
		//line unql.y:763
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line unql.y:772
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 95:
		//line unql.y:780
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
	case 96:
		//line unql.y:794
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 97:
		//line unql.y:798
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 98:
		//line unql.y:804
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 99:
		//line unql.y:810
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:817
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line unql.y:828
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 102:
		//line unql.y:833
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
	case 103:
		//line unql.y:847
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 104:
		//line unql.y:851
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 105:
		//line unql.y:860
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 106:
		//line unql.y:866
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 107:
		//line unql.y:876
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 108:
		//line unql.y:882
		{
		logDebugGrammar("NUMBER")
	}
	case 109:
		//line unql.y:886
		{
		logDebugGrammar("OBJECT")
	}
	case 110:
		//line unql.y:890
		{
		logDebugGrammar("ARRAY")
	}
	case 111:
		//line unql.y:894
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 112:
		//line unql.y:900
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 113:
		//line unql.y:906
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line unql.y:914
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line unql.y:920
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 116:
		//line unql.y:928
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 117:
		//line unql.y:934
		{
		logDebugGrammar("OBJECT")
	}
	case 118:
		//line unql.y:940
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 119:
		//line unql.y:944
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 120:
		//line unql.y:956
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 121:
		//line unql.y:966
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 122:
		//line unql.y:972
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 123:
		//line unql.y:981
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 124:
		//line unql.y:988
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
