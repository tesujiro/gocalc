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
	//return run(stmts, env)
	result, err := run(stmts, env)
	if err != nil {
		return err
	}

	if result == nil {
		env.Block().NewRet(constant.NewInt(types.I32, 0))
		return nil
	}
	env.Block().NewRet(constant.NewInt(types.I32, 0))
	return nil

	/*
		// LLIR: %y = load i32, i32* %x
		r := env.Block().NewLoad(result)

		// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), i32 %7)
		zero := constant.NewInt(types.I32, 0)
		env.Block().NewCall(env.lib["printf"], constant.NewGetElementPtr(env.defs[".result"], zero, zero), r)

		if r.Type() == types.I32 {
			// LLIR: ret i32 %y
			env.Block().NewRet(r)
		} else {
			r32 := env.Block().NewZExt(r, types.I32)
			// LLIR: ret i32 %y
			env.Block().NewRet(r32)
		}
		return nil
	*/
}

func run(stmts []ast.Stmt, env *Env) (result value.Value, err error) {
	for _, stmt := range stmts {
		result, err = runSingleStmt(stmt, env)
		if err != nil {
			return
		}
	}
	return
}

func runSingleStmt(stmt ast.Stmt, env *Env) (value.Value, error) {
	switch stmt.(type) {
	case *ast.ExprStmt:
		return evalExpr(stmt.(*ast.ExprStmt).Expr, env)
	case *ast.IfStmt:
		var result value.Value
		ifStmt := stmt.(*ast.IfStmt)
		child := env.NewEnv()
		cond, err := evalExpr(ifStmt.If, child)
		if err != nil {
			return nil, err
		}
		result = cond

		thenBlock := child.GetNewBlock("then")
		elseBlock := child.GetNewBlock("else")
		nextBlock := child.GetNewBlock("")
		cond_r := env.Block().NewLoad(cond)
		child.Block().NewCondBr(cond_r, thenBlock, elseBlock)

		// then
		child.SetCurrentBlock(thenBlock)
		r, err := run(ifStmt.Then, child)
		if err != nil {
			return nil, err
		}
		result = r
		child.Block().NewBr(nextBlock)

		// else
		child.SetCurrentBlock(elseBlock)
		if len(stmt.(*ast.IfStmt).Else) > 0 {
			r, err := run(ifStmt.Else, child)
			if err != nil {
				return nil, err
			}
			result = r
		}
		child.Block().NewBr(nextBlock)

		//thenBlock.NewBr(nextBlock)
		child.SetCurrentBlock(nextBlock)
		return result, nil

	case *ast.PrintStmt:
		v, err := evalExpr(stmt.(*ast.PrintStmt).Expr, env)
		if err != nil {
			return nil, err
		}
		// LLIR: %y = load i32, i32* %x
		r := env.Block().NewLoad(v)

		switch r.Type() {
		case types.I1, types.I32:
			// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), i32 %7)
			zero := constant.NewInt(types.I32, 0)
			env.Block().NewCall(env.lib["printf"], constant.NewGetElementPtr(env.defs[".print_int"], zero, zero), r)
			return v, nil
		default:
			return nil, fmt.Errorf("print invalid value type : %v", v.Type())
		}
	default:
		return nil, fmt.Errorf("invalid statement")
	}
}
