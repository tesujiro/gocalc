package vm

import (
	"errors"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Env struct {
	env    map[string]value.Value
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
	return &Env{
		env:    make(map[string]value.Value),
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
