
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
const CREATE = 57347
const VIEW = 57348
const INDEX = 57349
const ON = 57350
const DISTINCT = 57351
const UNIQUE = 57352
const SELECT = 57353
const AS = 57354
const FROM = 57355
const WHERE = 57356
const ORDER = 57357
const BY = 57358
const ASC = 57359
const DESC = 57360
const LIMIT = 57361
const OFFSET = 57362
const GROUP = 57363
const HAVING = 57364
const LBRACE = 57365
const RBRACE = 57366
const LBRACKET = 57367
const RBRACKET = 57368
const COMMA = 57369
const COLON = 57370
const TRUE = 57371
const FALSE = 57372
const NULL = 57373
const INT = 57374
const NUMBER = 57375
const IDENTIFIER = 57376
const STRING = 57377
const PLUS = 57378
const MINUS = 57379
const MULT = 57380
const DIV = 57381
const CONCAT = 57382
const AND = 57383
const OR = 57384
const NOT = 57385
const EQ = 57386
const NE = 57387
const GT = 57388
const GTE = 57389
const LT = 57390
const LTE = 57391
const LPAREN = 57392
const RPAREN = 57393
const LIKE = 57394
const IS = 57395
const VALUED = 57396
const MISSING = 57397
const DOT = 57398
const CASE = 57399
const WHEN = 57400
const THEN = 57401
const ELSE = 57402
const END = 57403
const ANY = 57404
const ALL = 57405
const OVER = 57406
const FIRST = 57407
const ARRAY = 57408
const MOD = 57409

