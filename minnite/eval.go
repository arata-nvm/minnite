package minnite

import "fmt"

func (p *Program) Eval(ctx *Context) Value {
	result := NewVoid()

	for _, stmt := range p.Statements {
		result = stmt.Eval(ctx)
	}

	return result
}

func (s *Statement) Eval(ctx *Context) Value {
	switch {
	case s.Let != nil:
		return s.Let.Eval(ctx)
	case s.Print != nil:
		return s.Print.Eval(ctx)
	}

	panic("unreachable")
}

func (s *LetStatement) Eval(ctx *Context) Value {
	ctx.AddVariable(s.Variable, s.Value.Eval(ctx))
	return NewVoid()
}

func (s *PrintStatement) Eval(ctx *Context) Value {
	fmt.Println(s.Value.Eval(ctx))
	return NewVoid()
}

func (e *Expression) Eval(ctx *Context) Value {
	return e.Expression.Eval(ctx)
}

func (e *ComparisonExpression) Eval(ctx *Context) Value {
	if e.Op == nil {
		return e.Lhs.Eval(ctx)
	}

	lhs := e.Lhs.Eval(ctx).(Integer)
	rhs := e.Rhs.Eval(ctx).(Integer)
	var result bool
	switch *e.Op {
	case "==":
		result = lhs == rhs
	}

	return NewBoolean(result)
}

func (e *AdditionExpression) Eval(ctx *Context) Value {
	if len(e.Rhs) == 0 {
		return e.Lhs.Eval(ctx)
	}

	lhs := e.Lhs.Eval(ctx).(Integer)
	for _, rhs := range e.Rhs {
		op := *rhs.Op
		rhs := rhs.Mul.Eval(ctx).(Integer)
		switch op {
		case "+":
			lhs += rhs
		}
	}

	return lhs
}

func (e *MultiplicationExpression) Eval(ctx *Context) Value {
	if len(e.Rhs) == 0 {
		return e.Lhs.Eval(ctx)
	}

	lhs := e.Lhs.Eval(ctx).(Integer)
	for _, rhs := range e.Rhs {
		op := *rhs.Op
		rhs := rhs.Term.Eval(ctx).(Integer)
		switch op {
		case "*":
			lhs *= rhs
		}
	}

	return lhs
}

func (t *TermExpression) Eval(ctx *Context) Value {
	switch {
	case t.Variable != nil:
		return ctx.FindVariable(*t.Variable)
	case t.Number != nil:
		return NewInteger(*t.Number)
	}

	panic("unreachable")
}
