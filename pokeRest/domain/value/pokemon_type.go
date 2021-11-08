package value

import "github.com/thoas/go-funk"

type PokemonType struct {
	name        TypeJapaneseName
	englishName TypeEnglishName
}

func (t PokemonType) EnglishName() string {
	return t.englishName.String()
}

func (t PokemonType) Name() string {
	return t.name.String()
}

func NewPokemonType(typeName string) PokemonType {
	return funk.Filter(listPokemonTypes(), func(t PokemonType) bool {
		return t.EnglishName() == typeName
	}).([]PokemonType)[0]
}

func listPokemonTypes() []PokemonType {
	return []PokemonType{
		{name: Normal, englishName: NormalEN},
		{name: Fighting, englishName: FightingEN},
		{name: Flying, englishName: FlyingEN},
		{name: Poison, englishName: PoisonEN},
		{name: Ground, englishName: GroundEN},
		{name: Rock, englishName: RockEN},
		{name: Bug, englishName: BugEN},
		{name: Ghost, englishName: GhostEN},
		{name: Steel, englishName: SteelEN},
		{name: Fire, englishName: FireEN},
		{name: Water, englishName: WaterEN},
		{name: Grass, englishName: GrassEN},
		{name: Electric, englishName: ElectricEN},
		{name: Psychic, englishName: PsychicEN},
		{name: Ice, englishName: IceEN},
		{name: Dragon, englishName: DragonEN},
		{name: Dark, englishName: DarkEN},
		{name: Fairy, englishName: FairyEN},
		{name: None, englishName: NoneEN},
	}
}

type TypeJapaneseName string

func (v TypeJapaneseName) String() string {
	return string(v)
}

type TypeEnglishName string

func (v TypeEnglishName) String() string {
	return string(v)
}

const (
	Normal   TypeJapaneseName = "ノーマル"
	Fighting TypeJapaneseName = "かくとう"
	Flying   TypeJapaneseName = "ひこう"
	Poison   TypeJapaneseName = "どく"
	Ground   TypeJapaneseName = "じめん"
	Rock     TypeJapaneseName = "いわ"
	Bug      TypeJapaneseName = "むし"
	Ghost    TypeJapaneseName = "ゴースト"
	Steel    TypeJapaneseName = "はがね"
	Fire     TypeJapaneseName = "ほのお"
	Water    TypeJapaneseName = "みず"
	Grass    TypeJapaneseName = "くさ"
	Electric TypeJapaneseName = "でんき"
	Psychic  TypeJapaneseName = "エスパー"
	Ice      TypeJapaneseName = "こおり"
	Dragon   TypeJapaneseName = "ドラゴン"
	Dark     TypeJapaneseName = "あく"
	Fairy    TypeJapaneseName = "フェアリー"
	None     TypeJapaneseName = "-"

	NormalEN   TypeEnglishName = "Normal"
	FightingEN TypeEnglishName = "Fighting"
	FlyingEN   TypeEnglishName = "Flying"
	PoisonEN   TypeEnglishName = "Poison"
	GroundEN   TypeEnglishName = "Ground"
	RockEN     TypeEnglishName = "Rock"
	BugEN      TypeEnglishName = "Bug"
	GhostEN    TypeEnglishName = "Ghost"
	SteelEN    TypeEnglishName = "Steel"
	FireEN     TypeEnglishName = "Fire"
	WaterEN    TypeEnglishName = "Water"
	GrassEN    TypeEnglishName = "Grass"
	ElectricEN TypeEnglishName = "Electric"
	PsychicEN  TypeEnglishName = "Psychic"
	IceEN      TypeEnglishName = "Ice"
	DragonEN   TypeEnglishName = "Dragon"
	DarkEN     TypeEnglishName = "Dark"
	FairyEN    TypeEnglishName = "Fairy"
	NoneEN     TypeEnglishName = "-"
)
