package builder

import "github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"

type BattleRecordBuilder struct {
	id                       uint
	partyId                  uint
	battleResult             battles.BattleResult
	battleOpponentPartyId    uint
	selfElectionPokemons     battles.ElectionPokemons[int]
	selfTrainedPokemons      battles.ElectionPokemons[uint]
	opponentElectionPokemons battles.ElectionPokemons[int]
	generation               int
	series                   int
	season                   int
}

func NewBattleRecordBuilder(id uint, partyId uint) {
	
}