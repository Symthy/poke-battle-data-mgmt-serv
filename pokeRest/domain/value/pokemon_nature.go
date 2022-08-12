package value

// 性格
type PokemonNature struct {
	englishName    NatureNameEN
	japaneseName   NatureNameJP
	riseCorrection *NatureCorrection
	downCorrection *NatureCorrection
}

func NewPokemonNature(natureNameEN string) PokemonNature {
	return PokemonNatureAll()[NatureNameEN(natureNameEN)]
}

func (n PokemonNature) ApplyCorrection(actualValues PokemonActualValues) int {
	// Todo
	return 0
}

func (n PokemonNature) String() string {
	return string(n.englishName)
}

func PokemonNatureAll() map[NatureNameEN]PokemonNature {
	return map[NatureNameEN]PokemonNature{
		LonelyEN:  NatureLonely(),
		AdamantEN: NatureAdamant(),
		NaughtyEN: NatureNaughty(),
		BraveEN:   NatureBrave(),

		BoldEN:    NatureBold(),
		ImpishEN:  NatureImpish(),
		LaxEN:     NatureLax(),
		RelaxedEN: NatureRelaxed(),

		ModestEN: NatureModest(),
		MildEN:   NatureMild(),
		RashEN:   NatureRash(),
		QuietEN:  NatureQuiet(),

		CalmEN:    NatureCalm(),
		GentleEN:  NatureGentle(),
		CarefulEN: NatureCareful(),
		SassyEN:   NatureSassy(),

		TimidEN: NatureTimid(),
		HastyEN: NatureHasty(),
		JollyEN: NatureJolly(),
		NaiveEN: NatureNaive(),

		HardyEN:   NatureHardy(),
		DocileEN:  NatureDocile(),
		SeriousEN: NatureSerious(),
		BashfulEN: NatureBashful(),
		QuirkyEN:  NatureQuirky(),
	}
}

func NatureLonely() PokemonNature { // さみしがりや（A↑B↓）
	return PokemonNature{
		englishName:    LonelyEN,
		japaneseName:   LonelyJP,
		riseCorrection: &NatureARise,
		downCorrection: &NatureBDown,
	}
}
func NatureAdamant() PokemonNature { // いじっぱり（A↑C↓）
	return PokemonNature{
		englishName:    AdamantEN,
		japaneseName:   AdamantJP,
		riseCorrection: &NatureARise,
		downCorrection: &NatureCDown,
	}
}
func NatureNaughty() PokemonNature { // やんちゃ（A↑D↓）
	return PokemonNature{
		englishName:    NaughtyEN,
		japaneseName:   NaughtyJP,
		riseCorrection: &NatureARise,
		downCorrection: &NatureDDown,
	}
}
func NatureBrave() PokemonNature { // ゆうかん（A↑S↓）
	return PokemonNature{
		englishName:    BraveEN,
		japaneseName:   BraveJP,
		riseCorrection: &NatureARise,
		downCorrection: &NatureSDown,
	}
}

func NatureBold() PokemonNature { // ずぶとい（B↑A↓）
	return PokemonNature{
		englishName:    BoldEN,
		japaneseName:   BoldJP,
		riseCorrection: &NatureBRise,
		downCorrection: &NatureADown,
	}
}
func NatureImpish() PokemonNature { // わんぱく（B↑C↓）
	return PokemonNature{
		englishName:    ImpishEN,
		japaneseName:   ImpishJP,
		riseCorrection: &NatureBRise,
		downCorrection: &NatureCDown,
	}
}
func NatureLax() PokemonNature { // のうてんき（B↑D↓）
	return PokemonNature{
		englishName:    LaxEN,
		japaneseName:   LaxJP,
		riseCorrection: &NatureBRise,
		downCorrection: &NatureDDown,
	}
}
func NatureRelaxed() PokemonNature { // のんき（B↑S↓）
	return PokemonNature{
		englishName:    RelaxedEN,
		japaneseName:   RelaxedJP,
		riseCorrection: &NatureBRise,
		downCorrection: &NatureSDown,
	}
}

