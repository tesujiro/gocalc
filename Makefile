all : calc.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go ./parser/grammar.go
	go build -o calc .

./parser/grammar.go : ./parser/grammar.go.y ./ast/*.go
	goyacc -o ./parser/grammar.go ./parser/grammar.go.y
	gofmt -s -w ./parser

.PHONY: test
test: ./*_test.go
	go vet ./...
	go test -v -count=1 . -coverpkg ./...

stack_unlimited: # for the error "Illegal instruction: 4"
	ulimit -Ss unlimited
	ulimit -Ss
