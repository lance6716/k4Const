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
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		checkWriteToConst(pass, node)
	})
	return nil, err
}

func checkWriteToConst(pass *analysis.Pass, node ast.Node) {
	switch v := node.(type) {
	case *ast.AssignStmt:
		for _, e := range v.Lhs {
			checkWriteToOneConst(pass, e)
		}
	case *ast.IncDecStmt:
		checkWriteToOneConst(pass, v.X)
	}
}

func checkWriteToOneConst(pass *analysis.Pass, e ast.Expr) {
	e = unwrap(e, derefAndIndex)
	switch v := e.(type) {
	case *ast.Ident:
		if name, ok := isConstIdent(v); ok {
			pass.Reportf(e.Pos(), "write to const variable '%s'", name)
		}
	case *ast.SelectorExpr:
		if _, ok := isConstIdent(v.Sel); ok {
			var b strings.Builder
			err := printer.Fprint(&b, pass.Fset, v)
			if err != nil {
				pass.Reportf(e.Pos(), "failed to get SelectorExpr name: %v", err)
				return
			}
			pass.Reportf(e.Pos(), "write to const variable '%s'", b.String())
		}
		if name, ok := isConstIdent(unwrap(v.X, deref)); ok {
			pass.Reportf(e.Pos(), "write to const variable '%s'", name)
		}
	}
}
