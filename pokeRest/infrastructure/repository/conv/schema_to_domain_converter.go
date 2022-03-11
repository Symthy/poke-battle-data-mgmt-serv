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
	abilityId1 := uint(schema.AbilityId1.Int16)
	if _, err := schema.AbilityId2.Value(); err != nil {
		return nil, err
	}
	abilityId2 := uint(schema.AbilityId2.Int16)
	if _, err := schema.HiddenAbilityId.Value(); err != nil {
		return nil, err
	}
	hiddenAbilityId := uint(schema.HiddenAbilityId.Int16)

	input := factory.NewPokemonInput(schema.ID, schema.PokedexNo, schema.FormNo, schema.FormName, schema.Name,
		schema.EnglishName, schema.Generation, schema.Type1.String(), schema.Type2.String(),
		abilityId1, abilityId2, hiddenAbilityId, schema.BaseStatsH, schema.BaseStatsA,
		schema.BaseStatsB, schema.BaseStatsC, schema.BaseStatsD, schema.BaseStatsS, schema.IsFinalEvolution)
	return input.BuildDomain()
}

func ToDomainAbility(ability schema.Ability) (*abilities.Ability, error) {
	input := factory.NewAbilityInput(ability.ID, ability.Name, ability.Description,
		toCorrectionValues(ability.CorrectionValue))
	return input.BuildDomain()
}

func ToDomainMove(schema schema.Move) (*moves.Move, error) {
	input := factory.NewMoveInput(schema.ID, schema.Name, schema.Species.String(), schema.Power,
		schema.Accuracy, schema.PP, *schema.IsContained, *schema.IsCanGuard)
	return input.BuildDomain()
}

func ToDomainHeldItem(schema schema.HeldItem) (*items.HeldItem, error) {
	input := factory.NewHeldItemInput(schema.ID, schema.Name, schema.Description,
		toCorrectionValues(schema.CorrectionValue))
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
	var userId uint = 0
	if schema.CreateUserId != nil {
		userId = *schema.CreateUserId
	}
	adjustmentInput := toTrainedPokemonAdjustmentInput(schema.TrainedPokemonAdjustment)
	input := factory.NewTrainedPokemonInput(schema.ID, schema.Gender.String(), nickname,
		description, schema.IsPrivate, userId, adjustmentInput)
	return input.BuildDomain()
}

func toTrainedPokemonAdjustmentInput(schema schema.TrainedPokemonAdjustment) factory.TrainedPokemonAdjustmentInput {
	input := factory.NewTrainedPokemonAdjustmentInput(
		schema.ID, schema.PokemonId, schema.Nature.String(), uint(*schema.AbilityId), uint(*schema.HeldItemId),
		schema.EffortValueH, schema.EffortValueA, schema.EffortValueB, schema.EffortValueC,
		schema.EffortValueD, schema.EffortValueS, uint(*schema.MoveId1), uint(*schema.MoveId2),
		uint(*schema.MoveId3), uint(*schema.MoveId4))
	return input
}

func ToDomainTrainedPokemonAdjustment(schema schema.TrainedPokemonAdjustment) (*trainings.TrainedPokemonAdjustment, error) {
	return toTrainedPokemonAdjustmentInput(schema).BuildDomain()
}

func ToDomainTrainedPokemonAttackTarget(schema schema.TrainedPokemonAttackTarget) (*trainings.TrainedPokemonAttackTarget, error) {
	input := factory.NewTrainedAttackTargetInput(schema.ID, schema.TrainedPokemonId, uint(schema.MoveId),
		uint(schema.TargetPokemonId), schema.TargetPokemonNature.String(), schema.TargetPokemonEffortValueH,
		schema.TargetPokemonEffortValueB, schema.TargetPokemonEffortValueD)
	return input.BuildDomain()
}

