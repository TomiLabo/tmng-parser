package ast

import (
	"github.com/TomiLabo/tmngparser/token"
)

// Tml ast
type Tml struct {
	Token     token.Code `json:"token"`
	Line      int        `json:"line"`
	Start     int        `json:"start"`
	End       int        `json:"end"`
	Statement string     `json:"statement"`
}

// New generate struct
func New(line int, token token.Code, statement string, start int, end int) Tml {
	return Tml{
		Line:      line + 1,
		Token:     token,
		Start:     start,
		End:       end,
		Statement: statement,
	}
}
