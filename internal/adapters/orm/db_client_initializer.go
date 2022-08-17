package orm

import (
	"github.com/Symthy/PokeRest/internal/config"
	"github.com/Symthy/PokeRest/internal/logs"
)

type DbClientInitializer struct {
	dbConfig          config.DbConfig
	gormLoggerfactory logs.DbGormLoggerFactory
}

func NewDbClientInitializer(dbConfig config.DbConfig, dbLoggerFactory logs.IDbLoggerFactory) DbClientInitializer {
	return DbClientInitializer{
		dbConfig:          dbConfig,
		gormLoggerfactory: logs.NewDbGormLoggerFactory(dbLoggerFactory),
	}
}

func (i DbClientInitializer) InitializeDbClient() *GormDbClient {
	return NewGormDbClient(i.dbConfig, i.gormLoggerfactory.BuildGormLogger())
}
