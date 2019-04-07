package main

import (
	"fmt"
	"os"

	"github.com/tesujiro/gocalc/parser"
	"github.com/tesujiro/gocalc/vm"
)

func main() {

	//parser.TraceLexer()

	if len(os.Args) < 2 {
		fmt.Println("No expression error!")
		fmt.Printf("ex: %v '(1+1)*3+10' | lli ; echo $?\n", os.Args[0])
		os.Exit(1)
	}

	for _, source := range os.Args[1:] {
		//fmt.Printf("source: %v\n", source)
		env := vm.NewEnv()
		ast, parseError := parser.ParseSrc(source)
		if parseError != nil {
			fmt.Printf("Syntax error: %v \n", parseError)
			continue
		}
		err := vm.Run(ast, env)
		if err != nil {
			fmt.Printf("Runtime error: %v \n", err)
		}

		env.Generate()
	}
}
