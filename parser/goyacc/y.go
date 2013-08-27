
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
const DROP = 57348
const VIEW = 57349
const INDEX = 57350
const ON = 57351
const USING = 57352
const DISTINCT = 57353
const UNIQUE = 57354
const SELECT = 57355
const AS = 57356
const FROM = 57357
const WHERE = 57358
const ORDER = 57359
const BY = 57360
const ASC = 57361
const DESC = 57362
const LIMIT = 57363
const OFFSET = 57364
const GROUP = 57365
const HAVING = 57366
const LBRACE = 57367
const RBRACE = 57368
const LBRACKET = 57369
const RBRACKET = 57370
const COMMA = 57371
const COLON = 57372
const TRUE = 57373
const FALSE = 57374
const NULL = 57375
const INT = 57376
const NUMBER = 57377
const IDENTIFIER = 57378
const STRING = 57379
const PLUS = 57380
const MINUS = 57381
const MULT = 57382
const DIV = 57383
const CONCAT = 57384
const AND = 57385
const OR = 57386
const NOT = 57387
const EQ = 57388
const NE = 57389
const GT = 57390
const GTE = 57391
const LT = 57392
const LTE = 57393
const LPAREN = 57394
const RPAREN = 57395
const LIKE = 57396
const IS = 57397
const VALUED = 57398
const MISSING = 57399
const DOT = 57400
const CASE = 57401
const WHEN = 57402
const THEN = 57403
const ELSE = 57404
const END = 57405
const ANY = 57406
const ALL = 57407
const OVER = 57408
const FIRST = 57409
const ARRAY = 57410
const IN = 57411
const MOD = 57412

