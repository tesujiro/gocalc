package vm

import (
	"errors"
	"fmt"

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

	// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), i32 %7)
	zero := constant.NewInt(types.I32, 0)
	env.entry.NewCall(env.lib["printf"], constant.NewGetElementPtr(env.defs[".result"], zero, zero), r)

	if r.Type() == types.I32 {
		// LLIR: ret i32 %y
		env.entry.NewRet(r)
	} else {
		r32 := env.entry.NewZExt(r, types.I32)
		// LLIR: ret i32 %y
		env.entry.NewRet(r32)
	}
	return nil
}

func runSingleStmt(stmt ast.Stmt, env *Env) (value.Value, error) {
	switch stmt.(type) {
	case *ast.ExprStmt:
		return evalExpr(stmt.(*ast.ExprStmt).Expr, env)
	case *ast.PrintStmt:
		v, err := evalExpr(stmt.(*ast.PrintStmt).Expr, env)
		if err != nil {
			return nil, err
		}
		// LLIR: %y = load i32, i32* %x
		r := env.entry.NewLoad(v)

		switch r.Type() {
		case types.I1, types.I32:
			// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), i32 %7)
			zero := constant.NewInt(types.I32, 0)
			env.entry.NewCall(env.lib["printf"], constant.NewGetElementPtr(env.defs[".print_int"], zero, zero), r)
			return v, nil
		default:
			return nil, fmt.Errorf("print invalid value type : %v", v.Type())
		}
	default:
		return nil, fmt.Errorf("invalid statement")
	}
}
