package conv

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/common/lists"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/factory"
)

func ToDomainPokemon(schema schema.Pokemon) (*pokemons.Pokemon, error) {
	if _, err := schema.AbilityId1.Value(); err != nil {
		// Todo: error wrap
		return nil, err
	}
	abilityId1 := uint16(schema.AbilityId1.Int16)
	if _, err := schema.AbilityId2.Value(); err != nil {
		return nil, err
	}
	abilityId2 := uint16(schema.AbilityId2.Int16)
	if _, err := schema.HiddenAbilityId.Value(); err != nil {
		return nil, err
	}
	hiddenAbilityId := uint16(schema.HiddenAbilityId.Int16)

	builder := factory.NewPokemonBuilder()
	builder.Id(schema.ID)
	builder.PokedexNo(schema.PokedexNo)
	builder.FormNo(schema.FormNo)
	builder.FormName(schema.FormName)
	builder.Name(schema.Name)
	builder.EnglishName(schema.EnglishName)
	builder.Generation(schema.Generation)
	builder.TypeOne(schema.Type1.String())
	builder.TypeTwo(schema.Type2.String())
	builder.AbilityIdOne(abilityId1)
	builder.AbilityIdTwo(abilityId2)
	builder.HiddenAbilityId(hiddenAbilityId)
	builder.BaseStatsH(schema.BaseStatsH)
	builder.BaseStatsA(schema.BaseStatsA)
	builder.BaseStatsB(schema.BaseStatsB)
	builder.BaseStatsC(schema.BaseStatsC)
	builder.BaseStatsD(schema.BaseStatsD)
	builder.BaseStatsS(schema.BaseStatsS)
	builder.SetIsFinalEvolution(schema.IsFinalEvolution)
	return builder.BuildDomain()
}

func ToDomainAbility(ability schema.Ability) (*abilities.Ability, error) {
	input := factory.NewAbilityInput(uint64(ability.ID), ability.Name, ability.Description,
		toBattleEffects(ability.BattleEffects))
	return input.BuildDomain()
}

func ToDomainMove(schema schema.Move) (*moves.Move, error) {
	input := factory.NewMoveInput(uint64(schema.ID), schema.Name, schema.Species.String(), schema.Type.String(),
		uint64(schema.Power), schema.Accuracy, uint64(schema.PP), *schema.IsContained, *schema.CanGuard)
	return input.BuildDomain()
}

func ToDomainHeldItem(schema schema.HeldItem) (*items.HeldItem, error) {
	input := factory.NewHeldItemInput(uint64(schema.ID), schema.Name, schema.Description,
		toBattleEffects(schema.BattleEffects))
	return input.BuildDomain()
}

func ToDomainTrainedPokemon(schema schema.TrainedPokemon) (*trainings.TrainedPokemon, error) {
	nickname := ""
	if schema.Nickname != nil {
		nickname = *schema.Nickname
	}
	description := ""
	if schema.Description != nil {
		description = *schema.Description
	}
	var userId uint64 = 0
	if schema.CreateUserId != nil {
		userId = *schema.CreateUserId
	}
	adjustmentInput := toTrainedPokemonAdjustmentInput(schema.TrainedPokemonAdjustment)
	input := factory.NewTrainedPokemonInput(schema.ID, schema.Gender.String(), nickname,
		description, schema.IsPrivate, userId, adjustmentInput)
	return input.BuildDomain()
}

