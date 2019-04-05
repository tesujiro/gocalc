package vm

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/tesujiro/gocalc/ast"
)

func Run(expr ast.Expr, env *Env) error {
	result, err := evalExpr(expr, env)
	if err != nil {
		return err
	}
	r := env.entry.NewLoad(result)
	env.entry.NewRet(r)
	return nil
}

func evalExpr(expr ast.Expr, env *Env) (value.Value, error) {
	//fmt.Printf("evalExpr(%#v)\n", expr)
	switch expr.(type) {
	case *ast.NumExpr:
		lit := expr.(*ast.NumExpr).Literal
		i64, err := strconv.ParseInt(lit, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("convert number err(%v):%v", lit, err)
		}
		// LLIR: %x = alloca i32
		tmp := env.entry.NewAlloca(types.I32)
		// LLIR: store i32 <u>, i32* %x
		i1 := constant.NewInt(types.I32, i64)
		env.entry.NewStore(i1, tmp)
		return value.Value(tmp), nil
	case *ast.BinOpExpr:
		var left, right value.Value
		var err error
		if left, err = evalExpr(expr.(*ast.BinOpExpr).Left, env); err != nil {
			return nil, err
		}
		if right, err = evalExpr(expr.(*ast.BinOpExpr).Right, env); err != nil {
			return nil, err
		}
		// LLIR: %x = load i32, i32* %y
		l_register := env.entry.NewLoad(left)
		// LLIR: %x = load i32, i32* %y
		r_register := env.entry.NewLoad(right)

		var result value.Value
		switch expr.(*ast.BinOpExpr).Operator {
		case "+":
			// LLIR: %r= add i32 %l, %r
			result = env.entry.NewAdd(l_register, r_register)
		case "-":
			// LLIR: %r= sub i32 %l, %r
			result = env.entry.NewSub(l_register, r_register)
		case "*":
			// LLIR: %r= mul i32 %l, %r
			result = env.entry.NewMul(l_register, r_register)
		/*
			case "/":
					num.Quo(lnum, rnum)
		*/
		default:
			return nil, fmt.Errorf("invalid binary operation: %v %v %v", left, expr.(*ast.BinOpExpr).Operator, right)
		}

		// LLIR: %x = alloca i32
		tmp := env.entry.NewAlloca(types.I32)
		// LLIR: store i32 <u>, i32* %x
		env.entry.NewStore(result, tmp)

		return value.Value(tmp), nil

	case *ast.ParentExpr:
		sub := expr.(*ast.ParentExpr).SubExpr
		return evalExpr(sub, env)
	default:
		return nil, fmt.Errorf("invalid expression: %v", reflect.TypeOf(expr))
	}
}
