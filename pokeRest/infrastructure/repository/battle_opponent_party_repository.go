package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/battles"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
	"gorm.io/gorm"
)

var _ repository.IBattleOpponentPartyRepository = (*BattleOpponentPartyRepository)(nil)

var emptyBattleOpponentPartySchemaBuilder = func() schema.BattleOpponentParty {
	return schema.BattleOpponentParty{}
}

type BattleOpponentPartyRepository struct {
	BaseWriteRepository[schema.BattleOpponentParty, battles.BattleOpponentParty]
	dbClient orm.IDbClient
}

func NewBattleOpponentParty(dbClient orm.IDbClient) *BattleOpponentPartyRepository {
	return &BattleOpponentPartyRepository{
		BaseWriteRepository: BaseWriteRepository[schema.BattleOpponentParty, battles.BattleOpponentParty]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyBattleOpponentPartySchemaBuilder,
			schemaConverter:    dto.ToSchemaBattleOpponentParty,
		},
		dbClient: dbClient,
	}
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository

func (rep BattleOpponentPartyRepository) FindParty(partyMember value.PartyPokemonIds) (*battles.BattleOpponentParty, error) {
	db := rep.dbClient.Db()
	searchParty := buildSearchPartySchema(partyMember)
	schema := schema.BattleOpponentParty{}
	tx := db.Where(searchParty).First(&schema)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if tx != nil {
		return nil, tx.Error
	}
	domain := schema.ConvertToDomain()
	return &domain, nil
}

func buildSearchPartySchema(partyMember value.PartyPokemonIds) schema.BattleOpponentParty {
	return schema.BattleOpponentParty{
		OpponentPokemonId1: partyMember.Get(0),
		OpponentPokemonId2: partyMember.Get(1),
		OpponentPokemonId3: partyMember.Get(2),
		OpponentPokemonId4: partyMember.Get(3),
		OpponentPokemonId5: partyMember.Get(4),
		OpponentPokemonId6: partyMember.Get(5),
	}
}
