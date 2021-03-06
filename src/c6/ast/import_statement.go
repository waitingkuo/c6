package ast

type ImportStatement struct {
	Url       interface{} // if it's wrapped with url(...) or "string"
	MediaList []string
}

func (self ImportStatement) CanBeStatement() {}

func (self ImportStatement) String() string {
	return ""
}

// for Url()
type Url string

type RelativeUrl string
