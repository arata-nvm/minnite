package main

import "fmt"

type Context map[string]int

func (p *Program) Eval(ctx Context) {
	for _, stmt := range p.Statements {
		stmt.Eval(ctx)
	}
}

func (s *Statement) Eval(ctx Context) {
	switch {
	case s.Let != nil:
		s.Let.Eval(ctx)
	case s.Print != nil:
		s.Print.Eval(ctx)
	}
}

func (s *LetStatement) Eval(ctx Context) {
	ctx[s.Variable] = s.Value.Eval(ctx)
}

func (s *PrintStatement) Eval(ctx Context) {
	fmt.Printf("%d\n", s.Value.Eval(ctx))

}

func (t *Term) Eval(ctx Context) int {
	switch {
	case t.Variable != nil:
		return ctx[*t.Variable]
	case t.Number != nil:
		return *t.Number
	}

	return 0
}
