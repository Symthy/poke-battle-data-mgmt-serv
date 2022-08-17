package logs

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// public for testing
func BuildZapLogger(writer io.Writer, level Level, options ...zap.Option) *zap.Logger {
	encoderConfig := buildZapCommonEncoderConfig()
	logCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(writer),
		convertLevel(level), // 出力エラーレベル
	)
	return zap.New(logCore, options...)
}

func buildZapCommonEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,                                      // 大文字表示
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.000(UTCZ0700)"), // 時刻フォーマット
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 出力元ソース出力形式
	}
}

func convertLevel(level Level) zapcore.Level {
	switch level {
	case Fatal:
		return zapcore.FatalLevel
	case Error:
		return zapcore.ErrorLevel
	case Warn:
		return zapcore.WarnLevel
	case Info:
		return zapcore.InfoLevel
	case Debug:
		return zapcore.DebugLevel
	default:
		return zapcore.WarnLevel
	}
}
