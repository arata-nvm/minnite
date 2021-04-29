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

func (e *Expression) Eval(ctx Context) int {
	return e.Expression.Eval(ctx)
}

func (e *AdditionExpression) Eval(ctx Context) int {
	lhs := e.Lhs.Eval(ctx)

	for _, rhs := range e.Rhs {
		switch *rhs.Op {
		case "+":
			lhs += rhs.Mul.Eval(ctx)
		case "-":
			lhs -= rhs.Mul.Eval(ctx)
		}
	}

	return lhs
}

func (e *MultiplicationExpression) Eval(ctx Context) int {
	lhs := e.Lhs.Eval(ctx)

	for _, rhs := range e.Rhs {
		switch *rhs.Op {
		case "*":
			lhs *= rhs.Term.Eval(ctx)
		case "/":
			lhs /= rhs.Term.Eval(ctx)
		case "%":
			lhs %= rhs.Term.Eval(ctx)
		}
	}

	return lhs
}

func (t *TermExpression) Eval(ctx Context) int {
	switch {
	case t.Variable != nil:
		return ctx[*t.Variable]
	case t.Number != nil:
		return *t.Number
	}

	return 0
}
