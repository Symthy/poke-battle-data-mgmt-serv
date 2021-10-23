package property

// ポケモンタイプ
type Types string

const (
	Normal   Types = "Normal"
	Fighting Types = "Fighting"
	Flying   Types = "Flying"
	Poison   Types = "Poison"
	Ground   Types = "Ground"
	Rock     Types = "Rock"
	Bug      Types = "Bug"
	Ghost    Types = "Ghost"
	Steel    Types = "Steel"
	Fire     Types = "Fire"
	Water    Types = "Water"
	Grass    Types = "Grass"
	Electric Types = "Electric"
	Psychic  Types = "Psychic"
	Ice      Types = "Ice"
	Dragon   Types = "Dragon"
	Dark     Types = "Dark"
	None     Types = "None"
)

func (Types) Values() (values []string) {
	for _, s := range []Types{Normal, Fighting, Flying, Poison, Ground, Rock, Bug, Ghost, Steel, Fire, Water, Grass, Electric, Psychic, Ice, Dragon, Dark, None} {
		values = append(values, string(s))
	}
	return
}

// ロール
type Role string

const (
	User                Role = "User"
	ReferenceableMaster Role = "ReferenceableMaster"
	Administrator       Role = "Administrator"
)

func (Role) Values() (values []string) {
	for _, s := range []Role{User, ReferenceableMaster, Administrator} {
		values = append(values, string(s))
	}
	return
}

// 性別
type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

func (Sex) Values() (values []string) {
	for _, s := range []Sex{Male, Female} {
		values = append(values, string(s))
	}
	return
}

// 分類
type Species string

const (
	Physical Species = "Physical" // 物理
	Special  Species = "Special"  // 特殊
)

func (Species) Values() (values []string) {
	for _, s := range []Species{Physical, Special} {
		values = append(values, string(s))
	}
	return
}

// 対戦形式
type BattleFormats string

const (
	Single BattleFormats = "Single"
	Double BattleFormats = "Double"
)

func (BattleFormats) Values() (values []string) {
	for _, s := range []BattleFormats{Single, Double} {
		values = append(values, string(s))
	}
	return
}
