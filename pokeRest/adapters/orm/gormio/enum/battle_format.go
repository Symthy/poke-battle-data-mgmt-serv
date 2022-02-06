package enum

import "database/sql/driver"

// 対戦形式
type BattleFormat string

const (
	Single BattleFormat = "Single"
	Double BattleFormat = "Double"
)

func (ba *BattleFormat) Scan(value interface{}) error {
	*ba = BattleFormat(value.([]byte))
	return nil
}

func (ba BattleFormat) Value() (driver.Value, error) {
	return string(ba), nil
}
