
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
	-1, 277,
	19, 69,
	96, 69,
	97, 69,
	98, 69,
	-2, 71,
	-1, 278,
	19, 70,
	96, 70,
	97, 70,
	98, 70,
	-2, 72,
}

const yyNprod = 174
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1666

var yyAct = []int{

	50, 285, 209, 91, 84, 230, 83, 148, 31, 143,
	103, 85, 133, 34, 76, 309, 147, 86, 200, 306,
	299, 133, 135, 261, 252, 290, 288, 195, 213, 81,
	284, 86, 173, 164, 45, 49, 79, 212, 153, 196,
	280, 98, 228, 96, 97, 334, 190, 254, 253, 105,
	198, 197, 350, 320, 296, 94, 129, 94, 295, 255,
	208, 136, 139, 140, 141, 243, 134, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 286, 142, 126, 130, 133, 95, 128, 95, 341,
	145, 321, 342, 145, 88, 89, 90, 161, 162, 154,
	155, 85, 319, 191, 318, 116, 85, 298, 88, 89,
	90, 129, 287, 175, 176, 177, 178, 179, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 174, 127,
	192, 191, 172, 316, 207, 226, 210, 170, 171, 130,
	129, 205, 128, 169, 276, 150, 274, 269, 267, 251,
	81, 112, 113, 114, 115, 117, 35, 79, 127, 231,
	46, 229, 227, 244, 224, 242, 35, 37, 130, 151,
	235, 128, 92, 36, 105, 96, 97, 240, 239, 32,
	232, 225, 245, 233, 234, 35, 191, 94, 163, 116,
	160, 248, 157, 109, 99, 94, 93, 82, 43, 293,
	282, 238, 207, 207, 159, 292, 281, 165, 158, 205,
	205, 102, 263, 264, 265, 266, 51, 268, 95, 270,
	256, 257, 236, 250, 237, 221, 95, 258, 223, 272,
	271, 220, 166, 149, 277, 278, 275, 322, 317, 279,
	273, 294, 283, 129, 219, 222, 13, 247, 48, 101,
	47, 167, 168, 111, 40, 291, 114, 115, 117, 207,
	289, 127, 300, 301, 41, 21, 205, 96, 97, 131,
	132, 130, 233, 234, 128, 27, 25, 297, 17, 353,
	314, 312, 313, 333, 315, 332, 3, 12, 10, 241,
	12, 10, 116, 107, 110, 108, 17, 106, 16, 17,
	42, 16, 29, 30, 324, 325, 19, 326, 327, 156,
	328, 329, 22, 26, 23, 144, 67, 2, 66, 330,
	331, 18, 65, 204, 210, 203, 335, 260, 57, 55,
	54, 100, 39, 104, 345, 346, 344, 87, 44, 348,
	33, 78, 349, 77, 75, 129, 28, 15, 246, 14,
	24, 347, 38, 20, 11, 354, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	7, 9, 126, 130, 8, 6, 128, 5, 338, 4,
	1, 339, 0, 0, 129, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 116, 112, 113, 114, 115, 117,
	118, 119, 127, 120, 125, 123, 124, 121, 122, 0,
	0, 126, 130, 0, 0, 128, 0, 310, 0, 0,
	311, 0, 0, 129, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 116, 112, 113, 114, 115, 117, 118,
	119, 127, 120, 125, 123, 124, 121, 122, 0, 0,
	126, 130, 0, 0, 128, 0, 307, 0, 0, 308,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 116, 112, 113, 114, 115, 117, 118, 119,
	127, 120, 125, 123, 124, 121, 122, 0, 0, 126,
	130, 0, 0, 128, 0, 0, 0, 0, 0, 0,
	0, 129, 0, 218, 0, 0, 0, 217, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 0, 0, 126, 130,
	0, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 216, 0, 0, 0, 215, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 0, 0, 126, 130, 0,
	0, 128, 0, 0, 0, 0, 352, 0, 0, 129,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 116,
	112, 113, 114, 115, 117, 118, 119, 127, 120, 125,
	123, 124, 121, 122, 0, 0, 126, 130, 0, 0,
	128, 0, 0, 0, 0, 351, 0, 0, 129, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 116, 112,
	113, 114, 115, 117, 118, 119, 127, 120, 125, 123,
	124, 121, 122, 0, 0, 126, 130, 0, 0, 128,
	0, 0, 0, 0, 343, 0, 0, 129, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 116, 112, 113,
	114, 115, 117, 118, 119, 127, 120, 125, 123, 124,
	121, 122, 0, 0, 126, 130, 0, 0, 128, 0,
	0, 0, 0, 340, 0, 0, 129, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 337, 0, 0, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 336, 0, 0, 129, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 116, 112, 113, 114, 115, 117,
	118, 119, 127, 120, 125, 123, 124, 121, 122, 0,
	0, 126, 130, 0, 0, 128, 0, 323, 0, 0,
	0, 0, 0, 129, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 116, 112, 113, 114, 115, 117, 118,
	119, 127, 120, 125, 123, 124, 121, 122, 0, 0,
	126, 130, 0, 0, 128, 0, 0, 0, 0, 305,
	0, 0, 129, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 116, 112, 113, 114, 115, 117, 118, 119,
	127, 120, 125, 123, 124, 121, 122, 0, 0, 126,
	130, 0, 0, 128, 0, 0, 0, 0, 0, 0,
	0, 129, 0, 0, 304, 0, 0, 0, 0, 0,
	0, 116, 112, 113, 114, 115, 117, 118, 119, 127,
	120, 125, 123, 124, 121, 122, 0, 0, 126, 130,
	0, 0, 128, 0, 0, 0, 0, 0, 0, 0,
	129, 0, 0, 303, 0, 0, 0, 0, 0, 0,
	116, 112, 113, 114, 115, 117, 118, 119, 127, 120,
	125, 123, 124, 121, 122, 0, 0, 126, 130, 0,
	0, 128, 0, 0, 0, 0, 302, 0, 0, 129,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 116,
	112, 113, 114, 115, 117, 118, 119, 127, 120, 125,
	123, 124, 121, 122, 0, 0, 126, 130, 0, 0,
	128, 0, 0, 262, 0, 0, 0, 0, 129, 249,
	0, 0, 0, 0, 0, 0, 0, 0, 116, 112,
	113, 114, 115, 117, 118, 119, 127, 120, 125, 123,
	124, 121, 122, 0, 0, 126, 130, 0, 0, 128,
	0, 0, 0, 0, 0, 0, 0, 129, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 116, 112, 113,
	114, 115, 117, 118, 119, 127, 120, 125, 123, 124,
	121, 122, 0, 0, 126, 130, 0, 0, 128, 0,
	0, 0, 0, 0, 0, 0, 129, 0, 0, 214,
	0, 0, 0, 0, 0, 0, 116, 112, 113, 114,
	115, 117, 118, 119, 127, 120, 125, 123, 124, 121,
	122, 0, 0, 126, 130, 0, 0, 128, 0, 0,
	0, 0, 0, 0, 0, 129, 0, 0, 211, 0,
	0, 0, 0, 0, 0, 116, 112, 113, 114, 115,
	117, 118, 119, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 0, 0, 0,
	0, 0, 0, 0, 129, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 116, 112, 113, 114, 115, 117,
	118, 119, 127, 120, 125, 123, 124, 121, 122, 0,
	0, 126, 130, 0, 0, 259, 0, 0, 0, 0,
	0, 0, 0, 129, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 116, 112, 113, 114, 115, 117, 118,
	119, 127, 120, 125, 123, 124, 121, 122, 0, 0,
	126, 130, 0, 0, 152, 129, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 112, 113, 114, 115,
	117, 118, 116, 127, 120, 125, 123, 124, 121, 122,
	0, 0, 126, 130, 0, 0, 128, 129, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 112, 113,
	114, 115, 117, 0, 116, 127, 120, 125, 123, 124,
	121, 122, 0, 0, 126, 130, 201, 202, 128, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 73, 0, 74, 0, 116, 0, 68, 69,
	70, 71, 72, 56, 64, 0, 53, 206, 0, 0,
	0, 0, 52, 0, 0, 0, 0, 0, 0, 58,
	199, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 60, 0, 62, 63, 0, 73, 61, 74, 0,
	0, 0, 68, 69, 70, 71, 72, 56, 64, 0,
	53, 206, 0, 0, 0, 0, 52, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 0, 0, 60, 0, 62, 63, 0,
	73, 61, 74, 0, 0, 0, 68, 69, 70, 71,
	72, 56, 64, 0, 53, 80, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 62, 63, 0, 73, 61, 74, 0, 0, 194,
	68, 69, 70, 193, 72, 56, 64, 0, 53, 0,
	0, 0, 0, 0, 52, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 0, 0, 0, 0, 59, 0,
	0, 0, 0, 60, 0, 62, 63, 0, 73, 61,
	74, 146, 0, 0, 68, 69, 70, 71, 72, 56,
	64, 0, 53, 0, 0, 0, 0, 0, 52, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 0, 0,
	0, 0, 59, 0, 0, 0, 0, 60, 0, 62,
	63, 0, 73, 61, 74, 0, 0, 0, 68, 69,
	70, 71, 72, 56, 64, 0, 53, 0, 0, 0,
	0, 0, 52, 0, 0, 0, 0, 0, 0, 58,
	0, 0, 0, 0, 0, 0, 59, 0, 0, 0,
	0, 60, 0, 62, 63, 0, 73, 61, 74, 0,
	0, 0, 68, 69, 70, 71, 72, 138, 64, 0,
	53, 0, 0, 0, 0, 0, 52, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 0, 0, 0, 0,
	59, 0, 0, 0, 0, 60, 0, 62, 63, 0,
	73, 61, 74, 0, 0, 0, 68, 69, 70, 71,
	72, 137, 64, 0, 53, 0, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 62, 63, 0, 0, 61,
}
var yyPact = []int{

	262, -1000, -1000, 265, -1000, -1000, -1000, -1000, -1000, -1000,
	277, 225, 285, 240, 238, 270, 126, -1000, -1000, 114,
	210, 223, 271, 139, 238, 107, 202, 1484, 1352, -1000,
	-1000, -1000, 138, 12, 137, -1000, -40, 135, -1000, 204,
	154, 1484, 267, 263, 202, -1000, 134, 244, 212, -1000,
	1075, -1000, 1484, 1484, -1000, -1000, 10, -1000, 1484, -61,
	1572, 1528, 1484, 1484, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 33, 1440, -1000, -1000, 181, -1000, 110,
	-1000, 1153, -43, -1000, -1000, 97, 97, 290, -1000, -1000,
	-1000, -1000, 133, -1000, 151, 131, 1484, 1484, 129, -48,
	-1000, 150, -1000, -1000, 180, 209, 84, 79, -1000, -49,
	-1000, 1484, 1484, 1484, 1484, 1484, 1484, 1484, 1484, 1484,
	1484, 1484, 1484, 1484, 1484, 1484, 1484, -31, 127, 1396,
	-29, -1000, -1000, 1264, -16, 1484, 1036, -54, -63, 997,
	451, 412, -1000, 195, 179, 172, -1000, 194, 176, 1352,
	122, -1000, 72, 97, 7, 145, 97, -1000, 171, 144,
	-1000, 1075, 1075, -1000, 119, -1000, 1484, -1000, -1000, 258,
	106, -10, 104, 97, 200, 193, 193, 61, 61, 61,
	61, 1217, 1185, 90, 90, 90, 90, 90, 90, 90,
	1484, -1000, 958, 170, 92, -1000, -32, -1000, -1000, -1000,
	-17, 1308, 1308, 175, -1000, -1000, -1000, 1114, -1000, -62,
	919, 1484, 1484, 1484, 1484, 89, 1484, 88, 1484, -1000,
	30, 1484, -1000, 1484, -1000, -1000, -1000, -1000, 87, -1000,
	-1000, -2, 85, 1484, 1484, 5, -1000, 149, 191, -51,
	-1000, 53, -55, 1484, -56, -1000, -1000, 1484, 90, -1000,
	148, 190, -1000, -1000, -1000, -1000, -18, -22, 1308, 44,
	-66, 1484, 1484, 880, 841, 802, 763, -72, 373, -76,
	334, -1000, -1000, -1000, -83, -1000, 234, 1075, 1075, -2,
	74, 187, -1000, -1000, 45, -1000, -1000, -1000, 43, -23,
	32, -1000, 186, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	1075, 724, -1000, 1484, 1484, -1000, 1484, 1484, -1000, 1484,
	1484, -1000, -1000, -1000, -2, -1000, 229, -1000, -1000, 254,
	252, -30, -1000, 1484, 685, 646, 295, 607, 6, 568,
	-1000, -2, 53, 53, 1484, -1000, -1000, -1000, 1484, -1000,
	-1000, 1484, -1000, -1000, -1000, -1000, -1000, -24, 529, 490,
	248, -1000, -1000, 53, -1000,
}
var yyPgo = []int{

	0, 380, 317, 379, 377, 375, 374, 371, 1, 16,
	370, 354, 353, 352, 246, 350, 313, 250, 349, 348,
	7, 347, 346, 344, 14, 343, 341, 0, 8, 340,
	6, 4, 13, 5, 3, 337, 10, 333, 332, 331,
	216, 330, 329, 328, 2, 327, 18, 325, 323, 322,
	318, 316, 9, 315,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 25,
	26, 26, 15, 15, 15, 18, 18, 28, 28, 28,
	30, 30, 30, 30, 31, 31, 31, 31, 31, 31,
	31, 31, 35, 35, 35, 29, 29, 29, 29, 34,
	34, 33, 33, 16, 16, 12, 12, 36, 36, 37,
	37, 37, 13, 13, 13, 38, 39, 20, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 40, 40, 40, 41,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	44, 44, 45, 45, 32, 32, 32, 32, 32, 32,
	46, 46, 47, 47, 48, 48, 43, 43, 43, 43,
	43, 43, 43, 49, 49, 50, 50, 52, 52, 53,
	51, 51, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 0, 2, 5, 2, 5, 1, 2, 2,
	2, 4, 3, 5, 3, 4, 5, 6, 4, 5,
	6, 7, 1, 1, 1, 1, 2, 3, 2, 2,
	2, 2, 2, 0, 2, 0, 3, 1, 3, 1,
	2, 2, 0, 1, 2, 2, 2, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 6, 5, 5, 3,
	4, 3, 4, 3, 4, 1, 2, 2, 1, 1,
	1, 1, 3, 5, 5, 7, 7, 5, 9, 7,
	7, 5, 9, 7, 7, 5, 3, 4, 5, 5,
	3, 5, 0, 2, 1, 4, 6, 5, 5, 3,
	1, 3, 1, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 1, 3, 3,
	2, 3, 1, 3,
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
	63, -27, 59, -30, -31, 94, 19, -35, 96, 97,
	98, -34, 35, 59, 50, 81, 38, 39, 81, 59,
	-39, 45, 57, -36, -37, -20, 30, 30, -17, 59,
	-14, 41, 61, 62, 63, 64, 99, 65, 66, 67,
	69, 73, 74, 71, 72, 70, 77, 68, 81, 50,
	78, -40, -40, 75, -20, 83, -27, 59, 59, -27,
	-27, -27, 49, -52, -53, 60, 51, -9, -20, 52,
	35, 59, 81, 81, -32, -32, 19, 59, 57, 53,
	59, -27, -27, 59, 81, 57, 52, 42, 43, 59,
	53, 59, 53, 81, -9, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	77, 59, -27, 57, 53, 56, 68, 80, 79, 76,
	-46, 32, 33, -47, -48, -20, 63, -27, 76, -44,
	-27, 92, 91, 91, 92, 95, 91, 95, 91, 49,
	52, 53, 51, 52, -24, 59, 63, -28, 35, -30,
	-33, -34, 35, 38, 39, -32, 51, 53, 57, 59,
	-36, 31, 59, 75, 59, -28, -19, 47, -27, 51,
	53, 57, 56, 80, 79, 76, -46, -46, 52, 81,
	-45, 85, 84, -27, -27, -27, -27, 59, -27, 59,
	-27, -52, -20, -9, 59, -31, 59, -27, -27, -34,
	35, 57, 51, 51, 81, -8, 28, 59, 81, -9,
	81, -20, 57, 51, 51, 76, 76, -46, 63, 86,
	-27, -27, 86, 92, 92, 86, 91, 83, 86, 91,
	83, 86, -30, -33, -34, -31, 59, 51, 59, 59,
	76, 59, 51, 83, -27, -27, -27, -27, -27, -27,
	-31, -34, 31, 31, 75, -44, 86, 86, 83, 86,
	86, 83, 86, 86, -31, -8, -8, -9, -27, -27,
	76, 86, 86, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 75, 0, 42, 73, 30, 0, 29, 2, 0,
	82, 0, 0, 0, 73, 0, 24, 0, 0, 31,
	32, 45, 0, 47, 65, 144, 0, 0, 21, 83,
	0, 0, 0, 0, 24, 43, 0, 0, 0, 74,
	87, 115, 0, 0, 118, 119, 120, 121, 0, 0,
	0, 0, 0, 0, 156, 157, 158, 159, 160, 161,
	162, 163, 164, 0, 0, 28, 33, 34, 36, 37,
	40, 87, 0, 48, 49, 0, 0, 0, 62, 63,
	64, 66, 0, 68, 0, 0, 0, 0, 0, 0,
	84, 0, 85, 76, 77, 79, 0, 0, 22, 0,
	23, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 116, 117, 0, 0, 0, 0, 120, 120, 0,
	0, 0, 165, 0, 167, 0, 170, 0, 172, 0,
	0, 39, 0, 0, 50, 0, 0, 67, 0, 0,
	149, 69, 70, 18, 0, 86, 0, 80, 81, 8,
	0, 0, 0, 0, 26, 88, 89, 90, 91, 92,
	93, 94, 95, 96, 97, 98, 99, 100, 101, 102,
	0, 104, 0, 163, 0, 109, 0, 111, 113, 136,
	0, 0, 0, 150, 152, 153, 154, 87, 122, 142,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 166,
	0, 0, 171, 0, 35, 38, 41, 46, 0, 52,
	54, 0, 0, 0, 0, 0, 145, 0, 0, 0,
	78, 0, 0, 0, 0, 44, 25, 0, 103, 105,
	0, 0, 110, 112, 114, 137, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 168, 169, 173, 51, 55, 0, -2, -2, 58,
	0, 0, 147, 148, 0, 10, 16, 17, 0, 0,
	0, 27, 0, 107, 108, 138, 139, 151, 155, 123,
	143, 140, 124, 0, 0, 127, 0, 0, 131, 0,
	0, 135, 53, 56, 0, 59, 0, 146, 19, 9,
	12, 0, 106, 0, 0, 0, 0, 0, 0, 0,
	57, 60, 0, 0, 0, 141, 125, 126, 0, 130,
	129, 0, 134, 133, 61, 11, 14, 0, 0, 0,
	13, 128, 132, 0, 15,
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
		last.Join = rest
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
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: "", Over:rest})
	}
	case 53:
		//line n1ql.y:487
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 54:
		//line n1ql.y:497
		{
	    logDebugGrammar("JOIN KEY")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:""})
	}
	case 55:
		//line n1ql.y:503
		{
	    logDebugGrammar("JOIN KEY")
	}
	case 56:
		//line n1ql.y:507
		{
	    logDebugGrammar("JOIN AS KEY") 
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-1].s})
	
	}
	case 57:
		//line n1ql.y:514
		{
	    logDebugGrammar("JOIN AS KEY") 
	}
	case 58:
		//line n1ql.y:518
		{
	    logDebugGrammar("TYPE JOIN KEY")
	}
	case 59:
		//line n1ql.y:522
		{
	    logDebugGrammar("TYPE JOIN KEY")
	}
	case 60:
		//line n1ql.y:526
		{
	    logDebugGrammar("TYPE JOIN AS KEY")
	}
	case 61:
		//line n1ql.y:530
		{
	    logDebugGrammar("TYPE JOIN AS KEY")
	}
	case 62:
		//line n1ql.y:536
		{
	    logDebugGrammar("INNER")
	}
	case 63:
		//line n1ql.y:540
		{
	    logDebugGrammar("OUTER")
	}
	case 64:
		//line n1ql.y:544
		{
	    logDebugGrammar("OUTER")
	}
	case 65:
		//line n1ql.y:550
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 66:
		//line n1ql.y:556
		{
	    logDebugGrammar("PATH KEY(S) EXPRESSION")
	}
	case 67:
		//line n1ql.y:560
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 68:
		//line n1ql.y:567
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 69:
		//line n1ql.y:576
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
	case 70:
		//line n1ql.y:589
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
	case 71:
		//line n1ql.y:603
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
	case 72:
		//line n1ql.y:614
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
	case 73:
		//line n1ql.y:626
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 74:
		//line n1ql.y:630
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
	case 76:
		//line n1ql.y:644
		{
	
	}
	case 77:
		//line n1ql.y:650
		{
	
	}
	case 78:
		//line n1ql.y:654
		{
	
	}
	case 79:
		//line n1ql.y:659
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
	case 80:
		//line n1ql.y:670
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
	case 81:
		//line n1ql.y:681
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
	case 82:
		//line n1ql.y:693
		{
	
	}
	case 83:
		//line n1ql.y:697
		{
	
	}
	case 84:
		//line n1ql.y:701
		{
	
	}
	case 85:
		//line n1ql.y:707
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
	case 86:
		//line n1ql.y:721
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
	case 87:
		//line n1ql.y:738
		{
		logDebugGrammar("EXPRESSION")
	}
	case 88:
		//line n1ql.y:743
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line n1ql.y:751
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line n1ql.y:759
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line n1ql.y:767
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line n1ql.y:775
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line n1ql.y:783
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line n1ql.y:791
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 95:
		//line n1ql.y:799
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line n1ql.y:807
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:815
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line n1ql.y:823
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line n1ql.y:831
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line n1ql.y:839
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line n1ql.y:847
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 102:
		//line n1ql.y:855
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line n1ql.y:863
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line n1ql.y:871
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 105:
		//line n1ql.y:879
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 106:
		//line n1ql.y:887
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 107:
		//line n1ql.y:894
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 108:
		//line n1ql.y:902
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 109:
		//line n1ql.y:909
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line n1ql.y:916
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 111:
		//line n1ql.y:923
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 112:
		//line n1ql.y:930
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 113:
		//line n1ql.y:937
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 114:
		//line n1ql.y:944
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line n1ql.y:951
		{
	
	}
	case 116:
		//line n1ql.y:957
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 117:
		//line n1ql.y:964
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 118:
		//line n1ql.y:971
		{
	
	}
	case 119:
		//line n1ql.y:976
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 120:
		//line n1ql.y:982
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 121:
		//line n1ql.y:988
		{
		logDebugGrammar("LITERAL")
	}
	case 122:
		//line n1ql.y:992
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 123:
		//line n1ql.y:996
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
	case 124:
		//line n1ql.y:1013
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 125:
		//line n1ql.y:1021
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 126:
		//line n1ql.y:1029
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 127:
		//line n1ql.y:1037
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 128:
		//line n1ql.y:1045
		{
		logDebugGrammar("FIRST FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 129:
		//line n1ql.y:1054
		{
		logDebugGrammar("FIRST IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 130:
		//line n1ql.y:1063
		{
		logDebugGrammar("FIRST FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 131:
		//line n1ql.y:1071
		{
		logDebugGrammar("FIRST IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 132:
		//line n1ql.y:1079
		{
		logDebugGrammar("ARRAY FOR IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 133:
		//line n1ql.y:1088
		{
		logDebugGrammar("ARRAY IN WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 134:
		//line n1ql.y:1097
		{
		logDebugGrammar("ARRAY FOR IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 135:
		//line n1ql.y:1105
		{
		logDebugGrammar("ARRAY IN")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 136:
		//line n1ql.y:1113
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 137:
		//line n1ql.y:1119
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 138:
		//line n1ql.y:1126
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 139:
		//line n1ql.y:1134
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 140:
		//line n1ql.y:1143
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 141:
		//line n1ql.y:1151
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
	case 142:
		//line n1ql.y:1165
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 143:
		//line n1ql.y:1169
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 144:
		//line n1ql.y:1175
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 145:
		//line n1ql.y:1181
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 146:
		//line n1ql.y:1188
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 147:
		//line n1ql.y:1195
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 148:
		//line n1ql.y:1203
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 149:
		//line n1ql.y:1210
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 150:
		//line n1ql.y:1221
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 151:
		//line n1ql.y:1226
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
	case 152:
		//line n1ql.y:1240
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 153:
		//line n1ql.y:1244
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 154:
		//line n1ql.y:1253
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 155:
		//line n1ql.y:1259
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 156:
		//line n1ql.y:1269
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 157:
		//line n1ql.y:1275
		{
		logDebugGrammar("NUMBER")
	}
	case 158:
		//line n1ql.y:1279
		{
		logDebugGrammar("OBJECT")
	}
	case 159:
		//line n1ql.y:1283
		{
		logDebugGrammar("ARRAY")
	}
	case 160:
		//line n1ql.y:1287
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 161:
		//line n1ql.y:1293
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 162:
		//line n1ql.y:1299
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 163:
		//line n1ql.y:1307
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 164:
		//line n1ql.y:1313
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 165:
		//line n1ql.y:1321
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 166:
		//line n1ql.y:1327
		{
		logDebugGrammar("OBJECT")
	}
	case 167:
		//line n1ql.y:1333
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 168:
		//line n1ql.y:1337
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 169:
		//line n1ql.y:1349
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 170:
		//line n1ql.y:1359
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 171:
		//line n1ql.y:1365
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 172:
		//line n1ql.y:1374
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 173:
		//line n1ql.y:1381
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
