
//line unql.y:2
package goyacc
import __yyfmt__ "fmt"
//line unql.y:2
		import "log"

func logDebugGrammar(format string, v ...interface{}) {
    if DebugGrammar && len(v) > 0 {
        log.Printf("DEBUG GRAMMAR " + format, v)
    } else if DebugGrammar {
        log.Printf("DEBUG GRAMMAR " + format)
    }
}

//line unql.y:14
type yySymType struct {
	yys int 
s string 
n int
f float64}

const SELECT = 57346
const AS = 57347
const FROM = 57348
const WHERE = 57349
const ORDER = 57350
const BY = 57351
const ASC = 57352
const DESC = 57353
const LIMIT = 57354
const OFFSET = 57355
const LBRACE = 57356
const RBRACE = 57357
const LBRACKET = 57358
const RBRACKET = 57359
const COMMA = 57360
const COLON = 57361
const DQUOTE = 57362
const TRUE = 57363
const FALSE = 57364
const NULL = 57365
const INT = 57366
const NUMBER = 57367
const IDENTIFIER = 57368
const STRING = 57369
const PLUS = 57370
const MINUS = 57371
const MULT = 57372
const DIV = 57373
const CONCAT = 57374
const AND = 57375
const OR = 57376
const NOT = 57377
const EQ = 57378
const NE = 57379
const GT = 57380
const GTE = 57381
const LT = 57382
const LTE = 57383
const LPAREN = 57384
const RPAREN = 57385
const LIKE = 57386
const IS = 57387
const VALUED = 57388
const MISSING = 57389
const DOT = 57390
const MOD = 57391

