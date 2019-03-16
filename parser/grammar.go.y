%{
	package parser
	import (
		//"fmt"
		"github.com/tesujiro/gocalc/ast"
	)
%}

%union{
	token	ast.Token
	expr	ast.Expr
}

%type	<expr>		program
%type	<expr>		expr

%token	<token>		NUMBER STRING
/*
%token	<token>	IDENT NUMBER STRING TRUE FALSE NIL
%token	<token>	EQEQ NEQ GE LE NOTTILDE ANDAND OROR LEN 
%token	<token>	PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ MODEQ
*/

%left NUMBER STRING
%left '+' '-'
%left '*' '/' '%'
/*
%right '!' UNARY
%left PLUSPLUS MINUSMINUS
*/
%left '(' ')'

%%

program
	: opt_term
	{
		$$ = nil
	}
	| expr opt_term
	{
		$$ = $1
		yylex.(*Lexer).result = $$
	}

expr
	: NUMBER
	{
		$$ = &ast.NumExpr{Literal: $1.Literal}
	}
	| '(' expr ')'
	{
		$$ = &ast.ParentExpr{SubExpr: $2}
	}
	| expr '+' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "+", Right: $3}
	}
	| expr '-' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "-", Right: $3}
	}
	| expr '*' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "*", Right: $3}
	}
	| expr '/' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "/", Right: $3}
	}
	| expr '%' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "%", Right: $3}
	}

opt_term
	: /* empty */
	| term

term
	: semi
	| term semi

semi
	: ';'

%%
