package items

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/items"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
)

type is = items.HeldItems
type i = items.HeldItem
type ii = identifier.HeldItemId

type ItemReadService struct {
	repo repository.IHeldItemRepository
	service.EntityAllFinder[is, i, ii, uint16]
}

func NewItemReadService(repo repository.IHeldItemRepository) ItemReadService {
	serv := ItemReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[is, i, ii, uint16](repo)
	return serv
}

// UC: 持ち物取得
func (s ItemReadService) FindHeldItem(itemId uint64) (*items.HeldItem, error) {
	// Todo: validate id upper limit
	return s.repo.FindById(uint16(itemId))
}

// UC: 持ち物一覧取得
// FindAll <- EntityAllFinder
