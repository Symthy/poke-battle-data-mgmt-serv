package battles

import "github.com/Symthy/PokeRest/pokeRest/domain/value"

type BattleOpponentParty struct {
	id               uint
	opponentPokemons value.PartyPokemons
}

func NewBattleOpponentParty(id uint) BattleOpponentParty {
	return BattleOpponentParty{
		id: id,
	}
}

func (b BattleOpponentParty) Id() uint {
	return b.id
}

func (b BattleOpponentParty) OpponentPokemons() value.PartyPokemons {
	return b.opponentPokemons
}
