package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let   *LetStatement   `( @@ `
	Print *PrintStatement `| @@ ) ";"`
}

type LetStatement struct {
	Variable string `"let" @Ident`
	Value    *Term  `"=" @@`
}

type PrintStatement struct {
	Value *Term `"print" @@`
}

type Term struct {
	Variable *string `@Ident`
	Number   *int    `| @Number`
}

func main() {
	repl()
}

func repl() {
	// vars := map[string]int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		ast := parse(line)
		fmt.Printf("%#v\n", ast)
	}
}

func parse(s string) *Program {
	lexerDef := stateful.MustSimple([]stateful.Rule{
		{`Ident`, `[a-zA-Z][a-zA-Z_\d]*`, nil},
		{`Number`, `\d+`, nil},
		{`Punct`, `[+\-*/!=?;]`, nil},
		{"whitespace", `[\n\r\s]+`, nil},
	})

	parser, err := participle.Build(&Program{}, participle.Lexer(lexerDef))
	if err != nil {
		log.Fatal(err)
	}
	program := &Program{}
	parser.ParseString("", s, program)

	return program
}
