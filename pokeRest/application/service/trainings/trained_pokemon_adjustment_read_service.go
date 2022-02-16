package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type tas = trainings.TrainedPokemonAdjustments
type ta = trainings.TrainedPokemonAdjustment

type TrainedPokemonAdjustmentReadService struct {
	service.EntityAllFinder[tas, ta]
	repo repository.ITrainedPokemonAdjustmentRepository
}

func NewTrainedPokemonAdjustmentReadService(
	repo repository.ITrainedPokemonAdjustmentRepository) TrainedPokemonAdjustmentReadService {
	serv := TrainedPokemonAdjustmentReadService{
		repo: repo,
	}
	serv.EntityAllFinder = service.NewEntityAllFinder[tas, ta](repo)
	return serv
}

// UC: 育成済み調整全取得
// FindAll <- EntityAllFinder
