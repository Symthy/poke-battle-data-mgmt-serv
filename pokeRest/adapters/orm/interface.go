package orm

import "gorm.io/gorm"

type IDbClientProvider interface {
	GetDbAccessor() *IDbClient
}

type IDbClient interface {
	GetDb() *gorm.DB
	Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB
}
