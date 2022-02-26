package inputs

type InputBattleRecord struct {
	id                         uint
	partyId                    uint
	generation                 int
	series                     int
	season                     int
	battleResult               string
	selfPokemonIds             []int
	selfTrainedPokemonIds      []uint
	opponentElectionPokemonIds []int
	opponentPartyMember        []int
}

func NewInputBattleRecord(
	id, partyId uint, generation, series, season int, battleResult string, selfPokemonIds []int,
	selfTrainedPokemonIds []uint, opponentElectionPokemonIds []int, opponentPartyMember []int,
) InputBattleRecord {
	return InputBattleRecord{
		id:                         id,
		partyId:                    partyId,
		generation:                 generation,
		series:                     series,
		season:                     season,
		battleResult:               battleResult,
		selfPokemonIds:             selfPokemonIds,
		selfTrainedPokemonIds:      selfTrainedPokemonIds,
		opponentElectionPokemonIds: opponentElectionPokemonIds,
		opponentPartyMember:        opponentPartyMember,
	}
}

func (i InputBattleRecord) Id() uint {
	return i.id
}
func (i InputBattleRecord) PartyId() uint {
	return i.partyId
}
func (i InputBattleRecord) Generation() int {
	return i.generation
}
func (i InputBattleRecord) Series() int {
	return i.series
}
func (i InputBattleRecord) Season() int {
	return i.season
}
func (i InputBattleRecord) BattleResult() string {
	return i.battleResult
}
func (i InputBattleRecord) SelfPokemonIds() []int {
	return i.selfPokemonIds
}
func (i InputBattleRecord) SelfTrainedPokemonIds() []uint {
	return i.selfTrainedPokemonIds
}
func (i InputBattleRecord) OpponentElectionPokemonIds() []int {
	return i.opponentElectionPokemonIds
}
func (i InputBattleRecord) OpponentPartyMember() []int {
	return i.opponentPartyMember
}
