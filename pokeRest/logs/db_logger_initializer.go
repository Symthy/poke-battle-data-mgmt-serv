package logs

import (
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

// Todo: refactor db initialize process
type IDbLoggerInitializer interface {
	BuildGormLogger() logger.Interface
}

type DbLoggerInitializer struct {
	factory IDbLoggerFactory
}

func NewDbLoggerInitializer(factory IDbLoggerFactory) IDbLoggerInitializer {
	return DbLoggerInitializer{factory: factory}
}

func (i DbLoggerInitializer) BuildGormLogger() logger.Interface {
	zapLogger := buildZapLogger(i.factory.BuildRotateDbLogger())
	zapGormLogger := zapgorm2.New(zapLogger)
	zapGormLogger.SetAsDefault()
	return zapGormLogger
}
