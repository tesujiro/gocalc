package vm

import (
	"errors"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/tesujiro/gocalc/debug"
)

type Env struct {
	path     string
	env      map[string]value.Value
	lib      map[string]*ir.Func
	defs     map[string]*ir.Global
	parent   *Env
	module   *ir.Module
	fnc      *ir.Func
	block    *MyBlock
	cntBlock *ir.Block
	brkBlock *ir.Block
}

var ErrUnknownSymbol = errors.New("unknown symbol")
var AlreadyKnownSymbol = errors.New("already known symbol")

//var ErrDivisionByZero = errors.New("division by zero") //TODO

func NewEnv() *Env {
	func_name := "main"
	module := ir.NewModule()
	m := module.NewFunc(func_name, types.I32)
	//entry := m.NewBlock("entry")
	entry := &MyBlock{m.NewBlock("entry"), false}
	lib := make(map[string]*ir.Func)
	defs := make(map[string]*ir.Global)

	// LLIR: declare i32 @printf(i8* %format, ...)
	i8ptr := types.NewPointer(types.I8)
	printf := module.NewFunc("printf", types.I32, ir.NewParam("format", i8ptr))
	printf.FuncAttrs = []ir.FuncAttribute{enum.FuncAttrNoUnwind}
	printf.Sig.Variadic = true
	lib["printf"] = printf

	// LLIR: declare i32 @strcmp(i8* nocapture, i8* nocapture) local_unnamed_addr
	strcmp := module.NewFunc("strcmp", types.I32, ir.NewParam("str1", i8ptr), ir.NewParam("str2", i8ptr))
	strcmp.FuncAttrs = []ir.FuncAttribute{enum.FuncAttrNoUnwind, enum.FuncAttrReadOnly}
	lib["strcmp"] = strcmp

	// LLIR: declare i64 @strlen(i8* nocapture) local_unnamed_addr #2
	strlen := module.NewFunc("strlen", types.I64, ir.NewParam("str", i8ptr))
	strlen.FuncAttrs = []ir.FuncAttribute{enum.FuncAttrNoUnwind, enum.FuncAttrReadOnly, enum.FuncAttrArgMemOnly}
	lib["strlen"] = strlen

	// LLIR:
	strcpy := module.NewFunc("strcpy", types.I32, ir.NewParam("str1", i8ptr), ir.NewParam("str2", i8ptr))
	strcpy.FuncAttrs = []ir.FuncAttribute{enum.FuncAttrNoUnwind, enum.FuncAttrReadOnly}
	lib["strcpy"] = strcpy

	// LLIR:
	strcat := module.NewFunc("strcat", types.I32, ir.NewParam("str1", i8ptr), ir.NewParam("str2", i8ptr))
	strcat.FuncAttrs = []ir.FuncAttribute{enum.FuncAttrNoUnwind, enum.FuncAttrReadOnly}
	lib["strcat"] = strcat

	// LLIR: @.str.result = global [12 x i8] c"Result : %d\0A"
	defs[".print_int"] = module.NewGlobalDef(".print_int", constant.NewCharArrayFromString("%d\n\x00"))
	defs[".print_float"] = module.NewGlobalDef(".print_float", constant.NewCharArrayFromString("%g\n\x00"))
	defs[".print_string"] = module.NewGlobalDef(".print_string", constant.NewCharArrayFromString("%s\n\x00"))
	defs[".result"] = module.NewGlobalDef(".result", constant.NewCharArrayFromString("Result : %d\n\x00"))
	defs[".error"] = module.NewGlobalDef(".error", constant.NewCharArrayFromString("Runtime error : %s\n\x00"))
	defs[".error_division_by_zero"] = module.NewGlobalDef(".error_division_by_zero", constant.NewCharArrayFromString("division by zero\x00"))

	return &Env{
		path:   "/" + func_name,
		env:    make(map[string]value.Value),
		lib:    lib,
		defs:   defs,
		parent: nil,
		module: module,
		fnc:    m,
		block:  entry,
	}
}

func (e *Env) NewEnv(stmt_name string) *Env {
	return &Env{
		path:   e.path + "/" + stmt_name,
		env:    make(map[string]value.Value),
		lib:    e.lib,
		defs:   e.defs,
		parent: e,
		module: nil,
		fnc:    nil,
		block:  e.block,
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

type MyBlock struct {
	*ir.Block
	branched bool
}

func (mb *MyBlock) NewBr(target *ir.Block) *ir.TermBr {
	if mb.branched {
		panic("ALREADY BRANCHED")
	}
	mb.branched = true
	return mb.Block.NewBr(target)
}

func (mb *MyBlock) NewCondBr(cond value.Value, targetTrue, targetFalse *ir.Block) *ir.TermCondBr {
	if mb.branched {
		panic("ALREADY BRANCHED")
	}
	mb.branched = true
	return mb.Block.NewCondBr(cond, targetTrue, targetFalse)
}

func (e *Env) Block() *MyBlock {
	return e.block
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
	if !e.Block().branched {
		panic("No BRANCH")
	}
	debug.Printf("%v:SetCurrentBlock(%v->%v)\n", e.path, e.block, b)
	*e.block = MyBlock{b, false}
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
	debug.Printf("%v:SetContinueBlock\n", e.path)
	e.cntBlock = b
}

func (e *Env) GetContinueBlock() *ir.Block {
	debug.Printf("%v:GetContinueBlock\n", e.path)
	if e.cntBlock != nil {
		return e.cntBlock
	}
	if e.parent == nil {
		return nil
	}
	return e.parent.GetContinueBlock()
}

func (e *Env) SetBreakBlock(b *ir.Block) {
	debug.Printf("%v:SetBreakBlock\n", e.path)
	e.brkBlock = b
}

func (e *Env) GetBreakBlock() *ir.Block {
	debug.Printf("%v:GetBreakBlock\n", e.path)
	if e.brkBlock != nil {
		return e.brkBlock
	}
	if e.parent == nil {
		return nil
	}
	return e.parent.GetBreakBlock()
}
