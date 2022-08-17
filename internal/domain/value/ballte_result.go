package value

type BattleResult string

const (
	WIN  BattleResult = "Win"
	LOSE BattleResult = "Lose"
	DRAW BattleResult = "Draw"
	None BattleResult = ""
)

func (b BattleResult) String() string {
	return string(b)
}
