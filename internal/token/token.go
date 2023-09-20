package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal   Type = "ILLEGAL"
	EOF       Type = "EOF"
	Ident     Type = "IDENT"
	Int       Type = "INT"
	Assign    Type = "="
	Plus      Type = "+"
	Comma     Type = ","
	Semicolon Type = ";"
	LParen    Type = "("
	RParen    Type = ")"
	LBrace    Type = "{"
	RBrace    Type = "}"
	Function  Type = "FUNCTION"
	Let       Type = "LET"
)
