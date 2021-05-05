package minnite

import "fmt"

type Context struct {
	variables map[string]Value
	parent    *Context
}

func NewContext() *Context {
	ctx := Context{variables: map[string]Value{}}
	return &ctx
}

func (ctx *Context) Clone() *Context {
	newCtx := NewContext()
	newCtx.parent = ctx
	return newCtx
}

func (ctx *Context) AddVariable(name string, value Value) {
	ctx.variables[name] = value
}

func (ctx *Context) SetVariable(name string, value Value) {
	currentCtx := ctx
	for {
		if _, ok := currentCtx.variables[name]; ok {
			currentCtx.variables[name] = value
			return
		}

		if currentCtx.parent == nil {
			panic(fmt.Sprintf("変数'%s'は存在しません", name))
		}

		currentCtx = currentCtx.parent
	}
}

func (ctx *Context) FindVariable(name string) Value {
	if value, ok := ctx.variables[name]; ok {
		return value
	}

	if ctx.parent == nil {
		panic(fmt.Sprintf("変数'%s'は存在しません", name))
	}

	return ctx.parent.FindVariable(name)
}
