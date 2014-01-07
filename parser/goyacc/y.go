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

const yyNprod = 212
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1701

var yyAct = []int{

	51, 313, 84, 214, 143, 205, 32, 103, 148, 77,
	240, 160, 337, 334, 74, 35, 75, 147, 156, 199,
	69, 70, 71, 198, 73, 57, 65, 133, 54, 82,
	133, 327, 46, 274, 53, 318, 50, 80, 316, 135,
	312, 59, 178, 218, 30, 31, 217, 195, 60, 91,
	169, 105, 153, 61, 98, 63, 64, 400, 129, 62,
	87, 360, 136, 139, 140, 141, 324, 323, 134, 112,
	113, 114, 115, 117, 118, 119, 127, 120, 125, 123,
	124, 121, 122, 265, 382, 126, 130, 268, 256, 128,
	87, 389, 213, 155, 390, 157, 133, 166, 167, 145,
	29, 154, 361, 158, 159, 177, 267, 266, 116, 96,
	97, 176, 359, 180, 181, 182, 183, 184, 185, 186,
	187, 188, 189, 190, 191, 192, 193, 194, 196, 179,
	197, 314, 326, 129, 212, 85, 215, 88, 89, 90,
	358, 142, 210, 162, 112, 113, 114, 115, 117, 196,
	82, 127, 145, 231, 200, 352, 129, 235, 80, 229,
	232, 130, 315, 349, 128, 85, 201, 88, 89, 90,
	245, 237, 238, 239, 127, 129, 175, 203, 202, 253,
	105, 150, 174, 116, 130, 258, 47, 128, 114, 115,
	117, 38, 36, 127, 33, 343, 261, 37, 307, 302,
	36, 287, 282, 130, 236, 151, 128, 212, 212, 280,
	36, 248, 269, 270, 257, 210, 210, 276, 277, 278,
	279, 255, 281, 252, 283, 116, 230, 196, 168, 165,
	284, 161, 109, 99, 83, 285, 44, 288, 164, 290,
	293, 321, 163, 301, 304, 305, 286, 320, 306, 295,
	298, 310, 264, 303, 251, 170, 102, 309, 308, 249,
	52, 250, 263, 226, 271, 228, 225, 171, 149, 319,
	362, 357, 212, 322, 317, 328, 329, 325, 300, 311,
	210, 243, 244, 227, 289, 297, 224, 294, 243, 244,
	340, 260, 342, 94, 49, 344, 13, 346, 347, 101,
	94, 350, 299, 41, 48, 111, 354, 348, 42, 296,
	351, 356, 21, 353, 131, 132, 172, 173, 355, 246,
	243, 244, 243, 244, 95, 241, 96, 97, 243, 244,
	27, 95, 364, 365, 94, 366, 367, 341, 368, 369,
	94, 25, 345, 247, 370, 110, 371, 17, 373, 242,
	108, 374, 403, 381, 376, 380, 378, 254, 379, 107,
	375, 12, 10, 377, 215, 95, 43, 383, 106, 19,
	17, 95, 16, 26, 22, 392, 23, 2, 393, 144,
	394, 18, 395, 396, 3, 12, 10, 398, 68, 67,
	399, 66, 209, 372, 17, 129, 16, 208, 45, 273,
	397, 58, 56, 55, 100, 404, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	40, 104, 126, 130, 86, 34, 128, 79, 386, 78,
	76, 387, 28, 15, 259, 129, 14, 24, 39, 20,
	11, 7, 9, 8, 6, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	5, 4, 126, 130, 1, 0, 128, 0, 338, 0,
	0, 339, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 335, 0,
	0, 336, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 0, 129, 223, 0, 0, 0,
	222, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 0, 129, 221, 0, 0, 0,
	220, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 402, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 401, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 391, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 388, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 385, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 384, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	129, 0, 126, 130, 0, 0, 128, 0, 363, 0,
	0, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 116, 0, 126, 130, 0,
	0, 128, 0, 0, 0, 0, 333, 0, 0, 0,
	129, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 0, 0, 126, 130, 0,
	0, 128, 0, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 332, 0, 0, 0, 0, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 0, 0, 126, 130, 0,
	0, 128, 0, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 331, 0, 0, 0, 0, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 0, 0, 126, 130, 0,
	0, 128, 0, 0, 0, 0, 330, 0, 0, 0,
	129, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 129, 262, 126, 130, 0,
	0, 128, 0, 0, 275, 0, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	116, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 0, 129, 0, 219, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 0, 129, 0, 216, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	129, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 116, 0, 126, 130, 0,
	0, 272, 0, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 129, 0, 126, 130, 0,
	0, 152, 0, 0, 0, 0, 112, 113, 114, 115,
	117, 118, 0, 127, 120, 125, 123, 124, 121, 122,
	116, 0, 126, 130, 0, 0, 128, 0, 129, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 112,
	113, 114, 115, 117, 0, 116, 127, 120, 125, 123,
	124, 121, 122, 0, 0, 126, 130, 206, 207, 128,
	0, 0, 92, 0, 0, 96, 97, 0, 0, 0,
	0, 0, 0, 74, 0, 75, 0, 94, 116, 69,
	70, 71, 72, 73, 57, 65, 93, 54, 211, 0,
	0, 0, 0, 53, 0, 0, 0, 0, 0, 0,
	59, 204, 0, 0, 0, 0, 0, 60, 95, 0,
	0, 0, 61, 0, 63, 64, 0, 74, 62, 75,
	0, 0, 0, 69, 70, 71, 72, 73, 57, 65,
	0, 54, 211, 0, 0, 0, 0, 53, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 0,
	0, 60, 0, 0, 0, 0, 61, 0, 63, 64,
	0, 74, 62, 75, 0, 0, 0, 69, 70, 71,
	72, 73, 57, 65, 0, 54, 81, 0, 0, 0,
	0, 53, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 0, 0, 0, 0, 60, 0, 0, 0, 0,
	61, 0, 63, 64, 0, 74, 62, 75, 146, 0,
	0, 69, 70, 71, 72, 73, 57, 65, 0, 54,
	0, 0, 0, 0, 0, 53, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 0, 0, 60,
	0, 0, 0, 0, 61, 0, 63, 64, 0, 74,
	62, 75, 0, 0, 0, 69, 70, 71, 72, 73,
	57, 65, 0, 54, 0, 0, 0, 0, 0, 53,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 0, 0, 60, 0, 0, 0, 0, 61, 0,
	63, 64, 0, 74, 62, 75, 0, 0, 0, 69,
	70, 71, 72, 73, 138, 65, 0, 54, 0, 0,
	87, 0, 0, 53, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 0, 0, 0, 291, 60, 0, 96,
	97, 0, 61, 0, 63, 64, 0, 74, 62, 75,
	0, 94, 0, 69, 70, 71, 72, 73, 137, 65,
	292, 54, 0, 0, 0, 0, 0, 53, 0, 0,
	0, 87, 0, 0, 59, 0, 0, 0, 0, 0,
	0, 60, 95, 0, 0, 0, 61, 233, 63, 64,
	96, 97, 62, 0, 0, 85, 0, 88, 89, 90,
	0, 0, 94, 0, 0, 0, 0, 0, 0, 0,
	0, 234, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 95, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 85, 0, 88, 89,
	90,
}
var yyPact = []int{

	360, -1000, -1000, 336, -1000, -1000, -1000, -1000, -1000, -1000,
	340, 272, 347, 305, 293, 12, 141, -1000, -1000, 138,
	259, 267, 337, 177, 293, 133, 248, 1461, 1373, -1000,
	-1000, -1000, -1000, 175, 41, 1287, -1000, -27, 174, -1000,
	254, 199, 1461, 338, 329, 248, -1000, 173, 313, 264,
	-1000, 1115, -1000, 1461, 1461, -1000, -1000, 21, -1000, 1461,
	-44, 1549, 1505, 1461, 1461, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 92, 1417, -1000, -1000, 216, -1000,
	146, -1000, 1180, -29, -1000, 151, -1, 151, 151, -1000,
	-88, -1000, 172, 288, 185, 170, 1461, 1461, 169, -31,
	-1000, 198, -1000, -1000, 215, 274, 123, 52, -1000, -39,
	-1000, 1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461, 1461,
	1461, 1461, 1461, 1461, 1461, 1461, 1461, -30, 168, -34,
	98, -1000, -1000, 1285, 16, 1461, 1075, -45, -48, 1035,
	505, 465, -1000, 237, 214, 210, -1000, 232, 213, 1373,
	167, -1000, 90, 151, 1602, 151, 151, 151, 290, 284,
	-1000, 288, -1000, 208, 197, -1000, 1115, 1115, -1000, 164,
	-1000, 1461, -1000, -1000, 326, 162, 13, 155, 151, 244,
	125, 125, 106, 106, 106, 106, 1238, 1205, 83, 83,
	83, 83, 83, 83, 83, 1461, -1000, 995, 209, 195,
	-1000, 27, -1000, -1000, -1000, 11, 1329, 1329, 212, -1000,
	-1000, -1000, 1140, -1000, -52, 970, 1461, 1461, 1461, 1461,
	150, 1461, 143, 1461, -1000, 39, 1461, -1000, 1461, -1000,
	-1000, -1000, -1000, 142, 71, -1000, 41, 1551, 250, 243,
	41, 140, 282, 1461, 1461, 41, 139, 282, -1000, -1000,
	200, 228, -41, -1000, 103, -43, 1461, -46, -1000, -1000,
	1461, 83, -1000, 190, 222, -1000, -1000, -1000, -1000, -9,
	-10, 1329, 69, -55, 1461, 1461, 930, 890, 850, 810,
	-78, 425, -79, 385, -1000, -1000, -1000, 71, -1000, 41,
	-1000, 136, 71, -1000, 41, 41, 282, 104, 41, 282,
	96, -1000, 282, 41, 1115, 1115, -1000, 282, 41, 220,
	-1000, -1000, 81, -1000, -1000, -1000, 53, -15, 43, -1000,
	219, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 1115, 785,
	-1000, 1461, 1461, -1000, 1461, 1461, -1000, 1461, 1461, -1000,
	-1000, 41, -1000, 71, -1000, 41, -1000, -1000, 41, 282,
	-1000, 41, 282, 41, -1000, 41, -1000, -1000, -1000, 324,
	322, 9, -1000, 1461, 745, 705, 345, 665, 8, 625,
	-1000, -1000, 41, -1000, -1000, 41, -1000, 41, -1000, -1000,
	103, 103, 1461, -1000, -1000, -1000, 1461, -1000, -1000, 1461,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -19, 585, 545,
	321, -1000, -1000, 103, -1000,
}
var yyPgo = []int{

	0, 464, 377, 461, 460, 444, 443, 442, 1, 17,
	441, 440, 439, 438, 296, 437, 373, 304, 436, 434,
	8, 433, 432, 430, 9, 429, 427, 0, 6, 425,
	2, 15, 49, 424, 10, 7, 421, 420, 404, 260,
	403, 402, 401, 3, 399, 5, 397, 392, 391, 389,
	388, 4, 379,
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
	0, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	2, 1, 3, 0, 2, 5, 2, 5, 1, 2,
	2, 4, 3, 3, 5, 4, 3, 4, 5, 4,
	5, 6, 3, 5, 4, 4, 6, 5, 4, 5,
	6, 5, 6, 7, 3, 5, 4, 4, 6, 5,
	4, 5, 5, 6, 6, 7, 3, 5, 4, 4,
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
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 88,
	32, 33, -28, 53, -29, -31, 59, 59, 53, -13,
	-37, 44, 41, 29, 59, -16, -28, 53, -17, 46,
	-20, -27, -39, 68, 62, -40, -41, 59, -42, 75,
	82, 87, 93, 89, 90, 60, -48, -49, -50, 54,
	55, 56, 57, 58, 48, 50, -23, -24, -25, -26,
	-20, 63, -27, 59, -30, 94, -33, 19, 96, 97,
	98, -32, 35, 59, 50, 81, 38, 39, 81, 59,
	-38, 45, 57, -35, -36, -20, 30, 30, -17, 59,
	-14, 41, 61, 62, 63, 64, 100, 65, 66, 67,
	69, 73, 74, 71, 72, 70, 77, 68, 81, 50,
	78, -39, -39, 75, -20, 83, -27, 59, 59, -27,
	-27, -27, 49, -51, -52, 60, 51, -9, -20, 52,
	35, 59, 81, 81, -31, 94, 19, 96, -31, -31,
	99, 59, -32, 57, 53, 59, -27, -27, 59, 81,
	57, 52, 42, 43, 59, 53, 59, 53, 81, -9,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, 77, 59, -27, 57, 53,
	56, 68, 80, 79, 76, -45, 32, 33, -46, -47,
	-20, 63, -27, 76, -43, -27, 92, 91, 91, 92,
	95, 91, 95, 91, 49, 52, 53, 51, 52, -24,
	59, 63, -28, 35, 59, -30, -32, -31, -31, -31,
	-34, 35, 59, 38, 39, -34, 35, 59, -32, 51,
	53, 57, 59, -35, 31, 59, 75, 59, -28, -19,
	47, -27, 51, 53, 57, 56, 80, 79, 76, -45,
	-45, 52, 81, -44, 85, 84, -27, -27, -27, -27,
	59, -27, 59, -27, -51, -20, -9, 59, -30, -32,
	-30, 35, 59, -30, -32, -34, 59, 35, -34, 59,
	35, -30, 59, -34, -27, -27, -30, 59, -34, 57,
	51, 51, 81, -8, 28, 59, 81, -9, 81, -20,
	57, 51, 51, 76, 76, -45, 63, 86, -27, -27,
	86, 92, 92, 86, 91, 83, 86, 91, 83, 86,
	-30, -32, -30, 59, -30, -32, -30, -30, -34, 59,
	-30, -34, 59, -34, -30, -34, -30, 51, 59, 59,
	76, 59, 51, 83, -27, -27, -27, -27, -27, -27,
	-30, -30, -32, -30, -30, -34, -30, -34, -30, -30,
	31, 31, 75, -43, 86, 86, 83, 86, 86, 83,
	86, 86, -30, -30, -30, -8, -8, -9, -27, -27,
	76, 86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 113, 0, 43, 111, 30, 0, 29, 2, 0,
	120, 0, 0, 0, 111, 0, 24, 0, 0, 31,
	32, 33, 46, 0, 48, 103, 182, 0, 0, 21,
	121, 0, 0, 0, 0, 24, 44, 0, 0, 0,
	112, 125, 153, 0, 0, 156, 157, 158, 159, 0,
	0, 0, 0, 0, 0, 194, 195, 196, 197, 198,
	199, 200, 201, 202, 0, 0, 28, 34, 35, 37,
	38, 41, 125, 0, 49, 0, 0, 0, 0, 100,
	101, 104, 0, 106, 0, 0, 0, 0, 0, 0,
	122, 0, 123, 114, 115, 117, 0, 0, 22, 0,
	23, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 154, 155, 0, 0, 0, 0, 158, 158, 0,
	0, 0, 203, 0, 205, 0, 208, 0, 210, 0,
	0, 40, 0, 0, 50, 0, 0, 0, 0, 0,
	102, 105, 108, 0, 0, 187, 109, 110, 18, 0,
	124, 0, 118, 119, 8, 0, 0, 0, 0, 26,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 0, 142, 0, 201, 0,
	147, 0, 149, 151, 174, 0, 0, 0, 188, 190,
	191, 192, 125, 160, 180, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 204, 0, 0, 209, 0, 36,
	39, 42, 47, 0, 52, 53, 56, 62, 0, 0,
	74, 0, 0, 0, 0, 86, 0, 0, 107, 183,
	0, 0, 0, 116, 0, 0, 0, 0, 45, 25,
	0, 141, 143, 0, 0, 148, 150, 152, 175, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 206, 207, 211, 51, 55, 57,
	59, 0, 64, 65, 68, 80, 0, 0, 92, 0,
	0, 77, 0, 76, 98, 99, 89, 0, 88, 0,
	185, 186, 0, 10, 16, 17, 0, 0, 0, 27,
	0, 145, 146, 176, 177, 189, 193, 161, 181, 178,
	162, 0, 0, 165, 0, 0, 169, 0, 0, 173,
	54, 58, 60, 63, 67, 69, 71, 81, 82, 0,
	93, 94, 0, 75, 79, 87, 91, 184, 19, 9,
	12, 0, 144, 0, 0, 0, 0, 0, 0, 0,
	61, 66, 70, 72, 83, 84, 95, 96, 78, 90,
	0, 0, 0, 179, 163, 164, 0, 168, 167, 0,
	172, 171, 73, 85, 97, 11, 14, 0, 0, 0,
	13, 166, 170, 0, 15,
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
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over: rest})
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
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 57:
		//line n1ql.y:515
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 58:
		//line n1ql.y:522
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 59:
		//line n1ql.y:529
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr, Over: rest})
		}
	case 60:
		//line n1ql.y:537
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 61:
		//line n1ql.y:545
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 62:
		//line n1ql.y:553
		{
			logDebugGrammar("UNNEST")
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type})
		}
	case 63:
		//line n1ql.y:561
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-1].s})
		}
	case 64:
		//line n1ql.y:569
		{
			logDebugGrammar("UNNEST AS")
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-1].s})
		}
	case 65:
		//line n1ql.y:577
		{
			logDebugGrammar("UNNEST nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: "", Over: rest})
		}
	case 66:
		//line n1ql.y:585
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-2].s, Over: rest})
		}
	case 67:
		//line n1ql.y:593
		{
			logDebugGrammar("UNNEST AS nested")
			rest := parsingStack.Pop().(*ast.From)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Type: Type, As: yyS[yypt-2].s, Over: rest})
		}
	case 68:
		//line n1ql.y:601
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr})
		}
	case 69:
		//line n1ql.y:609
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr})
		}
	case 70:
		//line n1ql.y:617
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr})
		}
	case 71:
		//line n1ql.y:625
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr, Over: rest})
		}
	case 72:
		//line n1ql.y:634
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 73:
		//line n1ql.y:643
		{
			logDebugGrammar("UNNEST KEY_EXPR")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 74:
		//line n1ql.y:652
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr})
		}
	case 75:
		//line n1ql.y:659
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 76:
		//line n1ql.y:666
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 77:
		//line n1ql.y:673
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Keys: key_expr, Over: rest})
		}
	case 78:
		//line n1ql.y:681
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 79:
		//line n1ql.y:689
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 80:
		//line n1ql.y:697
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr})

		}
	case 81:
		//line n1ql.y:706
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Type: Type, Keys: key_expr, Over: rest})
		}
	case 82:
		//line n1ql.y:715
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})

		}
	case 83:
		//line n1ql.y:724
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 84:
		//line n1ql.y:733
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Type: Type, Keys: key_expr})
		}
	case 85:
		//line n1ql.y:741
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Type: Type, Keys: key_expr, Over: rest})
		}
	case 86:
		//line n1ql.y:750
		{
			logDebugGrammar("JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Keys: key_expr})
		}
	case 87:
		//line n1ql.y:757
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 88:
		//line n1ql.y:764
		{
			logDebugGrammar("JOIN AS KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-1].s, Keys: key_expr})
		}
	case 89:
		//line n1ql.y:771
		{
			logDebugGrammar("JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Keys: key_expr, Over: rest})
		}
	case 90:
		//line n1ql.y:779
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 91:
		//line n1ql.y:787
		{
			logDebugGrammar("JOIN AS KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: yyS[yypt-2].s, Keys: key_expr, Over: rest})
		}
	case 92:
		//line n1ql.y:795
		{
			logDebugGrammar("TYPE JOIN KEY")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, Oper: "NEST", As: "", Type: Type, Keys: key_expr})

		}
	case 93:
		//line n1ql.y:804
		{
			logDebugGrammar("TYPE JOIN KEY NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: "", Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 94:
		//line n1ql.y:813
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Oper: "NEST", Type: Type, Keys: key_expr})

		}
	case 95:
		//line n1ql.y:822
		{
			logDebugGrammar("TYPE JOIN KEY IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 96:
		//line n1ql.y:831
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER")
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Oper: "NEST", Type: Type, Keys: key_expr})
		}
	case 97:
		//line n1ql.y:839
		{
			logDebugGrammar("TYPE JOIN KEY AS IDENTIFIER NESTED")
			rest := parsingStack.Pop().(*ast.From)
			key_expr := parsingStack.Pop().(*ast.KeyExpression)
			proj := parsingStack.Pop().(ast.Expression)
			Type := parsingStack.Pop().(string)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s, Oper: "NEST", Type: Type, Keys: key_expr, Over: rest})
		}
	case 98:
		//line n1ql.y:850
		{
			logDebugGrammar("FROM JOIN DATASOURCE with KEY")
			key := parsingStack.Pop().(ast.Expression)
			key_expr := ast.NewKeyExpression(key, "KEY")
			parsingStack.Push(key_expr)
		}
	case 99:
		//line n1ql.y:857
		{
			logDebugGrammar("FROM DATASOURCE with KEYS")
			keys := parsingStack.Pop().(ast.Expression)
			keys_expr := ast.NewKeyExpression(keys, "KEYS")
			parsingStack.Push(keys_expr)

		}
	case 100:
		//line n1ql.y:866
		{
			logDebugGrammar("INNER")
			parsingStack.Push("INNER")
		}
	case 101:
		//line n1ql.y:871
		{
			logDebugGrammar("OUTER")
			parsingStack.Push("LEFT")
		}
	case 102:
		//line n1ql.y:876
		{
			logDebugGrammar("LEFT OUTER")
			parsingStack.Push("LEFT")
		}
	case 103:
		//line n1ql.y:883
		{
			logDebugGrammar("FROM DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 104:
		//line n1ql.y:889
		{
			logDebugGrammar("FROM KEY(S) DATASOURCE")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj})
		}
	case 105:
		//line n1ql.y:895
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE AS ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 106:
		//line n1ql.y:902
		{
			// fixme support over as
			logDebugGrammar("FROM DATASOURCE ID")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
		}
	case 107:
		//line n1ql.y:909
		{
			logDebugGrammar("FROM DATASOURCE AS ID KEY(S)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 108:
		//line n1ql.y:916
		{
			logDebugGrammar("FROM DATASOURCE ID KEY(s)")
			proj := parsingStack.Pop().(ast.Expression)
			parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s})

		}
	case 109:
		//line n1ql.y:925
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
		//line n1ql.y:936
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
		//line n1ql.y:950
		{
			logDebugGrammar("SELECT WHERE - EMPTY")
		}
	case 112:
		//line n1ql.y:954
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
		//line n1ql.y:968
		{

		}
	case 115:
		//line n1ql.y:974
		{

		}
	case 116:
		//line n1ql.y:978
		{

		}
	case 117:
		//line n1ql.y:983
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
		//line n1ql.y:994
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
		//line n1ql.y:1005
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
		//line n1ql.y:1017
		{

		}
	case 121:
		//line n1ql.y:1021
		{

		}
	case 122:
		//line n1ql.y:1025
		{

		}
	case 123:
		//line n1ql.y:1031
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
		//line n1ql.y:1045
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
		//line n1ql.y:1062
		{
			logDebugGrammar("EXPRESSION")
		}
	case 126:
		//line n1ql.y:1068
		{
			logDebugGrammar("EXPR - PLUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 127:
		//line n1ql.y:1076
		{
			logDebugGrammar("EXPR - MINUS")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 128:
		//line n1ql.y:1084
		{
			logDebugGrammar("EXPR - MULT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 129:
		//line n1ql.y:1092
		{
			logDebugGrammar("EXPR - DIV")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 130:
		//line n1ql.y:1100
		{
			logDebugGrammar("EXPR - MOD")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 131:
		//line n1ql.y:1108
		{
			logDebugGrammar("EXPR - CONCAT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 132:
		//line n1ql.y:1116
		{
			logDebugGrammar("EXPR - AND")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 133:
		//line n1ql.y:1124
		{
			logDebugGrammar("EXPR - OR")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
			parsingStack.Push(thisExpression)
		}
	case 134:
		//line n1ql.y:1132
		{
			logDebugGrammar("EXPR - EQ")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 135:
		//line n1ql.y:1140
		{
			logDebugGrammar("EXPR - LT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 136:
		//line n1ql.y:1148
		{
			logDebugGrammar("EXPR - LTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 137:
		//line n1ql.y:1156
		{
			logDebugGrammar("EXPR - GT")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 138:
		//line n1ql.y:1164
		{
			logDebugGrammar("EXPR - GTE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 139:
		//line n1ql.y:1172
		{
			logDebugGrammar("EXPR - NE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 140:
		//line n1ql.y:1180
		{
			logDebugGrammar("EXPR - LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 141:
		//line n1ql.y:1188
		{
			logDebugGrammar("EXPR - NOT LIKE")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 142:
		//line n1ql.y:1196
		{
			logDebugGrammar("EXPR DOT MEMBER")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 143:
		//line n1ql.y:1204
		{
			logDebugGrammar("EXPR BRACKET MEMBER")
			right := parsingStack.Pop()
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 144:
		//line n1ql.y:1212
		{
			logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 145:
		//line n1ql.y:1219
		{
			logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 146:
		//line n1ql.y:1227
		{
			logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 147:
		//line n1ql.y:1234
		{
			logDebugGrammar("SUFFIX_EXPR IS NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 148:
		//line n1ql.y:1241
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 149:
		//line n1ql.y:1248
		{
			logDebugGrammar("SUFFIX_EXPR IS MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 150:
		//line n1ql.y:1255
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 151:
		//line n1ql.y:1262
		{
			logDebugGrammar("SUFFIX_EXPR IS VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 152:
		//line n1ql.y:1269
		{
			logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
			operand := parsingStack.Pop()
			thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 153:
		//line n1ql.y:1276
		{

		}
	case 154:
		//line n1ql.y:1282
		{
			logDebugGrammar("EXPR - NOT")
			operand := parsingStack.Pop()
			thisExpression := ast.NewNotOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 155:
		//line n1ql.y:1289
		{
			logDebugGrammar("EXPR - CHANGE SIGN")
			operand := parsingStack.Pop()
			thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
			parsingStack.Push(thisExpression)
		}
	case 156:
		//line n1ql.y:1296
		{

		}
	case 157:
		//line n1ql.y:1301
		{
			logDebugGrammar("SUFFIX_EXPR")
		}
	case 158:
		//line n1ql.y:1307
		{
			logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 159:
		//line n1ql.y:1313
		{
			logDebugGrammar("LITERAL")
		}
	case 160:
		//line n1ql.y:1317
		{
			logDebugGrammar("NESTED EXPR")
		}
	case 161:
		//line n1ql.y:1321
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
	case 162:
		//line n1ql.y:1338
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 163:
		//line n1ql.y:1346
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 164:
		//line n1ql.y:1354
		{
			logDebugGrammar("ANY IN SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
			parsingStack.Push(collectionAny)
		}
	case 165:
		//line n1ql.y:1362
		{
			logDebugGrammar("ANY SATISFIES")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
			parsingStack.Push(collectionAny)
		}
	case 166:
		//line n1ql.y:1370
		{
			logDebugGrammar("FIRST FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 167:
		//line n1ql.y:1379
		{
			logDebugGrammar("FIRST IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 168:
		//line n1ql.y:1388
		{
			logDebugGrammar("FIRST FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionFirst)
		}
	case 169:
		//line n1ql.y:1396
		{
			logDebugGrammar("FIRST IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
			parsingStack.Push(collectionFirst)
		}
	case 170:
		//line n1ql.y:1404
		{
			logDebugGrammar("ARRAY FOR IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
			parsingStack.Push(collectionArray)
		}
	case 171:
		//line n1ql.y:1413
		{
			logDebugGrammar("ARRAY IN WHEN")
			condition := parsingStack.Pop().(ast.Expression)
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 172:
		//line n1ql.y:1422
		{
			logDebugGrammar("ARRAY FOR IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
			parsingStack.Push(collectionArray)
		}
	case 173:
		//line n1ql.y:1430
		{
			logDebugGrammar("ARRAY IN")
			sub := parsingStack.Pop().(ast.Expression)
			output := parsingStack.Pop().(ast.Expression)
			collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
			parsingStack.Push(collectionArray)
		}
	case 174:
		//line n1ql.y:1438
		{
			logDebugGrammar("FUNCTION EXPR NOPARAM")
			thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 175:
		//line n1ql.y:1444
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 176:
		//line n1ql.y:1451
		{
			logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			function.SetDistinct(true)
			parsingStack.Push(function)
		}
	case 177:
		//line n1ql.y:1459
		{
			logDebugGrammar("FUNCTION EXPR PARAM")
			funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
			thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
			parsingStack.Push(thisExpression)
		}
	case 178:
		//line n1ql.y:1468
		{
			logDebugGrammar("THEN_LIST - SINGLE")
			when_then_list := make([]*ast.WhenThen, 0)
			when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
			when_then_list = append(when_then_list, &when_then)
			parsingStack.Push(when_then_list)
		}
	case 179:
		//line n1ql.y:1476
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
	case 180:
		//line n1ql.y:1490
		{
			logDebugGrammar("ELSE - EMPTY")
		}
	case 181:
		//line n1ql.y:1494
		{
			logDebugGrammar("ELSE - EXPR")
		}
	case 182:
		//line n1ql.y:1500
		{
			logDebugGrammar("PATH - %v", yyS[yypt-0].s)
			thisExpression := ast.NewProperty(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 183:
		//line n1ql.y:1506
		{
			logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 184:
		//line n1ql.y:1513
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s, yyS[yypt-3].n, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 185:
		//line n1ql.y:1520
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
			parsingStack.Push(thisExpression)

		}
	case 186:
		//line n1ql.y:1528
		{
			logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
			left := parsingStack.Pop()
			thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
			parsingStack.Push(thisExpression)
		}
	case 187:
		//line n1ql.y:1535
		{
			logDebugGrammar("PATH DOT PATH - $1.s")
			right := ast.NewProperty(yyS[yypt-0].s)
			left := parsingStack.Pop()
			thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
			parsingStack.Push(thisExpression)
		}
	case 188:
		//line n1ql.y:1546
		{
			funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
			parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
		}
	case 189:
		//line n1ql.y:1551
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
	case 190:
		//line n1ql.y:1565
		{
			logDebugGrammar("FUNARG STAR")
		}
	case 191:
		//line n1ql.y:1569
		{
			logDebugGrammar("FUNARG EXPR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 192:
		//line n1ql.y:1578
		{
			logDebugGrammar("FUNSTAR")
			funarg_expr := ast.NewStarFunctionArgExpression()
			parsingStack.Push(funarg_expr)
		}
	case 193:
		//line n1ql.y:1584
		{
			logDebugGrammar("FUN PATH DOT STAR")
			expr_part := parsingStack.Pop().(ast.Expression)
			funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
			parsingStack.Push(funarg_expr)
		}
	case 194:
		//line n1ql.y:1594
		{
			logDebugGrammar("STRING %s", yyS[yypt-0].s)
			thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
			parsingStack.Push(thisExpression)
		}
	case 195:
		//line n1ql.y:1600
		{
			logDebugGrammar("NUMBER")
		}
	case 196:
		//line n1ql.y:1604
		{
			logDebugGrammar("OBJECT")
		}
	case 197:
		//line n1ql.y:1608
		{
			logDebugGrammar("ARRAY")
		}
	case 198:
		//line n1ql.y:1612
		{
			logDebugGrammar("TRUE")
			thisExpression := ast.NewLiteralBool(true)
			parsingStack.Push(thisExpression)
		}
	case 199:
		//line n1ql.y:1618
		{
			logDebugGrammar("FALSE")
			thisExpression := ast.NewLiteralBool(false)
			parsingStack.Push(thisExpression)
		}
	case 200:
		//line n1ql.y:1624
		{
			logDebugGrammar("NULL")
			thisExpression := ast.NewLiteralNull()
			parsingStack.Push(thisExpression)
		}
	case 201:
		//line n1ql.y:1632
		{
			logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
			thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
			parsingStack.Push(thisExpression)
		}
	case 202:
		//line n1ql.y:1638
		{
			logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
			thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
			parsingStack.Push(thisExpression)
		}
	case 203:
		//line n1ql.y:1646
		{
			logDebugGrammar("EMPTY OBJECT")
			emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
			parsingStack.Push(emptyObject)
		}
	case 204:
		//line n1ql.y:1652
		{
			logDebugGrammar("OBJECT")
		}
	case 205:
		//line n1ql.y:1658
		{
			logDebugGrammar("NAMED EXPR LIST SINGLE")
		}
	case 206:
		//line n1ql.y:1662
		{
			logDebugGrammar("NAMED EXPR LIST COMPOUND")
			last := parsingStack.Pop().(*ast.LiteralObject)
			rest := parsingStack.Pop().(*ast.LiteralObject)
			for k, v := range last.Val {
				rest.Val[k] = v
			}
			parsingStack.Push(rest)
		}
	case 207:
		//line n1ql.y:1674
		{
			logDebugGrammar("NAMED EXPR SINGLE")
			thisKey := yyS[yypt-2].s
			thisValue := parsingStack.Pop().(ast.Expression)
			thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
			parsingStack.Push(thisExpression)
		}
	case 208:
		//line n1ql.y:1684
		{
			logDebugGrammar("EMPTY ARRAY")
			thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
			parsingStack.Push(thisExpression)
		}
	case 209:
		//line n1ql.y:1690
		{
			logDebugGrammar("ARRAY")
			exp_list := parsingStack.Pop().(ast.ExpressionList)
			thisExpression := ast.NewLiteralArray(exp_list)
			parsingStack.Push(thisExpression)
		}
	case 210:
		//line n1ql.y:1699
		{
			logDebugGrammar("EXPRESSION LIST SINGLE")
			exp_list := make(ast.ExpressionList, 0)
			exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
			parsingStack.Push(exp_list)
		}
	case 211:
		//line n1ql.y:1706
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
