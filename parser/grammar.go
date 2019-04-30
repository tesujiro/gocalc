//line ./parser/grammar.go.y:2
package parser

import __yyfmt__ "fmt"

//line ./parser/grammar.go.y:2
import (
	//"fmt"
	"github.com/tesujiro/gocalc/ast"
)

//line ./parser/grammar.go.y:9
type yySymType struct {
	yys   int
	token ast.Token
	expr  ast.Expr
	exprs []ast.Expr
	stmt  ast.Stmt
	stmts []ast.Stmt
}

const NUMBER = 57346
const STRING = 57347
const IDENT = 57348
const PRINT = 57349
const PRINTF = 57350
const TRUE = 57351
const FALSE = 57352
const EQEQ = 57353
const NEQ = 57354
const GE = 57355
const LE = 57356
const ANDAND = 57357
const OROR = 57358
const IF = 57359
const ELSE = 57360
const FOR = 57361
const BREAK = 57362
const CONTINUE = 57363
const RETURN = 57364
const EXIT = 57365
const PLUSPLUS = 57366
const MINUSMINUS = 57367
const PLUSEQ = 57368
const MINUSEQ = 57369
const MULEQ = 57370
const DIVEQ = 57371
const MODEQ = 57372
const LEN = 57373
const vars = 57374
const UNARY = 57375

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"STRING",
	"IDENT",
	"PRINT",
	"PRINTF",
	"TRUE",
	"FALSE",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"ANDAND",
	"OROR",
	"IF",
	"ELSE",
	"FOR",
	"BREAK",
	"CONTINUE",
	"RETURN",
	"EXIT",
	"PLUSPLUS",
	"MINUSMINUS",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"MODEQ",
	"LEN",
	"'='",
	"','",
	"vars",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"UNARY",
	"'('",
	"')'",
	"';'",
	"'{'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/grammar.go.y:337

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 15,
	32, 53,
	33, 53,
	-2, 21,
	-1, 80,
	48, 5,
	-2, 56,
	-1, 92,
	48, 5,
	-2, 56,
	-1, 109,
	48, 5,
	-2, 56,
}

const yyPrivate = 57344

const yyLast = 277

