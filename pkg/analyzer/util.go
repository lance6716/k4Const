package analyzer

import (
	"go/ast"
	"unicode"
)

func isConstIdent(node ast.Node) (bool, string) {
	ident, ok := node.(*ast.Ident)
	if !ok {
		return false, ""
	}

	num := 0
FORLOOP:
	for _, r := range ident.Name {
		switch num {
		case 0:
			if r != 'k' {
				return false, ""
			}
		case 1:
			if unicode.IsLower(r) {
				return false, ""
			}
		default:
			break FORLOOP
		}
		num++
	}
	return num > 1, ident.Name
}

func deref(node ast.Node) ast.Node {
	ret := node

	derefOnce := func(n ast.Node) (ast.Node, bool) {
		switch v := n.(type) {
		case *ast.StarExpr:
			return v.X, true
		case *ast.ParenExpr:
			return v.X, true
		default:
			return nil, false
		}
	}

	tmp, ok := derefOnce(ret)
	for ok {
		ret = tmp
		tmp, ok = derefOnce(ret)
	}
	return ret
}
