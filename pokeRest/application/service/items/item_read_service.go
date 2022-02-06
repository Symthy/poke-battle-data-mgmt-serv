package items

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/model/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type is = items.HeldItems
type i = items.HeldItem

type ItemReadService struct {
	repo repository.IItemRepository
	service.EntityAllFinder[is, i]
}

func NewItemReadService(repo repository.IItemRepository) ItemReadService {
	serv := ItemReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[is, i](repo)
	return serv
}

// UC: 持ち物一覧取得 EntityAllFinder
