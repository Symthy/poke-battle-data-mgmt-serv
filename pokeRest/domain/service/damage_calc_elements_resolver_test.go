package service

import (
	"testing"

	s_damages "github.com/Symthy/PokeRest/pokeRest/application/service/damages"
	"github.com/Symthy/PokeRest/pokeRest/application/service/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/test/mock/move"
	"github.com/Symthy/PokeRest/test/mock/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func battleEffectsPM15SM15() *battles.BattleEffects {
	effectsBuilder := battles.NewBattleEffectsBuilder()
	effectsBuilder.AddCorrectionValues(
		battles.NewDefaultCorrectionValue(battles.CorrectionPhysicalMove, 1.5, nil),
		battles.NewDefaultCorrectionValue(battles.CorrectionSpecialMove, 1.5, nil))
	return effectsBuilder.Build()
}

func battleEffectsPP15() *battles.BattleEffects {
	effectsBuilder := battles.NewBattleEffectsBuilder()
	effectsBuilder.AddCorrectionValues(battles.NewDefaultCorrectionValue(battles.CorrectionPhysicalPower, 1.5, nil))
	return effectsBuilder.Build()
}

func battleEffectsSP15() *battles.BattleEffects {
	effectsBuilder := battles.NewBattleEffectsBuilder()
	effectsBuilder.AddCorrectionValues(battles.NewDefaultCorrectionValue(battles.CorrectionSpecialPower, 1.5, nil))
	return effectsBuilder.Build()
}

var _ battleAttackSideResolver = (*mockAttackSideReolver)(nil)

type mockAttackSideReolver struct{}

func (_m mockAttackSideReolver) resolve(adjustment *s_damages.BattlePokemonAdjustment, moveId identifier.MoveId, resultChan chan<- AttackSide) error {
	actualValues := value.NewPokemonActualValues(177, 178, 200, 70, 80, 71)
	attackSide := NewAttackSide(actualValues, value.NewPokemonTypeSet(value.Steel(), value.Rock()), value.NatureAdamant(),
		battleEffectsPM15SM15(), move.MoveHeadSmash(), true)
	resultChan <- attackSide
	return nil
}

var _ battleDefenseSideResolver = (*mockDefenseSideResolver)(nil)

type mockDefenseSideResolver struct{}

func (_m mockDefenseSideResolver) resolve(adjustment *s_damages.BattlePokemonAdjustment, resultChan chan<- DefenseSide) error {
	actualValues := value.NewPokemonActualValues(213, 152, 161, 90, 105, 123)
	defenseSide := NewDefenseSide(actualValues, value.NewPokemonTypeSet(value.Dragon(), value.Ground()), value.NatureImpish(),
		battleEffectsPM15SM15(), false)
	resultChan <- defenseSide
	return nil
}

