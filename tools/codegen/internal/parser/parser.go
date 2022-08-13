package parser

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
)

func Parse(filename string) ([]*internal.FieldInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(0))

	if err != nil {
		return nil, err
	}
	// for _, d := range f.Decls {
	// 	ast.Print(fset, d)
	// 	fmt.Println()
	// }

	fields := []*internal.FieldInfo{}

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
			// ast.Print(fset, field)
			// fmt.Println()

			var typeInfo *internal.TypeInfo = nil
			typeNode := field.Type
			if e, ok := typeNode.(*ast.Ident); ok {
				// 標準型or内部クラス
				typeInfo = internal.NewStandardTypeInfo(e.Name)
			} else if e, ok := typeNode.(*ast.SelectorExpr); ok {
				// 外部クラス
				pkgName := ""
				if ec, ok := e.X.(*ast.Ident); ok {
					pkgName = ec.Name
				}
				sel := e.Sel
				typeInfo = internal.NewTypeInfo(pkgName, sel.Name)
			} else if e, ok := typeNode.(*ast.ArrayType); ok {
				// 配列
				elt := e.Elt
				if e, ok := elt.(*ast.Ident); ok {
					typeInfo = internal.NewArrayTypeInfo("", e.Name)
				} else if e, ok := elt.(*ast.StarExpr); ok {
					x := e.X
					if e, ok := x.(*ast.Ident); ok {
						typeInfo = internal.NewPointerArrayTypeInfo("", e.Name)
					}
					if e, ok := x.(*ast.SelectorExpr); ok {
						pkgName := ""
						if ec, ok := e.X.(*ast.Ident); ok {
							pkgName = ec.Name
						}
						sel := e.Sel
						typeInfo = internal.NewPointerArrayTypeInfo(pkgName, sel.Name)
					}
				} else if e, ok := elt.(*ast.SelectorExpr); ok {
					pkgName := ""
					if ec, ok := e.X.(*ast.Ident); ok {
						pkgName = ec.Name
					}
					sel := e.Sel
					typeInfo = internal.NewArrayTypeInfo(pkgName, sel.Name)
				}
			} else if e, ok := typeNode.(*ast.StarExpr); ok {
				// ポインタ
				if x, ok := e.X.(*ast.SelectorExpr); ok {
					pkgName := ""
					if ec, ok := x.X.(*ast.Ident); ok {
						pkgName = ec.Name
					}
					sel := x.Sel
					typeInfo = internal.NewPointerTypeInfo(pkgName, sel.Name)
				} else if x, ok := e.X.(*ast.Ident); ok {
					typeInfo = internal.NewPointerTypeInfo("", x.Name)
				}
			}

			if typeInfo == nil {
				continue
			}
			if len(field.Names) == 0 {
				fields = append(fields, internal.NewEmbeddedFieldInfo(typeInfo))
				continue
			}
			for _, nameNode := range field.Names {
				fields = append(fields, internal.NewFieldInfo(nameNode.Name, typeInfo))
			}

		}
		return true
	})

	return fields, nil
}
