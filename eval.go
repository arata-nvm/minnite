package main

import "fmt"

type Context map[string]int

func (p *Program) Eval(ctx Context) int {
	result := 0

	for _, stmt := range p.Statements {
		result = stmt.Eval(ctx)
	}

	return result
}

func (s *Statement) Eval(ctx Context) int {
	switch {
	case s.Let != nil:
		s.Let.Eval(ctx)
	case s.Print != nil:
		s.Print.Eval(ctx)
	case s.If != nil:
		return s.If.Eval(ctx)
	case s.Expression != nil:
		return s.Expression.Eval(ctx)
	}

	return 0
}

func (s *BlockStatement) Eval(ctx Context) int {
	result := 0

	for _, stmt := range s.Body {
		result = stmt.Eval(ctx)
	}

	return result
}

func (s *LetStatement) Eval(ctx Context) {
	ctx[s.Variable] = s.Value.Eval(ctx)
}

func (s *PrintStatement) Eval(ctx Context) {
	fmt.Printf("%d\n", s.Value.Eval(ctx))
}

func (s *IfStatement) Eval(ctx Context) int {
	cond := s.Cond.Eval(ctx)
	result := 0

	if cond != 0 {
		result = s.Then.Eval(ctx)
	} else if s.Else != nil {
		result = s.Else.Eval(ctx)
	}

	return result
}

func (s *ExpressionStatement) Eval(ctx Context) int {
	return s.Expression.Eval(ctx)
}

func (e *Expression) Eval(ctx Context) int {
	return e.Expression.Eval(ctx)
}

func (e *ComparisonExpression) Eval(ctx Context) int {
	lhs := e.Lhs.Eval(ctx)
	if e.Op == nil {
		return lhs
	}

	rhs := e.Rhs.Eval(ctx)

	var result bool
	switch *e.Op {
	case "==":
		result = lhs == rhs
	case "!=":
		result = lhs != rhs
	case "<":
		result = lhs < rhs
	case "<=":
		result = lhs <= rhs
	case ">":
		result = lhs > rhs
	case ">=":
		result = lhs >= rhs
	}

	if result {
		return 1
	} else {
		return 0
	}
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
	case t.Expression != nil:
		return t.Expression.Eval(ctx)
	}

	return 0
}
