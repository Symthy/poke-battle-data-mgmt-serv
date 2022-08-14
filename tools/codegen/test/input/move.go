package input

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/mixin"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	ident "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type Move struct {
	id               ident.MoveId
	name             string
	description      *string
	effects          *battles.BattleEffects
	acquiredMember   []*Member
	characters       []Character
	leanablePokemons []*ident.PokemonId
	mixin.UpdateTimes
}

type Test struct {
	test string
}
