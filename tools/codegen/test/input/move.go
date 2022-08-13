package input

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type Move struct {
	id               identifier.MoveId
	name             string
	description      *string
	effects          *battles.BattleEffects
	acquiredMember   []*Member
	characters       []Character
	leanablePokemons []*identifier.PokemonId
	mixin.UpdateTimes
}
