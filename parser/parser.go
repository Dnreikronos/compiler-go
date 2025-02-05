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
func (p *Parser) parseTerm() Expr {
	left := p.parseFactor()

	for {
		if p.match(lexer.Plus, lexer.Minus) {
			op := p.previous().Type
			right := p.parseFactor()
			left = BinaryExpr{Left: left, Op: op, Right: right}
		} else {
			break
		}
	}
	return left
}

func (p *Parser) parseFactor() Expr {
	left := p.parsePrimary()

	for {
		if p.match(lexer.Multiply, lexer.Divide) {
			op := p.previous().Type
			right := p.parsePrimary()
			left = BinaryExpr{Left: left, Op: op, Right: right}
		} else {
			break
		}
	}
	return left
}
func (p *Parser) parsePrimary() Expr {
	if p.match(lexer.Number) {
		return NumberExpr{Value: p.previous().Value}
	}
	panic("Unexpected token")
}

func (p *Parser) match(types ...lexer.TokenType) bool {
	for _, t := range types {
		if p.peek().Type == t {
			p.pos++
			return true
		}
	}
	return false
}
func (p *Parser) peek() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) previous() lexer.Token {
	return p.tokens[p.pos-1]
}
