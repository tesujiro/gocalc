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
	fnc    *ir.Func
	block  *ir.Block
}

var ErrUnknownSymbol = errors.New("unknown symbol")
var AlreadyKnownSymbol = errors.New("already known symbol")

func NewEnv() *Env {
	module := ir.NewModule()
	m := module.NewFunc("main", types.I32)
	entry := m.NewBlock("entry")
	lib := make(map[string]*ir.Func)
	defs := make(map[string]*ir.Global)

	// LLIR: declare i32 @printf(i8* %format, ...)
	i8ptr := types.NewPointer(types.I8)
	printf := module.NewFunc("printf", types.I32, ir.NewParam("format", i8ptr))
	printf.Sig.Variadic = true
	lib["printf"] = printf

	// LLIR: @.str.result = global [12 x i8] c"Result : %d\0A"
	defs[".print_int"] = module.NewGlobalDef(".print_int", constant.NewCharArrayFromString("%d\n\x00"))
	defs[".print_float"] = module.NewGlobalDef(".print_float", constant.NewCharArrayFromString("%g\n\x00"))
	defs[".result"] = module.NewGlobalDef(".result", constant.NewCharArrayFromString("Result : %d\n\x00"))

	return &Env{
		env:    make(map[string]value.Value),
		lib:    lib,
		defs:   defs,
		parent: nil,
		module: module,
		fnc:    m,
		block:  entry,
	}
}

func (e *Env) NewEnv() *Env {
	return &Env{
		env:    make(map[string]value.Value),
		lib:    e.lib,
		defs:   e.defs,
		parent: e,
		module: nil,
		fnc:    nil,
		block:  nil,
	}
}

func (e *Env) GetVar(id string) (value.Value, error) {
	if v, ok := e.env[id]; ok {
		return v, nil
	}
	if e.parent == nil {
		return nil, ErrUnknownSymbol
	}
	return e.parent.GetVar(id)
}

func (e *Env) SetVar(k string, v value.Value) error {
	e.env[k] = v
	return nil
}

func (e *Env) moduleScope() *Env {
	if e.module != nil || e.parent == nil {
		return e
	}
	return e.parent.moduleScope()
}

func (e *Env) funcScope() *Env {
	if e.fnc != nil || e.parent == nil {
		return e
	}
	return e.parent.funcScope()
}

func (e *Env) blockScope() *Env {
	if e.block != nil || e.parent == nil {
		return e
	}
	return e.parent.blockScope()
}

func (e *Env) Module() *ir.Module {
	if e.module != nil || e.parent == nil {
		return e.module
	}
	return e.parent.Module()
}

func (e *Env) Func() *ir.Func {
	if e.fnc != nil || e.parent == nil {
		return e.fnc
	}
	return e.parent.Func()
}

func (e *Env) Block() *ir.Block {
	if e.block != nil || e.parent == nil {
		return e.block
	}
	return e.parent.Block()
}

func (e *Env) Generate() string {
	return fmt.Sprintln(e.module)
}

//TODO: GetNewFunc
//TODO: SetCurrentFunc

func (e *Env) GetNewBlock(id string) *ir.Block {
	// LLIR: ; <label>:(id)xx
	block := e.funcScope().fnc.NewBlock(id)
	return block
}

func (e *Env) SetCurrentBlock(b *ir.Block) {
	//fmt.Printf("SetCurrentBlock: %#v\n", b)
	e.funcScope().block = b
}
