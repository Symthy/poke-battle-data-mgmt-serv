package damages

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
)

type BattlePokemonResolver struct {
	pokemonRepo  repository.IPokemonRepository
	abilityRepo  repository.IAbilityRepository
	heldItemRepo repository.IHeldItemRepository
	moveRepo     repository.IMoveRepository
}

func (r BattlePokemonResolver) ResolveAttackSide(adjustment BattlePokemonAdjustment, moveId uint) (*BattleAttackPokemon, error) {
	pokemon, err := r.pokemonRepo.FindById(adjustment.pokemonId)
	if err != nil {
		return nil, err
	}
	// Todo: effort value
	actualValues := pokemon.ResolveActualValues()

	effects := value.BattleEffects{}
	ability, err := r.abilityRepo.FindById(adjustment.abilityId)
	if err != nil {
		return nil, err
	}
	ability.NotifyBattleEffects(effects)
	item, err := r.heldItemRepo.FindById(adjustment.heldItemId)
	if err != nil {
		return nil, err
	}
	item.NotifyBattleEffects(effects)
	move, err := r.moveRepo.FindById(moveId)
	if err != nil {
		return nil, err
	}
	move.NotifyBattleEffects(effects)

}

func (r BattlePokemonResolver) ResolveDefenceSide() BattleDefencePokemon {

}

var _ moves.IMoveNotificationForDamageCalc = (*BattleAttackPokemon)(nil)

type BattleAttackPokemon struct {
	value.PokemonActualValues
	pokemonType     value.PokemonTypeSet
	moveSpecies     value.MoveSpecies
	moveType        value.PokemonType
	movePower       int // 威力
	isMoveContained bool
	canMoveGuard    bool
}

func (b BattleAttackPokemon) SetMoveType(moveType value.PokemonType) {
	b.moveType = moveType
}
func (b BattleAttackPokemon) SetMoveSpecies(moveSpecies value.MoveSpecies) {
	b.moveSpecies = moveSpecies
}
func (b BattleAttackPokemon) SetMovePower(movePower int) {
	b.movePower = movePower
}
func (b BattleAttackPokemon) SetIsMoveContained(isMoveContained bool) {
	b.isMoveContained = isMoveContained
}
func (b BattleAttackPokemon) SetCanMoveGuard(canMoveGuard bool) {
	b.canMoveGuard = canMoveGuard
}

func (p BattleAttackPokemon) Power() int {
	return 0
}

type BattleDefencePokemon struct {
	value.PokemonActualValues
	pokemonType       value.PokemonTypeSet
	additionalEffects value.BattleEffects
}

type DamageCalculation struct {
	attackSide   BattleAttackPokemon
	deffenceSide BattleDefencePokemon
}
