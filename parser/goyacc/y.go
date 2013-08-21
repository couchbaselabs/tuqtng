
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
const IN = 57410
const MOD = 57411

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
	"IN",
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

const yyNprod = 127
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 993

var yyAct = []int{

	41, 171, 162, 118, 25, 82, 123, 224, 67, 223,
	89, 90, 91, 92, 94, 222, 221, 104, 74, 218,
	122, 203, 112, 72, 36, 157, 137, 107, 40, 70,
	104, 75, 127, 154, 225, 215, 214, 158, 120, 197,
	107, 84, 93, 170, 76, 187, 106, 110, 160, 159,
	38, 208, 113, 114, 115, 116, 111, 89, 90, 91,
	92, 94, 95, 96, 104, 97, 102, 100, 101, 98,
	99, 9, 117, 103, 107, 77, 42, 105, 207, 128,
	155, 155, 244, 120, 217, 184, 85, 206, 205, 93,
	139, 140, 141, 142, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 239, 29, 156, 183, 138,
	87, 169, 194, 172, 91, 92, 94, 167, 37, 104,
	108, 109, 26, 155, 29, 72, 132, 131, 29, 107,
	129, 70, 185, 182, 240, 196, 195, 86, 73, 30,
	188, 84, 189, 133, 93, 130, 81, 179, 200, 181,
	178, 134, 124, 186, 180, 192, 177, 191, 39, 80,
	135, 136, 20, 33, 169, 169, 198, 199, 88, 34,
	167, 167, 17, 21, 19, 3, 7, 125, 13, 231,
	7, 35, 209, 13, 78, 12, 210, 13, 15, 12,
	23, 24, 2, 119, 58, 57, 14, 56, 213, 166,
	165, 169, 211, 216, 219, 220, 202, 167, 212, 48,
	46, 45, 79, 32, 83, 28, 27, 69, 68, 66,
	22, 11, 227, 228, 229, 230, 190, 172, 232, 106,
	10, 18, 31, 16, 8, 6, 241, 5, 242, 4,
	89, 90, 91, 92, 94, 95, 96, 104, 97, 102,
	100, 101, 98, 99, 1, 0, 103, 107, 0, 0,
	105, 0, 237, 106, 0, 238, 0, 0, 0, 0,
	0, 0, 93, 0, 89, 90, 91, 92, 94, 95,
	96, 104, 97, 102, 100, 101, 98, 99, 0, 0,
	103, 107, 0, 0, 105, 0, 235, 106, 0, 236,
	0, 0, 0, 0, 0, 0, 93, 0, 89, 90,
	91, 92, 94, 95, 96, 104, 97, 102, 100, 101,
	98, 99, 0, 0, 103, 107, 0, 0, 105, 0,
	106, 0, 0, 243, 0, 0, 0, 0, 0, 0,
	93, 89, 90, 91, 92, 94, 95, 96, 104, 97,
	102, 100, 101, 98, 99, 0, 0, 103, 107, 0,
	0, 105, 0, 106, 0, 0, 234, 0, 0, 0,
	0, 0, 0, 93, 89, 90, 91, 92, 94, 95,
	96, 104, 97, 102, 100, 101, 98, 99, 0, 0,
	103, 107, 0, 0, 105, 0, 106, 0, 0, 233,
	0, 0, 0, 0, 0, 0, 93, 89, 90, 91,
	92, 94, 95, 96, 104, 97, 102, 100, 101, 98,
	99, 0, 0, 103, 107, 0, 0, 105, 0, 226,
	106, 0, 0, 0, 0, 0, 0, 0, 0, 93,
	0, 89, 90, 91, 92, 94, 95, 96, 104, 97,
	102, 100, 101, 98, 99, 0, 0, 103, 107, 0,
	0, 105, 0, 0, 204, 106, 193, 0, 0, 0,
	0, 0, 0, 93, 0, 0, 89, 90, 91, 92,
	94, 95, 96, 104, 97, 102, 100, 101, 98, 99,
	0, 0, 103, 107, 0, 0, 105, 0, 106, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 93, 89,
	90, 91, 92, 94, 95, 96, 104, 97, 102, 100,
	101, 98, 99, 0, 0, 103, 107, 0, 0, 105,
	0, 106, 0, 0, 0, 0, 0, 176, 0, 0,
	0, 93, 89, 90, 91, 92, 94, 95, 96, 104,
	97, 102, 100, 101, 98, 99, 0, 0, 103, 107,
	0, 0, 105, 0, 106, 0, 0, 0, 0, 0,
	175, 0, 0, 0, 93, 89, 90, 91, 92, 94,
	95, 96, 104, 97, 102, 100, 101, 98, 99, 0,
	0, 103, 107, 0, 0, 105, 0, 106, 0, 0,
	0, 0, 0, 174, 0, 0, 0, 93, 89, 90,
	91, 92, 94, 95, 96, 104, 97, 102, 100, 101,
	98, 99, 0, 0, 103, 107, 0, 0, 105, 0,
	106, 0, 0, 0, 0, 0, 173, 0, 0, 0,
	93, 89, 90, 91, 92, 94, 95, 96, 104, 97,
	102, 100, 101, 98, 99, 0, 0, 103, 107, 0,
	0, 105, 0, 106, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 89, 90, 91, 92, 94, 95,
	96, 104, 97, 102, 100, 101, 98, 99, 163, 164,
	103, 107, 0, 0, 201, 0, 0, 0, 0, 0,
	0, 0, 64, 0, 65, 0, 93, 0, 59, 60,
	61, 62, 63, 47, 55, 0, 44, 168, 0, 0,
	0, 0, 43, 0, 0, 0, 0, 0, 0, 49,
	161, 0, 0, 0, 0, 106, 50, 0, 0, 0,
	0, 51, 52, 0, 53, 54, 89, 90, 91, 92,
	94, 95, 96, 104, 97, 102, 100, 101, 98, 99,
	0, 106, 103, 107, 0, 0, 126, 0, 0, 0,
	0, 0, 89, 90, 91, 92, 94, 95, 93, 104,
	97, 102, 100, 101, 98, 99, 0, 106, 103, 107,
	0, 0, 105, 0, 0, 0, 0, 0, 89, 90,
	91, 92, 94, 0, 93, 104, 97, 102, 100, 101,
	98, 99, 0, 0, 103, 107, 0, 64, 105, 65,
	0, 0, 0, 59, 60, 61, 62, 63, 47, 55,
	93, 44, 168, 0, 0, 0, 0, 43, 0, 0,
	0, 0, 0, 0, 49, 0, 0, 0, 0, 0,
	0, 50, 0, 0, 0, 0, 51, 52, 0, 53,
	54, 64, 0, 65, 0, 0, 0, 59, 60, 61,
	62, 63, 47, 55, 0, 44, 71, 0, 0, 0,
	0, 43, 0, 0, 0, 0, 0, 0, 49, 0,
	0, 0, 0, 0, 0, 50, 0, 0, 0, 0,
	51, 52, 0, 53, 54, 64, 0, 65, 121, 0,
	0, 59, 60, 61, 62, 63, 47, 55, 0, 44,
	0, 0, 0, 0, 0, 43, 0, 0, 0, 0,
	0, 0, 49, 0, 0, 0, 0, 0, 0, 50,
	0, 0, 0, 0, 51, 52, 0, 53, 54, 64,
	0, 65, 0, 0, 0, 59, 60, 61, 62, 63,
	47, 55, 0, 44, 0, 0, 0, 0, 0, 43,
	0, 0, 0, 0, 0, 0, 49, 0, 0, 0,
	0, 0, 0, 50, 0, 0, 0, 0, 51, 52,
	0, 53, 54,
}
var yyPact = []int{

	171, -1000, -1000, 175, -1000, -1000, -1000, 181, 156, 160,
	158, 180, 93, -1000, -1000, 104, 143, 152, 158, 89,
	136, 925, 837, -1000, -1000, -1000, 103, -47, 18, -1000,
	176, -1000, 138, 113, 925, 136, -1000, 102, 166, 151,
	-1000, 604, -1000, 925, 925, -1000, -1000, -4, -1000, 925,
	-37, 925, 925, 925, 925, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 47, 881, -1000, -1000, 124, -1000,
	164, -1000, 709, -25, 71, 95, 112, 92, 91, -1000,
	110, -1000, -1000, 123, 142, -1000, -31, -1000, 925, 925,
	925, 925, 925, 925, 925, 925, 925, 925, 925, 925,
	925, 925, 925, 925, -20, 88, 925, -7, -1000, -1000,
	678, -9, 925, 571, 538, 505, 472, -1000, 131, 122,
	118, -1000, 127, 121, 837, 73, 46, 71, -1000, -1000,
	126, -1000, -6, -1000, 925, -1000, -1000, 71, 134, 75,
	75, -14, -14, -14, -14, 761, 735, -27, -27, -27,
	-27, -27, -27, -27, 925, -1000, 439, -1000, 80, -1000,
	-1000, -1000, -13, 793, 793, 120, -1000, -1000, -1000, 637,
	-1000, -40, 404, 53, 52, 43, 16, -1000, 2, 925,
	-1000, 925, -1000, -1000, -1000, -1000, -1000, 925, -1000, -1000,
	-1000, 925, -27, -1000, -1000, -1000, -1000, -1000, -16, -17,
	793, 45, -43, 925, 925, -52, -53, -59, -61, -1000,
	-1000, -1000, -18, -1000, -1000, -1000, -1000, -1000, -1000, 604,
	370, 925, 925, 925, 925, 170, 925, 337, 304, 237,
	203, 99, -1000, -1000, -1000, 925, -1000, 925, -1000, -1000,
	-1000, 271, 20, -1000, -1000,
}
var yyPgo = []int{

	0, 254, 192, 239, 237, 20, 235, 234, 233, 232,
	71, 231, 162, 50, 230, 226, 6, 221, 220, 219,
	8, 218, 217, 0, 4, 216, 215, 5, 214, 213,
	212, 76, 211, 210, 209, 1, 206, 2, 200, 199,
	197, 195, 194, 3, 193,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 4, 4, 4, 3, 6,
	7, 7, 13, 13, 15, 15, 10, 17, 18, 18,
	18, 19, 20, 20, 21, 21, 21, 22, 22, 11,
	11, 11, 14, 14, 24, 24, 25, 25, 12, 12,
	8, 8, 27, 27, 28, 28, 28, 9, 9, 9,
	29, 30, 16, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 31, 31,
	31, 32, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 35, 35, 36, 36,
	26, 26, 26, 37, 37, 38, 38, 39, 39, 34,
	34, 34, 34, 34, 34, 34, 40, 40, 41, 41,
	43, 43, 44, 42, 42, 5, 5,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 8, 10, 10, 1, 3,
	4, 4, 0, 4, 0, 2, 3, 1, 0, 1,
	1, 1, 1, 3, 1, 1, 3, 1, 3, 0,
	2, 5, 2, 5, 1, 3, 1, 3, 0, 2,
	0, 3, 1, 3, 1, 2, 2, 0, 1, 2,
	2, 2, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 1, 2, 2,
	1, 1, 1, 1, 3, 5, 7, 7, 9, 7,
	9, 7, 3, 4, 5, 5, 3, 5, 0, 2,
	1, 4, 3, 1, 3, 1, 1, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -6, 5, -7, -10,
	-14, -17, 14, 12, -2, 7, -8, 16, -11, 14,
	-12, 15, -18, 10, 11, -24, 29, -25, -26, 35,
	35, -9, -29, 20, 17, -12, -24, 29, -13, 22,
	-16, -23, -31, 44, 38, -32, -33, 35, -34, 51,
	58, 63, 64, 66, 67, 36, -40, -41, -42, 30,
	31, 32, 33, 34, 24, 26, -19, -20, -21, -22,
	-16, 39, -23, 35, 65, 13, 26, 57, 8, -30,
	21, 33, -27, -28, -16, -13, 35, -10, 17, 37,
	38, 39, 40, 69, 41, 42, 43, 45, 49, 50,
	47, 48, 46, 53, 44, 57, 26, 54, -31, -31,
	51, -16, 59, -23, -23, -23, -23, 25, -43, -44,
	36, 27, -5, -16, 28, 13, 57, 57, -24, 35,
	33, 35, 35, 33, 28, 18, 19, 57, -5, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, 53, 35, -23, 32, 44, 56,
	55, 52, -37, 10, 11, -38, -39, -16, 39, -23,
	52, -35, -23, 65, 65, 65, 65, 25, 28, 29,
	27, 28, -20, 35, 39, -24, 27, 51, -27, -24,
	-15, 23, -23, 27, 32, 56, 55, 52, -37, -37,
	28, 57, -36, 61, 60, 35, 35, 35, 35, -43,
	-16, -5, -5, -16, 52, 52, -37, 39, 62, -23,
	-23, 68, 68, 68, 68, 52, 59, -23, -23, -23,
	-23, 9, -35, 62, 62, 59, 62, 59, 62, 6,
	35, -23, -23, 62, 62,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 8, 0, 40, 29,
	38, 18, 0, 17, 2, 0, 47, 0, 38, 0,
	12, 0, 0, 19, 20, 32, 0, 34, 36, 100,
	0, 9, 48, 0, 0, 12, 30, 0, 0, 0,
	39, 52, 77, 0, 0, 80, 81, 82, 83, 0,
	0, 0, 0, 0, 0, 109, 110, 111, 112, 113,
	114, 115, 116, 117, 0, 0, 16, 21, 22, 24,
	25, 27, 52, 0, 0, 0, 0, 0, 0, 49,
	0, 50, 41, 42, 44, 10, 0, 11, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 78, 79,
	0, 0, 0, 0, 0, 0, 0, 118, 0, 120,
	0, 123, 0, 125, 0, 0, 0, 0, 35, 37,
	0, 102, 0, 51, 0, 45, 46, 0, 14, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
	64, 65, 66, 67, 0, 69, 0, 71, 0, 73,
	75, 92, 0, 0, 0, 103, 105, 106, 107, 52,
	84, 98, 0, 0, 0, 0, 0, 119, 0, 0,
	124, 0, 23, 26, 28, 33, 101, 0, 43, 31,
	13, 0, 68, 70, 72, 74, 76, 93, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 121,
	122, 126, 0, 15, 94, 95, 104, 108, 85, 99,
	96, 0, 0, 0, 0, 5, 0, 0, 0, 0,
	0, 0, 97, 86, 87, 0, 89, 0, 91, 6,
	7, 0, 0, 88, 90,
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
	62, 63, 64, 65, 66, 67, 68, 69,
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
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 31:
		//line unql.y:282
		{
		logDebugGrammar("SELECT FROM - DATASOURCE WITH POOL")
		from := parsingStack.Pop().(*ast.From)
		from.Pool = yyS[yypt-2].s
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
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
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 33:
		//line unql.y:307
		{
		logDebugGrammar("SELECT FROM - DATASOURCE WITH POOL")
		from := parsingStack.Pop().(*ast.From)
		from.Pool = yyS[yypt-2].s
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 34:
		//line unql.y:321
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 35:
		//line unql.y:325
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 36:
		//line unql.y:335
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 37:
		//line unql.y:341
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 38:
		//line unql.y:349
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 39:
		//line unql.y:353
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
	case 41:
		//line unql.y:367
		{
		
	}
	case 42:
		//line unql.y:373
		{
		
	}
	case 43:
		//line unql.y:377
		{
		
	}
	case 44:
		//line unql.y:382
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
	case 45:
		//line unql.y:393
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
	case 46:
		//line unql.y:404
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
	case 47:
		//line unql.y:416
		{
		
	}
	case 48:
		//line unql.y:420
		{
		
	}
	case 49:
		//line unql.y:424
		{
		
	}
	case 50:
		//line unql.y:430
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
	case 51:
		//line unql.y:444
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
	case 52:
		//line unql.y:460
		{
		logDebugGrammar("EXPRESSION")
	}
	case 53:
		//line unql.y:465
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 54:
		//line unql.y:473
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 55:
		//line unql.y:481
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 56:
		//line unql.y:489
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 57:
		//line unql.y:497
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:505
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:513
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:521
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:529
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:537
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:545
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:553
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:561
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:569
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:577
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:585
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:593
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:601
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:609
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:616
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:623
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:630
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line unql.y:637
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:644
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line unql.y:651
		{
		
	}
	case 78:
		//line unql.y:657
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line unql.y:664
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line unql.y:671
		{
		
	}
	case 81:
		//line unql.y:676
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 82:
		//line unql.y:682
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 83:
		//line unql.y:688
		{
		logDebugGrammar("LITERAL")
	}
	case 84:
		//line unql.y:692
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 85:
		//line unql.y:696
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
	case 86:
		//line unql.y:713
		{
		logDebugGrammar("ANY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 87:
		//line unql.y:721
		{
		logDebugGrammar("ALL OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 88:
		//line unql.y:729
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 89:
		//line unql.y:738
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 90:
		//line unql.y:746
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 91:
		//line unql.y:755
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 92:
		//line unql.y:763
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line unql.y:769
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line unql.y:776
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 95:
		//line unql.y:784
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line unql.y:793
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 97:
		//line unql.y:801
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
	case 98:
		//line unql.y:815
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 99:
		//line unql.y:819
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 100:
		//line unql.y:825
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 101:
		//line unql.y:831
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line unql.y:838
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line unql.y:849
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 104:
		//line unql.y:854
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
	case 105:
		//line unql.y:868
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 106:
		//line unql.y:872
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 107:
		//line unql.y:881
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 108:
		//line unql.y:887
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 109:
		//line unql.y:897
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line unql.y:903
		{
		logDebugGrammar("NUMBER")
	}
	case 111:
		//line unql.y:907
		{
		logDebugGrammar("OBJECT")
	}
	case 112:
		//line unql.y:911
		{
		logDebugGrammar("ARRAY")
	}
	case 113:
		//line unql.y:915
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line unql.y:921
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line unql.y:927
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 116:
		//line unql.y:935
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 117:
		//line unql.y:941
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 118:
		//line unql.y:949
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 119:
		//line unql.y:955
		{
		logDebugGrammar("OBJECT")
	}
	case 120:
		//line unql.y:961
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 121:
		//line unql.y:965
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 122:
		//line unql.y:977
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 123:
		//line unql.y:987
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line unql.y:993
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 125:
		//line unql.y:1002
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 126:
		//line unql.y:1009
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
