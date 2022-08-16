package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Symthy/PokeRest/tools/codegen/internal/ast"
	"github.com/Symthy/PokeRest/tools/codegen/internal/builder"
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

	targetEntities = map[string]string{
		"abilities/Ability.go":                        "",
		"battles/battle_record.go":                    "",
		"battles/battle_opponent_party.go":            "",
		"battles/season.go":                           "",
		"items/held_item.go":                          "",
		"moves/move.go":                               "",
		"parties/party.go":                            "",
		"parties/party_battle_result.go":              "",
		"parties/party_tag.go":                        "",
		"pokemons/pokemon.go":                         "",
		"trainings/trained_pokemon_adjustment.go":     "",
		"trainings/trained_pokemon_attack_target.go":  "",
		"trainings/trained_pokemon_defense_target.go": "",
		"trainings/trained_pokemon.go":                "",
		"users/user.go":                               "",
	}
)

func main() {
	rootPath := filepath.Join(filepath.Dir(thisPath), "../../..")
	homePath := filepath.Join(rootPath, "tools/codegen")

	schemaStructs, err := ast.ParseFiles(filepath.Join(rootPath, schemaPath), targetSchemas)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = builder.GenerateFuncSchemaFieldsGetter(homePath, schemaStructs)
	if err != nil {
		fmt.Println(err)
		return
	}

	entityStructs, err := ast.ParseFiles(filepath.Join(rootPath, entityPath), targetEntities)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = builder.GenerateEntityFactoryStructs(rootPath, homePath, entityStructs)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println()
	fmt.Print("Code generation success.")
}
