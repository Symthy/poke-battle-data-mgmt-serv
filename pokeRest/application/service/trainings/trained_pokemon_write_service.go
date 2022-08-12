package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings/command"
	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings/spec"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type TrainedPokemonWriteService struct {
	trainedRepo repository.ITrainedPokemonRepository
	spec        spec.TrainedPokemonSpecification
}

func NewTrainedPokemonWriteService(
	trainedRepo repository.ITrainedPokemonRepository, userRepo repository.IUserRepository,
) TrainedPokemonWriteService {
	serv := TrainedPokemonWriteService{
		trainedRepo: trainedRepo,
		spec:        spec.NewTrainedPokemonSpecification(userRepo),
	}
	return serv
}

// UC: 育成済み個体登録
func (s TrainedPokemonWriteService) SaveTrainedPokemon(
	cmd command.CreateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if err := s.spec.ValidateForCreate(domain); err != nil {
		return nil, err
	}
	created, err := s.trainedRepo.Create(domain)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// UC: 育成済み個体更新
func (s TrainedPokemonWriteService) UpdateTrainedPokemon(
	cmd command.UpdateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	trainedPokemon, err := cmd.BuildDomain()
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if err := s.spec.ValidateForUpdate(domain); err != nil {
		return nil, err
	}
	updated, err := s.trainedRepo.Update(trainedPokemon)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

// UC: 育成済み個体削除
func (s TrainedPokemonWriteService) DeleteTrainedPokemon(
	cmd command.DeleteTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	if err := s.spec.ValidateForUpdate(domain); err != nil {
		return nil, err
	}
	return s.trainedRepo.Delete(domain.Id().Value())
}
