
//line n1ql.y:2
package goyacc
import __yyfmt__ "fmt"
//line n1ql.y:2
		import "github.com/couchbaselabs/clog"
import "github.com/couchbaselabs/tuqtng/parser"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
    clog.To(parser.PARSER_CHANNEL, format, v...)
}

//line n1ql.y:12
type yySymType struct {
	yys int
s string
n int
f float64}

const ALTER = 57346
const BETWEEN = 57347
const BUCKET = 57348
const CAST = 57349
const COLLATE = 57350
const DATABASE = 57351
const DELETE = 57352
const EACH = 57353
const EXCEPT = 57354
const EXISTS = 57355
const IF = 57356
const INLINE = 57357
const INSERT = 57358
const INTERSECT = 57359
const INTO = 57360
const JOIN = 57361
const PATH = 57362
const UNION = 57363
const UPDATE = 57364
const POOL = 57365
const EXPLAIN = 57366
const CREATE = 57367
const DROP = 57368
const PRIMARY = 57369
const VIEW = 57370
const INDEX = 57371
const ON = 57372
const USING = 57373
const DISTINCT = 57374
const UNIQUE = 57375
const SELECT = 57376
const AS = 57377
const FROM = 57378
const WHERE = 57379
const ORDER = 57380
const BY = 57381
const ASC = 57382
const DESC = 57383
const LIMIT = 57384
const OFFSET = 57385
const GROUP = 57386
const HAVING = 57387
const LBRACE = 57388
const RBRACE = 57389
const LBRACKET = 57390
const RBRACKET = 57391
const COMMA = 57392
const COLON = 57393
const TRUE = 57394
const FALSE = 57395
const NULL = 57396
const INT = 57397
const NUMBER = 57398
const IDENTIFIER = 57399
const STRING = 57400
const PLUS = 57401
const MINUS = 57402
const MULT = 57403
const DIV = 57404
const CONCAT = 57405
const AND = 57406
const OR = 57407
const NOT = 57408
const EQ = 57409
const NE = 57410
const GT = 57411
const GTE = 57412
const LT = 57413
const LTE = 57414
const LPAREN = 57415
const RPAREN = 57416
const LIKE = 57417
const IS = 57418
const VALUED = 57419
const MISSING = 57420
const DOT = 57421
const CASE = 57422
const WHEN = 57423
const THEN = 57424
const ELSE = 57425
const END = 57426
const ANY = 57427
const ALL = 57428
const OVER = 57429
const FIRST = 57430
const ARRAY = 57431
const IN = 57432
const MOD = 57433

