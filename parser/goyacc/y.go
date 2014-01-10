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
	s   string
	n   int
	f   float64
}

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
const NEST = 57438
const INNER = 57439
const LEFT = 57440
const OUTER = 57441
const MOD = 57442

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
}

const yyNprod = 215
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1795

var yyAct = []int{

	51, 319, 85, 221, 52, 32, 151, 4, 212, 134,
	244, 104, 78, 163, 140, 35, 343, 150, 88, 340,
	140, 333, 282, 142, 269, 130, 159, 30, 31, 83,
	225, 46, 88, 324, 50, 81, 224, 92, 322, 318,
	181, 172, 156, 128, 99, 388, 198, 271, 270, 106,
	406, 97, 98, 131, 366, 330, 129, 329, 130, 276,
	220, 260, 132, 140, 143, 146, 147, 148, 141, 113,
	114, 115, 116, 118, 119, 120, 128, 121, 126, 124,
	125, 122, 123, 29, 199, 127, 131, 136, 332, 129,
	367, 395, 365, 86, 396, 89, 90, 91, 169, 170,
	180, 158, 157, 160, 161, 162, 179, 86, 117, 89,
	90, 91, 364, 320, 183, 184, 185, 186, 187, 188,
	189, 190, 191, 192, 193, 194, 195, 196, 197, 358,
	182, 200, 165, 306, 199, 133, 247, 248, 235, 355,
	203, 219, 153, 222, 321, 178, 136, 217, 95, 47,
	38, 177, 204, 83, 130, 36, 37, 305, 349, 81,
	239, 313, 236, 206, 205, 233, 154, 115, 116, 118,
	308, 33, 128, 249, 241, 242, 243, 36, 130, 96,
	293, 106, 131, 290, 288, 129, 257, 262, 36, 113,
	114, 115, 116, 118, 261, 240, 128, 259, 256, 265,
	234, 199, 252, 171, 117, 168, 131, 164, 110, 129,
	100, 84, 272, 44, 219, 219, 53, 275, 268, 274,
	217, 217, 277, 278, 284, 285, 286, 287, 117, 289,
	17, 291, 16, 167, 255, 173, 303, 166, 103, 247,
	248, 294, 267, 296, 299, 133, 210, 307, 310, 311,
	292, 95, 312, 301, 304, 327, 136, 309, 316, 368,
	302, 326, 314, 253, 315, 254, 279, 232, 209, 174,
	152, 325, 137, 139, 363, 328, 295, 317, 323, 300,
	219, 132, 96, 334, 335, 231, 217, 250, 331, 208,
	247, 248, 207, 273, 13, 264, 346, 49, 348, 102,
	41, 350, 95, 352, 353, 245, 112, 356, 247, 248,
	48, 251, 360, 354, 175, 176, 357, 362, 42, 359,
	95, 21, 93, 27, 361, 97, 98, 247, 248, 246,
	25, 347, 17, 96, 97, 98, 351, 95, 370, 371,
	409, 372, 373, 111, 374, 375, 94, 17, 387, 16,
	376, 96, 377, 386, 379, 258, 109, 380, 108, 107,
	382, 22, 384, 23, 385, 43, 381, 19, 96, 383,
	222, 12, 10, 389, 26, 2, 135, 70, 69, 18,
	17, 398, 16, 68, 399, 216, 400, 378, 401, 402,
	3, 12, 10, 404, 215, 281, 405, 60, 58, 45,
	17, 130, 16, 57, 101, 40, 403, 105, 87, 34,
	80, 410, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 79, 77, 127, 131,
	28, 15, 129, 263, 392, 14, 24, 393, 39, 20,
	11, 130, 7, 9, 8, 6, 5, 1, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 344, 0, 0, 345, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 341, 0, 0, 342, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 230, 0, 0, 0, 229, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 228, 0, 0, 0, 227, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 408, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 407, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 397, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 394, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 391, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 390, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 130, 0, 127, 131,
	0, 0, 129, 0, 369, 0, 0, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 117, 0, 127, 131, 0, 0, 129, 0, 0,
	0, 0, 339, 0, 0, 0, 130, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 117, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 0, 0, 127, 131, 0, 0, 129, 0, 0,
	0, 0, 0, 0, 0, 0, 130, 0, 338, 0,
	0, 0, 0, 0, 0, 0, 117, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 0, 0, 127, 131, 0, 0, 129, 0, 0,
	0, 0, 0, 0, 0, 0, 130, 0, 337, 0,
	0, 0, 0, 0, 0, 0, 117, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 0, 0, 127, 131, 0, 0, 129, 0, 0,
	0, 0, 336, 0, 0, 0, 130, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 117, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 130, 266, 127, 131, 0, 0, 129, 0, 0,
	283, 0, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 117, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 0, 226, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 0, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 130, 0, 223, 0, 0, 0, 0, 0, 0,
	0, 117, 113, 114, 115, 116, 118, 119, 120, 128,
	121, 126, 124, 125, 122, 123, 130, 0, 127, 131,
	0, 0, 129, 0, 0, 0, 0, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 117, 0, 127, 131, 0, 0, 280, 0, 0,
	0, 0, 0, 0, 0, 0, 130, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 117, 113, 114, 115,
	116, 118, 119, 120, 128, 121, 126, 124, 125, 122,
	123, 130, 0, 127, 131, 0, 0, 155, 0, 0,
	0, 0, 113, 114, 115, 116, 118, 119, 0, 128,
	121, 126, 124, 125, 122, 123, 117, 0, 127, 131,
	0, 0, 129, 0, 130, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 114, 115, 116, 118,
	0, 117, 128, 121, 126, 124, 125, 122, 123, 0,
	0, 127, 131, 213, 214, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 54,
	0, 76, 0, 0, 117, 71, 72, 73, 74, 75,
	59, 67, 0, 56, 218, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 211, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 54, 64, 76, 0, 0, 0, 71,
	72, 73, 74, 75, 59, 67, 0, 56, 218, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 62, 0, 0,
	0, 0, 63, 0, 65, 66, 0, 54, 64, 76,
	0, 0, 0, 71, 72, 73, 74, 75, 59, 67,
	0, 56, 82, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 0, 0, 61, 0, 0, 0, 0, 0,
	0, 62, 0, 0, 0, 0, 63, 0, 65, 66,
	0, 138, 64, 76, 0, 0, 202, 71, 72, 73,
	201, 75, 59, 67, 0, 56, 0, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 0, 0, 61, 0,
	0, 0, 0, 0, 0, 62, 0, 0, 0, 0,
	63, 0, 65, 66, 0, 54, 64, 76, 149, 0,
	0, 71, 72, 73, 74, 75, 59, 67, 0, 56,
	0, 0, 0, 0, 0, 55, 0, 0, 0, 0,
	0, 0, 61, 0, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 63, 0, 65, 66, 0, 138,
	64, 76, 0, 0, 0, 71, 72, 73, 74, 75,
	59, 67, 0, 56, 0, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 54, 64, 76, 0, 0, 0, 71,
	72, 73, 74, 75, 59, 67, 0, 56, 0, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 62, 0, 0,
	0, 0, 63, 0, 65, 66, 0, 138, 64, 76,
	0, 0, 0, 71, 72, 73, 74, 75, 145, 67,
	0, 56, 0, 0, 88, 0, 0, 55, 0, 0,
	0, 0, 0, 0, 61, 0, 0, 0, 0, 0,
	297, 62, 0, 97, 98, 0, 63, 0, 65, 66,
	0, 138, 64, 76, 0, 95, 0, 71, 72, 73,
	74, 75, 144, 67, 298, 56, 0, 0, 0, 0,
	0, 55, 0, 0, 0, 88, 0, 0, 61, 0,
	0, 0, 0, 0, 0, 62, 96, 0, 0, 0,
	63, 237, 65, 66, 97, 98, 64, 0, 0, 86,
	0, 89, 90, 91, 0, 0, 95, 0, 0, 0,
	0, 0, 0, 0, 0, 238, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 96, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 89, 90, 91,
}
var yyPact = []int{

	366, -1000, -1000, 346, -1000, -1000, -1000, -1000, -1000, -1000,
	338, 281, 334, 294, 286, -5, 118, -1000, -1000, 97,
	256, 277, 336, 154, 286, 96, 251, 1555, 1379, -1000,
	-1000, -1000, -1000, 152, -1, 287, -1000, -37, 151, -1000,
	254, 181, 1555, 329, 328, 251, -1000, 149, 298, 265,
	-1000, 1121, -1000, -1000, 196, 1511, 1511, -1000, -1000, -12,
	-1000, 1555, -60, 1643, 1599, 1511, 1511, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1467, -1000, -1000, 218,
	-1000, 107, -1000, 1186, -39, -1000, 129, 7, 129, 129,
	-1000, -86, -1000, 148, 296, 180, 146, 1511, 1511, 144,
	-40, -1000, 178, -1000, -1000, 217, 272, 92, 47, -1000,
	-41, -1000, 1555, 1511, 1511, 1511, 1511, 1511, 1511, 1511,
	1511, 1511, 1511, 1511, 1511, 1511, 1511, 1511, -31, 142,
	1423, 84, 243, -1000, 240, 216, 193, -1000, 86, -1000,
	1291, -16, 1511, 1081, -55, -61, 1041, 511, 471, -1000,
	234, 215, 1379, 141, -1000, 75, 129, 1696, 129, 129,
	129, 270, 252, -1000, 296, -1000, 212, 177, -1000, 1121,
	1121, -1000, 139, -1000, 1555, -1000, -1000, 324, 138, -14,
	135, 129, 248, 104, 104, -25, -25, -25, -25, 1244,
	1211, 128, 128, 128, 128, 128, 128, 128, 1511, -1000,
	1001, 189, 161, -1000, -32, -1000, -1000, 245, -1000, 27,
	1555, -1000, -17, 1335, 1335, 214, -1000, -1000, -1000, 1146,
	-1000, -63, 976, 1511, 1511, 1511, 1511, 125, 1511, 124,
	1511, -1000, 1555, -1000, -1000, -1000, -1000, 121, 13, -1000,
	-1, 1645, 201, 98, -1, 111, 289, 1511, 1511, -1,
	102, 289, -1000, -1000, 207, 226, -42, -1000, 85, -43,
	1555, -48, -1000, -1000, 1555, 128, -1000, 204, 224, -1000,
	-1000, -1000, -1000, 313, -1000, -1000, -1000, -19, -21, 1335,
	25, -65, 1511, 1511, 936, 896, 856, 816, -72, 431,
	-75, 391, -1000, 13, -1000, -1, -1000, 99, 13, -1000,
	-1, -1, 289, 80, -1, 289, 70, -1000, 289, -1,
	1121, 1121, -1000, 289, -1, 223, -1000, -1000, 53, -1000,
	-1000, -1000, 33, -22, 31, -1000, 208, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 1121, 791, -1000, 1511, 1511, -1000,
	1511, 1511, -1000, 1511, 1511, -1000, -1000, -1, -1000, 13,
	-1000, -1, -1000, -1000, -1, 289, -1000, -1, 289, -1,
	-1000, -1, -1000, -1000, -1000, 322, 317, -30, -1000, 1511,
	751, 711, 351, 671, 8, 631, -1000, -1000, -1, -1000,
	-1000, -1, -1000, -1, -1000, -1000, 85, 85, 1555, -1000,
	-1000, -1000, 1511, -1000, -1000, 1511, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -26, 591, 551, 309, -1000, -1000, 85,
	-1000,
}
var yyPgo = []int{

	0, 447, 375, 7, 446, 445, 444, 443, 1, 17,
	442, 440, 439, 438, 294, 436, 374, 310, 435, 433,
	6, 431, 430, 427, 12, 426, 410, 0, 5, 409,
	2, 15, 37, 408, 10, 11, 407, 405, 404, 4,
	216, 403, 398, 397, 3, 395, 8, 394, 385, 383,
	378, 377, 9, 376,
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
	30, 30, 30, 30, 30, 30, 30, 30, 34, 34,
	33, 33, 33, 29, 29, 29, 29, 29, 29, 32,
	32, 16, 16, 12, 12, 35, 35, 36, 36, 36,
	13, 13, 13, 37, 38, 20, 20, 39, 39, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 40, 40, 40,
	41, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 44, 44, 45, 45, 31, 31, 31, 31, 31,
	31, 46, 46, 47, 47, 48, 48, 43, 43, 43,
	43, 43, 43, 43, 49, 49, 50, 50, 52, 52,
	53, 51, 51, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	2, 1, 3, 0, 2, 5, 2, 5, 1, 2,
	2, 4, 3, 3, 5, 4, 3, 4, 5, 4,
	5, 6, 3, 5, 4, 4, 6, 5, 4, 5,
	6, 5, 6, 7, 3, 5, 4, 4, 6, 5,
	4, 5, 5, 6, 6, 7, 3, 5, 4, 4,
	6, 5, 4, 5, 5, 6, 6, 7, 2, 2,
	1, 1, 2, 1, 2, 3, 2, 4, 3, 2,
	2, 0, 2, 0, 3, 1, 3, 1, 2, 2,
	0, 1, 2, 2, 2, 1, 1, 3, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 6, 5, 5,
	3, 4, 3, 4, 3, 4, 1, 2, 2, 1,
	1, 1, 1, 3, 5, 5, 7, 7, 5, 9,
	7, 7, 5, 9, 7, 7, 5, 3, 4, 5,
	5, 3, 5, 0, 2, 1, 4, 6, 5, 5,
	3, 1, 3, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 3, 1, 3,
	3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 88,
	32, 33, -28, 53, -29, -31, 59, 59, 53, -13,
	-37, 44, 41, 29, 59, -16, -28, 53, -17, 46,
	-20, -27, -39, -40, 48, 68, 62, -41, -42, 59,
	-43, 75, 82, 87, 93, 89, 90, 60, -49, -50,
	-51, 54, 55, 56, 57, 58, 50, -23, -24, -25,
	-26, -20, 63, -27, 59, -30, 94, -33, 19, 96,
	97, 98, -32, 35, 59, 50, 81, 38, 39, 81,
	59, -38, 45, 57, -35, -36, -20, 30, 30, -17,
	59, -14, 41, 61, 62, 63, 64, 100, 65, 66,
	67, 69, 73, 74, 71, 72, 70, 77, 68, 81,
	50, 78, -3, 49, -52, -53, 60, -40, 48, -40,
	75, -20, 83, -27, 59, 59, -27, -27, -27, 51,
	-9, -20, 52, 35, 59, 81, 81, -31, 94, 19,
	96, -31, -31, 99, 59, -32, 57, 53, 59, -27,
	-27, 59, 81, 57, 52, 42, 43, 59, 53, 59,
	53, 81, -9, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, 77, 59,
	-27, 57, 53, 56, 68, 80, 79, 49, 49, 52,
	53, 76, -46, 32, 33, -47, -48, -20, 63, -27,
	76, -44, -27, 92, 91, 91, 92, 95, 91, 95,
	91, 51, 52, -24, 59, 63, -28, 35, 59, -30,
	-32, -31, -31, -31, -34, 35, 59, 38, 39, -34,
	35, 59, -32, 51, 53, 57, 59, -35, 31, 59,
	75, 59, -28, -19, 47, -27, 51, 53, 57, 56,
	80, 79, -39, 48, -52, -20, 76, -46, -46, 52,
	81, -45, 85, 84, -27, -27, -27, -27, 59, -27,
	59, -27, -9, 59, -30, -32, -30, 35, 59, -30,
	-32, -34, 59, 35, -34, 59, 35, -30, 59, -34,
	-27, -27, -30, 59, -34, 57, 51, 51, 81, -8,
	28, 59, 81, -9, 81, -20, 57, 51, 51, 76,
	76, -46, 63, 86, -27, -27, 86, 92, 92, 86,
	91, 83, 86, 91, 83, 86, -30, -32, -30, 59,
	-30, -32, -30, -30, -34, 59, -30, -34, 59, -34,
	-30, -34, -30, 51, 59, 59, 76, 59, 51, 83,
	-27, -27, -27, -27, -27, -27, -30, -30, -32, -30,
	-30, -34, -30, -34, -30, -30, 31, 31, 75, -44,
	86, 86, 83, 86, 86, 83, 86, 86, -30, -30,
	-30, -8, -8, -9, -27, -27, 76, 86, 86, 31,
	-8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 113, 0, 43, 111, 30, 0, 29, 2, 0,
	120, 0, 0, 0, 111, 0, 24, 0, 0, 31,
	32, 33, 46, 0, 48, 103, 185, 0, 0, 21,
	121, 0, 0, 0, 0, 24, 44, 0, 0, 0,
	112, 125, 126, 156, 0, 0, 0, 159, 160, 161,
	162, 0, 0, 0, 0, 0, 0, 197, 198, 199,
	200, 201, 202, 203, 204, 205, 0, 28, 34, 35,
	37, 38, 41, 125, 0, 49, 0, 0, 0, 0,
	100, 101, 104, 0, 106, 0, 0, 0, 0, 0,
	0, 122, 0, 123, 114, 115, 117, 0, 0, 22,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 206, 0, 208, 0, 157, 0, 158,
	0, 0, 0, 0, 161, 161, 0, 0, 0, 211,
	0, 213, 0, 0, 40, 0, 0, 50, 0, 0,
	0, 0, 0, 102, 105, 108, 0, 0, 190, 109,
	110, 18, 0, 124, 0, 118, 119, 8, 0, 0,
	0, 0, 26, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 142, 143, 0, 145,
	0, 204, 0, 150, 0, 152, 154, 127, 207, 0,
	0, 177, 0, 0, 0, 191, 193, 194, 195, 125,
	163, 183, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 212, 0, 36, 39, 42, 47, 0, 52, 53,
	56, 62, 0, 0, 74, 0, 0, 0, 0, 86,
	0, 0, 107, 186, 0, 0, 0, 116, 0, 0,
	0, 0, 45, 25, 0, 144, 146, 0, 0, 151,
	153, 155, 128, 0, 209, 210, 178, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 214, 51, 55, 57, 59, 0, 64, 65,
	68, 80, 0, 0, 92, 0, 0, 77, 0, 76,
	98, 99, 89, 0, 88, 0, 188, 189, 0, 10,
	16, 17, 0, 0, 0, 27, 0, 148, 149, 179,
	180, 192, 196, 164, 184, 181, 165, 0, 0, 168,
	0, 0, 172, 0, 0, 176, 54, 58, 60, 63,
	67, 69, 71, 81, 82, 0, 93, 94, 0, 75,
	79, 87, 91, 187, 19, 9, 12, 0, 147, 0,
	0, 0, 0, 0, 0, 0, 61, 66, 70, 72,
	83, 84, 95, 96, 78, 90, 0, 0, 0, 182,
	166, 167, 0, 171, 170, 0, 175, 174, 73, 85,
	97, 11, 14, 0, 0, 0, 13, 169, 173, 0,
	15,
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
			logDebugGrammar("SELECT FROM - DATASOURCE over here")
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
			parsingStack.Push(&ast.From{Projection: proj, As: ""})
		}
	case 51:
		//line n1ql.y:474
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 52:
		//line n1ql.y:481
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 53:
		//line n1ql.y:488
		{
			logDebugGrammar("UNNEST nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Over: rest})
		}
	case 54:
		//line n1ql.y:495
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
		}
	case 55:
		//line n1ql.y:502
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
		}
	case 56:
		//line n1ql.y:509
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 57:
		//line n1ql.y:516
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 58:
		//line n1ql.y:523
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 59:
		//line n1ql.y:530
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr, Over: rest})
		}
	case 60:
		//line n1ql.y:538
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 61:
		//line n1ql.y:546
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 62:
		//line n1ql.y:554
		{
			logDebugGrammar("UNNEST")
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type})
		}
	case 63:
		//line n1ql.y:562
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-0].s})
		}
	case 64:
		//line n1ql.y:570
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-0].s})
		}
	case 65:
		//line n1ql.y:578
		{
			logDebugGrammar("UNNEST nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: "", Over: rest})
		}
	case 66:
		//line n1ql.y:586
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-1].s, Over: rest})
		}
	case 67:
		//line n1ql.y:594
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-1].s, Over: rest})
		}
	case 68:
		//line n1ql.y:602
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr})
		}
	case 69:
		//line n1ql.y:610
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})
		}
	case 70:
		//line n1ql.y:618
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})
		}
	case 71:
		//line n1ql.y:626
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr, Over: rest})
		}
	case 72:
		//line n1ql.y:635
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 73:
		//line n1ql.y:644
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 74:
		//line n1ql.y:653
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 75:
		//line n1ql.y:660
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 76:
		//line n1ql.y:667
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 77:
		//line n1ql.y:674
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr, Over: rest})
		}
	case 78:
		//line n1ql.y:682
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 79:
		//line n1ql.y:690
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 80:
		//line n1ql.y:698
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr})

		}
	case 81:
		//line n1ql.y:707
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr, Over: rest})
		}
	case 82:
		//line n1ql.y:716
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})

		}
	case 83:
		//line n1ql.y:725
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 84:
		//line n1ql.y:734
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})
		}
	case 85:
		//line n1ql.y:742
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 86:
		//line n1ql.y:751
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Keys: key_expr})
		}
	case 87:
		//line n1ql.y:758
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 88:
		//line n1ql.y:765
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 89:
		//line n1ql.y:772
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Keys: key_expr, Over: rest})
		}
	case 90:
		//line n1ql.y:780
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 91:
		//line n1ql.y:788
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 92:
		//line n1ql.y:796
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Type: Type, Keys: key_expr})

		}
	case 93:
		//line n1ql.y:805
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 94:
		//line n1ql.y:814
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Oper: "NEST", Type: Type, Keys: key_expr})

		}
	case 95:
		//line n1ql.y:823
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 96:
		//line n1ql.y:832
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Oper: "NEST", Type: Type, Keys: key_expr})
		}
	case 97:
		//line n1ql.y:840
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 98:
		//line n1ql.y:851
		{
			logDebugGrammar("FROM JOIN DATASOURCE with KEY")
			key := parsingStack.Pop().(ast.Expression)
			key_expr := ast.NewKeyExpression(key, "KEY")
			parsingStack.Push(key_expr)
		}
	case 99:
		//line n1ql.y:858
		{
			logDebugGrammar("FROM DATASOURCE with KEYS")
			keys := parsingStack.Pop().(ast.Expression)
			keys_expr := ast.NewKeyExpression(keys, "KEYS")
			parsingStack.Push(keys_expr)

		}
	case 100:
		//line n1ql.y:867
		{
			logDebugGrammar("INNER")
			parsingStack.Push("INNER")
		}
	case 101:
		//line n1ql.y:872
		{
			logDebugGrammar("OUTER")
			parsingStack.Push("LEFT")
		}
	case 102:
		//line n1ql.y:877
		{
			logDebugGrammar("LEFT OUTER")
			parsingStack.Push("LEFT")
		}
	case 103:
		//line n1ql.y:884
		{
			logDebugGrammar("FROM DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 104:
		//line n1ql.y:890
		{
			logDebugGrammar("FROM KEY(S) DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 105:
		//line n1ql.y:896
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE AS ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 106:
		//line n1ql.y:903
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 107:
		//line n1ql.y:910
		{
			logDebugGrammar("FROM DATASOURCE AS ID KEY(S)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 108:
		//line n1ql.y:917
		{
			logDebugGrammar("FROM DATASOURCE ID KEY(s)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 109:
		//line n1ql.y:926
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
	case 110:
		//line n1ql.y:937
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
	case 111:
		//line n1ql.y:951
		{
			logDebugGrammar("SELECT WHERE - EMPTY")
		}
	case 112:
		//line n1ql.y:955
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
	case 114:
		//line n1ql.y:969
		{

		}
	case 115:
		//line n1ql.y:975
		{

		}
	case 116:
		//line n1ql.y:979
		{

		}
	case 117:
		//line n1ql.y:984
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
	case 118:
		//line n1ql.y:995
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
	case 119:
		//line n1ql.y:1006
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
	case 120:
		//line n1ql.y:1018
		{

		}
	case 121:
		//line n1ql.y:1022
		{

		}
	case 122:
		//line n1ql.y:1026
		{

		}
	case 123:
		//line n1ql.y:1032
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
	case 124:
		//line n1ql.y:1046
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
	case 125:
		//line n1ql.y:1063
		{
			logDebugGrammar("EXPRESSION")
		}
	case 126:
		//line n1ql.y:1067
		{
		}
	case 127:
		//line n1ql.y:1071
		{
			logDebugGrammar("sub-query EXPRESSION")

		}
	case 128:
		//line n1ql.y:1076
		{
			logDebugGrammar("sub-query NESTED EXPRESSION")
		}
	case 129:
		//line n1ql.y:1082
		{
			logDebugGrammar("EXPR - PLUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 130:
		//line n1ql.y:1090
		{
			logDebugGrammar("EXPR - MINUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 131:
		//line n1ql.y:1098
		{
			logDebugGrammar("EXPR - MULT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 132:
		//line n1ql.y:1106
		{
			logDebugGrammar("EXPR - DIV")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 133:
		//line n1ql.y:1114
		{
			logDebugGrammar("EXPR - MOD")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 134:
		//line n1ql.y:1122
		{
			logDebugGrammar("EXPR - CONCAT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 135:
		//line n1ql.y:1130
		{
			logDebugGrammar("EXPR - AND")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 136:
		//line n1ql.y:1138
		{
			logDebugGrammar("EXPR - OR")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 137:
		//line n1ql.y:1146
		{
			logDebugGrammar("EXPR - EQ")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 138:
		//line n1ql.y:1154
		{
			logDebugGrammar("EXPR - LT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 139:
		//line n1ql.y:1162
		{
			logDebugGrammar("EXPR - LTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 140:
		//line n1ql.y:1170
		{
			logDebugGrammar("EXPR - GT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 141:
		//line n1ql.y:1178
		{
			logDebugGrammar("EXPR - GTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 142:
		//line n1ql.y:1186
		{
			logDebugGrammar("EXPR - NE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 143:
		//line n1ql.y:1194
		{
			logDebugGrammar("EXPR - LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 144:
		//line n1ql.y:1202
		{
			logDebugGrammar("EXPR - NOT LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 145:
		//line n1ql.y:1210
		{
			logDebugGrammar("EXPR DOT MEMBER")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 146:
		//line n1ql.y:1218
		{
			logDebugGrammar("EXPR BRACKET MEMBER")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 147:
		//line n1ql.y:1226
		{
			logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 148:
		//line n1ql.y:1233
		{
			logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 149:
		//line n1ql.y:1241
		{
			logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 150:
		//line n1ql.y:1248
		{
			logDebugGrammar("SUFFIX_EXPR IS NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 151:
		//line n1ql.y:1255
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 152:
		//line n1ql.y:1262
		{
			logDebugGrammar("SUFFIX_EXPR IS MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 153:
		//line n1ql.y:1269
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 154:
		//line n1ql.y:1276
		{
			logDebugGrammar("SUFFIX_EXPR IS VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 155:
		//line n1ql.y:1283
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 156:
		//line n1ql.y:1290
		{

		}
	case 157:
		//line n1ql.y:1296
		{
			logDebugGrammar("EXPR - NOT")
			operand := parsingStack.Pop()
			thisExpression := ast.NewNotOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 158:
		//line n1ql.y:1303
		{
			logDebugGrammar("EXPR - CHANGE SIGN")
			operand := parsingStack.Pop()
			thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 159:
		//line n1ql.y:1310
		{

		}
	case 160:
		//line n1ql.y:1315
		{
			logDebugGrammar("SUFFIX_EXPR")
		}
	case 161:
		//line n1ql.y:1321
		{
			logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 162:
		//line n1ql.y:1327
		{
			logDebugGrammar("LITERAL")
		}
	case 163:
		//line n1ql.y:1331
		{
			logDebugGrammar("NESTED EXPR")
		}
	case 164:
		//line n1ql.y:1335
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
	case 165:
		//line n1ql.y:1352
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 166:
		//line n1ql.y:1360
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 167:
		//line n1ql.y:1368
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 168:
		//line n1ql.y:1376
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 169:
		//line n1ql.y:1384
		{
			logDebugGrammar("FIRST FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 170:
		//line n1ql.y:1393
		{
			logDebugGrammar("FIRST IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 171:
		//line n1ql.y:1402
		{
			logDebugGrammar("FIRST FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 172:
		//line n1ql.y:1410
		{
			logDebugGrammar("FIRST IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 173:
		//line n1ql.y:1418
		{
			logDebugGrammar("ARRAY FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionArray)
		}
	case 174:
		//line n1ql.y:1427
		{
			logDebugGrammar("ARRAY IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 175:
		//line n1ql.y:1436
		{
			logDebugGrammar("ARRAY FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionArray)
		}
	case 176:
		//line n1ql.y:1444
		{
			logDebugGrammar("ARRAY IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 177:
		//line n1ql.y:1452
		{
			logDebugGrammar("FUNCTION EXPR NOPARAM")
			thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 178:
		//line n1ql.y:1458
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 179:
		//line n1ql.y:1465
		{
			logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			function.SetDistinct(true)
			parsingStack.Push(function)
		}
	case 180:
		//line n1ql.y:1473
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 181:
		//line n1ql.y:1482
		{
			logDebugGrammar("THEN_LIST - SINGLE")
			when_then_list := make([]*ast.WhenThen, 0)
			when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
			when_then_list = append(when_then_list, &when_then)
			parsingStack.Push(when_then_list)
		}
	case 182:
		//line n1ql.y:1490
		{
			logDebugGrammar("THEN_LIST - COMPOUND")
			rest := parsingStack.Pop().([]*ast.WhenThen)
			last := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
			new_list := make([]*ast.WhenThen, 0, len(rest)+1)
			new_list = append(new_list, &last)
			for _, v := range rest {
				new_list = append(new_list, v)
			}
			parsingStack.Push(new_list)
		}
	case 183:
		//line n1ql.y:1504
		{
			logDebugGrammar("ELSE - EMPTY")
		}
	case 184:
		//line n1ql.y:1508
		{
			logDebugGrammar("ELSE - EXPR")
		}
	case 185:
		//line n1ql.y:1514
		{
			logDebugGrammar("PATH - %v", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 186:
		//line n1ql.y:1520
		{
			logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 187:
		//line n1ql.y:1527
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s, yyS[yypt-3].n, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 188:
		//line n1ql.y:1534
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 189:
		//line n1ql.y:1542
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 190:
		//line n1ql.y:1549
		{
			logDebugGrammar("PATH DOT PATH - $1.s")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 191:
		//line n1ql.y:1560
		{
			funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
			parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
		}
	case 192:
		//line n1ql.y:1565
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
	case 193:
		//line n1ql.y:1579
		{
			logDebugGrammar("FUNARG STAR")
		}
	case 194:
		//line n1ql.y:1583
		{
			logDebugGrammar("FUNARG EXPR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 195:
		//line n1ql.y:1592
		{
			logDebugGrammar("FUNSTAR")
			funarg_expr := ast.NewStarFunctionArgExpression()
			parsingStack.Push(funarg_expr)
		}
	case 196:
		//line n1ql.y:1598
		{
			logDebugGrammar("FUN PATH DOT STAR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 197:
		//line n1ql.y:1608
		{
			logDebugGrammar("STRING %s", yyS[yypt-0].s)
			thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 198:
		//line n1ql.y:1614
		{
			logDebugGrammar("NUMBER")
		}
	case 199:
		//line n1ql.y:1618
		{
			logDebugGrammar("OBJECT")
		}
	case 200:
		//line n1ql.y:1622
		{
			logDebugGrammar("ARRAY")
		}
	case 201:
		//line n1ql.y:1626
		{
			logDebugGrammar("TRUE")
			thisExpression := ast.NewLiteralBool(true)
			parsingStack.Push(thisExpression)
		}
	case 202:
		//line n1ql.y:1632
		{
			logDebugGrammar("FALSE")
			thisExpression := ast.NewLiteralBool(false)
			parsingStack.Push(thisExpression)
		}
	case 203:
		//line n1ql.y:1638
		{
			logDebugGrammar("NULL")
			thisExpression := ast.NewLiteralNull()
			parsingStack.Push(thisExpression)
		}
	case 204:
		//line n1ql.y:1646
		{
			logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
			thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
			parsingStack.Push(thisExpression)
		}
	case 205:
		//line n1ql.y:1652
		{
			logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
			thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
			parsingStack.Push(thisExpression)
		}
	case 206:
		//line n1ql.y:1660
		{
			logDebugGrammar("EMPTY OBJECT")
			emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
			parsingStack.Push(emptyObject)
		}
	case 207:
		//line n1ql.y:1666
		{
			logDebugGrammar("OBJECT")
		}
	case 208:
		//line n1ql.y:1672
		{
			logDebugGrammar("NAMED EXPR LIST SINGLE")
		}
	case 209:
		//line n1ql.y:1676
		{
			logDebugGrammar("NAMED EXPR LIST COMPOUND")
			last := parsingStack.Pop().(*ast.LiteralObject)
			rest := parsingStack.Pop().(*ast.LiteralObject)
			for k, v := range last.Val {
				rest.Val[k] = v
			}
			parsingStack.Push(rest)
		}
	case 210:
		//line n1ql.y:1688
		{
			logDebugGrammar("NAMED EXPR SINGLE")
			thisKey := yyS[yypt-2].s
			thisValue := parsingStack.Pop().(ast.Expression)
			thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
			parsingStack.Push(thisExpression)
		}
	case 211:
		//line n1ql.y:1698
		{
			logDebugGrammar("EMPTY ARRAY")
			thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 212:
		//line n1ql.y:1704
		{
			logDebugGrammar("ARRAY")
			exp_list := parsingStack.Pop().(ast.ExpressionList)
			thisExpression := ast.NewLiteralArray(exp_list)
			parsingStack.Push(thisExpression)
		}
	case 213:
		//line n1ql.y:1713
		{
			logDebugGrammar("EXPRESSION LIST SINGLE")
			exp_list := make(ast.ExpressionList, 0)
			exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
			parsingStack.Push(exp_list)
		}
	case 214:
		//line n1ql.y:1720
		{
			logDebugGrammar("EXPRESSION LIST COMPOUND")
			rest := parsingStack.Pop().(ast.ExpressionList)
			last := parsingStack.Pop()
			new_list := make(ast.ExpressionList, 0, len(rest)+1)
			new_list = append(new_list, last.(ast.Expression))
			for _, v := range rest {
				new_list = append(new_list, v)
			}
			parsingStack.Push(new_list)
		}
	}
	goto yystack /* stack new state and value */
}
