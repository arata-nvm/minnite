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
	ctx := map[string]int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		ast := parse(line)
		ast.eval(ctx)
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

type Context map[string]int

func (p *Program) eval(ctx Context) {
	for _, stmt := range p.Statements {
		stmt.eval(ctx)
	}
}

func (s *Statement) eval(ctx Context) {
	switch {
	case s.Let != nil:
		s.Let.eval(ctx)
	case s.Print != nil:
		s.Print.eval(ctx)
	}
}

func (s *LetStatement) eval(ctx Context) {
	ctx[s.Variable] = s.Value.eval(ctx)
}

func (s *PrintStatement) eval(ctx Context) {
	fmt.Printf("%d\n", s.Value.eval(ctx))

}

func (t *Term) eval(ctx Context) int {
	switch {
	case t.Variable != nil:
		return ctx[*t.Variable]
	case t.Number != nil:
		return *t.Number
	}

	return 0
}