func ToDomainTrainedPokemonDefenceTarget(schema schema.TrainedPokemonDefenceTarget) (*trainings.TrainedPokemonDefenceTarget, error) {
	input := factory.NewTrainedDefenceTargetInput(schema.ID, schema.TrainedPokemonId,
		uint(schema.MoveId), uint(schema.TargetPokemonId), schema.TargetPokemonNature.String(),
		schema.TargetPokemonEffortValueA, schema.TargetPokemonEffortValueC)
	return input.BuildDomain()
}

func ToDomainParty(party schema.Party) (*parties.Party, error) {
	var userId uint = 0
	if party.CreateUserId != nil {
		userId = *party.CreateUserId
	}
	partyResultIds := lists.Map(
		party.PartyResult,
		func(result schema.PartySeasonResult) uint {
			return result.ID
		})
	partyTagIds := lists.Map(
		party.PartyTag,
		func(tag schema.PartyTag) uint {
			return tag.ID
		})
	trainedPokemonIds := lists.Map(
		party.TrainedPokemon,
		func(tp schema.TrainedPokemon) uint {
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
	input := factory.NewPartyBattleResultInput(
		schema.ID, schema.Generation, schema.Series, schema.Season, schema.MaxRate,
		schema.SeasonRanking, schema.MaxSeasonRanking, schema.WinCount, schema.LoseCount)
	return input.BuildDomain()
}

func ToDomainBattleRecord(schema schema.BattleRecord) (*battles.BattleRecord, error) {
	selfElectionPokemonIds := []uint{}
	lists.AddsToList(selfElectionPokemonIds, schema.SelfElectionPokemonId1,
		schema.SelfElectionPokemonId2, schema.SelfElectionPokemonId3, schema.SelfElectionPokemonId4)
	selfTrainedPokemonIds := []uint{}
	lists.AddsToList(selfTrainedPokemonIds, schema.SelfTrainedPokemonId1,
		schema.SelfTrainedPokemonId2, schema.SelfTrainedPokemonId3, schema.SelfTrainedPokemonId4)
	opponentElectionPokemonIds := []uint{}
	lists.AddsToList(opponentElectionPokemonIds, schema.OpponentElectionPokemonId1,
		schema.OpponentElectionPokemonId2, schema.OpponentElectionPokemonId3, schema.OpponentElectionPokemonId4)

	opponentPartyInput := toBattleOpponentPartyInput(schema.BattleOpponentParty)

	input := factory.NewBattleRecordInput(schema.ID, schema.PartyId, schema.UserId, schema.Generation,
		schema.Series, schema.Season, schema.Result.String(), selfElectionPokemonIds,
		selfTrainedPokemonIds, opponentElectionPokemonIds, opponentPartyInput)
	return input.BuildDomain()
}

func ToDomainBattleOpponentParty(schema schema.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	input := toBattleOpponentPartyInput(schema)
	return input.BuildDomain()
}

func toBattleOpponentPartyInput(schema schema.BattleOpponentParty) factory.BattleOpponentPartyInput {
	opponentPokemonIds := []uint{}
	lists.AddsToList(opponentPokemonIds,
		schema.OpponentPokemonId1, schema.OpponentPokemonId2, schema.OpponentPokemonId3,
		schema.OpponentPokemonId4, schema.OpponentPokemonId5, schema.OpponentPokemonId6)
	input := factory.NewBattleOpponentPartyInput(schema.ID, opponentPokemonIds...)
	return input
}

func ToDomainSeasonPeriod(schema schema.BattleSeason) (*battles.SeasonPeriod, error) {
	return battles.NewSeasonPeriod(schema.Generation, schema.Series, schema.Season,
		schema.StartDateTime, schema.EndDateTime)
}

func ToDomainSeason(schema schema.BattleSeason) (*battles.Season, error) {
	return battles.NewSeason(schema.Generation, schema.Series, schema.Season)
}

func ToDomainUser(schema schema.User) (*users.User, error) {
	input := factory.NewUserInput(schema.ID, schema.Name, schema.Role.String())
	return input.BuildDomain()
}
