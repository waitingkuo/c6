package ast

/**
An property may contains interpolation
*/
type Property struct {
	Name *PropertyName
	/**
	property value can be something like:

		`padding: 3px 3px;`
	*/
	Values []Expression
}

/**
Property is one of the declaration
*/
func (self Property) CanBeDeclaration() {}

func (self Property) AppendValue(value Expression) {
	self.Values = append(self.Values, value)
}

type PropertyName struct {
	String string
	// If there is an interpolation in the property name
	Interpolation bool
	Token         *Token
}

func NewPropertyName(tok *Token) *PropertyName {
	return &PropertyName{tok.Str, tok.ContainsInterpolation, tok}
}

func NewProperty(nameTok *Token) *Property {
	return &Property{NewPropertyName(nameTok), []Expression{}}
}
