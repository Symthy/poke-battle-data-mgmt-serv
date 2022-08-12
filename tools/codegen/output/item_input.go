package factory

import "github.com/Symthy/PokeRest/pokeRest/domain/value/battles"

type HeldItemInput struct {
	id            uint64
	name          string
	description   string
	battleEffects *battles.BattleEffects
}
