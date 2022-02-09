package migration

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/config"
)

func RunAutoMigration(dbConfig config.DbConfig) {
	db := orm.NewGormDbClientForStdOut(dbConfig).Db()

	db.AutoMigrate(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartyBattleResult{},
		&schema.Pokemon{},
		&schema.Tag{},
		&schema.TrainedPokemonMoveSet{},
		&schema.TrainedPokemon{},
		&schema.BattleRecord{},
	)
}

func RunDropTables(dbConfig config.DbConfig) {
	db := orm.NewGormDbClientForStdOut(dbConfig).Db()
	db.Migrator().DropTable(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.BattleRecord{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartyBattleResult{},
		&schema.Pokemon{},
		&schema.Tag{},
		&schema.TrainedPokemonMoveSet{},
		&schema.TrainedPokemon{},
	)
}
