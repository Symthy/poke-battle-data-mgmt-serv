package builder

import (
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

func BuildItemFactory(homePath string) error {
	f := jen.NewFile("factory")

	importNames := map[string]string{
		"github.com/Symthy/PokeRest/pokeRest/domain/entity/items":     "items",
		"github.com/Symthy/PokeRest/pokeRest/domain/value/battles":    "battles",
		"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier": "identifier",
	}
	f.ImportNames(importNames)

	f.Type().Id("HeldItemInput").Struct(
		jen.Id("id").Uint64(),
		jen.Id("name").String(),
		jen.Id("description").String(),
		jen.Id("battleEffects").Op("*").Qual("github.com/Symthy/PokeRest/pokeRest/domain/value/battles", "BattleEffects"),
	)

	return f.Save(filepath.Join(homePath, "output/item_input.go"))
}
