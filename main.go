package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

func main() {
	repl()
}

func repl() {
	ctx := map[string]int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		ast := parse(line)
		ast.Eval(ctx)
	}
}

func parse(s string) *Program {
	lexerDef := stateful.MustSimple([]stateful.Rule{
		{`Ident`, `[a-zA-Z][a-zA-Z_\d]*`, nil},
		{`Number`, `\d+`, nil},
		{`Operator`, `(==|!=|<=|<|>=|>)`, nil},
		{`Punct`, `[+\-*/%()=?;]`, nil},
		{"whitespace", `[\n\r\s]+`, nil},
	})

	parser, err := participle.Build(&Program{}, participle.Lexer(lexerDef))
	if err != nil {
		log.Fatal(err)
	}
	program := &Program{}
	err = parser.ParseString("", s, program)
	if err != nil {
		fmt.Print(err)
	}

	return program
}
