package parties

import (
	"github.com/Symthy/PokeRest/pokeRest/domain/entity"
)

var _ entity.IDomain = (*Party)(nil)

// Todo: field change to value model
type Party struct {
	id                uint
	name              string
	battleFormat      BattleFormat
	partyTagIds       []uint
	trainedPokemonIds []uint
	isPrivate         bool
	userId            *uint
}

func NewParty(id uint) Party {
	return Party{
		id: id,
	}
}

func NewPartyOfUnregistered(name string, battleFormat string, partyTagIds []uint,
	trainedPokemonIds []uint, isPrivate bool, userId uint) Party {
	// Todo: validate
	return Party{
		id:                0,
		name:              name,
		battleFormat:      resolveBattleFormat(battleFormat),
		partyTagIds:       partyTagIds,
		trainedPokemonIds: trainedPokemonIds,
		isPrivate:         isPrivate,
		userId:            &userId,
	}
}

func NewPartyForUpdated(id uint, name string, battleFormat string, partyTagIds []uint,
	trainedPokemonIds []uint, isPrivate bool) Party {
	// Todo: validate
	return Party{
		id:                id,
		name:              name,
		battleFormat:      resolveBattleFormat(battleFormat),
		partyTagIds:       partyTagIds,
		trainedPokemonIds: trainedPokemonIds,
		isPrivate:         isPrivate,
		userId:            nil,
	}
}

func (p Party) Id() uint {
	return p.id
}

func (p Party) Name() string {
	return p.name
}

func (p Party) BattleFormat() BattleFormat {
	return p.battleFormat
}

func (p Party) PartyTagIds() []uint {
	return p.partyTagIds
}

func (p Party) TrainedPokemonIds() []uint {
	return p.trainedPokemonIds
}

func (p Party) IsPrivate() bool {
	return p.isPrivate
}

func (p Party) UserId() *uint {
	return p.userId
}

type BattleFormat string

const (
	Single BattleFormat = "Single"
	Double BattleFormat = "Double"
)

func resolveBattleFormat(value string) BattleFormat {
	if value == "Double" {
		return Double
	}
	return Single
}
