package vm

import (
	"errors"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/tesujiro/gocalc/ast"
)

var (
	ErrBreak    = errors.New("unexpected break")
	ErrContinue = errors.New("unexpected continue")
	ErrNext     = errors.New("unexpected next")
	ErrReturn   = errors.New("unexpected return")
	ErrExit     = errors.New("unexpected exit")
)

func Run(stmts []ast.Stmt, env *Env) error {
	return run(stmts, env)
}

func run(stmts []ast.Stmt, env *Env) error {
	var result value.Value
	var err error
	for _, stmt := range stmts {
		result, err = runSingleStmt(stmt, env)
		if err != nil {
			return err
		}
	}

	if result == nil {
		env.entry.NewRet(constant.NewInt(types.I32, 0))
		return nil
	}

	// LLIR: %y = load i32, i32* %x
	r := env.entry.NewLoad(result)

	// LLIR: declare i32 @printf(i8* %format, ...)
	i8ptr := types.NewPointer(types.I8)
	zero := constant.NewInt(types.I32, 0)
	printf := env.module.NewFunc("printf", types.I32, ir.NewParam("format", i8ptr))
	printf.Sig.Variadic = true
	// LLIR: @.str.result = global [12 x i8] c"Result : %d\0A"
	str := env.module.NewGlobalDef(".str.result", constant.NewCharArrayFromString("Result : %d\n"))
	// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), i32 %7)
	env.entry.NewCall(printf, constant.NewGetElementPtr(str, zero, zero), r)

	// LLIR: ret i32 %y
	env.entry.NewRet(r)
	return nil
}

func runSingleStmt(stmt ast.Stmt, env *Env) (value.Value, error) {
	switch stmt.(type) {
	case *ast.ExprStmt:
		return evalExpr(stmt.(*ast.ExprStmt).Expr, env)
	default:
		return nil, fmt.Errorf("invalid statement")
	}
}
