package enum

import "database/sql/driver"

// 分類
type MoveSpecies string

const (
	Physical MoveSpecies = "Physical" // 物理
	Special  MoveSpecies = "Special"  // 特殊
	Status   MoveSpecies = "Status"   // 変化
)

func (mo *MoveSpecies) Scan(value interface{}) error {
	*mo = MoveSpecies(value.([]byte))
	return nil
}

func (mo MoveSpecies) Value() (driver.Value, error) {
	return string(mo), nil
}
