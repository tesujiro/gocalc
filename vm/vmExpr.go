package vm

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/tesujiro/gocalc/ast"
)

func Run(expr ast.Expr, env *Env) (interface{}, error) {
	return evalExpr(expr, env)
}

func isNumber(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Type() != reflect.TypeOf(big.NewFloat(0)) {
		return false
	}
	return true
}

func toNumber(i interface{}) *big.Float {
	v, _ := i.(*big.Float)
	return v
}

func evalExpr(expr ast.Expr, env *Env) (interface{}, error) {
	switch expr.(type) {
	case *ast.NumExpr:
		lit := expr.(*ast.NumExpr).Literal
		num := new(big.Float)
		num.SetPrec(env.prec)
		if _, ok := num.SetString(lit); !ok {
			return nil, fmt.Errorf("invalid number format:%v", lit)
		}
		return num, nil
	case *ast.BinOpExpr:
		var left, right interface{}
		var err error
		if left, err = evalExpr(expr.(*ast.BinOpExpr).Left, env); err != nil {
			return nil, err
		}
		if right, err = evalExpr(expr.(*ast.BinOpExpr).Right, env); err != nil {
			return nil, err
		}
		switch expr.(*ast.BinOpExpr).Operator {
		case "+", "-", "*", "/":
			switch {
			case isNumber(left) && isNumber(right):
				lnum := toNumber(left)
				rnum := toNumber(right)
				num := new(big.Float)
				num.SetPrec(env.prec)
				switch expr.(*ast.BinOpExpr).Operator {
				case "+":
					num.Add(lnum, rnum)
				case "-":
					num.Sub(lnum, rnum)
				case "*":
					num.Mul(lnum, rnum)
				case "/":
					num.Quo(lnum, rnum)
				}
				return num, nil
			default:
				return nil, fmt.Errorf("invalid binary operation: %v %v %v", left, expr.(*ast.BinOpExpr).Operator, right)
			}
		default:
			return nil, fmt.Errorf("invalid binary operator: %v", expr.(*ast.BinOpExpr).Operator)
		}
	case *ast.ParentExpr:
		sub := expr.(*ast.ParentExpr).SubExpr
		return evalExpr(sub, env)
	default:
		return nil, fmt.Errorf("invalid expression: %v", reflect.TypeOf(expr))
	}
}
