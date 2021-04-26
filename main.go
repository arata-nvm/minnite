package main

import "fmt"

func main() {
	eval("a=1+3;b=a+a;?b;")
}

func eval(s string) {
	var vars [256]int
	for i := 0; i < 10; i += 1 {
		vars['0'+i] = i
	}

	for i := 0; i < len(s); i += 1 {
		if s[i] == ';' {
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
