package battles

import (
	"github.com/Symthy/PokeRest/internal/domain/entity"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

var _ entity.IDomain[identifier.BattleRecordId, uint64] = (*BattleRecord)(nil)

type BattleRecord struct {
	id                       identifier.BattleRecordId
	partyId                  identifier.PartyId
	userId                   identifier.UserId
	battleResult             value.BattleResult
	selfElectionPokemons     *ElectionPokemons
	selfTrainedPokemons      *ElectionTrainedPokemons
	opponentElectionPokemons *ElectionPokemons
	opponentParty            *BattleOpponentParty
	*Season
}

func NewBattleRecord(
	id identifier.BattleRecordId, partyId identifier.PartyId, userId identifier.UserId,
	season *Season, battleResult value.BattleResult, selfElectionPokemons *ElectionPokemons,
	selfTrainedPokemons *ElectionTrainedPokemons, opponentElectionPokemons *ElectionPokemons,
	opponentParty *BattleOpponentParty) *BattleRecord {
	return &BattleRecord{
		id:                       id,
		partyId:                  partyId,
		battleResult:             battleResult,
		selfElectionPokemons:     selfElectionPokemons,
		selfTrainedPokemons:      selfTrainedPokemons,
		opponentElectionPokemons: opponentElectionPokemons,
		opponentParty:            opponentParty,
		Season:                   season,
	}
}

func (b BattleRecord) Id() identifier.BattleRecordId {
	return b.id
}

func (b BattleRecord) PartyId() identifier.PartyId {
	return b.partyId
}

func (b BattleRecord) UserId() identifier.UserId {
	return b.userId
}

func (b BattleRecord) OpponentParty() *BattleOpponentParty {
	return b.opponentParty
}

func (b *BattleRecord) ApplyOpponentParty(opponentParty *BattleOpponentParty) {
	b.opponentParty = opponentParty
}

func (b BattleRecord) Notify(note IBattleRecordNotification) {
	note.SetId(b.id)
	note.SetPartyId(b.partyId)
	note.SetUserId(b.userId)
	note.SetBattleResult(b.battleResult)
	note.SetSelfElectionPokemons(b.selfElectionPokemons)
	note.SetSelfTrainedPokemons(b.selfTrainedPokemons)
	note.SetOpponentElectionPokemons(b.opponentElectionPokemons)
	b.opponentParty.Notify(note)
	b.Season.Notify(note)
}
