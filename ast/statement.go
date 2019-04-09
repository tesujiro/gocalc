package ast

type Stmt interface{}

type ExprStmt struct {
	Expr Expr
}

type PrintStmt struct {
	Expr Expr
}
