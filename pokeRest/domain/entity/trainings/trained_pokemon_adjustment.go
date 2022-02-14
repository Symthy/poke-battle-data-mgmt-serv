package trainings

type TrainedPokemonAdjustment struct {
	id           uint
	pokemonId    int
	nature       Nature
	abilityId    *int
	heldItemId   *int
	effortValueH int
	effortValueA int
	effortValueB int
	effortValueC int
	effortValueD int
	effortValueS int
	moveId1      *int
	moveId2      *int
	moveId3      *int
	moveId4      *int
}

func NewTrainedPokemonAdjustment(id uint) TrainedPokemonAdjustment {
	// Todo: add and validate
	return TrainedPokemonAdjustment{
		id: id,
	}
}

func (t TrainedPokemonAdjustment) Id() uint {
	return t.id
}
func (t TrainedPokemonAdjustment) PokemonId() int {
	return t.pokemonId
}
func (t TrainedPokemonAdjustment) Nature() Nature {
	return t.nature
}
func (t TrainedPokemonAdjustment) AbilityId() *int {
	return t.abilityId
}
func (t TrainedPokemonAdjustment) HeldItemId() *int {
	return t.heldItemId
}

func (t TrainedPokemonAdjustment) EffortValueH() int {
	return t.effortValueH
}
func (t TrainedPokemonAdjustment) EffortValueA() int {
	return t.effortValueA
}
func (t TrainedPokemonAdjustment) EffortValueB() int {
	return t.effortValueB
}
func (t TrainedPokemonAdjustment) EffortValueC() int {
	return t.effortValueC
}
func (t TrainedPokemonAdjustment) EffortValueD() int {
	return t.effortValueD
}
func (t TrainedPokemonAdjustment) EffortValueS() int {
	return t.effortValueS
}

func (t TrainedPokemonAdjustment) MoveIdFirst() *int {
	return t.moveId1
}
func (t TrainedPokemonAdjustment) MoveIdSecond() *int {
	return t.moveId2
}
func (t TrainedPokemonAdjustment) MoveIdThird() *int {
	return t.moveId3
}
func (t TrainedPokemonAdjustment) MoveIdForth() *int {
	return t.moveId4
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

func resolveNature(value string) Nature {
	return Nature(value)
}
