package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/application/service"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/parties"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type pts = parties.PartyTags
type pt = parties.PartyTag

type PartyTagReadService struct {
	repo repository.IPartyTagRepository
	service.EntityAllFinder[pts, pt]
}

func NewPartyTagReadService(repo repository.IPartyTagRepository) PartyTagReadService {
	serv := PartyTagReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[pts, pt](repo)
	return serv
}

// UC: 持ち物取得
func (s PartyTagReadService) FindPartyTag(itemId uint) (*parties.PartyTag, error) {
	return s.repo.FindById(itemId)
}

// UC: 持ち物一覧取得
// FindAll <- EntityAllFinder
