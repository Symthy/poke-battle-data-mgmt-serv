package builder

import (
	"os"
	"path/filepath"

	"github.com/Symthy/PokeRest/tools/codegen/internal"
	"github.com/dave/jennifer/jen"
)

func BuildItemFactory(homePath string) error {
	f := jen.NewFile("factory")

	// importNames := map[string]string{
	// 	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items":     "items",
	// 	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles":    "battles",
	// 	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier": "identifier",
	// }
	// f.ImportNames(importNames)
	f.ImportAlias("github.com/Symthy/PokeRest/pokeRest/domain/entity/items", "items")
	f.ImportAlias("github.com/Symthy/PokeRest/pokeRest/domain/value/battles", "battles")
	f.ImportAlias("github.com/Symthy/PokeRest/pokeRest/domain/value/identifier", "ident")

	f.Type().Id("HeldItemInput").Struct(
		jen.Id("id").Qual("github.com/Symthy/PokeRest/pokeRest/domain/value/identifier", "HeldItemId"),
		jen.Id("name").String(),
		jen.Id("description").String(),
		jen.Id("battleEffects").Op("*").Qual("github.com/Symthy/PokeRest/pokeRest/domain/value/battles", "BattleEffects"),
	)

	return f.Save(filepath.Join(homePath, "output/item_input.go"))
}

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

	if !existsDir(outPath) {
		if err := os.Mkdir(outPath, 0666); err != nil {
			return err
		}
	}
	newFilePath := filepath.Join(outPath, "schema_fields.go")
	return f.Save(newFilePath)
}

func existsDir(dirPath string) bool {
	if f, err := os.Stat(dirPath); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}
