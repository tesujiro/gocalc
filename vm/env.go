package vm

type Env struct {
	env    map[string]interface{}
	parent *Env
	prec   uint
}

func NewEnv() *Env {
	return &Env{
		env:    make(map[string]interface{}),
		parent: nil,
		prec:   64,
	}
}
