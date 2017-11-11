package ast

import (
	"github.com/tomilabo/tmngparser/category"
)

// Tml ast
type Tml struct {
	Category  category.Code `json:"category"`
	Line      int           `json:"line"`
	Start     int           `json:"start"`
	End       int           `json:"end"`
	Statement string        `json:"statement"`
}

// New generate struct
func New(line int, category category.Code, statement string, start int, end int) Tml {
	return Tml{
		Line:      line + 1,
		Category:  category,
		Start:     start,
		End:       end,
		Statement: statement,
	}
}
