package value

import "github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"

type PokemonMoveIdSet struct {
	moveId1 identifier.MoveId
	moveId2 identifier.MoveId
	moveId3 identifier.MoveId
	moveId4 identifier.MoveId
}

func NewPokemonMoveIdSet(ids ...uint) (*PokemonMoveIdSet, error) {
	var id1 uint = 0
	var id2 uint = 0
	var id3 uint = 0
	var id4 uint = 0
	if len(ids) >= 1 {
		id1 = ids[0]
	}
	if len(ids) >= 2 {
		id2 = ids[1]
	}
	if len(ids) >= 3 {
		id3 = ids[2]
	}
	if len(ids) >= 4 {
		id4 = ids[3]
	}

	moveId1, err := identifier.NewMoveId(id1)
	if err != nil {
		return nil, err
	}
	moveId2, err := identifier.NewMoveId(id2)
	if err != nil {
		return nil, err
	}
	moveId3, err := identifier.NewMoveId(id3)
	if err != nil {
		return nil, err
	}
	moveId4, err := identifier.NewMoveId(id4)
	if err != nil {
		return nil, err
	}
	return &PokemonMoveIdSet{*moveId1, *moveId2, *moveId3, *moveId4}, nil
}

func (m PokemonMoveIdSet) GetMoveIds() (identifier.MoveId, identifier.MoveId, identifier.MoveId, identifier.MoveId) {
	return m.moveId1, m.moveId2, m.moveId3, m.moveId4
}
