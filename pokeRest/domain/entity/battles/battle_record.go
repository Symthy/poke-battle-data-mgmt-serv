package battles

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

var _ entity.IDomain[identifier.BattleRecordId] = (*BattleRecord)(nil)

type BattleRecord struct {
	id                       identifier.BattleRecordId
	partyId                  identifier.PartyId
	battleResult             value.BattleResult
	selfElectionPokemons     ElectionPokemons[int]
	selfTrainedPokemons      ElectionPokemons[uint]
	opponentPartyId          identifier.BattleOpponentPartyId
	opponentElectionPokemons ElectionPokemons[int]
	Season
}

func NewBattleRecord(
	id identifier.BattleRecordId, partyId identifier.PartyId, season Season, battleResult value.BattleResult,
	selfElectionPokemons ElectionPokemons[int], selfTrainedPokemons ElectionPokemons[uint],
	opponentPartyId identifier.BattleOpponentPartyId, opponentElectionPokemons ElectionPokemons[int],
) BattleRecord {
	return BattleRecord{
		id:                       id,
		partyId:                  partyId,
		battleResult:             battleResult,
		selfElectionPokemons:     selfElectionPokemons,
		selfTrainedPokemons:      selfTrainedPokemons,
		opponentPartyId:          opponentPartyId,
		opponentElectionPokemons: opponentElectionPokemons,
		Season:                   season,
	}
}

func (b BattleRecord) Id() identifier.BattleRecordId {
	return b.id
}

func (b BattleRecord) PartyId() identifier.PartyId {
	return b.partyId
}

func (b BattleRecord) Notify(note IBattleRecordNotification) {
	note.SetId(b.id)
	note.SetPartyId(b.partyId)
	note.SetBattleResult(b.battleResult)
	note.SetBattleOpponentPartyId(b.opponentPartyId)
	note.SetSelfElectionPokemons(b.selfElectionPokemons)
	note.SetSelfTrainedPokemons(b.selfTrainedPokemons)
	note.SetOpponentElectionPokemons(b.opponentElectionPokemons)
	b.Season.Notify(note)
}

func (b *BattleRecord) ApplyOpponentPartyId(opponentPartyId identifier.BattleOpponentPartyId) {
	b.opponentPartyId = opponentPartyId
}
