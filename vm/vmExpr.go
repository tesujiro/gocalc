package vm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

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
		var num_type types.Type
		// float
		if strings.Contains(lit, ".") || strings.Contains(lit, "e") {
			f, err := strconv.ParseFloat(lit, 64)
			if err != nil {
				return nil, err
			}
			num_type = types.Double // double float NOT types.Float
			// LLIR: %x = alloca f32
			tmp := env.Block().NewAlloca(num_type)
			// LLIR: store f32 <u>, f32* %x
			i1 := constant.NewFloat(num_type.(*types.FloatType), f)
			env.Block().NewStore(i1, tmp)
			return value.Value(tmp), nil
		} else {
			// integer
			i64, err := strconv.ParseInt(lit, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("convert number err(%v):%v", lit, err)
			}
			num_type = types.I32
			// LLIR: %x = alloca i32
			tmp := env.Block().NewAlloca(num_type)
			// LLIR: store i32 <u>, i32* %x
			i1 := constant.NewInt(num_type.(*types.IntType), i64)
			env.Block().NewStore(i1, tmp)
			return value.Value(tmp), nil
		}

	case *ast.IdentExpr:
		id := expr.(*ast.IdentExpr).Literal
		v, err := env.GetVar(id)
		if err != nil {
			return nil, err
		}
		return v, nil

	case *ast.UnaryExpr:
		val, err := evalExpr(expr.(*ast.UnaryExpr).Expr, env)
		if err != nil {
			return nil, err
		}
		// LLIR: %y = load i32, i32* %x
		r := env.Block().NewLoad(val)

		/* switch r.Type() {
		case types.I1, types.I32:
		default:
		} */

		var result value.Value
		switch expr.(*ast.UnaryExpr).Operator {
		case "+":
			return val, nil
		case "-":
			// LLIR: %r= sub i32 0, %r
			result = env.Block().NewSub(constant.NewInt(types.I32, 0), r)
		case "!":
			// LLIR: %r= fneg %r
			result = env.Block().NewICmp(enum.IPredEQ, constant.NewInt(types.I1, 0), r)
		default:
			return nil, fmt.Errorf("invalid unary type")
		}

		// LLIR: %x = alloca i32
		tmp := env.Block().NewAlloca(result.Type())
		// LLIR: store i32 <u>, i32* %x
		env.Block().NewStore(result, tmp)

		return value.Value(tmp), nil

	case *ast.AssExpr:
		assExpr := expr.(*ast.AssExpr)
		key, exp := assExpr.Left, assExpr.Right
		var val value.Value
		var err error
		if val, err = evalExpr(exp, env); err != nil {
			return nil, err
		}

		_, err = evalAssExpr(key, val, env)
		if err != nil {
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

		/*
			if l_register.Type() == types.Double || r_register.Type() == types.Double {
				arithmetic_type = types.Double
			}
		*/
		l_type, r_type := l_register.Type(), r_register.Type()
		arithmetic_type := precedenceOfTypes(l_type, r_type)
		compare := func(ipred enum.IPred, fpred enum.FPred) value.Value {
			if arithmetic_type != types.Double {
				return env.Block().NewICmp(ipred, l_register, r_register)
			} else {
				l := toDouble(env, l_register)
				r := toDouble(env, r_register)
				return env.Block().NewFCmp(fpred, l, r)
			}
		}

		var result value.Value
		switch expr.(*ast.BinOpExpr).Operator {
		case "+":
			if arithmetic_type != types.Double {
				// LLIR: %r= add i32 %l, %r
				result = env.Block().NewAdd(l_register, r_register)
			} else {
				l := toDouble(env, l_register)
				r := toDouble(env, r_register)
				// LLIR: %r= add i32 %l, %r
				result = env.Block().NewFAdd(l, r)
			}
		case "-":
			if arithmetic_type != types.Double {
				// LLIR: %r= sub i32 %l, %r
				result = env.Block().NewSub(l_register, r_register)
			} else {
				l := toDouble(env, l_register)
				r := toDouble(env, r_register)
				// LLIR: %r= add i32 %l, %r
				result = env.Block().NewFSub(l, r)
			}
		case "*":
			if arithmetic_type != types.Double {
				// LLIR: %r= mul i32 %l, %r
				result = env.Block().NewMul(l_register, r_register)
			} else {
				l := toDouble(env, l_register)
				r := toDouble(env, r_register)
				// LLIR: %r= add i32 %l, %r
				result = env.Block().NewFMul(l, r)
			}
		case "/":
			nextBlock := env.GetNewBlock("next")
			errBlock := env.GetNewErrorBlock(".error_division_by_zero")
			if arithmetic_type != types.Double {
				cmp := env.Block().NewICmp(enum.IPredEQ, r_register, constant.NewInt(types.I32, 0))
				env.Block().NewCondBr(cmp, errBlock, nextBlock)
				env.SetCurrentBlock(nextBlock)
				// LLIR: %r= sdiv i32 %l, %r
				result = env.Block().NewSDiv(l_register, r_register)
			} else {
				l := toDouble(env, l_register)
				r := toDouble(env, r_register)
				cmp := env.Block().NewFCmp(enum.FPredOEQ, r, constant.NewFloat(types.Double, 0))
				env.Block().NewCondBr(cmp, errBlock, nextBlock)
				env.SetCurrentBlock(nextBlock)

				// LLIR: %r= fdiv double %l, %r
				result = env.Block().NewFDiv(l, r)
			}
		case "%":
			if arithmetic_type != types.Double {
				// LLIR: %r= srem i32 %l, %r
				result = env.Block().NewSRem(l_register, r_register)
			} else {
				l := toDouble(env, l_register)
				r := toDouble(env, r_register)
				//TODO: if r is zero "division by zero"

				// LLIR: %r= frem double %l, %r
				result = env.Block().NewFRem(l, r)
			}
		case "<":
			result = compare(enum.IPredSLT, enum.FPredOLT)
		case ">":
			result = compare(enum.IPredSGT, enum.FPredOGT)
		case "<=":
			result = compare(enum.IPredSLE, enum.FPredOLE)
		case ">=":
			result = compare(enum.IPredSGE, enum.FPredOGE)
		case "==":
			result = compare(enum.IPredEQ, enum.FPredOEQ)
		case "!=":
			result = compare(enum.IPredNE, enum.FPredONE)
		case "&&":
			result = env.Block().NewAnd(l_register, r_register)
		case "||":
			result = env.Block().NewOr(l_register, r_register)
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
	if err == ErrUnknownSymbol || !val.Type().Equal(stored_value.Type()) {
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
