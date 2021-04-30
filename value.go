package main

import "fmt"

type Value interface {
	Type() ValueType
	String() string
}

type ValueType int

const (
	VOID    ValueType = iota
	INTEGER ValueType = iota
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