var yyToknames = []string{
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
	"LBRACE",
	"RBRACE",
	"LBRACKET",
	"RBRACKET",
	"COMMA",
	"COLON",
	"DQUOTE",
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

const yyNprod = 84
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 281

var yyAct = []int{

	16, 70, 81, 76, 72, 125, 68, 105, 18, 49,
	50, 51, 52, 54, 55, 111, 64, 57, 62, 60,
	61, 58, 59, 107, 13, 63, 64, 69, 127, 126,
	53, 71, 89, 118, 74, 108, 88, 89, 77, 122,
	87, 45, 83, 74, 84, 80, 110, 109, 18, 115,
	119, 90, 91, 92, 93, 94, 95, 96, 97, 98,
	99, 100, 101, 102, 103, 104, 117, 106, 65, 124,
	116, 86, 114, 19, 46, 113, 79, 120, 121, 41,
	49, 50, 51, 52, 54, 55, 56, 64, 57, 62,
	60, 61, 58, 59, 66, 67, 63, 40, 9, 43,
	112, 53, 11, 85, 47, 7, 73, 123, 30, 29,
	28, 65, 25, 23, 22, 78, 129, 39, 77, 128,
	83, 130, 131, 49, 50, 51, 52, 54, 55, 56,
	64, 57, 62, 60, 61, 58, 59, 65, 82, 63,
	44, 15, 14, 112, 53, 12, 6, 42, 10, 49,
	50, 51, 52, 54, 55, 56, 64, 57, 62, 60,
	61, 58, 59, 5, 4, 63, 38, 8, 3, 48,
	53, 49, 50, 51, 52, 54, 2, 1, 64, 57,
	62, 60, 61, 58, 59, 0, 0, 63, 0, 0,
	0, 0, 53, 49, 50, 51, 52, 54, 0, 0,
	64, 0, 0, 0, 0, 0, 0, 0, 36, 0,
	37, 0, 0, 0, 53, 31, 32, 33, 34, 35,
	24, 27, 0, 21, 17, 0, 0, 0, 0, 20,
	36, 0, 37, 75, 0, 0, 26, 31, 32, 33,
	34, 35, 24, 27, 0, 21, 0, 0, 0, 0,
	0, 20, 36, 0, 37, 0, 0, 0, 26, 31,
	32, 33, 34, 35, 24, 27, 0, 21, 0, 0,
	0, 0, 0, 20, 0, 0, 0, 0, 0, 0,
	26,
}
var yyPact = []int{

	101, -1000, -1000, 90, -1000, 96, 194, -1000, 85, 70,
	92, 15, -1000, -1000, 56, -1000, 99, -1000, 121, -1000,
	238, 238, -1000, -39, -1000, -1000, 238, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 16, 216, -1000, 63,
	21, 238, -1000, 238, -1000, 98, 194, 14, 6, 238,
	238, 238, 238, 238, 238, 238, 238, 238, 238, 238,
	238, 238, 238, 238, -37, 238, -1000, -1000, 0, -28,
	95, -1000, 60, 54, 30, -1000, 53, 48, -1000, 9,
	-1000, -1000, 32, 67, -1000, 13, -1000, -1000, -1000, -1000,
	-9, -9, -9, -9, -9, -9, 143, -19, 165, 165,
	165, 165, 165, 165, 165, 238, 52, -1000, -18, -1000,
	-1000, -1000, 11, -1000, 7, 238, -1000, 238, -1000, 238,
	-1000, -1000, -1000, 165, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000,
}
var yyPgo = []int{

	0, 177, 176, 168, 167, 166, 164, 163, 148, 147,
	146, 145, 24, 142, 141, 0, 1, 140, 2, 138,
	117, 115, 73, 114, 113, 112, 110, 109, 108, 4,
	106, 3,
}
var yyR1 = []int{

	0, 1, 2, 3, 6, 7, 10, 11, 12, 12,
	13, 13, 13, 14, 14, 8, 8, 17, 17, 9,
	9, 4, 4, 18, 18, 19, 19, 19, 5, 5,
	5, 20, 21, 15, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 22, 22, 22, 23, 23, 23, 23,
	23, 23, 23, 24, 24, 24, 25, 25, 25, 25,
	25, 25, 25, 26, 26, 27, 27, 29, 29, 30,
	28, 28, 31, 31,
}
var yyR2 = []int{

	0, 1, 3, 1, 3, 2, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 0, 2, 1, 3, 0,
	2, 0, 3, 1, 3, 1, 2, 2, 0, 1,
	2, 2, 2, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 1, 2, 2, 1, 1, 3, 4, 3,
	4, 3, 4, 1, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 1, 3, 3,
	2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -6, -7, -10, 4, -4, 8,
	-8, 6, -11, -12, -13, -14, -15, 30, -16, -22,
	35, 29, -23, -24, 26, -25, 42, 27, -26, -27,
	-28, 21, 22, 23, 24, 25, 14, 16, -5, -20,
	12, 9, -9, 7, -17, 26, 18, 5, 48, 28,
	29, 30, 31, 49, 32, 33, 34, 36, 40, 41,
	38, 39, 37, 44, 35, 16, -22, -22, 45, -15,
	-16, 15, -29, -30, 27, 17, -31, -15, -21, 13,
	24, -18, -19, -15, -15, 5, -12, 26, 30, 26,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, 44, -16, 23, 35, 47,
	46, 43, 48, 15, 18, 19, 17, 18, 24, 18,
	10, 11, 26, -16, 17, 23, 47, 46, -29, -15,
	-31, -18,
}
var yyDef = []int{

	0, -2, 1, 21, 3, 15, 0, 6, 28, 0,
	19, 0, 5, 7, 8, 10, 11, 13, 33, 52,
	0, 0, 55, 56, 63, 64, 0, 66, 67, 68,
	69, 70, 71, 72, 73, 74, 0, 0, 2, 29,
	0, 0, 4, 0, 16, 17, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 53, 54, 0, 0,
	33, 75, 0, 77, 0, 80, 0, 82, 30, 0,
	31, 22, 23, 25, 20, 0, 9, 12, 14, 50,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 0, 0, 57, 0, 59,
	61, 65, 0, 76, 0, 0, 81, 0, 32, 0,
	26, 27, 18, 49, 51, 58, 60, 62, 78, 79,
	83, 24,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49,
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
		//line unql.y:43
		{ 
		logDebugGrammar("INPUT") 
	}
	case 2:
		//line unql.y:50
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 3:
		//line unql.y:56
		{ 
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND") 
	}
	case 4:
		//line unql.y:63
		{ 
		logDebugGrammar("SELECT_CORE")
	}
	case 5:
		//line unql.y:69
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 6:
		//line unql.y:75
		{ 
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 7:
		//line unql.y:81
		{ 
		logDebugGrammar("SELECT SELECT TAIL - EXPR")
	}
	case 8:
		//line unql.y:87
		{
		
	}
	case 9:
		//line unql.y:91
		{
		
	}
	case 10:
		//line unql.y:96
		{
		logDebugGrammar("RESULT STAR")
	}
	case 11:
		//line unql.y:100
		{ 
		logDebugGrammar("RESULT EXPR")
	}
	case 12:
		//line unql.y:104
		{ 
		logDebugGrammar("SORT EXPR ASC")
	}
	case 13:
		//line unql.y:110
		{
		logDebugGrammar("STAR")
	}
	case 14:
		//line unql.y:114
		{
		logDebugGrammar("PATH DOT STAR")
	}
	case 15:
		//line unql.y:119
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 16:
		//line unql.y:123
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
	}
	case 17:
		//line unql.y:128
		{
		logDebugGrammar("FROM DATASOURCE")
	}
	case 18:
		//line unql.y:132
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
	}
	case 19:
		//line unql.y:138
		{ 
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 20:
		//line unql.y:142
		{
		logDebugGrammar("SELECT WHERE - EXPR")
	}
	case 22:
		//line unql.y:149
		{
		
	}
	case 23:
		//line unql.y:155
		{
		
	}
	case 24:
		//line unql.y:159
		{
		
	}
	case 25:
		//line unql.y:164
		{ 
		logDebugGrammar("SORT EXPR")
	}
	case 26:
		//line unql.y:168
		{ 
		logDebugGrammar("SORT EXPR ASC")
	}
	case 27:
		//line unql.y:172
		{ 
		logDebugGrammar("SORT EXPR DESC")
	}
	case 28:
		//line unql.y:177
		{
		
	}
	case 29:
		//line unql.y:181
		{
		
	}
	case 30:
		//line unql.y:185
		{
		
	}
	case 31:
		//line unql.y:191
		{
		logDebugGrammar("LIMIT")
	}
	case 32:
		//line unql.y:196
		{ 
		logDebugGrammar("OFFSET")
	}
	case 33:
		//line unql.y:203
		{
		logDebugGrammar("EXPRESSION")
	}
	case 34:
		//line unql.y:208
		{
		logDebugGrammar("EXPR - PLUS")
	}
	case 35:
		//line unql.y:212
		{
		logDebugGrammar("EXPR - MINUS")
	}
	case 36:
		//line unql.y:216
		{
		logDebugGrammar("EXPR - MULT")
	}
	case 37:
		//line unql.y:220
		{
		logDebugGrammar("EXPR - DIV")
	}
	case 38:
		//line unql.y:224
		{
		logDebugGrammar("EXPR - MOD")
	}
	case 39:
		//line unql.y:228
		{
		logDebugGrammar("EXPR - DIV")
	}
	case 40:
		//line unql.y:232
		{
		logDebugGrammar("EXPR - AND")
	}
	case 41:
		//line unql.y:236
		{
		logDebugGrammar("EXPR - OR")
	}
	case 42:
		//line unql.y:240
		{
		logDebugGrammar("EXPR - EQ")
	}
	case 43:
		//line unql.y:244
		{
		logDebugGrammar("EXPR - LT")
	}
	case 44:
		//line unql.y:248
		{
		logDebugGrammar("EXPR - LTE")
	}
	case 45:
		//line unql.y:252
		{
		logDebugGrammar("EXPR - GT")
	}
	case 46:
		//line unql.y:256
		{
		logDebugGrammar("EXPR - GTE")
	}
	case 47:
		//line unql.y:260
		{
		logDebugGrammar("EXPR - NE")
	}
	case 48:
		//line unql.y:264
		{
		logDebugGrammar("EXPR - LIKE")
	}
	case 49:
		//line unql.y:268
		{
		logDebugGrammar("EXPR - NOT LIKE")
	}
	case 50:
		//line unql.y:272
		{
		logDebugGrammar("EXPR DOT MEMBER")
	}
	case 51:
		//line unql.y:276
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
	}
	case 52:
		//line unql.y:280
		{
		
	}
	case 53:
		//line unql.y:286
		{
		logDebugGrammar("EXPR - NOT")
	}
	case 54:
		//line unql.y:290
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
	}
	case 55:
		//line unql.y:294
		{
		
	}
	case 56:
		//line unql.y:299
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 57:
		//line unql.y:303
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
	}
	case 58:
		//line unql.y:307
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
	}
	case 59:
		//line unql.y:311
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
	}
	case 60:
		//line unql.y:315
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
	}
	case 61:
		//line unql.y:319
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
	}
	case 62:
		//line unql.y:323
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
	}
	case 64:
		//line unql.y:331
		{
		logDebugGrammar("LITERAL")
	}
	case 65:
		//line unql.y:335
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 66:
		//line unql.y:342
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
	}
	case 67:
		//line unql.y:346
		{
		logDebugGrammar("NUMBER")
	}
	case 68:
		//line unql.y:350
		{
		logDebugGrammar("OBJECT")
	}
	case 69:
		//line unql.y:354
		{
		logDebugGrammar("ARRAY")
	}
	case 70:
		//line unql.y:358
		{
		logDebugGrammar("TRUE")
	}
	case 71:
		//line unql.y:362
		{
		logDebugGrammar("FALSE")
	}
	case 72:
		//line unql.y:366
		{
		logDebugGrammar("NULL")
	}
	case 73:
		//line unql.y:372
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
	}
	case 74:
		//line unql.y:376
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
	}
	case 75:
		//line unql.y:382
		{
		logDebugGrammar("EMPTY OBJECT")
	}
	case 76:
		//line unql.y:386
		{
		logDebugGrammar("OBJECT")
	}
	case 77:
		//line unql.y:392
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 78:
		//line unql.y:396
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
	}
	case 79:
		//line unql.y:402
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
	}
	case 80:
		//line unql.y:408
		{
		logDebugGrammar("EMPTY ARRAY")
	}
	case 81:
		//line unql.y:412
		{
		logDebugGrammar("ARRAY")
	}
	case 82:
		//line unql.y:418
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
	}
	case 83:
		//line unql.y:422
		{ 
		logDebugGrammar("EXPRESSION LIST COMPOUND")
	}
	}
	goto yystack /* stack new state and value */
}
