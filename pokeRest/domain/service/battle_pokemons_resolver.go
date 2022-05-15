package service

import (
	s_damages "github.com/Symthy/PokeRest/pokeRest/application/service/damages"
	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
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
	attackChan := make(chan AttackSideData, 1)
	defer close(attackChan)
	defenseChan := make(chan DefenseSideData, 1)
	defer close(defenseChan)
	errAttackChan := make(chan error, 1)
	defer close(errAttackChan)
	errDefenseChan := make(chan error, 1)

	go r.resolveAttackSide(adjustment, moveId, attackChan, errAttackChan)
	go r.resolveDefenseSide(adjustment, defenseChan, errDefenseChan)

	select {
	case err := <-errAttackChan:
		return nil, err
	case err := <-errDefenseChan:
		return nil, err
	default:
		attackSide := <-attackChan
		defenseSide := <-defenseChan
		battlePokemons := damages.NewPokemonBattleDataSet(
			attackSide.toAttackSidePokemon(),
			attackSide.toAttackSideBattleEffects(),
			defenseSide.toDefenseSidePokemon(),
			defenseSide.toDefenseSideBattleEffects(),
			attackSide.toAttackMove(),
			r.resolveTypeCompatibility(attackSide.MoveType(), defenseSide.PokemonTypes()))
		return &battlePokemons, nil
	}
}

func (r BattlePokemonResolver) resolveAttackSide(adjustment s_damages.BattlePokemonAdjustment,
	moveId uint16, resultChan chan<- AttackSideData, errChan chan<- error) {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId())
	if err != nil {
		errChan <- err
		return
	}
	effects, err := r.buildBattleEffects(adjustment)
	if err != nil {
		errChan <- err
		return
	}
	move, err := r.moveRepo.FindById(moveId)
	if err != nil {
		errChan <- err
		return
	}
	move.NotifyBattleEffects(effects)
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	attackSide := NewAttackSideData(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects, *move)
	resultChan <- attackSide
}

func (r BattlePokemonResolver) resolveDefenseSide(adjustment s_damages.BattlePokemonAdjustment,
	resultChan chan<- DefenseSideData, errChan chan<- error) {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId())
	if err != nil {
		errChan <- err
	}
	effects, err := r.buildBattleEffects(adjustment)
	if err != nil {
		errChan <- err
	}
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	defensePokemon := NewDefenseSideData(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects)
	resultChan <- defensePokemon
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
