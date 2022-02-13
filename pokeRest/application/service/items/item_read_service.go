package items

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type is = items.HeldItems
type i = items.HeldItem

type ItemReadService struct {
	repo repository.IHeldItemRepository
	service.EntityAllFinder[is, i]
}

func NewItemReadService(repo repository.IHeldItemRepository) ItemReadService {
	serv := ItemReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[is, i](repo)
	return serv
}

// UC: 持ち物取得
func (s ItemReadService) FindHeldItem(itemId uint) (*items.HeldItem, error) {
	return s.repo.FindById(itemId)
}

// UC: 持ち物一覧取得
// FindAll <- EntityAllFinder
