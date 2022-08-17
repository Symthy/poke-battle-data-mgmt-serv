package trainings

import "github.com/Symthy/PokeRest/internal/application/query"

type TrainedPokemonMultiSearchService struct {
	qs query.ITrainedPokemonQueryService
}

// UC: 育成済み個体検索
func (s TrainedPokemonMultiSearchService) SearchTrainedPokemons() {
}