var yyToknames = []string{
	"EXPLAIN",
	"CREATE",
	"DROP",
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

const yyNprod = 135
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 936

var yyAct = []int{

	46, 260, 182, 79, 31, 126, 131, 28, 173, 90,
	241, 72, 240, 69, 239, 70, 130, 238, 197, 64,
	65, 66, 67, 68, 52, 60, 77, 49, 179, 80,
	41, 45, 75, 48, 235, 217, 120, 81, 229, 228,
	54, 148, 143, 135, 85, 165, 92, 55, 266, 82,
	82, 114, 56, 57, 243, 58, 59, 121, 122, 123,
	124, 119, 97, 98, 99, 100, 102, 103, 104, 112,
	105, 110, 108, 109, 106, 107, 112, 252, 111, 115,
	83, 83, 113, 232, 258, 231, 115, 259, 80, 211,
	181, 199, 118, 208, 101, 43, 47, 245, 150, 151,
	152, 153, 154, 155, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 149, 11, 167, 210, 209, 128, 180,
	69, 183, 70, 244, 32, 178, 64, 65, 66, 67,
	68, 52, 60, 77, 49, 76, 93, 125, 166, 75,
	48, 168, 234, 196, 193, 116, 117, 54, 128, 261,
	222, 221, 92, 169, 55, 202, 203, 166, 95, 56,
	57, 195, 58, 59, 171, 170, 206, 141, 42, 220,
	219, 35, 201, 140, 32, 180, 180, 34, 262, 29,
	200, 178, 178, 212, 213, 32, 194, 166, 142, 139,
	137, 136, 94, 86, 78, 223, 33, 224, 144, 190,
	138, 89, 226, 97, 98, 99, 100, 102, 214, 225,
	112, 192, 230, 189, 145, 180, 227, 132, 236, 237,
	115, 178, 198, 233, 191, 188, 205, 44, 88, 23,
	242, 146, 147, 96, 38, 101, 39, 20, 24, 247,
	248, 249, 250, 3, 8, 9, 22, 183, 15, 253,
	133, 40, 15, 269, 14, 8, 9, 264, 251, 265,
	84, 114, 18, 15, 17, 14, 26, 27, 127, 263,
	63, 270, 97, 98, 99, 100, 102, 103, 104, 112,
	105, 110, 108, 109, 106, 107, 2, 62, 111, 115,
	16, 61, 113, 177, 256, 114, 176, 257, 216, 53,
	51, 50, 87, 37, 101, 91, 97, 98, 99, 100,
	102, 103, 104, 112, 105, 110, 108, 109, 106, 107,
	30, 74, 111, 115, 73, 71, 113, 25, 114, 13,
	204, 268, 12, 21, 36, 19, 10, 7, 101, 97,
	98, 99, 100, 102, 103, 104, 112, 105, 110, 108,
	109, 106, 107, 6, 5, 111, 115, 4, 1, 113,
	0, 114, 0, 0, 267, 0, 0, 0, 0, 0,
	0, 101, 97, 98, 99, 100, 102, 103, 104, 112,
	105, 110, 108, 109, 106, 107, 0, 0, 111, 115,
	0, 0, 113, 0, 114, 0, 0, 255, 0, 0,
	0, 0, 0, 0, 101, 97, 98, 99, 100, 102,
	103, 104, 112, 105, 110, 108, 109, 106, 107, 0,
	0, 111, 115, 0, 0, 113, 0, 114, 0, 0,
	254, 0, 0, 0, 0, 0, 0, 101, 97, 98,
	99, 100, 102, 103, 104, 112, 105, 110, 108, 109,
	106, 107, 0, 0, 111, 115, 0, 0, 113, 0,
	246, 114, 0, 0, 0, 0, 0, 0, 0, 0,
	101, 0, 97, 98, 99, 100, 102, 103, 104, 112,
	105, 110, 108, 109, 106, 107, 0, 0, 111, 115,
	0, 0, 113, 0, 0, 218, 114, 207, 0, 0,
	0, 0, 0, 0, 101, 0, 0, 97, 98, 99,
	100, 102, 103, 104, 112, 105, 110, 108, 109, 106,
	107, 0, 0, 111, 115, 0, 0, 113, 0, 114,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 101,
	97, 98, 99, 100, 102, 103, 104, 112, 105, 110,
	108, 109, 106, 107, 0, 0, 111, 115, 0, 0,
	113, 0, 114, 0, 0, 0, 0, 0, 187, 0,
	0, 0, 101, 97, 98, 99, 100, 102, 103, 104,
	112, 105, 110, 108, 109, 106, 107, 0, 0, 111,
	115, 0, 0, 113, 0, 114, 0, 0, 0, 0,
	0, 186, 0, 0, 0, 101, 97, 98, 99, 100,
	102, 103, 104, 112, 105, 110, 108, 109, 106, 107,
	0, 0, 111, 115, 0, 0, 113, 0, 114, 0,
	0, 0, 0, 0, 185, 0, 0, 0, 101, 97,
	98, 99, 100, 102, 103, 104, 112, 105, 110, 108,
	109, 106, 107, 0, 0, 111, 115, 0, 0, 113,
	0, 114, 0, 0, 0, 0, 0, 184, 0, 0,
	0, 101, 97, 98, 99, 100, 102, 103, 104, 112,
	105, 110, 108, 109, 106, 107, 0, 0, 111, 115,
	0, 0, 113, 0, 114, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 101, 97, 98, 99, 100, 102,
	103, 104, 112, 105, 110, 108, 109, 106, 107, 174,
	175, 111, 115, 0, 0, 215, 0, 0, 0, 0,
	0, 0, 0, 69, 0, 70, 0, 101, 0, 64,
	65, 66, 67, 68, 52, 60, 0, 49, 179, 0,
	0, 0, 0, 48, 0, 0, 0, 0, 0, 0,
	54, 172, 0, 0, 0, 0, 114, 55, 0, 0,
	0, 0, 56, 57, 0, 58, 59, 97, 98, 99,
	100, 102, 103, 104, 112, 105, 110, 108, 109, 106,
	107, 0, 114, 111, 115, 0, 0, 134, 0, 0,
	0, 0, 0, 97, 98, 99, 100, 102, 103, 101,
	112, 105, 110, 108, 109, 106, 107, 0, 114, 111,
	115, 0, 0, 113, 0, 0, 0, 0, 0, 97,
	98, 99, 100, 102, 0, 101, 112, 105, 110, 108,
	109, 106, 107, 0, 0, 111, 115, 0, 69, 113,
	70, 129, 0, 0, 64, 65, 66, 67, 68, 52,
	60, 101, 49, 0, 0, 99, 100, 102, 48, 0,
	112, 0, 0, 0, 0, 54, 0, 0, 0, 0,
	115, 0, 55, 0, 0, 0, 0, 56, 57, 0,
	58, 59, 69, 0, 70, 101, 0, 0, 64, 65,
	66, 67, 68, 52, 60, 0, 49, 0, 0, 0,
	0, 0, 48, 0, 0, 0, 0, 0, 0, 54,
	0, 0, 0, 0, 0, 0, 55, 0, 0, 0,
	0, 56, 57, 0, 58, 59,
}
var yyPact = []int{

	239, -1000, -1000, 250, -1000, -1000, -1000, -1000, 256, 254,
	220, 231, 222, 255, 149, -1000, -1000, 160, 141, 213,
	218, 222, 138, 204, 867, 95, -1000, -1000, -1000, 158,
	-37, 23, -1000, 251, -14, 157, -1000, 206, 167, 867,
	204, -1000, 156, 235, 215, -1000, 634, -1000, 867, 867,
	-1000, -1000, 40, -1000, 867, -24, 867, 867, 867, 867,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 111,
	823, -1000, -1000, 188, -1000, 236, -1000, 739, -15, -1000,
	155, 154, 166, 153, 137, 152, -16, -1000, 164, -1000,
	-1000, 185, 212, -1000, -17, -1000, 867, 867, 867, 867,
	867, 867, 867, 867, 867, 867, 867, 867, 867, 867,
	867, 867, -9, 151, 867, 108, -1000, -1000, 708, 37,
	867, 601, 568, 535, 502, -1000, 199, 184, 169, -1000,
	196, 182, 95, 150, 121, 88, -51, -1000, 194, -1000,
	39, 144, -1000, 136, -1000, 867, -1000, -1000, 88, 202,
	825, 825, 31, 31, 31, 31, 791, 765, 165, 165,
	165, 165, 165, 165, 165, 867, -1000, 469, -1000, 60,
	-1000, -1000, -1000, 36, -12, -12, 179, -1000, -1000, -1000,
	667, -1000, -27, 434, 134, 133, 115, 114, -1000, 81,
	867, -1000, 867, -1000, -1000, -1000, -1000, 88, -1000, 867,
	-19, -20, -1000, -1000, -1000, 867, 165, -1000, -1000, -1000,
	-1000, -1000, 32, 30, -12, 102, -29, 867, 867, -52,
	-55, -57, -59, -1000, -1000, -1000, 22, 1, 87, 61,
	-1000, -1000, -1000, -1000, -1000, -1000, 634, 400, 867, 867,
	867, 867, -1000, 248, 25, -1000, 867, 367, 334, 234,
	24, 142, 867, -1000, -1000, -1000, 867, -1000, 867, -1000,
	-1000, -1000, -1000, -5, 301, 268, 243, -1000, -1000, 142,
	-1000,
}
var yyPgo = []int{

	0, 358, 286, 357, 354, 353, 16, 1, 337, 336,
	335, 334, 114, 333, 229, 95, 332, 330, 6, 329,
	327, 325, 11, 324, 321, 0, 7, 320, 3, 4,
	9, 305, 303, 302, 96, 301, 300, 299, 2, 298,
	8, 296, 293, 291, 287, 270, 5, 268,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 4, 4,
	7, 7, 5, 5, 3, 8, 9, 9, 15, 15,
	17, 17, 12, 19, 20, 20, 20, 21, 22, 22,
	23, 23, 23, 24, 24, 13, 13, 13, 16, 16,
	26, 26, 28, 28, 27, 27, 14, 14, 10, 10,
	30, 30, 31, 31, 31, 11, 11, 11, 32, 33,
	18, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 34, 34, 34, 35,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 38, 38, 39, 39, 29, 29,
	29, 40, 40, 41, 41, 42, 42, 37, 37, 37,
	37, 37, 37, 37, 43, 43, 44, 44, 46, 46,
	47, 45, 45, 6, 6,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 8, 11, 10, 13,
	1, 1, 5, 8, 1, 3, 4, 4, 0, 4,
	0, 2, 3, 1, 0, 1, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 0, 2, 5, 2, 5,
	1, 2, 4, 5, 1, 3, 0, 2, 0, 3,
	1, 3, 1, 2, 2, 0, 1, 2, 2, 2,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 1, 2, 2, 1, 1,
	1, 1, 3, 5, 7, 7, 9, 7, 9, 7,
	3, 4, 5, 5, 3, 5, 0, 2, 1, 4,
	3, 1, 3, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 3, 1, 3,
	3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -5, -8, 5, 6,
	-9, -12, -16, -19, 15, 13, -2, 8, 8, -10,
	17, -13, 15, -14, 16, -20, 11, 12, -26, 30,
	-27, -29, 36, 36, 36, 30, -11, -32, 21, 18,
	-14, -26, 30, -15, 23, -18, -25, -34, 45, 39,
	-35, -36, 36, -37, 52, 59, 64, 65, 67, 68,
	37, -43, -44, -45, 31, 32, 33, 34, 35, 25,
	27, -21, -22, -23, -24, -18, 40, -25, 36, -28,
	66, 14, 27, 58, 9, 58, 36, -33, 22, 34,
	-30, -31, -18, -15, 36, -12, 18, 38, 39, 40,
	41, 70, 42, 43, 44, 46, 50, 51, 48, 49,
	47, 54, 45, 58, 27, 55, -34, -34, 52, -18,
	60, -25, -25, -25, -25, 26, -46, -47, 37, 28,
	-6, -18, 29, 14, 58, 58, 36, 36, 34, 36,
	36, 30, 36, 58, 34, 29, 19, 20, 58, -6,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, 54, 36, -25, 33, 45,
	57, 56, 53, -40, 11, 12, -41, -42, -18, 40,
	-25, 53, -38, -25, 66, 66, 66, 66, 26, 29,
	30, 28, 29, -22, 36, 40, -26, 69, 28, 52,
	36, 36, -30, -26, -17, 24, -25, 28, 33, 57,
	56, 53, -40, -40, 29, 58, -39, 62, 61, 36,
	36, 36, 36, -46, -18, -6, -29, -6, 58, 58,
	-18, 53, 53, -40, 40, 63, -25, -25, 69, 69,
	69, 69, -28, 53, 36, 36, 60, -25, -25, -25,
	-25, 10, 52, -38, 63, 63, 60, 63, 60, 63,
	-7, 7, 36, -6, -25, -25, 53, 63, 63, 10,
	-7,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 14, 0, 0,
	48, 35, 46, 24, 0, 23, 2, 0, 0, 55,
	0, 46, 0, 18, 0, 0, 25, 26, 38, 0,
	40, 44, 108, 0, 0, 0, 15, 56, 0, 0,
	18, 36, 0, 0, 0, 47, 60, 85, 0, 0,
	88, 89, 90, 91, 0, 0, 0, 0, 0, 0,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 0,
	0, 22, 27, 28, 30, 31, 33, 60, 0, 41,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 58,
	49, 50, 52, 16, 0, 17, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 86, 87, 0, 0,
	0, 0, 0, 0, 0, 126, 0, 128, 0, 131,
	0, 133, 0, 0, 0, 0, 0, 45, 0, 110,
	0, 0, 12, 0, 59, 0, 53, 54, 0, 20,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 0, 77, 0, 79, 0,
	81, 83, 100, 0, 0, 0, 111, 113, 114, 115,
	60, 92, 106, 0, 0, 0, 0, 0, 127, 0,
	0, 132, 0, 29, 32, 34, 39, 0, 109, 0,
	0, 0, 51, 37, 19, 0, 76, 78, 80, 82,
	84, 101, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 129, 130, 134, 42, 0, 0, 0,
	21, 102, 103, 112, 116, 93, 107, 104, 0, 0,
	0, 0, 43, 6, 0, 13, 0, 0, 0, 0,
	0, 0, 0, 105, 94, 95, 0, 97, 0, 99,
	8, 10, 11, 0, 0, 0, 7, 96, 98, 0,
	9,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70,
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
		//line unql.y:67
		{
		logDebugGrammar("STMT - DROP INDEX")
	}
	case 6:
		//line unql.y:74
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
	case 7:
		//line unql.y:85
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-3].s
		pool := yyS[yypt-5].s
		name := yyS[yypt-8].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		parsingStatement = createIndexStmt
	}
	case 8:
		//line unql.y:98
		{
		method := parsingStack.Pop().(string)
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-5].s
		name := yyS[yypt-7].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = method
		parsingStatement = createIndexStmt
	}
	case 9:
		//line unql.y:111
		{
		method := parsingStack.Pop().(string)
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-5].s
		pool := yyS[yypt-7].s
		name := yyS[yypt-10].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = method
		parsingStatement = createIndexStmt
	}
	case 10:
		//line unql.y:128
		{
		parsingStack.Push("view")
	}
	case 11:
		//line unql.y:132
		{
		parsingStack.Push(yyS[yypt-0].s)
	}
	case 12:
		//line unql.y:138
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 13:
		//line unql.y:147
		{
		bucket := yyS[yypt-2].s
		pool := yyS[yypt-4].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Pool = pool
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 14:
		//line unql.y:161
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 15:
		//line unql.y:167
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 16:
		//line unql.y:174
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 17:
		//line unql.y:178
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 18:
		//line unql.y:184
		{
	}
	case 19:
		//line unql.y:187
		{
		group_by := parsingStack.Pop().(ast.ExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.GroupBy = group_by
		default:
			logDebugGrammar("This statement does not support GROUP BY")
		}
	}
	case 20:
		//line unql.y:199
		{
	}
	case 21:
		//line unql.y:202
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
	case 22:
		//line unql.y:215
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 23:
		//line unql.y:221
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 24:
		//line unql.y:227
		{
	}
	case 25:
		//line unql.y:230
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 26:
		//line unql.y:240
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 27:
		//line unql.y:252
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
	case 28:
		//line unql.y:266
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 29:
		//line unql.y:271
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
	case 30:
		//line unql.y:284
		{
		logDebugGrammar("RESULT STAR")
	}
	case 31:
		//line unql.y:288
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 32:
		//line unql.y:295
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 33:
		//line unql.y:304
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 34:
		//line unql.y:310
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 35:
		//line unql.y:319
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 36:
		//line unql.y:323
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
	case 37:
		//line unql.y:334
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
	case 38:
		//line unql.y:348
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
	case 39:
		//line unql.y:359
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
	case 40:
		//line unql.y:373
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 41:
		//line unql.y:377
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 42:
		//line unql.y:387
		{
		logDebugGrammar("OVER IN")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s})
	}
	case 43:
		//line unql.y:393
		{
		logDebugGrammar("OVER IN nested")
		rest := parsingStack.Pop().(*ast.From)
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Over:rest})
	}
	case 44:
		//line unql.y:402
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 45:
		//line unql.y:408
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 46:
		//line unql.y:416
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 47:
		//line unql.y:420
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
	case 49:
		//line unql.y:434
		{
		
	}
	case 50:
		//line unql.y:440
		{
		
	}
	case 51:
		//line unql.y:444
		{
		
	}
	case 52:
		//line unql.y:449
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
	case 53:
		//line unql.y:460
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
	case 54:
		//line unql.y:471
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
	case 55:
		//line unql.y:483
		{
		
	}
	case 56:
		//line unql.y:487
		{
		
	}
	case 57:
		//line unql.y:491
		{
		
	}
	case 58:
		//line unql.y:497
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
	case 59:
		//line unql.y:511
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
	case 60:
		//line unql.y:527
		{
		logDebugGrammar("EXPRESSION")
	}
	case 61:
		//line unql.y:532
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:540
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:548
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:556
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:564
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:572
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:580
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:588
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:596
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:604
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:612
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:620
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:628
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:636
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line unql.y:644
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:652
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line unql.y:660
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line unql.y:668
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line unql.y:676
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line unql.y:683
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line unql.y:690
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line unql.y:697
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line unql.y:704
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line unql.y:711
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line unql.y:718
		{
		
	}
	case 86:
		//line unql.y:724
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 87:
		//line unql.y:731
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line unql.y:738
		{
		
	}
	case 89:
		//line unql.y:743
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 90:
		//line unql.y:749
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 91:
		//line unql.y:755
		{
		logDebugGrammar("LITERAL")
	}
	case 92:
		//line unql.y:759
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 93:
		//line unql.y:763
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
	case 94:
		//line unql.y:780
		{
		logDebugGrammar("ANY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 95:
		//line unql.y:788
		{
		logDebugGrammar("ALL OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 96:
		//line unql.y:796
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 97:
		//line unql.y:805
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 98:
		//line unql.y:813
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 99:
		//line unql.y:822
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 100:
		//line unql.y:830
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line unql.y:836
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line unql.y:843
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 103:
		//line unql.y:851
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line unql.y:860
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 105:
		//line unql.y:868
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
	case 106:
		//line unql.y:882
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 107:
		//line unql.y:886
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 108:
		//line unql.y:892
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 109:
		//line unql.y:898
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line unql.y:905
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 111:
		//line unql.y:916
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 112:
		//line unql.y:921
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
	case 113:
		//line unql.y:935
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 114:
		//line unql.y:939
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 115:
		//line unql.y:948
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 116:
		//line unql.y:954
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 117:
		//line unql.y:964
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 118:
		//line unql.y:970
		{
		logDebugGrammar("NUMBER")
	}
	case 119:
		//line unql.y:974
		{
		logDebugGrammar("OBJECT")
	}
	case 120:
		//line unql.y:978
		{
		logDebugGrammar("ARRAY")
	}
	case 121:
		//line unql.y:982
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 122:
		//line unql.y:988
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 123:
		//line unql.y:994
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line unql.y:1002
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 125:
		//line unql.y:1008
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 126:
		//line unql.y:1016
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 127:
		//line unql.y:1022
		{
		logDebugGrammar("OBJECT")
	}
	case 128:
		//line unql.y:1028
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 129:
		//line unql.y:1032
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 130:
		//line unql.y:1044
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 131:
		//line unql.y:1054
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 132:
		//line unql.y:1060
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 133:
		//line unql.y:1069
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 134:
		//line unql.y:1076
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
