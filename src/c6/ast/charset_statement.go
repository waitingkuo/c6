package ast

type CharsetStatement struct {
	Charset string
	Token   *Token
}

func (self CharsetStatement) CanBeStatement() {}

func NewCharsetStatement(token *Token) *CharsetStatement {
	return &CharsetStatement{token.Str, token}
}
