package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/abilities"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/moves"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/pokemons"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/tags"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/types"
)

type IBasicRepository interface {
	FindById(id uint)
}

type IEntityAllRepository[L model.IDomains[T], T model.IDomain] interface {
	FindAll(page int, pageSize int) (*L, error)
}

type IValueOfPokemonRepository[L model.IDomains[T], T model.IDomain] interface {
	FindOfPokemon(pokemonId uint) (*L, error)
}

type IPokemonRepository interface {
	IEntityAllRepository[pokemons.Pokemons, pokemons.Pokemon]
	FindById(id uint) (*pokemons.Pokemon, error)
	Create(pokemon *pokemons.Pokemon) (*pokemons.Pokemon, error)
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
}

type IPartyRepository interface {
	Save()
	Update()
	Delete()
}

type IPartySeasonResultRepository interface {
	Save()
	Update()
	Delete()
}

type IBattleRecordRepository interface {
	Save()
	Update()
	Delete()
}

type IBattleOpponentPartyRepository interface {
	Save()
	Update()
	Delete()
}

type ITrainedPokemonRepository interface {
	Save()
	Update()
	Delete()
}

type ITrainedPokemonBaseRepository interface {
	Find()
	Save()
	Delete()
}

type IUserRepository interface {
	FindById(id uint) (*model.User, error)
	FindByName(targetName string, filterFields ...string) (*model.User, error)
	Save(user model.User) (*model.User, error)
}
