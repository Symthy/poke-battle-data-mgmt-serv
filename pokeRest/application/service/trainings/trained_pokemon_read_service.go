package trainings

import "github.com/Symthy/PokeRest/pokeRest/application/query"

type TrainedPokemonReadService struct {
	qs query.ITrainedPokemonQueryService
}

// UC: ユーザの育成済み個体一覧取得
func (s TrainedPokemonReadService) FindByUser(userId uint) {
}

// UC: 育成済み個体詳細取得
func (s TrainedPokemonReadService) FindById(id uint) {
}

// UC: パーティの育成済み個体取得
func (s TrainedPokemonReadService) FindByParty(partyId uint) {
}
