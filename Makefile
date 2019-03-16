all : calc.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go ./parser/grammar.go
	go build -o calc .

./parser/grammar.go : ./parser/grammar.go.y ./ast/*.go
	goyacc -o ./parser/grammar.go ./parser/grammar.go.y
	gofmt -s -w ./parser
