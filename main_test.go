package main

import (
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42;", 42},
		{"1 + 2 * 3;", 7},
		{"(1 + 2) * 3;", 9},
		{"5 / 2 + 5 % 2;", 3},
		{"1 == 1;", 1},
		{"1 != 1;", 0},
		{"1 < 1;", 0},
		{"1 <= 1;", 1},
		{"1 < 1;", 0},
		{"1 <= 1;", 1},
		{"let hoge = 2; hoge;", 2},
		{"if 1 == 1 { 2; };", 2},
		{"if 1 != 1 { 2; };", 0},
		{"if 1 == 1 { 2; } else { 3; }", 2},
		{"if 1 != 1 { 2; } else { 3; }", 3},
		{"let i = 0; let sum = 0; while i < 10 { let sum = sum + i; let i = i + 1; }; sum;", 45},
	}

	for _, test := range tests {
		ctx := Context{}
		actual := Exec(test.input, ctx)
		if actual != test.expected {
			t.Errorf("[FAILED] `%s` -> %d\n", test.input, actual)
		}
	}

}
