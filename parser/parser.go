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
