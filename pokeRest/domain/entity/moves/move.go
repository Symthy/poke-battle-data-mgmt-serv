package moves

type Move struct {
	id           uint
	name         string
	species      MoveSpecies
	power        int     // 威力
	accuracyRate float32 // 命中率
	pp           int
	isContained  bool
	isCanGuard   bool
}

func NewMove(id uint, name string, species string, power int, accuracyRate float32, pp int,
	isContained bool, isCanGuard bool) Move {
	return Move{
		id:           id,
		name:         name,
		species:      MoveSpecies(species),
		power:        power,
		accuracyRate: accuracyRate,
		pp:           pp,
		isContained:  isContained,
		isCanGuard:   isCanGuard,
	}
}

func (m Move) Id() uint {
	return m.id
}

func (m Move) Name() string {
	return m.name
}

func (m Move) Species() MoveSpecies {
	return m.species
}

func (m Move) Power() int {
	return m.power
}

func (m Move) AccuracyRate() float32 {
	return m.accuracyRate
}

func (m Move) PP() int {
	return m.pp
}

func (m Move) IsContained() bool {
	return m.isContained
}

func (m Move) IsCanGuard() bool {
	return m.isCanGuard
}

type MoveSpecies string

const (
	Physical MoveSpecies = "Physical" // 物理
	Special  MoveSpecies = "Special"  // 特殊
	Status   MoveSpecies = "Status"   // 変化
)
