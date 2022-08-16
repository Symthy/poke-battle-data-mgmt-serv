package builder

import (
	"fmt"
	"path/filepath"

	"github.com/Symthy/PokeRest/tools/codegen/internal"

	"github.com/dave/jennifer/jen"
)

const outputGoFileName = "schema_fields.gen.go"

func GenerateFuncSchemaFieldsGetter(homePath, outPath string, structs []*internal.StructInfo) error {
	f := jen.NewFile("schema")

	for _, structInfo := range structs {
		fieldsVarName := "fields"
		funcName := structInfo.Name() + "Fields"
		f.Func().Id(funcName).Params().Index().String().Block(
			jen.Id(fieldsVarName).Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
				for _, fieldName := range structInfo.FieldNames() {
					g.Lit(fieldName)
				}
			}),
			jen.Return(jen.Id(fieldsVarName)),
		)
	}
	newFilePath := filepath.Join(outPath, outputGoFileName)
	err := f.Save(newFilePath)
	if err != nil {
		return fmt.Errorf("[Error] failed to generate: %v", err)
	}
	fmt.Printf("success to code generation: %s\n", outputGoFileName)
	return nil
}
