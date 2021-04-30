package main

import (
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		input    string
		expected Value
	}{
		{"42;", NewInteger(42)},
		{"1 + 2 * 3;", NewInteger(7)},
		{"(1 + 2) * 3;", NewInteger(9)},
		{"5 / 2 + 5 % 2;", NewInteger(3)},
		{"1 == 1;", NewBoolean(true)},
		{"1 != 1;", NewBoolean(false)},
		{"1 < 1;", NewBoolean(false)},
		{"1 <= 1;", NewBoolean(true)},
		{"1 < 1;", NewBoolean(false)},
		{"1 <= 1;", NewBoolean(true)},
		{"let hoge = 2; hoge;", NewInteger(2)},
		{"if 1 == 1 { 2; };", NewInteger(2)},
		{"if 1 != 1 { 2; };", NewVoid()},
		{"if 1 == 1 { 2; } else { 3; }", NewInteger(2)},
		{"if 1 != 1 { 2; } else { 3; }", NewInteger(3)},
		{"let i = 0; let sum = 0; while i < 10 { let sum = sum + i; let i = i + 1; }; sum;", NewInteger(45)},
	}

	for _, test := range tests {
		ctx := Context{}
		actual := Exec(test.input, ctx)
		if actual != test.expected {
			t.Errorf("[FAILED] `%s` -> %d\n", test.input, actual)
		}
	}

}
