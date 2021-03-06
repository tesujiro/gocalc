package ast

type Expr interface{}

type IdentExpr struct {
	Literal string
}

type AssExpr struct {
	Left  []Expr
	Right []Expr
}
type NumExpr struct {
	Literal string
}

type StringExpr struct {
	Literal string
}

type ConstExpr struct {
	Literal string
}

type ParentExpr struct {
	SubExpr Expr
}

type UnaryExpr struct {
	Operator string
	Expr     Expr
}

type BinOpExpr struct {
	Left     Expr
	Operator string
	Right    Expr
}

type CompExpr struct {
	Left     Expr
	Operator string
	Right    Expr
	//After    bool
}

type LenExpr struct {
	Expr Expr
}
