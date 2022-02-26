package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/dto"
	"gorm.io/gorm"
)

var _ repository.ITrainedPokemonAdjustmentRepository = (*TrainedPokemonAdjustmentRepository)(nil)

var emptyAdjustmentSchemaBuilder = func() schema.TrainedPokemonAdjustment { return schema.TrainedPokemonAdjustment{} }

type adjustmentDomain = trainings.TrainedPokemonAdjustment
type adjustmentId = identifier.TrainedAdjustmentId

type TrainedPokemonAdjustmentRepository struct {
	BaseReadRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, trainings.TrainedPokemonAdjustments, adjustmentId]
	BaseWriteRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, adjustmentId]
	dbClient orm.IDbClient
}

func NewTrainedPokemonAdjustmentRepository(dbClient orm.IDbClient) *TrainedPokemonAdjustmentRepository {
	return &TrainedPokemonAdjustmentRepository{
		BaseReadRepository: BaseReadRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, trainings.TrainedPokemonAdjustments, identifier.TrainedAdjustmentId]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyAdjustmentSchemaBuilder,
			schemaConverter:    dto.ToSchemaTrainedPokemonAdjustment,
		},
		BaseWriteRepository: BaseWriteRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, adjustmentId]{
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

// FindAll <- BaseReadRepository

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
