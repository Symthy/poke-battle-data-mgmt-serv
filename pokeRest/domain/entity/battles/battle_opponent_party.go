package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.BattleOpponentPartyId, uint64] = (*BattleOpponentParty)(nil)

type BattleOpponentParty struct {
	id         identifier.BattleOpponentPartyId
	pokemonIds *value.PartyPokemonIds
}

func NewBattleOpponentPartyOfUnregister(pokemonIds *value.PartyPokemonIds) BattleOpponentParty {
	return NewBattleOpponentParty(identifier.NewEmptyBattleOpponentPartyId(), pokemonIds)
}

func NewBattleOpponentParty(
	id identifier.BattleOpponentPartyId, pokemonIds *value.PartyPokemonIds,
) BattleOpponentParty {
	return BattleOpponentParty{
		id:         id,
		pokemonIds: pokemonIds,
	}
}

func (b BattleOpponentParty) Id() identifier.BattleOpponentPartyId {
	return b.id
}

func (b BattleOpponentParty) Notify(note IBattleOpponentPartyNotification) {
	note.SetOpponentPartyId(b.id)
	note.SetPokemonIds(b.pokemonIds)
}
