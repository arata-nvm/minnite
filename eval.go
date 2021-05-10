package main

import (
	"bytes"
	"fmt"
	"strings"
)

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
	case s.If != nil:
		return s.If.Eval(ctx)
	case s.While != nil:
		return s.While.Eval(ctx)
	case s.Return != nil:
		return s.Return.Eval(ctx)
	case s.Assign != nil:
		return s.Assign.Eval(ctx)
	case s.Expression != nil:
		return s.Expression.Eval(ctx)
	}

	panic("unreachable")
}

func (s *BlockStatement) Eval(ctx *Context) Value {
	result := NewVoid()

	newCtx := ctx.Clone()
	for _, stmt := range s.Body {
		result = stmt.Eval(newCtx)

		if stmt.Return != nil {
			break
		}
	}

	return result
}

func (s *LetStatement) Eval(ctx *Context) Value {
	ctx.AddVariable(s.Variable, s.Value.Eval(ctx))
	return NewVoid()
}

func (s *PrintStatement) Eval(ctx *Context) Value {
	fmt.Printf("%v\n", s.Value.Eval(ctx))
	return NewVoid()
}

func (s *IfStatement) Eval(ctx *Context) Value {
	cond := s.Cond.Eval(ctx).(Boolean)
	result := NewVoid()

	if cond {
		result = s.Then.Eval(ctx)
	} else if s.Else != nil {
		result = s.Else.Eval(ctx)
	}

	return result
}

func (s *WhileStatement) Eval(ctx *Context) Value {
	for {
		cond := s.Cond.Eval(ctx).(Boolean)
		if !cond {
			break
		}

		s.Body.Eval(ctx)
	}

	return NewVoid()
}

func (s *ReturnStatement) Eval(ctx *Context) Value {
	return s.Value.Eval(ctx)
}

func (s *AssignStatement) Eval(ctx *Context) Value {
	ctx.SetVariable(s.Variable, s.Value.Eval(ctx))
	return NewVoid()
}

func (s *ExpressionStatement) Eval(ctx *Context) Value {
	return s.Expression.Eval(ctx)
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
		case "-":
			lhs -= rhs
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
		case "/":
			lhs /= rhs
		case "%":
			lhs %= rhs
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
	case t.String != nil:
		return EvalString(*t.String, ctx)
	case t.Expression != nil:
		return t.Expression.Eval(ctx)
	case t.Function != nil:
		return t.Function.Eval(ctx)
	case t.Call != nil:
		return t.Call.Eval(ctx)
	case t.List != nil:
		return t.List.Eval(ctx)
	case t.Index != nil:
		return t.Index.Eval(ctx)
	}

	panic("unreachable")
}

func EvalString(s string, ctx *Context) Value {
	// 前後の"を除去する
	content := strings.TrimSuffix(strings.TrimPrefix(s, "\""), "\"")

	if !strings.Contains(content, "$") {
		return NewString(content)
	}

	var buf bytes.Buffer
	for i := 0; i < len(content); i++ {
		c := content[i]
		switch c {
		case '$':
			nextSpace := strings.IndexByte(content[i:], ' ')
			if nextSpace == -1 {
				nextSpace = len(content)
			} else {
				nextSpace += i
			}

			varName := content[i+1 : nextSpace]
			buf.WriteString(ctx.FindVariable(varName).String())

			i = nextSpace - 1
		default:
			buf.WriteByte(c)
		}
	}

	return NewString(buf.String())
}

func (f *FunctionExpression) Eval(ctx *Context) Value {
	return NewFunction(f.Params, f.Body)
}

func (c *CallExpression) Eval(ctx *Context) Value {
	f := ctx.FindVariable(*c.Name).(*Function)

	if len(f.Params) != len(c.Args) {
		panic(fmt.Sprintf("%d個の引数に対し、%d個の値が与えられています", len(f.Params), len(c.Args)))
	}

	newCtx := ctx.Clone()
	for i, param := range f.Params {
		newCtx.AddVariable(param, c.Args[i].Eval(newCtx))
	}

	return f.Body.Eval(newCtx)
}

func (l *ListExpression) Eval(ctx *Context) Value {
	items := []Value{}

	for _, item := range l.Items {
		items = append(items, item.Eval(ctx))
	}

	return NewList(items)
}

func (i *IndexExpression) Eval(ctx *Context) Value {
	value := ctx.FindVariable(*i.Variable).(*List)
	index := i.Index.Eval(ctx).(Integer)

	return value.Items[index]
}
