
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

//line n1ql.y:12
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
const ORDER = 57380
const BY = 57381
const ASC = 57382
const DESC = 57383
const LIMIT = 57384
const OFFSET = 57385
const GROUP = 57386
const HAVING = 57387
const LBRACE = 57388
const RBRACE = 57389
const LBRACKET = 57390
const RBRACKET = 57391
const COMMA = 57392
const COLON = 57393
const TRUE = 57394
const FALSE = 57395
const NULL = 57396
const INT = 57397
const NUMBER = 57398
const IDENTIFIER = 57399
const STRING = 57400
const PLUS = 57401
const MINUS = 57402
const MULT = 57403
const DIV = 57404
const CONCAT = 57405
const AND = 57406
const OR = 57407
const NOT = 57408
const EQ = 57409
const NE = 57410
const GT = 57411
const GTE = 57412
const LT = 57413
const LTE = 57414
const LPAREN = 57415
const RPAREN = 57416
const LIKE = 57417
const IS = 57418
const VALUED = 57419
const MISSING = 57420
const DOT = 57421
const CASE = 57422
const WHEN = 57423
const THEN = 57424
const ELSE = 57425
const END = 57426
const ANY = 57427
const ALL = 57428
const OVER = 57429
const FIRST = 57430
const ARRAY = 57431
const IN = 57432
const SATISFIES = 57433
const EVERY = 57434
const UNNEST = 57435
const FOR = 57436
const MOD = 57437

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
	"OVER",
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

const yyNprod = 160
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1507

