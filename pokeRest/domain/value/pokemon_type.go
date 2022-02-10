package value

import "github.com/thoas/go-funk"

var (
	PokemonTypeNormal   = PokemonType{japaneseName: NormalJP, englishName: Normal}
	PokemonTypeFighting = PokemonType{japaneseName: FightingJP, englishName: Fighting}
	PokemonTypeFlying   = PokemonType{japaneseName: FlyingJP, englishName: Flying}
	PokemonTypePoison   = PokemonType{japaneseName: PoisonJP, englishName: Poison}
	PokemonTypeGround   = PokemonType{japaneseName: GroundJP, englishName: Ground}
	PokemonTypeRock     = PokemonType{japaneseName: RockJP, englishName: Rock}
	PokemonTypeBug      = PokemonType{japaneseName: BugJP, englishName: Bug}
	PokemonTypeGhost    = PokemonType{japaneseName: GhostJP, englishName: Ghost}
	PokemonTypeSteel    = PokemonType{japaneseName: SteelJP, englishName: Steel}
	PokemonTypeFire     = PokemonType{japaneseName: FireJP, englishName: Fire}
	PokemonTypeWater    = PokemonType{japaneseName: WaterJP, englishName: Water}
	PokemonTypeGrass    = PokemonType{japaneseName: GrassJP, englishName: Grass}
	PokemonTypeElectric = PokemonType{japaneseName: ElectricJP, englishName: Electric}
	PokemonTypePsychic  = PokemonType{japaneseName: PsychicJP, englishName: Psychic}
	PokemonTypeIce      = PokemonType{japaneseName: IceJP, englishName: Ice}
	PokemonTypeDragon   = PokemonType{japaneseName: DragonJP, englishName: Dragon}
	PokemonTypeDark     = PokemonType{japaneseName: DarkJP, englishName: Dark}
	PokemonTypeFairy    = PokemonType{japaneseName: FairyJP, englishName: Fairy}
	PokemonTypeUnknown  = PokemonType{japaneseName: NoneJP, englishName: None}
)

type PokemonType struct {
	englishName  typeNameEN
	japaneseName typeNameJP
}

func (t PokemonType) NameEN() string {
	return t.englishName.String()
}

func (t PokemonType) NameJP() string {
	return t.japaneseName.String()
}

func NewPokemonType(typeName string) PokemonType {
	return funk.Filter(GetPokemonTypes(), func(t PokemonType) bool {
		return t.NameEN() == typeName
	}).([]PokemonType)[0]
}

func GetPokemonTypes() []PokemonType {
	return []PokemonType{
		PokemonTypeNormal,
		PokemonTypeFire,
		PokemonTypeWater,
		PokemonTypeElectric,
		PokemonTypeGrass,
		PokemonTypeIce,
		PokemonTypeFighting,
		PokemonTypePoison,
		PokemonTypeGround,
		PokemonTypeFlying,
		PokemonTypePsychic,
		PokemonTypeGhost,
		PokemonTypeRock,
		PokemonTypeGhost,
		PokemonTypeDragon,
		PokemonTypeDark,
		PokemonTypeSteel,
		PokemonTypeFairy,
		PokemonTypeUnknown,
	}
}

type typeNameJP string

func (v typeNameJP) String() string {
	return string(v)
}

type typeNameEN string

func (v typeNameEN) String() string {
	return string(v)
}

const (
	NormalJP   typeNameJP = "ノーマル"
	FightingJP typeNameJP = "かくとう"
	FlyingJP   typeNameJP = "ひこう"
	PoisonJP   typeNameJP = "どく"
	GroundJP   typeNameJP = "じめん"
	RockJP     typeNameJP = "いわ"
	BugJP      typeNameJP = "むし"
	GhostJP    typeNameJP = "ゴースト"
	SteelJP    typeNameJP = "はがね"
	FireJP     typeNameJP = "ほのお"
	WaterJP    typeNameJP = "みず"
	GrassJP    typeNameJP = "くさ"
	ElectricJP typeNameJP = "でんき"
	PsychicJP  typeNameJP = "エスパー"
	IceJP      typeNameJP = "こおり"
	DragonJP   typeNameJP = "ドラゴン"
	DarkJP     typeNameJP = "あく"
	FairyJP    typeNameJP = "フェアリー"
	NoneJP     typeNameJP = "-"

	Normal   typeNameEN = "Normal"
	Fighting typeNameEN = "Fighting"
	Flying   typeNameEN = "Flying"
	Poison   typeNameEN = "Poison"
	Ground   typeNameEN = "Ground"
	Rock     typeNameEN = "Rock"
	Bug      typeNameEN = "Bug"
	Ghost    typeNameEN = "Ghost"
	Steel    typeNameEN = "Steel"
	Fire     typeNameEN = "Fire"
	Water    typeNameEN = "Water"
	Grass    typeNameEN = "Grass"
	Electric typeNameEN = "Electric"
	Psychic  typeNameEN = "Psychic"
	Ice      typeNameEN = "Ice"
	Dragon   typeNameEN = "Dragon"
	Dark     typeNameEN = "Dark"
	Fairy    typeNameEN = "Fairy"
	None     typeNameEN = "-"
)