func NatureModest() PokemonNature { // ひかえめ（C↑A↓）
	return PokemonNature{
		englishName:    ModestEN,
		japaneseName:   ModestJP,
		riseCorrection: &NatureCRise,
		downCorrection: &NatureADown,
	}
}
func NatureMild() PokemonNature { // おっとり（C↑B↓）
	return PokemonNature{
		englishName:    MildEN,
		japaneseName:   MildJP,
		riseCorrection: &NatureCRise,
		downCorrection: &NatureBDown,
	}
}
func NatureRash() PokemonNature { // うっかりや（C↑D↓）
	return PokemonNature{
		englishName:    RashEN,
		japaneseName:   RashJP,
		riseCorrection: &NatureCRise,
		downCorrection: &NatureDDown,
	}
}
func NatureQuiet() PokemonNature { // れいせい（C↑S↓）
	return PokemonNature{
		englishName:    QuietEN,
		japaneseName:   QuietJP,
		riseCorrection: &NatureCRise,
		downCorrection: &NatureSDown,
	}
}

func NatureCalm() PokemonNature { // おだやか（D↑A↓）
	return PokemonNature{
		englishName:    CalmEN,
		japaneseName:   CalmJP,
		riseCorrection: &NatureDRise,
		downCorrection: &NatureADown,
	}
}
func NatureGentle() PokemonNature { // おとなしい（D↑B↓）
	return PokemonNature{
		englishName:    GentleEN,
		japaneseName:   GentleJP,
		riseCorrection: &NatureDRise,
		downCorrection: &NatureBDown,
	}
}
func NatureCareful() PokemonNature { // しんちょう（D↑C↓）
	return PokemonNature{
		englishName:    CarefulEN,
		japaneseName:   CarefulJP,
		riseCorrection: &NatureDRise,
		downCorrection: &NatureCDown,
	}
}
func NatureSassy() PokemonNature { // なまいき（D↑S↓）
	return PokemonNature{
		englishName:    SassyEN,
		japaneseName:   SassyJP,
		riseCorrection: &NatureDRise,
		downCorrection: &NatureSDown,
	}
}

func NatureTimid() PokemonNature { // おくびょう（S↑A↓）
	return PokemonNature{
		englishName:    TimidEN,
		japaneseName:   TimidJP,
		riseCorrection: &NatureSRise,
		downCorrection: &NatureADown,
	}
}
func NatureHasty() PokemonNature { // せっかち（S↑B↓）
	return PokemonNature{
		englishName:    HastyEN,
		japaneseName:   HastyJP,
		riseCorrection: &NatureSRise,
		downCorrection: &NatureBDown,
	}
}
func NatureJolly() PokemonNature { // ようき（S↑C↓）
	return PokemonNature{
		englishName:    HastyEN,
		japaneseName:   HastyJP,
		riseCorrection: &NatureSRise,
		downCorrection: &NatureCDown,
	}
}
func NatureNaive() PokemonNature { // むじゃき（S↑D↓）
	return PokemonNature{
		englishName:    NaiveEN,
		japaneseName:   NaiveJP,
		riseCorrection: &NatureSRise,
		downCorrection: &NatureDDown,
	}
}

func NatureHardy() PokemonNature { // がんばりや（補正無し）
	return PokemonNature{
		englishName:    HardyEN,
		japaneseName:   HardyJP,
		riseCorrection: nil,
		downCorrection: nil,
	}
}
func NatureDocile() PokemonNature { // すなお（補正無し）
	return PokemonNature{
		englishName:    DocileEN,
		japaneseName:   DocileJP,
		riseCorrection: nil,
		downCorrection: nil,
	}
}
func NatureSerious() PokemonNature { // まじめ（補正無し）
	return PokemonNature{
		englishName:    SeriousEN,
		japaneseName:   SeriousJP,
		riseCorrection: nil,
		downCorrection: nil,
	}
}
func NatureBashful() PokemonNature { // てれや（補正無し）
	return PokemonNature{
		englishName:    BashfulEN,
		japaneseName:   BashfulJP,
		riseCorrection: nil,
		downCorrection: nil,
	}
}
func NatureQuirky() PokemonNature { // きまぐれ（補正無し）
	return PokemonNature{
		englishName:    QuirkyEN,
		japaneseName:   QuirkyJP,
		riseCorrection: nil,
		downCorrection: nil,
	}
}

