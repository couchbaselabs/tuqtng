
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

const yyNprod = 177
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1667

var yyAct = []int{

	50, 284, 204, 276, 84, 226, 145, 31, 195, 140,
	100, 83, 34, 76, 86, 86, 144, 88, 89, 90,
	130, 130, 308, 305, 298, 257, 126, 190, 351, 81,
	132, 289, 287, 45, 49, 79, 208, 207, 277, 191,
	283, 278, 279, 168, 124, 159, 150, 95, 102, 185,
	193, 192, 335, 93, 127, 126, 321, 125, 295, 294,
	251, 133, 136, 137, 138, 131, 109, 110, 111, 112,
	114, 115, 116, 124, 117, 122, 120, 121, 118, 119,
	203, 239, 123, 127, 94, 248, 125, 130, 342, 85,
	85, 343, 88, 89, 90, 139, 186, 186, 151, 152,
	297, 221, 142, 285, 113, 246, 142, 322, 250, 249,
	170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
	180, 181, 182, 183, 184, 169, 167, 187, 126, 165,
	320, 202, 166, 205, 286, 164, 126, 200, 46, 109,
	110, 111, 112, 114, 35, 147, 124, 81, 319, 111,
	112, 114, 315, 79, 124, 91, 127, 272, 222, 125,
	219, 270, 86, 225, 127, 37, 231, 125, 102, 148,
	93, 36, 236, 265, 32, 263, 241, 113, 223, 92,
	35, 35, 240, 238, 235, 113, 244, 220, 186, 158,
	157, 154, 106, 93, 96, 82, 227, 202, 202, 229,
	230, 94, 224, 200, 200, 252, 253, 259, 260, 261,
	262, 93, 264, 43, 266, 292, 281, 247, 51, 234,
	228, 291, 280, 268, 94, 267, 160, 99, 156, 216,
	274, 275, 155, 254, 273, 269, 218, 85, 271, 232,
	323, 233, 94, 215, 161, 146, 318, 293, 282, 217,
	290, 214, 13, 243, 48, 202, 288, 98, 299, 300,
	47, 200, 40, 296, 162, 163, 108, 41, 278, 279,
	27, 128, 129, 229, 230, 25, 21, 17, 312, 316,
	317, 314, 311, 12, 10, 313, 3, 12, 10, 29,
	30, 354, 17, 104, 16, 334, 17, 333, 16, 237,
	107, 103, 42, 325, 326, 105, 327, 328, 19, 329,
	330, 22, 153, 23, 26, 141, 2, 67, 66, 332,
	18, 65, 199, 198, 331, 205, 256, 336, 57, 55,
	54, 97, 39, 101, 87, 346, 347, 345, 33, 44,
	349, 78, 77, 350, 75, 28, 126, 15, 242, 14,
	24, 38, 348, 20, 11, 7, 355, 109, 110, 111,
	112, 114, 115, 116, 124, 117, 122, 120, 121, 118,
	119, 9, 8, 123, 127, 6, 5, 125, 4, 339,
	1, 0, 340, 0, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 109, 110, 111, 112,
	114, 115, 116, 124, 117, 122, 120, 121, 118, 119,
	0, 0, 123, 127, 0, 0, 125, 0, 309, 0,
	0, 310, 0, 0, 126, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 109, 110, 111, 112, 114,
	115, 116, 124, 117, 122, 120, 121, 118, 119, 0,
	0, 123, 127, 0, 0, 125, 0, 306, 0, 0,
	307, 0, 0, 126, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 113, 109, 110, 111, 112, 114, 115,
	116, 124, 117, 122, 120, 121, 118, 119, 0, 0,
	123, 127, 0, 0, 125, 0, 0, 0, 0, 0,
	0, 0, 126, 0, 213, 0, 0, 0, 212, 0,
	0, 0, 113, 109, 110, 111, 112, 114, 115, 116,
	124, 117, 122, 120, 121, 118, 119, 0, 0, 123,
	127, 0, 0, 125, 0, 0, 0, 0, 0, 0,
	0, 126, 0, 211, 0, 0, 0, 210, 0, 0,
	0, 113, 109, 110, 111, 112, 114, 115, 116, 124,
	117, 122, 120, 121, 118, 119, 0, 0, 123, 127,
	0, 0, 125, 0, 0, 0, 0, 353, 0, 0,
	126, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 109, 110, 111, 112, 114, 115, 116, 124, 117,
	122, 120, 121, 118, 119, 0, 0, 123, 127, 0,
	0, 125, 0, 0, 0, 0, 352, 0, 0, 126,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 113,
	109, 110, 111, 112, 114, 115, 116, 124, 117, 122,
	120, 121, 118, 119, 0, 0, 123, 127, 0, 0,
	125, 0, 0, 0, 0, 344, 0, 0, 126, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 113, 109,
	110, 111, 112, 114, 115, 116, 124, 117, 122, 120,
	121, 118, 119, 0, 0, 123, 127, 0, 0, 125,
	0, 0, 0, 0, 341, 0, 0, 126, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 113, 109, 110,
	111, 112, 114, 115, 116, 124, 117, 122, 120, 121,
	118, 119, 0, 0, 123, 127, 0, 0, 125, 0,
	0, 0, 0, 338, 0, 0, 126, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 113, 109, 110, 111,
	112, 114, 115, 116, 124, 117, 122, 120, 121, 118,
	119, 0, 0, 123, 127, 0, 0, 125, 0, 0,
	0, 0, 337, 0, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 109, 110, 111, 112,
	114, 115, 116, 124, 117, 122, 120, 121, 118, 119,
	0, 0, 123, 127, 0, 0, 125, 0, 324, 0,
	0, 0, 0, 0, 126, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 109, 110, 111, 112, 114,
	115, 116, 124, 117, 122, 120, 121, 118, 119, 0,
	0, 123, 127, 0, 0, 125, 0, 0, 0, 0,
	304, 0, 0, 126, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 113, 109, 110, 111, 112, 114, 115,
	116, 124, 117, 122, 120, 121, 118, 119, 0, 0,
	123, 127, 0, 0, 125, 0, 0, 0, 0, 0,
	0, 0, 126, 0, 0, 303, 0, 0, 0, 0,
	0, 0, 113, 109, 110, 111, 112, 114, 115, 116,
	124, 117, 122, 120, 121, 118, 119, 0, 0, 123,
	127, 0, 0, 125, 0, 0, 0, 0, 0, 0,
	0, 126, 0, 0, 302, 0, 0, 0, 0, 0,
	0, 113, 109, 110, 111, 112, 114, 115, 116, 124,
	117, 122, 120, 121, 118, 119, 0, 0, 123, 127,
	0, 0, 125, 0, 0, 0, 0, 301, 0, 0,
	126, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 109, 110, 111, 112, 114, 115, 116, 124, 117,
	122, 120, 121, 118, 119, 0, 0, 123, 127, 0,
	0, 125, 0, 0, 258, 0, 0, 0, 0, 126,
	245, 0, 0, 0, 0, 0, 0, 0, 0, 113,
	109, 110, 111, 112, 114, 115, 116, 124, 117, 122,
	120, 121, 118, 119, 0, 0, 123, 127, 0, 0,
	125, 0, 0, 0, 0, 0, 0, 0, 126, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 113, 109,
	110, 111, 112, 114, 115, 116, 124, 117, 122, 120,
	121, 118, 119, 0, 0, 123, 127, 0, 0, 125,
	0, 0, 0, 0, 0, 0, 0, 126, 0, 0,
	209, 0, 0, 0, 0, 0, 0, 113, 109, 110,
	111, 112, 114, 115, 116, 124, 117, 122, 120, 121,
	118, 119, 0, 0, 123, 127, 0, 0, 125, 0,
	0, 0, 0, 0, 0, 0, 126, 0, 0, 206,
	0, 0, 0, 0, 0, 0, 113, 109, 110, 111,
	112, 114, 115, 116, 124, 117, 122, 120, 121, 118,
	119, 0, 0, 123, 127, 0, 0, 125, 0, 0,
	0, 0, 0, 0, 0, 126, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 109, 110, 111, 112,
	114, 115, 116, 124, 117, 122, 120, 121, 118, 119,
	0, 0, 123, 127, 0, 0, 255, 0, 0, 0,
	0, 0, 0, 0, 126, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 109, 110, 111, 112, 114,
	115, 116, 124, 117, 122, 120, 121, 118, 119, 0,
	0, 123, 127, 0, 0, 149, 126, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 109, 110, 111,
	112, 114, 115, 113, 124, 117, 122, 120, 121, 118,
	119, 0, 0, 123, 127, 0, 0, 125, 126, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 109,
	110, 111, 112, 114, 0, 113, 124, 117, 122, 120,
	121, 118, 119, 0, 0, 123, 127, 196, 197, 125,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 73, 0, 74, 0, 113, 0, 68,
	69, 70, 71, 72, 56, 64, 0, 53, 201, 0,
	0, 0, 0, 52, 0, 0, 0, 0, 0, 0,
	58, 194, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 60, 0, 62, 63, 0, 73, 61, 74,
	0, 0, 0, 68, 69, 70, 71, 72, 56, 64,
	0, 53, 201, 0, 0, 0, 0, 52, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 60, 0, 62, 63,
	0, 73, 61, 74, 0, 0, 0, 68, 69, 70,
	71, 72, 56, 64, 0, 53, 80, 0, 0, 0,
	0, 52, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	60, 0, 62, 63, 0, 73, 61, 74, 0, 0,
	189, 68, 69, 70, 188, 72, 56, 64, 0, 53,
	0, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 62, 63, 0, 73,
	61, 74, 143, 0, 0, 68, 69, 70, 71, 72,
	56, 64, 0, 53, 0, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 0, 0, 60, 0,
	62, 63, 0, 73, 61, 74, 0, 0, 0, 68,
	69, 70, 71, 72, 56, 64, 0, 53, 0, 0,
	0, 0, 0, 52, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 60, 0, 62, 63, 0, 73, 61, 74,
	0, 0, 0, 68, 69, 70, 71, 72, 135, 64,
	0, 53, 0, 0, 0, 0, 0, 52, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 60, 0, 62, 63,
	0, 73, 61, 74, 0, 0, 0, 68, 69, 70,
	71, 72, 134, 64, 0, 53, 0, 0, 0, 0,
	0, 52, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	60, 0, 62, 63, 0, 0, 61,
}
var yyPact = []int{

	262, -1000, -1000, 258, -1000, -1000, -1000, -1000, -1000, -1000,
	279, 236, 284, 239, 233, 257, 121, -1000, -1000, 112,
	218, 226, 273, 154, 233, 85, 208, 1485, 1353, -1000,
	-1000, -1000, 136, -4, 120, -1000, -34, 135, -1000, 212,
	170, 1485, 271, 263, 208, -1000, 133, 243, 225, -1000,
	1076, -1000, 1485, 1485, -1000, -1000, 12, -1000, 1485, -53,
	1573, 1529, 1485, 1485, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 46, 1441, -1000, -1000, 193, -1000, 110,
	-1000, 1154, -35, -1000, -1000, 122, 122, 293, -1000, -1000,
	-1000, 132, -1000, 175, 131, 130, -36, -1000, 169, -1000,
	-1000, 192, 222, 76, 73, -1000, -38, -1000, 1485, 1485,
	1485, 1485, 1485, 1485, 1485, 1485, 1485, 1485, 1485, 1485,
	1485, 1485, 1485, 1485, -28, 129, 1397, -29, -1000, -1000,
	1265, 4, 1485, 1037, -54, -55, 998, 452, 413, -1000,
	202, 191, 176, -1000, 198, 184, 1353, 128, -1000, 38,
	122, 143, 161, 122, -1000, 188, 162, -1000, -1000, 125,
	-1000, 1485, -1000, -1000, 268, 124, 6, 123, 122, 206,
	86, 86, -24, -24, -24, -24, 1218, 1186, 78, 78,
	78, 78, 78, 78, 78, 1485, -1000, 959, 52, 160,
	-1000, 29, -1000, -1000, -1000, -16, 1309, 1309, 181, -1000,
	-1000, -1000, 1115, -1000, -60, 920, 1485, 1485, 1485, 1485,
	116, 1485, 114, 1485, -1000, 42, 1485, -1000, 1485, -1000,
	-1000, -1000, -1000, 102, -1000, -1000, -5, 98, 235, 1485,
	1485, 3, -1000, 165, 197, -41, -1000, 75, -49, 1485,
	-50, -1000, -1000, 1485, 78, -1000, 164, 196, -1000, -1000,
	-1000, -1000, -17, -18, 1309, 37, -62, 1485, 1485, 881,
	842, 803, 764, -68, 374, -69, 335, -1000, -1000, -1000,
	-5, -1000, 235, -5, 1076, 1076, -79, 93, 1485, 1485,
	195, -1000, -1000, 89, -1000, -1000, -1000, 71, -20, 48,
	-1000, 189, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1076,
	725, -1000, 1485, 1485, -1000, 1485, 1485, -1000, 1485, 1485,
	-1000, -1000, -5, -1000, -1000, 230, 1076, 1076, -1000, -1000,
	266, 264, -23, -1000, 1485, 686, 647, 296, 608, 5,
	569, -1000, -79, 75, 75, 1485, -1000, -1000, -1000, 1485,
	-1000, -1000, 1485, -1000, -1000, -1000, -1000, -1000, -48, 530,
	491, 260, -1000, -1000, 75, -1000,
}
var yyPgo = []int{

	0, 380, 316, 378, 376, 375, 372, 371, 1, 16,
	355, 354, 353, 351, 252, 350, 314, 260, 349, 348,
	6, 347, 345, 344, 13, 342, 341, 0, 7, 338,
	11, 4, 12, 5, 334, 3, 10, 333, 332, 331,
	218, 330, 329, 328, 2, 326, 8, 323, 322, 321,
	318, 317, 9, 315,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 25,
	26, 26, 15, 15, 15, 18, 18, 28, 28, 28,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 31, 31, 31, 31, 33, 33, 33, 34, 34,
	34, 29, 29, 29, 35, 35, 16, 16, 12, 12,
	36, 36, 37, 37, 37, 13, 13, 13, 38, 39,
	20, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 40,
	40, 40, 41, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 44, 44, 45, 45, 32, 32, 32,
	32, 32, 32, 46, 46, 47, 47, 48, 48, 43,
	43, 43, 43, 43, 43, 43, 49, 49, 50, 50,
	52, 52, 53, 51, 51, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 0, 2, 5, 2, 5, 1, 2, 2,
	2, 4, 3, 3, 5, 3, 5, 4, 4, 6,
	5, 4, 5, 6, 7, 2, 2, 0, 1, 1,
	1, 1, 3, 2, 2, 2, 0, 2, 0, 3,
	1, 3, 1, 2, 2, 0, 1, 2, 2, 2,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 6,
	5, 5, 3, 4, 3, 4, 3, 4, 1, 2,
	2, 1, 1, 1, 1, 3, 5, 5, 7, 7,
	5, 9, 7, 7, 5, 9, 7, 7, 5, 3,
	4, 5, 5, 3, 5, 0, 2, 1, 4, 6,
	5, 5, 3, 1, 3, 1, 1, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 32,
	33, -28, 53, -29, -32, 59, 59, 53, -13, -38,
	44, 41, 29, 59, -16, -28, 53, -17, 46, -20,
	-27, -40, 68, 62, -41, -42, 59, -43, 75, 82,
	87, 93, 89, 90, 60, -49, -50, -51, 54, 55,
	56, 57, 58, 48, 50, -23, -24, -25, -26, -20,
	63, -27, 59, -30, -31, 94, 19, -34, 96, 97,
	98, 35, 59, 50, 81, 81, 59, -39, 45, 57,
	-36, -37, -20, 30, 30, -17, 59, -14, 41, 61,
	62, 63, 64, 99, 65, 66, 67, 69, 73, 74,
	71, 72, 70, 77, 68, 81, 50, 78, -40, -40,
	75, -20, 83, -27, 59, 59, -27, -27, -27, 49,
	-52, -53, 60, 51, -9, -20, 52, 35, 59, 81,
	81, -32, -32, 19, 59, 57, 53, 59, 59, 81,
	57, 52, 42, 43, 59, 53, 59, 53, 81, -9,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, 77, 59, -27, 57, 53,
	56, 68, 80, 79, 76, -46, 32, 33, -47, -48,
	-20, 63, -27, 76, -44, -27, 92, 91, 91, 92,
	95, 91, 95, 91, 49, 52, 53, 51, 52, -24,
	59, 63, -28, 35, 59, -30, -33, 35, 59, 38,
	39, -32, 51, 53, 57, 59, -36, 31, 59, 75,
	59, -28, -19, 47, -27, 51, 53, 57, 56, 80,
	79, 76, -46, -46, 52, 81, -45, 85, 84, -27,
	-27, -27, -27, 59, -27, 59, -27, -52, -20, -9,
	59, -30, 59, -33, -27, -27, -35, 35, 38, 39,
	57, 51, 51, 81, -8, 28, 59, 81, -9, 81,
	-20, 57, 51, 51, 76, 76, -46, 63, 86, -27,
	-27, 86, 92, 92, 86, 91, 83, 86, 91, 83,
	86, -30, -33, -30, -31, 59, -27, -27, 51, 59,
	59, 76, 59, 51, 83, -27, -27, -27, -27, -27,
	-27, -30, -35, 31, 31, 75, -44, 86, 86, 83,
	86, 86, 83, 86, 86, -31, -8, -8, -9, -27,
	-27, 76, 86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 78, 0, 42, 76, 30, 0, 29, 2, 0,
	85, 0, 0, 0, 76, 0, 24, 0, 0, 31,
	32, 45, 0, 47, 71, 147, 0, 0, 21, 86,
	0, 0, 0, 0, 24, 43, 0, 0, 0, 77,
	90, 118, 0, 0, 121, 122, 123, 124, 0, 0,
	0, 0, 0, 0, 159, 160, 161, 162, 163, 164,
	165, 166, 167, 0, 0, 28, 33, 34, 36, 37,
	40, 90, 0, 48, 49, 0, 0, 0, 68, 69,
	70, 0, 73, 0, 0, 0, 0, 87, 0, 88,
	79, 80, 82, 0, 0, 22, 0, 23, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 119, 120,
	0, 0, 0, 0, 123, 123, 0, 0, 0, 168,
	0, 170, 0, 173, 0, 175, 0, 0, 39, 0,
	0, 50, 67, 0, 72, 0, 0, 152, 18, 0,
	89, 0, 83, 84, 8, 0, 0, 0, 0, 26,
	91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
	101, 102, 103, 104, 105, 0, 107, 0, 166, 0,
	112, 0, 114, 116, 139, 0, 0, 0, 153, 155,
	156, 157, 90, 125, 145, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 169, 0, 0, 174, 0, 35,
	38, 41, 46, 0, 52, 53, 55, 0, 67, 0,
	0, 0, 148, 0, 0, 0, 81, 0, 0, 0,
	0, 44, 25, 0, 106, 108, 0, 0, 113, 115,
	117, 140, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 171, 172, 176,
	51, 58, 67, 57, 65, 66, 61, 0, 0, 0,
	0, 150, 151, 0, 10, 16, 17, 0, 0, 0,
	27, 0, 110, 111, 141, 142, 154, 158, 126, 146,
	143, 127, 0, 0, 130, 0, 0, 134, 0, 0,
	138, 54, 56, 60, 62, 0, 74, 75, 149, 19,
	9, 12, 0, 109, 0, 0, 0, 0, 0, 0,
	0, 59, 63, 0, 0, 0, 144, 128, 129, 0,
	133, 132, 0, 137, 136, 64, 11, 14, 0, 0,
	0, 13, 131, 135, 0, 15,
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
		//line n1ql.y:455
		{
	    	logDebugGrammar("FROM DATASOURCE WITH JOIN") 
	        rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 50:
		//line n1ql.y:466
		{
	    logDebugGrammar("UNNEST")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:""})
	}
	case 51:
		//line n1ql.y:473
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 52:
		//line n1ql.y:480
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 53:
		//line n1ql.y:487
		{
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: "", Over:rest})
	}
	case 54:
		//line n1ql.y:494
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 55:
		//line n1ql.y:501
		{
	    logDebugGrammar("JOIN KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr})
	}
	case 56:
		//line n1ql.y:508
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 57:
		//line n1ql.y:515
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 58:
		//line n1ql.y:522
		{
	    logDebugGrammar("JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr, Over: rest})
	}
	case 59:
		//line n1ql.y:530
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 60:
		//line n1ql.y:538
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 61:
		//line n1ql.y:549
		{
	    logDebugGrammar("TYPE JOIN KEY")
	}
	case 62:
		//line n1ql.y:553
		{
	    logDebugGrammar("TYPE JOIN KEY")
	}
	case 63:
		//line n1ql.y:557
		{
	    logDebugGrammar("TYPE JOIN AS KEY")
	}
	case 64:
		//line n1ql.y:561
		{
	    logDebugGrammar("TYPE JOIN AS KEY")
	}
	case 65:
		//line n1ql.y:567
		{
	        logDebugGrammar("FROM JOIN DATASOURCE with KEY")
	        key := parsingStack.Pop().(ast.Expression)
	        key_expr := ast.NewKeyExpression(key, "KEY")
	        parsingStack.Push(key_expr)
	}
	case 66:
		//line n1ql.y:574
		{
	        logDebugGrammar("FROM DATASOURCE with KEYS")
	        keys := parsingStack.Pop().(ast.Expression)
	        keys_expr := ast.NewKeyExpression(keys, "KEYS")
	        parsingStack.Push(keys_expr)
	
	}
	case 68:
		//line n1ql.y:584
		{
	    logDebugGrammar("INNER")
	}
	case 69:
		//line n1ql.y:588
		{
	    logDebugGrammar("OUTER")
	}
	case 70:
		//line n1ql.y:592
		{
	    logDebugGrammar("OUTER")
	}
	case 71:
		//line n1ql.y:598
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 72:
		//line n1ql.y:604
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 73:
		//line n1ql.y:611
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 74:
		//line n1ql.y:620
		{
	        logDebugGrammar("FROM DATASOURCE with KEY")
	        keys := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Keys = ast.NewKeyExpression(keys, "KEY") 
		default:
			logDebugGrammar("This statement does not support KEY")
		}
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection:proj})
	}
	case 75:
		//line n1ql.y:633
		{
	        logDebugGrammar("FROM DATASOURCE with KEYS")
	        keys := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Keys = ast.NewKeyExpression(keys, "KEYS")
		default:
			logDebugGrammar("This statement does not support KEYS")
		}
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 76:
		//line n1ql.y:648
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 77:
		//line n1ql.y:652
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
	case 79:
		//line n1ql.y:666
		{
	
	}
	case 80:
		//line n1ql.y:672
		{
	
	}
	case 81:
		//line n1ql.y:676
		{
	
	}
	case 82:
		//line n1ql.y:681
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
	case 83:
		//line n1ql.y:692
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
	case 84:
		//line n1ql.y:703
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
	case 85:
		//line n1ql.y:715
		{
	
	}
	case 86:
		//line n1ql.y:719
		{
	
	}
	case 87:
		//line n1ql.y:723
		{
	
	}
	case 88:
		//line n1ql.y:729
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
	case 89:
		//line n1ql.y:743
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
	case 90:
		//line n1ql.y:760
		{
		logDebugGrammar("EXPRESSION")
	}
	case 91:
		//line n1ql.y:765
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line n1ql.y:773
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line n1ql.y:781
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line n1ql.y:789
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 95:
		//line n1ql.y:797
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line n1ql.y:805
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:813
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line n1ql.y:821
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line n1ql.y:829
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line n1ql.y:837
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line n1ql.y:845
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line n1ql.y:853
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line n1ql.y:861
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line n1ql.y:869
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 105:
		//line n1ql.y:877
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line n1ql.y:885
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line n1ql.y:893
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 108:
		//line n1ql.y:901
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 109:
		//line n1ql.y:909
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 110:
		//line n1ql.y:916
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 111:
		//line n1ql.y:924
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 112:
		//line n1ql.y:931
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 113:
		//line n1ql.y:938
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line n1ql.y:945
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line n1ql.y:952
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 116:
		//line n1ql.y:959
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 117:
		//line n1ql.y:966
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 118:
		//line n1ql.y:973
		{
	
	}
	case 119:
		//line n1ql.y:979
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 120:
		//line n1ql.y:986
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 121:
		//line n1ql.y:993
		{
	
	}
	case 122:
		//line n1ql.y:998
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 123:
		//line n1ql.y:1004
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line n1ql.y:1010
		{
		logDebugGrammar("LITERAL")
	}
	case 125:
		//line n1ql.y:1014
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 126:
		//line n1ql.y:1018
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
	case 127:
		//line n1ql.y:1035
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 128:
		//line n1ql.y:1043
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 129:
		//line n1ql.y:1051
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 130:
		//line n1ql.y:1059
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 131:
		//line n1ql.y:1067
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 132:
		//line n1ql.y:1076
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 133:
		//line n1ql.y:1085
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 134:
		//line n1ql.y:1093
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 135:
		//line n1ql.y:1101
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 136:
		//line n1ql.y:1110
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 137:
		//line n1ql.y:1119
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 138:
		//line n1ql.y:1127
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 139:
		//line n1ql.y:1135
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 140:
		//line n1ql.y:1141
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 141:
		//line n1ql.y:1148
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 142:
		//line n1ql.y:1156
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 143:
		//line n1ql.y:1165
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 144:
		//line n1ql.y:1173
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
	case 145:
		//line n1ql.y:1187
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 146:
		//line n1ql.y:1191
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 147:
		//line n1ql.y:1197
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1203
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 149:
		//line n1ql.y:1210
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 150:
		//line n1ql.y:1217
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 151:
		//line n1ql.y:1225
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 152:
		//line n1ql.y:1232
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 153:
		//line n1ql.y:1243
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 154:
		//line n1ql.y:1248
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
	case 155:
		//line n1ql.y:1262
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 156:
		//line n1ql.y:1266
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 157:
		//line n1ql.y:1275
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 158:
		//line n1ql.y:1281
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 159:
		//line n1ql.y:1291
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 160:
		//line n1ql.y:1297
		{
		logDebugGrammar("NUMBER")
	}
	case 161:
		//line n1ql.y:1301
		{
		logDebugGrammar("OBJECT")
	}
	case 162:
		//line n1ql.y:1305
		{
		logDebugGrammar("ARRAY")
	}
	case 163:
		//line n1ql.y:1309
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 164:
		//line n1ql.y:1315
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 165:
		//line n1ql.y:1321
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 166:
		//line n1ql.y:1329
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 167:
		//line n1ql.y:1335
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 168:
		//line n1ql.y:1343
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 169:
		//line n1ql.y:1349
		{
		logDebugGrammar("OBJECT")
	}
	case 170:
		//line n1ql.y:1355
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 171:
		//line n1ql.y:1359
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 172:
		//line n1ql.y:1371
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 173:
		//line n1ql.y:1381
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 174:
		//line n1ql.y:1387
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 175:
		//line n1ql.y:1396
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 176:
		//line n1ql.y:1403
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
