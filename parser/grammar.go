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
	stmts []ast.Stmt
	stmt  ast.Stmt
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
const PLUSPLUS = 57364
const MINUSMINUS = 57365
const PLUSEQ = 57366
const MINUSEQ = 57367
const MULEQ = 57368
const DIVEQ = 57369
const MODEQ = 57370
const UNARY = 57371

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
	"PLUSPLUS",
	"MINUSMINUS",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"MODEQ",
	"'='",
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

//line ./parser/grammar.go.y:283

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 72,
	43, 5,
	-2, 47,
	-1, 81,
	43, 5,
	-2, 47,
	-1, 93,
	43, 5,
	-2, 47,
}

const yyPrivate = 57344

const yyLast = 209

var yyAct = [...]int{

	8, 82, 85, 5, 95, 84, 2, 23, 24, 22,
	38, 91, 88, 93, 72, 6, 90, 73, 53, 54,
	55, 56, 57, 29, 30, 31, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 49,
	39, 4, 51, 74, 75, 76, 77, 78, 35, 37,
	79, 27, 28, 29, 30, 31, 40, 32, 33, 35,
	37, 26, 25, 83, 3, 34, 36, 27, 28, 29,
	30, 31, 10, 1, 86, 0, 34, 36, 27, 28,
	29, 30, 31, 87, 0, 0, 0, 23, 81, 89,
	0, 86, 0, 92, 0, 94, 32, 33, 35, 37,
	26, 25, 47, 48, 42, 43, 44, 45, 46, 16,
	0, 13, 9, 0, 0, 34, 36, 27, 28, 29,
	30, 31, 21, 7, 11, 80, 0, 14, 15, 32,
	33, 35, 37, 26, 25, 41, 0, 18, 19, 0,
	0, 0, 20, 0, 17, 12, 0, 58, 34, 36,
	27, 28, 29, 30, 31, 32, 33, 35, 37, 26,
	50, 52, 0, 16, 0, 13, 0, 0, 0, 0,
	0, 0, 0, 0, 34, 36, 27, 28, 29, 30,
	31, 14, 15, 32, 33, 35, 37, 0, 0, 0,
	0, 18, 19, 0, 0, 0, 20, 0, 17, 0,
	0, 0, 34, 36, 27, 28, 29, 30, 31,
}
var yyPact = [...]int{

	-26, -1000, 105, -26, -26, -1000, -1000, -1000, 118, 159,
	22, 105, 80, 10, 36, 36, -1000, 159, 159, 159,
	159, 159, -1000, 105, -1000, 159, 159, 159, 159, 159,
	159, 159, 159, 159, 159, 159, 159, 159, 118, -28,
	-24, -1000, 159, 159, 159, 159, 159, -1000, -1000, 159,
	-1000, -1000, -1000, 85, -1000, -1000, -1000, 46, -1000, 144,
	172, -11, -11, -1000, -1000, -1000, 35, 35, 19, 19,
	19, 19, -26, 159, 118, 118, 118, 118, 118, 118,
	-1000, -26, -31, -26, 105, -25, 118, -32, -1000, -1000,
	159, -1000, -29, -26, -39, -1000,
}
var yyPgo = [...]int{

	0, 73, 123, 72, 63, 56, 1, 0, 2, 145,
	5, 3, 41,
}
var yyR1 = [...]int{

	0, 1, 1, 4, 4, 6, 6, 2, 2, 2,
	2, 5, 5, 3, 3, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	7, 7, 7, 7, 8, 8, 9, 10, 10, 12,
	12, 11,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 3, 0, 2, 1, 2, 1,
	9, 0, 1, 5, 5, 1, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 0, 1, 1, 0, 1, 1,
	2, 1,
}
var yyChk = [...]int{

	-1000, -1, -10, -4, -12, -11, 41, -2, -7, 7,
	-3, 19, -9, 6, 22, 23, 4, 39, 32, 33,
	37, 17, -10, -11, -11, 16, 15, 32, 33, 34,
	35, 36, 11, 12, 30, 13, 31, 14, -7, 18,
	-5, -2, 24, 25, 26, 27, 28, 22, 23, 29,
	-9, 6, -9, -7, -7, -7, -7, -7, -2, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, 42, 41, -7, -7, -7, -7, -7, -7,
	40, 42, -6, -4, -10, -8, -7, -6, 43, -10,
	41, 43, -8, 42, -6, 43,
}
var yyDef = [...]int{

	47, -2, 1, 47, 48, 49, 51, 3, 7, 0,
	9, 11, 15, 46, 0, 0, 26, 0, 0, 0,
	0, 0, 2, 49, 50, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 8, 0,
	0, 12, 0, 0, 0, 0, 0, 22, 23, 0,
	24, 46, 25, 0, 41, 42, 43, 0, 4, 27,
	28, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, -2, 44, 17, 18, 19, 20, 21, 16,
	29, -2, 0, 47, 0, 0, 45, 0, 14, 6,
	44, 13, 0, -2, 0, 10,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 37, 3, 3, 3, 36, 3, 3,
	39, 40, 34, 32, 3, 33, 3, 35, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 41,
	31, 29, 30, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 42, 3, 43,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 38,
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
		//line ./parser/grammar.go.y:56
		{
			yyVAL.stmts = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:60
		{
			yyVAL.stmts = yyDollar[1].stmts
			yylex.(*Lexer).result = yyVAL.stmts
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:67
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:71
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:77
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:81
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:87
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:91
		{
			yyVAL.stmt = &ast.PrintStmt{Expr: yyDollar[2].expr}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:95
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 10:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:99
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[2].stmt, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:105
		{
			yyVAL.stmt = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:109
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:115
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:125
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt.(*ast.IfStmt).Else = yyDollar[4].stmts
			}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:136
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:140
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].token.Literal, Right: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:145
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:149
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:153
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:157
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:161
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:165
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++"}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:169
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--"}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:173
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:177
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:181
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:186
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:190
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:194
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:198
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:202
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:206
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:210
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:214
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:219
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:223
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:227
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:231
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:235
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:239
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:244
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:248
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:252
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:258
		{
			yyVAL.expr = nil
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:262
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:268
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	}
	goto yystack /* stack new state and value */
}
