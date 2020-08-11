package analys

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Analys(filepath string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filepath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed parse file with path %s: %w", filepath, err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			fmt.Printf("return statement found on line %d:\n\t", fileSet.Position(ret.Pos()).Line)
			return true
		}
		return true
	})

	return nil
}
