
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

//line n1ql.y:13
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
const KEY = 57380
const KEYS = 57381
const ORDER = 57382
const BY = 57383
const ASC = 57384
const DESC = 57385
const LIMIT = 57386
const OFFSET = 57387
const GROUP = 57388
const HAVING = 57389
const LBRACE = 57390
const RBRACE = 57391
const LBRACKET = 57392
const RBRACKET = 57393
const COMMA = 57394
const COLON = 57395
const TRUE = 57396
const FALSE = 57397
const NULL = 57398
const INT = 57399
const NUMBER = 57400
const IDENTIFIER = 57401
const STRING = 57402
const PLUS = 57403
const MINUS = 57404
const MULT = 57405
const DIV = 57406
const CONCAT = 57407
const AND = 57408
const OR = 57409
const NOT = 57410
const EQ = 57411
const NE = 57412
const GT = 57413
const GTE = 57414
const LT = 57415
const LTE = 57416
const LPAREN = 57417
const RPAREN = 57418
const LIKE = 57419
const IS = 57420
const VALUED = 57421
const MISSING = 57422
const DOT = 57423
const CASE = 57424
const WHEN = 57425
const THEN = 57426
const ELSE = 57427
const END = 57428
const ANY = 57429
const ALL = 57430
const FIRST = 57431
const ARRAY = 57432
const IN = 57433
const SATISFIES = 57434
const EVERY = 57435
const UNNEST = 57436
const FOR = 57437
const INNER = 57438
const LEFT = 57439
const OUTER = 57440
const MOD = 57441

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
	"KEY",
	"KEYS",
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
	"FIRST",
	"ARRAY",
	"IN",
	"SATISFIES",
	"EVERY",
	"UNNEST",
	"FOR",
	"INNER",
	"LEFT",
	"OUTER",
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

const yyNprod = 182
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1630