var yyToknames = []string{
	"EXPLAIN",
	"CREATE",
	"VIEW",
	"INDEX",
	"ON",
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

const yyNprod = 123
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 824

var yyAct = []int{

	41, 169, 160, 117, 82, 67, 28, 122, 73, 218,
	201, 111, 103, 152, 105, 234, 226, 215, 214, 121,
	119, 195, 106, 168, 72, 88, 89, 90, 91, 93,
	40, 70, 103, 96, 101, 99, 100, 97, 98, 26,
	212, 102, 106, 84, 185, 104, 90, 91, 93, 109,
	155, 103, 112, 113, 114, 115, 92, 110, 192, 153,
	37, 106, 156, 217, 236, 225, 9, 223, 116, 237,
	222, 221, 74, 158, 157, 92, 233, 75, 75, 119,
	75, 194, 193, 75, 75, 75, 132, 184, 42, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 151, 29, 86, 154, 136, 76, 76,
	167, 76, 170, 126, 76, 76, 76, 165, 88, 89,
	90, 91, 93, 153, 72, 103, 231, 183, 230, 181,
	38, 70, 107, 108, 229, 106, 186, 182, 187, 153,
	130, 84, 129, 127, 78, 30, 128, 81, 178, 92,
	198, 180, 177, 190, 133, 123, 179, 176, 189, 39,
	80, 34, 167, 167, 196, 197, 87, 85, 21, 165,
	165, 134, 135, 35, 18, 206, 3, 7, 203, 204,
	205, 208, 207, 13, 7, 12, 209, 22, 36, 20,
	13, 13, 12, 124, 24, 25, 131, 213, 77, 167,
	210, 216, 219, 220, 118, 211, 165, 16, 15, 31,
	2, 58, 57, 56, 14, 164, 163, 200, 48, 46,
	45, 79, 105, 33, 83, 27, 69, 68, 66, 170,
	235, 232, 227, 88, 89, 90, 91, 93, 94, 95,
	103, 96, 101, 99, 100, 97, 98, 23, 11, 102,
	106, 188, 10, 104, 105, 174, 19, 32, 17, 8,
	6, 175, 5, 4, 92, 88, 89, 90, 91, 93,
	94, 95, 103, 96, 101, 99, 100, 97, 98, 1,
	0, 102, 106, 0, 0, 104, 105, 228, 0, 0,
	0, 0, 0, 0, 0, 0, 92, 88, 89, 90,
	91, 93, 94, 95, 103, 96, 101, 99, 100, 97,
	98, 0, 0, 102, 106, 0, 0, 104, 105, 0,
	0, 0, 0, 0, 0, 224, 0, 0, 92, 88,
	89, 90, 91, 93, 94, 95, 103, 96, 101, 99,
	100, 97, 98, 0, 0, 102, 106, 0, 0, 104,
	105, 191, 202, 0, 0, 0, 0, 0, 0, 0,
	92, 88, 89, 90, 91, 93, 94, 95, 103, 96,
	101, 99, 100, 97, 98, 0, 0, 102, 106, 0,
	0, 104, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 92, 88, 89, 90, 91, 93, 94, 95,
	103, 96, 101, 99, 100, 97, 98, 0, 0, 102,
	106, 0, 0, 104, 105, 0, 0, 0, 0, 0,
	0, 173, 0, 0, 92, 88, 89, 90, 91, 93,
	94, 95, 103, 96, 101, 99, 100, 97, 98, 0,
	0, 102, 106, 0, 0, 104, 105, 0, 0, 0,
	0, 0, 0, 172, 0, 0, 92, 88, 89, 90,
	91, 93, 94, 95, 103, 96, 101, 99, 100, 97,
	98, 0, 0, 102, 106, 161, 162, 104, 0, 0,
	0, 0, 0, 0, 0, 171, 0, 0, 92, 64,
	0, 65, 0, 0, 0, 59, 60, 61, 62, 63,
	47, 55, 0, 44, 166, 0, 0, 0, 0, 43,
	0, 0, 0, 0, 0, 0, 49, 159, 0, 0,
	0, 0, 105, 50, 0, 0, 0, 0, 51, 52,
	0, 53, 54, 88, 89, 90, 91, 93, 94, 95,
	103, 96, 101, 99, 100, 97, 98, 0, 0, 102,
	106, 0, 0, 104, 105, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 92, 88, 89, 90, 91, 93,
	94, 95, 103, 96, 101, 99, 100, 97, 98, 0,
	0, 102, 106, 0, 0, 199, 105, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 92, 88, 89, 90,
	91, 93, 94, 95, 103, 96, 101, 99, 100, 97,
	98, 105, 0, 102, 106, 0, 0, 125, 0, 0,
	0, 0, 88, 89, 90, 91, 93, 94, 92, 103,
	96, 101, 99, 100, 97, 98, 0, 0, 102, 106,
	0, 0, 104, 0, 0, 0, 0, 0, 64, 0,
	65, 0, 0, 92, 59, 60, 61, 62, 63, 47,
	55, 0, 44, 166, 0, 0, 0, 0, 43, 0,
	0, 0, 0, 0, 0, 49, 0, 0, 0, 0,
	0, 0, 50, 0, 0, 0, 0, 51, 52, 0,
	53, 54, 64, 0, 65, 0, 0, 0, 59, 60,
	61, 62, 63, 47, 55, 0, 44, 71, 0, 0,
	0, 0, 43, 0, 0, 0, 0, 0, 0, 49,
	0, 0, 0, 0, 0, 0, 50, 0, 0, 0,
	0, 51, 52, 0, 53, 54, 64, 0, 65, 120,
	0, 0, 59, 60, 61, 62, 63, 47, 55, 0,
	44, 0, 0, 0, 0, 0, 43, 0, 0, 0,
	0, 0, 0, 49, 0, 0, 0, 0, 0, 0,
	50, 0, 0, 0, 0, 51, 52, 0, 53, 54,
	64, 0, 65, 0, 0, 0, 59, 60, 61, 62,
	63, 47, 55, 0, 44, 0, 0, 0, 0, 0,
	43, 0, 0, 0, 0, 0, 0, 49, 0, 0,
	0, 0, 0, 0, 50, 0, 0, 0, 0, 51,
	52, 0, 53, 54,
}
var yyPact = []int{

	172, -1000, -1000, 179, -1000, -1000, -1000, 201, 159, 176,
	173, 185, 70, -1000, -1000, 111, 202, 142, 157, 173,
	70, 138, 757, 669, -1000, -1000, -1000, -56, 60, -1000,
	190, 110, -1000, 140, 115, 757, 138, -1000, 180, 150,
	-1000, 497, -1000, 757, 757, -1000, -1000, -1, -1000, 757,
	-47, 757, 757, 757, 757, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 44, 713, -1000, -1000, 128, -1000,
	181, -1000, 561, 70, 109, 114, 108, 106, 188, -1000,
	54, -1000, -1000, 127, 154, -1000, -1000, 757, 757, 757,
	757, 757, 757, 757, 757, 757, 757, 757, 757, 757,
	757, 757, 757, -39, 105, 757, 19, -1000, -1000, 466,
	-28, 757, 421, 389, 357, 197, -1000, 133, 125, 120,
	-1000, 130, 124, 669, 103, 89, -1000, -1000, 61, -1000,
	-6, 102, -1000, 757, -1000, -1000, 136, 8, 8, -31,
	-31, -31, -31, -11, 586, 82, 82, 82, 82, 82,
	82, 82, 757, -1000, 325, -1000, 27, -1000, -1000, -1000,
	-30, 625, 625, 123, -1000, -1000, -1000, 529, -1000, -50,
	293, 70, 70, 70, 757, 70, -1000, -15, 757, -1000,
	757, -1000, -1000, -1000, -1000, 757, -10, -1000, -1000, 757,
	82, -1000, -1000, -1000, -1000, -1000, -33, -34, 625, 25,
	-52, 757, 757, 59, 58, 55, 261, 53, -1000, -1000,
	-1000, -35, 757, -1000, -1000, -1000, -1000, -1000, -1000, 497,
	229, 100, 94, 92, 70, 42, -1000, -36, 757, -1000,
	-1000, -1000, 52, -1000, -1000, -1000, 35, -1000,
}
var yyPgo = []int{

	0, 279, 210, 263, 262, 19, 260, 259, 258, 257,
	66, 256, 168, 130, 252, 251, 7, 248, 247, 228,
	5, 227, 226, 0, 39, 225, 6, 4, 224, 223,
	221, 88, 220, 219, 218, 1, 217, 2, 216, 215,
	213, 212, 211, 3, 204,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 4, 4, 3, 6, 7,
	7, 13, 13, 15, 15, 10, 17, 18, 18, 18,
	19, 20, 20, 21, 21, 21, 22, 22, 11, 11,
	14, 24, 24, 25, 25, 12, 12, 8, 8, 27,
	27, 28, 28, 28, 9, 9, 9, 29, 30, 16,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 31, 31, 31, 32, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 35, 35, 36, 36, 26, 26, 26, 37,
	37, 38, 38, 39, 39, 34, 34, 34, 34, 34,
	34, 34, 40, 40, 41, 41, 43, 43, 44, 42,
	42, 5, 5,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 8, 9, 1, 3, 4,
	4, 0, 4, 0, 2, 3, 1, 0, 1, 1,
	1, 1, 3, 1, 1, 3, 1, 3, 0, 2,
	2, 1, 3, 1, 3, 0, 2, 0, 3, 1,
	3, 1, 2, 2, 0, 1, 2, 2, 2, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 1, 2, 2, 1, 1, 1,
	1, 3, 5, 6, 6, 6, 8, 6, 3, 4,
	5, 5, 3, 5, 0, 2, 1, 4, 3, 1,
	3, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 1, 3, 3, 2,
	3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -6, 5, -7, -10,
	-14, -17, 13, 11, -2, 7, 6, -8, 15, -11,
	13, -12, 14, -18, 9, 10, -24, -25, -26, 34,
	34, 7, -9, -29, 19, 16, -12, -24, -13, 21,
	-16, -23, -31, 43, 37, -32, -33, 34, -34, 50,
	57, 62, 63, 65, 66, 35, -40, -41, -42, 29,
	30, 31, 32, 33, 23, 25, -19, -20, -21, -22,
	-16, 38, -23, 64, 12, 25, 56, 8, 34, -30,
	20, 32, -27, -28, -16, -13, -10, 16, 36, 37,
	38, 39, 67, 40, 41, 42, 44, 48, 49, 46,
	47, 45, 52, 43, 56, 25, 53, -31, -31, 50,
	-16, 58, -23, -23, -23, -23, 24, -43, -44, 35,
	26, -5, -16, 27, 12, 56, -24, 34, 32, 34,
	34, 8, 32, 27, 17, 18, -5, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, 52, 34, -23, 31, 43, 55, 54, 51,
	-37, 9, 10, -38, -39, -16, 38, -23, 51, -35,
	-23, 64, 64, 64, 58, 64, 24, 27, 28, 26,
	27, -20, 34, 38, 26, 50, 34, -27, -15, 22,
	-23, 26, 31, 55, 54, 51, -37, -37, 27, 56,
	-36, 60, 59, -26, -26, -26, -23, -26, -43, -16,
	-5, -5, 50, -16, 51, 51, -37, 38, 61, -23,
	-23, 12, 12, 12, 64, 12, 51, -5, 58, 34,
	34, 34, -26, 34, 51, -35, 12, 34,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 7, 0, 37, 28,
	35, 17, 0, 16, 2, 0, 0, 44, 0, 35,
	0, 11, 0, 0, 18, 19, 30, 31, 33, 96,
	0, 0, 8, 45, 0, 0, 11, 29, 0, 0,
	36, 49, 74, 0, 0, 77, 78, 79, 80, 0,
	0, 0, 0, 0, 0, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 0, 0, 15, 20, 21, 23,
	24, 26, 49, 0, 0, 0, 0, 0, 0, 46,
	0, 47, 38, 39, 41, 9, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 75, 76, 0,
	0, 0, 0, 0, 0, 0, 114, 0, 116, 0,
	119, 0, 121, 0, 0, 0, 32, 34, 0, 98,
	0, 0, 48, 0, 42, 43, 13, 50, 51, 52,
	53, 54, 55, 56, 57, 58, 59, 60, 61, 62,
	63, 64, 0, 66, 0, 68, 0, 70, 72, 88,
	0, 0, 0, 99, 101, 102, 103, 49, 81, 94,
	0, 0, 0, 0, 0, 0, 115, 0, 0, 120,
	0, 22, 25, 27, 97, 0, 0, 40, 12, 0,
	65, 67, 69, 71, 73, 89, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 117, 118,
	122, 0, 0, 14, 90, 91, 100, 104, 82, 95,
	92, 0, 0, 0, 0, 0, 5, 0, 0, 83,
	84, 85, 0, 87, 6, 93, 0, 86,
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
	62, 63, 64, 65, 66, 67,
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
		//line unql.y:50
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line unql.y:54
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line unql.y:60
		{ 
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line unql.y:64
		{
		logDebugGrammar("STMT - CREATE INDEX")
	}
	case 5:
		//line unql.y:71
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-2].s
		name := yyS[yypt-4].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		parsingStatement = createIndexStmt
	}
	case 6:
		//line unql.y:82
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-3].s
		name := yyS[yypt-5].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.View = true
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		parsingStatement = createIndexStmt
	}
	case 7:
		//line unql.y:97
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 8:
		//line unql.y:103
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 9:
		//line unql.y:110
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 10:
		//line unql.y:114
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 11:
		//line unql.y:120
		{
	}
	case 12:
		//line unql.y:123
		{
		group_by := parsingStack.Pop().(ast.ExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.GroupBy = group_by
		default:
			logDebugGrammar("This statement does not support GROUP BY")
		}
	}
	case 13:
		//line unql.y:135
		{
	}
	case 14:
		//line unql.y:138
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
	case 15:
		//line unql.y:151
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 16:
		//line unql.y:157
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 17:
		//line unql.y:163
		{
	}
	case 18:
		//line unql.y:166
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 19:
		//line unql.y:176
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
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
		logDebugGrammar("SELECT SELECT TAIL - EXPR")
		result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Select = result_expr_list
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	
	}
	case 21:
		//line unql.y:202
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 22:
		//line unql.y:207
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
	case 23:
		//line unql.y:220
		{
		logDebugGrammar("RESULT STAR")
	}
	case 24:
		//line unql.y:224
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 25:
		//line unql.y:231
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 26:
		//line unql.y:240
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 27:
		//line unql.y:246
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 28:
		//line unql.y:254
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 29:
		//line unql.y:258
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
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 32:
		//line unql.y:286
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 33:
		//line unql.y:296
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 34:
		//line unql.y:302
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 35:
		//line unql.y:310
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 36:
		//line unql.y:314
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
	case 38:
		//line unql.y:328
		{
		
	}
	case 39:
		//line unql.y:334
		{
		
	}
	case 40:
		//line unql.y:338
		{
		
	}
	case 41:
		//line unql.y:343
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
	case 42:
		//line unql.y:354
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
	case 43:
		//line unql.y:365
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
	case 44:
		//line unql.y:377
		{
		
	}
	case 45:
		//line unql.y:381
		{
		
	}
	case 46:
		//line unql.y:385
		{
		
	}
	case 47:
		//line unql.y:391
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 48:
		//line unql.y:402
		{ 
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 49:
		//line unql.y:415
		{
		logDebugGrammar("EXPRESSION")
	}
	case 50:
		//line unql.y:420
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 51:
		//line unql.y:428
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 52:
		//line unql.y:436
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 53:
		//line unql.y:444
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:452
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:460
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:468
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:476
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:484
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:492
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:500
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:508
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:516
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:524
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:532
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:540
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:548
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:556
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:564
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:571
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:578
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:585
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:592
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:599
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:606
		{
		
	}
	case 75:
		//line unql.y:612
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:619
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line unql.y:626
		{
		
	}
	case 78:
		//line unql.y:631
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 79:
		//line unql.y:637
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 80:
		//line unql.y:643
		{
		logDebugGrammar("LITERAL")
	}
	case 81:
		//line unql.y:647
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 82:
		//line unql.y:651
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
	case 83:
		//line unql.y:668
		{
		logDebugGrammar("ANY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 84:
		//line unql.y:676
		{
		logDebugGrammar("ALL OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionAny)
	}
	case 85:
		//line unql.y:684
		{
		logDebugGrammar("FIRST OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, path, yyS[yypt-0].s)
		parsingStack.Push(collectionFirst)
	}
	case 86:
		//line unql.y:692
		{
		logDebugGrammar("ARRAY WHEN OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, path, yyS[yypt-0].s, output)
		parsingStack.Push(collectionArray)
	}
	case 87:
		//line unql.y:701
		{
		logDebugGrammar("ARRAY OVER AS IDENTIFIER")
		path := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, path, yyS[yypt-0].s, output)
		parsingStack.Push(collectionArray)
	}
	case 88:
		//line unql.y:709
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line unql.y:715
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line unql.y:722
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 91:
		//line unql.y:730
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line unql.y:739
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 93:
		//line unql.y:747
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
	case 94:
		//line unql.y:761
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 95:
		//line unql.y:765
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 96:
		//line unql.y:771
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 97:
		//line unql.y:777
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line unql.y:784
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line unql.y:795
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 100:
		//line unql.y:800
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
	case 101:
		//line unql.y:814
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 102:
		//line unql.y:818
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 103:
		//line unql.y:827
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 104:
		//line unql.y:833
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 105:
		//line unql.y:843
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line unql.y:849
		{
		logDebugGrammar("NUMBER")
	}
	case 107:
		//line unql.y:853
		{
		logDebugGrammar("OBJECT")
	}
	case 108:
		//line unql.y:857
		{
		logDebugGrammar("ARRAY")
	}
	case 109:
		//line unql.y:861
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line unql.y:867
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 111:
		//line unql.y:873
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 112:
		//line unql.y:881
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 113:
		//line unql.y:887
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line unql.y:895
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 115:
		//line unql.y:901
		{
		logDebugGrammar("OBJECT")
	}
	case 116:
		//line unql.y:907
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 117:
		//line unql.y:911
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 118:
		//line unql.y:923
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 119:
		//line unql.y:933
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 120:
		//line unql.y:939
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 121:
		//line unql.y:948
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 122:
		//line unql.y:955
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
