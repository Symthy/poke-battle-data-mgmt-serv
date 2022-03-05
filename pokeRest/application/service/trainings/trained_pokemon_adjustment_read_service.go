package trainings

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type tas = trainings.TrainedPokemonAdjustments
type ta = trainings.TrainedPokemonAdjustment
type tai = identifier.TrainedAdjustmentId

type TrainedPokemonAdjustmentReadService struct {
	service.EntityAllFinder[tas, ta, tai]
	repo repository.ITrainedPokemonAdjustmentRepository
}

func NewTrainedPokemonAdjustmentReadService(
	repo repository.ITrainedPokemonAdjustmentRepository) TrainedPokemonAdjustmentReadService {
	serv := TrainedPokemonAdjustmentReadService{
		repo: repo,
	}
	serv.EntityAllFinder = service.NewEntityAllFinder[tas, ta, tai](repo)
	return serv
}

// UC: 育成済み調整全取得
// FindAll <- EntityAllFinder
