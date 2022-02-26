package trainings

import (
	"fmt"

	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/service"
)

type trainedPokemonTransactionalRepository = func(repository.ITrainedPokemonRepository) repository.ITrainedPokemonTransactionalRepository

type TrainedPokemonWriteService struct {
	trainedParamRepo               repository.ITrainedPokemonRepository
	adjustmentRepo                 repository.ITrainedPokemonAdjustmentRepository
	trainedPokemonResolver         service.TrainedPokemonResolver
	transactionalRepositoryBuilder trainedPokemonTransactionalRepository
}

func NewTrainedPokemonWriteService(
	trainedParamRepo repository.ITrainedPokemonRepository,
	adjustmentRepo repository.ITrainedPokemonAdjustmentRepository,
	transactionalRepositoryBuilder trainedPokemonTransactionalRepository) TrainedPokemonWriteService {

	serv := TrainedPokemonWriteService{
		trainedParamRepo:               trainedParamRepo,
		adjustmentRepo:                 adjustmentRepo,
		trainedPokemonResolver:         service.NewTrainedPokemonResolver(adjustmentRepo),
		transactionalRepositoryBuilder: transactionalRepositoryBuilder,
	}
	return serv
}

// UC: 育成済み個体登録
func (s TrainedPokemonWriteService) SaveTrainedPokemon(
	cmd command.CreateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	inputTrainedPokemon := cmd.ToDomain()
	transactionalRepo := s.transactionalRepositoryBuilder(s.trainedParamRepo)

	if err := transactionalRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer transactionalRepo.PanicPostProcess()

	trainedPokemon, err := s.trainedPokemonResolver.Resolve(inputTrainedPokemon)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}
	createdParam, err := transactionalRepo.Create(trainedPokemon.TrainedPokemonParam)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}
	trainedPokemon.TrainedPokemonParam = *createdParam

	if err := transactionalRepo.FinishTransaction(); err != nil {
		return nil, err
	}
	return trainedPokemon, nil
}

// UC: 育成済み個体更新
func (s TrainedPokemonWriteService) UpdateTrainedPokemon(cmd command.UpdateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	inputTrainedPokemon := cmd.ToDomain()
	transactionalRepo := s.transactionalRepositoryBuilder(s.trainedParamRepo)

	if err := transactionalRepo.StartTransaction(); err != nil {
		return nil, err
	}
	defer transactionalRepo.PanicPostProcess()

	ret, err := transactionalRepo.FindById(inputTrainedPokemon.Id())
	if err != nil {
		return nil, err
	}
	if ret == nil {
		// Todo: replace error
		return nil, fmt.Errorf("trained pokemon not found")
	}

	trainedPokemon, err := s.trainedPokemonResolver.Resolve(inputTrainedPokemon)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}
	updatedParam, err := transactionalRepo.Update(trainedPokemon.TrainedPokemonParam)
	if err != nil {
		transactionalRepo.CancelTransaction()
		return nil, err
	}
	trainedPokemon.TrainedPokemonParam = *updatedParam

	if err := transactionalRepo.FinishTransaction(); err != nil {
		return nil, err
	}
	return trainedPokemon, nil
}

// UC: 育成済み個体削除
func (s TrainedPokemonWriteService) DeleteTrainedPokemon(id uint) (*trainings.TrainedPokemon, error) {
	param, err := s.trainedParamRepo.Delete(id)
	// Todo:
	// delete used trained pokemon adjustment if ref num is zero
	isRefZero := false
	if isRefZero {
		_, err := s.adjustmentRepo.Delete(param.AdjustmentId().Value())
		if err != nil {
			return nil, err
		}
	}
	trainedPoke := trainings.NewTrainedPokemon(*param, trainings.TrainedPokemonAdjustment{})
	return &trainedPoke, err
}
