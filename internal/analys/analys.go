package analys

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type AnalysResult struct {
	DeclCount    int
	CallCount    int
	AssignCount  int
	ImportsCount int
}

func Analys(filepath string) (*AnalysResult, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	importCount := len(node.Imports)

	var assignCount int
	ast.Inspect(node, func(n ast.Node) bool {
		if _, ok := n.(*ast.AssignStmt); ok {
			assignCount++
		}
		return true
	})
	var callCount int
	ast.Inspect(node, func(n ast.Node) bool {
		if _, ok := n.(*ast.CallExpr); ok {
			callCount++
		}

		return true
	})
	var declarationCount int
	ast.Inspect(node, func(n ast.Node) bool {
		if gen, ok := n.(*ast.GenDecl); ok && gen.Tok != token.IMPORT {
			declarationCount++
		}
		return true
	})

	result := &AnalysResult{
		DeclCount:    declarationCount,
		ImportsCount: importCount,
		CallCount:    callCount,
		AssignCount:  assignCount,
	}

	return result, nil
}
