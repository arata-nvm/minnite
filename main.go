package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

func main() {
	repl()
}

func repl() {
	vars := map[string]int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		tokens := tokenize(line)
		eval(tokens, vars)
	}
}

// 文字列を受け取り、トークン列にする
func tokenize(s string) []lexer.Token {
	// トークンの定義
	lexerDef := stateful.MustSimple([]stateful.Rule{
		{`Ident`, `[a-zA-Z][a-zA-Z_\d]*`, nil},
		{`Number`, `\d+`, nil},
		{`Punct`, `[+\-*/!=?;]`, nil},
		{"whitespace", `[\n\r\s]+`, nil},
	})

	// トークンをスライスに保持する
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

// トークン列を受け取り、それを評価する
func eval(tokens []lexer.Token, vars map[string]int) {
	i := 0

	// EOF=文字列の終わりまで繰り返す
	for !tokens[i].EOF() {

		// 文の最初の文字で分岐する
		switch tokens[i].String() {
		// let文
		case "let":
			name := tokens[i+1].String()
			value := evalValue(tokens[i+3], vars)

			if tokens[i+4].String() == "+" {
				value += evalValue(tokens[i+5], vars)
			}

			vars[name] = value

		// print文
		case "print":
			value := evalValue(tokens[i+1], vars)
			fmt.Printf("%d\n", value)

		// それ以外はエラーとする
		default:
			fmt.Println("エラー")
			return
		}

		// セミコロンの次まで進める
		for {
			i += 1
			if tokens[i].String() == ";" {
				i += 1
				break
			}
		}
	}
}

// トークンを受けとり、数値にして返す
func evalValue(token lexer.Token, vars map[string]int) int {
	var (
		ident  rune = -2
		number rune = -3
	)

	// 変数ならばその値を取得する
	if token.Type == ident {
		return vars[token.String()]
	}

	// 数字なら数値に変換する
	if token.Type == number {
		value, _ := strconv.Atoi(token.String())
		return value
	}

	// 変数・数字以外はエラーとする
	panic("エラー")
}
