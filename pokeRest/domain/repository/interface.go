package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"gorm.io/gorm"
)

// common
type ISingleRecordFinder[T entity.IDomain[K], K entity.IValueId] interface {
	FindById(id uint) (*T, error)
}

type IAllRecordRepository[L entity.IDomains[T, K], T entity.IDomain[K], K entity.IValueId] interface {
	FindAll(page int, pageSize int) (*L, error)
}

type IPokemonStatsRepository[L entity.IDomains[T, K], T entity.IDomain[K], K entity.IValueId] interface {
	FindOfPokemon(pokemonId uint) (*L, error)
}

type IWritableRepository[TD entity.IDomain[K], K entity.IValueId] interface {
	Create(TD) (*TD, error)
	Update(TD) (*TD, error)
	Delete(id uint) (*TD, error)
}

type ITransactionalRepository interface {
	StartTransaction() error
	CancelTransaction() error
	FinishTransaction() error
	PanicPostProcess()
}

// special
type IPokemonRepository interface {
	ISingleRecordFinder[pokemons.Pokemon, identifier.PokemonId]
	IAllRecordRepository[pokemons.Pokemons, pokemons.Pokemon, identifier.PokemonId]
	IWritableRepository[pokemons.Pokemon, identifier.PokemonId]
}

type IAbilityRepository interface {
	ISingleRecordFinder[abilities.Ability, identifier.AbilityId]
	IAllRecordRepository[abilities.Abilities, abilities.Ability, identifier.AbilityId]
	IPokemonStatsRepository[abilities.Abilities, abilities.Ability, identifier.AbilityId]
}

type IMoveRepository interface {
	ISingleRecordFinder[moves.Move, identifier.MoveId]
	IAllRecordRepository[moves.Moves, moves.Move, identifier.MoveId]
	IPokemonStatsRepository[moves.Moves, moves.Move, identifier.MoveId]
}

type IHeldItemRepository interface {
	ISingleRecordFinder[items.HeldItem, identifier.HeldItemId]
	IAllRecordRepository[items.HeldItems, items.HeldItem, identifier.HeldItemId]
}

type IPartyTagRepository interface {
	ISingleRecordFinder[parties.PartyTag, identifier.PartyTagId]
	IAllRecordRepository[parties.PartyTags, parties.PartyTag, identifier.PartyTagId]
	IWritableRepository[parties.PartyTag, identifier.PartyTagId]
}

type IPartyRepository interface {
	ISingleRecordFinder[parties.Party, identifier.PartyId]
	IWritableRepository[parties.Party, identifier.PartyId]
}

type IPartyBattleResultRepository interface {
	IWritableRepository[parties.PartyBattleResult, identifier.PartyBattleResultId]
}

type IBattleRecordRepository interface {
	ISingleRecordFinder[battles.BattleRecord, identifier.BattleRecordId]
	IWritableRepository[battles.BattleRecord, identifier.BattleRecordId]
}

type IBattleSeasonRepository interface {
	FindCurrent() (*battles.Season, error)
	Find(battles.Season) (*battles.SeasonPeriods, error)
	FindAll() (*battles.SeasonPeriods, error)
}

type IBattleRecordTransactionalRepository interface {
	IBattleRecordRepository
	ITransactionalRepository
}

type IBattleOpponentPartyRepository interface {
	IWritableRepository[battles.BattleOpponentParty, identifier.BattleOpponentPartyId]
	FindParty(battles.BattleOpponentParty) (*battles.BattleOpponentParty, error)
}

type ITrainedPokemonRepository interface {
	IWritableRepository[trainings.TrainedPokemon, identifier.TrainedPokemonId]
}

// Todo: Don't want to depend gorm. wrap?
type ITrainedPokemonAdjustmentRepository interface {
	Find(*gorm.DB, trainings.TrainedPokemonAdjustment) (*trainings.TrainedPokemonAdjustment, error)
	IAllRecordRepository[trainings.TrainedPokemonAdjustments, trainings.TrainedPokemonAdjustment, identifier.TrainedAdjustmentId]
	IWritableRepository[trainings.TrainedPokemonAdjustment, identifier.TrainedAdjustmentId]
}

type ITrainedPokemonAttackRepository interface {
	IWritableRepository[trainings.TrainedPokemonAttackTarget, identifier.TrainedAttackTargetId]
}

type ITrainedPokemonDeffenceRepository interface {
	IWritableRepository[trainings.TrainedPokemonDefenceTarget, identifier.TrainedDefenseTargetId]
}

type IUserRepository interface {
	FindById(id uint) (*users.User, error)
	FindByName(targetName string, filterFields ...string) (*users.User, error)
	IWritableRepository[users.User, identifier.UserId]
}
