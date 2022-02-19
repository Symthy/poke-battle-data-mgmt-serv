package value

type PokemonType struct {
	englishName  typeNameEN
	japaneseName typeNameJP
}

func (t PokemonType) ResolveTypeName(lang string) string {
	if lang == "ja-JP" {
		return t.NameJP()
	}
	return t.NameEN()
}

func (t PokemonType) NameEN() string {
	return t.englishName.String()
}

func (t PokemonType) NameJP() string {
	return t.japaneseName.String()
}

func NewPokemonType(typeName string) PokemonType {
	for _, t := range PokemonTypeAll() {
		if t.NameEN() == typeName || t.NameJP() == typeName {
			return t
		}
	}
	return UnknownType()
}

func PokemonTypeAll() []PokemonType {
	return []PokemonType{
		Normal(),
		Fire(),
		Water(),
		Electric(),
		Grass(),
		Ice(),
		Fighting(),
		Poison(),
		Ground(),
		Flying(),
		Psychic(),
		Bug(),
		Rock(),
		Ghost(),
		Dragon(),
		Dark(),
		Steel(),
		Fairy(),
	}
}

func Normal() PokemonType      { return PokemonType{japaneseName: NormalJP, englishName: NormalEN} }
func Fire() PokemonType        { return PokemonType{japaneseName: FireJP, englishName: FireEN} }
func Water() PokemonType       { return PokemonType{japaneseName: WaterJP, englishName: WaterEN} }
func Electric() PokemonType    { return PokemonType{japaneseName: ElectricJP, englishName: ElectricEN} }
func Grass() PokemonType       { return PokemonType{japaneseName: GrassJP, englishName: GrassEN} }
func Ice() PokemonType         { return PokemonType{japaneseName: IceJP, englishName: IceEN} }
func Fighting() PokemonType    { return PokemonType{japaneseName: FightingJP, englishName: FightingEN} }
func Poison() PokemonType      { return PokemonType{japaneseName: PoisonJP, englishName: PoisonEN} }
func Ground() PokemonType      { return PokemonType{japaneseName: GroundJP, englishName: GroundEN} }
func Flying() PokemonType      { return PokemonType{japaneseName: FlyingJP, englishName: FlyingEN} }
func Psychic() PokemonType     { return PokemonType{japaneseName: PsychicJP, englishName: PsychicEN} }
func Bug() PokemonType         { return PokemonType{japaneseName: BugJP, englishName: BugEN} }
func Rock() PokemonType        { return PokemonType{japaneseName: RockJP, englishName: RockEN} }
func Ghost() PokemonType       { return PokemonType{japaneseName: GhostJP, englishName: GhostEN} }
func Dragon() PokemonType      { return PokemonType{japaneseName: DragonJP, englishName: DragonEN} }
func Dark() PokemonType        { return PokemonType{japaneseName: DarkJP, englishName: DarkEN} }
func Steel() PokemonType       { return PokemonType{japaneseName: SteelJP, englishName: SteelEN} }
func Fairy() PokemonType       { return PokemonType{japaneseName: FairyJP, englishName: FairyEN} }
func UnknownType() PokemonType { return PokemonType{japaneseName: NoneJP, englishName: NoneEN} }

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

	NormalEN   typeNameEN = "Normal"
	FireEN     typeNameEN = "Fire"
	WaterEN    typeNameEN = "Water"
	ElectricEN typeNameEN = "Electric"
	IceEN      typeNameEN = "Ice"
	GrassEN    typeNameEN = "Grass"
	FightingEN typeNameEN = "Fighting"
	PoisonEN   typeNameEN = "Poison"
	GroundEN   typeNameEN = "Ground"
	FlyingEN   typeNameEN = "Flying"
	PsychicEN  typeNameEN = "Psychic"
	BugEN      typeNameEN = "Bug"
	RockEN     typeNameEN = "Rock"
	GhostEN    typeNameEN = "Ghost"
	DragonEN   typeNameEN = "Dragon"
	DarkEN     typeNameEN = "Dark"
	SteelEN    typeNameEN = "Steel"
	FairyEN    typeNameEN = "Fairy"
	NoneEN     typeNameEN = ""
)