type NatureNameEN string

const (
	LonelyEN  NatureNameEN = "Lonely"  // さみしがりや（A↑B↓）
	AdamantEN NatureNameEN = "Adamant" // いじっぱり（A↑C↓）
	NaughtyEN NatureNameEN = "Naughty" // やんちゃ（A↑D↓）
	BraveEN   NatureNameEN = "Brave"   // ゆうかん（A↑S↓）

	BoldEN    NatureNameEN = "Bold"    // ずぶとい（B↑A↓）
	ImpishEN  NatureNameEN = "Impish"  // わんぱく（B↑C↓）
	LaxEN     NatureNameEN = "Lax"     // のうてんき（B↑D↓）
	RelaxedEN NatureNameEN = "Relaxed" // のんき（B↑S↓）

	ModestEN NatureNameEN = "Modest" // ひかえめ（C↑A↓）
	MildEN   NatureNameEN = "Mild"   // おっとり（C↑B↓）
	RashEN   NatureNameEN = "Rash"   // うっかりや（C↑D↓）
	QuietEN  NatureNameEN = "Quiet"  // れいせい（C↑S↓）

	CalmEN    NatureNameEN = "Calm"    // おだやか（D↑A↓）
	GentleEN  NatureNameEN = "Gentle"  // おとなしい（D↑B↓）
	CarefulEN NatureNameEN = "Careful" // しんちょう（D↑C↓）
	SassyEN   NatureNameEN = "Sassy"   // なまいき（D↑S↓）

	TimidEN NatureNameEN = "Timid" // おくびょう（S↑A↓）
	HastyEN NatureNameEN = "Hasty" // せっかち（S↑B↓）
	JollyEN NatureNameEN = "Jolly" // ようき（S↑C↓）
	NaiveEN NatureNameEN = "Naive" // むじゃき（S↑D↓）

	HardyEN   NatureNameEN = "Hardy"   // がんばりや（補正無し）
	DocileEN  NatureNameEN = "Docile"  // すなお（補正無し）
	SeriousEN NatureNameEN = "Serious" // まじめ（補正無し）
	BashfulEN NatureNameEN = "Bashful" // てれや（補正無し）
	QuirkyEN  NatureNameEN = "Quirky"  // きまぐれ（補正無し）
)

type NatureNameJP string

const (
	LonelyJP  NatureNameJP = "さみしがりや"
	AdamantJP NatureNameJP = "いじっぱり"
	NaughtyJP NatureNameJP = "やんちゃ"
	BraveJP   NatureNameJP = "ゆうかん"

	BoldJP    NatureNameJP = "ずぶとい"
	ImpishJP  NatureNameJP = "わんぱく"
	LaxJP     NatureNameJP = "のうてんき"
	RelaxedJP NatureNameJP = "のんき"

	ModestJP NatureNameJP = "ひかえめ"
	MildJP   NatureNameJP = "おっとり"
	RashJP   NatureNameJP = "うっかりや"
	QuietJP  NatureNameJP = "れいせい"

	CalmJP    NatureNameJP = "おだやか"
	GentleJP  NatureNameJP = "おとなしい"
	CarefulJP NatureNameJP = "しんちょう"
	SassyJP   NatureNameJP = "なまいき"

	TimidJP NatureNameJP = "おくびょう"
	HastyJP NatureNameJP = "せっかち"
	JollyJP NatureNameJP = "ようき"
	NaiveJP NatureNameJP = "むじゃき"

	HardyJP   NatureNameJP = "がんばりや"
	DocileJP  NatureNameJP = "すなお"
	SeriousJP NatureNameJP = "まじめ"
	BashfulJP NatureNameJP = "てれや"
	QuirkyJP  NatureNameJP = "きまぐれ"
)
