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

type GlobalServerZapLogger struct {
	logger *zap.Logger
}

func InitGlobalServerLogger(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
}

func GetGlobalServerLogger() IGlobalServerLogger {
	return GlobalServerZapLogger{logger: zap.L()}
}

func (l GlobalServerZapLogger) Debug(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Debug(msg, zapField...)
}

func (l GlobalServerZapLogger) Info(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Info(msg, zapField...)
}

func (l GlobalServerZapLogger) Warn(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Warn(msg, zapField...)
}

func (l GlobalServerZapLogger) Error(msg string, fields ...ILogField) {
	zapField := convertZapField(fields...)
	l.logger.Error(msg, zapField...)
}

func (l GlobalServerZapLogger) Fatal(msg string, fields ...ILogField) {
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
