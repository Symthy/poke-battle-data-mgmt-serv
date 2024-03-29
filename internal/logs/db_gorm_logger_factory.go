package logs

import (
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type DbGormLoggerFactory struct {
	factory IDbLoggerFactory
}

func NewDbGormLoggerFactory(factory IDbLoggerFactory) DbGormLoggerFactory {
	return DbGormLoggerFactory{factory: factory}
}

func (i DbGormLoggerFactory) BuildGormLogger() logger.Interface {
	zapLogger := BuildZapLogger(i.factory.BuildBaseDbLogger(), i.factory.ResolveLogLevel())
	zapGormLogger := zapgorm2.New(zapLogger)
	zapGormLogger.SetAsDefault()
	return zapGormLogger
}
