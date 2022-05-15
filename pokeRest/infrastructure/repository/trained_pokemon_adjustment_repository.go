package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/domain/value/identifier"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
	"gorm.io/gorm"
)

var _ repository.ITrainedPokemonAdjustmentRepository = (*TrainedPokemonAdjustmentRepository)(nil)

var emptyAdjustmentSchemaBuilder = func() schema.TrainedPokemonAdjustment { return schema.TrainedPokemonAdjustment{} }

type adjustmentDomain = trainings.TrainedPokemonAdjustment
type adjustmentId = identifier.TrainedAdjustmentId

type TrainedPokemonAdjustmentRepository struct {
	BaseReadRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, trainings.TrainedPokemonAdjustments, adjustmentId, uint64]
	BaseWriteRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, adjustmentId, uint64]
	dbClient orm.IDbClient
}

func NewTrainedPokemonAdjustmentRepository(dbClient orm.IDbClient) *TrainedPokemonAdjustmentRepository {
	return &TrainedPokemonAdjustmentRepository{
		BaseReadRepository: BaseReadRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, trainings.TrainedPokemonAdjustments, identifier.TrainedAdjustmentId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyAdjustmentSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaTrainedPokemonAdjustment,
		},
		BaseWriteRepository: BaseWriteRepository[schema.TrainedPokemonAdjustment, adjustmentDomain, adjustmentId, uint64]{
			dbClient:           dbClient,
			emptySchemaBuilder: emptyAdjustmentSchemaBuilder,
			toSchemaConverter:  conv.ToSchemaTrainedPokemonAdjustment,
			toDomainConverter:  conv.ToDomainTrainedPokemonAdjustment,
		},
		dbClient: dbClient,
	}
}

func (rep TrainedPokemonAdjustmentRepository) Find(db *gorm.DB, adjustment adjustmentDomain) (*adjustmentDomain, error) {
	schema := emptyAdjustmentSchemaBuilder()
	tx := db.Where(adjustment).First(&schema)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rep.toDomainConverter(schema)
}

// FindAll <- BaseReadRepository

// Create <- BaseWriteRepository
// Update <- BaseWriteRepository
// Delete <- BaseWriteRepository
