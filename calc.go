package main

import (
	"fmt"

	"github.com/tesujiro/gocalc/parser"
	"github.com/tesujiro/gocalc/vm"
)

func main() {
	env := vm.NewEnv()

	scripts := []string{
		"1",
		"1+1",
		"1.1+1.2",
		"1.7-1.5",
		"1.1*1.2",
		"1.81+1.19",
		"1/3",
		"(1+1.2)*2.1+3.5",
	}

	//parser.TraceLexer()

	for _, source := range scripts {
		ast, parseError := parser.ParseSrc(source)
		if parseError != nil {
			fmt.Printf("Syntax error: %v \n", parseError)
			continue
		}
		result, err := vm.Run(ast, env)
		if err != nil {
			fmt.Printf("Runtime error: %v \n", err)
		}

		fmt.Printf("%v = %.10g\n", source, result)
	}
}
