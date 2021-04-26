package main

import (
	"bufio"
	"fmt"
	"os"
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
		eval(line, vars[:])
	}
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
