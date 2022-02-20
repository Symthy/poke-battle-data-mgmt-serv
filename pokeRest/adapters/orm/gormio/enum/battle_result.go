package enum

// 対戦結果
type BattleResult string

const (
	Win  BattleResult = "Win"
	Lose BattleResult = "Lose"
	Draw BattleResult = "Draw"
)

func (b BattleResult) String() string {
	return string(b)
}
