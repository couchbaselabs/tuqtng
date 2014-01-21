
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

const yyNprod = 209
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1731

var yyAct = []int{

	51, 316, 85, 221, 4, 52, 151, 32, 212, 92,
	243, 104, 78, 130, 163, 159, 150, 340, 134, 35,
	140, 140, 337, 330, 113, 114, 115, 116, 118, 83,
	130, 128, 281, 46, 50, 81, 225, 224, 30, 31,
	142, 131, 321, 319, 129, 315, 268, 181, 128, 106,
	172, 156, 99, 198, 400, 361, 382, 88, 131, 132,
	130, 129, 327, 117, 143, 146, 147, 148, 141, 270,
	269, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 326, 259, 127, 131, 203,
	158, 129, 160, 389, 29, 275, 390, 220, 169, 170,
	140, 204, 136, 199, 165, 180, 157, 329, 161, 162,
	117, 179, 206, 205, 183, 184, 185, 186, 187, 188,
	189, 190, 191, 192, 193, 194, 195, 196, 197, 182,
	88, 200, 86, 199, 89, 90, 91, 235, 153, 362,
	17, 219, 16, 222, 317, 133, 294, 217, 360, 97,
	98, 359, 178, 83, 130, 133, 136, 88, 177, 81,
	239, 95, 154, 266, 236, 233, 136, 115, 116, 118,
	295, 353, 128, 248, 251, 318, 97, 98, 240, 241,
	242, 106, 131, 350, 344, 129, 256, 303, 47, 261,
	246, 247, 96, 310, 36, 38, 305, 33, 292, 264,
	289, 37, 95, 36, 117, 86, 287, 89, 90, 91,
	36, 302, 260, 271, 219, 219, 258, 274, 255, 234,
	217, 217, 276, 277, 283, 284, 285, 286, 273, 288,
	199, 290, 86, 96, 89, 90, 91, 171, 168, 164,
	110, 293, 100, 296, 84, 44, 304, 307, 308, 291,
	297, 309, 298, 301, 167, 267, 306, 254, 166, 173,
	300, 311, 53, 246, 247, 249, 324, 313, 246, 247,
	322, 103, 323, 312, 210, 95, 320, 132, 278, 219,
	95, 232, 331, 332, 299, 217, 209, 328, 244, 250,
	174, 246, 247, 152, 252, 343, 253, 363, 345, 358,
	347, 348, 325, 95, 351, 346, 96, 314, 231, 355,
	349, 96, 245, 352, 357, 208, 354, 207, 137, 139,
	13, 356, 272, 263, 49, 48, 102, 175, 176, 112,
	41, 42, 246, 247, 96, 365, 366, 21, 367, 368,
	27, 369, 370, 97, 98, 25, 17, 371, 16, 373,
	17, 403, 374, 381, 372, 376, 380, 378, 257, 379,
	108, 375, 107, 88, 377, 222, 43, 19, 383, 111,
	22, 109, 23, 26, 135, 392, 70, 2, 393, 237,
	394, 18, 395, 396, 69, 68, 216, 398, 215, 280,
	399, 12, 10, 60, 95, 130, 58, 57, 45, 397,
	17, 101, 16, 238, 40, 404, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	105, 87, 127, 131, 34, 96, 129, 80, 386, 79,
	77, 387, 28, 15, 262, 130, 14, 24, 86, 39,
	89, 90, 91, 20, 11, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	7, 9, 127, 131, 8, 6, 129, 5, 341, 1,
	0, 342, 3, 12, 10, 130, 0, 0, 0, 0,
	0, 0, 17, 0, 16, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 338, 0,
	0, 339, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 130, 230, 0, 0, 0,
	229, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 130, 228, 0, 0, 0,
	227, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 402, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 401, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 391, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 388, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 385, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 384, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	130, 0, 127, 131, 0, 0, 129, 0, 364, 0,
	0, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 117, 0, 127, 131, 0,
	0, 129, 0, 0, 0, 0, 336, 0, 0, 0,
	130, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	117, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 0, 0, 127, 131, 0,
	0, 129, 0, 0, 0, 0, 0, 0, 0, 0,
	130, 0, 335, 0, 0, 0, 0, 0, 0, 0,
	117, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 0, 0, 127, 131, 0,
	0, 129, 0, 0, 0, 0, 0, 0, 0, 0,
	130, 0, 334, 0, 0, 0, 0, 0, 0, 0,
	117, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 0, 0, 127, 131, 0,
	0, 129, 0, 0, 0, 0, 333, 0, 0, 0,
	130, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	117, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 130, 265, 127, 131, 0,
	0, 129, 0, 0, 282, 0, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	117, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 130, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 130, 0, 226, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	0, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 130, 0, 223, 0, 0,
	0, 0, 0, 0, 0, 117, 113, 114, 115, 116,
	118, 119, 120, 128, 121, 126, 124, 125, 122, 123,
	130, 0, 127, 131, 0, 0, 129, 0, 0, 0,
	0, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 117, 0, 127, 131, 0,
	0, 279, 0, 0, 0, 0, 0, 0, 0, 0,
	130, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	117, 113, 114, 115, 116, 118, 119, 120, 128, 121,
	126, 124, 125, 122, 123, 130, 0, 127, 131, 0,
	0, 155, 0, 0, 0, 0, 113, 114, 115, 116,
	118, 119, 0, 128, 121, 126, 124, 125, 122, 123,
	117, 0, 127, 131, 0, 0, 129, 0, 130, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 113,
	114, 115, 116, 118, 0, 117, 128, 121, 126, 124,
	125, 122, 123, 0, 0, 127, 131, 213, 214, 129,
	0, 0, 93, 0, 0, 97, 98, 0, 0, 0,
	0, 0, 0, 54, 0, 76, 0, 95, 117, 71,
	72, 73, 74, 75, 59, 67, 94, 56, 218, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 211, 0, 0, 0, 0, 0, 62, 96, 0,
	0, 0, 63, 0, 65, 66, 0, 54, 64, 76,
	0, 0, 0, 71, 72, 73, 74, 75, 59, 67,
	0, 56, 218, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 0, 0, 61, 0, 0, 0, 0, 0,
	0, 62, 0, 0, 0, 0, 63, 0, 65, 66,
	0, 54, 64, 76, 0, 0, 0, 71, 72, 73,
	74, 75, 59, 67, 0, 56, 82, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 0, 0, 61, 0,
	0, 0, 0, 0, 0, 62, 0, 0, 0, 0,
	63, 0, 65, 66, 0, 138, 64, 76, 0, 0,
	202, 71, 72, 73, 201, 75, 59, 67, 0, 56,
	0, 0, 0, 0, 0, 55, 0, 0, 0, 0,
	0, 0, 61, 0, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 63, 0, 65, 66, 0, 54,
	64, 76, 149, 0, 0, 71, 72, 73, 74, 75,
	59, 67, 0, 56, 0, 0, 0, 0, 0, 55,
	0, 0, 0, 0, 0, 0, 61, 0, 0, 0,
	0, 0, 0, 62, 0, 0, 0, 0, 63, 0,
	65, 66, 0, 138, 64, 76, 0, 0, 0, 71,
	72, 73, 74, 75, 59, 67, 0, 56, 0, 0,
	0, 0, 0, 55, 0, 0, 0, 0, 0, 0,
	61, 0, 0, 0, 0, 0, 0, 62, 0, 0,
	0, 0, 63, 0, 65, 66, 0, 54, 64, 76,
	0, 0, 0, 71, 72, 73, 74, 75, 59, 67,
	0, 56, 0, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 0, 0, 61, 0, 0, 0, 0, 0,
	0, 62, 0, 0, 0, 0, 63, 0, 65, 66,
	0, 138, 64, 76, 0, 0, 0, 71, 72, 73,
	74, 75, 145, 67, 0, 56, 0, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 0, 0, 61, 0,
	0, 0, 0, 0, 0, 62, 0, 0, 0, 0,
	63, 0, 65, 66, 0, 138, 64, 76, 0, 0,
	0, 71, 72, 73, 74, 75, 144, 67, 0, 56,
	0, 0, 0, 0, 0, 55, 0, 0, 0, 0,
	0, 0, 61, 0, 0, 0, 0, 0, 0, 62,
	0, 0, 0, 0, 63, 0, 65, 66, 0, 0,
	64,
}
var yyPact = []int{

	448, -1000, -1000, 366, -1000, -1000, -1000, -1000, -1000, -1000,
	338, 297, 343, 309, 303, 6, 144, -1000, -1000, 142,
	286, 290, 337, 186, 303, 135, 278, 1549, 1373, -1000,
	-1000, -1000, -1000, 185, 38, 1287, -1000, -29, 183, -1000,
	281, 214, 1549, 332, 330, 278, -1000, 181, 316, 288,
	-1000, 1115, -1000, -1000, 106, 1505, 1505, -1000, -1000, 25,
	-1000, 1549, -43, 1637, 1593, 1505, 1505, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1461, -1000, -1000, 241,
	-1000, 103, -1000, 1180, -30, -1000, 151, -4, 151, 151,
	-1000, -85, -1000, 180, 305, 201, 179, 1505, 1505, 178,
	-31, -1000, 202, -1000, -1000, 238, 285, 99, 52, -1000,
	-34, -1000, 1549, 1505, 1505, 1505, 1505, 1505, 1505, 1505,
	1505, 1505, 1505, 1505, 1505, 1505, 1505, 1505, -24, 171,
	1417, 33, 268, -1000, 266, 234, 221, -1000, 96, -1000,
	1285, 21, 1505, 1075, -54, -55, 1035, 505, 465, -1000,
	257, 229, 1373, 160, -1000, 74, 151, 344, 151, 151,
	151, 253, 230, -1000, 305, -1000, 243, 200, -1000, 1115,
	1115, -1000, 159, -1000, 1549, -1000, -1000, 327, 157, 11,
	153, 151, 276, 104, 104, -20, -20, -20, -20, 1238,
	1205, -37, -37, -37, -37, -37, -37, -37, 1505, -1000,
	995, 110, 198, -1000, -10, -1000, -1000, 274, -1000, 42,
	1549, -1000, 19, 1329, 1329, 226, -1000, -1000, -1000, 1140,
	-1000, -53, 970, 1505, 1505, 1505, 1505, 147, 1505, 141,
	1505, -1000, 1549, -1000, -1000, -1000, -1000, 139, 38, -1000,
	111, 225, 152, 38, 137, 294, 1505, 1505, 38, 134,
	294, -1000, -1000, 216, 256, -36, -1000, 116, -38, 1549,
	-39, -1000, -1000, 1549, -37, -1000, 215, 251, -1000, -1000,
	-1000, -1000, 312, -1000, -1000, -1000, 9, -14, 1329, 44,
	-63, 1505, 1505, 930, 890, 850, 810, -69, 425, -74,
	385, -1000, 38, -1000, 125, 138, -1000, 38, 38, 294,
	124, 38, 294, 112, -1000, 294, 38, 1115, 1115, -1000,
	294, 38, 248, -1000, -1000, 92, -1000, -1000, -1000, 89,
	-21, 80, -1000, 246, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 1115, 785, -1000, 1505, 1505, -1000, 1505, 1505, -1000,
	1505, 1505, -1000, -1000, 138, -1000, 38, -1000, -1000, 38,
	294, -1000, 38, 294, 38, -1000, 38, -1000, -1000, -1000,
	325, 322, -19, -1000, 1505, 745, 705, 345, 665, 10,
	625, -1000, 38, -1000, -1000, 38, -1000, 38, -1000, -1000,
	116, 116, 1549, -1000, -1000, -1000, 1505, -1000, -1000, 1505,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -22, 585, 545,
	320, -1000, -1000, 116, -1000,
}
var yyPgo = []int{

	0, 469, 377, 4, 467, 465, 464, 461, 1, 16,
	460, 444, 443, 439, 320, 437, 373, 325, 436, 434,
	6, 433, 432, 430, 12, 429, 427, 0, 7, 424,
	2, 19, 421, 9, 10, 11, 420, 404, 401, 5,
	262, 397, 396, 393, 3, 389, 8, 388, 386, 385,
	384, 376, 18, 374,
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
	20, 39, 39, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 40, 40, 40, 41, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 44, 44, 45, 45, 31,
	31, 31, 31, 31, 31, 46, 46, 47, 47, 48,
	48, 43, 43, 43, 43, 43, 43, 43, 49, 49,
	50, 50, 52, 52, 53, 51, 51, 9, 9,
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
	1, 3, 4, 3, 3, 3, 3, 3, 3, 3,
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
	-12, 40, 27, 29, -15, 36, -16, 37, -22, 88,
	32, 33, -28, 53, -29, -31, 59, 59, 53, -13,
	-37, 44, 41, 29, 59, -16, -28, 53, -17, 46,
	-20, -27, -39, -40, 48, 68, 62, -41, -42, 59,
	-43, 75, 82, 87, 93, 89, 90, 60, -49, -50,
	-51, 54, 55, 56, 57, 58, 50, -23, -24, -25,
	-26, -20, 63, -27, 59, -30, 94, -32, 19, 96,
	97, 98, -33, 35, 59, 50, 81, 38, 39, 81,
	59, -38, 45, 57, -35, -36, -20, 30, 30, -17,
	59, -14, 41, 61, 62, 63, 64, 100, 65, 66,
	67, 69, 73, 74, 71, 72, 70, 77, 68, 81,
	50, 78, -3, 49, -52, -53, 60, -40, 48, -40,
	75, -20, 83, -27, 59, 59, -27, -27, -27, 51,
	-9, -20, 52, 35, 59, 81, 81, -31, 94, 19,
	96, -31, -31, 99, 59, -33, 57, 53, 59, -27,
	-27, 59, 81, 57, 52, 42, 43, 59, 53, 59,
	53, 81, -9, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, 77, 59,
	-27, 57, 53, 56, 68, 80, 79, 49, 49, 52,
	53, 76, -46, 32, 33, -47, -48, -20, 63, -27,
	76, -44, -27, 92, 91, 91, 92, 95, 91, 95,
	91, 51, 52, -24, 59, 63, -28, 35, 59, -30,
	-31, -31, -31, -34, 35, 59, 38, 39, -34, 35,
	59, -33, 51, 53, 57, 59, -35, 31, 59, 75,
	59, -28, -19, 47, -27, 51, 53, 57, 56, 80,
	79, -39, 48, -52, -20, 76, -46, -46, 52, 81,
	-45, 85, 84, -27, -27, -27, -27, 59, -27, 59,
	-27, -9, 59, -30, 35, 59, -30, -33, -34, 59,
	35, -34, 59, 35, -30, 59, -34, -27, -27, -30,
	59, -34, 57, 51, 51, 81, -8, 28, 59, 81,
	-9, 81, -20, 57, 51, 51, 76, 76, -46, 63,
	86, -27, -27, 86, 92, 92, 86, 91, 83, 86,
	91, 83, 86, -30, 59, -30, -33, -30, -30, -34,
	59, -30, -34, 59, -34, -30, -34, -30, 51, 59,
	59, 76, 59, 51, 83, -27, -27, -27, -27, -27,
	-27, -30, -33, -30, -30, -34, -30, -34, -30, -30,
	31, 31, 75, -44, 86, 86, 83, 86, 86, 83,
	86, 86, -30, -30, -30, -8, -8, -9, -27, -27,
	76, 86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 107, 0, 43, 105, 30, 0, 29, 2, 0,
	114, 0, 0, 0, 105, 0, 24, 0, 0, 31,
	32, 33, 46, 0, 48, 97, 179, 0, 0, 21,
	115, 0, 0, 0, 0, 24, 44, 0, 0, 0,
	106, 119, 120, 150, 0, 0, 0, 153, 154, 155,
	156, 0, 0, 0, 0, 0, 0, 191, 192, 193,
	194, 195, 196, 197, 198, 199, 0, 28, 34, 35,
	37, 38, 41, 119, 0, 49, 0, 0, 0, 0,
	94, 95, 98, 0, 100, 0, 0, 0, 0, 0,
	0, 116, 0, 117, 108, 109, 111, 0, 0, 22,
	0, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 200, 0, 202, 0, 151, 0, 152,
	0, 0, 0, 0, 155, 155, 0, 0, 0, 205,
	0, 207, 0, 0, 40, 0, 0, 50, 0, 0,
	0, 0, 0, 96, 99, 102, 0, 0, 184, 103,
	104, 18, 0, 118, 0, 112, 113, 8, 0, 0,
	0, 0, 26, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 0, 139,
	0, 198, 0, 144, 0, 146, 148, 121, 201, 0,
	0, 171, 0, 0, 0, 185, 187, 188, 189, 119,
	157, 177, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 206, 0, 36, 39, 42, 47, 0, 52, 53,
	56, 0, 0, 68, 0, 0, 0, 0, 80, 0,
	0, 101, 180, 0, 0, 0, 110, 0, 0, 0,
	0, 45, 25, 0, 138, 140, 0, 0, 145, 147,
	149, 122, 0, 203, 204, 172, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 208, 51, 55, 0, 58, 59, 62, 74, 0,
	0, 86, 0, 0, 71, 0, 70, 92, 93, 83,
	0, 82, 0, 182, 183, 0, 10, 16, 17, 0,
	0, 0, 27, 0, 142, 143, 173, 174, 186, 190,
	158, 178, 175, 159, 0, 0, 162, 0, 0, 166,
	0, 0, 170, 54, 57, 61, 63, 65, 75, 76,
	0, 87, 88, 0, 69, 73, 81, 85, 181, 19,
	9, 12, 0, 141, 0, 0, 0, 0, 0, 0,
	0, 60, 64, 66, 77, 78, 89, 90, 72, 84,
	0, 0, 0, 176, 160, 161, 0, 165, 164, 0,
	169, 168, 67, 79, 91, 11, 14, 0, 0, 0,
	13, 163, 167, 0, 15,
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
	}
	case 121:
		//line n1ql.y:1026
		{
	    logDebugGrammar("sub-query EXPRESSION")
	    
	}
	case 122:
		//line n1ql.y:1031
		{
	    logDebugGrammar("sub-query NESTED EXPRESSION")
	}
	case 123:
		//line n1ql.y:1037
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line n1ql.y:1045
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 125:
		//line n1ql.y:1053
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 126:
		//line n1ql.y:1061
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 127:
		//line n1ql.y:1069
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 128:
		//line n1ql.y:1077
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 129:
		//line n1ql.y:1085
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 130:
		//line n1ql.y:1093
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 131:
		//line n1ql.y:1101
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 132:
		//line n1ql.y:1109
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 133:
		//line n1ql.y:1117
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 134:
		//line n1ql.y:1125
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 135:
		//line n1ql.y:1133
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 136:
		//line n1ql.y:1141
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 137:
		//line n1ql.y:1149
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 138:
		//line n1ql.y:1157
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 139:
		//line n1ql.y:1165
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 140:
		//line n1ql.y:1173
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 141:
		//line n1ql.y:1181
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 142:
		//line n1ql.y:1188
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 143:
		//line n1ql.y:1196
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 144:
		//line n1ql.y:1203
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 145:
		//line n1ql.y:1210
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 146:
		//line n1ql.y:1217
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 147:
		//line n1ql.y:1224
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1231
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 149:
		//line n1ql.y:1238
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 150:
		//line n1ql.y:1245
		{
	
	}
	case 151:
		//line n1ql.y:1251
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 152:
		//line n1ql.y:1258
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 153:
		//line n1ql.y:1265
		{
	
	}
	case 154:
		//line n1ql.y:1270
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 155:
		//line n1ql.y:1276
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 156:
		//line n1ql.y:1282
		{
		logDebugGrammar("LITERAL")
	}
	case 157:
		//line n1ql.y:1286
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 158:
		//line n1ql.y:1290
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
	case 159:
		//line n1ql.y:1307
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 160:
		//line n1ql.y:1315
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 161:
		//line n1ql.y:1323
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 162:
		//line n1ql.y:1331
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 163:
		//line n1ql.y:1339
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 164:
		//line n1ql.y:1348
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 165:
		//line n1ql.y:1357
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 166:
		//line n1ql.y:1365
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 167:
		//line n1ql.y:1373
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 168:
		//line n1ql.y:1382
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 169:
		//line n1ql.y:1391
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 170:
		//line n1ql.y:1399
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 171:
		//line n1ql.y:1407
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 172:
		//line n1ql.y:1413
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 173:
		//line n1ql.y:1420
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 174:
		//line n1ql.y:1428
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 175:
		//line n1ql.y:1437
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 176:
		//line n1ql.y:1445
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
	case 177:
		//line n1ql.y:1459
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 178:
		//line n1ql.y:1463
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 179:
		//line n1ql.y:1469
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 180:
		//line n1ql.y:1475
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 181:
		//line n1ql.y:1482
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 182:
		//line n1ql.y:1489
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 183:
		//line n1ql.y:1497
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 184:
		//line n1ql.y:1504
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 185:
		//line n1ql.y:1515
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 186:
		//line n1ql.y:1520
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
	case 187:
		//line n1ql.y:1534
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 188:
		//line n1ql.y:1538
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 189:
		//line n1ql.y:1547
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 190:
		//line n1ql.y:1553
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 191:
		//line n1ql.y:1563
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 192:
		//line n1ql.y:1569
		{
		logDebugGrammar("NUMBER")
	}
	case 193:
		//line n1ql.y:1573
		{
		logDebugGrammar("OBJECT")
	}
	case 194:
		//line n1ql.y:1577
		{
		logDebugGrammar("ARRAY")
	}
	case 195:
		//line n1ql.y:1581
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 196:
		//line n1ql.y:1587
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 197:
		//line n1ql.y:1593
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 198:
		//line n1ql.y:1601
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 199:
		//line n1ql.y:1607
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 200:
		//line n1ql.y:1615
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 201:
		//line n1ql.y:1621
		{
		logDebugGrammar("OBJECT")
	}
	case 202:
		//line n1ql.y:1627
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 203:
		//line n1ql.y:1631
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 204:
		//line n1ql.y:1643
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 205:
		//line n1ql.y:1653
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 206:
		//line n1ql.y:1659
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 207:
		//line n1ql.y:1668
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 208:
		//line n1ql.y:1675
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