var yyAct = []int{

	50, 265, 199, 83, 34, 135, 140, 31, 190, 95,
	76, 73, 289, 74, 286, 125, 139, 68, 69, 70,
	71, 72, 56, 64, 85, 53, 196, 125, 220, 81,
	84, 52, 203, 45, 49, 79, 279, 270, 58, 246,
	127, 268, 264, 218, 202, 59, 163, 154, 97, 145,
	60, 90, 180, 62, 63, 121, 88, 61, 324, 297,
	276, 128, 131, 132, 133, 126, 104, 105, 106, 107,
	109, 110, 111, 119, 112, 117, 115, 116, 113, 114,
	185, 137, 118, 122, 275, 240, 120, 89, 316, 146,
	147, 317, 186, 88, 88, 85, 198, 309, 228, 298,
	125, 84, 108, 188, 187, 165, 166, 167, 168, 169,
	170, 171, 172, 173, 174, 175, 176, 177, 178, 179,
	164, 162, 182, 121, 89, 89, 197, 161, 200, 121,
	296, 121, 195, 85, 104, 105, 106, 107, 109, 84,
	236, 119, 81, 86, 106, 107, 109, 119, 79, 119,
	219, 122, 214, 217, 120, 142, 88, 122, 295, 122,
	120, 237, 120, 97, 134, 87, 225, 266, 181, 181,
	108, 230, 278, 216, 160, 137, 46, 143, 108, 35,
	159, 233, 35, 37, 239, 238, 259, 89, 32, 36,
	254, 252, 197, 197, 35, 229, 267, 227, 195, 195,
	241, 242, 248, 249, 250, 251, 224, 253, 215, 255,
	181, 153, 152, 149, 148, 101, 256, 91, 257, 82,
	43, 273, 262, 223, 155, 260, 94, 272, 261, 151,
	258, 235, 221, 150, 222, 211, 243, 213, 210, 271,
	51, 156, 141, 299, 197, 269, 294, 280, 281, 274,
	195, 263, 277, 212, 232, 121, 13, 209, 48, 93,
	47, 157, 158, 292, 293, 40, 104, 105, 106, 107,
	109, 110, 111, 119, 112, 117, 115, 116, 113, 114,
	103, 41, 118, 122, 301, 302, 120, 303, 304, 21,
	305, 306, 27, 123, 124, 25, 17, 327, 284, 29,
	30, 200, 108, 310, 102, 100, 3, 12, 10, 319,
	320, 12, 10, 99, 322, 308, 17, 323, 16, 121,
	17, 307, 16, 226, 98, 22, 321, 23, 42, 328,
	104, 105, 106, 107, 109, 110, 111, 119, 112, 117,
	115, 116, 113, 114, 19, 2, 118, 122, 136, 18,
	120, 67, 313, 66, 65, 314, 121, 26, 194, 193,
	245, 57, 55, 54, 92, 39, 108, 104, 105, 106,
	107, 109, 110, 111, 119, 112, 117, 115, 116, 113,
	114, 96, 44, 118, 122, 33, 78, 120, 77, 290,
	75, 28, 291, 121, 15, 231, 14, 24, 38, 20,
	11, 7, 9, 108, 104, 105, 106, 107, 109, 110,
	111, 119, 112, 117, 115, 116, 113, 114, 8, 6,
	118, 122, 5, 4, 120, 1, 287, 0, 0, 288,
	121, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	108, 104, 105, 106, 107, 109, 110, 111, 119, 112,
	117, 115, 116, 113, 114, 0, 0, 118, 122, 0,
	0, 120, 0, 0, 0, 0, 0, 121, 0, 0,
	0, 0, 208, 0, 0, 0, 207, 108, 104, 105,
	106, 107, 109, 110, 111, 119, 112, 117, 115, 116,
	113, 114, 0, 0, 118, 122, 0, 0, 120, 0,
	0, 0, 0, 0, 121, 0, 0, 0, 0, 206,
	0, 0, 0, 205, 108, 104, 105, 106, 107, 109,
	110, 111, 119, 112, 117, 115, 116, 113, 114, 0,
	0, 118, 122, 0, 0, 120, 0, 0, 0, 0,
	326, 121, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 108, 104, 105, 106, 107, 109, 110, 111, 119,
	112, 117, 115, 116, 113, 114, 0, 0, 118, 122,
	0, 0, 120, 0, 0, 0, 0, 325, 121, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 108, 104,
	105, 106, 107, 109, 110, 111, 119, 112, 117, 115,
	116, 113, 114, 0, 0, 118, 122, 0, 0, 120,
	0, 0, 0, 0, 318, 121, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 108, 104, 105, 106, 107,
	109, 110, 111, 119, 112, 117, 115, 116, 113, 114,
	0, 0, 118, 122, 0, 0, 120, 0, 0, 0,
	0, 315, 121, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 108, 104, 105, 106, 107, 109, 110, 111,
	119, 112, 117, 115, 116, 113, 114, 0, 0, 118,
	122, 0, 0, 120, 0, 0, 0, 0, 312, 121,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	104, 105, 106, 107, 109, 110, 111, 119, 112, 117,
	115, 116, 113, 114, 0, 0, 118, 122, 0, 0,
	120, 0, 0, 0, 0, 311, 121, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 108, 104, 105, 106,
	107, 109, 110, 111, 119, 112, 117, 115, 116, 113,
	114, 0, 0, 118, 122, 0, 0, 120, 0, 300,
	0, 0, 0, 121, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 108, 104, 105, 106, 107, 109, 110,
	111, 119, 112, 117, 115, 116, 113, 114, 0, 0,
	118, 122, 0, 0, 120, 0, 0, 0, 0, 285,
	121, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	108, 104, 105, 106, 107, 109, 110, 111, 119, 112,
	117, 115, 116, 113, 114, 0, 0, 118, 122, 0,
	0, 120, 0, 0, 0, 0, 0, 121, 0, 0,
	0, 0, 0, 283, 0, 0, 0, 108, 104, 105,
	106, 107, 109, 110, 111, 119, 112, 117, 115, 116,
	113, 114, 0, 0, 118, 122, 0, 0, 120, 0,
	0, 0, 0, 282, 121, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 108, 104, 105, 106, 107, 109,
	110, 111, 119, 112, 117, 115, 116, 113, 114, 0,
	0, 118, 122, 0, 0, 120, 0, 0, 247, 0,
	0, 121, 234, 0, 0, 0, 0, 0, 0, 0,
	0, 108, 104, 105, 106, 107, 109, 110, 111, 119,
	112, 117, 115, 116, 113, 114, 0, 0, 118, 122,
	0, 0, 120, 0, 0, 0, 0, 0, 121, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 108, 104,
	105, 106, 107, 109, 110, 111, 119, 112, 117, 115,
	116, 113, 114, 0, 0, 118, 122, 0, 0, 120,
	0, 0, 0, 0, 0, 121, 0, 0, 0, 0,
	0, 204, 0, 0, 0, 108, 104, 105, 106, 107,
	109, 110, 111, 119, 112, 117, 115, 116, 113, 114,
	0, 0, 118, 122, 0, 0, 120, 0, 0, 0,
	0, 0, 121, 0, 0, 0, 0, 0, 201, 0,
	0, 0, 108, 104, 105, 106, 107, 109, 110, 111,
	119, 112, 117, 115, 116, 113, 114, 0, 0, 118,
	122, 0, 0, 120, 0, 0, 0, 0, 0, 121,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 108,
	104, 105, 106, 107, 109, 110, 111, 119, 112, 117,
	115, 116, 113, 114, 0, 0, 118, 122, 0, 0,
	244, 0, 0, 0, 0, 0, 121, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 108, 104, 105, 106,
	107, 109, 110, 111, 119, 112, 117, 115, 116, 113,
	114, 0, 0, 118, 122, 0, 121, 144, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 104, 105, 106,
	107, 109, 110, 108, 119, 112, 117, 115, 116, 113,
	114, 0, 0, 118, 122, 191, 192, 120, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 73,
	0, 74, 0, 108, 0, 68, 69, 70, 71, 72,
	56, 64, 0, 53, 196, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 58, 189, 0, 0,
	0, 0, 0, 59, 0, 121, 0, 0, 60, 0,
	0, 62, 63, 0, 0, 61, 104, 105, 106, 107,
	109, 0, 0, 119, 112, 117, 115, 116, 113, 114,
	0, 0, 118, 122, 0, 73, 120, 74, 0, 0,
	0, 68, 69, 70, 71, 72, 56, 64, 0, 53,
	80, 0, 108, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 0, 62, 63, 0,
	73, 61, 74, 0, 0, 184, 68, 69, 70, 183,
	72, 56, 64, 0, 53, 0, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 0, 62, 63, 0, 73, 61, 74, 138, 0,
	0, 68, 69, 70, 71, 72, 56, 64, 0, 53,
	0, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 0, 62, 63, 0,
	73, 61, 74, 0, 0, 0, 68, 69, 70, 71,
	72, 56, 64, 0, 53, 0, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 0, 62, 63, 0, 73, 61, 74, 0, 0,
	0, 68, 69, 70, 71, 72, 130, 64, 0, 53,
	0, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 0, 0, 62, 63, 0,
	73, 61, 74, 0, 0, 0, 68, 69, 70, 71,
	72, 129, 64, 0, 53, 0, 0, 0, 0, 0,
	52, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	0, 0, 0, 0, 59, 0, 0, 0, 0, 60,
	0, 0, 62, 63, 0, 0, 61,
}
var yyPact = []int{

	282, -1000, -1000, 286, -1000, -1000, -1000, -1000, -1000, -1000,
	315, 251, 298, 259, 255, 267, 137, -1000, -1000, 132,
	223, 242, 299, 163, 255, 125, 214, 1324, 1189, -1000,
	-1000, -1000, 162, -63, 108, -1000, -28, 160, -1000, 216,
	171, 1324, 294, 283, 214, -1000, 158, 262, 241, -1000,
	974, -1000, 1324, 1324, -1000, -1000, 27, -1000, 1324, -41,
	1414, 1369, 1324, 1324, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 117, 1279, -1000, -1000, 192, -1000, 120,
	-1000, 1048, -30, -1000, 122, 157, 156, -1000, 178, 155,
	154, -32, -1000, 169, -1000, -1000, 191, 221, 123, 70,
	-1000, -33, -1000, 1324, 1324, 1324, 1324, 1324, 1324, 1324,
	1324, 1324, 1324, 1324, 1324, 1324, 1324, 1324, 1324, -23,
	153, 1234, 26, -1000, -1000, 1123, 22, 1324, 937, -46,
	-58, 900, 419, 382, -1000, 210, 188, 184, -1000, 204,
	187, 1189, 151, -1000, 112, 122, 8, 45, -62, -1000,
	183, 168, -1000, -1000, 149, -1000, 1324, -1000, -1000, 292,
	140, 25, 138, 122, 209, 83, 83, 81, 81, 81,
	81, 1157, 1078, 75, 75, 75, 75, 75, 75, 75,
	1324, -1000, 863, 180, 85, -1000, 107, -1000, -1000, -1000,
	11, -35, -35, 186, -1000, -1000, -1000, 1011, -1000, -44,
	826, 1324, 1324, 1324, 1324, 134, 1324, 133, 1324, -1000,
	23, 1324, -1000, 1324, -1000, -1000, -1000, -1000, 129, -1000,
	122, -1000, 173, 202, -37, -1000, 139, -38, 1324, -42,
	-1000, -1000, 1324, 75, -1000, 172, 200, -1000, -1000, -1000,
	-1000, 10, -14, -35, 111, -48, 1324, 1324, 789, 752,
	207, 715, -76, 345, -78, 308, -1000, -1000, -1000, -63,
	46, 197, -1000, -1000, 101, -1000, -1000, -1000, 73, -15,
	42, -1000, 194, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	974, 678, -1000, 1324, 1324, -1000, 1324, 1324, -1000, 1324,
	1324, -1000, -1000, -1000, -1000, -1000, 290, 284, 24, -1000,
	1324, 641, 604, 271, 567, 7, 530, 139, 139, 1324,
	-1000, -1000, -1000, 1324, -1000, -1000, 1324, -1000, -1000, -1000,
	-1000, -16, 493, 456, 266, -1000, -1000, 139, -1000,
}
var yyPgo = []int{

	0, 425, 345, 423, 422, 419, 418, 402, 1, 16,
	401, 400, 399, 398, 256, 397, 357, 260, 396, 395,
	6, 394, 391, 390, 10, 388, 386, 0, 7, 385,
	3, 4, 9, 381, 365, 364, 240, 363, 362, 361,
	2, 360, 8, 359, 358, 354, 353, 351, 5, 348,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 25,
	26, 26, 15, 15, 15, 18, 18, 28, 28, 30,
	30, 30, 30, 30, 30, 30, 29, 29, 29, 16,
	16, 12, 12, 32, 32, 33, 33, 33, 13, 13,
	13, 34, 35, 20, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 36, 36, 36, 37, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 40, 40, 41, 41,
	31, 31, 31, 31, 31, 31, 42, 42, 43, 43,
	44, 44, 39, 39, 39, 39, 39, 39, 39, 45,
	45, 46, 46, 48, 48, 49, 47, 47, 9, 9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 2,
	1, 3, 0, 2, 5, 2, 5, 1, 2, 2,
	4, 3, 5, 2, 4, 5, 1, 3, 2, 0,
	2, 0, 3, 1, 3, 1, 2, 2, 0, 1,
	2, 2, 2, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 6, 5, 5, 3, 4, 3, 4, 3,
	4, 1, 2, 2, 1, 1, 1, 1, 3, 5,
	5, 7, 7, 5, 9, 7, 7, 5, 9, 7,
	7, 5, 3, 4, 5, 5, 3, 5, 0, 2,
	1, 4, 6, 5, 5, 3, 1, 3, 1, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 3, 1, 3, 3, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, 24, -3, -4, -5, -10, -6, -7,
	26, -11, 25, -14, -18, -21, 36, 34, -2, 29,
	-12, 38, 27, 29, -15, 36, -16, 37, -22, 32,
	33, -28, 51, -29, -31, 57, 57, 51, -13, -34,
	42, 39, 29, 57, -16, -28, 51, -17, 44, -20,
	-27, -36, 66, 60, -37, -38, 57, -39, 73, 80,
	85, 92, 88, 89, 58, -45, -46, -47, 52, 53,
	54, 55, 56, 46, 48, -23, -24, -25, -26, -20,
	61, -27, 57, -30, 93, 87, 35, 57, 48, 79,
	79, 57, -35, 43, 55, -32, -33, -20, 30, 30,
	-17, 57, -14, 39, 59, 60, 61, 62, 95, 63,
	64, 65, 67, 71, 72, 69, 70, 68, 75, 66,
	79, 48, 76, -36, -36, 73, -20, 81, -27, 57,
	57, -27, -27, -27, 47, -48, -49, 58, 49, -9,
	-20, 50, 35, 57, 79, 79, -31, -31, 57, 57,
	55, 51, 57, 57, 79, 55, 50, 40, 41, 57,
	51, 57, 51, 79, -9, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	75, 57, -27, 55, 51, 54, 66, 78, 77, 74,
	-42, 32, 33, -43, -44, -20, 61, -27, 74, -40,
	-27, 91, 90, 90, 91, 94, 90, 94, 90, 47,
	50, 51, 49, 50, -24, 57, 61, -28, 35, -30,
	90, 49, 51, 55, 57, -32, 31, 57, 73, 57,
	-28, -19, 45, -27, 49, 51, 55, 54, 78, 77,
	74, -42, -42, 50, 79, -41, 83, 82, -27, -27,
	-27, -27, 57, -27, 57, -27, -48, -20, -9, 57,
	-31, 55, 49, 49, 79, -8, 28, 57, 79, -9,
	79, -20, 55, 49, 49, 74, 74, -42, 61, 84,
	-27, -27, 84, 91, 91, 84, 90, 81, 84, 90,
	81, 84, -30, -30, 49, 57, 57, 74, 57, 49,
	81, -27, -27, -27, -27, -27, -27, 31, 31, 73,
	-40, 84, 84, 81, 84, 84, 81, 84, 84, -8,
	-8, -9, -27, -27, 74, 84, 84, 31, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 61, 0, 42, 59, 30, 0, 29, 2, 0,
	68, 0, 0, 0, 59, 0, 24, 0, 0, 31,
	32, 45, 0, 47, 56, 130, 0, 0, 21, 69,
	0, 0, 0, 0, 24, 43, 0, 0, 0, 60,
	73, 101, 0, 0, 104, 105, 106, 107, 0, 0,
	0, 0, 0, 0, 142, 143, 144, 145, 146, 147,
	148, 149, 150, 0, 0, 28, 33, 34, 36, 37,
	40, 73, 0, 48, 0, 0, 0, 58, 0, 0,
	0, 0, 70, 0, 71, 62, 63, 65, 0, 0,
	22, 0, 23, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 102, 103, 0, 0, 0, 0, 106,
	106, 0, 0, 0, 151, 0, 153, 0, 156, 0,
	158, 0, 0, 39, 0, 0, 49, 53, 130, 57,
	0, 0, 135, 18, 0, 72, 0, 66, 67, 8,
	0, 0, 0, 0, 26, 74, 75, 76, 77, 78,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	0, 90, 0, 149, 0, 95, 0, 97, 99, 122,
	0, 0, 0, 136, 138, 139, 140, 73, 108, 128,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 152,
	0, 0, 157, 0, 35, 38, 41, 46, 0, 51,
	0, 131, 0, 0, 0, 64, 0, 0, 0, 0,
	44, 25, 0, 89, 91, 0, 0, 96, 98, 100,
	123, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 154, 155, 159, 50,
	54, 0, 133, 134, 0, 10, 16, 17, 0, 0,
	0, 27, 0, 93, 94, 124, 125, 137, 141, 109,
	129, 126, 110, 0, 0, 113, 0, 0, 117, 0,
	0, 121, 52, 55, 132, 19, 9, 12, 0, 92,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	127, 111, 112, 0, 116, 115, 0, 120, 119, 11,
	14, 0, 0, 0, 13, 114, 118, 0, 15,
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
	92, 93, 94, 95,
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
		//line n1ql.y:53
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line n1ql.y:57
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line n1ql.y:63
		{
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line n1ql.y:67
		{
	}
	case 5:
		//line n1ql.y:70
		{
		logDebugGrammar("STMT - DROP INDEX")
	}
	case 6:
		//line n1ql.y:77
		{
		logDebugGrammar("STMT - CREATE PRIMARY INDEX")
	}
	case 7:
		//line n1ql.y:81
		{
		logDebugGrammar("STMT - CREATE SECONDARY INDEX")
	}
	case 8:
		//line n1ql.y:87
		{
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 9:
		//line n1ql.y:95
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
		//line n1ql.y:105
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
		//line n1ql.y:115
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
		//line n1ql.y:129
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
		//line n1ql.y:141
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
		//line n1ql.y:155
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
		//line n1ql.y:169
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
		//line n1ql.y:187
		{
		parsingStack.Push("view")
	}
	case 17:
		//line n1ql.y:191
		{
		parsingStack.Push(yyS[yypt-0].s)
	}
	case 18:
		//line n1ql.y:197
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 19:
		//line n1ql.y:206
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
		//line n1ql.y:220
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 21:
		//line n1ql.y:226
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND")
	}
	case 22:
		//line n1ql.y:233
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 23:
		//line n1ql.y:237
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 24:
		//line n1ql.y:243
		{
	}
	case 25:
		//line n1ql.y:246
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
		//line n1ql.y:258
		{
	}
	case 27:
		//line n1ql.y:261
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
		//line n1ql.y:274
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 29:
		//line n1ql.y:280
		{
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 30:
		//line n1ql.y:286
		{
	}
	case 31:
		//line n1ql.y:289
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
		//line n1ql.y:299
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
		//line n1ql.y:311
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
		//line n1ql.y:325
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 35:
		//line n1ql.y:330
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
		//line n1ql.y:343
		{
		logDebugGrammar("RESULT STAR")
	}
	case 37:
		//line n1ql.y:347
		{
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 38:
		//line n1ql.y:354
		{
		logDebugGrammar("RESULT EXPR AS ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 39:
		//line n1ql.y:361
		{
		logDebugGrammar("RESULT EXPR ID")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 40:
		//line n1ql.y:370
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 41:
		//line n1ql.y:376
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 42:
		//line n1ql.y:385
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 43:
		//line n1ql.y:389
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
		//line n1ql.y:400
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
		//line n1ql.y:414
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
	case 46:
		//line n1ql.y:425
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
		//line n1ql.y:439
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 48:
		//line n1ql.y:443
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 49:
		//line n1ql.y:454
		{
	    logDebugGrammar("UNNEST")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:""})
	}
	case 50:
		//line n1ql.y:461
		{
	    logDebugGrammar("UNNEST AS")
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As:yyS[yypt-0].s})
	}
	case 51:
		//line n1ql.y:468
		{
	    logDebugGrammar("UNNEST nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: "", Over:rest})
	}
	case 52:
		//line n1ql.y:475
		{
	    logDebugGrammar("UNNEST AS nested")
	    rest := parsingStack.Pop().(*ast.From)
	    proj := parsingStack.Pop().(ast.Expression)
	    parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-1].s, Over:rest})
	}
	case 53:
		//line n1ql.y:483
		{
		logDebugGrammar("OVER ")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: ""})
	}
	case 54:
		//line n1ql.y:489
		{
		logDebugGrammar("OVER IN")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s})
	}
	case 55:
		//line n1ql.y:495
		{
		logDebugGrammar("OVER IN nested")
		rest := parsingStack.Pop().(*ast.From)
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Over:rest})
	}
	case 56:
		//line n1ql.y:504
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 57:
		//line n1ql.y:510
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 58:
		//line n1ql.y:517
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE ID")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 59:
		//line n1ql.y:526
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 60:
		//line n1ql.y:530
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
	case 62:
		//line n1ql.y:544
		{
	
	}
	case 63:
		//line n1ql.y:550
		{
	
	}
	case 64:
		//line n1ql.y:554
		{
	
	}
	case 65:
		//line n1ql.y:559
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
	case 66:
		//line n1ql.y:570
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
	case 67:
		//line n1ql.y:581
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
	case 68:
		//line n1ql.y:593
		{
	
	}
	case 69:
		//line n1ql.y:597
		{
	
	}
	case 70:
		//line n1ql.y:601
		{
	
	}
	case 71:
		//line n1ql.y:607
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
	case 72:
		//line n1ql.y:621
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
	case 73:
		//line n1ql.y:637
		{
		logDebugGrammar("EXPRESSION")
	}
	case 74:
		//line n1ql.y:642
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line n1ql.y:650
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line n1ql.y:658
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line n1ql.y:666
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line n1ql.y:674
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line n1ql.y:682
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line n1ql.y:690
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line n1ql.y:698
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line n1ql.y:706
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line n1ql.y:714
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line n1ql.y:722
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line n1ql.y:730
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line n1ql.y:738
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 87:
		//line n1ql.y:746
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line n1ql.y:754
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line n1ql.y:762
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line n1ql.y:770
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line n1ql.y:778
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 92:
		//line n1ql.y:786
		{
	    logDebugGrammar("EXPR COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 93:
		//line n1ql.y:793
		{
	    logDebugGrammar("EXPR COLON SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 94:
		//line n1ql.y:801
		{
	    logDebugGrammar("COLON EXPR SLICE BRACKET MEMBER")
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 95:
		//line n1ql.y:808
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 96:
		//line n1ql.y:815
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:822
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 98:
		//line n1ql.y:829
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 99:
		//line n1ql.y:836
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 100:
		//line n1ql.y:843
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 101:
		//line n1ql.y:850
		{
	
	}
	case 102:
		//line n1ql.y:856
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 103:
		//line n1ql.y:863
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 104:
		//line n1ql.y:870
		{
	
	}
	case 105:
		//line n1ql.y:875
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 106:
		//line n1ql.y:881
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line n1ql.y:887
		{
		logDebugGrammar("LITERAL")
	}
	case 108:
		//line n1ql.y:891
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 109:
		//line n1ql.y:895
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
	case 110:
		//line n1ql.y:912
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 111:
		//line n1ql.y:920
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 112:
		//line n1ql.y:928
		{
	    logDebugGrammar("ANY IN SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-5].s)
	    parsingStack.Push(collectionAny)
	}
	case 113:
		//line n1ql.y:936
		{
	    logDebugGrammar("ANY SATISFIES")
	    condition := parsingStack.Pop().(ast.Expression)
	    sub := parsingStack.Pop().(ast.Expression)
	    collectionAny := ast.NewCollectionAllOperator(condition, sub, "")
	    parsingStack.Push(collectionAny)
	}
	case 114:
		//line n1ql.y:944
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 115:
		//line n1ql.y:953
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 116:
		//line n1ql.y:962
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 117:
		//line n1ql.y:970
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, "", output)
		parsingStack.Push(collectionFirst)
	}
	case 118:
		//line n1ql.y:978
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 119:
		//line n1ql.y:987
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 120:
		//line n1ql.y:996
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 121:
		//line n1ql.y:1004
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, "", output)
		parsingStack.Push(collectionArray)
	}
	case 122:
		//line n1ql.y:1012
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 123:
		//line n1ql.y:1018
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line n1ql.y:1025
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 125:
		//line n1ql.y:1033
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 126:
		//line n1ql.y:1042
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 127:
		//line n1ql.y:1050
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
	case 128:
		//line n1ql.y:1064
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 129:
		//line n1ql.y:1068
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 130:
		//line n1ql.y:1074
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 131:
		//line n1ql.y:1080
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 132:
		//line n1ql.y:1087
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v-%v]", yyS[yypt-5].s,yyS[yypt-3].n, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-3].n)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 133:
		//line n1ql.y:1094
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER - %v[%v:]", yyS[yypt-4].s, yyS[yypt-2].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-2].n)), ast.NewLiteralNumber(float64(0)))
	    parsingStack.Push(thisExpression)
	
	}
	case 134:
		//line n1ql.y:1102
		{
	    logDebugGrammar("PATH SLICE BRACKET MEMBER -%v[:%v]", yyS[yypt-4].s, yyS[yypt-1].n)
	    left := parsingStack.Pop()
	    thisExpression := ast.NewBracketSliceMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(0)), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
	    parsingStack.Push(thisExpression)
	}
	case 135:
		//line n1ql.y:1109
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 136:
		//line n1ql.y:1120
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 137:
		//line n1ql.y:1125
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
	case 138:
		//line n1ql.y:1139
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 139:
		//line n1ql.y:1143
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 140:
		//line n1ql.y:1152
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 141:
		//line n1ql.y:1158
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 142:
		//line n1ql.y:1168
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 143:
		//line n1ql.y:1174
		{
		logDebugGrammar("NUMBER")
	}
	case 144:
		//line n1ql.y:1178
		{
		logDebugGrammar("OBJECT")
	}
	case 145:
		//line n1ql.y:1182
		{
		logDebugGrammar("ARRAY")
	}
	case 146:
		//line n1ql.y:1186
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 147:
		//line n1ql.y:1192
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 148:
		//line n1ql.y:1198
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 149:
		//line n1ql.y:1206
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 150:
		//line n1ql.y:1212
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 151:
		//line n1ql.y:1220
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 152:
		//line n1ql.y:1226
		{
		logDebugGrammar("OBJECT")
	}
	case 153:
		//line n1ql.y:1232
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 154:
		//line n1ql.y:1236
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 155:
		//line n1ql.y:1248
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 156:
		//line n1ql.y:1258
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 157:
		//line n1ql.y:1264
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 158:
		//line n1ql.y:1273
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 159:
		//line n1ql.y:1280
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
