package battles

type BattleRecord struct {
	id                       uint
	partyId                  uint
	generation               int
	series                   int
	season                   int
	battleResult             BattleResult
	battleOpponentPartyId    uint
	selfElectionPokemons     ElectionPokemons
	selfTrainedPokemons      ElectionPokemons
	opponentElectionPokemons ElectionPokemons
}

func NewBattleRecord(id uint) BattleRecord {
	return BattleRecord{
		id: id,
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

func (b BattleRecord) SelfElectionPokemons() ElectionPokemons {
	return b.selfElectionPokemons
}

func (b BattleRecord) SelfTrainedPokemons() ElectionPokemons {
	return b.selfTrainedPokemons
}

func (b BattleRecord) OpponentPokemons() ElectionPokemons {
	return b.opponentElectionPokemons
}

type BattleResult int

const (
	WIN  BattleResult = 0
	LOSE BattleResult = 1
	DRAW BattleResult = 2
)
