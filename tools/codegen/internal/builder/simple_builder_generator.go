package builder

import (
	"fmt"
	"path/filepath"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/Symthy/PokeRest/tools/codegen/internal/pkg/strs"
	"github.com/dave/jennifer/jen"
)

func GenerateSimpleBuilderStruct(homePath, outPath, pkgName string, responseModelStructs []*internal.StructInfo) error {
	isErr := false
	for _, responseModelStruct := range responseModelStructs {
		f := jen.NewFile(pkgName)
		addImportStatement(f, responseModelStruct)

		fieldStatements := []*jen.Statement{}
		funcStatements := []*jen.Statement{}
		builderTypeName := responseModelStruct.Name() + "Builder"
		for _, field := range responseModelStruct.Fields() {
			fieldName := strs.ToTopLower(field.Name())
			fieldStatement := jen.Id(fieldName)
			typeStatement := resolveTypeStatement(field, responseModelStruct)

			fieldStatement.Add(typeStatement)
			fieldStatements = append(fieldStatements, fieldStatement)

			funcStatement := buildSetterMethodStatement(
				"b", builderTypeName, field.Name(), fieldName, field.TypeName(), typeStatement, field)
			funcStatements = append(funcStatements, funcStatement)
		}

		f.Type().Id(builderTypeName).StructFunc(func(g *jen.Group) {
			for _, fieldStatement := range fieldStatements {
				g.Add(fieldStatement)
			}
		})

		f.Add(buildBuilderConstructorStatement(responseModelStruct.Name(), builderTypeName))

		for _, funcStatement := range funcStatements {
			f.Add(funcStatement)
		}

		err := f.Save(filepath.Join(outPath, strs.ToSnakeCase(builderTypeName)+".gen.go"))
		if err != nil {
			fmt.Printf("[Error] failed to simple builder code gen (%s): %v\n", responseModelStruct.Name(), err)
			isErr = true
		}
	}

	if isErr {
		return fmt.Errorf("failed to simple builder structs code generation.")
	}
	fmt.Println("success to simple builder structs code generation.")
	return nil
}
