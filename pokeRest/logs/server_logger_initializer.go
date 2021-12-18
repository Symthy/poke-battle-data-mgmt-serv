package logs

import "go.uber.org/zap"

type IServerLoggerInitializer interface {
	AcceptLogger()
}

type ServerGlobalLoggerInitializer struct {
	factory IServerLoggerFactory
}

func NewServerGlobalLoggerInitializer(factory IServerLoggerFactory) IServerLoggerInitializer {
	return ServerGlobalLoggerInitializer{factory: factory}
}

func (i ServerGlobalLoggerInitializer) AcceptLogger() {
	rotateLogger := i.factory.BuildRotateServerLogger()
	zapLogger := buildZapLogger(rotateLogger)
	zap.ReplaceGlobals(zapLogger)
}
