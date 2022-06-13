package service

import (
	s_damages "github.com/Symthy/PokeRest/pokeRest/application/service/damages"
	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"golang.org/x/sync/errgroup"
)

type BattlePokemonResolver struct {
	pokemonRepo  repository.IPokemonRepository
	abilityRepo  repository.IAbilityRepository
	heldItemRepo repository.IHeldItemRepository
	moveRepo     repository.IMoveRepository
	typeServ     types.TypeReadService
}

// Todo: use WaitGroup
func (r BattlePokemonResolver) Resolve(
	adjustment s_damages.BattlePokemonAdjustment, moveId uint16,
) (*damages.PokemonBattleDataSet, error) {
	eg := new(errgroup.Group)

	attackChan := make(chan AttackSide, 1)
	defer close(attackChan)
	defenseChan := make(chan DefenseSide, 1)
	defer close(defenseChan)

	eg.Go(func() error {
		return r.resolveAttackSide(adjustment, moveId, attackChan)
	})
	eg.Go(func() error {
		return r.resolveDefenseSide(adjustment, defenseChan)
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	attackSide := <-attackChan
	defenseSide := <-defenseChan
	battlePokemons := damages.NewPokemonBattleDataSet(
		attackSide.toAttackSidePokemon(),
		attackSide.toAttackSideBattleEffects(),
		defenseSide.toDefenseSidePokemon(),
		defenseSide.toDefenseSideBattleEffects(),
		attackSide.toAttackMove(),
		r.resolveTypeCompatibility(attackSide.MoveType(), defenseSide.PokemonTypes()))
	return battlePokemons, nil
}

func (r BattlePokemonResolver) resolveAttackSide(adjustment s_damages.BattlePokemonAdjustment,
	moveId uint16, resultChan chan<- AttackSide) error {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId())
	if err != nil {
		return err
	}
	effects, err := r.buildBattleEffects(adjustment)
	if err != nil {
		return err
	}
	move, err := r.moveRepo.FindById(moveId)
	if err != nil {
		return err
	}
	move.NotifyBattleEffects(effects)
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	attackSide := NewAttackSide(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects, *move)
	resultChan <- attackSide
	return nil
}

func (r BattlePokemonResolver) resolveDefenseSide(adjustment s_damages.BattlePokemonAdjustment,
	resultChan chan<- DefenseSide) error {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId())
	if err != nil {
		return err
	}
	effects, err := r.buildBattleEffects(adjustment)
	if err != nil {
		return err
	}
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	defensePokemon := NewDefenseSide(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects)
	resultChan <- defensePokemon
	return nil
}

func (r BattlePokemonResolver) buildBattleEffects(
	adjustment s_damages.BattlePokemonAdjustment) (*value.BattleEffects, error) {
	effects := &value.BattleEffects{}
	ability, err := r.abilityRepo.FindById(adjustment.AbilityId())
	if err != nil {
		return nil, err
	}
	ability.NotifyBattleEffects(effects)
	item, err := r.heldItemRepo.FindById(adjustment.HeldItemId())
	if err != nil {
		return nil, err
	}
	item.NotifyBattleEffects(effects)
	return effects, err
}

func (r BattlePokemonResolver) resolveTypeCompatibility(
	attackMoveType value.PokemonType, defensePokemonType value.PokemonTypeSet) float32 {
	typeCompatibility := r.typeServ.GetAttackTypeCompatibility(attackMoveType.NameEN())
	return typeCompatibility.GetTotalDamageRate(defensePokemonType)
}
