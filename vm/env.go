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
	env      map[string]value.Value
	lib      map[string]*ir.Func
	defs     map[string]*ir.Global
	parent   *Env
	module   *ir.Module
	fnc      *ir.Func
	block    *ir.Block
	cntBlock *ir.Block
	brkBlock *ir.Block
}

var ErrUnknownSymbol = errors.New("unknown symbol")
var AlreadyKnownSymbol = errors.New("already known symbol")

//var ErrDivisionByZero = errors.New("division by zero") //TODO

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
	defs[".error"] = module.NewGlobalDef(".error", constant.NewCharArrayFromString("Runtime error : %s\n\x00"))
	defs[".error_division_by_zero"] = module.NewGlobalDef(".error_division_by_zero", constant.NewCharArrayFromString("division by zero\x00"))

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

var label_number = 0 //TODO move global variable to env struct member

func (e *Env) newLabel(id string) string {
	label_number++
	return fmt.Sprintf("%d:%s", label_number, id)
}

func (e *Env) GetNewBlock(id string) *ir.Block {
	// LLIR: ; <label>:(id)xx
	block := e.funcScope().fnc.NewBlock(e.newLabel(id))
	return block
}

func (e *Env) SetCurrentBlock(b *ir.Block) {
	//fmt.Printf("SetCurrentBlock: %#v\n", b)
	e.funcScope().block = b
}

func (e *Env) GetNewErrorBlock(msg_key string) *ir.Block {
	// LLIR: ; <label>:(id)xx
	block := e.funcScope().fnc.NewBlock(e.newLabel("error"))

	// LLIR: %8 = call i32 (i8*, ...) @printf(i8* getelementptr ([12 x i8], [12 x i8]* @.str.result, i32 0, i32 0), i32 %7)
	zero := constant.NewInt(types.I32, 0)
	msg := constant.NewGetElementPtr(e.defs[msg_key], zero, zero)
	block.NewCall(e.lib["printf"], constant.NewGetElementPtr(e.defs[".error"], zero, zero), msg)

	// EXIT 1
	block.NewRet(constant.NewInt(types.I32, 1))

	return block
}

func (e *Env) SetContinueBlock(b *ir.Block) {
	e.funcScope().cntBlock = b
}

func (e *Env) GetContinueBlock() *ir.Block {
	if e.funcScope().cntBlock != nil {
		return e.funcScope().cntBlock
	}
	if e.parent == nil {
		return nil
	}
	return e.parent.GetContinueBlock()
}

func (e *Env) SetBreakBlock(b *ir.Block) {
	e.funcScope().brkBlock = b
}

func (e *Env) GetBreakBlock() *ir.Block {
	if e.funcScope().brkBlock != nil {
		return e.funcScope().brkBlock
	}
	if e.parent == nil {
		return nil
	}
	return e.parent.GetBreakBlock()
}
