package migration

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/db"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
)

func RunAutoMigration() {
	db := db.GetDb()

	db.AutoMigrate(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartyResult{},
		&schema.Pokemon{},
		&schema.Tag{},
		&schema.TrainedPokemonBase{},
		&schema.TrainedPokemon{},
		&schema.BattleRecord{},
	)
}

func RunDropTables() {
	db := db.GetDb()
	db.Migrator().DropTable(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.BattleRecord{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartyResult{},
		&schema.Pokemon{},
		&schema.Tag{},
		&schema.TrainedPokemonBase{},
		&schema.TrainedPokemon{},
	)
}