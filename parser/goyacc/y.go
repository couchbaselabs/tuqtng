
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

const LBRACE = 57346
const RBRACE = 57347
const LBRACKET = 57348
const RBRACKET = 57349
const COMMA = 57350
const COLON = 57351
const DQUOTE = 57352
const TRUE = 57353
const FALSE = 57354
const NULL = 57355
const INT = 57356
const NUMBER = 57357
const IDENTIFIER = 57358
const STRING = 57359
const PLUS = 57360
const MINUS = 57361
const MULT = 57362
const DIV = 57363
const CONCAT = 57364
const AND = 57365
const OR = 57366
const NOT = 57367
const EQ = 57368
const NE = 57369
const GT = 57370
const GTE = 57371
const LT = 57372
const LTE = 57373
const LPAREN = 57374
const RPAREN = 57375

var yyToknames = []string{
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

const yyNprod = 41
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 116

var yyAct = []int{

	42, 43, 2, 3, 38, 22, 23, 24, 25, 26,
	27, 28, 36, 29, 34, 32, 33, 30, 31, 57,
	22, 23, 24, 25, 26, 40, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 20,
	60, 21, 41, 37, 62, 59, 15, 16, 17, 18,
	19, 8, 11, 61, 58, 40, 39, 14, 13, 4,
	5, 12, 64, 65, 63, 35, 20, 10, 21, 9,
	7, 6, 1, 15, 16, 17, 18, 19, 8, 11,
	0, 0, 0, 0, 0, 0, 0, 5, 22, 23,
	24, 25, 26, 27, 10, 0, 29, 34, 32, 33,
	30, 31, 22, 23, 24, 25, 26, 0, 0, 0,
	29, 34, 32, 33, 30, 31,
}
var yyPact = []int{

	62, -1000, -1000, -13, -1000, 62, -1000, -1000, -1000, -1000,
	62, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	38, 35, 62, 62, 62, 62, 62, 62, 62, 62,
	62, 62, 62, 62, 62, -1000, -14, -1000, 49, 37,
	31, -1000, 46, 36, -1000, -1000, -1000, -1000, -1000, 84,
	70, 2, 2, 2, 2, 2, 2, -1000, -1000, 8,
	62, -1000, 62, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 72, 1, 3, 59, 71, 70, 69, 61, 58,
	57, 4, 56, 0,
}
var yyR1 = []int{

	0, 1, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 4, 5,
	6, 6, 6, 7, 7, 7, 7, 7, 7, 7,
	8, 8, 9, 9, 11, 11, 12, 10, 10, 13,
	13,
}
var yyR2 = []int{

	0, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 2, 1, 1,
	1, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 3, 1, 3, 3, 2, 3, 1,
	3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, 25, -5, -6, 16, -7,
	32, 17, -8, -9, -10, 11, 12, 13, 14, 15,
	4, 6, 18, 19, 20, 21, 22, 23, 24, 26,
	30, 31, 28, 29, 27, -4, -2, 5, -11, -12,
	17, 7, -13, -2, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, 33, 5, 8,
	9, 7, 8, -11, -2, -13,
}
var yyDef = []int{

	0, -2, 1, 2, 16, 0, 18, 19, 20, 21,
	0, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 17, 0, 32, 0, 34,
	0, 37, 0, 39, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 22, 33, 0,
	0, 38, 0, 35, 36, 40,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33,
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
		//line unql.y:37
		{ 
		logDebugGrammar("INPUT") 
	}
	case 2:
		//line unql.y:45
		{
		logDebugGrammar("EXPRESSION")
	}
	case 3:
		//line unql.y:50
		{
		logDebugGrammar("EXPR - PLUS")
	}
	case 4:
		//line unql.y:54
		{
		logDebugGrammar("EXPR - MINUS")
	}
	case 5:
		//line unql.y:58
		{
		logDebugGrammar("EXPR - MULT")
	}
	case 6:
		//line unql.y:62
		{
		logDebugGrammar("EXPR - DIV")
	}
	case 7:
		//line unql.y:66
		{
		logDebugGrammar("EXPR - DIV")
	}
	case 8:
		//line unql.y:70
		{
		logDebugGrammar("EXPR - AND")
	}
	case 9:
		//line unql.y:74
		{
		logDebugGrammar("EXPR - OR")
	}
	case 10:
		//line unql.y:78
		{
		logDebugGrammar("EXPR - EQ")
	}
	case 11:
		//line unql.y:82
		{
		logDebugGrammar("EXPR - LT")
	}
	case 12:
		//line unql.y:86
		{
		logDebugGrammar("EXPR - LTE")
	}
	case 13:
		//line unql.y:90
		{
		logDebugGrammar("EXPR - GT")
	}
	case 14:
		//line unql.y:94
		{
		logDebugGrammar("EXPR - GTE")
	}
	case 15:
		//line unql.y:98
		{
		logDebugGrammar("EXPR - NE")
	}
	case 16:
		//line unql.y:102
		{
		
	}
	case 17:
		//line unql.y:108
		{
		logDebugGrammar("EXPR - NOT")
	}
	case 18:
		//line unql.y:112
		{
		
	}
	case 19:
		//line unql.y:117
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 21:
		//line unql.y:124
		{
		logDebugGrammar("LITERAL")
	}
	case 22:
		//line unql.y:128
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 23:
		//line unql.y:135
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
	}
	case 24:
		//line unql.y:139
		{
		logDebugGrammar("NUMBER")
	}
	case 25:
		//line unql.y:143
		{
		logDebugGrammar("OBJECT")
	}
	case 26:
		//line unql.y:147
		{
		logDebugGrammar("ARRAY")
	}
	case 27:
		//line unql.y:151
		{
		logDebugGrammar("TRUE")
	}
	case 28:
		//line unql.y:155
		{
		logDebugGrammar("FALSE")
	}
	case 29:
		//line unql.y:159
		{
		logDebugGrammar("NULL")
	}
	case 30:
		//line unql.y:165
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
	}
	case 31:
		//line unql.y:169
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
	}
	case 32:
		//line unql.y:175
		{
		logDebugGrammar("EMPTY OBJECT")
	}
	case 33:
		//line unql.y:179
		{
		logDebugGrammar("OBJECT")
	}
	case 34:
		//line unql.y:185
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 35:
		//line unql.y:189
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
	}
	case 36:
		//line unql.y:195
		{  
		logDebugGrammar("NAMED EXPR SINGLE")
	}
	case 37:
		//line unql.y:201
		{
		logDebugGrammar("EMPTY ARRAY")
	}
	case 38:
		//line unql.y:205
		{
		logDebugGrammar("ARRAY")
	}
	case 39:
		//line unql.y:211
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
	}
	case 40:
		//line unql.y:215
		{ 
		logDebugGrammar("EXPRESSION LIST COMPOUND")
	}
	}
	goto yystack /* stack new state and value */
}
