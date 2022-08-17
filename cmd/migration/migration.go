package migration

import (
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/internal/config"
)

func RunAutoMigration(dbConfig config.DbConfig) {
	db := orm.NewGormDbClientForStdOut(dbConfig).Db()
	db.AutoMigrate(
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
}

func RunDropTables(dbConfig config.DbConfig) {
	db := orm.NewGormDbClientForStdOut(dbConfig).Db()
	db.Migrator().DropTable(
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
}