var yyAct = []int{

	50, 291, 209, 141, 31, 83, 146, 101, 232, 89,
	127, 76, 34, 155, 315, 85, 145, 131, 200, 312,
	131, 110, 111, 112, 113, 115, 305, 195, 125, 81,
	45, 264, 133, 213, 49, 79, 212, 296, 128, 196,
	294, 126, 290, 173, 164, 151, 96, 190, 103, 358,
	198, 197, 327, 302, 301, 127, 258, 255, 208, 114,
	342, 134, 137, 138, 139, 132, 110, 111, 112, 113,
	115, 116, 117, 125, 118, 123, 121, 122, 119, 120,
	257, 256, 124, 128, 246, 131, 126, 143, 349, 328,
	84, 350, 87, 88, 140, 161, 162, 152, 153, 191,
	292, 157, 191, 304, 114, 143, 226, 326, 325, 323,
	280, 175, 176, 177, 178, 179, 180, 181, 182, 183,
	184, 185, 186, 187, 188, 189, 174, 172, 192, 127,
	170, 293, 207, 171, 210, 46, 169, 127, 205, 148,
	277, 35, 272, 37, 270, 35, 254, 125, 81, 36,
	112, 113, 115, 247, 79, 125, 227, 128, 230, 224,
	126, 245, 231, 149, 85, 128, 238, 237, 126, 242,
	225, 32, 286, 103, 243, 235, 236, 35, 248, 191,
	228, 163, 160, 94, 95, 156, 114, 92, 107, 97,
	82, 251, 43, 299, 288, 92, 285, 159, 241, 298,
	287, 158, 207, 207, 229, 165, 100, 233, 205, 205,
	235, 236, 266, 267, 268, 269, 51, 271, 93, 273,
	259, 260, 92, 253, 274, 239, 93, 240, 275, 221,
	90, 234, 261, 94, 95, 278, 282, 283, 279, 84,
	276, 87, 88, 281, 223, 92, 284, 99, 220, 166,
	147, 329, 324, 93, 91, 300, 289, 297, 222, 219,
	13, 250, 207, 295, 48, 306, 307, 47, 205, 129,
	130, 167, 168, 109, 40, 41, 93, 235, 236, 27,
	303, 94, 95, 318, 21, 25, 17, 320, 361, 319,
	321, 29, 30, 341, 322, 73, 340, 74, 244, 105,
	194, 68, 69, 70, 193, 72, 56, 64, 108, 53,
	331, 332, 106, 333, 334, 52, 335, 336, 104, 22,
	42, 23, 58, 19, 26, 337, 154, 142, 338, 59,
	67, 210, 339, 343, 60, 2, 62, 63, 66, 18,
	61, 65, 353, 354, 204, 352, 203, 356, 263, 44,
	357, 12, 10, 127, 57, 55, 54, 98, 39, 355,
	17, 102, 16, 362, 110, 111, 112, 113, 115, 116,
	117, 125, 118, 123, 121, 122, 119, 120, 86, 33,
	124, 128, 78, 77, 126, 75, 346, 28, 15, 347,
	249, 14, 127, 24, 38, 20, 11, 7, 9, 8,
	6, 5, 114, 110, 111, 112, 113, 115, 116, 117,
	125, 118, 123, 121, 122, 119, 120, 4, 1, 124,
	128, 0, 0, 126, 0, 316, 0, 0, 317, 0,
	0, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 114, 110, 111, 112, 113, 115, 116, 117, 125,
	118, 123, 121, 122, 119, 120, 0, 0, 124, 128,
	0, 0, 126, 0, 313, 0, 0, 314, 0, 0,
	127, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	114, 110, 111, 112, 113, 115, 116, 117, 125, 118,
	123, 121, 122, 119, 120, 0, 0, 124, 128, 0,
	0, 126, 3, 12, 10, 0, 0, 0, 0, 127,
	0, 218, 17, 0, 16, 217, 0, 0, 0, 114,
	110, 111, 112, 113, 115, 116, 117, 125, 118, 123,
	121, 122, 119, 120, 0, 0, 124, 128, 0, 0,
	126, 0, 0, 0, 0, 0, 0, 0, 127, 0,
	216, 0, 0, 0, 215, 0, 0, 0, 114, 110,
	111, 112, 113, 115, 116, 117, 125, 118, 123, 121,
	122, 119, 120, 0, 0, 124, 128, 0, 0, 126,
	0, 0, 0, 0, 360, 0, 0, 127, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 114, 110, 111,
	112, 113, 115, 116, 117, 125, 118, 123, 121, 122,
	119, 120, 0, 0, 124, 128, 0, 0, 126, 0,
	0, 0, 0, 359, 0, 0, 127, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 114, 110, 111, 112,
	113, 115, 116, 117, 125, 118, 123, 121, 122, 119,
	120, 0, 0, 124, 128, 0, 0, 126, 0, 0,
	0, 0, 351, 0, 0, 127, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 114, 110, 111, 112, 113,
	115, 116, 117, 125, 118, 123, 121, 122, 119, 120,
	0, 0, 124, 128, 0, 0, 126, 0, 0, 0,
	0, 348, 0, 0, 127, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 114, 110, 111, 112, 113, 115,
	116, 117, 125, 118, 123, 121, 122, 119, 120, 0,
	0, 124, 128, 0, 0, 126, 0, 0, 0, 0,
	345, 0, 0, 127, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 114, 110, 111, 112, 113, 115, 116,
	117, 125, 118, 123, 121, 122, 119, 120, 0, 0,
	124, 128, 0, 0, 126, 0, 0, 0, 0, 344,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 110, 111, 112, 113, 115, 116, 117,
	125, 118, 123, 121, 122, 119, 120, 0, 0, 124,
	128, 0, 0, 126, 0, 330, 0, 0, 0, 0,
	0, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 114, 110, 111, 112, 113, 115, 116, 117, 125,
	118, 123, 121, 122, 119, 120, 0, 0, 124, 128,
	0, 0, 126, 0, 0, 0, 0, 311, 0, 0,
	127, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	114, 110, 111, 112, 113, 115, 116, 117, 125, 118,
	123, 121, 122, 119, 120, 0, 0, 124, 128, 0,
	0, 126, 0, 0, 0, 0, 0, 0, 0, 127,
	0, 0, 310, 0, 0, 0, 0, 0, 0, 114,
	110, 111, 112, 113, 115, 116, 117, 125, 118, 123,
	121, 122, 119, 120, 0, 0, 124, 128, 0, 0,
	126, 0, 0, 0, 0, 0, 0, 0, 127, 0,
	0, 309, 0, 0, 0, 0, 0, 0, 114, 110,
	111, 112, 113, 115, 116, 117, 125, 118, 123, 121,
	122, 119, 120, 0, 0, 124, 128, 0, 0, 126,
	0, 0, 0, 0, 308, 0, 0, 127, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 114, 110, 111,
	112, 113, 115, 116, 117, 125, 118, 123, 121, 122,
	119, 120, 0, 0, 124, 128, 0, 0, 126, 0,
	0, 265, 0, 0, 0, 0, 127, 252, 0, 0,
	0, 0, 0, 0, 0, 0, 114, 110, 111, 112,
	113, 115, 116, 117, 125, 118, 123, 121, 122, 119,
	120, 0, 0, 124, 128, 0, 0, 126, 0, 0,
	0, 0, 0, 0, 0, 127, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 114, 110, 111, 112, 113,
	115, 116, 117, 125, 118, 123, 121, 122, 119, 120,
	0, 0, 124, 128, 0, 0, 126, 0, 0, 0,
	0, 0, 0, 0, 127, 0, 0, 214, 0, 0,
	0, 0, 0, 0, 114, 110, 111, 112, 113, 115,
	116, 117, 125, 118, 123, 121, 122, 119, 120, 0,
	0, 124, 128, 0, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 127, 0, 0, 211, 0, 0, 0,
	0, 0, 0, 114, 110, 111, 112, 113, 115, 116,
	117, 125, 118, 123, 121, 122, 119, 120, 0, 0,
	124, 128, 0, 0, 126, 0, 0, 0, 0, 0,
	0, 0, 127, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 114, 110, 111, 112, 113, 115, 116, 117,
	125, 118, 123, 121, 122, 119, 120, 0, 0, 124,
	128, 0, 0, 262, 0, 0, 0, 0, 0, 0,
	0, 127, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 114, 110, 111, 112, 113, 115, 116, 117, 125,
	118, 123, 121, 122, 119, 120, 0, 0, 124, 128,
	0, 0, 150, 127, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 110, 111, 112, 113, 115, 116,
	114, 125, 118, 123, 121, 122, 119, 120, 0, 0,
	124, 128, 0, 0, 126, 127, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 110, 111, 112, 113,
	115, 0, 114, 125, 118, 123, 121, 122, 119, 120,
	0, 0, 124, 128, 201, 202, 126, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	73, 0, 74, 0, 114, 0, 68, 69, 70, 71,
	72, 56, 64, 0, 53, 206, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 199, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 62, 63, 0, 73, 61, 74, 0, 0, 0,
	68, 69, 70, 71, 72, 56, 64, 0, 53, 206,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 0, 0, 60, 0, 62, 63, 0, 73, 61,
	74, 0, 0, 0, 68, 69, 70, 71, 72, 56,
	64, 0, 53, 80, 0, 0, 0, 0, 52, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 60, 0, 62,
	63, 0, 73, 61, 74, 144, 0, 0, 68, 69,
	70, 71, 72, 56, 64, 0, 53, 0, 0, 0,
	0, 0, 52, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 60, 0, 62, 63, 0, 73, 61, 74, 0,
	0, 0, 68, 69, 70, 71, 72, 56, 64, 0,
	53, 0, 0, 0, 0, 0, 52, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 0, 0, 60, 0, 62, 63, 0,
	73, 61, 74, 0, 0, 0, 68, 69, 70, 71,
	72, 136, 64, 0, 53, 0, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 62, 63, 0, 73, 61, 74, 0, 0, 0,
	68, 69, 70, 71, 72, 135, 64, 0, 53, 0,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 0, 0, 60, 0, 62, 63, 0, 0, 61,
}
var yyPact = []int{

	478, -1000, -1000, 326, -1000, -1000, -1000, -1000, -1000, -1000,
	294, 244, 292, 249, 242, 259, 118, -1000, -1000, 90,
	230, 234, 291, 133, 242, 82, 218, 1448, 1360, -1000,
	-1000, -1000, 131, -4, 195, -1000, -35, 130, -1000, 202,
	149, 1448, 288, 269, 218, -1000, 129, 252, 232, -1000,
	1083, -1000, 1448, 1448, -1000, -1000, 10, -1000, 1448, -51,
	1536, 1492, 1448, 1448, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 45, 1404, -1000, -1000, 198, -1000, 104,
	-1000, 1161, -36, -1000, 86, 86, 307, -1000, -85, -1000,
	126, 243, 144, 123, 1448, 1448, 122, -37, -1000, 148,
	-1000, -1000, 197, 229, 77, 74, -1000, -38, -1000, 1448,
	1448, 1448, 1448, 1448, 1448, 1448, 1448, 1448, 1448, 1448,
	1448, 1448, 1448, 1448, 1448, -30, 120, 247, -29, -1000,
	-1000, 1272, -18, 1448, 1044, -55, -58, 1005, 459, 420,
	-1000, 210, 196, 176, -1000, 207, 192, 1360, 111, -1000,
	43, 86, 145, 172, 86, -1000, 243, -1000, 174, 141,
	-1000, 1083, 1083, -1000, 110, -1000, 1448, -1000, -1000, 267,
	102, 9, 94, 86, 214, 87, 87, 79, 79, 79,
	79, 1225, 1193, -40, -40, -40, -40, -40, -40, -40,
	1448, -1000, 966, 170, 89, -1000, 1, -1000, -1000, -1000,
	-20, 1316, 1316, 180, -1000, -1000, -1000, 1122, -1000, -54,
	927, 1448, 1448, 1448, 1448, 85, 1448, 83, 1448, -1000,
	27, 1448, -1000, 1448, -1000, -1000, -1000, -1000, 81, -4,
	-1000, -1000, -4, 51, 239, 1448, 1448, 137, -1000, -1000,
	143, 205, -39, -1000, 72, -41, 1448, -44, -1000, -1000,
	1448, -40, -1000, 142, 204, -1000, -1000, -1000, -1000, -22,
	-23, 1316, 40, -60, 1448, 1448, 888, 849, 810, 771,
	-72, 381, -77, 342, -1000, -1000, -1000, -4, -1000, -1000,
	239, -4, 1083, 1083, -4, 239, 50, 201, -1000, -1000,
	49, -1000, -1000, -1000, 48, -24, 30, -1000, 200, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1083, 732, -1000, 1448,
	1448, -1000, 1448, 1448, -1000, 1448, 1448, -1000, -1000, -4,
	-1000, -1000, -4, 239, -1000, -1000, 265, 262, -15, -1000,
	1448, 693, 654, 303, 615, 5, 576, -1000, -1000, -4,
	72, 72, 1448, -1000, -1000, -1000, 1448, -1000, -1000, 1448,
	-1000, -1000, -1000, -1000, -1000, -27, 537, 498, 257, -1000,
	-1000, 72, -1000,
}
var yyPgo = []int{

	0, 418, 335, 417, 401, 400, 399, 398, 1, 16,
	397, 396, 395, 394, 260, 393, 324, 267, 391, 390,
	6, 388, 387, 385, 11, 383, 382, 0, 4, 379,
	5, 12, 9, 8, 378, 7, 361, 358, 357, 216,
	356, 355, 354, 2, 348, 18, 346, 344, 341, 338,
	330, 3, 327,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 25,
	26, 26, 15, 15, 15, 18, 18, 28, 28, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 33, 33,
	34, 34, 34, 29, 29, 29, 29, 29, 29, 32,
	32, 16, 16, 12, 12, 35, 35, 36, 36, 36,
	13, 13, 13, 37, 38, 20, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 39, 39, 39, 40, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 43, 43,
	44, 44, 31, 31, 31, 31, 31, 31, 45, 45,
	46, 46, 47, 47, 42, 42, 42, 42, 42, 42,
	42, 48, 48, 49, 49, 51, 51, 52, 50, 50,
	9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 0, 2, 5, 2, 5, 1, 2, 2,
	4, 3, 3, 3, 5, 4, 3, 5, 4, 4,
	6, 5, 4, 5, 5, 6, 6, 7, 2, 2,
	1, 1, 2, 1, 2, 3, 2, 4, 3, 2,
	2, 0, 2, 0, 3, 1, 3, 1, 2, 2,
	0, 1, 2, 2, 2, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 6, 5, 5, 3, 4, 3,
	4, 3, 4, 1, 2, 2, 1, 1, 1, 1,
	3, 5, 5, 7, 7, 5, 9, 7, 7, 5,
	9, 7, 7, 5, 3, 4, 5, 5, 3, 5,
	0, 2, 1, 4, 6, 5, 5, 3, 1, 3,
	1, 1, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 3, 1, 3, 3, 2, 3,
	1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 32,
	33, -28, 53, -29, -31, 59, 59, 53, -13, -37,
	44, 41, 29, 59, -16, -28, 53, -17, 46, -20,
	-27, -39, 68, 62, -40, -41, 59, -42, 75, 82,
	87, 93, 89, 90, 60, -48, -49, -50, 54, 55,
	56, 57, 58, 48, 50, -23, -24, -25, -26, -20,
	63, -27, 59, -30, 94, 19, -34, 96, 97, -32,
	35, 59, 50, 81, 38, 39, 81, 59, -38, 45,
	57, -35, -36, -20, 30, 30, -17, 59, -14, 41,
	61, 62, 63, 64, 99, 65, 66, 67, 69, 73,
	74, 71, 72, 70, 77, 68, 81, 50, 78, -39,
	-39, 75, -20, 83, -27, 59, 59, -27, -27, -27,
	49, -51, -52, 60, 51, -9, -20, 52, 35, 59,
	81, 81, -31, -31, 19, 98, 59, -32, 57, 53,
	59, -27, -27, 59, 81, 57, 52, 42, 43, 59,
	53, 59, 53, 81, -9, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	77, 59, -27, 57, 53, 56, 68, 80, 79, 76,
	-45, 32, 33, -46, -47, -20, 63, -27, 76, -43,
	-27, 92, 91, 91, 92, 95, 91, 95, 91, 49,
	52, 53, 51, 52, -24, 59, 63, -28, 35, 59,
	-30, -32, -33, 35, 59, 38, 39, -31, -32, 51,
	53, 57, 59, -35, 31, 59, 75, 59, -28, -19,
	47, -27, 51, 53, 57, 56, 80, 79, 76, -45,
	-45, 52, 81, -44, 85, 84, -27, -27, -27, -27,
	59, -27, 59, -27, -51, -20, -9, 59, -30, -30,
	59, -33, -27, -27, -33, 59, 35, 57, 51, 51,
	81, -8, 28, 59, 81, -9, 81, -20, 57, 51,
	51, 76, 76, -45, 63, 86, -27, -27, 86, 92,
	92, 86, 91, 83, 86, 91, 83, 86, -30, -33,
	-30, -30, -33, 59, 51, 59, 59, 76, 59, 51,
	83, -27, -27, -27, -27, -27, -27, -30, -30, -33,
	31, 31, 75, -43, 86, 86, 83, 86, 86, 83,
	86, 86, -30, -8, -8, -9, -27, -27, 76, 86,
	86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 83, 0, 42, 81, 30, 0, 29, 2, 0,
	90, 0, 0, 0, 81, 0, 24, 0, 0, 31,
	32, 45, 0, 47, 73, 152, 0, 0, 21, 91,
	0, 0, 0, 0, 24, 43, 0, 0, 0, 82,
	95, 123, 0, 0, 126, 127, 128, 129, 0, 0,
	0, 0, 0, 0, 164, 165, 166, 167, 168, 169,
	170, 171, 172, 0, 0, 28, 33, 34, 36, 37,
	40, 95, 0, 48, 0, 0, 0, 70, 71, 74,
	0, 76, 0, 0, 0, 0, 0, 0, 92, 0,
	93, 84, 85, 87, 0, 0, 22, 0, 23, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 124,
	125, 0, 0, 0, 0, 128, 128, 0, 0, 0,
	173, 0, 175, 0, 178, 0, 180, 0, 0, 39,
	0, 0, 49, 0, 0, 72, 75, 78, 0, 0,
	157, 79, 80, 18, 0, 94, 0, 88, 89, 8,
	0, 0, 0, 0, 26, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	0, 112, 0, 171, 0, 117, 0, 119, 121, 144,
	0, 0, 0, 158, 160, 161, 162, 95, 130, 150,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 174,
	0, 0, 179, 0, 35, 38, 41, 46, 0, 51,
	52, 53, 56, 0, 0, 0, 0, 0, 77, 153,
	0, 0, 0, 86, 0, 0, 0, 0, 44, 25,
	0, 111, 113, 0, 0, 118, 120, 122, 145, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 176, 177, 181, 50, 55, 59,
	0, 58, 68, 69, 62, 0, 0, 0, 155, 156,
	0, 10, 16, 17, 0, 0, 0, 27, 0, 115,
	116, 146, 147, 159, 163, 131, 151, 148, 132, 0,
	0, 135, 0, 0, 139, 0, 0, 143, 54, 57,
	61, 63, 64, 0, 154, 19, 9, 12, 0, 114,
	0, 0, 0, 0, 0, 0, 0, 60, 65, 66,
	0, 0, 0, 149, 133, 134, 0, 138, 137, 0,
	142, 141, 67, 11, 14, 0, 0, 0, 13, 136,
	140, 0, 15,
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
	92, 93, 94, 95, 96, 97, 98, 99,
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
		//line n1ql.y:55
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line n1ql.y:59
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line n1ql.y:65
		{
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line n1ql.y:69
		{
	}
	case 5:
		//line n1ql.y:72
		{
		logDebugGrammar("STMT - DROP INDEX")
	}
	case 6:
		//line n1ql.y:79
		{
		logDebugGrammar("STMT - CREATE PRIMARY INDEX")
	}
	case 7:
		//line n1ql.y:83
		{
		logDebugGrammar("STMT - CREATE SECONDARY INDEX")
	}
	case 8:
		//line n1ql.y:89
		{
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 9:
		//line n1ql.y:97
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
		//line n1ql.y:107
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
		//line n1ql.y:117
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
		//line n1ql.y:131
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
		//line n1ql.y:143
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
		//line n1ql.y:157
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
		//line n1ql.y:171
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
		//line n1ql.y:190
		{
		parsingStack.Push("view")
	}
	case 17:
		//line n1ql.y:194
		{
		parsingStack.Push(yyS[yypt-0].s)
	}
	case 18:
		//line n1ql.y:200
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 19:
		//line n1ql.y:209
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
		//line n1ql.y:223
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 21:
		//line n1ql.y:229
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND")
	}
	case 22:
		//line n1ql.y:236
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 23:
		//line n1ql.y:240
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 24:
		//line n1ql.y:246
		{
	}
	case 25:
		//line n1ql.y:249
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
		//line n1ql.y:261
		{
	}
	case 27:
		//line n1ql.y:264
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
		//line n1ql.y:277
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 29:
		//line n1ql.y:283
		{
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 30:
		//line n1ql.y:289
		{
	}
	case 31:
		//line n1ql.y:292
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
		//line n1ql.y:302
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
		//line n1ql.y:314
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
		//line n1ql.y:328
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 35:
		//line n1ql.y:333
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
		//line n1ql.y:346
		{
		logDebugGrammar("RESULT STAR")
	}
	case 37:
		//line n1ql.y:350
		{
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 38:
		//line n1ql.y:357
		{
		logDebugGrammar("RESULT EXPR AS ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 39:
		//line n1ql.y:364
		{
		logDebugGrammar("RESULT EXPR ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 40:
		//line n1ql.y:373
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 41:
		//line n1ql.y:379
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 42:
		//line n1ql.y:388
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 43:
		//line n1ql.y:392
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
		//line n1ql.y:403
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
		//line n1ql.y:418
		{
		logDebugGrammar("SELECT FROM - DATASOURCE over here")
		from := parsingStack.Pop().(*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 46:
		//line n1ql.y:429
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
		//line n1ql.y:443
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT UNNEST")
	}
	case 48:
		//line n1ql.y:447
		{
		logDebugGrammar("FROM DATASOURCE WITH UNNEST")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 49:
		//line n1ql.y:458
		{
	    logDebugGrammar("UNNEST")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:""})
	}
	case 50:
		//line n1ql.y:465
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 51:
		//line n1ql.y:472
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 52:
		//line n1ql.y:479
		{
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: "", Over:rest})
	}
	case 53:
		//line n1ql.y:486
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr}) 
	}
	case 54:
		//line n1ql.y:493
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 55:
		//line n1ql.y:500
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 56:
		//line n1ql.y:507
		{
	    logDebugGrammar("JOIN KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr})
	}
	case 57:
		//line n1ql.y:514
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 58:
		//line n1ql.y:521
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 59:
		//line n1ql.y:528
		{
	    logDebugGrammar("JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr, Over: rest})
	}
	case 60:
		//line n1ql.y:536
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 61:
		//line n1ql.y:544
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 62:
		//line n1ql.y:552
		{
	    logDebugGrammar("TYPE JOIN KEY")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type, Keys: key_expr})
	
	}
	case 63:
		//line n1ql.y:561
		{
	    logDebugGrammar("TYPE JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type, Keys: key_expr, Over: rest})
	}
	case 64:
		//line n1ql.y:570
		{
	    logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Type:Type, Keys: key_expr})
	
	}
	case 65:
		//line n1ql.y:579
		{
	    logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Type:Type, Keys: key_expr, Over: rest})
	}
	case 66:
		//line n1ql.y:588
		{
	    logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Type:Type, Keys: key_expr})
	}
	case 67:
		//line n1ql.y:596
		{
	    logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Type:Type, Keys: key_expr, Over: rest})
	}
	case 68:
		//line n1ql.y:607
		{
	        logDebugGrammar("FROM JOIN DATASOURCE with KEY")
	        key := parsingStack.Pop().(ast.Expression)
	        key_expr := ast.NewKeyExpression(key, "KEY")
	        parsingStack.Push(key_expr)
	}
	case 69:
		//line n1ql.y:614
		{
	        logDebugGrammar("FROM DATASOURCE with KEYS")
	        keys := parsingStack.Pop().(ast.Expression)
	        keys_expr := ast.NewKeyExpression(keys, "KEYS")
	        parsingStack.Push(keys_expr)
	
	}
	case 70:
		//line n1ql.y:623
		{
	    logDebugGrammar("INNER")
	    parsingStack.Push("INNER")
	}
	case 71:
		//line n1ql.y:628
		{
	    logDebugGrammar("OUTER")
	    parsingStack.Push("LEFT")
	}
	case 72:
		//line n1ql.y:633
		{
	    logDebugGrammar("LEFT OUTER")
	    parsingStack.Push("LEFT")
	}
	case 73:
		//line n1ql.y:640
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 74:
		//line n1ql.y:646
		{
	    logDebugGrammar("FROM KEY(S) DATASOURCE")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection:proj})
	}
	case 75:
		//line n1ql.y:652
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 76:
		//line n1ql.y:659
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 77:
		//line n1ql.y:666
		{
	        logDebugGrammar("FROM DATASOURCE AS ID KEY(S)")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})
	
	}
	case 78:
		//line n1ql.y:673
		{
	        logDebugGrammar("FROM DATASOURCE ID KEY(s)")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})
	
	}
	case 79:
		//line n1ql.y:682
		{
	        logDebugGrammar("FROM DATASOURCE with KEY")
	        keys := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Keys = ast.NewKeyExpression(keys, "KEY") 
		default:
			logDebugGrammar("This statement does not support KEY")
		}
	}
	case 80:
		//line n1ql.y:693
		{
	        logDebugGrammar("FROM DATASOURCE with KEYS")
	        keys := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Keys = ast.NewKeyExpression(keys, "KEYS")
		default:
			logDebugGrammar("This statement does not support KEYS")
		}
	}
	case 81:
		//line n1ql.y:707
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 82:
		//line n1ql.y:711
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
	case 84:
		//line n1ql.y:725
		{
	
	}
	case 85:
		//line n1ql.y:731
		{
	
	}
	case 86:
		//line n1ql.y:735
		{
	
	}
	case 87:
		//line n1ql.y:740
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
	case 88:
		//line n1ql.y:751
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
	case 89:
		//line n1ql.y:762
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
	case 90:
		//line n1ql.y:774
		{
	
	}
	case 91:
		//line n1ql.y:778
		{
	
	}
	case 92:
		//line n1ql.y:782
		{
	
	}
	case 93:
		//line n1ql.y:788
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
	case 94:
		//line n1ql.y:802
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
	case 95:
		//line n1ql.y:819
		{
		logDebugGrammar("EXPRESSION")
	}
	case 96:
		//line n1ql.y:824
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:832
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line n1ql.y:840
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line n1ql.y:848
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line n1ql.y:856
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line n1ql.y:864
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line n1ql.y:872
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line n1ql.y:880
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line n1ql.y:888
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 105:
		//line n1ql.y:896
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line n1ql.y:904
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line n1ql.y:912
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 108:
		//line n1ql.y:920
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 109:
		//line n1ql.y:928
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line n1ql.y:936
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 111:
		//line n1ql.y:944
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 112:
		//line n1ql.y:952
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 113:
		//line n1ql.y:960
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line n1ql.y:968
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 115:
		//line n1ql.y:975
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 116:
		//line n1ql.y:983
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 117:
		//line n1ql.y:990
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 118:
		//line n1ql.y:997
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 119:
		//line n1ql.y:1004
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 120:
		//line n1ql.y:1011
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 121:
		//line n1ql.y:1018
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 122:
		//line n1ql.y:1025
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 123:
		//line n1ql.y:1032
		{
	
	}
	case 124:
		//line n1ql.y:1038
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 125:
		//line n1ql.y:1045
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 126:
		//line n1ql.y:1052
		{
	
	}
	case 127:
		//line n1ql.y:1057
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 128:
		//line n1ql.y:1063
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 129:
		//line n1ql.y:1069
		{
		logDebugGrammar("LITERAL")
	}
	case 130:
		//line n1ql.y:1073
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 131:
		//line n1ql.y:1077
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
	case 132:
		//line n1ql.y:1094
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 133:
		//line n1ql.y:1102
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 134:
		//line n1ql.y:1110
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 135:
		//line n1ql.y:1118
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 136:
		//line n1ql.y:1126
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 137:
		//line n1ql.y:1135
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 138:
		//line n1ql.y:1144
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 139:
		//line n1ql.y:1152
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 140:
		//line n1ql.y:1160
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 141:
		//line n1ql.y:1169
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 142:
		//line n1ql.y:1178
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 143:
		//line n1ql.y:1186
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 144:
		//line n1ql.y:1194
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 145:
		//line n1ql.y:1200
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 146:
		//line n1ql.y:1207
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 147:
		//line n1ql.y:1215
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1224
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 149:
		//line n1ql.y:1232
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
	case 150:
		//line n1ql.y:1246
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 151:
		//line n1ql.y:1250
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 152:
		//line n1ql.y:1256
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 153:
		//line n1ql.y:1262
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 154:
		//line n1ql.y:1269
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 155:
		//line n1ql.y:1276
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 156:
		//line n1ql.y:1284
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 157:
		//line n1ql.y:1291
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 158:
		//line n1ql.y:1302
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 159:
		//line n1ql.y:1307
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
	case 160:
		//line n1ql.y:1321
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 161:
		//line n1ql.y:1325
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 162:
		//line n1ql.y:1334
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 163:
		//line n1ql.y:1340
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 164:
		//line n1ql.y:1350
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 165:
		//line n1ql.y:1356
		{
		logDebugGrammar("NUMBER")
	}
	case 166:
		//line n1ql.y:1360
		{
		logDebugGrammar("OBJECT")
	}
	case 167:
		//line n1ql.y:1364
		{
		logDebugGrammar("ARRAY")
	}
	case 168:
		//line n1ql.y:1368
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 169:
		//line n1ql.y:1374
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 170:
		//line n1ql.y:1380
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 171:
		//line n1ql.y:1388
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 172:
		//line n1ql.y:1394
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 173:
		//line n1ql.y:1402
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 174:
		//line n1ql.y:1408
		{
		logDebugGrammar("OBJECT")
	}
	case 175:
		//line n1ql.y:1414
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 176:
		//line n1ql.y:1418
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 177:
		//line n1ql.y:1430
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 178:
		//line n1ql.y:1440
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 179:
		//line n1ql.y:1446
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 180:
		//line n1ql.y:1455
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 181:
		//line n1ql.y:1462
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
