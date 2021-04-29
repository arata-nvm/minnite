package main

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let        *LetStatement        `( @@ `
	Print      *PrintStatement      `| @@ `
	Expression *ExpressionStatement `| @@ ) ";"`
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
	Variable   *string     `@Ident`
	Number     *int        `| @Number`
	Expression *Expression `| ( "(" @@ ")" )`
}
