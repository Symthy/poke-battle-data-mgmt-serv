package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleRecordInput struct {
	id                            uint64
	partyId                       uint64
	userId                        uint64
	generation                    uint64
	series                        uint64
	season                        uint64
	battleResult                  string
	selfElectionPokemonIds        []uint64
	selfElectionTrainedPokemonIds []uint64
	opponentElectionPokemonIds    []uint64
	BattleOpponentPartyInput
}

func NewBattleRecordInput(
	id, partyId, userId, generation, series, season uint64, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds []uint64,
	opponentPartyInput BattleOpponentPartyInput,
) BattleRecordInput {
	return BattleRecordInput{
		id:      id,
		partyId: partyId,
		userId:  userId,
		// Todo
		battleResult:                  battleResult,
		selfElectionPokemonIds:        selfElectionPokemonIds,
		selfElectionTrainedPokemonIds: selfElectionTrainedPokemonIds,
		opponentElectionPokemonIds:    opponentElectionPokemonIds,
		BattleOpponentPartyInput:      opponentPartyInput,
	}
}

func NewBattleRecordBuilder() BattleRecordInput {
	return BattleRecordInput{}
}

func (i *BattleRecordInput) Id(id uint64) {
	i.id = id
}
func (i *BattleRecordInput) PartyId(partyId uint64) {
	i.partyId = partyId
}
func (i *BattleRecordInput) UserId(userId uint64) {
	i.userId = userId
}
func (i *BattleRecordInput) Generation(generation uint64) {
	i.generation = generation
}
func (i *BattleRecordInput) Series(series uint64) {
	i.series = series
}
func (i *BattleRecordInput) Season(season uint64) {
	i.season = season
}
func (i *BattleRecordInput) BattleResult(battleResult string) {
	i.battleResult = battleResult
}
func (i *BattleRecordInput) SelfElectionPokemonIds(ids []uint64) {
	i.selfElectionPokemonIds = ids
}
func (i *BattleRecordInput) OpponentElectionPokemonIds(ids []uint64) {
	i.opponentElectionPokemonIds = ids
}
func (i *BattleRecordInput) BattleOpponentParty(input BattleOpponentPartyInput) {
	i.BattleOpponentPartyInput = input
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
	// Todo
	season, err := battles.NewSeason(0, 0, 0)
	if err != nil {
		return nil, err
	}
	result := value.BattleResult(b.battleResult)
	selfPokemons := battles.NewElectionPokemons(b.selfElectionPokemonIds)
	selfTrainedPokemons := battles.NewElectionTrainedPokemons(b.selfElectionTrainedPokemonIds)
	opponentPokemons := battles.NewElectionPokemons(b.opponentElectionPokemonIds)
	opponentParty, err := b.BattleOpponentPartyInput.BuildDomain()
	if err != nil {
		return nil, err
	}
	domain := battles.NewBattleRecord(*id, *partyId, *userId, *season, result, selfPokemons, selfTrainedPokemons,
		opponentPokemons, *opponentParty)
	return &domain, nil
}
