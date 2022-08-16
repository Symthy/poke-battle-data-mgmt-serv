package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/Symthy/PokeRest/pokeRest/common/lists"
	"github.com/Symthy/PokeRest/tools/codegen/internal"
)

func ParseFiles(targetPkgPath string, fileToStruct map[string]string) ([]*internal.StructInfo, error) {
	isErr := false
	structs := []*internal.StructInfo{}
	for file, structName := range fileToStruct {
		path := filepath.Join(targetPkgPath, file)

		var sts []*internal.StructInfo
		var err error
		if structName == "" {
			sts, err = ParseSingleFile(path)
		} else {
			sts, err = ParseSingleFile(path, structName)
		}

		if err != nil {
			isErr = true
			fmt.Printf("[Error] %s %s: %v\n", file, structName, err)
			continue
		}
		structs = append(structs, sts...)
		fmt.Printf("[Info]  %s %s: success to parse\n", file, structName)
	}
	if isErr {
		return structs, fmt.Errorf("Failed to Parse.\n")
	}
	return structs, nil
}

func ParseSingleStruct(gofilePath string, structName string) (*internal.StructInfo, error) {
	structs, err := ParseSingleFile(gofilePath, structName)
	if err != nil {
		return nil, err
	}
	return structs[0], nil
}

func ParseSingleFile(gofilePath string, filterTypeNames ...string) ([]*internal.StructInfo, error) {
	fileNode, err := parseGoFile(gofilePath)
	if err != nil {
		return nil, err
	}
	// for _, d := range f.Decls {
	// 	ast.Print(fset, d)
	// 	fmt.Println()
	// }

	structs := parseStruct(fileNode, filterTypeNames...)
	if len(structs) == 0 {
		return nil, fmt.Errorf("Specied Type noting.")
	}

	return structs, nil
}

func parseGoFile(gofilePath string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, gofilePath, nil, parser.Mode(0))
}

func parseImport(fileNode *ast.File) []*internal.ImportInfo {
	imports := []*internal.ImportInfo{}
	for _, i := range fileNode.Imports {
		alias := ""
		if i.Name != nil {
			alias = i.Name.Name
		}
		imports = append(imports, internal.NewImportInfo(alias, i.Path.Value))
	}
	return imports
}

func parseStruct(fileNode *ast.File, filterTypeNames ...string) []*internal.StructInfo {
	structs := []*internal.StructInfo{}
	imports := parseImport(fileNode)

	// ファイル内にあるtypeの数だけ呼び出される
	ast.Inspect(fileNode, func(node ast.Node) bool {
		t, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}
		if len(filterTypeNames) > 0 && !lists.Contains(filterTypeNames, t.Name.Name) {
			return true
		}
		typeName := t.Name.Name
		fields := []*internal.FieldInfo{}

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
		structInfo := internal.NewStructInfo(typeName, imports, fields)
		structs = append(structs, structInfo)
		return true
	})
	return structs
}
