package vm

import (
	"errors"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Env struct {
	env    map[string]value.Value
	lib    map[string]*ir.Func
	defs   map[string]*ir.Global
	parent *Env
	module *ir.Module
	entry  *ir.Block
}

var ErrUnknownSymbol = errors.New("unknown symbol")
var AlreadyKnownSymbol = errors.New("already known symbol")

func NewEnv() *Env {
	module := ir.NewModule()
	m := module.NewFunc("main", types.I32)
	entry := m.NewBlock("")
	lib := make(map[string]*ir.Func)
	defs := make(map[string]*ir.Global)

	// LLIR: declare i32 @printf(i8* %format, ...)
	i8ptr := types.NewPointer(types.I8)
	printf := module.NewFunc("printf", types.I32, ir.NewParam("format", i8ptr))
	printf.Sig.Variadic = true
	lib["printf"] = printf

	// LLIR: @.str.result = global [12 x i8] c"Result : %d\0A"
	defs[".print_int"] = module.NewGlobalDef(".print_int", constant.NewCharArrayFromString("%d\n\x00"))
	defs[".result"] = module.NewGlobalDef(".result", constant.NewCharArrayFromString("Result : %d\n\x00"))

	return &Env{
		env:    make(map[string]value.Value),
		lib:    lib,
		defs:   defs,
		parent: nil,
		module: module,
		entry:  entry,
	}
}

func (e *Env) Get(id string) (value.Value, error) {
	if v, ok := e.env[id]; ok {
		return v, nil
	}
	if e.parent == nil {
		return nil, ErrUnknownSymbol
	}
	return e.parent.Get(id)
}

func (e *Env) Set(k string, v value.Value) error {
	e.env[k] = v
	return nil
}

func (e *Env) Generate() {
	fmt.Println(e.module)
}
