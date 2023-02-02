package migration

import (
	"fmt"
	"log"

	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/config"
)

func CreateTable() {
	conf, err := config.LoadConfigYaml()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print("Start Migration")
	runAutoMigration(conf.DbConfig)
	log.Print("End Migration")
}

func AllTableDrop() {
	conf, err := config.LoadConfigYaml()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print("Start Drop All Tables")
	runDropTables(conf.DbConfig)
	log.Print("End Drop All Tables")
}

func runAutoMigration(dbConfig config.DbConfig) {
	db := orm.NewGormDbClientForStdOut(dbConfig).Db()
	err := db.AutoMigrate(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.BattleSeason{},
		&schema.BattleRecord{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartySeasonResult{},
		&schema.Pokemon{},
		&schema.PartyTag{},
		&schema.TrainedPokemon{},
		&schema.TrainedPokemonAdjustment{},
		&schema.TrainedPokemonAttackTarget{},
		&schema.TrainedPokemonDefenseTarget{},
		&schema.User{},
	)
	if err != nil {
		fmt.Println(err)
	}
}

func runDropTables(dbConfig config.DbConfig) {
	db := orm.NewGormDbClientForStdOut(dbConfig).Db()
	err := db.Migrator().DropTable(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.BattleRecord{},
		&schema.BattleSeason{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartySeasonResult{},
		&schema.Pokemon{},
		&schema.PartyTag{},
		&schema.TrainedPokemon{},
		&schema.TrainedPokemonAdjustment{},
		&schema.TrainedPokemonAttackTarget{},
		&schema.TrainedPokemonDefenseTarget{},
		&schema.User{},
		"pokemons_moves",
		"trained_pokemons_parties",
		"parties_party_tags",
	)
	if err != nil {
		fmt.Println(err)
	}
}
