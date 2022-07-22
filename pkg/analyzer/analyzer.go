package analyzer

import (
	"go/ast"
	"go/printer"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "k4const",
	Doc:      "Checks that k-started parameters should not be written.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	var err error
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.AssignStmt)(nil),
		(*ast.IncDecStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		switch v := node.(type) {
		case *ast.AssignStmt:
			for _, e := range v.Lhs {
				err = checkWriteToConst(pass, e)
			}
		case *ast.IncDecStmt:
			err = checkWriteToConst(pass, v.X)
		}
	})
	return nil, err
}

func checkWriteToConst(pass *analysis.Pass, e ast.Expr) error {
	switch v := e.(type) {
	case *ast.IndexExpr:
		if ok, name := isConstIdent(deref(v.X)); ok {
			pass.Reportf(e.Pos(), "write to const variable '%s'", name)
		}
	case *ast.SelectorExpr:
		if ok, _ := isConstIdent(v.Sel); ok {
			var b strings.Builder
			err := printer.Fprint(&b, pass.Fset, v)
			if err != nil {
				return err
			}
			pass.Reportf(e.Pos(), "write to const variable '%s'", b.String())
			return nil
		}
		if ok, name := isConstIdent(deref(v.X)); ok {
			pass.Reportf(e.Pos(), "write to const variable '%s'", name)
		}
	case *ast.StarExpr:
		if ok, name := isConstIdent(deref(v.X)); ok {
			pass.Reportf(e.Pos(), "write to const variable '%s'", name)
		}
	}
	return nil
}
