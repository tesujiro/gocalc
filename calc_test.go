package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"
)

const command_name = "calc"

type test struct {
	script   string
	options  []string
	ok       string
	ok_regex string
	rc       int
}

func TestCalc(t *testing.T) {
	tests := []test{
		//STATEMENT
		{script: "1", ok: ""},
		{script: "", rc: 1, ok: ""},
		{script: "1\n2", ok: ""},
		{script: "xxx", rc: 1, ok: "Compile error: unknown symbol\n"},
		{script: "((", rc: 1, ok: "syntax error\n"},
		//OPTIONS
		{script: "1", options: []string{"-a"}, ok_regex: `ast.NumExpr{Literal:"1"}`},
		{script: "1", options: []string{"-i"}, ok_regex: `define i32 @main\(\) {`},
		{script: "for i=0;i<1;i++{print i}", options: []string{"-d"}, ok_regex: `debug option`},
		{script: "1", options: []string{"-n"}, ok: ""},

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
		{script: "print 7%3.1", ok: "0.8\n"},
		{script: "print 7.1%3.1", ok: "0.9\n"},
		{script: "print -1*-2", ok: "2\n"},
		{script: "print (1+30)*2", ok: "62\n"},
		{rc: 1, ok: "No expression error!\nex: calc '(1+1)*3+10' ; echo $?\n"},
		{script: "xxxx", rc: 1, ok: "Compile error: unknown symbol\n"},

		// String Assign
		{script: `s="ABC";print s`, ok: "ABC\n"},
		{script: `s='ABC';print s`, ok: "ABC\n"},
		{script: `s="\n";print s`, ok: "\n\n"},
		{script: `s="a\tb";print s`, ok: "a\tb\n"},
		{script: `s="\b";print s`, ok: "\b\n"},
		{script: `s="\r";print s`, ok: "\r\n"},
		{script: `s="a\fb";print s`, ok: "a\fb\n"},
		{script: `s=1;s="ABC";print s`, ok: "ABC\n"},
		{script: `s="ABC";s=123;print s`, ok: "123\n"},
		// Length
		{script: `s="ABC";print len(s)`, ok: "3\n"},
		{script: `s=123;print len(s)`, ok: "Compile error: arg type error: i32*\n"},
		{script: `print len(s)`, ok: "Compile error: unknown symbol\n"},
		// String Compare
		{script: `s="ABC";print s=="ABC"`, ok: "1\n"},
		{script: `s="ABC";print s=="DEF"`, ok: "0\n"},
		{script: `s="ABC";print s!="ABC"`, ok: "0\n"},
		{script: `s="ABC";print s>"ABC"`, ok: "0\n"},
		{script: `s="ABC";print s>="ABC"`, ok: "1\n"},
		{script: `s="ABC";print s<="ABC"`, ok: "1\n"},
		{script: `s="ABC";print s<"ABC"`, ok: "0\n"},
		{script: `s="ABC";print s>"AB"`, ok: "1\n"},
		{script: `s="ABC";print s>"ABCD"`, ok: "0\n"},
		{script: `s="ABC";print s<"AB"`, ok: "0\n"},
		{script: `s="ABC";print s<"ABCD"`, ok: "1\n"},
		{script: `s="ABC";print s==1`, ok: "0\n"},
		{script: `s=1;print s=="ABC"`, ok: "0\n"},
		// String operation
		{script: `s="A"+"BC";print s`, ok: "ABC\n"},
		{script: `s="A";s+="BC";print s`, ok: "ABC\n"},
		{script: `s="A";s+="A";print s`, ok: "AA\n"},
		//{script: `s="A";for i=0;i<3;i++{s=s+"A"};print s`, ok: "AAAA\n"},
		//{script: `s="A";for i=0;i<3;i++{s+="A"};print s`, ok: "AAAA\n"},

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
		{script: "print 1/0", rc: 1, ok: "Runtime error : division by zero\n"},
		{script: "print 1.23/0", rc: 1, ok: "Runtime error : division by zero\n"},
		{script: "print 1%0", rc: 1, ok: "Runtime error : division by zero\n"},
		{script: "print 1.23%0", rc: 1, ok: "Runtime error : division by zero\n"},

		//BOOL EXPRESSION
		{script: "print 1==1", ok: "1\n"},
		{script: "print 1.1==1.1", ok: "1\n"},
		{script: "print 1.1==1", ok: "0\n"},
		{script: "print 1==0", ok: "0\n"},
		{script: "print 1!=1", ok: "0\n"},
		{script: "print 1.1!=1.1", ok: "0\n"},
		{script: "print 1!=0", ok: "1\n"},
		{script: "print !(1==1)", ok: "0\n"},
		{script: "print !(1==0)", ok: "1\n"},
		{script: "print 2>1", ok: "1\n"},
		{script: "print 2>1", ok: "1\n"},
		{script: "print 2>1.1", ok: "1\n"},
		{script: "print 2<1", ok: "0\n"},
		{script: "print 1>=1", ok: "1\n"},
		{script: "print 0>=1", ok: "0\n"},
		{script: "print 0.1>=1", ok: "0\n"},
		{script: "print 1<=1", ok: "1\n"},
		{script: "print 1.1<=1.1", ok: "1\n"},
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
		{script: "i=1;i=123;print i", ok: "123\n"},
		{script: "i=1;i=2.3;print i", ok: "2.3\n"},
		{script: "i=1;i=2.3;i=1;print i", ok: "1\n"},
		{script: "i=1.2;print i", ok: "1.2\n"},
		{script: "i=1;j=2;print i*10+j", ok: "12\n"},
		{script: "i=1>0;print i", ok: "1\n"},
		{script: "print j", ok: "Compile error: unknown symbol\n", rc: 1},
		{script: "i,j=1,2;print i;print j", ok: "1\n2\n"},
		{script: "i,j=1,2,3;print i;print j", ok: "1\n2\n"},
		{script: "i,j=1;print i", ok: "1\n"},
		{script: "print i", ok: "Compile error: unknown symbol\n"},
		{script: "print -i", ok: "Compile error: unknown symbol\n"},
		{script: "i=-j", ok: "Compile error: unknown symbol\n"},
		{script: "i,j=1,x;print i;print j", ok: "Compile error: unknown symbol\n"},
		{script: "print i+1", ok: "Compile error: unknown symbol\n"},
		{script: "print 1+i", ok: "Compile error: unknown symbol\n"},

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
		{script: "i+=j", ok: "Compile error: unknown symbol\n"},

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
		{script: "if i>0 {print 1}", ok: "Compile error: unknown symbol\n"},
		{script: "if 1>0 {print i}", ok: "Compile error: unknown symbol\n"},
		{script: "if 1>0 {}else{print j}", ok: "Compile error: unknown symbol\n"},

		//FOR STMT
		{script: "for i=1;i<5;i++{print i}", ok: "1\n2\n3\n4\n"},
		{script: "for i=1;i<5;i++{for j=1;j<3;j++{print i}}", ok: "1\n1\n2\n2\n3\n3\n4\n4\n"},
		{script: "for i=1;i<5;i++{print i;if i==2 {break}}", ok: "1\n2\n"},
		{script: "for i=1;i<5;i++{if i<=2 {continue};print i}", ok: "3\n4\n"},
		{script: "for i=1;i<5;i++{print i;break}", ok: "1\n"},
		{script: "for i=1;i<5;i++{print i;continue}", ok: "1\n2\n3\n4\n"},
		{script: "for i=1;;i++{print i;if i==2 {break}else{continue}}", ok: "1\n2\n"},
		{script: "for i=1;i<5;i++{print i;if i==2 {continue}else{break}}", ok: "1\n"},
		{script: "for i=1;i<5;i++{if 1==1{ if i<=2 {continue};print i}}", ok: "3\n4\n"},
		{script: "for i=1;i<5;i++{if i<4{ if i<=2 {continue};print i}else{break}}", ok: "3\n"},
		{script: "for i=1;i<5;i++{for j=1;j<3;j++{break};print i;break}", ok: "1\n"},
		//{script: "for i=1;i<5;i++{for break;j<3;j++{print i}}", ok: "\n"}, //TODO??
		{script: "break", ok: "Compile error: break not inside loop\n"},
		{script: "continue", ok: "Compile error: continue not inside loop\n"},
		{script: "for break;i<5;i++{print i}", ok: "Compile error: for init stmt error: unexpected break\n"},
		{script: "for i=1;j<5;i++{print i}", ok: "Compile error: for condition expr error: unknown symbol\n"},
		{script: "for i=1;i<5;j++{print i}", ok: "Compile error: for final expr error: unknown symbol\n"},
		{script: "for i=1;i<5;i++{print j}", ok: "Compile error: for loop stmts error: unknown symbol\n"},
		{script: `count=1;for i=0;i<3;i++{count+=1};print count`, ok: "4\n"},
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
			//fmt.Fprintf(realStdout, "case:%d os.Args=%v *print_ast=%v\n", case_number, os.Args, *print_ast)
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
		if test.ok_regex != "" {
			r := regexp.MustCompile(test.ok_regex)
			if !r.MatchString(resultOut) {
				t.Errorf("Case:[%v] received: %v - expected(regexp): %v - runSource: %v", case_number, resultOut, test.ok_regex, test.script)
			}
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
