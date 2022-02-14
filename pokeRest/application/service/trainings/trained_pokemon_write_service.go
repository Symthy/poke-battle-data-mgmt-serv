package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service/trainings/command"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type TrainedPokemonWriteService struct {
	trainedParamRepo repository.ITrainedPokemonRepository
	adjustmentRepo   repository.ITrainedPokemonAdjustmentRepository
}

func NewTrainedPokemonWriteService(
	trainedParamRepo repository.ITrainedPokemonRepository,
	adjustmentRepo repository.ITrainedPokemonAdjustmentRepository) TrainedPokemonWriteService {
	serv := TrainedPokemonWriteService{
		trainedParamRepo: trainedParamRepo,
		adjustmentRepo:   adjustmentRepo,
	}
	return serv
}

// UC: 育成済み個体登録
func (s TrainedPokemonWriteService) SaveTrainedPokemon(cmd command.CreateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	trainedPokemonEntity := cmd.ToDomain()
	// Todo: transaction
	adjustment, err := s.adjustmentRepo.Find(trainedPokemonEntity.TrainedPokemonAdjustment)
	if err != nil {
		return nil, err
	}
	param := trainedPokemonEntity.TrainedPokemonParam
	if adjustment == nil {
		createdAdjustment, err := s.adjustmentRepo.Create(trainedPokemonEntity.TrainedPokemonAdjustment)
		if err != nil {
			return nil, err
		}
		param = param.ApplyAdjustmentId(createdAdjustment.Id())
		adjustment = createdAdjustment
	}
	createdParam, err := s.trainedParamRepo.Create(param)
	if err != nil {
		return nil, err
	}
	trainedPoke := trainings.NewTrainedPokemon(*createdParam, *adjustment)
	return &trainedPoke, nil
}

// UC: 育成済み個体更新
func (s TrainedPokemonWriteService) UpdateTrainedPokemon(cmd command.UpdateTrainedPokemonCommand) (*trainings.TrainedPokemon, error) {
	trainedPokemonEntity := cmd.ToDomain()
	// Todo: transaction
	// Todo: refactor: find Adjustment id base
	adjustment, err := s.adjustmentRepo.Find(trainedPokemonEntity.TrainedPokemonAdjustment)
	if err != nil {
		return nil, err
	}
	param := trainedPokemonEntity.TrainedPokemonParam
	if adjustment == nil {
		createdAdjustment, err := s.adjustmentRepo.Create(trainedPokemonEntity.TrainedPokemonAdjustment)
		if err != nil {
			return nil, err
		}
		param = param.ApplyAdjustmentId(createdAdjustment.Id())
		adjustment = createdAdjustment
	}
	updatedParam, err := s.trainedParamRepo.Create(param)
	if err != nil {
		return nil, err
	}
	trainedPoke := trainings.NewTrainedPokemon(*updatedParam, *adjustment)
	return &trainedPoke, nil
}

// UC: 育成済み個体削除
func (s TrainedPokemonWriteService) DeleteTrainedPokemon(id uint) (*trainings.TrainedPokemonParam, error) {
	param, err := s.trainedParamRepo.Delete(id)
	// Todo:
	// delete used trained pokemon adjustment if ref num is zero
	isRefZero := false
	if isRefZero {
		_, err := s.adjustmentRepo.Delete(*param.AdjustmentId())
		if err != nil {
			return nil, err
		}
	}
	return param, err
}