func toTrainedPokemonAdjustmentInput(schema schema.TrainedPokemonAdjustment) factory.TrainedPokemonAdjustmentInput {
	builder := factory.NewTrainedPokemonAdjustmentBuilder()
	builder.Id(schema.ID)
	builder.PokemonId(uint64(schema.PokemonId))
	builder.Nature(schema.Nature.String())
	var ability uint64 = 0
	if schema.AbilityId != nil {
		ability = uint64(*schema.AbilityId)
	}
	builder.AbilityId(ability)
	var heldItem uint64 = 0
	if schema.HeldItemId != nil {
		heldItem = uint64(*schema.HeldItemId)
	}
	builder.HeldItemId(heldItem)
	builder.EffortValueH(uint64(schema.EffortValueH))
	builder.EffortValueA(uint64(schema.EffortValueA))
	builder.EffortValueB(uint64(schema.EffortValueB))
	builder.EffortValueC(uint64(schema.EffortValueC))
	builder.EffortValueD(uint64(schema.EffortValueD))
	builder.EffortValueS(uint64(schema.EffortValueS))
	var move1 uint64 = 0
	var move2 uint64 = 0
	var move3 uint64 = 0
	var move4 uint64 = 0
	if schema.MoveId1 != nil {
		move1 = uint64(*schema.MoveId1)
	}
	if schema.MoveId2 != nil {
		move2 = uint64(*schema.MoveId2)
	}
	if schema.MoveId2 != nil {
		move3 = uint64(*schema.MoveId2)
	}
	if schema.MoveId3 != nil {
		move4 = uint64(*schema.MoveId3)
	}
	builder.MoveIdFirst(move1)
	builder.MoveIdSecond(move2)
	builder.MoveIdThird(move3)
	builder.MoveIdFourth(move4)
	return builder
}

func ToDomainTrainedPokemonAdjustment(schema schema.TrainedPokemonAdjustment) (*trainings.TrainedPokemonAdjustment, error) {
	return toTrainedPokemonAdjustmentInput(schema).BuildDomain()
}

func ToDomainTrainedPokemonAttackTarget(schema schema.TrainedPokemonAttackTarget) (*trainings.TrainedPokemonAttackTarget, error) {
	builder := factory.NewTrainedAttackTargetBuilder()
	builder.Id(schema.ID)
	builder.TrainedPokemonId(schema.TrainedPokemonId)
	builder.MoveId(uint64(schema.MoveId))
	builder.TargetPokemonId(uint64(schema.TargetPokemonId))
	builder.TargetPokemonNature(schema.TargetPokemonNature.String())
	builder.TargetPokemonAbilityId(uint64(schema.TargetPokemonAbilityId))
	builder.TargetPokemonHeldItemId(uint64(schema.TargetPokemonHeldItemId))
	builder.TargetPokemonEffortValueH(uint64(schema.TargetPokemonEffortValueH))
	builder.TargetPokemonEffortValueB(uint64(schema.TargetPokemonEffortValueB))
	builder.TargetPokemonEffortValueD(uint64(schema.TargetPokemonEffortValueD))
	return builder.BuildDomain()
}

func ToDomainTrainedPokemonDefenceTarget(schema schema.TrainedPokemonDefenceTarget) (*trainings.TrainedPokemonDefenceTarget, error) {
	builder := factory.NewTrainedDefenseTargetBuilder()
	builder.Id(schema.ID)
	builder.TrainedPokemonId(schema.TrainedPokemonId)
	builder.MoveId(uint64(schema.MoveId))
	builder.TargetPokemonId(uint64(schema.TargetPokemonId))
	builder.TargetPokemonNature(schema.TargetPokemonNature.String())
	builder.TargetPokemonAbilityId(uint64(schema.TargetPokemonAbilityId))
	builder.TargetPokemonHeldItemId(uint64(schema.TargetPokemonHeldItemId))
	builder.TargetPokemonEffortValueA(uint64(schema.TargetPokemonEffortValueA))
	builder.TargetPokemonEffortValueC(uint64(schema.TargetPokemonEffortValueC))
	return builder.BuildDomain()
}

func ToDomainParty(party schema.Party) (*parties.Party, error) {
	var userId uint64 = 0
	if party.CreateUserId != nil {
		userId = *party.CreateUserId
	}
	partyResultIds := lists.Map(
		party.PartyResult,
		func(result schema.PartySeasonResult) uint64 {
			return result.ID
		})
	partyTagIds := lists.Map(
		party.PartyTag,
		func(tag schema.PartyTag) uint64 {
			return tag.ID
		})
	trainedPokemonIds := lists.Map(
		party.TrainedPokemon,
		func(tp schema.TrainedPokemon) uint64 {
			return tp.ID
		})
	input := factory.NewPartyInput(party.ID, party.Name, string(party.BattleFormat), party.IsPrivate,
		userId, partyResultIds, partyTagIds, trainedPokemonIds)
	return input.BuildDomain()
}

