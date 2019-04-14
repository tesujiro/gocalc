package vm

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/tesujiro/gocalc/ast"
)

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
		tmp := env.Block().NewAlloca(types.I32)
		// LLIR: store i32 <u>, i32* %x
		i1 := constant.NewInt(types.I32, i64)
		env.Block().NewStore(i1, tmp)
		return value.Value(tmp), nil

	case *ast.IdentExpr:
		id := expr.(*ast.IdentExpr).Literal
		v, err := env.GetVar(id)
		if err != nil {
			return nil, err
		}
		return v, nil

	case *ast.AssExpr:
		assExpr := expr.(*ast.AssExpr)
		key, exp := assExpr.Left, assExpr.Right
		var val value.Value
		var err error
		if val, err = evalExpr(exp, env); err != nil {
			return nil, err
		}

		stored_value, err := env.GetVar(key)
		if err == ErrUnknownSymbol {
			// LLIR: %x = alloca i32
			tmp := env.Block().NewAlloca(val.Type())
			// LLIR: store i32 <u>, i32* %x
			env.Block().NewStore(val, tmp)
		} else if err != nil {
			return nil, err
		} else {
			// LLIR: %x = load i32, i32* %y
			v_value := env.Block().NewLoad(val)
			// LLIR: store i32 <u>, i32* %x
			env.Block().NewStore(v_value, stored_value)
		}
		if err := env.SetVar(key, val); err != nil {
			return nil, err
		}
		return val, nil

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
		l_register := env.Block().NewLoad(left)
		// LLIR: %x = load i32, i32* %y
		r_register := env.Block().NewLoad(right)

		var result value.Value
		switch expr.(*ast.BinOpExpr).Operator {
		case "+":
			// LLIR: %r= add i32 %l, %r
			result = env.Block().NewAdd(l_register, r_register)
		case "-":
			// LLIR: %r= sub i32 %l, %r
			result = env.Block().NewSub(l_register, r_register)
		case "*":
			// LLIR: %r= mul i32 %l, %r
			result = env.Block().NewMul(l_register, r_register)
		case "<":
			result = env.Block().NewICmp(enum.IPredSLT, l_register, r_register)
		case ">":
			result = env.Block().NewICmp(enum.IPredSGT, l_register, r_register)
		case "<=":
			result = env.Block().NewICmp(enum.IPredSLE, l_register, r_register)
		case ">=":
			result = env.Block().NewICmp(enum.IPredSGE, l_register, r_register)
		case "==":
			result = env.Block().NewICmp(enum.IPredEQ, l_register, r_register)
		case "!=":
			result = env.Block().NewICmp(enum.IPredNE, l_register, r_register)
		case "&&":
			result = env.Block().NewAnd(l_register, r_register)
		case "||":
			result = env.Block().NewOr(l_register, r_register)
		/*
			case "/":
					num.Quo(lnum, rnum)
		*/
		default:
			return nil, fmt.Errorf("invalid binary operation: %v %v %v", left, expr.(*ast.BinOpExpr).Operator, right)
		}

		// LLIR: %x = alloca i32
		tmp := env.Block().NewAlloca(result.Type())
		// LLIR: store i32 <u>, i32* %x
		env.Block().NewStore(result, tmp)

		return value.Value(tmp), nil

	case *ast.CompExpr:
		left := expr.(*ast.CompExpr).Left
		right := expr.(*ast.CompExpr).Right
		operator := expr.(*ast.CompExpr).Operator

		if operator == "++" || operator == "--" {
			right = &ast.NumExpr{Literal: "1"}
		}
		result, err := evalExpr(&ast.BinOpExpr{Left: left, Operator: operator[0:1], Right: right}, env)
		if err != nil {
			return nil, err
		}

		after_val, err := evalAssExpr(left.(*ast.IdentExpr).Literal, result, env)
		if err != nil {
			return nil, err
		}
		return after_val, nil

	case *ast.ParentExpr:
		sub := expr.(*ast.ParentExpr).SubExpr
		return evalExpr(sub, env)
	default:
		return nil, fmt.Errorf("invalid expression: %v", reflect.TypeOf(expr))
	}
}

func evalAssExpr(key string, val value.Value, env *Env) (value.Value, error) {
	stored_value, err := env.GetVar(key)
	if err == ErrUnknownSymbol {
		// LLIR: %x = alloca i32
		tmp := env.Block().NewAlloca(val.Type())
		// LLIR: store i32 <u>, i32* %x
		env.Block().NewStore(val, tmp)
	} else if err != nil {
		return nil, err
	} else {
		// LLIR: %x = load i32, i32* %y
		v_value := env.Block().NewLoad(val)
		// LLIR: store i32 <u>, i32* %x
		env.Block().NewStore(v_value, stored_value)
	}
	if err := env.SetVar(key, val); err != nil {
		return nil, err
	}
	return val, nil
}
