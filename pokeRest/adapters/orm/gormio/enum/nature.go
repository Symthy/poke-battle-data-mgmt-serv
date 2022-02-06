package enum

import (
	"database/sql/driver"
)

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
