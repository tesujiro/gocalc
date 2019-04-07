package ast

type Stmt interface{}

type ExprStmt struct {
	Expr Expr
}
