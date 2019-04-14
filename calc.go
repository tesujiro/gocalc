package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

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
		//result := runScript(strings.NewReader(source))
		result := runScript(source)
		if result != 0 {
			os.Exit(result)
		}
	}
}

func runScript(source string) int {
	env := vm.NewEnv()
	ast, parseError := parser.ParseSrc(source)
	if parseError != nil {
		fmt.Printf("Syntax error: %v \n", parseError)
		return 1
	}
	err := vm.Run(ast, env)
	if err != nil {
		fmt.Printf("Runtime error: %v \n", err)
		return 1
	}

	src_lli := env.Generate()

	// run llvm
	llvm_run := []string{"lli"}
	command := exec.Command(llvm_run[0], llvm_run[1:]...)

	stdin, _ := command.StdinPipe()
	io.WriteString(stdin, src_lli)
	stdin.Close()
	out, _ := command.Output()
	fmt.Printf("%s", out)

	return 0
}