func TestBattleDataSetResolver(t *testing.T) {
	cases := []struct {
		inputAttackSideResolver             battleAttackSideResolver
		inputDefenseSideResolver            battleDefenseSideResolver
		expectedAttackPokemonTypeFirst      value.PokemonType
		expectedAttackPokemonTypeSecond     value.PokemonType
		expectedAttackPokemonActualValueS   uint16
		expectedDefensePokemonTypeFirst     value.PokemonType
		expectedDefencePokemonTypeSecond    value.PokemonType
		expectedDefensePokemonActualValueS  uint16
		expectedMoveType                    value.PokemonType
		expectedHasItemAttackSide           bool
		expectedHasItemDefenseSide          bool
		expectedTypeCompatibilityDamageRate float32
	}{
		{
			inputAttackSideResolver:             mockAttackSideReolver{},
			inputDefenseSideResolver:            mockDefenseSideResolver{},
			expectedAttackPokemonTypeFirst:      value.Steel(),
			expectedAttackPokemonTypeSecond:     value.Rock(),
			expectedAttackPokemonActualValueS:   71,
			expectedDefensePokemonTypeFirst:     value.Dragon(),
			expectedDefencePokemonTypeSecond:    value.Ground(),
			expectedDefensePokemonActualValueS:  123,
			expectedMoveType:                    value.Rock(),
			expectedHasItemAttackSide:           true,
			expectedHasItemDefenseSide:          false,
			expectedTypeCompatibilityDamageRate: 0.5,
		},
	}

	dummyId, _ := identifier.NewMoveId(1)
	for _, tt := range cases {
		resolver := DamageCalcElementsService{tt.inputAttackSideResolver, tt.inputDefenseSideResolver, types.NewTypeReadService()}
		elements, err := resolver.Resolve(nil, *dummyId)
		actual := elements.GetPokemonBattleDataSet()
		assert.NoError(t, err)
		assert.Equal(t, tt.expectedAttackPokemonTypeFirst, actual.AttackPokemonTypeOfFirst())
		assert.Equal(t, tt.expectedAttackPokemonTypeSecond, actual.AttackPokemonTypeOfSecond())
		assert.Equal(t, tt.expectedAttackPokemonActualValueS, actual.AttackPokemonCorrectedActualValueS())
		assert.Equal(t, tt.expectedDefensePokemonTypeFirst, actual.DefensePokemonTypeOfFirst())
		assert.Equal(t, tt.expectedDefencePokemonTypeSecond, actual.DefensePokemonTypeOfSecond())
		assert.Equal(t, tt.expectedDefensePokemonActualValueS, actual.DefensePokemonCorrectedActualValueS())
		assert.Equal(t, tt.expectedMoveType, actual.MovePokemonType())
		assert.Equal(t, tt.expectedHasItemAttackSide, actual.HasItemAttackSide())
		assert.Equal(t, tt.expectedHasItemDefenseSide, actual.HasItemDefenseSide())
		assert.Equal(t, tt.expectedTypeCompatibilityDamageRate, actual.TypeCompatibilityDamageRate())
	}
}

var _ repository.IAbilityRepository = (*MockAbilityRepository)(nil)

type MockAbilityRepository struct {
	mock.Mock
}

func (_m *MockAbilityRepository) FindById(abilityId uint16) (*abilities.Ability, error) {
	ret := _m.Called(abilityId)
	return ret.Get(0).(*abilities.Ability), ret.Error(1)
}

func (_m MockAbilityRepository) FindAll(page uint32, pageSize uint16) (*abilities.Abilities, error) {
	return nil, nil
}

func (_m MockAbilityRepository) FindOfPokemon(pokemonId uint16) (*abilities.Abilities, error) {
	return nil, nil
}

var _ repository.IHeldItemRepository = (*MockHeldItemRepository)(nil)

type MockHeldItemRepository struct {
	mock.Mock
}

func (_m *MockHeldItemRepository) FindById(itemId uint16) (*items.HeldItem, error) {
	ret := _m.Called(itemId)
	return ret.Get(0).(*items.HeldItem), ret.Error(1)
}

func (_m MockHeldItemRepository) FindAll(page uint32, pageSize uint16) (*items.HeldItems, error) {
	return nil, nil
}

var _ repository.IPokemonRepository = (*MockPokemonRepository)(nil)

type MockPokemonRepository struct {
	mock.Mock
}

func (_m MockPokemonRepository) FindById(id uint16) (*pokemons.Pokemon, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*pokemons.Pokemon), ret.Error(1)
}
func (_m MockPokemonRepository) FindAll(page uint32, pageSize uint16) (*pokemons.Pokemons, error) {
	return nil, nil
}

var _ repository.IMoveRepository = (*MockMoveRepository)(nil)

type MockMoveRepository struct {
	mock.Mock
}

func (_m MockMoveRepository) FindById(id uint16) (*moves.Move, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*moves.Move), ret.Error(1)
}
func (_m MockMoveRepository) FindAll(page uint32, pageSize uint16) (*moves.Moves, error) {
	return nil, nil
}

