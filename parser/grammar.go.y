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
	stmts	[]ast.Stmt
	stmt	ast.Stmt
}

%type	<stmts>		program
%type	<stmt>		stmt
%type	<stmts>		stmts
%type	<expr>		expr

%token	<token>		NUMBER STRING IDENT
%token	<token>		PRINT PRINTF
%token	<token>		TRUE FALSE
%token	<token>		EQEQ NEQ GE LE ANDAND OROR
%token	<token>		IF ELSE
/*
%token	<token>	IDENT NUMBER STRING TRUE FALSE NIL
%token	<token>	EQEQ NEQ GE LE NOTTILDE ANDAND OROR LEN 
%token	<token>	PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ MODEQ
*/

%right '='
%left OROR
%left ANDAND
%left IDENT
%left EQEQ NEQ
%left '>' '<' GE LE

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
	| stmts opt_term
	{
		$$ = $1
		yylex.(*Lexer).result = $$
	}

stmts
	: opt_term stmt 
	{
		$$ = []ast.Stmt{$2}
	}
	| stmts semi stmt
	{
		$$ = append($1,$3)
	}

stmt
	: expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| PRINT expr
	{
		$$ = &ast.PrintStmt{Expr: $2}
	}

expr
	: IDENT
	{
		$$ = &ast.IdentExpr{Literal: $1.Literal}
	}
	| IDENT '=' expr
	{
		$$ = &ast.AssExpr{Left: $1.Literal, Right: $3}
	}
	| NUMBER
	{
		$$ = &ast.NumExpr{Literal: $1.Literal}
	}
	/* BOOL EXPRESSION */
	| expr OROR expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "||", Right: $3}
	}
	| expr ANDAND expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "&&", Right: $3}
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
	/* RELATION EXPRESSION */
	| expr EQEQ expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "==", Right: $3}
	}
	| expr NEQ expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "!=", Right: $3}
	}
	| expr '>' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">", Right: $3}
	}
	| expr GE expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: ">=", Right: $3}
	}
	| expr '<' expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<", Right: $3}
	}
	| expr LE expr
	{
		$$ = &ast.BinOpExpr{Left: $1, Operator: "<=", Right: $3}
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