func ToDomainPartyTag(schema schema.PartyTag) (*parties.PartyTag, error) {
	input := factory.NewPartyTagInput(schema.ID, schema.Name, *schema.IsGeneration, *schema.IsSeason)
	return input.BuildDomain()
}

func ToDomainPartyBattleResult(schema schema.PartySeasonResult) (*parties.PartyBattleResult, error) {
	builder := factory.NewPartyBattleResultBuilder()
	builder.Id(schema.ID)
	builder.Generation(uint64(schema.Generation))
	builder.Series(uint64(schema.Series))
	builder.Season(uint64(schema.Season))
	builder.MaxRate(uint64(schema.MaxRate))
	builder.SeasonRanking(uint64(schema.SeasonRanking))
	builder.MaxSeasonRanking(uint64(schema.MaxSeasonRanking))
	builder.WinCount(uint64(schema.WinCount))
	builder.LoseCount(uint64(schema.LoseCount))
	return builder.BuildDomain()
}

func ToDomainBattleRecord(schema schema.BattleRecord) (*battles.BattleRecord, error) {
	selfElectionPokemonIds := []uint16{}
	lists.AddsToList(selfElectionPokemonIds, schema.SelfElectionPokemonId1,
		schema.SelfElectionPokemonId2, schema.SelfElectionPokemonId3, schema.SelfElectionPokemonId4)
	selfTrainedPokemonIds := []uint64{}
	lists.AddsToList(selfTrainedPokemonIds, schema.SelfTrainedPokemonId1,
		schema.SelfTrainedPokemonId2, schema.SelfTrainedPokemonId3, schema.SelfTrainedPokemonId4)
	opponentElectionPokemonIds := []uint16{}
	lists.AddsToList(opponentElectionPokemonIds, schema.OpponentElectionPokemonId1,
		schema.OpponentElectionPokemonId2, schema.OpponentElectionPokemonId3, schema.OpponentElectionPokemonId4)

	opponentPartyInput := toBattleOpponentPartyInput(schema.BattleOpponentParty)

	input := factory.NewBattleRecordInput(schema.ID, schema.PartyId, schema.UserId,
		uint64(schema.Generation), uint64(schema.Series), uint64(schema.Season), schema.Result.String(),
		lists.ConvertTypeUint16To64(selfElectionPokemonIds), selfTrainedPokemonIds,
		lists.ConvertTypeUint16To64(opponentElectionPokemonIds), opponentPartyInput)
	return input.BuildDomain()
}

func ToDomainBattleOpponentParty(schema schema.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	input := toBattleOpponentPartyInput(schema)
	return input.BuildDomain()
}

func toBattleOpponentPartyInput(schema schema.BattleOpponentParty) factory.BattleOpponentPartyInput {
	opponentPokemonIds := []uint16{}
	lists.AddsToList(opponentPokemonIds,
		schema.OpponentPokemonId1, schema.OpponentPokemonId2, schema.OpponentPokemonId3,
		schema.OpponentPokemonId4, schema.OpponentPokemonId5, schema.OpponentPokemonId6)
	input := factory.NewBattleOpponentPartyInput(schema.ID, lists.ConvertTypeUint16To64(opponentPokemonIds)...)
	return input
}

func ToDomainSeasonPeriod(schema schema.BattleSeason) (*battles.SeasonPeriod, error) {
	return battles.NewSeasonPeriod(uint64(schema.Generation), uint64(schema.Series), uint64(schema.Season),
		schema.StartDateTime, schema.EndDateTime)
}

func ToDomainSeason(schema schema.BattleSeason) (*battles.Season, error) {
	return battles.NewSeason(uint64(schema.Generation), uint64(schema.Series), uint64(schema.Season))
}

func ToDomainUser(schema schema.User) (*users.User, error) {
	input := factory.NewUserInput(schema.ID, schema.Name, schema.Role.String())
	return input.BuildDomain()
}
