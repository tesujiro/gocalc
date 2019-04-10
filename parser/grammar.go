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
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'('",
	"')'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/grammar.go.y:164

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 124

var yyAct = [...]int{

	8, 18, 19, 20, 21, 22, 6, 11, 30, 10,
	29, 4, 3, 31, 20, 21, 22, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
	12, 46, 23, 24, 26, 28, 17, 16, 1, 0,
	0, 25, 27, 18, 19, 20, 21, 22, 0, 47,
	23, 24, 26, 28, 17, 16, 7, 0, 0, 25,
	27, 18, 19, 20, 21, 22, 23, 24, 26, 28,
	17, 32, 0, 0, 0, 25, 27, 18, 19, 20,
	21, 22, 23, 24, 26, 28, 2, 0, 0, 0,
	13, 25, 27, 18, 19, 20, 21, 22, 26, 28,
	11, 0, 10, 9, 0, 25, 27, 18, 19, 20,
	21, 22, 5, 0, 0, 0, 14, 15, 0, 0,
	0, 0, 0, 12,
}
var yyPact = [...]int{

	-23, -1000, 96, -23, -23, -1000, -1000, -1000, 39, 3,
	-11, -1000, 3, -1000, 96, -1000, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 39,
	3, 21, -1000, 55, 71, -10, -10, -1000, -1000, -1000,
	85, 85, -21, -21, -21, -21, 39, -1000,
}
var yyPgo = [...]int{

	0, 38, 56, 12, 0, 86, 112, 11,
}
var yyR1 = [...]int{

	0, 1, 1, 3, 3, 2, 2, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 5, 5, 7, 7, 6,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 3, 1, 2, 1, 3, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 0, 1, 1, 2, 1,
}
var yyChk = [...]int{

	-1000, -1, -5, -3, -7, -6, 29, -2, -4, 7,
	6, 4, 27, -5, -6, -6, 16, 15, 22, 23,
	24, 25, 26, 11, 12, 20, 13, 21, 14, -4,
	19, -4, -2, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, 28,
}
var yyDef = [...]int{

	24, -2, 1, 24, 25, 26, 28, 3, 5, 0,
	7, 9, 0, 2, 26, 27, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 6,
	0, 0, 4, 10, 11, 13, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 23, 8, 12,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 26, 3, 3,
	27, 28, 24, 22, 3, 23, 3, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 29,
	21, 19, 20,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
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
		//line ./parser/grammar.go.y:52
		{
			yyVAL.stmts = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:56
		{
			yyVAL.stmts = yyDollar[1].stmts
			yylex.(*Lexer).result = yyVAL.stmts
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:63
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:67
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:73
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:77
		{
			yyVAL.stmt = &ast.PrintStmt{Expr: yyDollar[2].expr}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:83
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:87
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].token.Literal, Right: yyDollar[3].expr}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:91
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:96
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:100
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:104
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:108
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:112
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:116
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:120
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:124
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:129
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:133
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:137
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:141
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:145
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:149
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	}
	goto yystack /* stack new state and value */
}
