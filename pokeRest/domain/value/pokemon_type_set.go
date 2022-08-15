package value

type PokemonTypeSet struct {
	first  PokemonType
	second PokemonType
}

func NewPokemonTypeSet(firstType, secondType PokemonType) *PokemonTypeSet {
	return &PokemonTypeSet{
		first:  firstType,
		second: secondType,
	}
}

func (t PokemonTypeSet) GetTypeNames(lang string) (string, string) {
	return t.first.ResolveTypeName(lang), t.second.ResolveTypeName(lang)
}

func (t PokemonTypeSet) GetTypeEnglishNames() (string, string) {
	return t.first.NameEN(), t.second.NameEN()
}

func (t PokemonTypeSet) FirstType() PokemonType {
	return t.first
}

func (t PokemonTypeSet) SecondType() PokemonType {
	return t.second
}
