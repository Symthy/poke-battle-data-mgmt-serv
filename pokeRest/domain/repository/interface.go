package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/tags"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/users"
)

// common
type IBasicRepository interface {
	FindById(id uint)
}

type IEntityAllRepository[L model.IDomains[T], T model.IDomain] interface {
	FindAll(page int, pageSize int) (*L, error)
}

type IValueOfPokemonRepository[L model.IDomains[T], T model.IDomain] interface {
	FindOfPokemon(pokemonId uint) (*L, error)
}

type IWritableRepository[TD model.IDomain] interface {
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
	IValueOfPokemonRepository[abilities.Abilities, abilities.Ability]
}

type IMoveRepository interface {
	IEntityAllRepository[moves.Moves, moves.Move]
	IValueOfPokemonRepository[moves.Moves, moves.Move]
}

type IItemRepository interface {
	IEntityAllRepository[items.HeldItems, items.HeldItem]
}

type ITypeRepository interface {
	BuildTypeCompatibility() (*types.TypeCompatibility, error)
}

type ITagRepository interface {
	IEntityAllRepository[tags.Tags, tags.Tag]
	IWritableRepository[tags.Tag]
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
	IWritableRepository[pokemons.TrainedPokemon]
}

type ITrainedPokemonMoveSetRepository interface {
	Find() (*pokemons.TrainedPokemonMoveSet, error)
	IWritableRepository[pokemons.TrainedPokemonMoveSet]
}

type ITrainedPokemonAttackRepository interface {
	IWritableRepository[pokemons.TrainedPokemonAttackTarget]
}

type ITrainedPokemonDeffenceRepository interface {
	IWritableRepository[pokemons.TrainedPokemonDeffenceTarget]
}

type IUserRepository interface {
	FindById(id uint) (*users.User, error)
	FindByName(targetName string, filterFields ...string) (*users.User, error)
	IWritableRepository[users.User]
}
