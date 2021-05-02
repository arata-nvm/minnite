package main

import "fmt"

type Value interface {
	Type() ValueType
	String() string
}

type ValueType int

const (
	VOID     ValueType = iota
	BOOLEAN  ValueType = iota
	INTEGER  ValueType = iota
	FUNCTION ValueType = iota
)

type Void struct{}

func NewVoid() Value {
	return &Void{}
}

func (v *Void) Type() ValueType {
	return VOID
}

func (v *Void) String() string {
	return "void"
}

type Boolean bool

func NewBoolean(b bool) Value {
	return Boolean(b)
}

func (b Boolean) Type() ValueType {
	return BOOLEAN
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b)
}

type Integer int

func NewInteger(i int) Value {
	return Integer(i)
}

func (i Integer) Type() ValueType {
	return INTEGER
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type Function struct {
	Params []string
	Body   *BlockStatement
}

func NewFunction(params []string, body *BlockStatement) Value {
	return &Function{
		Params: params,
		Body:   body,
	}
}

func (f *Function) Type() ValueType {
	return FUNCTION
}

func (f *Function) String() string {
	return "func"
}