var yyToknames = []string{
	"ALTER",
	"BETWEEN",
	"BUCKET",
	"CAST",
	"COLLATE",
	"DATABASE",
	"DELETE",
	"EACH",
	"EXCEPT",
	"EXISTS",
	"IF",
	"INLINE",
	"INSERT",
	"INTERSECT",
	"INTO",
	"JOIN",
	"PATH",
	"UNION",
	"UPDATE",
	"POOL",
	"EXPLAIN",
	"CREATE",
	"DROP",
	"PRIMARY",
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

const yyNprod = 150
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1406

var yyAct = []int{

	50, 244, 192, 83, 34, 132, 137, 31, 183, 94,
	76, 207, 87, 120, 84, 124, 255, 229, 126, 124,
	136, 124, 124, 220, 103, 104, 105, 106, 108, 81,
	178, 118, 265, 45, 49, 79, 262, 85, 260, 258,
	249, 121, 179, 88, 119, 247, 222, 221, 96, 87,
	87, 84, 243, 181, 180, 120, 107, 158, 149, 86,
	142, 127, 128, 129, 130, 125, 103, 104, 105, 106,
	108, 109, 110, 118, 111, 116, 114, 115, 112, 113,
	88, 88, 117, 121, 89, 175, 119, 297, 289, 143,
	271, 290, 252, 251, 223, 191, 282, 213, 107, 124,
	134, 120, 272, 270, 160, 161, 162, 163, 164, 165,
	166, 167, 168, 169, 170, 171, 172, 173, 174, 118,
	120, 177, 176, 159, 176, 190, 254, 193, 205, 121,
	131, 188, 119, 105, 106, 108, 245, 157, 118, 81,
	155, 134, 139, 156, 269, 79, 154, 46, 121, 203,
	206, 119, 35, 35, 37, 214, 32, 212, 96, 209,
	36, 210, 35, 107, 140, 246, 215, 204, 176, 148,
	147, 145, 144, 100, 90, 82, 218, 43, 150, 146,
	93, 200, 226, 51, 202, 190, 190, 199, 151, 138,
	208, 188, 188, 224, 225, 232, 234, 236, 238, 201,
	198, 217, 48, 92, 40, 239, 13, 240, 152, 153,
	47, 102, 242, 41, 21, 27, 3, 12, 10, 25,
	17, 29, 30, 241, 250, 300, 17, 190, 16, 120,
	256, 257, 281, 188, 248, 253, 122, 123, 280, 211,
	103, 104, 105, 106, 108, 109, 268, 118, 111, 116,
	114, 115, 112, 113, 101, 99, 117, 121, 98, 274,
	119, 275, 97, 276, 277, 42, 278, 279, 12, 10,
	19, 22, 107, 23, 193, 26, 283, 17, 133, 16,
	2, 67, 292, 293, 18, 66, 65, 295, 187, 186,
	296, 228, 57, 120, 55, 54, 91, 39, 95, 33,
	44, 78, 301, 294, 103, 104, 105, 106, 108, 109,
	110, 118, 111, 116, 114, 115, 112, 113, 77, 75,
	117, 121, 28, 15, 119, 216, 286, 120, 14, 287,
	24, 38, 20, 11, 7, 9, 107, 8, 103, 104,
	105, 106, 108, 109, 110, 118, 111, 116, 114, 115,
	112, 113, 6, 5, 117, 121, 4, 1, 119, 0,
	266, 120, 0, 267, 0, 0, 0, 0, 0, 0,
	107, 0, 103, 104, 105, 106, 108, 109, 110, 118,
	111, 116, 114, 115, 112, 113, 0, 0, 117, 121,
	0, 0, 119, 0, 263, 120, 0, 264, 0, 0,
	0, 0, 0, 0, 107, 0, 103, 104, 105, 106,
	108, 109, 110, 118, 111, 116, 114, 115, 112, 113,
	0, 0, 117, 121, 0, 0, 119, 0, 120, 0,
	0, 299, 0, 0, 0, 0, 0, 0, 107, 103,
	104, 105, 106, 108, 109, 110, 118, 111, 116, 114,
	115, 112, 113, 0, 0, 117, 121, 0, 0, 119,
	0, 120, 0, 0, 298, 0, 0, 0, 0, 0,
	0, 107, 103, 104, 105, 106, 108, 109, 110, 118,
	111, 116, 114, 115, 112, 113, 0, 0, 117, 121,
	0, 0, 119, 0, 120, 0, 0, 291, 0, 0,
	0, 0, 0, 0, 107, 103, 104, 105, 106, 108,
	109, 110, 118, 111, 116, 114, 115, 112, 113, 0,
	0, 117, 121, 0, 0, 119, 0, 120, 0, 0,
	288, 0, 0, 0, 0, 0, 0, 107, 103, 104,
	105, 106, 108, 109, 110, 118, 111, 116, 114, 115,
	112, 113, 0, 0, 117, 121, 0, 0, 119, 0,
	120, 0, 0, 285, 0, 0, 0, 0, 0, 0,
	107, 103, 104, 105, 106, 108, 109, 110, 118, 111,
	116, 114, 115, 112, 113, 0, 0, 117, 121, 0,
	0, 119, 0, 120, 0, 0, 284, 0, 0, 0,
	0, 0, 0, 107, 103, 104, 105, 106, 108, 109,
	110, 118, 111, 116, 114, 115, 112, 113, 0, 0,
	117, 121, 0, 0, 119, 0, 273, 120, 0, 0,
	0, 0, 0, 0, 0, 0, 107, 0, 103, 104,
	105, 106, 108, 109, 110, 118, 111, 116, 114, 115,
	112, 113, 0, 0, 117, 121, 0, 0, 119, 0,
	120, 0, 0, 261, 0, 0, 0, 0, 0, 0,
	107, 103, 104, 105, 106, 108, 109, 110, 118, 111,
	116, 114, 115, 112, 113, 0, 0, 117, 121, 0,
	0, 119, 0, 120, 0, 0, 259, 0, 0, 0,
	0, 0, 0, 107, 103, 104, 105, 106, 108, 109,
	110, 118, 111, 116, 114, 115, 112, 113, 0, 0,
	117, 121, 0, 0, 119, 0, 0, 230, 120, 219,
	0, 0, 0, 0, 0, 0, 107, 0, 0, 103,
	104, 105, 106, 108, 109, 110, 118, 111, 116, 114,
	115, 112, 113, 0, 0, 117, 121, 0, 0, 119,
	0, 120, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 107, 103, 104, 105, 106, 108, 109, 110, 118,
	111, 116, 114, 115, 112, 113, 0, 0, 117, 121,
	0, 0, 119, 0, 120, 0, 0, 0, 0, 0,
	197, 0, 0, 0, 107, 103, 104, 105, 106, 108,
	109, 110, 118, 111, 116, 114, 115, 112, 113, 0,
	0, 117, 121, 0, 0, 119, 0, 120, 0, 0,
	0, 0, 0, 196, 0, 0, 0, 107, 103, 104,
	105, 106, 108, 109, 110, 118, 111, 116, 114, 115,
	112, 113, 0, 0, 117, 121, 0, 0, 119, 0,
	120, 0, 0, 0, 0, 0, 195, 0, 0, 0,
	107, 103, 104, 105, 106, 108, 109, 110, 118, 111,
	116, 114, 115, 112, 113, 0, 0, 117, 121, 0,
	0, 119, 0, 120, 0, 0, 0, 0, 0, 194,
	0, 0, 0, 107, 103, 104, 105, 106, 108, 109,
	110, 118, 111, 116, 114, 115, 112, 113, 0, 0,
	117, 121, 0, 0, 119, 0, 120, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 107, 103, 104, 105,
	106, 108, 109, 110, 118, 111, 116, 114, 115, 112,
	113, 184, 185, 117, 121, 0, 0, 227, 0, 0,
	0, 0, 0, 0, 0, 73, 0, 74, 0, 107,
	0, 68, 69, 70, 71, 72, 56, 64, 0, 53,
	189, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 182, 0, 0, 0, 0, 120, 59,
	0, 0, 0, 0, 60, 61, 0, 62, 63, 103,
	104, 105, 106, 108, 109, 110, 118, 111, 116, 114,
	115, 112, 113, 0, 120, 117, 121, 0, 0, 141,
	0, 0, 0, 0, 0, 103, 104, 105, 106, 108,
	0, 107, 118, 111, 116, 114, 115, 112, 113, 0,
	0, 117, 121, 0, 73, 119, 74, 0, 0, 0,
	68, 69, 70, 71, 72, 56, 64, 107, 53, 189,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 0, 0, 60, 61, 0, 62, 63, 73, 0,
	74, 0, 0, 0, 68, 69, 70, 71, 72, 56,
	64, 0, 53, 80, 0, 0, 0, 0, 52, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 60, 61, 0,
	62, 63, 73, 0, 74, 135, 0, 0, 68, 69,
	70, 71, 72, 56, 64, 0, 53, 0, 0, 0,
	0, 0, 52, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 60, 61, 0, 62, 63, 73, 0, 74, 0,
	0, 0, 68, 69, 70, 71, 72, 56, 64, 0,
	53, 0, 0, 0, 0, 0, 52, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 0, 0, 60, 61, 0, 62, 63,
	73, 0, 74, 0, 0, 0, 68, 69, 70, 71,
	72, 237, 64, 0, 53, 0, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	61, 0, 62, 63, 73, 0, 74, 0, 0, 0,
	68, 69, 70, 71, 72, 235, 64, 0, 53, 0,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 0, 0, 60, 61, 0, 62, 63, 73, 0,
	74, 0, 0, 0, 68, 69, 70, 71, 72, 233,
	64, 0, 53, 0, 0, 0, 0, 0, 52, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 60, 61, 0,
	62, 63, 73, 0, 74, 0, 0, 0, 68, 69,
	70, 71, 72, 231, 64, 0, 53, 0, 0, 0,
	0, 0, 52, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 60, 61, 0, 62, 63,
}
var yyPact = []int{

	192, -1000, -1000, 243, -1000, -1000, -1000, -1000, -1000, -1000,
	241, 176, 244, 183, 178, 189, 105, -1000, -1000, 103,
	162, 174, 236, 120, 178, 96, 158, 1140, 1052, -1000,
	-1000, -1000, 118, -73, 2, -1000, 5, 117, -1000, 160,
	125, 1140, 232, 228, 158, -1000, 116, 186, 172, -1000,
	845, -1000, 1140, 1140, -1000, -1000, 26, -1000, 1140, -63,
	1140, 1140, 1140, 1140, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 83, 1096, -1000, -1000, 139, -1000, 107,
	-1000, 950, -19, -1000, 115, 114, -1000, 124, 113, 112,
	-21, -1000, 123, -1000, -1000, 138, 168, 89, 86, -1000,
	-22, -1000, 1140, 1140, 1140, 1140, 1140, 1140, 1140, 1140,
	1140, 1140, 1140, 1140, 1140, 1140, 1140, 1140, 10, 111,
	1140, -24, -1000, -1000, 919, 21, 1140, 812, 779, 746,
	713, -1000, 153, 137, 130, -1000, 150, 134, 1052, 110,
	-1000, 67, 95, 1, -79, -1000, 141, -1000, -1000, 102,
	-1000, 1140, -1000, -1000, 208, 100, 24, 98, 95, 156,
	72, 72, 53, 53, 53, 53, 976, 181, -35, -35,
	-35, -35, -35, -35, -35, 1140, -1000, 680, -1000, -31,
	-1000, -1000, -1000, 20, 1008, 1008, 132, -1000, -1000, -1000,
	878, -1000, -66, 645, 1316, 1272, 1228, 1184, -1000, 42,
	1140, -1000, 1140, -1000, -1000, -1000, -1000, 95, -1000, -27,
	-1000, 108, -34, 1140, -39, -1000, -1000, 1140, -35, -1000,
	-1000, -1000, -1000, -1000, 19, 18, 1008, 65, -68, 1140,
	1140, -51, 612, -52, 579, -54, 313, -58, 279, -1000,
	-1000, -1000, -36, 87, -1000, -1000, -1000, 46, 16, 45,
	-1000, -1000, -1000, -1000, -1000, -1000, 845, 545, 1140, -1000,
	1140, -1000, 1140, 1140, -1000, 1140, 1140, -1000, -1000, -1000,
	207, 201, 23, 1140, 512, 479, 245, 446, 7, 413,
	108, 108, 1140, -1000, -1000, -1000, 1140, -1000, -1000, 1140,
	-1000, -1000, -1000, -1000, 13, 380, 347, 194, -1000, -1000,
	108, -1000,
}
var yyPgo = []int{

	0, 357, 280, 356, 353, 352, 337, 335, 1, 20,
	334, 333, 332, 331, 206, 330, 275, 210, 328, 325,
	6, 323, 322, 319, 10, 318, 301, 0, 7, 299,
	3, 4, 9, 298, 297, 296, 183, 295, 294, 292,
	2, 291, 8, 289, 288, 286, 285, 281, 5, 278,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 25,
	26, 26, 15, 15, 15, 18, 18, 28, 28, 30,
	30, 30, 29, 29, 29, 16, 16, 12, 12, 32,
	32, 33, 33, 33, 13, 13, 13, 34, 35, 20,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 36, 36, 36, 37, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 40,
	40, 41, 41, 31, 31, 31, 42, 42, 43, 43,
	44, 44, 39, 39, 39, 39, 39, 39, 39, 45,
	45, 46, 46, 48, 48, 49, 47, 47, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 0, 2, 5, 2, 5, 1, 2, 2,
	4, 5, 1, 3, 2, 0, 2, 0, 3, 1,
	3, 1, 2, 2, 0, 1, 2, 2, 2, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 1, 2, 2, 1, 1, 1,
	1, 3, 5, 7, 5, 7, 5, 9, 7, 7,
	5, 9, 7, 7, 5, 3, 4, 5, 5, 3,
	5, 0, 2, 1, 4, 3, 1, 3, 1, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 3, 1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 38, 27, 29, -15, 36, -16, 37, -22, 32,
	33, -28, 51, -29, -31, 57, 57, 51, -13, -34,
	42, 39, 29, 57, -16, -28, 51, -17, 44, -20,
	-27, -36, 66, 60, -37, -38, 57, -39, 73, 80,
	85, 86, 88, 89, 58, -45, -46, -47, 52, 53,
	54, 55, 56, 46, 48, -23, -24, -25, -26, -20,
	61, -27, 57, -30, 87, 35, 57, 48, 79, 79,
	57, -35, 43, 55, -32, -33, -20, 30, 30, -17,
	57, -14, 39, 59, 60, 61, 62, 91, 63, 64,
	65, 67, 71, 72, 69, 70, 68, 75, 66, 79,
	48, 76, -36, -36, 73, -20, 81, -27, -27, -27,
	-27, 47, -48, -49, 58, 49, -9, -20, 50, 35,
	57, 79, 79, -31, 57, 57, 55, 57, 57, 79,
	55, 50, 40, 41, 57, 51, 57, 51, 79, -9,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, 75, 57, -27, 54, 66,
	78, 77, 74, -42, 32, 33, -43, -44, -20, 61,
	-27, 74, -40, -27, 87, 87, 87, 87, 47, 50,
	51, 49, 50, -24, 57, 61, -28, 90, 49, 57,
	-32, 31, 57, 73, 57, -28, -19, 45, -27, 49,
	54, 78, 77, 74, -42, -42, 50, 79, -41, 83,
	82, 57, -27, 57, -27, 57, -27, 57, -27, -48,
	-20, -9, -31, 79, -8, 28, 57, 79, -9, 79,
	-20, 74, 74, -42, 61, 84, -27, -27, 90, 84,
	90, 84, 90, 81, 84, 90, 81, 84, -30, 57,
	57, 74, 57, 81, -27, -27, -27, -27, -27, -27,
	31, 31, 73, -40, 84, 84, 81, 84, 84, 81,
	84, 84, -8, -8, -9, -27, -27, 74, 84, 84,
	31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 57, 0, 42, 55, 30, 0, 29, 2, 0,
	64, 0, 0, 0, 55, 0, 24, 0, 0, 31,
	32, 45, 0, 47, 52, 123, 0, 0, 21, 65,
	0, 0, 0, 0, 24, 43, 0, 0, 0, 56,
	69, 94, 0, 0, 97, 98, 99, 100, 0, 0,
	0, 0, 0, 0, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 0, 0, 28, 33, 34, 36, 37,
	40, 69, 0, 48, 0, 0, 54, 0, 0, 0,
	0, 66, 0, 67, 58, 59, 61, 0, 0, 22,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 95, 96, 0, 0, 0, 0, 0, 0,
	0, 141, 0, 143, 0, 146, 0, 148, 0, 0,
	39, 0, 0, 49, 123, 53, 0, 125, 18, 0,
	68, 0, 62, 63, 8, 0, 0, 0, 0, 26,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 82, 83, 84, 0, 86, 0, 88, 0,
	90, 92, 115, 0, 0, 0, 126, 128, 129, 130,
	69, 101, 121, 0, 0, 0, 0, 0, 142, 0,
	0, 147, 0, 35, 38, 41, 46, 0, 124, 0,
	60, 0, 0, 0, 0, 44, 25, 0, 85, 87,
	89, 91, 93, 116, 0, 0, 0, 0, 0, 0,
	0, 99, 0, 99, 0, 99, 0, 99, 0, 144,
	145, 149, 50, 0, 10, 16, 17, 0, 0, 0,
	27, 117, 118, 127, 131, 102, 122, 119, 0, 104,
	0, 106, 0, 0, 110, 0, 0, 114, 51, 19,
	9, 12, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 120, 103, 105, 0, 109, 108, 0,
	113, 112, 11, 14, 0, 0, 0, 13, 107, 111,
	0, 15,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
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
		//line n1ql.y:53
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line n1ql.y:57
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line n1ql.y:63
		{
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line n1ql.y:67
		{
	}
	case 5:
		//line n1ql.y:70
		{
		logDebugGrammar("STMT - DROP INDEX")
	}
	case 6:
		//line n1ql.y:77
		{
		logDebugGrammar("STMT - CREATE PRIMARY INDEX")
	}
	case 7:
		//line n1ql.y:81
		{
		logDebugGrammar("STMT - CREATE SECONDARY INDEX")
	}
	case 8:
		//line n1ql.y:87
		{
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 9:
		//line n1ql.y:95
		{
		pool := yyS[yypt-2].s
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 10:
		//line n1ql.y:105
		{
		method := parsingStack.Pop().(string)
		bucket := yyS[yypt-2].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Method = method
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 11:
		//line n1ql.y:115
		{
		method := parsingStack.Pop().(string)
		bucket := yyS[yypt-2].s
		pool := yyS[yypt-4].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Method = method
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 12:
		//line n1ql.y:129
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-3].s
		name := yyS[yypt-5].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 13:
		//line n1ql.y:141
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
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 14:
		//line n1ql.y:155
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
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 15:
		//line n1ql.y:169
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
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 16:
		//line n1ql.y:187
		{
		parsingStack.Push("view")
	}
	case 17:
		//line n1ql.y:191
		{
		parsingStack.Push(yyS[yypt-0].s)
	}
	case 18:
		//line n1ql.y:197
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 19:
		//line n1ql.y:206
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
	case 20:
		//line n1ql.y:220
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 21:
		//line n1ql.y:226
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND")
	}
	case 22:
		//line n1ql.y:233
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 23:
		//line n1ql.y:237
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 24:
		//line n1ql.y:243
		{
	}
	case 25:
		//line n1ql.y:246
		{
		group_by := parsingStack.Pop().(ast.ExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.GroupBy = group_by
		default:
			logDebugGrammar("This statement does not support GROUP BY")
		}
	}
	case 26:
		//line n1ql.y:258
		{
	}
	case 27:
		//line n1ql.y:261
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
	case 28:
		//line n1ql.y:274
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 29:
		//line n1ql.y:280
		{
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 30:
		//line n1ql.y:286
		{
	}
	case 31:
		//line n1ql.y:289
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 32:
		//line n1ql.y:299
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 33:
		//line n1ql.y:311
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
	case 34:
		//line n1ql.y:325
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 35:
		//line n1ql.y:330
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
	case 36:
		//line n1ql.y:343
		{
		logDebugGrammar("RESULT STAR")
	}
	case 37:
		//line n1ql.y:347
		{
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 38:
		//line n1ql.y:354
		{
		logDebugGrammar("RESULT EXPR AS ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 39:
		//line n1ql.y:361
		{
		logDebugGrammar("RESULT EXPR ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 40:
		//line n1ql.y:370
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 41:
		//line n1ql.y:376
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 42:
		//line n1ql.y:385
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 43:
		//line n1ql.y:389
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
	case 44:
		//line n1ql.y:400
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
	case 45:
		//line n1ql.y:414
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
	case 46:
		//line n1ql.y:425
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
	case 47:
		//line n1ql.y:439
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 48:
		//line n1ql.y:443
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 49:
		//line n1ql.y:453
		{
		logDebugGrammar("OVER IN")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: ""})
	}
	case 50:
		//line n1ql.y:459
		{
		logDebugGrammar("OVER IN")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s})
	}
	case 51:
		//line n1ql.y:465
		{
		logDebugGrammar("OVER IN nested")
		rest := parsingStack.Pop().(*ast.From)
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Over:rest})
	}
	case 52:
		//line n1ql.y:474
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 53:
		//line n1ql.y:480
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 54:
		//line n1ql.y:487
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 55:
		//line n1ql.y:496
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 56:
		//line n1ql.y:500
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
	case 58:
		//line n1ql.y:514
		{
	
	}
	case 59:
		//line n1ql.y:520
		{
	
	}
	case 60:
		//line n1ql.y:524
		{
	
	}
	case 61:
		//line n1ql.y:529
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
	case 62:
		//line n1ql.y:540
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
	case 63:
		//line n1ql.y:551
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
	case 64:
		//line n1ql.y:563
		{
	
	}
	case 65:
		//line n1ql.y:567
		{
	
	}
	case 66:
		//line n1ql.y:571
		{
	
	}
	case 67:
		//line n1ql.y:577
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
	case 68:
		//line n1ql.y:591
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
	case 69:
		//line n1ql.y:607
		{
		logDebugGrammar("EXPRESSION")
	}
	case 70:
		//line n1ql.y:612
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line n1ql.y:620
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line n1ql.y:628
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line n1ql.y:636
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line n1ql.y:644
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line n1ql.y:652
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line n1ql.y:660
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line n1ql.y:668
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line n1ql.y:676
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line n1ql.y:684
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line n1ql.y:692
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line n1ql.y:700
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line n1ql.y:708
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line n1ql.y:716
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line n1ql.y:724
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line n1ql.y:732
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line n1ql.y:740
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 87:
		//line n1ql.y:748
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line n1ql.y:756
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line n1ql.y:763
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line n1ql.y:770
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line n1ql.y:777
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line n1ql.y:784
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line n1ql.y:791
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line n1ql.y:798
		{
	
	}
	case 95:
		//line n1ql.y:804
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line n1ql.y:811
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:818
		{
	
	}
	case 98:
		//line n1ql.y:823
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 99:
		//line n1ql.y:829
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line n1ql.y:835
		{
		logDebugGrammar("LITERAL")
	}
	case 101:
		//line n1ql.y:839
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 102:
		//line n1ql.y:843
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
	case 103:
		//line n1ql.y:860
		{
		logDebugGrammar("ANY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 104:
		//line n1ql.y:868
		{
		logDebugGrammar("ANY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
		parsingStack.Push(collectionAny)
	}
	case 105:
		//line n1ql.y:876
		{
		logDebugGrammar("ALL OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 106:
		//line n1ql.y:884
		{
		logDebugGrammar("ALL OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
		parsingStack.Push(collectionAny)
	}
	case 107:
		//line n1ql.y:892
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 108:
		//line n1ql.y:901
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 109:
		//line n1ql.y:910
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 110:
		//line n1ql.y:918
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 111:
		//line n1ql.y:926
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 112:
		//line n1ql.y:935
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 113:
		//line n1ql.y:944
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 114:
		//line n1ql.y:952
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 115:
		//line n1ql.y:960
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 116:
		//line n1ql.y:966
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 117:
		//line n1ql.y:973
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 118:
		//line n1ql.y:981
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 119:
		//line n1ql.y:990
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 120:
		//line n1ql.y:998
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
	case 121:
		//line n1ql.y:1012
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 122:
		//line n1ql.y:1016
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 123:
		//line n1ql.y:1022
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line n1ql.y:1028
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 125:
		//line n1ql.y:1035
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 126:
		//line n1ql.y:1046
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 127:
		//line n1ql.y:1051
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
	case 128:
		//line n1ql.y:1065
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 129:
		//line n1ql.y:1069
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 130:
		//line n1ql.y:1078
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 131:
		//line n1ql.y:1084
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 132:
		//line n1ql.y:1094
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 133:
		//line n1ql.y:1100
		{
		logDebugGrammar("NUMBER")
	}
	case 134:
		//line n1ql.y:1104
		{
		logDebugGrammar("OBJECT")
	}
	case 135:
		//line n1ql.y:1108
		{
		logDebugGrammar("ARRAY")
	}
	case 136:
		//line n1ql.y:1112
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 137:
		//line n1ql.y:1118
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 138:
		//line n1ql.y:1124
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 139:
		//line n1ql.y:1132
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 140:
		//line n1ql.y:1138
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 141:
		//line n1ql.y:1146
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 142:
		//line n1ql.y:1152
		{
		logDebugGrammar("OBJECT")
	}
	case 143:
		//line n1ql.y:1158
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 144:
		//line n1ql.y:1162
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 145:
		//line n1ql.y:1174
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 146:
		//line n1ql.y:1184
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 147:
		//line n1ql.y:1190
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1199
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 149:
		//line n1ql.y:1206
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
