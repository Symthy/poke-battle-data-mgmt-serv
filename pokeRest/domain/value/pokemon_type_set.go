package value

type PokemonTypeSet struct {
	firstType  PokemonType
	secondType PokemonType
}

func NewPokemonTypeSet(firstType, secondType PokemonType) PokemonTypeSet {
	return PokemonTypeSet{}
}

func (t PokemonTypeSet) GetTypeNames(lang string) (string, string) {
	return t.firstType.ResolveTypeName(lang), t.secondType.ResolveTypeName(lang)
}

func (t PokemonTypeSet) GetTypeEnglishNames() (string, string) {
	return t.firstType.NameEN(), t.secondType.NameEN()
}
