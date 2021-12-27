package logs

import (
	"io"
)

type IAccessLoggerFactory interface {
	BuildBaseAccessLogger() io.Writer
}

type IServerLoggerFactory interface {
	BuildBaseServerLogger() io.Writer
}

type IDbLoggerFactory interface {
	BuildBaseDbLogger() io.Writer
}

type IGlobalServerLogger interface {
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
