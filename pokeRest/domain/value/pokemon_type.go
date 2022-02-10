package value

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
	for _, t := range GetPokemonTypes() {
		if t.NameEN() == typeName {
			return t
		}
	}
	return PokemonTypeUnknown
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
		PokemonTypeBug,
		PokemonTypeRock,
		PokemonTypeGhost,
		PokemonTypeDragon,
		PokemonTypeDark,
		PokemonTypeSteel,
		PokemonTypeFairy,
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
	FireJP     typeNameJP = "ほのお"
	WaterJP    typeNameJP = "みず"
	ElectricJP typeNameJP = "でんき"
	GrassJP    typeNameJP = "くさ"
	IceJP      typeNameJP = "こおり"
	FightingJP typeNameJP = "かくとう"
	PoisonJP   typeNameJP = "どく"
	GroundJP   typeNameJP = "じめん"
	FlyingJP   typeNameJP = "ひこう"
	PsychicJP  typeNameJP = "エスパー"
	BugJP      typeNameJP = "むし"
	RockJP     typeNameJP = "いわ"
	GhostJP    typeNameJP = "ゴースト"
	DragonJP   typeNameJP = "ドラゴン"
	DarkJP     typeNameJP = "あく"
	SteelJP    typeNameJP = "はがね"
	FairyJP    typeNameJP = "フェアリー"
	NoneJP     typeNameJP = ""

	Normal   typeNameEN = "Normal"
	Fire     typeNameEN = "Fire"
	Water    typeNameEN = "Water"
	Electric typeNameEN = "Electric"
	Ice      typeNameEN = "Ice"
	Grass    typeNameEN = "Grass"
	Fighting typeNameEN = "Fighting"
	Poison   typeNameEN = "Poison"
	Ground   typeNameEN = "Ground"
	Flying   typeNameEN = "Flying"
	Psychic  typeNameEN = "Psychic"
	Bug      typeNameEN = "Bug"
	Rock     typeNameEN = "Rock"
	Ghost    typeNameEN = "Ghost"
	Dragon   typeNameEN = "Dragon"
	Dark     typeNameEN = "Dark"
	Steel    typeNameEN = "Steel"
	Fairy    typeNameEN = "Fairy"
	None     typeNameEN = ""
)
