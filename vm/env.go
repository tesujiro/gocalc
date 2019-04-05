package vm

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Env struct {
	env    map[string]interface{}
	parent *Env
	module *ir.Module
	entry  *ir.Block
}

func NewEnv() *Env {
	module := ir.NewModule()
	m := module.NewFunc("main", types.I32)
	entry := m.NewBlock("")
	return &Env{
		env:    make(map[string]interface{}),
		parent: nil,
		module: module,
		entry:  entry,
	}
}

func (env *Env) Generate() {
	fmt.Println(env.module)
}
