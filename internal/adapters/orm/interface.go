package orm

import (
	"gorm.io/gorm"
)

type IDbClientProvider interface {
	GetDbAccessor() *IDbClient
}

type IDbClient interface {
	Db() *gorm.DB
	Close()
	Paginate(page uint32, pageSize uint16) func(db *gorm.DB) *gorm.DB
}
