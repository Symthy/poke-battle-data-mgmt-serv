package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type TrainedPokemonWriteService struct {
	trainedRepo repository.ITrainedPokemonRepository
}

func NewTrainedPokemonWriteService(
	trainedRepo repository.ITrainedPokemonRepository) TrainedPokemonWriteService {
	serv := TrainedPokemonWriteService{
		trainedRepo: trainedRepo,
	}
	return serv
}

// UC: 育成済み個体登録
func (s TrainedPokemonWriteService) SaveTrainedPokemon(
	cmd command.CreateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	trainedPokemon := cmd.ToDomain()

	created, err := s.trainedRepo.Create(trainedPokemon)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// UC: 育成済み個体更新
func (s TrainedPokemonWriteService) UpdateTrainedPokemon(cmd command.UpdateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	trainedPokemon := cmd.ToDomain()

	updated, err := s.trainedRepo.Update(trainedPokemon)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

// UC: 育成済み個体削除
func (s TrainedPokemonWriteService) DeleteTrainedPokemon(id uint) (*trainings.TrainedPokemon, error) {
	deleted, err := s.trainedRepo.Delete(id)
	return deleted, err
}
