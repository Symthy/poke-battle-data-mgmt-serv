package factory

import (
	"github.com/Symthy/PokeRest/internal/domain/entity/battles"
	"github.com/Symthy/PokeRest/internal/domain/value"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type BattleRecordInput struct {
	id                            uint64
	partyId                       uint64
	userId                        uint64
	season                        *SeasonInput
	battleResult                  string
	selfElectionPokemonIds        []uint64
	selfElectionTrainedPokemonIds []uint64
	opponentElectionPokemonIds    []uint64
	opponentParty                 *BattleOpponentPartyInput
}

func NewBattleRecordInput(
	id, partyId, userId uint64, battleResult string, seasonInput *SeasonInput,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds []uint64,
	opponentPartyInput *BattleOpponentPartyInput,
) BattleRecordInput {
	return BattleRecordInput{
		id:                            id,
		partyId:                       partyId,
		userId:                        userId,
		season:                        seasonInput,
		battleResult:                  battleResult,
		selfElectionPokemonIds:        selfElectionPokemonIds,
		selfElectionTrainedPokemonIds: selfElectionTrainedPokemonIds,
		opponentElectionPokemonIds:    opponentElectionPokemonIds,
		opponentParty:                 opponentPartyInput,
	}
}

func NewBattleRecordBuilder() *BattleRecordInput {
	return &BattleRecordInput{}
}

func (i *BattleRecordInput) Id(id uint64) *BattleRecordInput {
	i.id = id
	return i
}
func (i *BattleRecordInput) PartyId(partyId uint64) *BattleRecordInput {
	i.partyId = partyId
	return i
}
func (i *BattleRecordInput) UserId(userId uint64) *BattleRecordInput {
	i.userId = userId
	return i
}
func (i *BattleRecordInput) Season(season *SeasonInput) *BattleRecordInput {
	i.season = season
	return i
}
func (i *BattleRecordInput) BattleResult(battleResult string) *BattleRecordInput {
	i.battleResult = battleResult
	return i
}
func (i *BattleRecordInput) SelfElectionPokemonIds(ids []uint64) *BattleRecordInput {
	i.selfElectionPokemonIds = ids
	return i
}
func (i *BattleRecordInput) SelfElectionTrainedPokemonIds(ids []uint64) *BattleRecordInput {
	i.selfElectionTrainedPokemonIds = ids
	return i
}
func (i *BattleRecordInput) OpponentElectionPokemonIds(ids []uint64) *BattleRecordInput {
	i.opponentElectionPokemonIds = ids
	return i
}
func (i *BattleRecordInput) BattleOpponentParty(input *BattleOpponentPartyInput) *BattleRecordInput {
	i.opponentParty = input
	return i
}

func (b BattleRecordInput) BuildDomain() (*battles.BattleRecord, error) {
	id, err := identifier.NewBattleRecordId(b.id)
	if err != nil {
		return nil, err
	}
	partyId, err := identifier.NewPartyId(b.partyId)
	if err != nil {
		return nil, err
	}
	userId, err := identifier.NewUserId(b.userId)
	if err != nil {
		return nil, err
	}
	season, err := b.season.BuildDomain()
	if err != nil {
		return nil, err
	}
	result := value.BattleResult(b.battleResult)
	selfPokemons := battles.NewElectionPokemons(b.selfElectionPokemonIds)
	selfTrainedPokemons := battles.NewElectionTrainedPokemons(b.selfElectionTrainedPokemonIds)
	opponentPokemons := battles.NewElectionPokemons(b.opponentElectionPokemonIds)
	opponentParty, err := b.opponentParty.BuildDomain()
	if err != nil {
		return nil, err
	}
	domain := battles.NewBattleRecord(*id, *partyId, *userId, season, result, selfPokemons, selfTrainedPokemons,
		opponentPokemons, opponentParty)
	return domain, nil
}
