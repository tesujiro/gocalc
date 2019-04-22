package ast

type Stmt interface{}

type ReturnStmt struct {
	Exprs []Expr
}

type ExitStmt struct {
	Expr Expr
}

type ExprStmt struct {
	Expr Expr
}

type PrintStmt struct {
	Expr Expr
}

type IfStmt struct {
	If     Expr
	Then   []Stmt
	Else   []Stmt
	ElseIf []Stmt
}

type CForLoopStmt struct {
	Stmt1 Stmt
	Expr2 Expr
	Expr3 Expr
	Stmts []Stmt
}

type BreakStmt struct {
}

type ContinueStmt struct {
}