var yyAct = [...]int{

	8, 93, 96, 95, 2, 5, 111, 28, 106, 29,
	30, 44, 102, 109, 80, 6, 104, 81, 56, 35,
	36, 37, 59, 60, 61, 62, 99, 65, 63, 64,
	46, 27, 67, 68, 69, 70, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 45, 94, 3, 15, 4,
	82, 83, 84, 85, 86, 25, 7, 87, 15, 33,
	34, 35, 36, 37, 90, 57, 58, 89, 91, 48,
	9, 47, 11, 1, 0, 15, 0, 0, 0, 0,
	0, 0, 97, 0, 0, 0, 66, 0, 38, 39,
	41, 43, 32, 31, 101, 0, 0, 0, 103, 0,
	29, 0, 0, 105, 0, 97, 108, 107, 0, 0,
	0, 110, 40, 42, 33, 34, 35, 36, 37, 16,
	17, 27, 10, 0, 92, 54, 55, 49, 50, 51,
	52, 53, 26, 0, 12, 13, 14, 100, 0, 19,
	20, 15, 0, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 22, 23, 0, 41, 43, 24, 0, 21,
	38, 39, 41, 43, 32, 31, 0, 0, 0, 0,
	0, 38, 39, 41, 43, 32, 31, 40, 42, 33,
	34, 35, 36, 37, 40, 42, 33, 34, 35, 36,
	37, 0, 0, 0, 98, 40, 42, 33, 34, 35,
	36, 37, 0, 0, 0, 88, 38, 39, 41, 43,
	32, 31, 0, 0, 0, 0, 0, 16, 17, 27,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	40, 42, 33, 34, 35, 36, 37, 19, 20, 38,
	39, 41, 43, 32, 18, 0, 38, 39, 41, 43,
	22, 23, 0, 0, 0, 24, 0, 21, 0, 0,
	0, 0, 0, 40, 42, 33, 34, 35, 36, 37,
	40, 42, 33, 34, 35, 36, 37,
}
var yyPact = [...]int{

	-31, -1000, 115, -31, -31, -1000, -1000, -1000, 195, -1000,
	213, 12, 115, -1000, -1000, 101, -1000, -1000, -26, 25,
	25, 213, 213, 213, 213, -4, 213, -1000, -1000, 115,
	-1000, 213, 213, 213, 213, 213, 213, 213, 213, 213,
	213, 213, 213, 213, 195, 101, -33, -29, -1000, 213,
	213, 213, 213, 213, -1000, -1000, 213, -1000, -1000, 160,
	-1000, -1000, -1000, 213, -31, 77, -1000, 228, 235, -20,
	-20, -1000, -1000, -1000, 142, 142, 22, 22, 22, 22,
	-31, 213, 195, 195, 195, 195, 195, 149, -1000, -7,
	195, 25, -31, -36, -31, 115, -30, 195, -1000, -31,
	-1000, -40, -1000, -1000, 213, 213, -1000, -34, 195, -31,
	-42, -1000,
}
var yyPgo = [...]int{

	0, 73, 56, 72, 46, 71, 1, 0, 2, 45,
	70, 67, 55, 3, 5, 49,
}
var yyR1 = [...]int{

	0, 1, 1, 4, 4, 6, 6, 2, 2, 2,
	2, 2, 2, 2, 5, 5, 3, 3, 10, 11,
	11, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 8, 8, 12, 12, 9, 13, 13, 15, 15,
	14,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 3, 0, 2, 1, 1, 2,
	1, 9, 1, 1, 0, 1, 5, 5, 3, 1,
	4, 1, 1, 1, 4, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 0, 1, 1, 4, 1, 0, 1, 1, 2,
	1,
}
var yyChk = [...]int{

	-1000, -1, -13, -4, -15, -14, 46, -2, -7, -10,
	7, -3, 19, 20, 21, -9, 4, 5, 31, 24,
	25, 44, 37, 38, 42, -12, 17, 6, -13, -14,
	-14, 16, 15, 37, 38, 39, 40, 41, 11, 12,
	35, 13, 36, 14, -7, -9, 18, -5, -2, 26,
	27, 28, 29, 30, 24, 25, 44, -9, -9, -7,
	-7, -7, -7, 32, 33, -7, -2, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	47, 46, -7, -7, -7, -7, -7, -7, 45, -11,
	-7, -13, 47, -6, -4, -13, -8, -7, 45, 33,
	-9, -6, 48, -13, 46, -13, 48, -8, -7, 47,
	-6, 48,
}
var yyDef = [...]int{

	56, -2, 1, 56, 57, 58, 60, 3, 7, 8,
	0, 10, 14, 12, 13, -2, 22, 23, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 55, 2, 58,
	59, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 21, 0, 0, 15, 0,
	0, 0, 0, 0, 30, 31, 0, 32, 33, 0,
	48, 49, 50, 0, 56, 0, 4, 34, 35, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	-2, 51, 25, 26, 27, 28, 29, 0, 36, 18,
	19, 0, -2, 0, 56, 0, 0, 52, 24, 56,
	54, 0, 17, 6, 51, 0, 16, 0, 20, -2,
	0, 11,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 42, 3, 3, 3, 41, 3, 3,
	44, 45, 39, 37, 33, 38, 3, 40, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 46,
	36, 32, 35, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 47, 3, 48,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	34, 43,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:62
		{
			yyVAL.stmts = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:66
		{
			yyVAL.stmts = yyDollar[1].stmts
			yylex.(*Lexer).result = yyVAL.stmts
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:73
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:77
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:83
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:87
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:93
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:97
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:101
		{
			yyVAL.stmt = &ast.PrintStmt{Expr: yyDollar[2].expr}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:105
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 11:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:109
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[2].stmt, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:113
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:117
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:123
		{
			yyVAL.stmt = nil
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:127
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 16:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:133
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:143
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt.(*ast.IfStmt).Else = yyDollar[4].stmts
			}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:154
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:160
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:164
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:170
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:180
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:184
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:188
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:193
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:197
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:201
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:205
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:209
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:213
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++"}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:217
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--"}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:221
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:225
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:230
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:234
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:238
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:242
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:246
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:250
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:254
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:258
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:263
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:267
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:271
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:275
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:279
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:283
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:288
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:292
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:296
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:302
		{
			yyVAL.expr = nil
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:306
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:312
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:316
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:322
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	}
	goto yystack /* stack new state and value */
}
