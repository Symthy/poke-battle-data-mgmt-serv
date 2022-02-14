package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
	"gorm.io/gorm"
)

var _ repository.ITrainedPokemonAdjustmentRepository = (*TrainedPokemonAdjustmentRepository)(nil)

var emptyAdjustmentSchemaBuilder = func() schema.TrainedPokemonAdjustment { return schema.TrainedPokemonAdjustment{} }

type adjustmentDomain = trainings.TrainedPokemonAdjustment

type TrainedPokemonAdjustmentRepository struct {
	BaseWriteRepository[schema.TrainedPokemonAdjustment, adjustmentDomain]
	dbClient orm.IDbClient
}

func NewTrainedPokemonAdjustmentRepository(dbClient orm.IDbClient) *TrainedPokemonAdjustmentRepository {
	return &TrainedPokemonAdjustmentRepository{
		BaseWriteRepository: BaseWriteRepository[schema.TrainedPokemonAdjustment, adjustmentDomain]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyAdjustmentSchemaBuilder,
			schemaConverter:    dto.ToSchemaTrainedPokemonAdjustment,
		},
		dbClient: dbClient,
	}
}

func (repo TrainedPokemonAdjustmentRepository) Find(adjustment adjustmentDomain) (*adjustmentDomain, error) {
	db := repo.dbClient.Db()
	schema := emptyAdjustmentSchemaBuilder()
	tx := db.Where(adjustment).First(&schema)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	s := schema.ConvertToDomain()
	return &s, nil
}

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
