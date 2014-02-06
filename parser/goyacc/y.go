
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
const BUCKET = 57347
const CAST = 57348
const COLLATE = 57349
const DATABASE = 57350
const DELETE = 57351
const EACH = 57352
const EXCEPT = 57353
const EXISTS = 57354
const IF = 57355
const INLINE = 57356
const INSERT = 57357
const INTERSECT = 57358
const INTO = 57359
const JOIN = 57360
const PATH = 57361
const UNION = 57362
const UPDATE = 57363
const POOL = 57364
const EXPLAIN = 57365
const CREATE = 57366
const DROP = 57367
const PRIMARY = 57368
const VIEW = 57369
const INDEX = 57370
const ON = 57371
const USING = 57372
const DISTINCT = 57373
const UNIQUE = 57374
const SELECT = 57375
const AS = 57376
const FROM = 57377
const WHERE = 57378
const KEY = 57379
const KEYS = 57380
const ORDER = 57381
const BY = 57382
const ASC = 57383
const DESC = 57384
const LIMIT = 57385
const OFFSET = 57386
const GROUP = 57387
const HAVING = 57388
const LBRACE = 57389
const RBRACE = 57390
const LBRACKET = 57391
const RBRACKET = 57392
const COMMA = 57393
const COLON = 57394
const TRUE = 57395
const FALSE = 57396
const NULL = 57397
const INT = 57398
const NUMBER = 57399
const IDENTIFIER = 57400
const STRING = 57401
const PLUS = 57402
const MINUS = 57403
const MULT = 57404
const DIV = 57405
const CONCAT = 57406
const AND = 57407
const OR = 57408
const NOT = 57409
const EQ = 57410
const NE = 57411
const GT = 57412
const GTE = 57413
const LT = 57414
const LTE = 57415
const LPAREN = 57416
const RPAREN = 57417
const LIKE = 57418
const IS = 57419
const VALUED = 57420
const MISSING = 57421
const BETWEEN = 57422
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
const NEST = 57438
const INNER = 57439
const LEFT = 57440
const OUTER = 57441
const MOD = 57442

var yyToknames = []string{
	"ALTER",
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
	"BETWEEN",
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
	"NEST",
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
	-1, 333,
	65, 133,
	66, 133,
	-2, 120,
	-1, 375,
	65, 133,
	66, 133,
	-2, 121,
}

const yyNprod = 213
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1835

