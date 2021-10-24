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
type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (Gender) Values() (values []string) {
	for _, s := range []Gender{Male, Female} {
		values = append(values, string(s))
	}
	return
}

// 分類
type MoveSpecies string

const (
	Physical MoveSpecies = "Physical" // 物理
	Special  MoveSpecies = "Special"  // 特殊
	Status   MoveSpecies = "Status"   // 変化
)

func (MoveSpecies) Values() (values []string) {
	for _, s := range []MoveSpecies{Physical, Special, Status} {
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

// 性格
type Nature string

const (
	Lonely  Nature = "Lonely"  // さみしがりや（A↑B↓）
	Adamant Nature = "Adamant" // いじっぱり（A↑C↓）
	Naughty Nature = "Naughty" // やんちゃ（A↑D↓）
	Brave   Nature = "Brave"   // ゆうかん（A↑S↓）

	Bold    Nature = "Bold"    // ずぶとい（B↑A↓）
	Impish  Nature = "Impish"  // わんぱく（B↑C↓）
	Lax     Nature = "Lax"     // のうてんき（B↑D↓）
	Relaxed Nature = "Relaxed" // のんき（B↑S↓）

	Modest Nature = "Modest" // ひかえめ（C↑A↓）
	Mild   Nature = "Mild"   // おっとり（C↑B↓）
	Rash   Nature = "Rash"   // うっかりや（C↑D↓）
	Quiet  Nature = "Quiet"  // れいせい（C↑S↓）

	Calm    Nature = "Calm"    // おだやか（D↑A↓）
	Gentle  Nature = "Gentle"  // おとなしい（D↑B↓）
	Careful Nature = "Careful" // しんちょう（D↑C↓）
	Sassy   Nature = "Sassy"   // なまいき（D↑S↓）

	Timid Nature = "Timid" // おくびょう（S↑A↓）
	Hasty Nature = "Hasty" // せっかち（S↑B↓）
	Jolly Nature = "Jolly" // ようき（S↑C↓）
	Naive Nature = "Naive" // むじゃき（S↑D↓）

	Hardy   Nature = "Hardy"   // がんばりや（補正無し）
	Docile  Nature = "Docile"  // すなお（補正無し）
	Serious Nature = "Serious" // まじめ（補正無し）
	Bashful Nature = "Bashful" // てれや（補正無し）
	Quirky  Nature = "Quirky"  // きまぐれ（補正無し）
)

func (Nature) Values() (values []string) {
	for _, s := range []Nature{
		Lonely, Adamant, Naughty, Brave,
		Bold, Impish, Lax, Relaxed,
		Modest, Mild, Rash, Quiet,
		Calm, Gentle, Careful, Sassy,
		Timid, Hasty, Jolly, Naive,
		Hardy,
	} {
		values = append(values, string(s))
	}
	return
}