func (_m MockMoveRepository) FindOfPokemon(pokemonId uint16) (*moves.Moves, error) {
	return nil, nil
}

func buildMockAbilityRepository(abilityId identifier.AbilityId, battleEffects *battles.BattleEffects) *MockAbilityRepository {
	mockAbilityRepo := new(MockAbilityRepository)
	abilityBuilder := factory.NewAbilityBuilder()
	abilityBuilder.Id(uint64(abilityId.Value()))
	abilityBuilder.BattleEffects(battleEffects)
	ability, _ := abilityBuilder.BuildDomain()
	mockAbilityRepo.On("FindById", abilityId.Value()).Return(ability, nil)
	return mockAbilityRepo
}

func buildMockHeldItemRepository(heldItemId identifier.HeldItemId, battleEffects *battles.BattleEffects) *MockHeldItemRepository {
	mockHeldItemRepo := new(MockHeldItemRepository)
	itemBuilder := factory.NewHeldItemBuilder()
	itemBuilder.Id(uint64(heldItemId.Value()))
	itemBuilder.BattleEffects(battleEffects)
	item, _ := itemBuilder.BuildDomain()
	mockHeldItemRepo.On("FindById", heldItemId.Value()).Return(item, nil)
	return mockHeldItemRepo
}

func buildMockPokemonRepository(inputPokemon *pokemons.Pokemon) *MockPokemonRepository {
	mockPokemonRepo := new(MockPokemonRepository)
	mockPokemonRepo.On("FindById", inputPokemon.Id().Value()).Return(inputPokemon, nil)
	return mockPokemonRepo
}

func TestAttackSideResolver(t *testing.T) {
	cases := []struct {
		inputPokemon          *pokemons.Pokemon
		inputAbilityEffects   *battles.BattleEffects
		inputHeldItemEffects  *battles.BattleEffects
		inputNature           value.PokemonNature
		inputMoveSpecies      value.MoveSpecies
		inputMoveType         value.PokemonType
		inputMovePower        uint16 // 威力
		inputIsMoveContained  bool
		inputCanMoveGuard     bool
		inputMoveEffects      *battles.BattleEffects
		expectedActualValues  *value.PokemonActualValues
		expectedBattleEffects *battles.BattleEffects
	}{
		{
			inputPokemon:         pokemon.Venusaur003(),
			inputAbilityEffects:  battleEffectsPM15SM15(),
			inputHeldItemEffects: battleEffectsSP15(),
			inputNature:          value.NatureModest(),
			inputMoveSpecies:     value.MoveSpeciesSpecial,
			inputMovePower:       uint16(90),
			inputIsMoveContained: false,
			inputCanMoveGuard:    true,
			inputMoveEffects:     battleEffectsPP15(),
			expectedActualValues: value.NewPokemonActualValues(187, 102, 103, 152, 120, 101),
			expectedBattleEffects: func() *battles.BattleEffects {
				effetcs := battles.NewEmptyBattleEffects()
				effetcs.Merge(battleEffectsPM15SM15())
				effetcs.Merge(battleEffectsSP15())
				effetcs.Merge(battleEffectsPP15())
				return effetcs
			}(),
		},
	}

	abilityId, _ := identifier.NewAbilityId(1)
	heldItemId, _ := identifier.NewHeldItemId(1)
	for _, tt := range cases {
		effortValues, _ := value.NewEffortValues(252, 0, 0, 252, 0, 4)

		adjustment := s_damages.NewBattlePokemonAdjustment(tt.inputPokemon.Id(), *abilityId, *heldItemId, tt.inputNature, effortValues)

		effectsResolver := NewBattleEffectsResolver(
			buildMockAbilityRepository(*abilityId, tt.inputAbilityEffects),
			buildMockHeldItemRepository(*heldItemId, tt.inputHeldItemEffects),
		)

		mockPokemonRepo := buildMockPokemonRepository(tt.inputPokemon)

		builder := factory.NewMoveBuilder()
		builder.Id(1)
		builder.Species(tt.inputMoveSpecies.String())
		builder.MovePokemonType(tt.inputMoveType.ToString())
		builder.Power(uint64(tt.inputMovePower))
		builder.SetCanGuard(tt.inputCanMoveGuard)
		builder.SetIsContained(tt.inputIsMoveContained)
		builder.BattleEffects(tt.inputMoveEffects)
		move, _ := builder.BuildDomain()
		mockMoveRepo := new(MockMoveRepository)
		mockMoveRepo.On("FindById", move.Id().Value()).Return(move, nil)

		resolver := attackSideResolver{effectsResolver, mockPokemonRepo, mockMoveRepo}
		attackSideChan := make(chan AttackSide, 1)
		resolver.resolve(adjustment, move.Id(), attackSideChan)
		actual := <-attackSideChan

		assert.Equal(t, tt.expectedActualValues, actual.actualValues)
		assert.Equal(t, tt.inputPokemon.TypeSet(), actual.pokemonTypeSet)
		assert.Equal(t, tt.inputNature, actual.nature)
		assert.Equal(t, tt.inputMoveSpecies, actual.moveSpecies)
		assert.Equal(t, tt.inputMovePower, actual.movePower)
		assert.Equal(t, tt.inputIsMoveContained, actual.isMoveContained)
		assert.Equal(t, tt.inputCanMoveGuard, actual.canMoveGuard)
		assert.True(t, actual.hasItem)
		assert.Equal(t, tt.expectedBattleEffects, actual.additionalEffects)
	}
}

