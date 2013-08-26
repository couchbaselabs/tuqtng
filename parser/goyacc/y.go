
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

const yyNprod = 131
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1002

var yyAct = []int{

	45, 178, 78, 31, 124, 28, 129, 88, 169, 71,
	95, 96, 97, 98, 100, 233, 232, 110, 97, 98,
	100, 231, 230, 110, 193, 128, 76, 113, 40, 79,
	227, 44, 74, 113, 118, 211, 110, 144, 80, 133,
	84, 161, 99, 235, 195, 90, 113, 224, 99, 81,
	112, 81, 202, 223, 205, 177, 119, 120, 121, 122,
	117, 95, 96, 97, 98, 100, 101, 102, 110, 103,
	108, 106, 107, 104, 105, 204, 203, 109, 113, 116,
	82, 111, 82, 247, 46, 11, 248, 162, 79, 42,
	162, 226, 126, 99, 191, 32, 146, 147, 148, 149,
	150, 151, 152, 153, 154, 155, 156, 157, 158, 159,
	160, 123, 164, 163, 249, 140, 216, 176, 215, 179,
	145, 41, 126, 174, 165, 214, 213, 32, 93, 91,
	190, 76, 114, 115, 162, 167, 166, 74, 29, 192,
	189, 139, 138, 250, 32, 137, 135, 134, 90, 196,
	197, 92, 77, 34, 33, 136, 87, 186, 208, 188,
	185, 141, 200, 130, 194, 187, 184, 199, 43, 86,
	23, 176, 176, 142, 143, 94, 37, 174, 174, 206,
	207, 38, 20, 24, 22, 3, 8, 9, 131, 15,
	217, 83, 39, 218, 15, 241, 14, 220, 26, 27,
	8, 9, 18, 2, 17, 125, 222, 16, 15, 176,
	14, 62, 228, 229, 219, 174, 61, 225, 60, 173,
	172, 221, 210, 234, 52, 50, 49, 85, 36, 89,
	30, 237, 238, 239, 240, 73, 72, 179, 242, 112,
	70, 25, 13, 198, 12, 21, 251, 35, 252, 19,
	95, 96, 97, 98, 100, 101, 102, 110, 103, 108,
	106, 107, 104, 105, 10, 7, 109, 113, 6, 5,
	111, 4, 245, 112, 1, 246, 0, 0, 0, 0,
	0, 0, 99, 0, 95, 96, 97, 98, 100, 101,
	102, 110, 103, 108, 106, 107, 104, 105, 0, 0,
	109, 113, 0, 0, 111, 0, 112, 0, 0, 254,
	0, 0, 0, 0, 0, 0, 99, 95, 96, 97,
	98, 100, 101, 102, 110, 103, 108, 106, 107, 104,
	105, 0, 0, 109, 113, 0, 0, 111, 0, 112,
	0, 0, 253, 0, 0, 0, 0, 0, 0, 99,
	95, 96, 97, 98, 100, 101, 102, 110, 103, 108,
	106, 107, 104, 105, 0, 0, 109, 113, 0, 0,
	111, 0, 112, 0, 0, 244, 0, 0, 0, 0,
	0, 0, 99, 95, 96, 97, 98, 100, 101, 102,
	110, 103, 108, 106, 107, 104, 105, 0, 0, 109,
	113, 0, 0, 111, 0, 112, 0, 0, 243, 0,
	0, 0, 0, 0, 0, 99, 95, 96, 97, 98,
	100, 101, 102, 110, 103, 108, 106, 107, 104, 105,
	0, 0, 109, 113, 0, 0, 111, 0, 236, 112,
	0, 0, 0, 0, 0, 0, 0, 0, 99, 0,
	95, 96, 97, 98, 100, 101, 102, 110, 103, 108,
	106, 107, 104, 105, 0, 0, 109, 113, 0, 0,
	111, 0, 0, 212, 112, 201, 0, 0, 0, 0,
	0, 0, 99, 0, 0, 95, 96, 97, 98, 100,
	101, 102, 110, 103, 108, 106, 107, 104, 105, 0,
	0, 109, 113, 0, 0, 111, 0, 112, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 99, 95, 96,
	97, 98, 100, 101, 102, 110, 103, 108, 106, 107,
	104, 105, 0, 0, 109, 113, 0, 0, 111, 0,
	112, 0, 0, 0, 0, 0, 183, 0, 0, 0,
	99, 95, 96, 97, 98, 100, 101, 102, 110, 103,
	108, 106, 107, 104, 105, 0, 0, 109, 113, 0,
	0, 111, 0, 112, 0, 0, 0, 0, 0, 182,
	0, 0, 0, 99, 95, 96, 97, 98, 100, 101,
	102, 110, 103, 108, 106, 107, 104, 105, 0, 0,
	109, 113, 0, 0, 111, 0, 112, 0, 0, 0,
	0, 0, 181, 0, 0, 0, 99, 95, 96, 97,
	98, 100, 101, 102, 110, 103, 108, 106, 107, 104,
	105, 0, 0, 109, 113, 0, 0, 111, 0, 112,
	0, 0, 0, 0, 0, 180, 0, 0, 0, 99,
	95, 96, 97, 98, 100, 101, 102, 110, 103, 108,
	106, 107, 104, 105, 0, 0, 109, 113, 0, 0,
	111, 0, 112, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 99, 95, 96, 97, 98, 100, 101, 102,
	110, 103, 108, 106, 107, 104, 105, 170, 171, 109,
	113, 0, 0, 209, 0, 0, 0, 0, 0, 0,
	0, 68, 0, 69, 0, 99, 0, 63, 64, 65,
	66, 67, 51, 59, 0, 48, 175, 0, 0, 0,
	0, 47, 0, 0, 0, 0, 0, 0, 53, 168,
	0, 0, 0, 0, 112, 54, 0, 0, 0, 0,
	55, 56, 0, 57, 58, 95, 96, 97, 98, 100,
	101, 102, 110, 103, 108, 106, 107, 104, 105, 0,
	112, 109, 113, 0, 0, 132, 0, 0, 0, 0,
	0, 95, 96, 97, 98, 100, 101, 99, 110, 103,
	108, 106, 107, 104, 105, 0, 112, 109, 113, 0,
	0, 111, 0, 0, 0, 0, 0, 95, 96, 97,
	98, 100, 0, 99, 110, 103, 108, 106, 107, 104,
	105, 0, 0, 109, 113, 0, 68, 111, 69, 0,
	0, 0, 63, 64, 65, 66, 67, 51, 59, 99,
	48, 175, 0, 0, 0, 0, 47, 0, 0, 0,
	0, 0, 0, 53, 0, 0, 0, 0, 0, 0,
	54, 0, 0, 0, 0, 55, 56, 0, 57, 58,
	68, 0, 69, 0, 0, 0, 63, 64, 65, 66,
	67, 51, 59, 0, 48, 75, 0, 0, 0, 0,
	47, 0, 0, 0, 0, 0, 0, 53, 0, 0,
	0, 0, 0, 0, 54, 0, 0, 0, 0, 55,
	56, 0, 57, 58, 68, 0, 69, 127, 0, 0,
	63, 64, 65, 66, 67, 51, 59, 0, 48, 0,
	0, 0, 0, 0, 47, 0, 0, 0, 0, 0,
	0, 53, 0, 0, 0, 0, 0, 0, 54, 0,
	0, 0, 0, 55, 56, 0, 57, 58, 68, 0,
	69, 0, 0, 0, 63, 64, 65, 66, 67, 51,
	59, 0, 48, 0, 0, 0, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 53, 0, 0, 0, 0,
	0, 0, 54, 0, 0, 0, 0, 55, 56, 0,
	57, 58,
}
var yyPact = []int{

	181, -1000, -1000, 195, -1000, -1000, -1000, -1000, 196, 194,
	165, 169, 167, 187, 108, -1000, -1000, 118, 117, 155,
	163, 167, 91, 145, 933, 845, -1000, -1000, -1000, 116,
	-37, 24, -1000, 182, -18, -1000, 147, 122, 933, 145,
	-1000, 115, 176, 157, -1000, 612, -1000, 933, 933, -1000,
	-1000, 27, -1000, 933, -26, 933, 933, 933, 933, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 85, 889,
	-1000, -1000, 134, -1000, 174, -1000, 717, -19, -1000, 111,
	110, 121, 109, 106, 105, -1000, 81, -1000, -1000, 132,
	154, -1000, -21, -1000, 933, 933, 933, 933, 933, 933,
	933, 933, 933, 933, 933, 933, 933, 933, 933, 933,
	-13, 98, 933, 79, -1000, -1000, 686, 2, 933, 579,
	546, 513, 480, -1000, 140, 131, 127, -1000, 137, 130,
	845, 94, 54, 59, -45, -1000, 136, -1000, -8, -1000,
	-1000, 933, -1000, -1000, 59, 143, -22, -22, -9, -9,
	-9, -9, 769, 743, -28, -28, -28, -28, -28, -28,
	-28, 933, -1000, 447, -1000, 19, -1000, -1000, -1000, 1,
	801, 801, 129, -1000, -1000, -1000, 645, -1000, -27, 412,
	90, 89, 82, 80, -1000, 55, 933, -1000, 933, -1000,
	-1000, -1000, -1000, 59, -1000, 933, -1000, -1000, -1000, 933,
	-28, -1000, -1000, -1000, -1000, -1000, 0, -6, 801, 51,
	-33, 933, 933, -47, -48, -53, -54, -1000, -1000, -1000,
	22, -10, -1000, -1000, -1000, -1000, -1000, -1000, 612, 378,
	933, 933, 933, 933, -1000, 185, 933, 345, 312, 212,
	23, 107, -1000, -1000, -1000, 933, -1000, 933, -1000, -1000,
	-1000, 279, 246, -1000, -1000,
}
var yyPgo = []int{

	0, 274, 203, 271, 269, 268, 25, 265, 264, 249,
	247, 85, 245, 170, 89, 244, 243, 6, 242, 241,
	240, 9, 236, 235, 0, 5, 230, 2, 3, 7,
	229, 228, 227, 84, 226, 225, 224, 1, 222, 8,
	220, 219, 218, 216, 211, 4, 205,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 4, 5,
	3, 7, 8, 8, 14, 14, 16, 16, 11, 18,
	19, 19, 19, 20, 21, 21, 22, 22, 22, 23,
	23, 12, 12, 12, 15, 15, 25, 25, 27, 27,
	26, 26, 13, 13, 9, 9, 29, 29, 30, 30,
	30, 10, 10, 10, 31, 32, 17, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
	24, 24, 33, 33, 33, 34, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	37, 37, 38, 38, 28, 28, 28, 39, 39, 40,
	40, 41, 41, 36, 36, 36, 36, 36, 36, 36,
	42, 42, 43, 43, 45, 45, 46, 44, 44, 6,
	6,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 8, 10, 10, 5,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 0, 2, 5, 2, 5, 1, 2, 4, 5,
	1, 3, 0, 2, 0, 3, 1, 3, 1, 2,
	2, 0, 1, 2, 2, 2, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 1, 2, 2, 1, 1, 1, 1, 3, 5,
	7, 7, 9, 7, 9, 7, 3, 4, 5, 5,
	3, 5, 0, 2, 1, 4, 3, 1, 3, 1,
	1, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 3, 1, 3, 3, 2, 3, 1,
	3,
}
var yyChk = []int{

	-1000, -1, -2, 4, -3, -4, -5, -7, 5, 6,
	-8, -11, -15, -18, 15, 13, -2, 8, 8, -9,
	17, -12, 15, -13, 16, -19, 11, 12, -25, 30,
	-26, -28, 36, 36, 36, -10, -31, 21, 18, -13,
	-25, 30, -14, 23, -17, -24, -33, 45, 39, -34,
	-35, 36, -36, 52, 59, 64, 65, 67, 68, 37,
	-42, -43, -44, 31, 32, 33, 34, 35, 25, 27,
	-20, -21, -22, -23, -17, 40, -24, 36, -27, 66,
	14, 27, 58, 9, 58, -32, 22, 34, -29, -30,
	-17, -14, 36, -11, 18, 38, 39, 40, 41, 70,
	42, 43, 44, 46, 50, 51, 48, 49, 47, 54,
	45, 58, 27, 55, -33, -33, 52, -17, 60, -24,
	-24, -24, -24, 26, -45, -46, 37, 28, -6, -17,
	29, 14, 58, 58, 36, 36, 34, 36, 36, 36,
	34, 29, 19, 20, 58, -6, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, 54, 36, -24, 33, 45, 57, 56, 53, -39,
	11, 12, -40, -41, -17, 40, -24, 53, -37, -24,
	66, 66, 66, 66, 26, 29, 30, 28, 29, -21,
	36, 40, -25, 69, 28, 52, -29, -25, -16, 24,
	-24, 28, 33, 57, 56, 53, -39, -39, 29, 58,
	-38, 62, 61, 36, 36, 36, 36, -45, -17, -6,
	-28, -6, -17, 53, 53, -39, 40, 63, -24, -24,
	69, 69, 69, 69, -27, 53, 60, -24, -24, -24,
	-24, 10, -37, 63, 63, 60, 63, 60, 63, 7,
	36, -24, -24, 63, 63,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 10, 0, 0,
	44, 31, 42, 20, 0, 19, 2, 0, 0, 51,
	0, 42, 0, 14, 0, 0, 21, 22, 34, 0,
	36, 40, 104, 0, 0, 11, 52, 0, 0, 14,
	32, 0, 0, 0, 43, 56, 81, 0, 0, 84,
	85, 86, 87, 0, 0, 0, 0, 0, 0, 113,
	114, 115, 116, 117, 118, 119, 120, 121, 0, 0,
	18, 23, 24, 26, 27, 29, 56, 0, 37, 0,
	0, 0, 0, 0, 0, 53, 0, 54, 45, 46,
	48, 12, 0, 13, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 82, 83, 0, 0, 0, 0,
	0, 0, 0, 122, 0, 124, 0, 127, 0, 129,
	0, 0, 0, 0, 0, 41, 0, 106, 0, 9,
	55, 0, 49, 50, 0, 16, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 0, 73, 0, 75, 0, 77, 79, 96, 0,
	0, 0, 107, 109, 110, 111, 56, 88, 102, 0,
	0, 0, 0, 0, 123, 0, 0, 128, 0, 25,
	28, 30, 35, 0, 105, 0, 47, 33, 15, 0,
	72, 74, 76, 78, 80, 97, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 125, 126, 130,
	38, 0, 17, 98, 99, 108, 112, 89, 103, 100,
	0, 0, 0, 0, 39, 6, 0, 0, 0, 0,
	0, 0, 101, 90, 91, 0, 93, 0, 95, 7,
	8, 0, 0, 92, 94,
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
		bucket := yyS[yypt-5].s
		name := yyS[yypt-7].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = "view"
		parsingStatement = createIndexStmt
	}
	case 8:
		//line unql.y:97
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
	case 9:
		//line unql.y:112
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 10:
		//line unql.y:124
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 11:
		//line unql.y:130
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 12:
		//line unql.y:137
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 13:
		//line unql.y:141
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 14:
		//line unql.y:147
		{
	}
	case 15:
		//line unql.y:150
		{
		group_by := parsingStack.Pop().(ast.ExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.GroupBy = group_by
		default:
			logDebugGrammar("This statement does not support GROUP BY")
		}
	}
	case 16:
		//line unql.y:162
		{
	}
	case 17:
		//line unql.y:165
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
	case 18:
		//line unql.y:178
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 19:
		//line unql.y:184
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 20:
		//line unql.y:190
		{
	}
	case 21:
		//line unql.y:193
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 22:
		//line unql.y:203
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 23:
		//line unql.y:215
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
	case 24:
		//line unql.y:229
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 25:
		//line unql.y:234
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
	case 26:
		//line unql.y:247
		{
		logDebugGrammar("RESULT STAR")
	}
	case 27:
		//line unql.y:251
		{ 
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 28:
		//line unql.y:258
		{ 
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 29:
		//line unql.y:267
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 30:
		//line unql.y:273
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 31:
		//line unql.y:282
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 32:
		//line unql.y:286
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
		//line unql.y:297
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
		//line unql.y:311
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
	case 35:
		//line unql.y:322
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
	case 36:
		//line unql.y:336
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 37:
		//line unql.y:340
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 38:
		//line unql.y:350
		{
		logDebugGrammar("OVER IN")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s})
	}
	case 39:
		//line unql.y:356
		{
		logDebugGrammar("OVER IN nested")
		rest := parsingStack.Pop().(*ast.From)
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Over:rest})
	}
	case 40:
		//line unql.y:365
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 41:
		//line unql.y:371
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 42:
		//line unql.y:379
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 43:
		//line unql.y:383
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
	case 45:
		//line unql.y:397
		{
		
	}
	case 46:
		//line unql.y:403
		{
		
	}
	case 47:
		//line unql.y:407
		{
		
	}
	case 48:
		//line unql.y:412
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
	case 49:
		//line unql.y:423
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
	case 50:
		//line unql.y:434
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
	case 51:
		//line unql.y:446
		{
		
	}
	case 52:
		//line unql.y:450
		{
		
	}
	case 53:
		//line unql.y:454
		{
		
	}
	case 54:
		//line unql.y:460
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
	case 55:
		//line unql.y:474
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
	case 56:
		//line unql.y:490
		{
		logDebugGrammar("EXPRESSION")
	}
	case 57:
		//line unql.y:495
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 58:
		//line unql.y:503
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 59:
		//line unql.y:511
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 60:
		//line unql.y:519
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 61:
		//line unql.y:527
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 62:
		//line unql.y:535
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 63:
		//line unql.y:543
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 64:
		//line unql.y:551
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 65:
		//line unql.y:559
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 66:
		//line unql.y:567
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 67:
		//line unql.y:575
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line unql.y:583
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line unql.y:591
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line unql.y:599
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line unql.y:607
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line unql.y:615
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line unql.y:623
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line unql.y:631
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line unql.y:639
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line unql.y:646
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line unql.y:653
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line unql.y:660
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line unql.y:667
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line unql.y:674
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line unql.y:681
		{
		
	}
	case 82:
		//line unql.y:687
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line unql.y:694
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression)) 
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line unql.y:701
		{
		
	}
	case 85:
		//line unql.y:706
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 86:
		//line unql.y:712
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 87:
		//line unql.y:718
		{
		logDebugGrammar("LITERAL")
	}
	case 88:
		//line unql.y:722
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 89:
		//line unql.y:726
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
	case 90:
		//line unql.y:743
		{
		logDebugGrammar("ANY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 91:
		//line unql.y:751
		{
		logDebugGrammar("ALL OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 92:
		//line unql.y:759
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 93:
		//line unql.y:768
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 94:
		//line unql.y:776
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 95:
		//line unql.y:785
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 96:
		//line unql.y:793
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line unql.y:799
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line unql.y:806
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 99:
		//line unql.y:814
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line unql.y:823
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 101:
		//line unql.y:831
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
	case 102:
		//line unql.y:845
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 103:
		//line unql.y:849
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 104:
		//line unql.y:855
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression) 
	}
	case 105:
		//line unql.y:861
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n))) 
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line unql.y:868
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s) 
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right) 
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line unql.y:879
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 108:
		//line unql.y:884
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
	case 109:
		//line unql.y:898
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 110:
		//line unql.y:902
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 111:
		//line unql.y:911
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 112:
		//line unql.y:917
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 113:
		//line unql.y:927
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s) 
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line unql.y:933
		{
		logDebugGrammar("NUMBER")
	}
	case 115:
		//line unql.y:937
		{
		logDebugGrammar("OBJECT")
	}
	case 116:
		//line unql.y:941
		{
		logDebugGrammar("ARRAY")
	}
	case 117:
		//line unql.y:945
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true) 
		parsingStack.Push(thisExpression)
	}
	case 118:
		//line unql.y:951
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false) 
		parsingStack.Push(thisExpression)
	}
	case 119:
		//line unql.y:957
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 120:
		//line unql.y:965
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 121:
		//line unql.y:971
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 122:
		//line unql.y:979
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 123:
		//line unql.y:985
		{
		logDebugGrammar("OBJECT")
	}
	case 124:
		//line unql.y:991
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 125:
		//line unql.y:995
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 126:
		//line unql.y:1007
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression) 
	}
	case 127:
		//line unql.y:1017
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 128:
		//line unql.y:1023
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 129:
		//line unql.y:1032
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 130:
		//line unql.y:1039
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
