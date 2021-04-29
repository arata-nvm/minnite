package main

type Program struct {
	Statements []*Statement `@@*`
}

type Statement struct {
	Let   *LetStatement   `( @@ `
	Print *PrintStatement `| @@ ) ";"`
}

type LetStatement struct {
	Variable string `"let" @Ident`
	Value    *Term  `"=" @@`
}

type PrintStatement struct {
	Value *Term `"print" @@`
}

type Term struct {
	Variable *string `@Ident`
	Number   *int    `| @Number`
}
