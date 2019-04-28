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
	exprs	[]ast.Expr
	stmt	ast.Stmt
	stmts	[]ast.Stmt
}

%type	<stmts>		program
%type	<stmt>		stmt
%type	<stmt>		stmt_if
%type	<stmts>		stmts
%type	<stmt>		opt_stmt
%type	<stmts>		opt_stmts
%type	<expr>		expr
%type	<expr>		opt_expr
%type	<expr>		variable
%type	<expr>		multi_val_assign
%type	<exprs>		exprs
%type	<exprs>		variables

%token	<token>		NUMBER STRING IDENT
%token	<token>		PRINT PRINTF
%token	<token>		TRUE FALSE
%token	<token>		EQEQ NEQ GE LE ANDAND OROR
%token	<token>		IF ELSE
%token	<token>		FOR BREAK CONTINUE
%token	<token>		RETURN EXIT
%token	<token>		PLUSPLUS MINUSMINUS PLUSEQ MINUSEQ MULEQ DIVEQ MODEQ
/*
%token	<token>	TRUE FALSE NIL
%token	<token>	LEN 
*/

%right '=' PLUSEQ MINUSEQ MULEQ DIVEQ MODEQ
%left OROR
%left ANDAND
/*%left IDENT*/
%nonassoc ',' vars
%left EQEQ NEQ
%left '>' '<' GE LE

%left NUMBER STRING
%left '+' '-'
%left '*' '/' '%'
%right '!' UNARY
%left PLUSPLUS MINUSMINUS
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

opt_stmts
	: /* empty */
	{
		$$ = []ast.Stmt{}
	}
	| stmts opt_term
	{
		$$ = $1
	}
	
stmt
	: expr
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| multi_val_assign
	{
		$$ = &ast.ExprStmt{Expr: $1}
	}
	| PRINT expr
	{
		$$ = &ast.PrintStmt{Expr: $2}
	}
	| stmt_if
	{
		$$ = $1
	}
	| FOR opt_stmt ';' opt_expr ';' opt_expr '{' opt_stmts '}'
	{
		$$ = &ast.CForLoopStmt{Stmt1: $2, Expr2: $4, Expr3: $6, Stmts: $8}
	}
	| BREAK
	{
		$$ = &ast.BreakStmt{}
	}
	| CONTINUE
	{
		$$ = &ast.ContinueStmt{}
	}

opt_stmt
	:
	{
		$$ = nil
	}
	| stmt
	{
		$$ = $1
	}

stmt_if
	: IF expr '{' opt_stmts '}'
	{
	    $$ = &ast.IfStmt{If: $2, Then: $4, Else: nil}
	}
	/*
	| stmt_if ELSE IF expr '{' opt_stmts '}'
	{
	        $$.(*ast.IfStmt).ElseIf = append($$.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: $4, Then: $6} )
	}
	*/
	| stmt_if ELSE '{' opt_stmts '}'
	{
		if $$.(*ast.IfStmt).Else != nil {
			yylex.Error("multiple else statement")
		} else {
			//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
			$$.(*ast.IfStmt).Else = $4
		}
	}

multi_val_assign
	: variables '=' exprs
	{
		$$ = &ast.AssExpr{Left: $1, Right: $3}
	}

exprs
	: expr
	{
		$$ = []ast.Expr{$1}
	}
	| exprs ',' opt_term expr
	{
		$$ = append($1,$4)
	}

expr
	: variable
	{
		$$ = $1
	}
	/*
	| variable '=' expr
	{
		$$ = &ast.AssExpr{Left: []ast.Expr{$1}, Right: []ast.Expr{$3}}
	}
	*/
	/* COMPOSITE EXPRESSION */
	| variable PLUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "+=", Right: $3}
	}
	| variable MINUSEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "-=", Right: $3}
	}
	| variable MULEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "*=", Right: $3}
	}
	| variable DIVEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "/=", Right: $3}
	}
	| variable MODEQ expr
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "%=", Right: $3}
	}
	| variable PLUSPLUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "++"}
	}
	| variable MINUSMINUS
	{
		$$ = &ast.CompExpr{Left: $1, Operator: "--"}
	}
	| PLUSPLUS variable
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "++"}
	}
	| MINUSMINUS variable
	{
		$$ = &ast.CompExpr{Left: $2, Operator: "--"}
	}
	| NUMBER
	{
		$$ = &ast.NumExpr{Literal: $1.Literal}
	}
	| STRING
	{
		$$ = &ast.StringExpr{Literal: $1.Literal}
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
	/* UNARY EXPRESSION */
	| '+' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "+", Expr:$2}
	}
	| '-' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "-", Expr:$2}
	}
	| '!' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Operator: "!", Expr:$2}
	}

opt_expr
	:
	{
		$$ = nil
	}
	| expr
	{
		$$ = $1
	}

variables
	: variable
	{
		$$ = []ast.Expr{$1}
	}
	| variables ',' opt_term variable
	{
		$$ = append($1,$4)
	}

variable
	: IDENT
	{
		$$ = &ast.IdentExpr{Literal: $1.Literal}
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
