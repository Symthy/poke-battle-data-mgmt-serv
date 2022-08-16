package builder

import (
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

func GenerateSample(outPath string) error {
	f := jen.NewFile("output")
	importNames := map[string]string{
		"github.com/Symthy/PokeRest/pokeRest/domain/entity/items":     "items",
		"github.com/Symthy/PokeRest/pokeRest/domain/value/battles":    "battles",
		"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier": "identifier",
	}
	f.ImportNames(importNames)
	f.ImportAlias("github.com/Symthy/PokeRest/pokeRest/domain/entity/items", "items")
	f.ImportAlias("github.com/Symthy/PokeRest/pokeRest/domain/value/battles", "battles")
	f.ImportAlias("github.com/Symthy/PokeRest/pokeRest/domain/value/identifier", "ident")

	f.Type().Id("Sample").Struct(
		jen.Id("id").Qual("github.com/Symthy/PokeRest/pokeRest/domain/value/identifier", "HeldItemId"),
		jen.Id("name").String(),
		jen.Id("description").String(),
		jen.Id("battleEffects").Op("*").Qual("github.com/Symthy/PokeRest/pokeRest/domain/value/battles", "BattleEffects"),
	)

	err := f.Save(filepath.Join(outPath, "sample.gen.go"))
	if err != nil {
		return err
	}
	return nil
}
