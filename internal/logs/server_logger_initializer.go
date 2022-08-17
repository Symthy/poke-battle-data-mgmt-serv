package logs

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
	rotateLogger := i.factory.BuildBaseServerLogger()
	zapLogger := BuildZapLogger(rotateLogger, i.factory.ResolveLogLevel())
	InitGlobalServerLogger(zapLogger)
}
