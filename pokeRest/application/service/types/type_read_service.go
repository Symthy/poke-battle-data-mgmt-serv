package types

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/model/types"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
)

type TypeReadService struct {
	repository repository.ITypeRepository
}

// UC: タイプ相性表取得
func (s TypeReadService) FindTypeCompatibility() (*types.TypeCompatibility, error) {
	return nil, nil
}
