package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// implements ILogField
type LogField struct {
	key   string
	value string
}

func (f LogField) GetKey() string {
	return f.key
}
func (f LogField) GetValue() string {
	return f.value
}

type ServerGlobalLogger struct {
	logger *zap.Logger
}

// public for testing
func InitGlobalServerLogger(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
}

func GetGlobalServerLogger() IServerGlobalLogger {
	return ServerGlobalLogger{logger: zap.L()}
}

func (l ServerGlobalLogger) Debug(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Debug(msg, zapField...)
}

func (l ServerGlobalLogger) Info(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Info(msg, zapField...)
}

func (l ServerGlobalLogger) Warn(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Warn(msg, zapField...)
}

func (l ServerGlobalLogger) Error(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Error(msg, zapField...)
}

func (l ServerGlobalLogger) Fatal(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Fatal(msg, zapField...)
}

func convertZapField(logFields ...ILogField) []zapcore.Field {
	zapFields := make([]zapcore.Field, len(logFields))
	for i, logField := range logFields {
		zapFields[i] = zap.String(logField.GetKey(), logField.GetValue())
	}
	return zapFields
}
