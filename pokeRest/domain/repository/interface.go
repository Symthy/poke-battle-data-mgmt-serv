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
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"gorm.io/gorm"
)

// common
type ISingleRecordFinder[T entity.IDomain] interface {
	FindById(id uint) (*T, error)
}

type IAllRecordRepository[L entity.IDomains[T], T entity.IDomain] interface {
	FindAll(page int, pageSize int) (*L, error)
}

type IPokemonStatsRepository[L entity.IDomains[T], T entity.IDomain] interface {
	FindOfPokemon(pokemonId uint) (*L, error)
}

type IWritableRepository[TD entity.IDomain] interface {
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
	ISingleRecordFinder[pokemons.Pokemon]
	IAllRecordRepository[pokemons.Pokemons, pokemons.Pokemon]
	IWritableRepository[pokemons.Pokemon]
}

type IAbilityRepository interface {
	ISingleRecordFinder[abilities.Ability]
	IAllRecordRepository[abilities.Abilities, abilities.Ability]
	IPokemonStatsRepository[abilities.Abilities, abilities.Ability]
}

type IMoveRepository interface {
	ISingleRecordFinder[moves.Move]
	IAllRecordRepository[moves.Moves, moves.Move]
	IPokemonStatsRepository[moves.Moves, moves.Move]
}

type IHeldItemRepository interface {
	ISingleRecordFinder[items.HeldItem]
	IAllRecordRepository[items.HeldItems, items.HeldItem]
}

type IPartyTagRepository interface {
	ISingleRecordFinder[parties.PartyTag]
	IAllRecordRepository[parties.PartyTags, parties.PartyTag]
	IWritableRepository[parties.PartyTag]
}

type IPartyRepository interface {
	ISingleRecordFinder[parties.Party]
	IWritableRepository[parties.Party]
}

type IPartyBattleResultRepository interface {
	IWritableRepository[parties.PartyBattleResult]
}

// Todo: Don't want to depend gorm. wrap?
type IBattleRecordRepository interface {
	ISingleRecordFinder[battles.BattleRecord]
	IWritableRepository[battles.BattleRecord]
	CreateRecord(*gorm.DB, battles.BattleRecord) (*battles.BattleRecord, error)
	UpdateRecord(*gorm.DB, battles.BattleRecord) (*battles.BattleRecord, error)
	DeleteRecord(*gorm.DB, uint) (*battles.BattleRecord, error)
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
	IWritableRepository[battles.BattleOpponentParty]
	FindParty(value.PartyPokemonIds) (*battles.BattleOpponentParty, error)
}

// Todo: Don't want to depend gorm. wrap?
type ITrainedPokemonRepository interface {
	IWritableRepository[trainings.TrainedPokemonParam]
	CreateRecord(*gorm.DB, trainings.TrainedPokemonParam) (*trainings.TrainedPokemonParam, error)
	UpdateRecord(*gorm.DB, trainings.TrainedPokemonParam) (*trainings.TrainedPokemonParam, error)
	DeleteRecord(*gorm.DB, uint) (*trainings.TrainedPokemonParam, error)
}
type ITrainedPokemonTransactionalRepository interface {
	ISingleRecordFinder[trainings.TrainedPokemonParam]
	ITrainedPokemonRepository
	ITransactionalRepository
}

type ITrainedPokemonAdjustmentRepository interface {
	Find(trainings.TrainedPokemonAdjustment) (*trainings.TrainedPokemonAdjustment, error)
	IAllRecordRepository[trainings.TrainedPokemonAdjustments, trainings.TrainedPokemonAdjustment]
	IWritableRepository[trainings.TrainedPokemonAdjustment]
}

type ITrainedPokemonAttackRepository interface {
	IWritableRepository[trainings.TrainedPokemonAttackTarget]
}

type ITrainedPokemonDeffenceRepository interface {
	IWritableRepository[trainings.TrainedPokemonDefenceTarget]
}

type IUserRepository interface {
	FindById(id uint) (*users.User, error)
	FindByName(targetName string, filterFields ...string) (*users.User, error)
	IWritableRepository[users.User]
}
