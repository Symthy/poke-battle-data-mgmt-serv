package trainings

import (
	"github.com/Symthy/PokeRest/internal/application/service"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type tas = trainings.TrainedPokemonAdjustments
type ta = trainings.TrainedPokemonAdjustment
type tai = identifier.TrainedAdjustmentId

type TrainedPokemonAdjustmentReadService struct {
	service.EntityAllFinder[tas, ta, tai, uint64]
	repo repository.ITrainedPokemonAdjustmentRepository
}

func NewTrainedPokemonAdjustmentReadService(
	repo repository.ITrainedPokemonAdjustmentRepository) TrainedPokemonAdjustmentReadService {
	serv := TrainedPokemonAdjustmentReadService{
		repo: repo,
	}
	serv.EntityAllFinder = service.NewEntityAllFinder[tas, ta, tai, uint64](repo)
	return serv
}

// UC: 育成済み調整全取得
// FindAll <- EntityAllFinder
