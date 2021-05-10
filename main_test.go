package main

import (
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		input    string
		expected Value
	}{
		{"42", NewInteger(42)},
		{"1 + 2 * 3", NewInteger(7)},
		{"(1 + 2) * 3", NewInteger(9)},
		{"5 / 2 + 5 % 2", NewInteger(3)},
		{"1 == 1", NewBoolean(true)},
		{"1 != 1", NewBoolean(false)},
		{"1 < 1", NewBoolean(false)},
		{"1 <= 1", NewBoolean(true)},
		{"1 < 1", NewBoolean(false)},
		{"1 <= 1", NewBoolean(true)},
		{"let hoge = 2; hoge", NewInteger(2)},
		{"if 1 == 1 { 2 }", NewInteger(2)},
		{"if 1 != 1 { 2 }", NewVoid()},
		{"if 1 == 1 { 2 } else { 3 }", NewInteger(2)},
		{"if 1 != 1 { 2 } else { 3 }", NewInteger(3)},
		{"let i = 0; let sum = 0; while i < 10 { sum = sum + i; i = i + 1 } sum", NewInteger(45)},
		{"return 42", NewInteger(42)},
		{"let hoge = func() { return 42 }; hoge()", NewInteger(42)},
		{"let hoge = func(a) { return a }; hoge(1)", NewInteger(1)},
		{"let hoge = func(a, b) { return a + b }; hoge(1, 2)", NewInteger(3)},
		{"let hoge = 1; hoge = 2; hoge", NewInteger(2)},
		{"let a = 2; let hoge = func() { let a = 4 return a }; hoge() + a", NewInteger(6)},
		{"let fib = func(x) { if x <= 1 { return 1 } else { return fib(x - 1) + fib(x - 2) } }; fib(5)", NewInteger(8)},
		{"let a = [1, 2, 3]; a[0] + a[1] + a[2]", NewInteger(6)},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			ctx := NewContext()
			actual := Exec(test.input, ctx)
			if actual != test.expected {
				t.Errorf("[FAILED] `%s` -> %d\n", test.input, actual)
			}
		})
	}

}
