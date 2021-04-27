package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

func main() {
	repl()
}

func repl() {
	var vars [256]int
	for i := 0; i < 10; i += 1 {
		vars['0'+i] = i
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		tokens := tokenize(line)
		fmt.Printf("%+v\n", tokens)
	}
}

func tokenize(s string) []lexer.Token {
	lexerDef := stateful.MustSimple([]stateful.Rule{
		{`Ident`, `[a-zA-Z][a-zA-Z_\d]*`, nil},
		{`Number`, `\d+`, nil},
		{`Punct`, `[+\-*/!=?;]`, nil},
		{"whitespace", `[\n\r\s]+`, nil},
	})

	var tokens []lexer.Token
	l, _ := lexerDef.LexString("", s)
	for {
		token, _ := l.Next()
		tokens = append(tokens, token)
		if token.EOF() {
			break
		}
	}

	return tokens
}

func eval(s string, vars []int) {
	for i := 0; i < len(s); i += 1 {
		if s[i] == ';' || s[i] == '\n' {
			i += 1
			continue
		}

		if s[i+1] == '=' && s[i+3] == ';' {
			name := s[i]
			value := vars[s[i+2]]
			vars[name] = value
		} else if s[i+1] == '=' && s[i+3] == '+' && s[i+5] == ';' {
			name := s[i]
			v1 := vars[s[i+2]]
			v2 := vars[s[i+4]]
			vars[name] = v1 + v2
		}

		if s[i] == '?' {
			value := s[i+1]
			fmt.Printf("%d\n", vars[value])
		}

		for s[i] != ';' {
			i += 1
		}
	}
}
