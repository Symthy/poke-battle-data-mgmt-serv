package errors

import "fmt"

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
	fatal Level = "fatal"
)

func (l Level) String() string {
	return string(l)
}

type AppError struct {
	err        error
	message    string
	level      Level
	errCode    int
	httpStatus int
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%s - %04d] %s", e.level, e.errCode, e.message)
}

func (e AppError) log() {
	fmt.Printf("%+v\n", e.err)
}
