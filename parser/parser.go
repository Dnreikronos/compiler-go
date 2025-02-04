package parser

import "github.com/Dnreikronos/compiler-go/lexer"

type Expr interface {
	exprNode()
}

type NumberExpr struct {
	Value string
}

func (n NumberExpr) exprNode() {}

type BinaryExpr struct {
	Left  Expr
	Op    lexer.TokenType
	Right Expr
}

func (b BinaryExpr) exprNode() {}

type Parser struct {
	tokens []lexer.Token
	pos    int
}
func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}
func (p *Parser) Parse() Expr {
	return p.parseExpr()
}
func (p *Parser) parseExpr() Expr {
	return p.parseTerm()
}
