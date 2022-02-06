package value

import "github.com/thoas/go-funk"

type PokemonType struct {
	name        TypeJapaneseName
	englishName TypeName
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
		{name: NormalJP, englishName: Normal},
		{name: FightingJP, englishName: Fighting},
		{name: FlyingJP, englishName: Flying},
		{name: PoisonJP, englishName: Poison},
		{name: GroundJP, englishName: Ground},
		{name: RockJP, englishName: Rock},
		{name: BugJP, englishName: Bug},
		{name: GhostJP, englishName: Ghost},
		{name: SteelJP, englishName: Steel},
		{name: FireJP, englishName: Fire},
		{name: WaterJP, englishName: Water},
		{name: GrassJP, englishName: Grass},
		{name: ElectricJP, englishName: Electric},
		{name: PsychicJP, englishName: Psychic},
		{name: IceJP, englishName: Ice},
		{name: DragonJP, englishName: Dragon},
		{name: DarkJP, englishName: Dark},
		{name: FairyJP, englishName: Fairy},
		{name: NoneJP, englishName: None},
	}
}

type TypeJapaneseName string

func (v TypeJapaneseName) String() string {
	return string(v)
}

type TypeName string

func (v TypeName) String() string {
	return string(v)
}

const (
	NormalJP   TypeJapaneseName = "ノーマル"
	FightingJP TypeJapaneseName = "かくとう"
	FlyingJP   TypeJapaneseName = "ひこう"
	PoisonJP   TypeJapaneseName = "どく"
	GroundJP   TypeJapaneseName = "じめん"
	RockJP     TypeJapaneseName = "いわ"
	BugJP      TypeJapaneseName = "むし"
	GhostJP    TypeJapaneseName = "ゴースト"
	SteelJP    TypeJapaneseName = "はがね"
	FireJP     TypeJapaneseName = "ほのお"
	WaterJP    TypeJapaneseName = "みず"
	GrassJP    TypeJapaneseName = "くさ"
	ElectricJP TypeJapaneseName = "でんき"
	PsychicJP  TypeJapaneseName = "エスパー"
	IceJP      TypeJapaneseName = "こおり"
	DragonJP   TypeJapaneseName = "ドラゴン"
	DarkJP     TypeJapaneseName = "あく"
	FairyJP    TypeJapaneseName = "フェアリー"
	NoneJP     TypeJapaneseName = "-"

	Normal   TypeName = "Normal"
	Fighting TypeName = "Fighting"
	Flying   TypeName = "Flying"
	Poison   TypeName = "Poison"
	Ground   TypeName = "Ground"
	Rock     TypeName = "Rock"
	Bug      TypeName = "Bug"
	Ghost    TypeName = "Ghost"
	Steel    TypeName = "Steel"
	Fire     TypeName = "Fire"
	Water    TypeName = "Water"
	Grass    TypeName = "Grass"
	Electric TypeName = "Electric"
	Psychic  TypeName = "Psychic"
	Ice      TypeName = "Ice"
	Dragon   TypeName = "Dragon"
	Dark     TypeName = "Dark"
	Fairy    TypeName = "Fairy"
	None     TypeName = "-"
)
