package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattleOpponentParty struct {
	id                 uint
	opponentPokemonIds value.PartyPokemonIds
}

func NewBattleOpponentPartyOfUnregister(pokemonIds value.PartyPokemonIds) BattleOpponentParty {
	return NewBattleOpponentParty(0, pokemonIds)
}

func NewBattleOpponentParty(id uint, pokemonIds value.PartyPokemonIds) BattleOpponentParty {
	return BattleOpponentParty{
		id:                 id,
		opponentPokemonIds: pokemonIds,
	}
}

func (b BattleOpponentParty) Id() uint {
	return b.id
}

func (b BattleOpponentParty) OpponentPokemonIds() value.PartyPokemonIds {
	return b.opponentPokemonIds
}
