package factory

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type BattleRecordInput struct {
	id                            uint
	partyId                       uint
	userId                        uint
	generation                    int
	series                        int
	season                        int
	battleResult                  string
	selfElectionPokemonIds        []uint
	selfElectionTrainedPokemonIds []uint
	opponentElectionPokemonIds    []uint
	BattleOpponentPartyInput
}

func NewBattleRecordInput(
	id, partyId, userId uint, generation, series, season int, battleResult string,
	selfElectionPokemonIds, selfElectionTrainedPokemonIds, opponentElectionPokemonIds []uint,
	opponentPartyInput BattleOpponentPartyInput,
) BattleRecordInput {
	return BattleRecordInput{
		id:                            id,
		partyId:                       partyId,
		userId:                        userId,
		generation:                    generation,
		series:                        series,
		season:                        season,
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

func (i BattleRecordInput) Id(id uint) {
	i.id = id
}
func (i BattleRecordInput) PartyId(partyId uint) {
	i.partyId = partyId
}
func (i BattleRecordInput) UserId(userId uint) {
	i.userId = userId
}
func (i BattleRecordInput) Generation(generation int) {
	i.generation = generation
}
func (i BattleRecordInput) Series(series int) {
	i.series = series
}
func (i BattleRecordInput) Season(season int) {
	i.season = season
}
func (i BattleRecordInput) BattleResult(battleResult string) {
	i.battleResult = battleResult
}
func (i BattleRecordInput) SelfElectionPokemonIds(ids []uint) {
	i.selfElectionPokemonIds = ids
}
func (i BattleRecordInput) OpponentElectionPokemonIds(ids []uint) {
	i.opponentElectionPokemonIds = ids
}
func (i BattleRecordInput) BattleOpponentParty(input BattleOpponentPartyInput) {
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
	season, err := battles.NewSeason(b.generation, b.series, b.season)
	if err != nil {
		return nil, err
	}
	result := value.BattleResult(b.battleResult)
	selfPokemons := battles.NewElectionPokemons(b.selfElectionPokemonIds)
	selfTrainedPokemons := battles.NewElectionPokemons(b.selfElectionTrainedPokemonIds)
	opponentPokemons := battles.NewElectionPokemons(b.opponentElectionPokemonIds)
	opponentParty, err := b.BattleOpponentPartyInput.BuildDomain()
	if err != nil {
		return nil, err
	}
	domain := battles.NewBattleRecord(*id, *partyId, *userId, *season, result, selfPokemons, selfTrainedPokemons,
		opponentPokemons, *opponentParty)
	return &domain, nil
}
