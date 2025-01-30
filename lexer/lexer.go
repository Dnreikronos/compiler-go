package lexer

type TokenType string

const (
	Number   TokenType = "Number"
	Plus     TokenType = "+"
	Minus    TokenType = "-"
	Multiply TokenType = "*"
	Divide   TokenType = "/"
	EOF      TokenType = "EOF"
)
type Token struct {
	Type  TokenType
	Value string
}