func TestToDefenseSideResolver(t *testing.T) {
	cases := []struct {
		inputPokemon          *pokemons.Pokemon
		inputAbilityEffects   *battles.BattleEffects
		inputHeldItemEffects  *battles.BattleEffects
		inputNature           value.PokemonNature
		expectedActualValues  *value.PokemonActualValues
		expectedBattleEffects *battles.BattleEffects
	}{
		{
			inputPokemon:         pokemon.Venusaur003(),
			inputAbilityEffects:  battleEffectsPM15SM15(),
			inputHeldItemEffects: battleEffectsSP15(),
			inputNature:          value.NatureModest(),
			expectedActualValues: value.NewPokemonActualValues(187, 102, 103, 152, 120, 101),
			expectedBattleEffects: func() *battles.BattleEffects {
				effetcs := battles.NewEmptyBattleEffects()
				effetcs.Merge(battleEffectsPM15SM15())
				effetcs.Merge(battleEffectsSP15())
				return effetcs
			}(),
		},
	}

	abilityId, _ := identifier.NewAbilityId(1)
	heldItemId, _ := identifier.NewHeldItemId(1)
	for _, tt := range cases {
		effortValues, _ := value.NewEffortValues(252, 0, 0, 252, 0, 4)

		adjustment := s_damages.NewBattlePokemonAdjustment(tt.inputPokemon.Id(), *abilityId, *heldItemId, tt.inputNature, effortValues)

		effectsResolver := NewBattleEffectsResolver(
			buildMockAbilityRepository(*abilityId, tt.inputAbilityEffects),
			buildMockHeldItemRepository(*heldItemId, tt.inputHeldItemEffects),
		)

		mockPokemonRepo := buildMockPokemonRepository(tt.inputPokemon)

		resolver := defenseSideResolver{effectsResolver, mockPokemonRepo}
		defenseSideChan := make(chan DefenseSide, 1)
		resolver.resolve(adjustment, defenseSideChan)
		actual := <-defenseSideChan

		assert.Equal(t, tt.expectedActualValues, actual.actualValues)
		assert.Equal(t, tt.inputPokemon.TypeSet(), actual.pokemonTypeSet)
		assert.Equal(t, tt.inputNature, actual.nature)
		assert.True(t, actual.hasItem)
		assert.Equal(t, tt.expectedBattleEffects, actual.additionalEffects)
	}

}
