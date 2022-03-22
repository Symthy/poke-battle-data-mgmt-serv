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

// Todo: pararel valid?
func (r BattlePokemonResolver) Resolve(
	adjustment s_damages.BattlePokemonAdjustment, moveId uint,
) (*damages.BattlePokemons, error) {
	attackChan := make(chan damages.BattleAttackPokemon, 1)
	defer close(attackChan)
	defenceChan := make(chan damages.BattleDefencePokemon, 1)
	defer close(defenceChan)
	errAttackChan := make(chan error, 1)
	defer close(errAttackChan)
	errDefenceChan := make(chan error, 1)

	go r.resolveAttackSide(adjustment, moveId, attackChan, errAttackChan)
	go r.resolveDefenceSide(adjustment, defenceChan, errDefenceChan)

	select {
	case err := <-errAttackChan:
		return nil, err
	case err := <-errDefenceChan:
		return nil, err
	default:
		attackSide := <-attackChan
		defenceSide := <-defenceChan
		battlePokemons := damages.NewBattlePokemons(attackSide, defenceSide,
			r.resolveTypeCompatibility(attackSide.MoveType(), defenceSide.PokemonTypes()))
		return &battlePokemons, nil
	}

}

func (r BattlePokemonResolver) resolveAttackSide(adjustment s_damages.BattlePokemonAdjustment,
	moveId uint, resultChan chan<- damages.BattleAttackPokemon, errChan chan<- error) {
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
	attackPokemon := damages.NewBattleAttackPokemon(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects, *move)
	resultChan <- attackPokemon
}

func (r BattlePokemonResolver) resolveDefenceSide(adjustment s_damages.BattlePokemonAdjustment,
	resultChan chan<- damages.BattleDefencePokemon, errChan chan<- error) {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId())
	if err != nil {
		errChan <- err
	}
	effects, err := r.buildBattleEffects(adjustment)
	if err != nil {
		errChan <- err
	}
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	defencePokemon := damages.NewBattleDefencePokemon(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects)
	resultChan <- defencePokemon
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
	attackMoveType value.PokemonType, defencePokemonType value.PokemonTypeSet) float32 {
	typeCompatibility := r.typeServ.GetAttackTypeCompatibility(attackMoveType.NameEN())
	return typeCompatibility.GetTotalDamageRate(defencePokemonType)
}
