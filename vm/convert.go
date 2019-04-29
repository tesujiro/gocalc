package vm

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// precedenceOfTypes returns the greater of two kinds
// String > Double > Int
func precedenceOfTypes(type1, type2 types.Type) types.Type {
	if type1.Equal(type2) {
		return type1
	}
	CharArray0 := &types.ArrayType{ElemType: types.I8, Len: 0}
	list := []types.Type{CharArray0, types.Double, types.I32, types.I1}
	index := func(t types.Type) int {
		//if types.IsPointer(t) && types.IsArray(t.(*types.PointerType).ElemType) {
		if types.IsArray(t) {
			t_copy := *t.(*types.ArrayType)
			t = &t_copy
			t.(*types.ArrayType).Len = 0
		}
		for i, v := range list {
			if t.Equal(v) {
				return i
			}
		}
		return -1
	}
	//fmt.Printf("index(%v)=%v\n", type1, index(type1))
	//fmt.Printf("index(%v)=%v\n", type2, index(type2))
	if index(type1) < index(type2) {
		return type1
	} else {
		return type2
	}
}

func isString(v value.Value) bool {
	// String is a Pointer to Array of I8
	if !types.IsPointer(v.Type()) {
		return false
	}
	pointer_type := v.Type().(*types.PointerType)
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
