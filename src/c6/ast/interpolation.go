package ast

type Interpolation struct {
	Expression Expression
	StartToken *Token
	EndToken   *Token
}

func (self Interpolation) CanBeExpression() {}
func (self Interpolation) CanBeNode()       {}

func NewInterpolation(expr Expression, startToken *Token, endToken *Token) *Interpolation {
	return &Interpolation{expr, startToken, endToken}
}