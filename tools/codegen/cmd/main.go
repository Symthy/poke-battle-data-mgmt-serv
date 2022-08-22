//go:generate go run .
//go:build tools

package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Symthy/PokeRest/tools/codegen/internal/ast"
	"github.com/Symthy/PokeRest/tools/codegen/internal/code"
	"github.com/Symthy/PokeRest/tools/codegen/internal/pkg/filesystem"
)

const (
	schemaPath         = "internal/adapters/orm/gormio/schema/"
	entityPath         = "internal/domain/entity/"
	serverResponsePath = "internal/adapters/rest/autogen/server/"
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

	targetResponses = []string{
		"Ability",
		"BattleOpponentParty",
		"BattleRecord",
		"BattleSeason",
		"CorrectionData",
		"DamageResult",
		"HeldItem",
		"Move",
		"Party",
		"PartyTag",
		"Pokemon",
		"PokemonDetail",
		"TrainedPokemon",
		"TrainedPokemonAttack",
		"TrainedPokemonDeffense",
		"TrainedPokemonDetail",
		"User",
	}
)

func main() {
	rootPath := filepath.Join(filepath.Dir(thisPath), "../../..")
	homePath := filepath.Join(rootPath, "tools/codegen")
	outputPath := filepath.Join(homePath, "output")
	if err := filesystem.MakeDirIfNotExists(outputPath); err != nil {
		fmt.Println(err)
		return
	}

	outSchemaPath := filepath.Join(outputPath, "schema")
	if err := filesystem.MakeDirIfNotExists(outSchemaPath); err != nil {
		fmt.Println(err)
		return
	}
	schemaStructs, err := ast.ParseFiles(filepath.Join(rootPath, schemaPath), targetSchemas)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := code.GenerateFuncSchemaFieldsGetter(homePath, outSchemaPath, schemaStructs); err != nil {
		fmt.Println(err)
		return
	}

	outDomainPath := filepath.Join(outputPath, "domain")
	if err := filesystem.MakeDirIfNotExists(outDomainPath); err != nil {
		fmt.Print(err)
		return
	}
	outFactoryPath := filepath.Join(outDomainPath, "factory")
	if err := filesystem.MakeDirIfNotExists(outFactoryPath); err != nil {
		fmt.Print(err)
		return
	}
	entityStructs, err := ast.ParseFiles(filepath.Join(rootPath, entityPath), targetEntities)
	if err != nil {
		fmt.Print(err)
		return
	}
	if err := code.GenerateEntityFactoryStructs(rootPath, homePath, outFactoryPath, entityStructs); err != nil {
		fmt.Print(err)
		return
	}

	outResponsePath := filepath.Join(outputPath, "response")
	if err := filesystem.MakeDirIfNotExists(outResponsePath); err != nil {
		fmt.Print(err)
		return
	}
	outResDtoPath := filepath.Join(outResponsePath, "dto")
	if err := filesystem.MakeDirIfNotExists(outResDtoPath); err != nil {
		fmt.Print(err)
		return
	}
	responseStructs, err := ast.ParseSingleFile(filepath.Join(rootPath, serverResponsePath, "types.gen.go"), targetResponses...)
	if err != nil {
		fmt.Print(err)
		return
	}
	if err := code.GenerateResponseModelBuilderStruct(homePath, outResDtoPath, responseStructs); err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println()
	fmt.Print("Code generation success.")
}
