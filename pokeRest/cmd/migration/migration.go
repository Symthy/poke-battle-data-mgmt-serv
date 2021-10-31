package migration

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/db"
	"github.com/Symthy/PokeRest/pokeRest/adapters/db/orm/gormio/schema"
)

func RunAutoMigration() {
	db := db.NewGormDbAccessor().GetDb()

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
	db := db.NewGormDbAccessor().GetDb()
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
