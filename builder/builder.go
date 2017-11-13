package builder

import (
	"strings"

	"github.com/TomiLabo/tmngparser/ast"
)

func BuildFromTree(tree []ast.Tml) string {
	var result []string
	for _, t := range tree {
		result = append(result, t.Statement)
	}
	return strings.Join(result, "")
}
