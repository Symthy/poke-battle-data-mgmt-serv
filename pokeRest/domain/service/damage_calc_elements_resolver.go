package service

import (
	s_damages "github.com/Symthy/PokeRest/pokeRest/application/service/damages"
	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/damages"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"golang.org/x/sync/errgroup"
)

type battleAttackSideResolver interface {
	resolve(adjustment *s_damages.BattlePokemonAdjustment, moveId identifier.MoveId, resultChan chan<- AttackSide) error
}

type battleDefenseSideResolver interface {
	resolve(adjustment *s_damages.BattlePokemonAdjustment, resultChan chan<- DefenseSide) error
}

type DamageCalcElementsService struct {
	attackSideResolver  battleAttackSideResolver
	defenseSideResolver battleDefenseSideResolver
	typeServ            types.TypeReadService
}

func NewDamageCalcPartialValuesFactory(effectsResolver *BattleEffectsResolver, pokemonRepo repository.IPokemonRepository,
	moveRepo repository.IMoveRepository) *DamageCalcElementsService {
	return &DamageCalcElementsService{
		attackSideResolver: attackSideResolver{
			effectsResolver: effectsResolver,
			pokemonRepo:     pokemonRepo,
			moveRepo:        moveRepo,
		},
		defenseSideResolver: defenseSideResolver{
			effectsResolver: effectsResolver,
			pokemonRepo:     pokemonRepo,
		},
		typeServ: types.NewTypeReadService(),
	}
}

func (r DamageCalcElementsService) Resolve(
	adjustment *s_damages.BattlePokemonAdjustment, moveId identifier.MoveId,
) (*damages.DamageCalcElements, error) {
	eg := new(errgroup.Group)

	attackChan := make(chan AttackSide, 1)
	defer close(attackChan)
	defenseChan := make(chan DefenseSide, 1)
	defer close(defenseChan)

	eg.Go(func() error {
		return r.attackSideResolver.resolve(adjustment, moveId, attackChan)
	})
	eg.Go(func() error {
		return r.defenseSideResolver.resolve(adjustment, defenseChan)
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	attackSide := <-attackChan
	defenseSide := <-defenseChan
	battlePokemons := damages.NewPokemonBattleDataSet(
		attackSide.toAttackSidePokemon(),
		defenseSide.toDefenseSidePokemon(),
		attackSide.toAttackMove(),
		r.resolveTypeCompatibility(attackSide.moveType, defenseSide.PokemonTypes()),
		attackSide.toAttackSideBattleEffects(),
		defenseSide.toDefenseSideBattleEffects())
	damageCalcElements := damages.NewDamageCalcElements(battlePokemons)
	return damageCalcElements, nil
}

func (r DamageCalcElementsService) resolveTypeCompatibility(
	attackMoveType value.PokemonType, defensePokemonType value.PokemonTypeSet) float32 {
	typeCompatibility := r.typeServ.GetAttackTypeCompatibility(attackMoveType.NameEN())
	return typeCompatibility.GetTotalDamageRate(defensePokemonType)
}

type attackSideResolver struct {
	effectsResolver *BattleEffectsResolver
	pokemonRepo     repository.IPokemonRepository
	moveRepo        repository.IMoveRepository
}

func (r attackSideResolver) resolve(adjustment *s_damages.BattlePokemonAdjustment,
	moveId identifier.MoveId, resultChan chan<- AttackSide) error {
	// Todo: receive type set
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId().Value())
	if err != nil {
		return err
	}
	effects, err := r.effectsResolver.resolve(adjustment)
	if err != nil {
		return err
	}
	move, err := r.moveRepo.FindById(moveId.Value())
	if err != nil {
		return err
	}
	move.NotifyBattleEffects(effects)
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	hasItem := !adjustment.HeldItemId().IsEmpty()
	attackSide := NewAttackSide(actualValues, pokemon.TypeSet(), adjustment.Nature(), effects, move, hasItem)
	resultChan <- attackSide
	return nil
}

type defenseSideResolver struct {
	effectsResolver *BattleEffectsResolver
	pokemonRepo     repository.IPokemonRepository
}

func (r defenseSideResolver) resolve(adjustment *s_damages.BattlePokemonAdjustment,
	resultChan chan<- DefenseSide) error {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId().Value())
	if err != nil {
		return err
	}
	effects, err := r.effectsResolver.resolve(adjustment)
	if err != nil {
		return err
	}
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	hasItem := !adjustment.HeldItemId().IsEmpty()
	defensePokemon := NewDefenseSide(actualValues, pokemon.TypeSet(), adjustment.Nature(), effects, hasItem)
	resultChan <- defensePokemon
	return nil
}
