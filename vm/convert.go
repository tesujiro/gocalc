package vm

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// precedenceOfTypes returns the greater of two kinds
// Double > Int
func precedenceOfTypes(type1, type2 types.Type) types.Type {
	if type1 == type2 {
		return type1
	}
	list := []types.Type{types.Double, types.I32, types.I1}
	index := func(t types.Type) int {
		for i, v := range list {
			if t == v {
				return i
			}
		}
		return -1
	}
	if index(type1) < index(type2) {
		return type1
	} else {
		return type2
	}
}

func toDouble(env *Env, v value.Value) value.Value {
	switch v.Type() {
	case types.Double:
		return v
	case types.I1, types.I32:
		d := env.Block().NewUIToFP(v, types.Double)
		return d
	}
	return nil
}