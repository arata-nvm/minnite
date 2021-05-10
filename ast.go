package main

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let        *LetStatement        `( @@ `
	Print      *PrintStatement      `| @@ `
	If         *IfStatement         `| @@ `
	While      *WhileStatement      `| @@ `
	Return     *ReturnStatement     `| @@ `
	Assign     *AssignStatement     `| @@ `
	Expression *ExpressionStatement `| @@ ) [ ";" ]`
}

type BlockStatement struct {
	Body []*Statement `"{" @@* "}"`
}

type LetStatement struct {
	Variable string      `"let" @Ident`
	Value    *Expression `"=" @@`
}

type PrintStatement struct {
	Value *Expression `"print" @@`
}

type ExpressionStatement struct {
	Expression *Expression `@@`
}

type IfStatement struct {
	Cond *Expression     `"if" @@ `
	Then *BlockStatement `@@`
	Else *BlockStatement `[ "else" @@ ]`
}

type WhileStatement struct {
	Cond *Expression     `"while" @@ `
	Body *BlockStatement `@@`
}

type ReturnStatement struct {
	Value *Expression `"return" @@`
}

type AssignStatement struct {
	Variable string      `@Ident`
	Value    *Expression `"=" @@`
}

type Expression struct {
	Expression *ComparisonExpression `@@`
}

type ComparisonExpression struct {
	Lhs *AdditionExpression `@@`
	Op  *string             `[ @Operator`
	Rhs *AdditionExpression `  @@ ]`
}

type AdditionExpression struct {
	Lhs *MultiplicationExpression `@@`
	Rhs []*OpAdditionExpression   `@@*`
}

type OpAdditionExpression struct {
	Op  *string                   `@("+" | "-")`
	Mul *MultiplicationExpression `  @@`
}

type MultiplicationExpression struct {
	Lhs *TermExpression               `@@`
	Rhs []*OpMultiplicationExpression `@@*`
}

type OpMultiplicationExpression struct {
	Op   *string         `@("*" | "/" | "%")`
	Term *TermExpression `@@`
}

type TermExpression struct {
	Expression *Expression         `( "(" @@ ")" )`
	List       *ListExpression     `| @@`
	Index      *IndexExpression    `| @@`
	Function   *FunctionExpression `| @@`
	Call       *CallExpression     `| @@`
	Variable   *string             `| @Ident`
	Number     *int                `| @Number`
	String     *string             `| @String`
}

type FunctionExpression struct {
	Params []string        `"func" "(" ( @Ident ( "," @Ident )* )? ")"`
	Body   *BlockStatement `@@`
}

type CallExpression struct {
	Name *string       `@Ident`
	Args []*Expression `"(" ( @@ ( "," @@ )* )? ")"`
}

type ListExpression struct {
	Items []*Expression `"[" ( @@ ( "," @@ )* )? "]"`
}

type IndexExpression struct {
	Variable *string     `@Ident`
	Index    *Expression `"[" @@ "]"`
}
