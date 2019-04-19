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
		//STATEMENT
		{script: "1", ok: ""},
		{script: "", rc: 1, ok: ""},
		{script: "1\n2", ok: ""},
		{script: "xxx", rc: 1, ok: "Compile error: unknown symbol\n"},
		{script: "((", rc: 1, ok: "syntax error\n"},
		//BASIC EXPRESSION
		{script: "print 1", ok: "1\n"},
		{script: "print 0", ok: "0\n"},
		{script: "print +1", ok: "1\n"},
		{script: "print -1", ok: "-1\n"},
		{script: "print -1", ok: "-1\n"},
		{script: "print 1+1", ok: "2\n"},
		{script: "print 1+1;#comment", ok: "2\n"},
		{script: "print 1+2", ok: "3\n"},
		{script: "print 2-1", ok: "1\n"},
		{script: "print 1-2", ok: "-1\n"},
		{script: "print 1*2", ok: "2\n"},
		{script: "print 2/1", ok: "2\n"},
		{script: "print 7/3", ok: "2\n"},
		{script: "print 7%3", ok: "1\n"},
		{script: "print -1*-2", ok: "2\n"},
		{script: "print (1+30)*2", ok: "62\n"},
		{rc: 1, ok: "No expression error!\nex: calc '(1+1)*3+10' ; echo $?\n"},
		{script: "xxxx", rc: 1, ok: "Compile error: unknown symbol\n"},

		// Float
		{script: "print 1.23", ok: "1.23\n"},
		{script: "print 1.23+2", ok: "3.23\n"},
		{script: "print 1.23+2.15", ok: "3.38\n"},
		{script: "print 1.23-2", ok: "-0.77\n"},
		{script: "print 1.23-2.15", ok: "-0.92\n"},
		{script: "print 1.23*2", ok: "2.46\n"},
		{script: "print 1.23*2.15", ok: "2.6445\n"},
		{script: "print 1.23/2", ok: "0.615\n"},
		{script: "print 2.23%2", ok: "0.23\n"},
		//{script: "print 1.23/0", rc: 1, ok: "Runtime error: division by zero\n"}, //TODO:division by zero

		//BOOL EXPRESSION
		{script: "print 1==1", ok: "1\n"},
		{script: "print 1==0", ok: "0\n"},
		{script: "print 1!=1", ok: "0\n"},
		{script: "print 1!=0", ok: "1\n"},
		{script: "print !(1==1)", ok: "0\n"},
		{script: "print !(1==0)", ok: "1\n"},
		{script: "print 2>1", ok: "1\n"},
		{script: "print 2>1", ok: "1\n"},
		{script: "print 2<1", ok: "0\n"},
		{script: "print 1>=1", ok: "1\n"},
		{script: "print 0>=1", ok: "0\n"},
		{script: "print 1<=1", ok: "1\n"},
		{script: "print 1<=0", ok: "0\n"},
		{script: "print 1<=1&&2>=2", ok: "1\n"},
		{script: "print 1<=1&&1>=2", ok: "0\n"},
		{script: "print 1<=0&&2>=2", ok: "0\n"},
		{script: "print 1<=0&&1>=2", ok: "0\n"},
		{script: "print 1<=1||2>=2", ok: "1\n"},
		{script: "print 1<=1||1>=2", ok: "1\n"},
		{script: "print 1<=0||2>=2", ok: "1\n"},
		{script: "print 1<=0||1>=2", ok: "0\n"},

		//ASSIGNMENT
		{script: "i=1;print i", ok: "1\n"},
		{script: "i=1;i=2.3;print i", ok: "2.3\n"},
		{script: "i=1;i=2.3;i=1;print i", ok: "1\n"},
		{script: "i=1.2;print i", ok: "1.2\n"},
		{script: "i=1;j=2;print i*10+j", ok: "12\n"},
		{script: "i=1>0;print i", ok: "1\n"},
		{script: "print j", ok: "Compile error: unknown symbol\n", rc: 1},

		//COMPOSITE EXPRESSION
		{script: "i=1;++i;print i", ok: "2\n"},
		{script: "i=1;i++;print i", ok: "2\n"},
		{script: "i=1;--i;print i", ok: "0\n"},
		{script: "i=1;i--;print i", ok: "0\n"},
		{script: "i=1;i=++i;print i", ok: "2\n"},
		{script: "i=1;i=--i;print i", ok: "0\n"},
		{script: "i=1;i=i++;print i", ok: "2\n"},
		{script: "i=1;i=i--;print i", ok: "0\n"},
		{script: "i=1;i+=2;print i", ok: "3\n"},
		{script: "i=1;i+=1.2;print i", ok: "2.2\n"},
		{script: "i=1;i-=2;print i", ok: "-1\n"},
		{script: "i=1.1;i-=1.2;print i", ok: "-0.1\n"},
		{script: "i=1.1;i*=1.2;print i", ok: "1.32\n"},
		{script: "i=17;i/=4;print i", ok: "4\n"},
		{script: "i=17;i%=4;print i", ok: "1\n"},

		//IF STMT
		{script: "if 1>0 {print 1};print 2", ok: "1\n2\n"},
		{script: "if 1<0 {print 1};print 2", ok: "2\n"},
		{script: "if 1>0 {print 1}else{print 2};print 3", ok: "1\n3\n"},
		{script: "if 1<0 {print 1}else{print 2};print 3", ok: "2\n3\n"},
		{script: "i=0;if i>0 {print 1};print 2", ok: "2\n"},
		{script: "i=1;if i>0 {print 1};print 2", ok: "1\n2\n"},
		{script: "i=0;if i>0 {if i>1 {print 1}else{print 2}};print 3", ok: "3\n"},
		{script: "i=1;if i>0 {if i>1 {print 1}else{print 2}};print 3", ok: "2\n3\n"},
		{script: "i=2;if i>0 {if i>1 {print 1}else{print 2}};print 3", ok: "1\n3\n"},

		//FOR STMT
		{script: "for i=1;i<5;i++{print i}", ok: "1\n2\n3\n4\n"},
		{script: "for i=1;i<5;i++{for j=1;j<3;j++{print i}}", ok: "1\n1\n2\n2\n3\n3\n4\n4\n"},
	}

	//realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr
	case_number := 0

	for _, test := range tests {
		case_number++
		wg := &sync.WaitGroup{}

		//fmt.Printf("TEST[%v] %v\n", case_number, test.script)

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
				t.Errorf("return code want:%v get:%v case:%v\n", test.rc, rc, test.script)
			}

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
