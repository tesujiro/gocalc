package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/tesujiro/gocalc/parser"
	"github.com/tesujiro/gocalc/vm"
)

var print_ir = flag.Bool("i", false, "print llvm ir")
var no_exec = flag.Bool("n", false, "no execution")

func main() {
	os.Exit(_main())
}

func _main() int {

	//parser.TraceLexer()
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("No expression error!")
		fmt.Printf("ex: %v '(1+1)*3+10' | lli ; echo $?\n", os.Args[0])
		return 1
	}

	for _, source := range args {
		//fmt.Printf("source: %v\n", source)
		//result := runScript(strings.NewReader(source))
		result := runScript(source)
		if result != 0 {
			return result
		}
	}
	return 0
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
		fmt.Printf("Compile error: %v \n", err)
		return 1
	}

	src_lli := env.Generate()

	if *print_ir {
		fmt.Println(src_lli)
	}

	if *no_exec {
		return 0
	}

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
