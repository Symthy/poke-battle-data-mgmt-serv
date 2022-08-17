package logs

import (
	"io"

	"github.com/labstack/echo/v4"
)

type IAccessLoggerInitializer interface {
	AcceptLogger(e *echo.Echo)
}

type AccessLoggerMiddlewareInitializer struct {
	factory IAccessLoggerFactory
}

func NewAccessLoggerMiddlewareInitializer(factory IAccessLoggerFactory) IAccessLoggerInitializer {
	return AccessLoggerMiddlewareInitializer{factory: factory}
}

func (i AccessLoggerMiddlewareInitializer) AcceptLogger(e *echo.Echo) {
	rotateLogger := i.factory.BuildBaseAccessLogger()
	zapLoggerEchoMiddleware := i.buildLoggerMiddleware(rotateLogger)
	e.Use(zapLoggerEchoMiddleware)
}

func (i AccessLoggerMiddlewareInitializer) buildLoggerMiddleware(writer io.Writer) echo.MiddlewareFunc {
	zapLoggerEchoMiddleware := zapEchoLogger(BuildZapLogger(writer, i.factory.ResolveLogLevel()))
	return zapLoggerEchoMiddleware
}
