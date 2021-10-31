package db

import "gorm.io/gorm"

type IDbAccessorProvider interface {
	GetDbAccessor() *IDbAccessor
}

type IDbAccessor interface {
	GetDb() *gorm.DB
	Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB
}
