package enum

import (
	"database/sql/driver"

	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

// ポケモンタイプ
type PokemonType string

const (
	Normal   PokemonType = "Normal"
	Fighting PokemonType = "Fighting"
	Flying   PokemonType = "Flying"
	Poison   PokemonType = "Poison"
	Ground   PokemonType = "Ground"
	Rock     PokemonType = "Rock"
	Bug      PokemonType = "Bug"
	Ghost    PokemonType = "Ghost"
	Steel    PokemonType = "Steel"
	Fire     PokemonType = "Fire"
	Water    PokemonType = "Water"
	Grass    PokemonType = "Grass"
	Electric PokemonType = "Electric"
	Psychic  PokemonType = "Psychic"
	Ice      PokemonType = "Ice"
	Dragon   PokemonType = "Dragon"
	Dark     PokemonType = "Dark"
	Fairy    PokemonType = "Fairy"
	None     PokemonType = "None"
)

func (ty *PokemonType) Scan(value interface{}) error {
	//fmt.Printf("value: %#v\n", value)
	*ty = PokemonType(value.([]byte))
	//fmt.Printf("value: %#v\n", *ty)
	return nil
}

func (ty PokemonType) Value() (driver.Value, error) {
	return string(ty), nil
}

func (ty PokemonType) String() string {
	return string(ty)
}

func (ty PokemonType) ConvertToDomain() value.PokemonType {
	return value.NewPokemonType(ty.String())
}

func Convert(t value.PokemonType) PokemonType {
	return PokemonType(t.EnglishName())
}

// ロール
type Role string

const (
	User                Role = "User"
	ReferenceableMaster Role = "ReferenceableMaster"
	Administrator       Role = "Administrator"
)

func (ro *Role) Scan(value interface{}) error {
	*ro = Role(value.([]byte))
	return nil
}

func (ro Role) Value() (driver.Value, error) {
	return string(ro), nil
}

// 性別
type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (ge *Gender) Scan(value interface{}) error {
	*ge = Gender(value.([]byte))
	return nil
}

func (ge Gender) Value() (driver.Value, error) {
	return string(ge), nil
}

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

func (na *Nature) Scan(value interface{}) error {
	*na = Nature(value.([]byte))
	return nil
}

func (na Nature) Value() (driver.Value, error) {
	return string(na), nil
}

// 対戦結果
type BattleResult string

const (
	Win  BattleResult = "Win"
	Lose BattleResult = "Lose"
	Draw BattleResult = "Draw"
)
