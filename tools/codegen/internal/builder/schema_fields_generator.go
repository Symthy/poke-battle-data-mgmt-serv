package builder

import (
	"path/filepath"

	"github.com/Symthy/PokeRest/tools/codegen/filesystem"
	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/dave/jennifer/jen"
)

func GenerateFuncSchemaFieldsGetter(homePath string, structs []*internal.StructInfo) error {
	outPath := filepath.Join(homePath, "output/schema")
	f := jen.NewFile("schema")

	for _, structInfo := range structs {
		fieldsVarName := "fields"
		f.Func().Id(structInfo.Name()+"Fields").Params().Index().String().Block(
			jen.Id(fieldsVarName).Op(":=").Index().String().ValuesFunc(func(g *jen.Group) {
				for _, fieldName := range structInfo.FieldNames() {
					g.Lit(fieldName)
				}
			}),
			jen.Return(jen.Id(fieldsVarName)),
		)
	}

	filesystem.MakeDirIfNotExists(outPath)
	newFilePath := filepath.Join(outPath, "schema_fields.go")
	return f.Save(newFilePath)
}
