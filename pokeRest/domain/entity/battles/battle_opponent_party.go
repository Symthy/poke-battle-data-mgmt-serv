package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattleOpponentParty struct {
	id                 uint
	opponentPokemonIds value.PartyPokemonIds
}

func NewBattleOpponentParty(id uint, pokemonIds ...int) BattleOpponentParty {
	return BattleOpponentParty{
		id:                 id,
		opponentPokemonIds: value.NewPartyPokemonIds(pokemonIds...),
	}
}

func (b BattleOpponentParty) Id() uint {
	return b.id
}

func (b BattleOpponentParty) OpponentPokemonIds() value.PartyPokemonIds {
	return b.opponentPokemonIds
}
