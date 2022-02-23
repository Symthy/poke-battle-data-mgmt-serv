package battles

type BattleRecord struct {
	id                       uint
	partyId                  uint
	battleResult             BattleResult
	battleOpponentPartyId    uint
	selfElectionPokemons     ElectionPokemons[int]
	selfTrainedPokemons      ElectionPokemons[uint]
	opponentElectionPokemons ElectionPokemons[int]
	Season
}

func NewBattleRecord(
	id, partyId uint, generation, series, season int, battleResult string, battleOpponentPartyId uint,
	selfElectionPokemons []int, selfTrainedPokemons []uint, opponentElectionPokemons []int) BattleRecord {
	return BattleRecord{
		id:                       id,
		partyId:                  partyId,
		battleResult:             BattleResult(battleResult),
		battleOpponentPartyId:    battleOpponentPartyId,
		selfElectionPokemons:     NewElectionPokemons(selfElectionPokemons),
		selfTrainedPokemons:      NewElectionPokemons(selfTrainedPokemons),
		opponentElectionPokemons: NewElectionPokemons(opponentElectionPokemons),
		Season: Season{
			generation: generation,
			series:     series,
			season:     season,
		},
	}
}

func (b BattleRecord) Id() uint {
	return b.id
}

func (b BattleRecord) PartyId() uint {
	return b.partyId
}

func (b BattleRecord) Notify(note *IBattleRecordNotification) {
	(*note).Id(b.id)
	(*note).PartyId(b.partyId)
	(*note).Generation(b.generation)
	(*note).Series(b.series)
	(*note).Season(b.season)
	(*note).BattleResult(b.battleResult.String())
	(*note).BattleOpponentPartyId(b.battleOpponentPartyId)
	(*note).SelfElectionPokemons(b.selfElectionPokemons.Ids())
	(*note).SelfTrainedPokemons(b.selfTrainedPokemons.Ids())
	(*note).OpponentElectionPokemons(b.opponentElectionPokemons.Ids())
}

func (b BattleRecord) solvedSeason() bool {
	return b.generation != 0 && b.series != 0 && b.season == 0
}

func (b BattleRecord) ApplyOpponentPartyId(opponentPartyId uint) {
	b.battleOpponentPartyId = opponentPartyId
}

type BattleResult string

const (
	WIN  BattleResult = "Win"
	LOSE BattleResult = "Lose"
	DRAW BattleResult = "Draw"
)

func (b BattleResult) String() string {
	return string(b)
}
