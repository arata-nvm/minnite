package main

import "fmt"

type Value interface {
	Type() ValueType
	String() string
}

type ValueType int

const (
	INTEGER ValueType = iota
)

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
