package pokemons

import "github.com/Symthy/PokeRest/pokeRest/domain/repository"

type TrainedPokemonWriteService struct {
	tr  repository.ITrainedPokemonRepository
	tbr repository.ITrainedPokemonBaseRepository
}

// 育成済み個体登録
func (s TrainedPokemonWriteService) SaveTrainedPokemon() {
}

// 育成済み個体更新
func (s TrainedPokemonWriteService) UpdateTrainedPokemon() {
}

// 育成済み個体削除
func (s TrainedPokemonWriteService) DeleteTrainedPokemon() {
}
