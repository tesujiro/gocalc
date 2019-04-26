package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/tesujiro/gocalc/debug"
	"github.com/tesujiro/gocalc/parser"
	"github.com/tesujiro/gocalc/vm"
)

var print_ast = flag.Bool("a", false, "print AST")
var print_ir = flag.Bool("i", false, "print llvm ir")
var no_exec = flag.Bool("n", false, "no execution")
var dbg = flag.Bool("d", false, "debug option")

func main() {
	os.Exit(_main())
}

func _main() int {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
			for depth := 0; ; depth++ {
				_, file, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				log.Printf("=>%d: %v:%d", depth, file, line)
			}
		}
	}()

	//parser.TraceLexer()
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("No expression error!")
		fmt.Printf("ex: %v '(1+1)*3+10' ; echo $?\n", os.Args[0])
		return 1
	}

	if *dbg {
		debug.On()
	} else {
		debug.Off()
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
		fmt.Printf("%v\n", parseError)
		return 1
	}
	if *print_ast {
		parser.Dump(ast)
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
