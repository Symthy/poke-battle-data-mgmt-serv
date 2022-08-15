package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleOpponentPartyInput struct {
	id                 uint64
	opponentPokemonIds []uint64
}

func NewBattleOpponentPartyInput(id uint64, pokemonIds []uint64) *BattleOpponentPartyInput {
	return &BattleOpponentPartyInput{id, pokemonIds}
}

func NewBattleOpponentPartyBuilder() *BattleOpponentPartyInput {
	return &BattleOpponentPartyInput{}
}

func (b *BattleOpponentPartyInput) Id(id uint64) *BattleOpponentPartyInput {
	b.id = id
	return b
}

func (b *BattleOpponentPartyInput) OpponentPokemonIds(ids []uint64) *BattleOpponentPartyInput {
	b.opponentPokemonIds = ids
	return b
}

func (i BattleOpponentPartyInput) BuildDomain() (*battles.BattleOpponentParty, error) {
	id, err := identifier.NewBattleOpponentPartyId(i.id)
	if err != nil {
		return nil, err
	}
	// Todo: validate upper limit
	opponentPokemonIds := value.NewPartyPokemonIds(i.opponentPokemonIds)
	domain := battles.NewBattleOpponentParty(*id, opponentPokemonIds)
	return &domain, nil
}
