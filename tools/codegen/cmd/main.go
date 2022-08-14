package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Symthy/PokeRest/tools/codegen/internal/builder"
	"github.com/Symthy/PokeRest/tools/codegen/internal/parser"
)

const (
	schemaPath    = "pokeRest/adapters/orm/gormio/schema/"
	entityPath    = "pokeRest/domain/entity"
	testInputPath = "tools/codegen/test/input"
)

var (
	_, thisPath, _, _ = runtime.Caller(0)

	targetSchemas = map[string]string{ // key: file name, value: struct name
		"ability.go":                        "AbilitySchema",
		"battle_opponent_party.go":          "BattleOpponentPartySchema",
		"battle_record.go":                  "BattleRecordSchema",
		"battle_season.go":                  "BattleSeasonSchema",
		"form.go":                           "FormSchema",
		"held_item.go":                      "HeldItemSchema",
		"move.go":                           "MoveSchema",
		"party_season_result.go":            "PartySeasonResultSchema",
		"party_tag.go":                      "PartyTagSchema",
		"party.go":                          "PartySchema",
		"trained_pokemon_adjustment.go":     "TrainedPokemonAdjustmentSchema",
		"trained_pokemon_attack_target.go":  "TrainedPokemonAttackTargetSchema",
		"trained_pokemon_defense_target.go": "TrainedPokemonDefenseTargetSchema",
		"trained_pokemon.go":                "TrainedPokemonSchema",
	}
)

func main() {
	rootPath := filepath.Join(filepath.Dir(thisPath), "../../..")
	// parser.Parse(filepath.Join(rootPath, testInputPath, "move.go"))

	schemaStructs, err := parser.ParseFiles(filepath.Join(rootPath, schemaPath), targetSchemas)
	if err != nil {
		fmt.Print(err)
		return
	}

	homePath := filepath.Join(rootPath, "tools/codegen")
	err = builder.GenerateFuncSchemaFieldsGetter(homePath, schemaStructs)
	// err := builder.BuildItemFactory(filepath.Join(rootPath, "tools/codegen"))
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print("Code generation success.")
}
