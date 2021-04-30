package main

import "fmt"

type Context map[string]Value

func (p *Program) Eval(ctx Context) Value {
	result := NewInteger(0)

	for _, stmt := range p.Statements {
		result = stmt.Eval(ctx)
	}

	return result
}

func (s *Statement) Eval(ctx Context) Value {
	switch {
	case s.Let != nil:
		s.Let.Eval(ctx)
	case s.Print != nil:
		s.Print.Eval(ctx)
	case s.If != nil:
		return s.If.Eval(ctx)
	case s.While != nil:
		s.While.Eval(ctx)
	case s.Expression != nil:
		return s.Expression.Eval(ctx)
	}

	return NewInteger(0)
}

func (s *BlockStatement) Eval(ctx Context) Value {
	result := NewInteger(0)

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

func (s *IfStatement) Eval(ctx Context) Value {
	cond := s.Cond.Eval(ctx).(Integer)
	result := NewInteger(0)

	if cond != 0 {
		result = s.Then.Eval(ctx)
	} else if s.Else != nil {
		result = s.Else.Eval(ctx)
	}

	return result
}

func (s *WhileStatement) Eval(ctx Context) {
	for {
		cond := s.Cond.Eval(ctx).(Integer)
		if cond == 0 {
			break
		}

		s.Body.Eval(ctx)
	}

}

func (s *ExpressionStatement) Eval(ctx Context) Value {
	return s.Expression.Eval(ctx)
}

func (e *Expression) Eval(ctx Context) Value {
	return e.Expression.Eval(ctx)
}

func (e *ComparisonExpression) Eval(ctx Context) Value {
	lhs := e.Lhs.Eval(ctx).(Integer)
	if e.Op == nil {
		return lhs
	}

	rhs := e.Rhs.Eval(ctx).(Integer)

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
		return NewInteger(1)
	} else {
		return NewInteger(0)
	}
}

func (e *AdditionExpression) Eval(ctx Context) Value {
	lhs := e.Lhs.Eval(ctx).(Integer)

	for _, rhs := range e.Rhs {
		op := *rhs.Op
		rhs := rhs.Mul.Eval(ctx).(Integer)
		switch op {
		case "+":
			lhs += rhs
		case "-":
			lhs -= rhs
		}
	}

	return lhs
}

func (e *MultiplicationExpression) Eval(ctx Context) Value {
	lhs := e.Lhs.Eval(ctx).(Integer)

	for _, rhs := range e.Rhs {
		op := *rhs.Op
		rhs := rhs.Term.Eval(ctx).(Integer)
		switch op {
		case "*":
			lhs *= rhs
		case "/":
			lhs /= rhs
		case "%":
			lhs %= rhs
		}
	}

	return lhs
}

func (t *TermExpression) Eval(ctx Context) Value {
	switch {
	case t.Variable != nil:
		return ctx[*t.Variable]
	case t.Number != nil:
		return NewInteger(*t.Number)
	case t.Expression != nil:
		return t.Expression.Eval(ctx)
	}

	return NewInteger(0)
}
