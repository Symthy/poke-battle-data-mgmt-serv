package logs

import (
	"io"

	"github.com/Symthy/PokeRest/internal/common"
)

type ILogLevelResolver interface {
	ResolveLogLevel() common.Level
}

type IAccessLoggerFactory interface {
	ILogLevelResolver
	BuildBaseAccessLogger() io.Writer
}

type IServerLoggerFactory interface {
	ILogLevelResolver
	BuildBaseServerLogger() io.Writer
}

type IDbLoggerFactory interface {
	ILogLevelResolver
	BuildBaseDbLogger() io.Writer
}

type IServerGlobalLogger interface {
	Debug(msg string, fields ...ILogField)
	Info(msg string, fields ...ILogField)
	Warn(msg string, fields ...ILogField)
	Error(msg string, fields ...ILogField)
	Fatal(msg string, fields ...ILogField)
}

type ILogField interface {
	GetKey() string
	GetValue() string
}
