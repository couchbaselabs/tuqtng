
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
const MOD = 57438

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

const yyNprod = 159
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1485

var yyAct = []int{

	50, 264, 200, 83, 136, 31, 141, 96, 191, 84,
	76, 288, 285, 73, 126, 74, 140, 278, 246, 68,
	69, 70, 71, 72, 56, 64, 126, 53, 197, 81,
	204, 45, 128, 52, 49, 79, 219, 34, 269, 267,
	58, 263, 203, 237, 164, 155, 322, 59, 98, 146,
	91, 89, 60, 181, 62, 63, 295, 122, 61, 275,
	274, 129, 132, 133, 134, 127, 239, 238, 105, 106,
	107, 108, 110, 111, 112, 120, 113, 118, 116, 117,
	114, 115, 90, 307, 119, 123, 148, 149, 121, 240,
	314, 199, 228, 315, 126, 84, 182, 182, 163, 138,
	277, 217, 296, 109, 162, 265, 166, 167, 168, 169,
	170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
	180, 165, 147, 183, 135, 161, 122, 198, 294, 201,
	122, 160, 122, 196, 235, 138, 266, 105, 106, 107,
	108, 110, 143, 81, 120, 107, 108, 110, 120, 79,
	120, 220, 218, 215, 123, 293, 186, 121, 123, 236,
	123, 121, 259, 121, 98, 225, 144, 46, 187, 254,
	230, 87, 109, 35, 85, 86, 252, 37, 109, 189,
	188, 32, 233, 36, 35, 229, 89, 35, 227, 224,
	216, 182, 154, 198, 198, 88, 153, 150, 102, 196,
	196, 241, 242, 248, 249, 250, 251, 92, 253, 82,
	255, 43, 272, 223, 261, 156, 256, 90, 271, 257,
	260, 152, 95, 212, 243, 151, 51, 221, 297, 222,
	214, 258, 211, 157, 142, 292, 273, 262, 213, 270,
	210, 232, 48, 13, 198, 268, 94, 279, 280, 47,
	196, 40, 276, 104, 122, 158, 159, 41, 21, 27,
	25, 17, 325, 291, 306, 105, 106, 107, 108, 110,
	111, 112, 120, 113, 118, 116, 117, 114, 115, 124,
	125, 119, 123, 299, 300, 121, 301, 302, 305, 303,
	304, 103, 29, 30, 101, 226, 283, 26, 100, 201,
	109, 308, 12, 10, 3, 12, 10, 317, 318, 99,
	42, 17, 320, 16, 17, 321, 16, 122, 22, 19,
	23, 2, 44, 137, 319, 18, 67, 326, 105, 106,
	107, 108, 110, 111, 112, 120, 113, 118, 116, 117,
	114, 115, 66, 65, 119, 123, 195, 194, 121, 245,
	311, 57, 55, 312, 122, 54, 93, 39, 97, 33,
	78, 77, 75, 109, 28, 105, 106, 107, 108, 110,
	111, 112, 120, 113, 118, 116, 117, 114, 115, 15,
	231, 119, 123, 14, 24, 121, 38, 289, 20, 11,
	290, 122, 7, 9, 8, 6, 5, 4, 1, 0,
	109, 0, 105, 106, 107, 108, 110, 111, 112, 120,
	113, 118, 116, 117, 114, 115, 0, 0, 119, 123,
	0, 0, 121, 0, 286, 0, 0, 287, 122, 0,
	0, 0, 0, 0, 0, 0, 0, 109, 0, 105,
	106, 107, 108, 110, 111, 112, 120, 113, 118, 116,
	117, 114, 115, 0, 0, 119, 123, 0, 0, 121,
	0, 0, 0, 0, 122, 0, 0, 0, 0, 209,
	0, 0, 0, 208, 109, 105, 106, 107, 108, 110,
	111, 112, 120, 113, 118, 116, 117, 114, 115, 0,
	0, 119, 123, 0, 0, 121, 0, 0, 0, 0,
	122, 0, 0, 0, 0, 207, 0, 0, 0, 206,
	109, 105, 106, 107, 108, 110, 111, 112, 120, 113,
	118, 116, 117, 114, 115, 0, 0, 119, 123, 0,
	0, 121, 0, 0, 0, 0, 324, 122, 0, 0,
	0, 0, 0, 0, 0, 0, 109, 0, 105, 106,
	107, 108, 110, 111, 112, 120, 113, 118, 116, 117,
	114, 115, 0, 0, 119, 123, 0, 0, 121, 0,
	0, 0, 0, 323, 122, 0, 0, 0, 0, 0,
	0, 0, 0, 109, 0, 105, 106, 107, 108, 110,
	111, 112, 120, 113, 118, 116, 117, 114, 115, 0,
	0, 119, 123, 0, 0, 121, 0, 0, 0, 0,
	316, 122, 0, 0, 0, 0, 0, 0, 0, 0,
	109, 0, 105, 106, 107, 108, 110, 111, 112, 120,
	113, 118, 116, 117, 114, 115, 0, 0, 119, 123,
	0, 0, 121, 0, 0, 0, 0, 313, 122, 0,
	0, 0, 0, 0, 0, 0, 0, 109, 0, 105,
	106, 107, 108, 110, 111, 112, 120, 113, 118, 116,
	117, 114, 115, 0, 0, 119, 123, 0, 0, 121,
	0, 0, 0, 0, 310, 122, 0, 0, 0, 0,
	0, 0, 0, 0, 109, 0, 105, 106, 107, 108,
	110, 111, 112, 120, 113, 118, 116, 117, 114, 115,
	0, 0, 119, 123, 0, 0, 121, 0, 0, 0,
	0, 309, 122, 0, 0, 0, 0, 0, 0, 0,
	0, 109, 0, 105, 106, 107, 108, 110, 111, 112,
	120, 113, 118, 116, 117, 114, 115, 0, 0, 119,
	123, 0, 0, 121, 0, 298, 0, 0, 122, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 109, 105,
	106, 107, 108, 110, 111, 112, 120, 113, 118, 116,
	117, 114, 115, 0, 0, 119, 123, 0, 0, 121,
	0, 0, 0, 0, 284, 122, 0, 0, 0, 0,
	0, 0, 0, 0, 109, 0, 105, 106, 107, 108,
	110, 111, 112, 120, 113, 118, 116, 117, 114, 115,
	0, 0, 119, 123, 0, 0, 121, 0, 0, 0,
	0, 122, 0, 0, 0, 0, 0, 282, 0, 0,
	0, 109, 105, 106, 107, 108, 110, 111, 112, 120,
	113, 118, 116, 117, 114, 115, 0, 0, 119, 123,
	0, 0, 121, 0, 0, 0, 0, 281, 122, 0,
	0, 0, 0, 0, 0, 0, 0, 109, 0, 105,
	106, 107, 108, 110, 111, 112, 120, 113, 118, 116,
	117, 114, 115, 0, 0, 119, 123, 0, 0, 121,
	0, 0, 247, 0, 122, 234, 0, 0, 0, 0,
	0, 0, 0, 0, 109, 105, 106, 107, 108, 110,
	111, 112, 120, 113, 118, 116, 117, 114, 115, 0,
	0, 119, 123, 0, 0, 121, 0, 0, 0, 0,
	122, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	109, 105, 106, 107, 108, 110, 111, 112, 120, 113,
	118, 116, 117, 114, 115, 0, 0, 119, 123, 0,
	0, 121, 0, 0, 0, 0, 122, 0, 0, 0,
	0, 0, 205, 0, 0, 0, 109, 105, 106, 107,
	108, 110, 111, 112, 120, 113, 118, 116, 117, 114,
	115, 0, 0, 119, 123, 0, 0, 121, 0, 0,
	0, 0, 122, 0, 0, 0, 0, 0, 202, 0,
	0, 0, 109, 105, 106, 107, 108, 110, 111, 112,
	120, 113, 118, 116, 117, 114, 115, 0, 0, 119,
	123, 0, 0, 121, 0, 0, 0, 0, 122, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 109, 105,
	106, 107, 108, 110, 111, 112, 120, 113, 118, 116,
	117, 114, 115, 0, 0, 119, 123, 0, 0, 244,
	0, 0, 0, 0, 122, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 109, 105, 106, 107, 108, 110,
	111, 112, 120, 113, 118, 116, 117, 114, 115, 0,
	0, 119, 123, 122, 0, 145, 0, 0, 0, 0,
	0, 0, 0, 0, 105, 106, 107, 108, 110, 111,
	109, 120, 113, 118, 116, 117, 114, 115, 192, 193,
	119, 123, 0, 0, 121, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 73, 0, 74, 0, 0, 109,
	68, 69, 70, 71, 72, 56, 64, 0, 53, 197,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 0,
	0, 58, 190, 0, 0, 0, 0, 0, 59, 122,
	0, 0, 0, 60, 0, 62, 63, 0, 0, 61,
	105, 106, 107, 108, 110, 0, 0, 120, 113, 118,
	116, 117, 114, 115, 0, 0, 119, 123, 0, 73,
	121, 74, 0, 0, 0, 68, 69, 70, 71, 72,
	56, 64, 0, 53, 80, 109, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 0, 0, 60, 0,
	62, 63, 0, 73, 61, 74, 0, 0, 185, 68,
	69, 70, 184, 72, 56, 64, 0, 53, 0, 0,
	0, 0, 0, 52, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 60, 0, 62, 63, 0, 73, 61, 74,
	139, 0, 0, 68, 69, 70, 71, 72, 56, 64,
	0, 53, 0, 0, 0, 0, 0, 52, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 60, 0, 62, 63,
	0, 73, 61, 74, 0, 0, 0, 68, 69, 70,
	71, 72, 56, 64, 0, 53, 0, 0, 0, 0,
	0, 52, 0, 0, 0, 0, 0, 0, 58, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	60, 0, 62, 63, 0, 73, 61, 74, 0, 0,
	0, 68, 69, 70, 71, 72, 131, 64, 0, 53,
	0, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 62, 63, 0, 73,
	61, 74, 0, 0, 0, 68, 69, 70, 71, 72,
	130, 64, 0, 53, 0, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 0, 0, 60, 0,
	62, 63, 0, 0, 61,
}
var yyPact = []int{

	280, -1000, -1000, 277, -1000, -1000, -1000, -1000, -1000, -1000,
	290, 218, 291, 224, 222, 260, 128, -1000, -1000, 124,
	207, 216, 281, 152, 222, 114, 196, 1303, 1171, -1000,
	-1000, -1000, 150, -85, 136, -1000, -31, 148, -1000, 201,
	165, 1303, 279, 268, 196, -1000, 139, 227, 212, -1000,
	962, -1000, 1303, 1303, -1000, -1000, 19, -1000, 1303, -51,
	1391, 1347, 1303, 1303, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 75, 1259, -1000, -1000, 182, -1000, 107,
	-1000, 1034, -32, -1000, 125, 1303, 1303, 138, -1000, 168,
	137, 133, -36, -1000, 158, -1000, -1000, 181, 213, 72,
	45, -1000, -37, -1000, 1303, 1303, 1303, 1303, 1303, 1303,
	1303, 1303, 1303, 1303, 1303, 1303, 1303, 1303, 1303, 1303,
	-24, 132, 1215, 100, -1000, -1000, 1106, 15, 1303, 926,
	-49, -61, 890, 414, 378, -1000, 191, 180, 170, -1000,
	187, 178, 1171, 131, -1000, 38, 125, 1, 962, 962,
	-1000, 176, 156, -1000, -1000, 130, -1000, 1303, -1000, -1000,
	264, 129, 17, 126, 125, 194, 82, 82, 80, 80,
	80, 80, 1139, 1063, 76, 76, 76, 76, 76, 76,
	76, 1303, -1000, 854, 81, 102, -1000, -13, -1000, -1000,
	-1000, 13, -35, -35, 172, -1000, -1000, -1000, 998, -1000,
	-67, 818, 1303, 1303, 1303, 1303, 117, 1303, 110, 1303,
	-1000, 39, 1303, -1000, 1303, -1000, -1000, -1000, -1000, 103,
	-1000, -1000, 163, 186, -40, -1000, 77, -42, 1303, -43,
	-1000, -1000, 1303, 76, -1000, 161, 185, -1000, -1000, -1000,
	-1000, -16, -17, -35, 37, -69, 1303, 1303, 781, 745,
	204, 708, -79, 341, -80, 304, -1000, -1000, -1000, -85,
	184, -1000, -1000, 96, -1000, -1000, -1000, 69, -20, 43,
	-1000, 177, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 962,
	672, -1000, 1303, 1303, -1000, 1303, 1303, -1000, 1303, 1303,
	-1000, -1000, -1000, -1000, 257, 233, 8, -1000, 1303, 635,
	598, 267, 561, 7, 524, 77, 77, 1303, -1000, -1000,
	-1000, 1303, -1000, -1000, 1303, -1000, -1000, -1000, -1000, -30,
	487, 450, 231, -1000, -1000, 77, -1000,
}
var yyPgo = []int{

	0, 398, 321, 397, 396, 395, 394, 393, 1, 16,
	392, 389, 388, 386, 243, 384, 297, 249, 383, 380,
	6, 379, 364, 362, 10, 361, 360, 0, 5, 359,
	3, 37, 7, 358, 357, 356, 226, 355, 352, 351,
	2, 349, 8, 347, 346, 343, 342, 326, 4, 323,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 25,
	26, 26, 15, 15, 15, 18, 18, 28, 28, 30,
	30, 30, 30, 29, 29, 29, 29, 29, 16, 16,
	12, 12, 32, 32, 33, 33, 33, 13, 13, 13,
	34, 35, 20, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 36, 36, 36, 37, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 40, 40, 41, 41, 31,
	31, 31, 31, 31, 31, 42, 42, 43, 43, 44,
	44, 39, 39, 39, 39, 39, 39, 39, 45, 45,
	46, 46, 48, 48, 49, 47, 47, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 0, 2, 5, 2, 5, 1, 2, 2,
	4, 3, 5, 1, 3, 3, 3, 2, 0, 2,
	0, 3, 1, 3, 1, 2, 2, 0, 1, 2,
	2, 2, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 6, 5, 5, 3, 4, 3, 4, 3, 4,
	1, 2, 2, 1, 1, 1, 1, 3, 5, 5,
	7, 7, 5, 9, 7, 7, 5, 9, 7, 7,
	5, 3, 4, 5, 5, 3, 5, 0, 2, 1,
	4, 6, 5, 5, 3, 1, 3, 1, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 3, 1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 32,
	33, -28, 53, -29, -31, 59, 59, 53, -13, -34,
	44, 41, 29, 59, -16, -28, 53, -17, 46, -20,
	-27, -36, 68, 62, -37, -38, 59, -39, 75, 82,
	87, 93, 89, 90, 60, -45, -46, -47, 54, 55,
	56, 57, 58, 48, 50, -23, -24, -25, -26, -20,
	63, -27, 59, -30, 94, 38, 39, 35, 59, 50,
	81, 81, 59, -35, 45, 57, -32, -33, -20, 30,
	30, -17, 59, -14, 41, 61, 62, 63, 64, 96,
	65, 66, 67, 69, 73, 74, 71, 72, 70, 77,
	68, 81, 50, 78, -36, -36, 75, -20, 83, -27,
	59, 59, -27, -27, -27, 49, -48, -49, 60, 51,
	-9, -20, 52, 35, 59, 81, 81, -31, -27, -27,
	59, 57, 53, 59, 59, 81, 57, 52, 42, 43,
	59, 53, 59, 53, 81, -9, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, 77, 59, -27, 57, 53, 56, 68, 80, 79,
	76, -42, 32, 33, -43, -44, -20, 63, -27, 76,
	-40, -27, 92, 91, 91, 92, 95, 91, 95, 91,
	49, 52, 53, 51, 52, -24, 59, 63, -28, 35,
	-30, 51, 53, 57, 59, -32, 31, 59, 75, 59,
	-28, -19, 47, -27, 51, 53, 57, 56, 80, 79,
	76, -42, -42, 52, 81, -41, 85, 84, -27, -27,
	-27, -27, 59, -27, 59, -27, -48, -20, -9, 59,
	57, 51, 51, 81, -8, 28, 59, 81, -9, 81,
	-20, 57, 51, 51, 76, 76, -42, 63, 86, -27,
	-27, 86, 92, 92, 86, 91, 83, 86, 91, 83,
	86, -30, 51, 59, 59, 76, 59, 51, 83, -27,
	-27, -27, -27, -27, -27, 31, 31, 75, -40, 86,
	86, 83, 86, 86, 83, 86, 86, -8, -8, -9,
	-27, -27, 76, 86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 60, 0, 42, 58, 30, 0, 29, 2, 0,
	67, 0, 0, 0, 58, 0, 24, 0, 0, 31,
	32, 45, 0, 47, 53, 129, 0, 0, 21, 68,
	0, 0, 0, 0, 24, 43, 0, 0, 0, 59,
	72, 100, 0, 0, 103, 104, 105, 106, 0, 0,
	0, 0, 0, 0, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 0, 0, 28, 33, 34, 36, 37,
	40, 72, 0, 48, 0, 0, 0, 0, 57, 0,
	0, 0, 0, 69, 0, 70, 61, 62, 64, 0,
	0, 22, 0, 23, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 101, 102, 0, 0, 0, 0,
	105, 105, 0, 0, 0, 150, 0, 152, 0, 155,
	0, 157, 0, 0, 39, 0, 0, 49, 54, 55,
	56, 0, 0, 134, 18, 0, 71, 0, 65, 66,
	8, 0, 0, 0, 0, 26, 73, 74, 75, 76,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 0, 89, 0, 148, 0, 94, 0, 96, 98,
	121, 0, 0, 0, 135, 137, 138, 139, 72, 107,
	127, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	151, 0, 0, 156, 0, 35, 38, 41, 46, 0,
	51, 130, 0, 0, 0, 63, 0, 0, 0, 0,
	44, 25, 0, 88, 90, 0, 0, 95, 97, 99,
	122, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 153, 154, 158, 50,
	0, 132, 133, 0, 10, 16, 17, 0, 0, 0,
	27, 0, 92, 93, 123, 124, 136, 140, 108, 128,
	125, 109, 0, 0, 112, 0, 0, 116, 0, 0,
	120, 52, 131, 19, 9, 12, 0, 91, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 126, 110,
	111, 0, 115, 114, 0, 119, 118, 11, 14, 0,
	0, 0, 13, 113, 117, 0, 15,
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
	92, 93, 94, 95, 96,
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
		//line n1ql.y:54
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line n1ql.y:58
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line n1ql.y:64
		{
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line n1ql.y:68
		{
	}
	case 5:
		//line n1ql.y:71
		{
		logDebugGrammar("STMT - DROP INDEX")
	}
	case 6:
		//line n1ql.y:78
		{
		logDebugGrammar("STMT - CREATE PRIMARY INDEX")
	}
	case 7:
		//line n1ql.y:82
		{
		logDebugGrammar("STMT - CREATE SECONDARY INDEX")
	}
	case 8:
		//line n1ql.y:88
		{
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 9:
		//line n1ql.y:96
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
		//line n1ql.y:106
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
		//line n1ql.y:116
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
		//line n1ql.y:130
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
		//line n1ql.y:142
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
		//line n1ql.y:156
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
		//line n1ql.y:170
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
		//line n1ql.y:189
		{
		parsingStack.Push("view")
	}
	case 17:
		//line n1ql.y:193
		{
		parsingStack.Push(yyS[yypt-0].s)
	}
	case 18:
		//line n1ql.y:199
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 19:
		//line n1ql.y:208
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
		//line n1ql.y:222
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 21:
		//line n1ql.y:228
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND")
	}
	case 22:
		//line n1ql.y:235
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 23:
		//line n1ql.y:239
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 24:
		//line n1ql.y:245
		{
	}
	case 25:
		//line n1ql.y:248
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
		//line n1ql.y:260
		{
	}
	case 27:
		//line n1ql.y:263
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
		//line n1ql.y:276
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 29:
		//line n1ql.y:282
		{
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 30:
		//line n1ql.y:288
		{
	}
	case 31:
		//line n1ql.y:291
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
		//line n1ql.y:301
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
		//line n1ql.y:313
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
		//line n1ql.y:327
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 35:
		//line n1ql.y:332
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
		//line n1ql.y:345
		{
		logDebugGrammar("RESULT STAR")
	}
	case 37:
		//line n1ql.y:349
		{
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 38:
		//line n1ql.y:356
		{
		logDebugGrammar("RESULT EXPR AS ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 39:
		//line n1ql.y:363
		{
		logDebugGrammar("RESULT EXPR ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 40:
		//line n1ql.y:372
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 41:
		//line n1ql.y:378
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 42:
		//line n1ql.y:387
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 43:
		//line n1ql.y:391
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
		//line n1ql.y:402
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
		//line n1ql.y:417
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
		//line n1ql.y:428
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
		//line n1ql.y:442
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT UNNEST")
	}
	case 48:
		//line n1ql.y:446
		{
		logDebugGrammar("FROM DATASOURCE WITH UNNEST")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 49:
		//line n1ql.y:457
		{
	    logDebugGrammar("UNNEST")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:""})
	}
	case 50:
		//line n1ql.y:464
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 51:
		//line n1ql.y:471
		{
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: "", Over:rest})
	}
	case 52:
		//line n1ql.y:478
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 53:
		//line n1ql.y:487
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 54:
		//line n1ql.y:493
		{
	        logDebugGrammar("FROM DATASOURCE with KEY")
	        keys := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Keys = ast.NewKeyExpression(keys, "key") 
		default:
			logDebugGrammar("This statement does not support KEY")
		}
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 55:
		//line n1ql.y:506
		{
	        logDebugGrammar("FROM DATASOURCE with KEYS")
	        keys := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Keys = ast.NewKeyExpression(keys, "keys")
		default:
			logDebugGrammar("This statement does not support KEYS")
		}
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, Keys: keys})
	}
	case 56:
		//line n1ql.y:519
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 57:
		//line n1ql.y:526
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 58:
		//line n1ql.y:535
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 59:
		//line n1ql.y:539
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
	case 61:
		//line n1ql.y:553
		{
	
	}
	case 62:
		//line n1ql.y:559
		{
	
	}
	case 63:
		//line n1ql.y:563
		{
	
	}
	case 64:
		//line n1ql.y:568
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
	case 65:
		//line n1ql.y:579
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
	case 66:
		//line n1ql.y:590
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
	case 67:
		//line n1ql.y:602
		{
	
	}
	case 68:
		//line n1ql.y:606
		{
	
	}
	case 69:
		//line n1ql.y:610
		{
	
	}
	case 70:
		//line n1ql.y:616
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
	case 71:
		//line n1ql.y:630
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
	case 72:
		//line n1ql.y:647
		{
		logDebugGrammar("EXPRESSION")
	}
	case 73:
		//line n1ql.y:652
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line n1ql.y:660
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line n1ql.y:668
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line n1ql.y:676
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line n1ql.y:684
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line n1ql.y:692
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line n1ql.y:700
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line n1ql.y:708
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line n1ql.y:716
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line n1ql.y:724
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line n1ql.y:732
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line n1ql.y:740
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line n1ql.y:748
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line n1ql.y:756
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 87:
		//line n1ql.y:764
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line n1ql.y:772
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line n1ql.y:780
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line n1ql.y:788
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line n1ql.y:796
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 92:
		//line n1ql.y:803
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 93:
		//line n1ql.y:811
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 94:
		//line n1ql.y:818
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 95:
		//line n1ql.y:825
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line n1ql.y:832
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:839
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line n1ql.y:846
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line n1ql.y:853
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line n1ql.y:860
		{
	
	}
	case 101:
		//line n1ql.y:866
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line n1ql.y:873
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line n1ql.y:880
		{
	
	}
	case 104:
		//line n1ql.y:885
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 105:
		//line n1ql.y:891
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line n1ql.y:897
		{
		logDebugGrammar("LITERAL")
	}
	case 107:
		//line n1ql.y:901
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 108:
		//line n1ql.y:905
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
	case 109:
		//line n1ql.y:922
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 110:
		//line n1ql.y:930
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 111:
		//line n1ql.y:938
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 112:
		//line n1ql.y:946
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 113:
		//line n1ql.y:954
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 114:
		//line n1ql.y:963
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 115:
		//line n1ql.y:972
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 116:
		//line n1ql.y:980
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 117:
		//line n1ql.y:988
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 118:
		//line n1ql.y:997
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 119:
		//line n1ql.y:1006
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 120:
		//line n1ql.y:1014
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 121:
		//line n1ql.y:1022
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 122:
		//line n1ql.y:1028
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 123:
		//line n1ql.y:1035
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 124:
		//line n1ql.y:1043
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 125:
		//line n1ql.y:1052
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 126:
		//line n1ql.y:1060
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
	case 127:
		//line n1ql.y:1074
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 128:
		//line n1ql.y:1078
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 129:
		//line n1ql.y:1084
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 130:
		//line n1ql.y:1090
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 131:
		//line n1ql.y:1097
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 132:
		//line n1ql.y:1104
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 133:
		//line n1ql.y:1112
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 134:
		//line n1ql.y:1119
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 135:
		//line n1ql.y:1130
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 136:
		//line n1ql.y:1135
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
	case 137:
		//line n1ql.y:1149
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 138:
		//line n1ql.y:1153
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 139:
		//line n1ql.y:1162
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 140:
		//line n1ql.y:1168
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 141:
		//line n1ql.y:1178
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 142:
		//line n1ql.y:1184
		{
		logDebugGrammar("NUMBER")
	}
	case 143:
		//line n1ql.y:1188
		{
		logDebugGrammar("OBJECT")
	}
	case 144:
		//line n1ql.y:1192
		{
		logDebugGrammar("ARRAY")
	}
	case 145:
		//line n1ql.y:1196
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 146:
		//line n1ql.y:1202
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 147:
		//line n1ql.y:1208
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1216
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 149:
		//line n1ql.y:1222
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 150:
		//line n1ql.y:1230
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 151:
		//line n1ql.y:1236
		{
		logDebugGrammar("OBJECT")
	}
	case 152:
		//line n1ql.y:1242
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 153:
		//line n1ql.y:1246
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 154:
		//line n1ql.y:1258
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 155:
		//line n1ql.y:1268
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 156:
		//line n1ql.y:1274
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 157:
		//line n1ql.y:1283
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 158:
		//line n1ql.y:1290
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