var yyAct = []int{

	51, 326, 85, 152, 227, 218, 136, 250, 32, 52,
	78, 104, 88, 4, 165, 92, 35, 161, 188, 352,
	153, 349, 186, 132, 342, 142, 142, 144, 291, 83,
	331, 97, 98, 187, 46, 329, 118, 119, 121, 325,
	183, 229, 232, 231, 30, 31, 174, 158, 50, 81,
	99, 133, 188, 219, 220, 131, 413, 373, 339, 88,
	338, 285, 278, 106, 145, 148, 149, 150, 134, 54,
	226, 76, 395, 138, 120, 71, 72, 73, 74, 75,
	59, 67, 143, 56, 224, 280, 279, 374, 86, 55,
	89, 90, 91, 160, 266, 162, 61, 217, 171, 172,
	142, 29, 135, 159, 62, 163, 164, 372, 155, 63,
	167, 65, 66, 138, 185, 64, 184, 190, 191, 192,
	193, 194, 195, 196, 197, 198, 199, 200, 201, 202,
	203, 204, 156, 206, 182, 86, 189, 89, 90, 91,
	181, 371, 327, 225, 313, 228, 310, 253, 254, 253,
	254, 180, 205, 132, 205, 83, 341, 179, 242, 95,
	47, 95, 246, 223, 365, 240, 36, 243, 312, 362,
	309, 229, 255, 328, 356, 81, 320, 247, 248, 249,
	256, 133, 258, 253, 254, 131, 209, 272, 263, 274,
	38, 96, 268, 96, 33, 95, 37, 106, 210, 315,
	36, 302, 299, 297, 257, 36, 267, 265, 273, 212,
	211, 262, 241, 205, 173, 251, 170, 166, 253, 254,
	225, 225, 283, 281, 110, 286, 287, 96, 100, 84,
	95, 293, 294, 295, 296, 44, 298, 284, 300, 252,
	223, 223, 277, 301, 53, 261, 336, 175, 303, 17,
	306, 16, 335, 314, 317, 318, 308, 311, 319, 103,
	316, 276, 96, 307, 135, 321, 216, 132, 214, 169,
	330, 323, 333, 168, 259, 138, 260, 322, 116, 117,
	118, 119, 121, 93, 288, 229, 97, 98, 239, 225,
	215, 332, 343, 344, 340, 133, 134, 176, 95, 131,
	139, 141, 154, 376, 370, 355, 337, 94, 357, 223,
	359, 360, 324, 238, 363, 213, 13, 361, 120, 367,
	364, 358, 282, 366, 369, 270, 88, 49, 368, 48,
	96, 102, 177, 178, 112, 375, 41, 42, 253, 254,
	97, 98, 304, 21, 25, 97, 98, 378, 379, 27,
	380, 381, 17, 382, 383, 416, 17, 95, 16, 384,
	394, 386, 393, 264, 387, 111, 305, 389, 108, 391,
	388, 392, 385, 390, 107, 109, 88, 22, 228, 23,
	43, 19, 396, 2, 26, 137, 70, 18, 405, 96,
	69, 406, 244, 407, 68, 408, 409, 222, 221, 410,
	411, 290, 86, 412, 89, 90, 91, 95, 132, 45,
	60, 58, 57, 101, 40, 105, 245, 87, 417, 116,
	117, 118, 119, 121, 122, 123, 229, 124, 129, 127,
	128, 125, 126, 34, 80, 130, 133, 12, 10, 96,
	131, 79, 402, 77, 28, 403, 17, 15, 16, 132,
	269, 14, 86, 24, 89, 90, 91, 39, 20, 120,
	116, 117, 118, 119, 121, 122, 123, 229, 124, 129,
	127, 128, 125, 126, 11, 7, 130, 133, 9, 8,
	6, 131, 5, 399, 1, 0, 400, 3, 12, 10,
	132, 0, 0, 0, 0, 0, 0, 17, 0, 16,
	120, 116, 117, 118, 119, 121, 122, 123, 229, 124,
	129, 127, 128, 125, 126, 0, 0, 130, 133, 0,
	0, 0, 131, 0, 353, 0, 0, 354, 0, 0,
	0, 132, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 120, 116, 117, 118, 119, 121, 122, 123, 229,
	124, 129, 127, 128, 125, 126, 0, 0, 130, 133,
	0, 0, 0, 131, 0, 350, 0, 0, 351, 0,
	0, 0, 132, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 116, 117, 118, 119, 121, 122, 123,
	114, 124, 129, 127, 128, 125, 126, 0, 0, 130,
	133, 0, 0, 113, 289, 0, 0, 0, 0, 0,
	0, 0, 0, 132, 115, 0, 0, 0, 0, 0,
	0, 0, 0, 120, 116, 117, 118, 119, 121, 122,
	123, 229, 124, 129, 127, 128, 125, 126, 0, 0,
	130, 133, 0, 0, 0, 131, 0, 0, 0, 0,
	0, 0, 0, 0, 132, 237, 0, 0, 0, 236,
	0, 0, 0, 0, 120, 116, 117, 118, 119, 121,
	122, 123, 229, 124, 129, 127, 128, 125, 126, 0,
	0, 130, 133, 0, 0, 0, 131, 0, 0, 0,
	0, 0, 0, 0, 0, 132, 235, 0, 0, 0,
	234, 0, 0, 0, 0, 120, 116, 117, 118, 119,
	121, 122, 123, 114, 124, 129, 127, 128, 125, 126,
	0, 0, 130, 133, 0, 0, 113, 157, 0, 0,
	0, 0, 0, 0, 0, 0, 132, 115, 0, 0,
	0, 0, 0, 0, 0, 0, 120, 116, 117, 118,
	119, 121, 122, 123, 114, 124, 129, 127, 128, 125,
	126, 0, 0, 130, 133, 0, 0, 113, 131, 0,
	0, 0, 0, 0, 0, 0, 0, 132, 115, 0,
	0, 0, 0, 0, 0, 0, 0, 120, 116, 117,
	118, 119, 121, 122, 123, 229, 124, 129, 127, 128,
	125, 126, 0, 0, 130, 133, 0, 0, 0, 131,
	0, 0, 0, 0, 415, 0, 0, 0, 132, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 120, 116,
	117, 118, 119, 121, 122, 123, 229, 124, 129, 127,
	128, 125, 126, 0, 0, 130, 133, 0, 0, 0,
	131, 0, 0, 0, 0, 414, 0, 0, 0, 132,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 120,
	116, 117, 118, 119, 121, 122, 123, 229, 124, 129,
	127, 128, 125, 126, 0, 0, 130, 133, 0, 0,
	0, 131, 0, 0, 0, 0, 404, 0, 0, 0,
	132, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	120, 116, 117, 118, 119, 121, 122, 123, 229, 124,
	129, 127, 128, 125, 126, 0, 0, 130, 133, 0,
	0, 0, 131, 0, 0, 0, 0, 401, 0, 0,
	0, 132, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 120, 116, 117, 118, 119, 121, 122, 123, 229,
	124, 129, 127, 128, 125, 126, 0, 0, 130, 133,
	0, 0, 0, 131, 0, 0, 0, 0, 398, 0,
	0, 0, 132, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 116, 117, 118, 119, 121, 122, 123,
	229, 124, 129, 127, 128, 125, 126, 0, 0, 130,
	133, 0, 0, 0, 131, 0, 0, 0, 0, 397,
	0, 0, 0, 132, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 120, 116, 117, 118, 119, 121, 122,
	123, 229, 124, 129, 127, 128, 125, 126, 132, 0,
	130, 133, 0, 0, 0, 131, 0, 377, 0, 116,
	117, 118, 119, 121, 122, 123, 229, 124, 129, 127,
	128, 125, 126, 0, 120, 130, 133, 0, 0, 0,
	131, 0, 0, 0, 0, 348, 0, 0, 0, 132,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 120,
	116, 117, 118, 119, 121, 122, 123, 229, 124, 129,
	127, 128, 125, 126, 0, 0, 130, 133, 0, 0,
	0, 131, 0, 0, 0, 0, 0, 0, 0, 0,
	132, 0, 347, 0, 0, 0, 0, 0, 0, 0,
	120, 116, 117, 118, 119, 121, 122, 123, 229, 124,
	129, 127, 128, 125, 126, 0, 0, 130, 133, 0,
	0, 0, 131, 0, 0, 0, 0, 0, 0, 0,
	0, 132, 0, 346, 0, 0, 0, 0, 0, 0,
	0, 120, 116, 117, 118, 119, 121, 122, 123, 229,
	124, 129, 127, 128, 125, 126, 0, 0, 130, 133,
	0, 0, 0, 131, 0, 0, 0, 0, 345, 0,
	0, 0, 132, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 116, 117, 118, 119, 121, 122, 123,
	229, 124, 129, 127, 128, 125, 126, 132, 275, 130,
	133, 0, 0, 0, 131, 0, 0, 292, 116, 117,
	118, 119, 121, 122, 123, 229, 124, 129, 127, 128,
	125, 126, 132, 120, 130, 133, 0, 0, 0, 131,
	0, 0, 0, 116, 117, 118, 119, 121, 122, 123,
	229, 124, 129, 127, 128, 125, 126, 0, 120, 130,
	133, 0, 0, 0, 131, 0, 0, 0, 0, 0,
	0, 0, 0, 132, 0, 233, 0, 0, 0, 0,
	0, 0, 0, 120, 116, 117, 118, 119, 121, 122,
	123, 229, 124, 129, 127, 128, 125, 126, 0, 0,
	130, 133, 0, 0, 0, 131, 0, 0, 0, 0,
	0, 0, 0, 0, 132, 0, 230, 0, 0, 0,
	0, 0, 0, 0, 120, 116, 117, 118, 119, 121,
	122, 123, 229, 124, 129, 127, 128, 125, 126, 132,
	0, 130, 133, 0, 0, 0, 131, 0, 0, 0,
	116, 117, 118, 119, 121, 334, 123, 229, 124, 129,
	127, 128, 125, 126, 132, 120, 130, 133, 0, 0,
	0, 131, 0, 0, 0, 116, 117, 118, 119, 121,
	271, 123, 229, 124, 129, 127, 128, 125, 126, 132,
	120, 130, 133, 0, 0, 0, 131, 0, 0, 0,
	116, 117, 118, 119, 121, 122, 0, 229, 124, 129,
	127, 128, 125, 126, 132, 120, 130, 133, 0, 0,
	0, 131, 0, 0, 0, 116, 117, 118, 119, 121,
	0, 0, 229, 124, 129, 127, 128, 125, 126, 0,
	120, 130, 133, 54, 0, 76, 131, 0, 0, 71,
	72, 73, 74, 75, 59, 67, 0, 56, 224, 0,
	0, 0, 0, 55, 0, 120, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 63, 0, 65, 66, 0, 54, 64,
	76, 0, 0, 0, 71, 72, 73, 74, 75, 59,
	67, 0, 56, 82, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 140, 64, 76, 0, 0, 208, 71,
	72, 73, 207, 75, 59, 67, 0, 56, 0, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 63, 0, 65, 66, 0, 54, 64,
	76, 151, 0, 0, 71, 72, 73, 74, 75, 59,
	67, 0, 56, 0, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 140, 64, 76, 0, 0, 0, 71,
	72, 73, 74, 75, 59, 67, 0, 56, 0, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 63, 0, 65, 66, 0, 54, 64,
	76, 0, 0, 0, 71, 72, 73, 74, 75, 59,
	67, 0, 56, 0, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 140, 64, 76, 0, 0, 0, 71,
	72, 73, 74, 75, 147, 67, 0, 56, 0, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 63, 0, 65, 66, 0, 140, 64,
	76, 0, 0, 0, 71, 72, 73, 74, 75, 146,
	67, 0, 56, 0, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 0, 64,
}
var yyPact = []int{

	464, -1000, -1000, 413, -1000, -1000, -1000, -1000, -1000, -1000,
	353, 304, 351, 309, 313, 13, 142, -1000, -1000, 138,
	293, 297, 352, 177, 313, 108, 282, 1651, 1471, -1000,
	-1000, -1000, -1000, 171, 41, 249, -1000, -31, 170, -1000,
	287, 203, 1651, 345, 339, 282, -1000, 166, 319, 294,
	-1000, 687, -1000, -1000, 216, 1606, 1606, -1000, -1000, 26,
	-1000, 1651, -56, 1741, 1696, 1606, 1606, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1561, -1000, -1000, 251,
	-1000, 74, -1000, 646, -34, -1000, 147, -1, 147, 147,
	-1000, -85, -1000, 159, 303, 217, 158, 1606, 1606, 156,
	-35, -1000, 191, -1000, -1000, 246, 291, 99, 82, -1000,
	-41, -1000, 1651, 1606, -58, 1651, 1606, 1606, 1606, 1606,
	1606, 1606, 1606, 1606, 1606, 1606, 1606, 1606, 1606, 1606,
	1606, 155, 1516, 131, 267, -1000, 220, 239, 214, -1000,
	54, -1000, 22, -5, 1606, 1254, -48, -49, 1213, 605,
	564, -1000, 263, 237, 1471, 154, -1000, 96, 147, 358,
	147, 147, 147, 181, 146, -1000, 303, -1000, 224, 189,
	-1000, 1295, 1295, -1000, 153, -1000, 1651, -1000, -1000, 333,
	149, 20, 148, 147, 279, 1345, 1606, 1651, 1606, -1000,
	-26, -26, 104, 104, 104, 104, 1395, 1370, 218, 218,
	218, 218, 218, 218, 218, -1000, 1188, 209, 186, -1000,
	7, -1000, -1000, 275, -1000, 14, 1651, -1000, -14, 1426,
	1426, 233, -1000, -1000, -1000, 523, -1000, -57, 1163, -24,
	1606, 1606, 1606, 1606, 145, 1606, 144, 1606, -1000, 1651,
	-1000, -1000, -1000, -1000, 143, 41, -1000, 308, 112, 110,
	41, 141, 301, 1606, 1606, 41, 118, 301, -1000, -1000,
	221, 262, -42, -1000, 115, -46, 1651, -51, -1000, -1000,
	1651, 1606, 1320, -1000, 218, -1000, 196, 256, -1000, -1000,
	-1000, -1000, 323, -1000, -1000, -1000, -15, -17, 1426, 94,
	-62, 1606, 1606, 1122, 1081, 1040, 999, -70, 482, -72,
	441, -1000, 41, -1000, 116, -6, -1000, 41, 41, 301,
	111, 41, 301, 106, -1000, 301, 41, 1295, 1295, -1000,
	301, 41, 254, -1000, -1000, 83, -1000, -1000, -1000, 49,
	-18, 29, -1000, 1395, 1606, 253, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 1295, 974, -1000, 1606, 1606, -1000, 1606,
	1606, -1000, 1606, 1606, -1000, -1000, -6, -1000, 41, -1000,
	-1000, 41, 301, -1000, 41, 301, 41, -1000, 41, -1000,
	-1000, -1000, 332, 330, -2, 1395, -1000, 1606, 933, 892,
	400, 851, 359, 810, -1000, 41, -1000, -1000, 41, -1000,
	41, -1000, -1000, 115, 115, 1651, -1000, -1000, -1000, 1606,
	-1000, -1000, 1606, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-19, 769, 728, 325, -1000, -1000, 115, -1000,
}
var yyPgo = []int{

	0, 484, 383, 13, 482, 480, 479, 478, 1, 3,
	475, 474, 458, 457, 316, 453, 384, 329, 451, 450,
	20, 447, 444, 443, 10, 441, 434, 0, 8, 433,
	2, 16, 417, 15, 7, 11, 415, 414, 413, 9,
	244, 412, 411, 410, 4, 401, 5, 398, 397, 394,
	390, 386, 6, 385,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 22, 23, 24, 24, 25, 25, 25,
	25, 26, 26, 15, 15, 15, 18, 18, 28, 28,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 34, 34, 32, 32, 32, 29, 29, 29,
	29, 29, 29, 33, 33, 16, 16, 12, 12, 35,
	35, 36, 36, 36, 13, 13, 13, 37, 38, 20,
	20, 20, 20, 20, 20, 39, 39, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 40, 40, 40, 41, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 44,
	44, 45, 45, 31, 31, 31, 31, 31, 31, 46,
	46, 47, 47, 48, 48, 43, 43, 43, 43, 43,
	43, 43, 49, 49, 50, 50, 52, 52, 53, 51,
	51, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	2, 1, 3, 0, 2, 5, 2, 5, 1, 2,
	2, 4, 3, 3, 5, 4, 3, 5, 4, 4,
	6, 5, 4, 5, 6, 5, 6, 7, 3, 5,
	4, 4, 6, 5, 4, 5, 5, 6, 6, 7,
	3, 5, 4, 4, 6, 5, 4, 5, 5, 6,
	6, 7, 2, 2, 1, 1, 2, 1, 2, 3,
	2, 4, 3, 2, 2, 0, 2, 0, 3, 1,
	3, 1, 2, 2, 0, 1, 2, 2, 2, 1,
	5, 6, 3, 4, 1, 3, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 6, 5, 5, 3, 4,
	3, 4, 3, 4, 1, 2, 2, 1, 1, 1,
	1, 3, 5, 5, 7, 7, 5, 9, 7, 7,
	5, 9, 7, 7, 5, 3, 4, 5, 5, 3,
	5, 0, 2, 1, 4, 6, 5, 5, 3, 1,
	3, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 3, 1, 3, 3, 2,
	3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 23, -3, -4, -5, -10, -6, -7,
	25, -11, 24, -14, -18, -21, 35, 33, -2, 28,
	-12, 39, 26, 28, -15, 35, -16, 36, -22, 88,
	31, 32, -28, 52, -29, -31, 58, 58, 52, -13,
	-37, 43, 40, 28, 58, -16, -28, 52, -17, 45,
	-20, -27, -39, -40, 47, 67, 61, -41, -42, 58,
	-43, 74, 82, 87, 93, 89, 90, 59, -49, -50,
	-51, 53, 54, 55, 56, 57, 49, -23, -24, -25,
	-26, -20, 62, -27, 58, -30, 94, -32, 18, 96,
	97, 98, -33, 34, 58, 49, 81, 37, 38, 81,
	58, -38, 44, 56, -35, -36, -20, 29, 29, -17,
	58, -14, 40, 80, 67, 91, 60, 61, 62, 63,
	100, 64, 65, 66, 68, 72, 73, 70, 71, 69,
	76, 81, 49, 77, -3, 48, -52, -53, 59, -40,
	47, -40, 74, -20, 83, -27, 58, 58, -27, -27,
	-27, 50, -9, -20, 51, 34, 58, 81, 81, -31,
	94, 18, 96, -31, -31, 99, 58, -33, 56, 52,
	58, -27, -27, 58, 81, 56, 51, 41, 42, 58,
	52, 58, 52, 81, -9, -27, 80, 91, 76, -20,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, 58, -27, 56, 52, 55,
	67, 79, 78, 48, 48, 51, 52, 75, -46, 31,
	32, -47, -48, -20, 62, -27, 75, -44, -27, 67,
	92, 91, 91, 92, 95, 91, 95, 91, 50, 51,
	-24, 58, 62, -28, 34, 58, -30, -31, -31, -31,
	-34, 34, 58, 37, 38, -34, 34, 58, -33, 50,
	52, 56, 58, -35, 30, 58, 74, 58, -28, -19,
	46, 65, -27, -20, -27, 50, 52, 56, 55, 79,
	78, -39, 47, -52, -20, 75, -46, -46, 51, 81,
	-45, 85, 84, -27, -27, -27, -27, 58, -27, 58,
	-27, -9, 58, -30, 34, 58, -30, -33, -34, 58,
	34, -34, 58, 34, -30, 58, -34, -27, -27, -30,
	58, -34, 56, 50, 50, 81, -8, 27, 58, 81,
	-9, 81, -20, -27, 65, 56, 50, 50, 75, 75,
	-46, 62, 86, -27, -27, 86, 92, 92, 86, 91,
	83, 86, 91, 83, 86, -30, 58, -30, -33, -30,
	-30, -34, 58, -30, -34, 58, -34, -30, -34, -30,
	50, 58, 58, 75, 58, -27, 50, 83, -27, -27,
	-27, -27, -27, -27, -30, -33, -30, -30, -34, -30,
	-34, -30, -30, 30, 30, 74, -44, 86, 86, 83,
	86, 86, 83, 86, 86, -30, -30, -30, -8, -8,
	-9, -27, -27, 75, 86, 86, 30, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 107, 0, 43, 105, 30, 0, 29, 2, 0,
	114, 0, 0, 0, 105, 0, 24, 0, 0, 31,
	32, 33, 46, 0, 48, 97, 183, 0, 0, 21,
	115, 0, 0, 0, 0, 24, 44, 0, 0, 0,
	106, 119, 124, 154, 0, 0, 0, 157, 158, 159,
	160, 0, 0, 0, 0, 0, 0, 195, 196, 197,
	198, 199, 200, 201, 202, 203, 0, 28, 34, 35,
	37, 38, 41, 119, 0, 49, 0, 0, 0, 0,
	94, 95, 98, 0, 100, 0, 0, 0, 0, 0,
	0, 116, 0, 117, 108, 109, 111, 0, 0, 22,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 204, 0, 206, 0, 155,
	0, 156, 0, 0, 0, 0, 159, 159, 0, 0,
	0, 209, 0, 211, 0, 0, 40, 0, 0, 50,
	0, 0, 0, 0, 0, 96, 99, 102, 0, 0,
	188, 103, 104, 18, 0, 118, 0, 112, 113, 8,
	0, 0, 0, 0, 26, 0, 0, 0, 0, 122,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 143, 0, 202, 0, 148,
	0, 150, 152, 125, 205, 0, 0, 175, 0, 0,
	0, 189, 191, 192, 193, 119, 161, 181, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 210, 0,
	36, 39, 42, 47, 0, 52, 53, 56, 0, 0,
	68, 0, 0, 0, 0, 80, 0, 0, 101, 184,
	0, 0, 0, 110, 0, 0, 0, 0, 45, 25,
	0, 0, 0, 123, 142, 144, 0, 0, 149, 151,
	153, 126, 0, 207, 208, 176, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 212, 51, 55, 0, 58, 59, 62, 74, 0,
	0, 86, 0, 0, 71, 0, 70, 92, 93, 83,
	0, 82, 0, 186, 187, 0, 10, 16, 17, 0,
	0, 0, 27, -2, 0, 0, 146, 147, 177, 178,
	190, 194, 162, 182, 179, 163, 0, 0, 166, 0,
	0, 170, 0, 0, 174, 54, 57, 61, 63, 65,
	75, 76, 0, 87, 88, 0, 69, 73, 81, 85,
	185, 19, 9, 12, 0, -2, 145, 0, 0, 0,
	0, 0, 0, 0, 60, 64, 66, 77, 78, 89,
	90, 72, 84, 0, 0, 0, 180, 164, 165, 0,
	169, 168, 0, 173, 172, 67, 79, 91, 11, 14,
	0, 0, 0, 13, 167, 171, 0, 15,
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
	92, 93, 94, 95, 96, 97, 98, 99, 100,
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
		//line n1ql.y:242
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 23:
		//line n1ql.y:246
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 24:
		//line n1ql.y:253
		{
	}
	case 25:
		//line n1ql.y:256
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
		//line n1ql.y:268
		{
	}
	case 27:
		//line n1ql.y:271
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
		//line n1ql.y:284
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 29:
		//line n1ql.y:290
		{
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 30:
		//line n1ql.y:296
		{
	}
	case 31:
		//line n1ql.y:299
		{
	/* empty */
}
	case 32:
		//line n1ql.y:302
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 33:
		//line n1ql.y:312
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 34:
		//line n1ql.y:324
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
	case 35:
		//line n1ql.y:338
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 36:
		//line n1ql.y:343
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
	case 37:
		//line n1ql.y:356
		{
		logDebugGrammar("RESULT STAR")
	}
	case 38:
		//line n1ql.y:360
		{
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 39:
		//line n1ql.y:367
		{
		logDebugGrammar("RESULT EXPR AS ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 40:
		//line n1ql.y:374
		{
		logDebugGrammar("RESULT EXPR ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 41:
		//line n1ql.y:383
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 42:
		//line n1ql.y:389
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 43:
		//line n1ql.y:398
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 44:
		//line n1ql.y:402
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
	case 45:
		//line n1ql.y:413
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
	case 46:
		//line n1ql.y:427
		{
		logDebugGrammar("SELECT FROM - DATASOURCE ")
		from := parsingStack.Pop().(*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 47:
		//line n1ql.y:438
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
	case 48:
		//line n1ql.y:452
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT UNNEST")
	}
	case 49:
		//line n1ql.y:456
		{
		logDebugGrammar("FROM DATASOURCE WITH UNNEST")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 50:
		//line n1ql.y:467
		{
	    logDebugGrammar("UNNEST")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:""})
	}
	case 51:
		//line n1ql.y:474
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 52:
		//line n1ql.y:481
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 53:
		//line n1ql.y:488
		{
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: "", Over:rest})
	}
	case 54:
		//line n1ql.y:495
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 55:
		//line n1ql.y:502
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 56:
		//line n1ql.y:509
		{
	    logDebugGrammar("UNNEST")
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type})
	}
	case 57:
		//line n1ql.y:517
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, Type: Type,As:yyS[yypt-0].s})
	}
	case 58:
		//line n1ql.y:525
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, Type: Type,As:yyS[yypt-0].s})
	}
	case 59:
		//line n1ql.y:533
		{
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: "", Over:rest})
	}
	case 60:
		//line n1ql.y:541
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-1].s, Over:rest})
	}
	case 61:
		//line n1ql.y:549
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-1].s, Over:rest})
	}
	case 62:
		//line n1ql.y:557
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type, Keys: key_expr}) 
	}
	case 63:
		//line n1ql.y:565
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Type: Type, Keys: key_expr}) 
	}
	case 64:
		//line n1ql.y:573
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Type: Type, Keys: key_expr}) 
	}
	case 65:
		//line n1ql.y:581
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type, Keys: key_expr, Over:rest}) 
	}
	case 66:
		//line n1ql.y:590
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Type: Type, Keys: key_expr, Over:rest}) 
	}
	case 67:
		//line n1ql.y:599
		{
	    logDebugGrammar("UNNEST KEY_EXPR")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Type: Type, Keys: key_expr, Over:rest}) 
	}
	case 68:
		//line n1ql.y:608
		{
	    logDebugGrammar("JOIN KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr})
	}
	case 69:
		//line n1ql.y:615
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 70:
		//line n1ql.y:622
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 71:
		//line n1ql.y:629
		{
	    logDebugGrammar("JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Keys: key_expr, Over: rest})
	}
	case 72:
		//line n1ql.y:637
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 73:
		//line n1ql.y:645
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 74:
		//line n1ql.y:653
		{
	    logDebugGrammar("TYPE JOIN KEY")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type, Keys: key_expr})
	
	}
	case 75:
		//line n1ql.y:662
		{
	    logDebugGrammar("TYPE JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Type: Type, Keys: key_expr, Over: rest})
	}
	case 76:
		//line n1ql.y:671
		{
	    logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Type:Type, Keys: key_expr})
	
	}
	case 77:
		//line n1ql.y:680
		{
	    logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Type:Type, Keys: key_expr, Over: rest})
	}
	case 78:
		//line n1ql.y:689
		{
	    logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Type:Type, Keys: key_expr})
	}
	case 79:
		//line n1ql.y:697
		{
	    logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Type:Type, Keys: key_expr, Over: rest})
	}
	case 80:
		//line n1ql.y:706
		{
	    logDebugGrammar("JOIN KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:"", Keys: key_expr})
	}
	case 81:
		//line n1ql.y:713
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 82:
		//line n1ql.y:720
		{
	    logDebugGrammar("JOIN AS KEY") 
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:yyS[yypt-1].s, Keys: key_expr})
	}
	case 83:
		//line n1ql.y:727
		{
	    logDebugGrammar("JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:"", Keys: key_expr, Over: rest})
	}
	case 84:
		//line n1ql.y:735
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 85:
		//line n1ql.y:743
		{
	    logDebugGrammar("JOIN AS KEY NESTED") 
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:yyS[yypt-2].s, Keys: key_expr, Over: rest})
	}
	case 86:
		//line n1ql.y:751
		{
	    logDebugGrammar("TYPE JOIN KEY")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As:"", Type: Type, Keys: key_expr})
	
	}
	case 87:
		//line n1ql.y:760
		{
	    logDebugGrammar("TYPE JOIN KEY NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:"", Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
	}
	case 88:
		//line n1ql.y:769
		{
	    logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Oper: "NEST", Type:Type, Keys: key_expr})
	
	}
	case 89:
		//line n1ql.y:778
		{
	    logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Oper: "NEST", Type:Type, Keys: key_expr, Over: rest})
	}
	case 90:
		//line n1ql.y:787
		{
	    logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s, Oper: "NEST", Type:Type, Keys: key_expr})
	}
	case 91:
		//line n1ql.y:795
		{
	    logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
	    rest := parsingStack.Pop().(*ast.From)
	    key_expr := parsingStack.Pop().(*ast.KeyExpression)
	    proj := parsingStack.Pop().(ast.Expression)
	    Type := parsingStack.Pop().(string)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-2].s, Oper: "NEST", Type:Type, Keys: key_expr, Over: rest})
	}
	case 92:
		//line n1ql.y:806
		{
	        logDebugGrammar("FROM JOIN DATASOURCE with KEY")
	        key := parsingStack.Pop().(ast.Expression)
	        key_expr := ast.NewKeyExpression(key, "KEY")
	        parsingStack.Push(key_expr)
	}
	case 93:
		//line n1ql.y:813
		{
	        logDebugGrammar("FROM DATASOURCE with KEYS")
	        keys := parsingStack.Pop().(ast.Expression)
	        keys_expr := ast.NewKeyExpression(keys, "KEYS")
	        parsingStack.Push(keys_expr)
	
	}
	case 94:
		//line n1ql.y:822
		{
	    logDebugGrammar("INNER")
	    parsingStack.Push("INNER")
	}
	case 95:
		//line n1ql.y:827
		{
	    logDebugGrammar("OUTER")
	    parsingStack.Push("LEFT")
	}
	case 96:
		//line n1ql.y:832
		{
	    logDebugGrammar("LEFT OUTER")
	    parsingStack.Push("LEFT")
	}
	case 97:
		//line n1ql.y:839
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 98:
		//line n1ql.y:845
		{
	    logDebugGrammar("FROM KEY(S) DATASOURCE")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection:proj})
	}
	case 99:
		//line n1ql.y:851
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 100:
		//line n1ql.y:858
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 101:
		//line n1ql.y:865
		{
	        logDebugGrammar("FROM DATASOURCE AS ID KEY(S)")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})
	
	}
	case 102:
		//line n1ql.y:872
		{
	        logDebugGrammar("FROM DATASOURCE ID KEY(s)")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})
	
	}
	case 103:
		//line n1ql.y:881
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
	case 104:
		//line n1ql.y:892
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
	case 105:
		//line n1ql.y:906
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 106:
		//line n1ql.y:910
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
	case 108:
		//line n1ql.y:924
		{
	
	}
	case 109:
		//line n1ql.y:930
		{
	
	}
	case 110:
		//line n1ql.y:934
		{
	
	}
	case 111:
		//line n1ql.y:939
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
	case 112:
		//line n1ql.y:950
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
	case 113:
		//line n1ql.y:961
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
	case 114:
		//line n1ql.y:973
		{
	
	}
	case 115:
		//line n1ql.y:977
		{
	
	}
	case 116:
		//line n1ql.y:981
		{
	
	}
	case 117:
		//line n1ql.y:987
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
	case 118:
		//line n1ql.y:1001
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
	case 119:
		//line n1ql.y:1018
		{
		logDebugGrammar("EXPRESSION")
	}
	case 120:
		//line n1ql.y:1022
		{
	    logDebugGrammar(" BETWEEN EXPRESSION")
	    high := parsingStack.Pop()
	    low := parsingStack.Pop()
	    element := parsingStack.Pop()
	    leftExpression := ast.NewGreaterThanOrEqualOperator(element.(ast.Expression), low.(ast.Expression))
	    rightExpression := ast.NewLessThanOrEqualOperator(element.(ast.Expression), high.(ast.Expression))
	    thisExpression := ast.NewAndOperator(ast.ExpressionList{leftExpression, rightExpression})
	    parsingStack.Push(thisExpression)
	}
	case 121:
		//line n1ql.y:1033
		{
	    logDebugGrammar(" BETWEEN EXPRESSION")
	    high := parsingStack.Pop()
	    low := parsingStack.Pop()
	    element := parsingStack.Pop()
	    leftExpression := ast.NewLessThanOperator(element.(ast.Expression), low.(ast.Expression))
	    rightExpression := ast.NewGreaterThanOperator(element.(ast.Expression), high.(ast.Expression))
	    thisExpression := ast.NewOrOperator(ast.ExpressionList{leftExpression, rightExpression})
	    parsingStack.Push(thisExpression)
	}
	case 122:
		//line n1ql.y:1044
		{
	    logDebugGrammar(" IN expression ")
	    right := parsingStack.Pop()
	    left := parsingStack.Pop()
	    thisExpression := ast.NewInOperator(left.(ast.Expression), right.(ast.Expression))
	    parsingStack.Push(thisExpression)
	}
	case 123:
		//line n1ql.y:1052
		{
	    logDebugGrammar(" IN expression ")
	    right := parsingStack.Pop()
	    left := parsingStack.Pop()
	    thisExpression := ast.NewNotInOperator(left.(ast.Expression), right.(ast.Expression))
	    parsingStack.Push(thisExpression)
	}
	case 124:
		//line n1ql.y:1060
		{
	}
	case 125:
		//line n1ql.y:1064
		{
	    logDebugGrammar("sub-query EXPRESSION")
	    
	}
	case 126:
		//line n1ql.y:1069
		{
	    logDebugGrammar("sub-query NESTED EXPRESSION")
	}
	case 127:
		//line n1ql.y:1075
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 128:
		//line n1ql.y:1083
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 129:
		//line n1ql.y:1091
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 130:
		//line n1ql.y:1099
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 131:
		//line n1ql.y:1107
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 132:
		//line n1ql.y:1115
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 133:
		//line n1ql.y:1123
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 134:
		//line n1ql.y:1131
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 135:
		//line n1ql.y:1149
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 136:
		//line n1ql.y:1157
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 137:
		//line n1ql.y:1165
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 138:
		//line n1ql.y:1173
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 139:
		//line n1ql.y:1181
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 140:
		//line n1ql.y:1189
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 141:
		//line n1ql.y:1197
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 142:
		//line n1ql.y:1205
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
	        parsingStack.Push(thisExpression)
	
	}
	case 143:
		//line n1ql.y:1214
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 144:
		//line n1ql.y:1222
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 145:
		//line n1ql.y:1230
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 146:
		//line n1ql.y:1237
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 147:
		//line n1ql.y:1245
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1252
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 149:
		//line n1ql.y:1259
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 150:
		//line n1ql.y:1266
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 151:
		//line n1ql.y:1273
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 152:
		//line n1ql.y:1280
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 153:
		//line n1ql.y:1287
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 154:
		//line n1ql.y:1294
		{
	
	}
	case 155:
		//line n1ql.y:1300
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 156:
		//line n1ql.y:1307
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 157:
		//line n1ql.y:1314
		{
	
	}
	case 158:
		//line n1ql.y:1319
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 159:
		//line n1ql.y:1325
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 160:
		//line n1ql.y:1331
		{
		logDebugGrammar("LITERAL")
	}
	case 161:
		//line n1ql.y:1335
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 162:
		//line n1ql.y:1339
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
	case 163:
		//line n1ql.y:1356
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 164:
		//line n1ql.y:1364
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 165:
		//line n1ql.y:1372
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 166:
		//line n1ql.y:1380
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 167:
		//line n1ql.y:1388
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 168:
		//line n1ql.y:1397
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 169:
		//line n1ql.y:1406
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 170:
		//line n1ql.y:1414
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 171:
		//line n1ql.y:1422
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 172:
		//line n1ql.y:1431
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 173:
		//line n1ql.y:1440
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 174:
		//line n1ql.y:1448
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 175:
		//line n1ql.y:1456
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 176:
		//line n1ql.y:1462
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 177:
		//line n1ql.y:1469
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 178:
		//line n1ql.y:1477
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 179:
		//line n1ql.y:1486
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 180:
		//line n1ql.y:1494
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
	case 181:
		//line n1ql.y:1508
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 182:
		//line n1ql.y:1512
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 183:
		//line n1ql.y:1518
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 184:
		//line n1ql.y:1524
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 185:
		//line n1ql.y:1531
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 186:
		//line n1ql.y:1538
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 187:
		//line n1ql.y:1546
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 188:
		//line n1ql.y:1553
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 189:
		//line n1ql.y:1564
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 190:
		//line n1ql.y:1569
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
	case 191:
		//line n1ql.y:1583
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 192:
		//line n1ql.y:1587
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 193:
		//line n1ql.y:1596
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 194:
		//line n1ql.y:1602
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 195:
		//line n1ql.y:1612
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 196:
		//line n1ql.y:1618
		{
		logDebugGrammar("NUMBER")
	}
	case 197:
		//line n1ql.y:1622
		{
		logDebugGrammar("OBJECT")
	}
	case 198:
		//line n1ql.y:1626
		{
		logDebugGrammar("ARRAY")
	}
	case 199:
		//line n1ql.y:1630
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 200:
		//line n1ql.y:1636
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 201:
		//line n1ql.y:1642
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 202:
		//line n1ql.y:1650
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 203:
		//line n1ql.y:1656
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 204:
		//line n1ql.y:1664
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 205:
		//line n1ql.y:1670
		{
		logDebugGrammar("OBJECT")
	}
	case 206:
		//line n1ql.y:1676
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 207:
		//line n1ql.y:1680
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 208:
		//line n1ql.y:1692
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 209:
		//line n1ql.y:1702
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 210:
		//line n1ql.y:1708
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 211:
		//line n1ql.y:1717
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 212:
		//line n1ql.y:1724
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
