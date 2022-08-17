package enum

import "database/sql/driver"

// 対戦結果
type BattleResult string

const (
	Win  BattleResult = "Win"
	Lose BattleResult = "Lose"
	Draw BattleResult = "Draw"
)

func (b *BattleResult) Scan(value interface{}) error {
	*b = BattleResult(value.([]byte))
	return nil
}

func (b BattleResult) Value() (driver.Value, error) {
	return string(b), nil
}

func (b BattleResult) String() string {
	return string(b)
}
