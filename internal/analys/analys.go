package analys

import (
	"fmt"
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

func Analys(filepath string) (AnalysResult, error) {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return AnalysResult{}, fmt.Errorf("failed parse file with path %s: %w", filepath, err)
	}

	result := AnalysResult{}

	var declCount int
	var assignCount int
	var callCount int

	ast.Inspect(node, func(n ast.Node) bool {
		if _, ok := n.(*ast.CallExpr); ok {
			callCount++
		}

		if _, ok := n.(*ast.DeclStmt); ok {
			declCount++
		}

		if _, ok := n.(*ast.AssignStmt); ok {
			assignCount++
		}

		return true
	})

	result.DeclCount = declCount
	result.CallCount = callCount
	result.AssignCount = assignCount
	result.ImportsCount = len(node.Imports)

	return result, nil
}
