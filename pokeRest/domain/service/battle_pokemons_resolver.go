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

type BattlePokemonResolver struct {
	IBattleEffectsResolver
	pokemonRepo repository.IPokemonRepository
	moveRepo    repository.IMoveRepository
	typeServ    types.TypeReadService
}

func NewBattlePokemonResolver(effectsResolver IBattleEffectsResolver, pokemonRepo repository.IPokemonRepository,
	moveRepo repository.IMoveRepository) *BattlePokemonResolver {
	return &BattlePokemonResolver{
		IBattleEffectsResolver: effectsResolver,
		pokemonRepo:            pokemonRepo,
		moveRepo:               moveRepo,
		typeServ:               types.NewTypeReadService(),
	}
}

func (r BattlePokemonResolver) Resolve(
	adjustment s_damages.BattlePokemonAdjustment, moveId identifier.MoveId,
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
	moveId identifier.MoveId, resultChan chan<- AttackSide) error {
	pokemon, err := r.pokemonRepo.FindById(adjustment.PokemonId())
	if err != nil {
		return err
	}
	effects, err := r.resolve(adjustment)
	if err != nil {
		return err
	}
	move, err := r.moveRepo.FindById(moveId.Value())
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
	effects, err := r.resolve(adjustment)
	if err != nil {
		return err
	}
	actualValues := pokemon.ResolveActualValues(adjustment.EffortValues())
	defensePokemon := NewDefenseSide(actualValues, pokemon.TypeSet(), adjustment.Nature(), *effects)
	resultChan <- defensePokemon
	return nil
}

func (r BattlePokemonResolver) resolveTypeCompatibility(
	attackMoveType value.PokemonType, defensePokemonType value.PokemonTypeSet) float32 {
	typeCompatibility := r.typeServ.GetAttackTypeCompatibility(attackMoveType.NameEN())
	return typeCompatibility.GetTotalDamageRate(defensePokemonType)
}
