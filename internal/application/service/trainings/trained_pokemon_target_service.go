package trainings

import (
	"github.com/Symthy/PokeRest/internal/application/service/trainings/command"
	"github.com/Symthy/PokeRest/internal/domain/entity/trainings"
	"github.com/Symthy/PokeRest/internal/domain/repository"
)

type TrainedPokemonTargetService struct {
	attackTargetRepo repository.ITrainedPokemonAttackRepository
}

// UC: 攻撃調整登録
func (s TrainedPokemonTargetService) AddAttackTarget(
	cmd command.AddAdjustmentAttackTargetCommand) (*trainings.TrainedPokemonAttackTarget, error) {
	domain, err := cmd.BuildDomain()
	if err != nil {
		return nil, err
	}
	created, err := s.attackTargetRepo.Create(domain)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// UC: 攻撃調整更新
func (s TrainedPokemonTargetService) UpdateAttackAdjustment() {
}

// UC: 攻撃調整削除
func (s TrainedPokemonTargetService) RemoveAttackAdjustment() {
}

// UC: 耐久調整登録
func (s TrainedPokemonTargetService) AddDefenseAdjustment() {
}

// UC: 耐久調整更新
func (s TrainedPokemonTargetService) UpdateDefenseAdjustment() {
}

// UC: 耐久調整削除
func (s TrainedPokemonTargetService) RemoveDefenseAdjustment() {
}
