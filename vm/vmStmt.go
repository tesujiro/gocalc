package vm

import (
	"errors"
	"fmt"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/tesujiro/gocalc/ast"
	"github.com/tesujiro/gocalc/debug"
)

var (
	ErrBreak    = errors.New("unexpected break")
	ErrContinue = errors.New("unexpected continue")
	ErrNext     = errors.New("unexpected next")
	ErrReturn   = errors.New("unexpected return")
	ErrExit     = errors.New("unexpected exit")
)

func Run(stmts []ast.Stmt, env *Env) error {
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

		//debug.Printf("run err:%v\n", err)
		if err != nil && err != ErrBreak && err != ErrContinue {
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
		child := env.NewEnv("If")
		debug.Println("NewEnv in IfStmt")
		cond, err := evalExpr(ifStmt.If, child)
		if err != nil {
			return nil, err
		}
		result = cond

		thenBlock := child.GetNewBlock("then")
		elseBlock := child.GetNewBlock("else")
		nextBlock := child.GetNewBlock("next")
		cond_r := child.Block().NewLoad(cond)
		child.Block().NewCondBr(cond_r, thenBlock, elseBlock)

		// then
		child.SetCurrentBlock(thenBlock)
		r, err := run(ifStmt.Then, child)
		if err == nil {
			child.Block().NewBr(nextBlock)
		} else if err != nil && err != ErrBreak && err != ErrContinue {
			return nil, err
		}
		result = r

		// else
		child.SetCurrentBlock(elseBlock)
		if len(stmt.(*ast.IfStmt).Else) > 0 {
			r, err := run(ifStmt.Else, child)
			if err == nil {
				child.Block().NewBr(nextBlock)
			} else if err != nil && err != ErrBreak && err != ErrContinue {
				return nil, err
			}
			result = r
		} else {
			child.Block().NewBr(nextBlock)
		}

		// next
		child.SetCurrentBlock(nextBlock)
		return result, nil

	case *ast.CForLoopStmt:
		var result value.Value
		stmt1 := stmt.(*ast.CForLoopStmt).Stmt1
		expr2 := stmt.(*ast.CForLoopStmt).Expr2
		expr3 := stmt.(*ast.CForLoopStmt).Expr3
		stmts := stmt.(*ast.CForLoopStmt).Stmts
		child := env.NewEnv("CForLoop")
		debug.Println("NewEnv in CForLoopStmt")
		condBlock := child.GetNewBlock("cond")
		loopBlock := child.GetNewBlock("loop")
		postBlock := child.GetNewBlock("post")
		nextBlock := child.GetNewBlock("next")
		child.SetContinueBlock(postBlock)
		child.SetBreakBlock(nextBlock)

		// init
		if stmt1 != nil {
			_, err := run([]ast.Stmt{stmt1}, child)
			if err != nil {
				return nil, fmt.Errorf("for init stmt error: %v", err)
			}
		}
		debug.Printf("%v.NewBr(%v)\n", child.path, condBlock)
		child.Block().NewBr(condBlock)

		// loop Condition
		child.SetCurrentBlock(condBlock)
		if expr2 == nil {
			debug.Printf("%v.NewBr(%v)\n", child.path, loopBlock)
			child.Block().NewBr(loopBlock)
		} else {
			cond, err := evalExpr(expr2, child)
			if err != nil {
				return nil, fmt.Errorf("for condition expr error: %v", err)
			}
			cond_r := env.Block().NewLoad(cond)
			debug.Printf("%v.NewCondBr(%v,%v)\n", child.path, loopBlock, nextBlock)
			child.Block().NewCondBr(cond_r, loopBlock, nextBlock)
		}

		// loop
		child.SetCurrentBlock(loopBlock)
		ret, err := run(stmts, child)
		if err == nil {
			debug.Printf("%v.NewBr(%v)\n", child.path, postBlock)
			child.Block().NewBr(postBlock)
		} else if err != nil && err != ErrBreak && err != ErrContinue {
			return nil, fmt.Errorf("for loop stmts error: %v", err)
		}
		result = ret

		// post statement
		child.SetCurrentBlock(postBlock)
		if expr3 != nil {
			_, err := evalExpr(expr3, child)
			if err != nil {
				return nil, fmt.Errorf("for final expr error: %v", err)
			}
		}
		debug.Printf("%v.NewBr(%v)\n", child.path, condBlock)
		child.Block().NewBr(condBlock)

		// next
		child.SetCurrentBlock(nextBlock)
		child.SetContinueBlock(nil)
		child.SetBreakBlock(nil)
		return result, nil

	case *ast.ContinueStmt:
		block := env.GetContinueBlock()
		if block == nil {
			return nil, fmt.Errorf("continue not inside loop")
		}
		env.Block().NewBr(block)
		return nil, ErrContinue

	case *ast.BreakStmt:
		block := env.GetBreakBlock()
		if block == nil {
			return nil, fmt.Errorf("break not inside loop")
		}
		env.Block().NewBr(block)
		return nil, ErrBreak

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
		case types.Double:
			// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), double %7)
			zero := constant.NewInt(types.I32, 0)
			env.Block().NewCall(env.lib["printf"], constant.NewGetElementPtr(env.defs[".print_float"], zero, zero), r)
			return v, nil
		default:
			return nil, fmt.Errorf("print invalid value type : %v", v.Type())
		}
	default:
		return nil, fmt.Errorf("invalid statement")
	}
}
