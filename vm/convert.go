package vm

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// precedenceOfTypes returns the greater of two kinds
// Double > Int
func precedenceOfTypes(type1, type2 types.Type) types.Type {
	if type1.Equal(type2) {
		return type1
	}
	list := []types.Type{types.Double, types.I32, types.I1}
	index := func(t types.Type) int {
		for i, v := range list {
			if t.Equal(v) {
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

func isString(v value.Value) bool {
	v_type := v.Type()
	if !types.IsPointer(v_type) {
		return false
	}
	pointer_type := v_type.(*types.PointerType)
	if !types.IsArray(pointer_type.ElemType) {
		return false
	}
	array_type := *pointer_type.ElemType.(*types.ArrayType) // Copy ArrayType
	array_type.Len = 0
	CharArray0 := &types.ArrayType{ElemType: types.I8, Len: 0}
	return array_type.Equal(CharArray0)
	//return strings.HasSuffix(v.Type().String(), "x i8]*")
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
