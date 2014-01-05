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

const yyNprod = 195
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1668

var yyAct = []int{

	51, 305, 84, 213, 143, 32, 148, 103, 236, 91,
	77, 159, 133, 156, 133, 35, 147, 329, 326, 204,
	319, 272, 129, 135, 199, 310, 308, 304, 217, 82,
	216, 46, 30, 31, 50, 80, 200, 177, 168, 153,
	127, 98, 194, 381, 145, 346, 316, 202, 201, 105,
	130, 315, 266, 128, 212, 364, 129, 254, 133, 269,
	86, 347, 136, 139, 140, 141, 134, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 129, 195, 126, 130, 176, 318, 128, 29, 371,
	157, 175, 372, 345, 114, 115, 117, 165, 166, 127,
	262, 154, 155, 161, 158, 344, 116, 195, 340, 130,
	337, 230, 128, 179, 180, 181, 182, 183, 184, 185,
	186, 187, 188, 189, 190, 191, 192, 193, 178, 297,
	196, 116, 239, 240, 211, 85, 214, 88, 89, 90,
	209, 142, 263, 174, 94, 306, 299, 150, 47, 173,
	82, 261, 145, 296, 36, 38, 80, 234, 288, 231,
	228, 37, 33, 285, 235, 265, 264, 243, 36, 280,
	246, 151, 241, 242, 129, 95, 307, 105, 251, 278,
	86, 36, 255, 256, 253, 112, 113, 114, 115, 117,
	250, 229, 127, 195, 167, 259, 232, 164, 160, 96,
	97, 109, 130, 99, 83, 128, 211, 211, 44, 313,
	163, 94, 209, 209, 162, 312, 274, 275, 276, 277,
	233, 279, 302, 281, 116, 267, 268, 249, 301, 282,
	169, 102, 283, 247, 225, 248, 286, 227, 224, 287,
	290, 291, 95, 170, 284, 294, 298, 289, 239, 240,
	292, 295, 52, 149, 300, 85, 348, 88, 89, 90,
	94, 223, 343, 244, 314, 311, 239, 240, 303, 293,
	211, 309, 226, 320, 321, 49, 209, 258, 94, 101,
	13, 171, 172, 111, 41, 42, 48, 245, 332, 317,
	21, 95, 334, 239, 240, 335, 27, 333, 338, 96,
	97, 25, 336, 342, 17, 339, 131, 132, 341, 95,
	384, 237, 12, 10, 239, 240, 363, 362, 43, 252,
	107, 17, 106, 16, 350, 351, 94, 352, 353, 110,
	354, 355, 108, 19, 144, 238, 356, 26, 22, 357,
	23, 2, 359, 68, 361, 18, 358, 67, 66, 360,
	214, 208, 207, 365, 271, 92, 58, 95, 96, 97,
	56, 374, 45, 375, 376, 377, 55, 100, 40, 379,
	94, 104, 380, 3, 12, 10, 129, 87, 34, 93,
	79, 378, 78, 17, 76, 16, 385, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 95, 28, 126, 130, 15, 257, 128, 14, 368,
	24, 39, 369, 20, 11, 7, 129, 9, 8, 6,
	5, 4, 1, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 330,
	0, 0, 331, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 327,
	0, 0, 328, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 129, 222, 0, 0,
	0, 221, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 129, 220, 0, 0,
	0, 219, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 383, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 382, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 373, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 370, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 367, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 366, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 129, 0, 126, 130, 0, 0, 128, 0, 349,
	0, 0, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 116, 0, 126, 130,
	0, 0, 128, 0, 0, 0, 0, 325, 0, 0,
	0, 129, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 0, 0, 126, 130,
	0, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	0, 129, 0, 324, 0, 0, 0, 0, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 0, 0, 126, 130,
	0, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	0, 129, 0, 323, 0, 0, 0, 0, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 0, 0, 126, 130,
	0, 0, 128, 0, 0, 0, 0, 322, 0, 0,
	0, 129, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 129, 260, 126, 130,
	0, 0, 128, 0, 0, 273, 0, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 116, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 129, 0, 218, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 0, 129, 0, 215, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 129, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 116, 0, 126, 130,
	0, 0, 270, 0, 0, 0, 0, 0, 0, 0,
	0, 129, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 129, 0, 126, 130,
	0, 0, 152, 0, 0, 0, 0, 112, 113, 114,
	115, 117, 118, 0, 127, 120, 125, 123, 124, 121,
	122, 116, 0, 126, 130, 0, 0, 128, 0, 129,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	112, 113, 114, 115, 117, 0, 116, 127, 120, 125,
	123, 124, 121, 122, 0, 0, 126, 130, 205, 206,
	128, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 74, 0, 75, 0, 0, 116,
	69, 70, 71, 72, 73, 57, 65, 0, 54, 210,
	0, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 59, 203, 0, 0, 0, 0, 0, 60, 0,
	0, 0, 0, 61, 0, 63, 64, 0, 74, 62,
	75, 0, 0, 0, 69, 70, 71, 72, 73, 57,
	65, 0, 54, 210, 0, 0, 0, 0, 53, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	0, 0, 60, 0, 0, 0, 0, 61, 0, 63,
	64, 0, 74, 62, 75, 0, 0, 0, 69, 70,
	71, 72, 73, 57, 65, 0, 54, 81, 0, 0,
	0, 0, 53, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 0, 0, 60, 0, 0, 0,
	0, 61, 0, 63, 64, 0, 74, 62, 75, 0,
	0, 198, 69, 70, 71, 197, 73, 57, 65, 0,
	54, 0, 0, 0, 0, 0, 53, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 0, 0, 0, 0,
	60, 0, 0, 0, 0, 61, 0, 63, 64, 0,
	74, 62, 75, 146, 0, 0, 69, 70, 71, 72,
	73, 57, 65, 0, 54, 0, 0, 0, 0, 0,
	53, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 0, 0, 60, 0, 0, 0, 0, 61,
	0, 63, 64, 0, 74, 62, 75, 0, 0, 0,
	69, 70, 71, 72, 73, 57, 65, 0, 54, 0,
	0, 0, 0, 0, 53, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 0, 0, 60, 0,
	0, 0, 0, 61, 0, 63, 64, 0, 74, 62,
	75, 0, 0, 0, 69, 70, 71, 72, 73, 138,
	65, 0, 54, 0, 0, 0, 0, 0, 53, 0,
	0, 0, 0, 0, 0, 59, 0, 0, 0, 0,
	0, 0, 60, 0, 0, 0, 0, 61, 0, 63,
	64, 0, 74, 62, 75, 0, 0, 0, 69, 70,
	71, 72, 73, 137, 65, 0, 54, 0, 0, 0,
	0, 0, 53, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 0, 0, 60, 0, 0, 0,
	0, 61, 0, 63, 64, 0, 0, 62,
}
var yyPact = []int{

	349, -1000, -1000, 287, -1000, -1000, -1000, -1000, -1000, -1000,
	304, 250, 311, 265, 259, 0, 109, -1000, -1000, 102,
	240, 244, 289, 149, 259, 95, 229, 1486, 1354, -1000,
	-1000, -1000, -1000, 145, 41, 320, -1000, -40, 144, -1000,
	234, 174, 1486, 292, 290, 229, -1000, 142, 270, 242,
	-1000, 1096, -1000, 1486, 1486, -1000, -1000, -17, -1000, 1486,
	-60, 1574, 1530, 1486, 1486, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 92, 1442, -1000, -1000, 201, -1000,
	112, -1000, 1161, -42, -1000, 122, 122, -6, 122, -1000,
	-88, -1000, 139, 261, 157, 138, 1486, 1486, 135, -43,
	-1000, 173, -1000, -1000, 191, 239, 90, 32, -1000, -44,
	-1000, 1486, 1486, 1486, 1486, 1486, 1486, 1486, 1486, 1486,
	1486, 1486, 1486, 1486, 1486, 1486, 1486, -35, 134, 1398,
	-32, -1000, -1000, 1266, -22, 1486, 1056, -61, -63, 1016,
	486, 446, -1000, 212, 186, 181, -1000, 221, 185, 1354,
	132, -1000, 48, 122, 161, 276, 122, 122, 228, -1000,
	261, -1000, 182, 170, -1000, 1096, 1096, -1000, 131, -1000,
	1486, -1000, -1000, 288, 125, -18, 123, 122, 230, 31,
	31, -28, -28, -28, -28, 1219, 1186, 124, 124, 124,
	124, 124, 124, 124, 1486, -1000, 976, 98, 43, -1000,
	86, -1000, -1000, -1000, -24, 1310, 1310, 7, -1000, -1000,
	-1000, 1121, -1000, -64, 951, 1486, 1486, 1486, 1486, 120,
	1486, 110, 1486, -1000, -16, 1486, -1000, 1486, -1000, -1000,
	-1000, -1000, 104, 41, -1000, -1000, 41, 99, 255, 1486,
	1486, 210, 94, 41, 87, 255, -1000, -1000, 171, 217,
	-54, -1000, 117, -55, 1486, -56, -1000, -1000, 1486, 124,
	-1000, 158, 213, -1000, -1000, -1000, -1000, -25, -30, 1310,
	23, -66, 1486, 1486, 911, 871, 831, 791, -73, 406,
	-74, 366, -1000, -1000, -1000, 41, -1000, -1000, 255, 41,
	1096, 1096, 41, 255, 51, 41, 255, 49, -1000, 255,
	41, 211, -1000, -1000, 46, -1000, -1000, -1000, 34, -31,
	2, -1000, 205, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	1096, 766, -1000, 1486, 1486, -1000, 1486, 1486, -1000, 1486,
	1486, -1000, -1000, 41, -1000, -1000, 41, 255, -1000, 41,
	255, 41, -1000, -1000, -1000, 286, 285, -20, -1000, 1486,
	726, 686, 326, 646, 6, 606, -1000, -1000, 41, -1000,
	41, -1000, 117, 117, 1486, -1000, -1000, -1000, 1486, -1000,
	-1000, 1486, -1000, -1000, -1000, -1000, -1000, -1000, -33, 566,
	526, 279, -1000, -1000, 117, -1000,
}
var yyPgo = []int{

	0, 422, 341, 421, 420, 419, 418, 417, 1, 16,
	415, 414, 413, 411, 280, 410, 337, 286, 408, 406,
	6, 405, 402, 384, 10, 382, 380, 0, 5, 378,
	2, 15, 9, 8, 377, 7, 371, 368, 367, 252,
	366, 360, 356, 3, 354, 19, 352, 351, 348, 347,
	343, 4, 334,
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
	30, 33, 33, 34, 34, 34, 29, 29, 29, 29,
	29, 29, 32, 32, 16, 16, 12, 12, 35, 35,
	36, 36, 36, 13, 13, 13, 37, 38, 20, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 39, 39, 39,
	40, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 43, 43, 44, 44, 31, 31, 31, 31, 31,
	31, 45, 45, 46, 46, 47, 47, 42, 42, 42,
	42, 42, 42, 42, 48, 48, 49, 49, 51, 51,
	52, 50, 50, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	2, 1, 3, 0, 2, 5, 2, 5, 1, 2,
	2, 4, 3, 3, 3, 5, 4, 3, 5, 4,
	4, 6, 5, 4, 5, 5, 6, 6, 7, 3,
	5, 4, 4, 6, 5, 4, 5, 5, 6, 6,
	7, 2, 2, 1, 1, 2, 1, 2, 3, 2,
	4, 3, 2, 2, 0, 2, 0, 3, 1, 3,
	1, 2, 2, 0, 1, 2, 2, 2, 1, 3,
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
	-20, -27, -39, 68, 62, -40, -41, 59, -42, 75,
	82, 87, 93, 89, 90, 60, -48, -49, -50, 54,
	55, 56, 57, 58, 48, 50, -23, -24, -25, -26,
	-20, 63, -27, 59, -30, 94, 19, -34, 96, 97,
	98, -32, 35, 59, 50, 81, 38, 39, 81, 59,
	-38, 45, 57, -35, -36, -20, 30, 30, -17, 59,
	-14, 41, 61, 62, 63, 64, 100, 65, 66, 67,
	69, 73, 74, 71, 72, 70, 77, 68, 81, 50,
	78, -39, -39, 75, -20, 83, -27, 59, 59, -27,
	-27, -27, 49, -51, -52, 60, 51, -9, -20, 52,
	35, 59, 81, 81, -31, -31, 19, 96, -31, 99,
	59, -32, 57, 53, 59, -27, -27, 59, 81, 57,
	52, 42, 43, 59, 53, 59, 53, 81, -9, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, 77, 59, -27, 57, 53, 56,
	68, 80, 79, 76, -45, 32, 33, -46, -47, -20,
	63, -27, 76, -43, -27, 92, 91, 91, 92, 95,
	91, 95, 91, 49, 52, 53, 51, 52, -24, 59,
	63, -28, 35, 59, -30, -32, -33, 35, 59, 38,
	39, -31, -31, -33, 35, 59, -32, 51, 53, 57,
	59, -35, 31, 59, 75, 59, -28, -19, 47, -27,
	51, 53, 57, 56, 80, 79, 76, -45, -45, 52,
	81, -44, 85, 84, -27, -27, -27, -27, 59, -27,
	59, -27, -51, -20, -9, 59, -30, -30, 59, -33,
	-27, -27, -33, 59, 35, -33, 59, 35, -30, 59,
	-33, 57, 51, 51, 81, -8, 28, 59, 81, -9,
	81, -20, 57, 51, 51, 76, 76, -45, 63, 86,
	-27, -27, 86, 92, 92, 86, 91, 83, 86, 91,
	83, 86, -30, -33, -30, -30, -33, 59, -30, -33,
	59, -33, -30, 51, 59, 59, 76, 59, 51, 83,
	-27, -27, -27, -27, -27, -27, -30, -30, -33, -30,
	-33, -30, 31, 31, 75, -43, 86, 86, 83, 86,
	86, 83, 86, 86, -30, -30, -8, -8, -9, -27,
	-27, 76, 86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 96, 0, 43, 94, 30, 0, 29, 2, 0,
	103, 0, 0, 0, 94, 0, 24, 0, 0, 31,
	32, 33, 46, 0, 48, 86, 165, 0, 0, 21,
	104, 0, 0, 0, 0, 24, 44, 0, 0, 0,
	95, 108, 136, 0, 0, 139, 140, 141, 142, 0,
	0, 0, 0, 0, 0, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 0, 0, 28, 34, 35, 37,
	38, 41, 108, 0, 49, 0, 0, 0, 0, 83,
	84, 87, 0, 89, 0, 0, 0, 0, 0, 0,
	105, 0, 106, 97, 98, 100, 0, 0, 22, 0,
	23, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 137, 138, 0, 0, 0, 0, 141, 141, 0,
	0, 0, 186, 0, 188, 0, 191, 0, 193, 0,
	0, 40, 0, 0, 50, 0, 0, 0, 0, 85,
	88, 91, 0, 0, 170, 92, 93, 18, 0, 107,
	0, 101, 102, 8, 0, 0, 0, 0, 26, 109,
	110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
	120, 121, 122, 123, 0, 125, 0, 184, 0, 130,
	0, 132, 134, 157, 0, 0, 0, 171, 173, 174,
	175, 108, 143, 163, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 187, 0, 0, 192, 0, 36, 39,
	42, 47, 0, 52, 53, 54, 57, 0, 0, 0,
	0, 0, 0, 69, 0, 0, 90, 166, 0, 0,
	0, 99, 0, 0, 0, 0, 45, 25, 0, 124,
	126, 0, 0, 131, 133, 135, 158, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 189, 190, 194, 51, 56, 60, 0, 59,
	81, 82, 63, 0, 0, 75, 0, 0, 72, 0,
	71, 0, 168, 169, 0, 10, 16, 17, 0, 0,
	0, 27, 0, 128, 129, 159, 160, 172, 176, 144,
	164, 161, 145, 0, 0, 148, 0, 0, 152, 0,
	0, 156, 55, 58, 62, 64, 65, 0, 76, 77,
	0, 70, 74, 167, 19, 9, 12, 0, 127, 0,
	0, 0, 0, 0, 0, 0, 61, 66, 67, 78,
	79, 73, 0, 0, 0, 162, 146, 147, 0, 151,
	150, 0, 155, 154, 68, 80, 11, 14, 0, 0,
	0, 13, 149, 153, 0, 15,
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
		//line n1ql.y:252
		{
		}
	case 25:
		//line n1ql.y:255
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
		//line n1ql.y:267
		{
		}
	case 27:
		//line n1ql.y:270
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
		//line n1ql.y:283
		{
			logDebugGrammar("SELECT_SELECT")
		}
	case 29:
		//line n1ql.y:289
		{
			logDebugGrammar("SELECT_SELECT_HEAD")
		}
	case 30:
		//line n1ql.y:295
		{
		}
	case 31:
		//line n1ql.y:298
		{
			/* empty */
		}
	case 32:
		//line n1ql.y:301
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
		//line n1ql.y:311
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
		//line n1ql.y:323
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
		//line n1ql.y:337
		{
			result_expr := parsingStack.Pop().(*ast.ResultExpression)
			parsingStack.Push(ast.ResultExpressionList{result_expr})
		}
	case 36:
		//line n1ql.y:342
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
		//line n1ql.y:355
		{
			logDebugGrammar("RESULT STAR")
		}
	case 38:
		//line n1ql.y:359
		{
			logDebugGrammar("RESULT EXPR")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewResultExpression(expr_part)
			parsingStack.Push(result_expr)
		}
	case 39:
		//line n1ql.y:366
		{
			logDebugGrammar("RESULT EXPR AS ID")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
			parsingStack.Push(result_expr)
		}
	case 40:
		//line n1ql.y:373
		{
			logDebugGrammar("RESULT EXPR ID")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
			parsingStack.Push(result_expr)
		}
	case 41:
		//line n1ql.y:382
		{
			logDebugGrammar("STAR")
			result_expr := ast.NewStarResultExpression()
			parsingStack.Push(result_expr)
		}
	case 42:
		//line n1ql.y:388
		{
			logDebugGrammar("PATH DOT STAR")
			expr_part := parsingStack.Pop().(ast.Expression)
			result_expr := ast.NewDotStarResultExpression(expr_part)
			parsingStack.Push(result_expr)
		}
	case 43:
		//line n1ql.y:397
		{
			logDebugGrammar("SELECT FROM - EMPTY")
		}
	case 44:
		//line n1ql.y:401
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
		//line n1ql.y:412
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
		//line n1ql.y:426
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
		//line n1ql.y:437
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
		//line n1ql.y:451
		{
			logDebugGrammar("FROM DATASOURCE WITHOUT UNNEST")
		}
	case 49:
		//line n1ql.y:455
		{
			logDebugGrammar("FROM DATASOURCE WITH UNNEST")
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
			parsingStack.Push(&ast.From{Projection: proj, As: ""})
		}
	case 51:
		//line n1ql.y:473
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 52:
		//line n1ql.y:480
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 53:
		//line n1ql.y:487
		{
			logDebugGrammar("UNNEST nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Over: rest})
		}
	case 54:
		//line n1ql.y:494
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 55:
		//line n1ql.y:501
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
		}
	case 56:
		//line n1ql.y:508
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
		}
	case 57:
		//line n1ql.y:515
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 58:
		//line n1ql.y:522
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 59:
		//line n1ql.y:529
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 60:
		//line n1ql.y:536
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr, Over: rest})
		}
	case 61:
		//line n1ql.y:544
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 62:
		//line n1ql.y:552
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 63:
		//line n1ql.y:560
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr})

		}
	case 64:
		//line n1ql.y:569
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr, Over: rest})
		}
	case 65:
		//line n1ql.y:578
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})

		}
	case 66:
		//line n1ql.y:587
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 67:
		//line n1ql.y:596
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})
		}
	case 68:
		//line n1ql.y:604
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 69:
		//line n1ql.y:613
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Keys: key_expr})
		}
	case 70:
		//line n1ql.y:620
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 71:
		//line n1ql.y:627
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 72:
		//line n1ql.y:634
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Keys: key_expr, Over: rest})
		}
	case 73:
		//line n1ql.y:642
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 74:
		//line n1ql.y:650
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 75:
		//line n1ql.y:658
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Type: Type, Keys: key_expr})

		}
	case 76:
		//line n1ql.y:667
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 77:
		//line n1ql.y:676
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Oper: "NEST", Type: Type, Keys: key_expr})

		}
	case 78:
		//line n1ql.y:685
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 79:
		//line n1ql.y:694
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Oper: "NEST", Type: Type, Keys: key_expr})
		}
	case 80:
		//line n1ql.y:702
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 81:
		//line n1ql.y:713
		{
			logDebugGrammar("FROM JOIN DATASOURCE with KEY")
			key := parsingStack.Pop().(ast.Expression)
			key_expr := ast.NewKeyExpression(key, "KEY")
			parsingStack.Push(key_expr)
		}
	case 82:
		//line n1ql.y:720
		{
			logDebugGrammar("FROM DATASOURCE with KEYS")
			keys := parsingStack.Pop().(ast.Expression)
			keys_expr := ast.NewKeyExpression(keys, "KEYS")
			parsingStack.Push(keys_expr)

		}
	case 83:
		//line n1ql.y:729
		{
			logDebugGrammar("INNER")
			parsingStack.Push("INNER")
		}
	case 84:
		//line n1ql.y:734
		{
			logDebugGrammar("OUTER")
			parsingStack.Push("LEFT")
		}
	case 85:
		//line n1ql.y:739
		{
			logDebugGrammar("LEFT OUTER")
			parsingStack.Push("LEFT")
		}
	case 86:
		//line n1ql.y:746
		{
			logDebugGrammar("FROM DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 87:
		//line n1ql.y:752
		{
			logDebugGrammar("FROM KEY(S) DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 88:
		//line n1ql.y:758
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE AS ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 89:
		//line n1ql.y:765
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 90:
		//line n1ql.y:772
		{
			logDebugGrammar("FROM DATASOURCE AS ID KEY(S)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 91:
		//line n1ql.y:779
		{
			logDebugGrammar("FROM DATASOURCE ID KEY(s)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 92:
		//line n1ql.y:788
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
	case 93:
		//line n1ql.y:799
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
	case 94:
		//line n1ql.y:813
		{
			logDebugGrammar("SELECT WHERE - EMPTY")
		}
	case 95:
		//line n1ql.y:817
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
	case 97:
		//line n1ql.y:831
		{

		}
	case 98:
		//line n1ql.y:837
		{

		}
	case 99:
		//line n1ql.y:841
		{

		}
	case 100:
		//line n1ql.y:846
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
	case 101:
		//line n1ql.y:857
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
	case 102:
		//line n1ql.y:868
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
	case 103:
		//line n1ql.y:880
		{

		}
	case 104:
		//line n1ql.y:884
		{

		}
	case 105:
		//line n1ql.y:888
		{

		}
	case 106:
		//line n1ql.y:894
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
	case 107:
		//line n1ql.y:908
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
	case 108:
		//line n1ql.y:925
		{
			logDebugGrammar("EXPRESSION")
		}
	case 109:
		//line n1ql.y:931
		{
			logDebugGrammar("EXPR - PLUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 110:
		//line n1ql.y:939
		{
			logDebugGrammar("EXPR - MINUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 111:
		//line n1ql.y:947
		{
			logDebugGrammar("EXPR - MULT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 112:
		//line n1ql.y:955
		{
			logDebugGrammar("EXPR - DIV")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 113:
		//line n1ql.y:963
		{
			logDebugGrammar("EXPR - MOD")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 114:
		//line n1ql.y:971
		{
			logDebugGrammar("EXPR - CONCAT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 115:
		//line n1ql.y:979
		{
			logDebugGrammar("EXPR - AND")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 116:
		//line n1ql.y:987
		{
			logDebugGrammar("EXPR - OR")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 117:
		//line n1ql.y:995
		{
			logDebugGrammar("EXPR - EQ")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 118:
		//line n1ql.y:1003
		{
			logDebugGrammar("EXPR - LT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 119:
		//line n1ql.y:1011
		{
			logDebugGrammar("EXPR - LTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 120:
		//line n1ql.y:1019
		{
			logDebugGrammar("EXPR - GT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 121:
		//line n1ql.y:1027
		{
			logDebugGrammar("EXPR - GTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 122:
		//line n1ql.y:1035
		{
			logDebugGrammar("EXPR - NE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 123:
		//line n1ql.y:1043
		{
			logDebugGrammar("EXPR - LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 124:
		//line n1ql.y:1051
		{
			logDebugGrammar("EXPR - NOT LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 125:
		//line n1ql.y:1059
		{
			logDebugGrammar("EXPR DOT MEMBER")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 126:
		//line n1ql.y:1067
		{
			logDebugGrammar("EXPR BRACKET MEMBER")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 127:
		//line n1ql.y:1075
		{
			logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 128:
		//line n1ql.y:1082
		{
			logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 129:
		//line n1ql.y:1090
		{
			logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 130:
		//line n1ql.y:1097
		{
			logDebugGrammar("SUFFIX_EXPR IS NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 131:
		//line n1ql.y:1104
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 132:
		//line n1ql.y:1111
		{
			logDebugGrammar("SUFFIX_EXPR IS MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 133:
		//line n1ql.y:1118
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 134:
		//line n1ql.y:1125
		{
			logDebugGrammar("SUFFIX_EXPR IS VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 135:
		//line n1ql.y:1132
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 136:
		//line n1ql.y:1139
		{

		}
	case 137:
		//line n1ql.y:1145
		{
			logDebugGrammar("EXPR - NOT")
			operand := parsingStack.Pop()
			thisExpression := ast.NewNotOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 138:
		//line n1ql.y:1152
		{
			logDebugGrammar("EXPR - CHANGE SIGN")
			operand := parsingStack.Pop()
			thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 139:
		//line n1ql.y:1159
		{

		}
	case 140:
		//line n1ql.y:1164
		{
			logDebugGrammar("SUFFIX_EXPR")
		}
	case 141:
		//line n1ql.y:1170
		{
			logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 142:
		//line n1ql.y:1176
		{
			logDebugGrammar("LITERAL")
		}
	case 143:
		//line n1ql.y:1180
		{
			logDebugGrammar("NESTED EXPR")
		}
	case 144:
		//line n1ql.y:1184
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
	case 145:
		//line n1ql.y:1201
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 146:
		//line n1ql.y:1209
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 147:
		//line n1ql.y:1217
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 148:
		//line n1ql.y:1225
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 149:
		//line n1ql.y:1233
		{
			logDebugGrammar("FIRST FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 150:
		//line n1ql.y:1242
		{
			logDebugGrammar("FIRST IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 151:
		//line n1ql.y:1251
		{
			logDebugGrammar("FIRST FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 152:
		//line n1ql.y:1259
		{
			logDebugGrammar("FIRST IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 153:
		//line n1ql.y:1267
		{
			logDebugGrammar("ARRAY FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionArray)
		}
	case 154:
		//line n1ql.y:1276
		{
			logDebugGrammar("ARRAY IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 155:
		//line n1ql.y:1285
		{
			logDebugGrammar("ARRAY FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionArray)
		}
	case 156:
		//line n1ql.y:1293
		{
			logDebugGrammar("ARRAY IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 157:
		//line n1ql.y:1301
		{
			logDebugGrammar("FUNCTION EXPR NOPARAM")
			thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 158:
		//line n1ql.y:1307
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 159:
		//line n1ql.y:1314
		{
			logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			function.SetDistinct(true)
			parsingStack.Push(function)
		}
	case 160:
		//line n1ql.y:1322
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 161:
		//line n1ql.y:1331
		{
			logDebugGrammar("THEN_LIST - SINGLE")
			when_then_list := make([]*ast.WhenThen, 0)
			when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
			when_then_list = append(when_then_list, &when_then)
			parsingStack.Push(when_then_list)
		}
	case 162:
		//line n1ql.y:1339
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
	case 163:
		//line n1ql.y:1353
		{
			logDebugGrammar("ELSE - EMPTY")
		}
	case 164:
		//line n1ql.y:1357
		{
			logDebugGrammar("ELSE - EXPR")
		}
	case 165:
		//line n1ql.y:1363
		{
			logDebugGrammar("PATH - %v", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 166:
		//line n1ql.y:1369
		{
			logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 167:
		//line n1ql.y:1376
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s, yyS[yypt-3].n, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 168:
		//line n1ql.y:1383
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 169:
		//line n1ql.y:1391
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 170:
		//line n1ql.y:1398
		{
			logDebugGrammar("PATH DOT PATH - $1.s")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 171:
		//line n1ql.y:1409
		{
			funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
			parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
		}
	case 172:
		//line n1ql.y:1414
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
	case 173:
		//line n1ql.y:1428
		{
			logDebugGrammar("FUNARG STAR")
		}
	case 174:
		//line n1ql.y:1432
		{
			logDebugGrammar("FUNARG EXPR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 175:
		//line n1ql.y:1441
		{
			logDebugGrammar("FUNSTAR")
			funarg_expr := ast.NewStarFunctionArgExpression()
			parsingStack.Push(funarg_expr)
		}
	case 176:
		//line n1ql.y:1447
		{
			logDebugGrammar("FUN PATH DOT STAR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 177:
		//line n1ql.y:1457
		{
			logDebugGrammar("STRING %s", yyS[yypt-0].s)
			thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 178:
		//line n1ql.y:1463
		{
			logDebugGrammar("NUMBER")
		}
	case 179:
		//line n1ql.y:1467
		{
			logDebugGrammar("OBJECT")
		}
	case 180:
		//line n1ql.y:1471
		{
			logDebugGrammar("ARRAY")
		}
	case 181:
		//line n1ql.y:1475
		{
			logDebugGrammar("TRUE")
			thisExpression := ast.NewLiteralBool(true)
			parsingStack.Push(thisExpression)
		}
	case 182:
		//line n1ql.y:1481
		{
			logDebugGrammar("FALSE")
			thisExpression := ast.NewLiteralBool(false)
			parsingStack.Push(thisExpression)
		}
	case 183:
		//line n1ql.y:1487
		{
			logDebugGrammar("NULL")
			thisExpression := ast.NewLiteralNull()
			parsingStack.Push(thisExpression)
		}
	case 184:
		//line n1ql.y:1495
		{
			logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
			thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
			parsingStack.Push(thisExpression)
		}
	case 185:
		//line n1ql.y:1501
		{
			logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
			thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
			parsingStack.Push(thisExpression)
		}
	case 186:
		//line n1ql.y:1509
		{
			logDebugGrammar("EMPTY OBJECT")
			emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
			parsingStack.Push(emptyObject)
		}
	case 187:
		//line n1ql.y:1515
		{
			logDebugGrammar("OBJECT")
		}
	case 188:
		//line n1ql.y:1521
		{
			logDebugGrammar("NAMED EXPR LIST SINGLE")
		}
	case 189:
		//line n1ql.y:1525
		{
			logDebugGrammar("NAMED EXPR LIST COMPOUND")
			last := parsingStack.Pop().(*ast.LiteralObject)
			rest := parsingStack.Pop().(*ast.LiteralObject)
			for k, v := range last.Val {
				rest.Val[k] = v
			}
			parsingStack.Push(rest)
		}
	case 190:
		//line n1ql.y:1537
		{
			logDebugGrammar("NAMED EXPR SINGLE")
			thisKey := yyS[yypt-2].s
			thisValue := parsingStack.Pop().(ast.Expression)
			thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
			parsingStack.Push(thisExpression)
		}
	case 191:
		//line n1ql.y:1547
		{
			logDebugGrammar("EMPTY ARRAY")
			thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 192:
		//line n1ql.y:1553
		{
			logDebugGrammar("ARRAY")
			exp_list := parsingStack.Pop().(ast.ExpressionList)
			thisExpression := ast.NewLiteralArray(exp_list)
			parsingStack.Push(thisExpression)
		}
	case 193:
		//line n1ql.y:1562
		{
			logDebugGrammar("EXPRESSION LIST SINGLE")
			exp_list := make(ast.ExpressionList, 0)
			exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
			parsingStack.Push(exp_list)
		}
	case 194:
		//line n1ql.y:1569
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
