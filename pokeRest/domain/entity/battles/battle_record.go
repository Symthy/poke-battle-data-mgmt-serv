package battles

type BattleRecord struct {
	id                       uint
	partyId                  uint
	generation               int
	series                   int
	season                   int
	battleResult             BattleResult
	battleOpponentPartyId    uint
	selfElectionPokemons     ElectionPokemons[int]
	selfTrainedPokemons      ElectionPokemons[uint]
	opponentElectionPokemons ElectionPokemons[int]
}

func NewBattleRecord(
	id, partyId uint, generation, series, season int, battleResult string, battleOpponentPartyId uint,
	selfElectionPokemons []int, selfTrainedPokemons []uint, opponentElectionPokemons []int) BattleRecord {
	return BattleRecord{
		id:                       id,
		partyId:                  partyId,
		generation:               generation,
		series:                   series,
		season:                   season,
		battleResult:             BattleResult(battleResult),
		battleOpponentPartyId:    battleOpponentPartyId,
		selfElectionPokemons:     NewElectionPokemons(selfElectionPokemons),
		selfTrainedPokemons:      NewElectionPokemons(selfTrainedPokemons),
		opponentElectionPokemons: NewElectionPokemons(opponentElectionPokemons),
	}
}

func (b BattleRecord) Id() uint {
	return b.id
}

func (b BattleRecord) PartyId() uint {
	return b.partyId
}

func (b BattleRecord) Generation() int {
	return b.generation
}

func (b BattleRecord) Series() int {
	return b.series
}

func (b BattleRecord) Season() int {
	return b.season
}

func (b BattleRecord) BattleResult() BattleResult {
	return b.battleResult
}

func (b BattleRecord) BattleOpponentPartyId() uint {
	return b.battleOpponentPartyId
}

func (b BattleRecord) SelfElectionPokemons() ElectionPokemons[int] {
	return b.selfElectionPokemons
}

func (b BattleRecord) SelfTrainedPokemons() ElectionPokemons[uint] {
	return b.selfTrainedPokemons
}

func (b BattleRecord) OpponentPokemons() ElectionPokemons[int] {
	return b.opponentElectionPokemons
}

type BattleResult string

const (
	WIN  BattleResult = "Win"
	LOSE BattleResult = "Lose"
	DRAW BattleResult = "Draw"
)
