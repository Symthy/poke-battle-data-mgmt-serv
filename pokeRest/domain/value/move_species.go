package value

type MoveSpecies string

const (
	Physical MoveSpecies = "Physical" // 物理
	Special  MoveSpecies = "Special"  // 特殊
	Status   MoveSpecies = "Status"   // 変化
)

func (m MoveSpecies) IsPhysical() bool {
	return m == Physical
}

func (m MoveSpecies) IsSpecial() bool {
	return m == Special
}
