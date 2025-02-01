package lexer

import "strings"
import (
	"strings"
	"unicode"
)

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

func Lex(input string) []Token {
	var tokens []Token
	r := strings.NewReader(input)

	for {
		ch, _, err := r.ReadRune()
		if err != nil { // End of input
			tokens = append(tokens, Token{EOF, ""})
			break
		}

		switch {
		case unicode.IsSpace(ch):
			continue // Skip whitespace
		case unicode.IsDigit(ch):
			r.UnreadRune()
			num := readNumber(r)
			tokens = append(tokens, Token{Number, num})
		case ch == '+':
			tokens = append(tokens, Token{Plus, "+"})
		case ch == '-':
			tokens = append(tokens, Token{Minus, "-"})
		case ch == '*':
			tokens = append(tokens, Token{Multiply, "*"})
		case ch == '/':
			tokens = append(tokens, Token{Divide, "/"})
		}
	}

	return tokens
}
