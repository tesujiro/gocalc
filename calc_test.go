package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

const command_name = "calc"

type test struct {
	script  string
	options []string
	ok      string
	rc      int
}

func TestCalc(t *testing.T) {
	tests := []test{
		//BASIC EXPRESSION
		{script: "print 1+1", ok: "2\n"},
		{script: "print 1+1;#comment", ok: "2\n"},
		{script: "print 1+2", ok: "3\n"},
	}

	//realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr
	case_number := 0

	for _, test := range tests {
		case_number++
		wg := &sync.WaitGroup{}

		// OUT PIPE
		readFromOut, writeToOut, err := os.Pipe()
		if err != nil {
			//os.Stdin = realStdin
			os.Stderr = realStderr
			t.Fatal("Pipe error:", err)
		}
		os.Stdout = writeToOut
		//logger.Print("pipe out created")

		// Read Stdout goroutine
		readerOut := bufio.NewScanner(readFromOut)
		chanOut := make(chan string)
		wg.Add(1)
		go func() {
			for readerOut.Scan() {
				chanOut <- readerOut.Text()
			}
			close(chanOut)
			wg.Done()
			return
		}()

		// Run Script goroutine
		wg.Add(1)
		go func() {

			os.Args = []string{command_name}
			os.Args = append(os.Args, test.options...)
			if test.script != "" {
				os.Args = append(os.Args, test.script)
			}
			rc := _main()
			if rc != test.rc && !strings.Contains(test.ok, "error") {
				t.Errorf("return code want:%v get:%v case:%v\n", test.rc, rc, test)
			}

			/*
				rc := runScript(script_reader, os.Stdin)
				if rc != 0 {
					t.Fatal("runscript return code:", rc)
				}
			*/
			//close(chanDone) //NG
			writeToOut.Close()
			wg.Done()
		}()

		// Get Result
		var resultOut string
	LOOP:
		for {
			select {
			case dataOut, ok := <-chanOut:
				if !ok {
					break LOOP
				}
				dataOut = strings.TrimSpace(dataOut)
				resultOut = fmt.Sprintf("%s%s%s", resultOut, dataOut, "\n")
			}
		}

		// Result Check
		//fmt.Fprintf(realStdout, "result:[%v]\ttest.ok:[%v]\n", resultOut, test.ok)
		if test.ok != "" && resultOut != strings.Replace(test.ok, "\r", "", -1) { //replace for Windows
			t.Errorf("Case:[%v] received: %v - expected: %v - runSource: %v", case_number, resultOut, test.ok, test.script)
		}

		wg.Wait()
		//readFromIn.Close()
		//writeToIn.Close()
		readFromOut.Close()
		writeToOut.Close()
		//os.Stdin = realStdin
		os.Stderr = realStderr
		os.Stdout = realStdout
	}

}
