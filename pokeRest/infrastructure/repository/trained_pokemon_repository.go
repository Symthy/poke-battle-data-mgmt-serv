package repository

import (
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm"
	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/domain/entity/trainings"
	"github.com/Symthy/PokeRest/pokeRest/domain/repository"
	"github.com/Symthy/PokeRest/pokeRest/infrastructure/repository/conv"
	"gorm.io/gorm"
)

var _ repository.ITrainedPokemonRepository = (*TrainedPokemonRepository)(nil)

var emptyTrainedPokemonSchemaBuilder = func() *schema.TrainedPokemon { return &schema.TrainedPokemon{} }

type TrainedPokemonRepository struct {
	//BaseWriteRepository[schema.TrainedPokemon, trainings.TrainedPokemonParam, identifier.TrainedPokemonId]
	adjustmentRepo TrainedPokemonAdjustmentRepository
	dbClient       orm.IDbClient
}

func NewTrainedPokemonRepository(dbClient orm.IDbClient) *TrainedPokemonRepository {
	return &TrainedPokemonRepository{
		// BaseWriteRepository: BaseWriteRepository[schema.TrainedPokemon, trainings.TrainedPokemonParam, identifier.TrainedPokemonId]{
		// 	dbClient:           dbClient,
		// 	emptySchemaBuilder: emptyTrainedPokemonSchemaBuilder,
		// 	toSchemaConverter:  conv.ToSchemaTrainedPokemon,
		// 	toDomainConverter:  conv.ToDomainTrainedPokemonParam,
		// },
		adjustmentRepo: *NewTrainedPokemonAdjustmentRepository(dbClient),
		dbClient:       dbClient,
	}
}

func (rep TrainedPokemonRepository) Create(domain *trainings.TrainedPokemon) (*trainings.TrainedPokemon, error) {
	schema := conv.ToSchemaTrainedPokemon(domain)
	db := rep.dbClient.Db()
	db.Transaction(func(tx *gorm.DB) error {
		adjustment, err := rep.adjustmentRepo.Find(tx, domain.TrainedPokemonAdjustment)
		if err != nil {
			return err
		}
		if adjustment == nil {
			adjustmentSchema := conv.ToSchemaTrainedPokemonAdjustment(adjustment)
			schema.TrainedPokemonAdjustment = *adjustmentSchema
		} else {
			schema.AdjustmentID = adjustment.Id().Value()
		}

		ret := tx.Create(&schema)
		if ret.Error != nil {
			return ret.Error
		}
		return nil
	})
	return conv.ToDomainTrainedPokemon(schema)
}

func (rep TrainedPokemonRepository) Update(domain *trainings.TrainedPokemon) (*trainings.TrainedPokemon, error) {
	trainedPoke := conv.ToSchemaTrainedPokemon(domain)
	db := rep.dbClient.Db()
	updated := emptyTrainedPokemonSchemaBuilder()
	db.Transaction(func(tx *gorm.DB) error {
		adjustment, err := rep.adjustmentRepo.Find(tx, domain.TrainedPokemonAdjustment)
		if err != nil {
			return err
		}
		if adjustment == nil {
			adjustmentDomain, err := rep.adjustmentRepo.CreateDelegate(tx, &domain.TrainedPokemonAdjustment)
			if err != nil {
				return err
			}
			trainedPoke.AdjustmentID = adjustmentDomain.Id().Value()
		} else {
			trainedPoke.AdjustmentID = adjustment.Id().Value()
		}

		ret := tx.Model(&updated).Updates(trainedPoke)
		if ret.Error != nil {
			return ret.Error
		}
		return nil
	})
	return conv.ToDomainTrainedPokemon(updated)
}

func (rep TrainedPokemonRepository) Delete(id uint64) (*trainings.TrainedPokemon, error) {
	db := rep.dbClient.Db()
	deleted := emptyTrainedPokemonSchemaBuilder()
	db.Transaction(func(tx *gorm.DB) error {
		if ret := db.First(schema.TrainedPokemon{}, id); tx.Error != nil {
			return ret.Error
		}
		ret := db.Delete(&deleted, id)
		if ret.Error != nil {
			return ret.Error
		}
		return nil
	})
	// Todo:
	// delete used trained pokemon adjustment if ref num is zero
	isRefZero := false
	if isRefZero {
		_, err := rep.adjustmentRepo.Delete(deleted.AdjustmentID)
		if err != nil {
			return nil, err
		}
	}
	return conv.ToDomainTrainedPokemon(deleted)
}
