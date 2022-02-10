package trainings

import "github.com/Symthy/PokeRest/pokeRest/application/query"

type TrainedPokemonReadService struct {
	qs query.ITrainedPokemonQueryService
}

// UC: ユーザの育成済み個体一覧取得
func (s TrainedPokemonReadService) FindTrainedPokemonsByUser(userId uint) {
}

// UC: 育成済み個体詳細取得
func (s TrainedPokemonReadService) FindTrainedPokemonById(id uint) {
}

// UC: 育成済み個体検索
func (s TrainedPokemonReadService) SearchTrainedPokemons() {
}

// UC: パーティの育成済み個体取得
func (s TrainedPokemonReadService) FindTrainedPokemonsByParty(partyId uint) {
}
