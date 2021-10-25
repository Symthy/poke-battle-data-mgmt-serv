package schema

import "github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/enum"

type Move struct {
	ID          uint `gorm:"primaryKey; autoIncrement"`
	Name        string
	Species     enum.MoveSpecies `sql:"type:species"`
	Power       int
	Accuracy    int
	PP          int
	IsContained bool `gorm:"default:false"`
	IsCanGuard  bool `gorm:"default:true"`
}

func (Move) TableName() string {
	return "move"
}
