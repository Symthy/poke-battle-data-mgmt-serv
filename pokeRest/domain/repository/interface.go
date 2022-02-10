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
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/users"
)

// common
type IBasicRepository interface {
	FindById(id uint)
}

type IEntityAllRepository[L entity.IDomains[T], T entity.IDomain] interface {
	FindAll(page int, pageSize int) (*L, error)
}

type IPokemonStatsRepository[L entity.IDomains[T], T entity.IDomain] interface {
	FindOfPokemon(pokemonId uint) (*L, error)
}

type IWritableRepository[TD entity.IDomain] interface {
	Create(TD) (*TD, error)
	Update(TD) (*TD, error)
	Delete(TD) (*TD, error)
}

// special
type IPokemonRepository interface {
	FindById(id uint) (*pokemons.Pokemon, error)
	IEntityAllRepository[pokemons.Pokemons, pokemons.Pokemon]
	IWritableRepository[pokemons.Pokemon]
}

type IAbilityRepository interface {
	IEntityAllRepository[abilities.Abilities, abilities.Ability]
	IPokemonStatsRepository[abilities.Abilities, abilities.Ability]
}

type IMoveRepository interface {
	IEntityAllRepository[moves.Moves, moves.Move]
	IPokemonStatsRepository[moves.Moves, moves.Move]
}

type IItemRepository interface {
	IEntityAllRepository[items.HeldItems, items.HeldItem]
}

type ITypeRepository interface {
	BuildTypeCompatibility() (*types.TypeCompatibility, error)
}

type IPartyTagRepository interface {
	IEntityAllRepository[parties.PartyTags, parties.PartyTag]
	IWritableRepository[parties.PartyTag]
}

type IPartyRepository interface {
	IWritableRepository[parties.Party]
}

type IPartyBattleResultRepository interface {
	IWritableRepository[parties.PartyBattleResult]
}

type IBattleRecordRepository interface {
	IWritableRepository[battles.BattleRecord]
}
type IBattleOpponentPartyRepository interface {
	IWritableRepository[battles.BattleOpponentParty]
}

type ITrainedPokemonRepository interface {
	IWritableRepository[trainings.TrainedPokemon]
}

type ITrainedPokemonAttackRepository interface {
	IWritableRepository[trainings.TrainedPokemonAttackTarget]
}

type ITrainedPokemonDeffenceRepository interface {
	IWritableRepository[trainings.TrainedPokemonDeffenceTarget]
}

type IUserRepository interface {
	FindById(id uint) (*users.User, error)
	FindByName(targetName string, filterFields ...string) (*users.User, error)
	IWritableRepository[users.User]
}
