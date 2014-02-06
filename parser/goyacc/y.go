
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
	-1, 336,
	65, 133,
	66, 133,
	-2, 120,
	-1, 379,
	65, 133,
	66, 133,
	-2, 121,
}

const yyNprod = 214
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1881

var yyAct = []int{

	51, 329, 85, 153, 292, 228, 219, 252, 136, 52,
	32, 104, 166, 35, 4, 92, 78, 88, 189, 356,
	142, 142, 187, 353, 382, 154, 345, 189, 132, 83,
	162, 30, 31, 188, 293, 334, 46, 234, 233, 116,
	117, 118, 119, 121, 122, 123, 114, 124, 129, 127,
	128, 125, 126, 50, 81, 130, 133, 332, 328, 113,
	291, 184, 132, 145, 146, 149, 150, 151, 106, 134,
	115, 175, 159, 116, 117, 118, 119, 121, 122, 120,
	230, 124, 129, 127, 128, 125, 126, 143, 29, 130,
	133, 99, 418, 86, 131, 89, 90, 91, 172, 173,
	160, 377, 164, 165, 342, 341, 161, 287, 163, 227,
	168, 400, 280, 120, 186, 268, 185, 191, 192, 193,
	194, 195, 196, 197, 198, 199, 200, 201, 202, 203,
	204, 205, 142, 207, 132, 282, 281, 132, 206, 135,
	206, 190, 344, 226, 244, 229, 132, 118, 119, 121,
	138, 330, 230, 138, 378, 230, 83, 116, 117, 118,
	119, 121, 133, 248, 230, 133, 131, 156, 224, 131,
	245, 376, 242, 257, 133, 249, 250, 251, 131, 375,
	183, 81, 331, 260, 210, 120, 182, 181, 274, 265,
	276, 157, 47, 180, 38, 270, 211, 120, 36, 88,
	37, 369, 316, 106, 33, 255, 256, 213, 212, 366,
	36, 360, 323, 318, 275, 305, 302, 95, 97, 98,
	300, 226, 226, 36, 283, 285, 315, 288, 289, 269,
	267, 264, 229, 296, 297, 298, 299, 295, 301, 278,
	303, 243, 206, 286, 174, 304, 224, 224, 171, 96,
	306, 167, 309, 110, 100, 317, 320, 321, 311, 314,
	322, 17, 319, 16, 84, 310, 44, 324, 279, 263,
	176, 103, 333, 53, 336, 86, 135, 89, 90, 91,
	339, 313, 326, 217, 255, 256, 338, 138, 325, 170,
	261, 226, 262, 169, 346, 347, 95, 343, 335, 134,
	348, 290, 241, 216, 177, 312, 155, 380, 359, 374,
	340, 361, 327, 363, 364, 240, 224, 367, 13, 215,
	365, 132, 371, 368, 362, 214, 370, 373, 96, 139,
	141, 372, 116, 117, 118, 119, 121, 284, 379, 230,
	124, 129, 127, 128, 125, 126, 272, 49, 130, 133,
	102, 383, 384, 131, 385, 386, 48, 387, 388, 178,
	179, 112, 41, 389, 42, 391, 21, 111, 392, 255,
	256, 394, 120, 396, 393, 397, 390, 395, 27, 88,
	97, 98, 229, 12, 10, 25, 17, 401, 16, 17,
	421, 108, 17, 410, 16, 246, 411, 399, 412, 398,
	413, 414, 109, 266, 415, 416, 107, 22, 417, 23,
	95, 43, 19, 132, 26, 137, 2, 70, 69, 247,
	18, 68, 223, 422, 116, 117, 118, 119, 121, 122,
	123, 230, 124, 129, 127, 128, 125, 126, 222, 45,
	130, 133, 96, 60, 58, 131, 57, 407, 101, 40,
	408, 3, 12, 10, 132, 86, 105, 89, 90, 91,
	87, 17, 34, 16, 120, 116, 117, 118, 119, 121,
	122, 123, 230, 124, 129, 127, 128, 125, 126, 80,
	79, 130, 133, 77, 28, 15, 131, 271, 404, 14,
	24, 405, 39, 20, 11, 132, 7, 9, 8, 6,
	5, 1, 0, 0, 0, 120, 116, 117, 118, 119,
	121, 122, 123, 230, 124, 129, 127, 128, 125, 126,
	0, 0, 130, 133, 0, 0, 0, 131, 0, 357,
	0, 0, 358, 0, 0, 0, 132, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 120, 116, 117, 118,
	119, 121, 122, 123, 230, 124, 129, 127, 128, 125,
	126, 0, 0, 130, 133, 0, 0, 0, 131, 0,
	354, 0, 0, 355, 0, 0, 0, 132, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 120, 116, 117,
	118, 119, 121, 122, 123, 230, 124, 129, 127, 128,
	125, 126, 0, 0, 130, 133, 0, 0, 0, 131,
	0, 0, 0, 0, 0, 0, 0, 0, 132, 239,
	0, 0, 0, 238, 0, 0, 0, 0, 120, 116,
	117, 118, 119, 121, 122, 123, 230, 124, 129, 127,
	128, 125, 126, 0, 0, 130, 133, 0, 0, 0,
	131, 0, 0, 0, 0, 0, 0, 0, 0, 132,
	237, 0, 0, 0, 236, 0, 0, 0, 0, 120,
	116, 117, 118, 119, 121, 122, 123, 114, 124, 129,
	127, 128, 125, 126, 0, 0, 130, 133, 0, 0,
	113, 158, 0, 0, 0, 0, 0, 0, 0, 0,
	132, 115, 0, 0, 0, 0, 0, 0, 0, 0,
	120, 116, 117, 118, 119, 121, 122, 123, 114, 124,
	129, 127, 128, 125, 126, 0, 0, 130, 133, 0,
	0, 113, 131, 0, 0, 0, 0, 0, 0, 0,
	0, 132, 115, 0, 0, 0, 0, 0, 0, 0,
	0, 120, 116, 117, 118, 119, 121, 122, 123, 230,
	124, 129, 127, 128, 125, 126, 0, 0, 130, 133,
	0, 0, 0, 131, 0, 0, 0, 0, 420, 0,
	0, 0, 132, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 116, 117, 118, 119, 121, 122, 123,
	230, 124, 129, 127, 128, 125, 126, 0, 0, 130,
	133, 0, 0, 0, 131, 0, 0, 0, 0, 419,
	0, 0, 0, 132, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 120, 116, 117, 118, 119, 121, 122,
	123, 230, 124, 129, 127, 128, 125, 126, 0, 0,
	130, 133, 0, 0, 0, 131, 0, 0, 0, 0,
	409, 0, 0, 0, 132, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 120, 116, 117, 118, 119, 121,
	122, 123, 230, 124, 129, 127, 128, 125, 126, 0,
	0, 130, 133, 0, 0, 0, 131, 0, 0, 0,
	0, 406, 0, 0, 0, 132, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 120, 116, 117, 118, 119,
	121, 122, 123, 230, 124, 129, 127, 128, 125, 126,
	0, 0, 130, 133, 0, 0, 0, 131, 0, 0,
	0, 0, 403, 0, 0, 0, 132, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 120, 116, 117, 118,
	119, 121, 122, 123, 230, 124, 129, 127, 128, 125,
	126, 0, 0, 130, 133, 0, 0, 0, 131, 0,
	0, 0, 0, 402, 0, 0, 0, 132, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 120, 116, 117,
	118, 119, 121, 122, 123, 230, 124, 129, 127, 128,
	125, 126, 132, 0, 130, 133, 0, 0, 0, 131,
	0, 381, 0, 116, 117, 118, 119, 121, 122, 123,
	230, 124, 129, 127, 128, 125, 126, 0, 120, 130,
	133, 0, 0, 0, 131, 0, 0, 0, 0, 352,
	0, 0, 0, 132, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 120, 116, 117, 118, 119, 121, 122,
	123, 230, 124, 129, 127, 128, 125, 126, 0, 0,
	130, 133, 0, 0, 0, 131, 0, 0, 0, 0,
	0, 0, 0, 0, 132, 0, 351, 0, 0, 0,
	0, 0, 0, 0, 120, 116, 117, 118, 119, 121,
	122, 123, 230, 124, 129, 127, 128, 125, 126, 0,
	0, 130, 133, 0, 0, 0, 131, 0, 0, 0,
	0, 0, 0, 0, 0, 132, 0, 350, 0, 0,
	0, 0, 0, 0, 0, 120, 116, 117, 118, 119,
	121, 122, 123, 230, 124, 129, 127, 128, 125, 126,
	0, 0, 130, 133, 0, 0, 0, 131, 0, 0,
	0, 0, 349, 0, 0, 0, 132, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 120, 116, 117, 118,
	119, 121, 122, 123, 230, 124, 129, 127, 128, 125,
	126, 132, 277, 130, 133, 0, 0, 0, 131, 0,
	0, 294, 116, 117, 118, 119, 121, 122, 123, 230,
	124, 129, 127, 128, 125, 126, 132, 120, 130, 133,
	0, 0, 0, 131, 0, 0, 0, 116, 117, 118,
	119, 121, 122, 123, 230, 124, 129, 127, 128, 125,
	126, 0, 120, 130, 133, 0, 0, 0, 131, 0,
	0, 0, 0, 0, 0, 0, 0, 132, 0, 235,
	0, 0, 0, 0, 0, 0, 0, 120, 116, 117,
	118, 119, 121, 122, 123, 230, 124, 129, 127, 128,
	125, 126, 0, 0, 130, 133, 0, 0, 0, 131,
	0, 0, 0, 0, 0, 0, 0, 0, 132, 0,
	232, 0, 0, 0, 0, 0, 0, 0, 120, 116,
	117, 118, 119, 121, 122, 123, 230, 124, 129, 127,
	128, 125, 126, 132, 0, 130, 133, 0, 0, 0,
	131, 0, 231, 0, 116, 117, 118, 119, 121, 122,
	123, 230, 124, 129, 127, 128, 125, 126, 132, 120,
	130, 133, 0, 0, 0, 131, 0, 0, 0, 116,
	117, 118, 119, 121, 337, 123, 230, 124, 129, 127,
	128, 125, 126, 132, 120, 130, 133, 0, 0, 0,
	131, 0, 0, 0, 116, 117, 118, 119, 121, 273,
	123, 230, 124, 129, 127, 128, 125, 126, 0, 120,
	130, 133, 0, 220, 221, 131, 0, 0, 258, 0,
	0, 255, 256, 0, 0, 0, 0, 0, 0, 54,
	0, 76, 0, 95, 120, 71, 72, 73, 74, 75,
	59, 67, 259, 56, 225, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 253, 61, 218, 255, 256,
	0, 0, 0, 93, 62, 96, 97, 98, 0, 63,
	95, 65, 66, 0, 54, 64, 76, 0, 95, 254,
	71, 72, 73, 74, 75, 59, 67, 94, 56, 225,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 0,
	0, 61, 96, 0, 0, 0, 0, 0, 0, 62,
	96, 0, 0, 0, 63, 0, 65, 66, 0, 54,
	64, 76, 0, 0, 0, 71, 72, 73, 74, 75,
	59, 67, 0, 56, 82, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 0, 0, 0, 63,
	0, 65, 66, 88, 140, 64, 76, 0, 0, 209,
	71, 72, 73, 208, 75, 59, 67, 0, 56, 307,
	0, 0, 97, 98, 55, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 95, 0, 0, 0, 0, 62,
	0, 0, 0, 308, 63, 0, 65, 66, 0, 54,
	64, 76, 152, 0, 0, 71, 72, 73, 74, 75,
	59, 67, 0, 56, 0, 0, 96, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 86,
	0, 89, 90, 91, 62, 0, 0, 0, 0, 63,
	0, 65, 66, 0, 140, 64, 76, 0, 0, 0,
	71, 72, 73, 74, 75, 59, 67, 0, 56, 0,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 0, 0, 0, 0, 0, 62,
	144, 0, 0, 0, 63, 0, 65, 66, 0, 140,
	64, 76, 0, 0, 0, 71, 72, 73, 74, 75,
	59, 67, 0, 56, 0, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 0, 0, 0, 63,
	0, 65, 66, 0, 54, 64, 76, 0, 0, 0,
	71, 72, 73, 74, 75, 59, 67, 0, 56, 0,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 63, 0, 65, 66, 0, 140,
	64, 76, 0, 0, 0, 71, 72, 73, 74, 75,
	148, 67, 0, 56, 0, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 0, 0, 0, 63,
	0, 65, 66, 0, 140, 64, 76, 0, 0, 0,
	71, 72, 73, 74, 75, 147, 67, 0, 56, 0,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 0,
	0, 61, 0, 0, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 63, 0, 65, 66, 0, 0,
	64,
}
var yyPact = []int{

	428, -1000, -1000, 359, -1000, -1000, -1000, -1000, -1000, -1000,
	384, 327, 381, 350, 342, 0, 152, -1000, -1000, 142,
	319, 324, 383, 208, 342, 140, 302, 1697, 1472, -1000,
	-1000, -1000, -1000, 206, -1, 1429, -1000, 10, 196, -1000,
	306, 215, 1697, 377, 362, 302, -1000, 195, 356, 321,
	-1000, 651, -1000, -1000, 228, 1652, 1652, -1000, -1000, 58,
	-1000, 1697, 1607, 1787, 1742, 1652, 1652, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1562, -1000, -1000, 255,
	-1000, 133, -1000, 610, -9, -1000, 165, 12, 165, 165,
	-1000, -87, -1000, 193, 343, 237, 190, 1652, 1652, 186,
	-10, -1000, 214, -1000, -1000, 253, 318, 135, 128, -1000,
	-20, -1000, 1697, 1652, -58, 1697, 1652, 1652, 1652, 1652,
	1652, 1652, 1652, 1652, 1652, 1652, 1652, 1652, 1652, 1652,
	1652, 184, 1517, 129, 277, -1000, 271, 252, 231, -1000,
	91, -1000, 1382, 34, 1652, 1259, 1218, -53, -54, 1177,
	569, 528, -1000, 265, 251, 1472, 183, -1000, 82, 165,
	361, 165, 165, 165, 1421, 1384, -1000, 343, -1000, 240,
	213, -1000, 1284, 1284, -1000, 173, -1000, 1697, -1000, -1000,
	373, 172, 41, 171, 165, 300, 1334, 1652, 1697, 1652,
	-1000, 85, 85, 88, 88, 88, 88, 272, 13, 97,
	97, 97, 97, 97, 97, 97, -1000, 1152, 187, 212,
	-1000, 57, -1000, -1000, 290, -1000, 94, 1697, -1000, 32,
	1427, 1427, 250, -1000, -1000, -1000, -21, -1000, -51, 1127,
	-49, 1652, 1652, 1652, 1652, 1652, 162, 1652, 158, 1652,
	-1000, 1697, -1000, -1000, -1000, -1000, 157, -1, -1000, 1545,
	247, 168, -1, 155, 332, 1652, 1652, -1, 154, 332,
	-1000, -1000, 232, 262, -23, -1000, 124, -24, 1697, -46,
	-1000, -1000, 1697, 1652, 1309, -1000, 97, -1000, 230, 260,
	-1000, -1000, -1000, -1000, 353, -1000, -1000, -1000, 30, 29,
	1427, 80, -60, 1652, 1652, -51, 1086, 1045, 1004, 963,
	-68, 487, -72, 446, -1000, -1, -1000, 153, 181, -1000,
	-1, -1, 332, 151, -1, 332, 143, -1000, 332, -1,
	1284, 1284, -1000, 332, -1, 259, -1000, -1000, 121, -1000,
	-1000, -1000, 113, 26, 96, -1000, 272, 1652, 257, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1284, 938, -62, -1000,
	1652, 1652, -1000, 1652, 1652, -1000, 1652, 1652, -1000, -1000,
	181, -1000, -1, -1000, -1000, -1, 332, -1000, -1, 332,
	-1, -1000, -1, -1000, -1000, -1000, 369, 367, 37, 272,
	-1000, 1652, -1000, 897, 856, 405, 815, 364, 774, -1000,
	-1, -1000, -1000, -1, -1000, -1, -1000, -1000, 124, 124,
	1697, -1000, -1000, -1000, 1652, -1000, -1000, 1652, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 17, 733, 692, 360, -1000,
	-1000, 124, -1000,
}
var yyPgo = []int{

	0, 501, 416, 14, 500, 499, 498, 497, 1, 3,
	496, 494, 493, 492, 318, 490, 414, 356, 489, 487,
	25, 485, 484, 483, 16, 480, 479, 0, 10, 462,
	2, 13, 460, 15, 7, 11, 456, 449, 448, 9,
	273, 446, 444, 443, 5, 4, 6, 438, 422, 421,
	418, 417, 8, 415,
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
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	44, 44, 45, 45, 31, 31, 31, 31, 31, 31,
	46, 46, 47, 47, 48, 48, 43, 43, 43, 43,
	43, 43, 43, 49, 49, 50, 50, 52, 52, 53,
	51, 51, 9, 9,
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
	1, 3, 5, 6, 5, 7, 7, 5, 9, 7,
	7, 5, 9, 7, 7, 5, 3, 4, 5, 5,
	3, 5, 0, 2, 1, 4, 6, 5, 5, 3,
	1, 3, 1, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 1, 3, 3,
	2, 3, 1, 3,
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
	47, -40, 74, -20, 83, -27, -27, 58, 58, -27,
	-27, -27, 50, -9, -20, 51, 34, 58, 81, 81,
	-31, 94, 18, 96, -31, -31, 99, 58, -33, 56,
	52, 58, -27, -27, 58, 81, 56, 51, 41, 42,
	58, 52, 58, 52, 81, -9, -27, 80, 91, 76,
	-20, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, 58, -27, 56, 52,
	55, 67, 79, 78, 48, 48, 51, 52, 75, -46,
	31, 32, -47, -48, -20, 62, -27, 75, -44, -27,
	67, 83, 92, 91, 91, 92, 95, 91, 95, 91,
	50, 51, -24, 58, 62, -28, 34, 58, -30, -31,
	-31, -31, -34, 34, 58, 37, 38, -34, 34, 58,
	-33, 50, 52, 56, 58, -35, 30, 58, 74, 58,
	-28, -19, 46, 65, -27, -20, -27, 50, 52, 56,
	55, 79, 78, -39, 47, -52, -20, 75, -46, -46,
	51, 81, -45, 85, 84, -44, -27, -27, -27, -27,
	58, -27, 58, -27, -9, 58, -30, 34, 58, -30,
	-33, -34, 58, 34, -34, 58, 34, -30, 58, -34,
	-27, -27, -30, 58, -34, 56, 50, 50, 81, -8,
	27, 58, 81, -9, 81, -20, -27, 65, 56, 50,
	50, 75, 75, -46, 62, 86, -27, -27, -45, 86,
	92, 92, 86, 91, 83, 86, 91, 83, 86, -30,
	58, -30, -33, -30, -30, -34, 58, -30, -34, 58,
	-34, -30, -34, -30, 50, 58, 58, 75, 58, -27,
	50, 83, 86, -27, -27, -27, -27, -27, -27, -30,
	-33, -30, -30, -34, -30, -34, -30, -30, 30, 30,
	74, -44, 86, 86, 83, 86, 86, 83, 86, 86,
	-30, -30, -30, -8, -8, -9, -27, -27, 75, 86,
	86, 30, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 107, 0, 43, 105, 30, 0, 29, 2, 0,
	114, 0, 0, 0, 105, 0, 24, 0, 0, 31,
	32, 33, 46, 0, 48, 97, 184, 0, 0, 21,
	115, 0, 0, 0, 0, 24, 44, 0, 0, 0,
	106, 119, 124, 154, 0, 0, 0, 157, 158, 159,
	160, 0, 0, 0, 0, 0, 0, 196, 197, 198,
	199, 200, 201, 202, 203, 204, 0, 28, 34, 35,
	37, 38, 41, 119, 0, 49, 0, 0, 0, 0,
	94, 95, 98, 0, 100, 0, 0, 0, 0, 0,
	0, 116, 0, 117, 108, 109, 111, 0, 0, 22,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 205, 0, 207, 0, 155,
	0, 156, 0, 0, 0, 0, 0, 159, 159, 0,
	0, 0, 210, 0, 212, 0, 0, 40, 0, 0,
	50, 0, 0, 0, 0, 0, 96, 99, 102, 0,
	0, 189, 103, 104, 18, 0, 118, 0, 112, 113,
	8, 0, 0, 0, 0, 26, 0, 0, 0, 0,
	122, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 143, 0, 203, 0,
	148, 0, 150, 152, 125, 206, 0, 0, 176, 0,
	0, 0, 190, 192, 193, 194, 119, 161, 182, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	211, 0, 36, 39, 42, 47, 0, 52, 53, 56,
	0, 0, 68, 0, 0, 0, 0, 80, 0, 0,
	101, 185, 0, 0, 0, 110, 0, 0, 0, 0,
	45, 25, 0, 0, 0, 123, 142, 144, 0, 0,
	149, 151, 153, 126, 0, 208, 209, 177, 0, 0,
	0, 0, 0, 0, 0, 182, 0, 0, 0, 0,
	0, 0, 0, 0, 213, 51, 55, 0, 58, 59,
	62, 74, 0, 0, 86, 0, 0, 71, 0, 70,
	92, 93, 83, 0, 82, 0, 187, 188, 0, 10,
	16, 17, 0, 0, 0, 27, -2, 0, 0, 146,
	147, 178, 179, 191, 195, 162, 183, 180, 0, 164,
	0, 0, 167, 0, 0, 171, 0, 0, 175, 54,
	57, 61, 63, 65, 75, 76, 0, 87, 88, 0,
	69, 73, 81, 85, 186, 19, 9, 12, 0, -2,
	145, 0, 163, 0, 0, 0, 0, 0, 0, 60,
	64, 66, 77, 78, 89, 90, 72, 84, 0, 0,
	0, 181, 165, 166, 0, 170, 169, 0, 174, 173,
	67, 79, 91, 11, 14, 0, 0, 0, 13, 168,
	172, 0, 15,
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
	        cwtee.Switch = parsingStack.Pop().(ast.Expression)
		parsingStack.Push(cwtee)
	}
	case 164:
		//line n1ql.y:1374
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 165:
		//line n1ql.y:1382
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 166:
		//line n1ql.y:1390
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 167:
		//line n1ql.y:1398
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 168:
		//line n1ql.y:1406
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 169:
		//line n1ql.y:1415
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 170:
		//line n1ql.y:1424
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 171:
		//line n1ql.y:1432
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 172:
		//line n1ql.y:1440
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 173:
		//line n1ql.y:1449
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 174:
		//line n1ql.y:1458
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 175:
		//line n1ql.y:1466
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 176:
		//line n1ql.y:1474
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 177:
		//line n1ql.y:1480
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 178:
		//line n1ql.y:1487
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 179:
		//line n1ql.y:1495
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 180:
		//line n1ql.y:1504
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 181:
		//line n1ql.y:1512
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
	case 182:
		//line n1ql.y:1526
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 183:
		//line n1ql.y:1530
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 184:
		//line n1ql.y:1536
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 185:
		//line n1ql.y:1542
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 186:
		//line n1ql.y:1549
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 187:
		//line n1ql.y:1556
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 188:
		//line n1ql.y:1564
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 189:
		//line n1ql.y:1571
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 190:
		//line n1ql.y:1582
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 191:
		//line n1ql.y:1587
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
	case 192:
		//line n1ql.y:1601
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 193:
		//line n1ql.y:1605
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 194:
		//line n1ql.y:1614
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 195:
		//line n1ql.y:1620
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 196:
		//line n1ql.y:1630
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 197:
		//line n1ql.y:1636
		{
		logDebugGrammar("NUMBER")
	}
	case 198:
		//line n1ql.y:1640
		{
		logDebugGrammar("OBJECT")
	}
	case 199:
		//line n1ql.y:1644
		{
		logDebugGrammar("ARRAY")
	}
	case 200:
		//line n1ql.y:1648
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 201:
		//line n1ql.y:1654
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 202:
		//line n1ql.y:1660
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 203:
		//line n1ql.y:1668
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 204:
		//line n1ql.y:1674
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 205:
		//line n1ql.y:1682
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 206:
		//line n1ql.y:1688
		{
		logDebugGrammar("OBJECT")
	}
	case 207:
		//line n1ql.y:1694
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 208:
		//line n1ql.y:1698
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 209:
		//line n1ql.y:1710
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 210:
		//line n1ql.y:1720
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 211:
		//line n1ql.y:1726
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 212:
		//line n1ql.y:1735
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 213:
		//line n1ql.y:1742
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
