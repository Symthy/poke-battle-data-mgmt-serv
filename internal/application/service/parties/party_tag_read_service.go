package parties

import (
	"github.com/Symthy/PokeRest/internal/application/service"
	"github.com/Symthy/PokeRest/internal/domain/entity/parties"
	"github.com/Symthy/PokeRest/internal/domain/repository"
	"github.com/Symthy/PokeRest/internal/domain/value/identifier"
)

type pts = parties.PartyTags
type pt = parties.PartyTag
type pti = identifier.PartyTagId

type PartyTagReadService struct {
	repo repository.IPartyTagRepository
	service.EntityAllFinder[pts, pt, pti, uint64]
}

func NewPartyTagReadService(repo repository.IPartyTagRepository) PartyTagReadService {
	serv := PartyTagReadService{repo: repo}
	serv.EntityAllFinder = service.NewEntityAllFinder[pts, pt, pti, uint64](repo)
	return serv
}

// UC: 持ち物取得
func (s PartyTagReadService) FindPartyTag(itemId uint64) (*parties.PartyTag, error) {
	return s.repo.FindById(itemId)
}

// UC: 持ち物一覧取得
// FindAll <- EntityAllFinder
