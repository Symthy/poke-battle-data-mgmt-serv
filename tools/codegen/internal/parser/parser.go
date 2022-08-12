package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(filename string) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(0))

	if err != nil {
		return err
	}

	// for _, d := range f.Decls {
	// 	ast.Print(fset, d)
	// 	fmt.Println()
	// }

	fieldNames := []string{}
	ast.Inspect(f, func(node ast.Node) bool {
		t, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := t.Type.(*ast.StructType)
		if !ok {
			return true
		}
		for _, field := range st.Fields.List {
			ast.Print(fset, field)
			fmt.Println()

			for _, nameNode := range field.Names {
				fieldNames = append(fieldNames, nameNode.Name)
			}
		}
		return true
	})
	fmt.Print(fieldNames)
	return nil
}
