package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleOpponentPartyInput struct {
	id                 uint
	opponentPokemonIds []uint
}

func NewBattleOpponentPartyInput(id uint, pokemonIds ...uint) BattleOpponentPartyInput {
	return BattleOpponentPartyInput{id, pokemonIds}
}

func (i BattleOpponentPartyInput) BuildDomain() (*battles.BattleOpponentParty, error) {
	id, err := identifier.NewBattleOpponentPartyId(i.id)
	if err != nil {
		return nil, err
	}
	opponentPokemonIds := value.NewPartyPokemonIds(i.opponentPokemonIds)
	domain := battles.NewBattleOpponentParty(*id, opponentPokemonIds)
	return &domain, nil
}
