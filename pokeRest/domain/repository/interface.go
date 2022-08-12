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
type ISingleRecordFinder[T entity.IDomain[K, I], K entity.IValueId[I], I uint16 | uint32 | uint64] interface {
	FindById(id I) (*T, error)
}

type IAllRecordRepository[L entity.IDomains[T, K, I], T entity.IDomain[K, I], K entity.IValueId[I], I uint16 | uint32 | uint64] interface {
	FindAll(page uint32, pageSize uint16) (*L, error)
}

type IPokemonStatsRepository[L entity.IDomains[T, K, uint16], T entity.IDomain[K, uint16], K entity.IValueId[uint16]] interface {
	FindOfPokemon(pokemonId uint16) (*L, error)
}

type IWritableRepository[TD entity.IDomain[K, I], K entity.IValueId[I], I uint16 | uint32 | uint64] interface {
	Create(*TD) (*TD, error)
	Update(*TD) (*TD, error)
	Delete(id I) (*TD, error)
}

type ITransactionalRepository interface {
	StartTransaction() error
	CancelTransaction() error
	FinishTransaction() error
	PanicPostProcess()
}

// special
type IPokemonRepository interface {
	ISingleRecordFinder[pokemons.Pokemon, identifier.PokemonId, uint16]
	IAllRecordRepository[pokemons.Pokemons, pokemons.Pokemon, identifier.PokemonId, uint16]
}

type IAbilityRepository interface {
	ISingleRecordFinder[abilities.Ability, identifier.AbilityId, uint16]
	IAllRecordRepository[abilities.Abilities, abilities.Ability, identifier.AbilityId, uint16]
	IPokemonStatsRepository[abilities.Abilities, abilities.Ability, identifier.AbilityId]
}

type IMoveRepository interface {
	ISingleRecordFinder[moves.Move, identifier.MoveId, uint16]
	IAllRecordRepository[moves.Moves, moves.Move, identifier.MoveId, uint16]
	IPokemonStatsRepository[moves.Moves, moves.Move, identifier.MoveId]
}

type IHeldItemRepository interface {
	ISingleRecordFinder[items.HeldItem, identifier.HeldItemId, uint16]
	IAllRecordRepository[items.HeldItems, items.HeldItem, identifier.HeldItemId, uint16]
}

type IPartyTagRepository interface {
	ISingleRecordFinder[parties.PartyTag, identifier.PartyTagId, uint64]
	IAllRecordRepository[parties.PartyTags, parties.PartyTag, identifier.PartyTagId, uint64]
	IWritableRepository[parties.PartyTag, identifier.PartyTagId, uint64]
}

type IPartyRepository interface {
	ISingleRecordFinder[parties.Party, identifier.PartyId, uint64]
	IWritableRepository[parties.Party, identifier.PartyId, uint64]
}

type IPartyBattleResultRepository interface {
	IWritableRepository[parties.PartyBattleResult, identifier.PartyBattleResultId, uint64]
}

type IBattleRecordRepository interface {
	ISingleRecordFinder[battles.BattleRecord, identifier.BattleRecordId, uint64]
	IWritableRepository[battles.BattleRecord, identifier.BattleRecordId, uint64]
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
	IWritableRepository[battles.BattleOpponentParty, identifier.BattleOpponentPartyId, uint64]
	FindParty(*battles.BattleOpponentParty) (*battles.BattleOpponentParty, error)
}

type ITrainedPokemonRepository interface {
	IWritableRepository[trainings.TrainedPokemon, identifier.TrainedPokemonId, uint64]
}

// Todo: Don't want to depend gorm. wrap?
type ITrainedPokemonAdjustmentRepository interface {
	Find(*gorm.DB, trainings.TrainedPokemonAdjustment) (*trainings.TrainedPokemonAdjustment, error)
	IAllRecordRepository[trainings.TrainedPokemonAdjustments, trainings.TrainedPokemonAdjustment, identifier.TrainedAdjustmentId, uint64]
	IWritableRepository[trainings.TrainedPokemonAdjustment, identifier.TrainedAdjustmentId, uint64]
}

type ITrainedPokemonAttackRepository interface {
	IWritableRepository[trainings.TrainedPokemonAttackTarget, identifier.TrainedAttackTargetId, uint64]
}

type ITrainedPokemonDeffenceRepository interface {
	IWritableRepository[trainings.TrainedPokemonDefenseTarget, identifier.TrainedDefenseTargetId, uint64]
}

type IUserRepository interface {
	FindById(id uint64) (*users.User, error)
	FindByName(targetName string, filterFields ...string) (*users.User, error)
	IWritableRepository[users.User, identifier.UserId, uint64]
}
