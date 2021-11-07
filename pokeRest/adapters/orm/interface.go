package orm

import "gorm.io/gorm"

type IDbClientProvider interface {
	GetDbAccessor() *IDbClient
}

type IDbClient interface {
	Db() *gorm.DB
	Close()
	Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB
}
