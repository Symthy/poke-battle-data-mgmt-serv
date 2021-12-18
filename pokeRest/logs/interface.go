package logs

import (
	"io"
)

type IAccessLoggerFactory interface {
	BuildRotateAccessLogger() io.Writer
}

type IServerLoggerFactory interface {
	BuildRotateServerLogger() io.Writer
}

type IDbLoggerFactory interface {
	BuildRotateDbLogger() io.Writer
}
