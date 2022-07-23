package analyzer

import (
	"go/ast"
	"unicode"
)

func isConstIdent(node ast.Expr) (string, bool) {
	ident, ok := node.(*ast.Ident)
	if !ok {
		return "", false
	}

	num := 0
FORLOOP:
	for _, r := range ident.Name {
		switch num {
		case 0:
			if r != 'k' {
				return "", false
			}
		case 1:
			if unicode.IsLower(r) {
				return "", false
			}
		default:
			break FORLOOP
		}
		num++
	}
	return ident.Name, num > 1
}

func unwrap(node ast.Expr, f func(n ast.Expr) (ast.Expr, bool)) ast.Expr {
	var (
		ret = node
		ok  bool
	)

	ret, ok = f(ret)
	for ok {
		ret, ok = f(ret)
	}
	return ret
}

func derefAndIndex(n ast.Expr) (ast.Expr, bool) {
	switch v := n.(type) {
	case *ast.StarExpr:
		return v.X, true
	case *ast.ParenExpr:
		return v.X, true
	case *ast.IndexExpr:
		return v.X, true
	default:
		return v, false
	}
}

func deref(n ast.Expr) (ast.Expr, bool) {
	switch v := n.(type) {
	case *ast.StarExpr:
		return v.X, true
	case *ast.ParenExpr:
		return v.X, true
	default:
		return v, false
	}
}
