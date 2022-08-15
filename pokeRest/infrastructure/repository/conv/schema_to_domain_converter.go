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

func ToDomainPokemon(schema *schema.Pokemon) (*pokemons.Pokemon, error) {
	if _, err := schema.AbilityID1.Value(); err != nil {
		// Todo: error wrap
		return nil, err
	}
	abilityId1 := uint16(schema.AbilityID1.Int16)
	if _, err := schema.AbilityID2.Value(); err != nil {
		return nil, err
	}
	abilityId2 := uint16(schema.AbilityID2.Int16)
	if _, err := schema.HiddenAbilityID.Value(); err != nil {
		return nil, err
	}
	hiddenAbilityId := uint16(schema.HiddenAbilityID.Int16)

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

func ToDomainAbility(ability *schema.Ability) (*abilities.Ability, error) {
	input := factory.NewAbilityInput(uint64(ability.ID), ability.Name, ability.Description,
		toBattleEffects(ability.BattleEffects))
	return input.BuildDomain()
}

func ToDomainMove(schema *schema.Move) (*moves.Move, error) {
	builder := factory.NewMoveBuilder()
	builder.Id(uint64(schema.ID))
	builder.Name(schema.Name)
	builder.Species(schema.Species.String())
	builder.MovePokemonType(schema.Type.String())
	builder.Power(uint64(schema.Power))
	builder.Accuracy(schema.Accuracy)
	builder.PP(uint64(schema.PP))
	builder.SetIsContained(*schema.IsContained)
	builder.SetCanGuard(*schema.CanGuard)
	return builder.BuildDomain()
}

func ToDomainHeldItem(schema *schema.HeldItem) (*items.HeldItem, error) {
	builder := factory.NewHeldItemBuilder()
	builder.Id(uint64(schema.ID))
	builder.Name(schema.Name)
	builder.Description(schema.Description)
	builder.BattleEffects(toBattleEffects(schema.BattleEffects))
	return builder.BuildDomain()
}

func ToDomainTrainedPokemon(schema *schema.TrainedPokemon) (*trainings.TrainedPokemon, error) {
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
	adjustment := toTrainedPokemonAdjustmentInput(&schema.TrainedPokemonAdjustment)
	builder := factory.NewTrainedPokemonBuilder()
	builder.Id(schema.ID)
	builder.Gender(schema.Gender.String())
	builder.Nickname(nickname)
	builder.Description(description)
	builder.SetIsPrivate(schema.IsPrivate)
	builder.UserId(userId)
	builder.Adjustment(adjustment)
	return builder.BuildDomain()
}

func toTrainedPokemonAdjustmentInput(schema *schema.TrainedPokemonAdjustment) factory.TrainedPokemonAdjustmentInput {
	builder := factory.NewTrainedPokemonAdjustmentBuilder()
	builder.Id(schema.ID)
	builder.PokemonId(uint64(schema.PokemonID))
	builder.Nature(schema.Nature.String())
	var ability uint64 = 0
	if schema.AbilityID != nil {
		ability = uint64(*schema.AbilityID)
	}
	builder.AbilityId(ability)
	var heldItem uint64 = 0
	if schema.HeldItemID != nil {
		heldItem = uint64(*schema.HeldItemID)
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
	if schema.MoveID1 != nil {
		move1 = uint64(*schema.MoveID1)
	}
	if schema.MoveID2 != nil {
		move2 = uint64(*schema.MoveID2)
	}
	if schema.MoveID2 != nil {
		move3 = uint64(*schema.MoveID2)
	}
	if schema.MoveID3 != nil {
		move4 = uint64(*schema.MoveID3)
	}
	builder.MoveIdFirst(move1)
	builder.MoveIdSecond(move2)
	builder.MoveIdThird(move3)
	builder.MoveIdFourth(move4)
	return builder
}

func ToDomainTrainedPokemonAdjustment(schema *schema.TrainedPokemonAdjustment) (*trainings.TrainedPokemonAdjustment, error) {
	return toTrainedPokemonAdjustmentInput(schema).BuildDomain()
}

func ToDomainTrainedPokemonAttackTarget(schema *schema.TrainedPokemonAttackTarget) (*trainings.TrainedPokemonAttackTarget, error) {
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

func ToDomainTrainedPokemonDefenceTarget(schema *schema.TrainedPokemonDefenseTarget) (*trainings.TrainedPokemonDefenseTarget, error) {
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

func ToDomainParty(party *schema.Party) (*parties.Party, error) {
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
	builder := factory.NewPartyBuilder()
	builder.Id(party.ID)
	builder.Name(party.Name)
	builder.BattleFormat(string(party.BattleFormat))
	builder.SetIsPrivate(party.IsPrivate)
	builder.UserId(userId)
	builder.PartyResultIds(partyResultIds)
	builder.PartyTagIds(partyTagIds)
	builder.TrainedPokemonIds(trainedPokemonIds)
	return builder.BuildDomain()
}

func ToDomainPartyTag(schema *schema.PartyTag) (*parties.PartyTag, error) {
	builder := factory.NewPartyTagBuilder()
	builder.Id(schema.ID)
	builder.Name(schema.Name)
	builder.SetIsGeneration(*schema.IsGeneration)
	builder.SetIsSeason(*schema.IsSeason)
	return builder.BuildDomain()
}

func ToDomainPartyBattleResult(schema *schema.PartySeasonResult) (*parties.PartyBattleResult, error) {
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

func ToDomainBattleRecord(schema *schema.BattleRecord) (*battles.BattleRecord, error) {
	selfElectionPokemonIds := []uint16{}
	lists.AddLiterals(selfElectionPokemonIds, schema.SelfElectionPokemonId1,
		schema.SelfElectionPokemonId2, schema.SelfElectionPokemonId3, schema.SelfElectionPokemonId4)
	selfTrainedPokemonIds := []uint64{}
	lists.AddLiterals(selfTrainedPokemonIds, schema.SelfTrainedPokemonId1,
		schema.SelfTrainedPokemonId2, schema.SelfTrainedPokemonId3, schema.SelfTrainedPokemonId4)
	opponentElectionPokemonIds := []uint16{}
	lists.AddLiterals(opponentElectionPokemonIds, schema.OpponentElectionPokemonId1,
		schema.OpponentElectionPokemonId2, schema.OpponentElectionPokemonId3, schema.OpponentElectionPokemonId4)

	season := factory.NewSeasonBuilder().
		Generation(uint64(schema.BattleSeason.Generation)).
		Series(uint64(schema.BattleSeason.Series)).
		Season(uint64(schema.BattleSeason.Season))

	opponentParty := toBattleOpponentPartyInput(&schema.BattleOpponentParty)

	builder := factory.NewBattleRecordBuilder().
		Id(schema.ID).
		PartyId(schema.PartyID).
		UserId(schema.UserID).
		Season(season).
		BattleResult(schema.Result.String()).
		SelfElectionPokemonIds(lists.ConvertTypeUint16To64(selfElectionPokemonIds)).
		SelfElectionTrainedPokemonIds(selfTrainedPokemonIds).
		OpponentElectionPokemonIds(lists.ConvertTypeUint16To64(opponentElectionPokemonIds)).
		BattleOpponentParty(opponentParty)
	return builder.BuildDomain()
}

func ToDomainBattleOpponentParty(schema *schema.BattleOpponentParty) (*battles.BattleOpponentParty, error) {
	input := toBattleOpponentPartyInput(schema)
	return input.BuildDomain()
}

func toBattleOpponentPartyInput(schema *schema.BattleOpponentParty) *factory.BattleOpponentPartyInput {
	opponentPokemonIds := []uint16{}
	lists.AddLiterals(opponentPokemonIds,
		schema.OpponentPokemonId1, schema.OpponentPokemonId2, schema.OpponentPokemonId3,
		schema.OpponentPokemonId4, schema.OpponentPokemonId5, schema.OpponentPokemonId6)
	input := factory.NewBattleOpponentPartyInput(schema.ID, lists.ConvertTypeUint16To64(opponentPokemonIds))
	return input
}

func ToDomainSeasonPeriod(schema schema.BattleSeason) (*battles.SeasonPeriod, error) {
	return battles.NewSeasonPeriod(uint64(schema.Generation), uint64(schema.Series), uint64(schema.Season),
		schema.StartDateTime, schema.EndDateTime)
}

func ToDomainSeason(schema schema.BattleSeason) (*battles.Season, error) {
	return battles.NewSeason(uint64(schema.Generation), uint64(schema.Series), uint64(schema.Season))
}

func ToDomainUser(schema *schema.User) (*users.User, error) {
	builder := factory.NewUserBuilder()
	builder.Id(schema.ID)
	builder.Name(schema.Name)
	builder.Role(schema.Role.String())
	return builder.BuildDomain()
}
