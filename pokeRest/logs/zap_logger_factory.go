package logs

import (
	"io"

	"github.com/Symthy/PokeRest/pokeRest/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// public for testing
func BuildZapLogger(writer io.Writer, level common.Level, options ...zap.Option) *zap.Logger {
	encoderConfig := buildZapCommonEncoderConfig()
	accessLogCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // フォーマット指定
		zapcore.AddSync(writer),
		convertLevel(level), // 出力エラーレベル
	)
	return zap.New(accessLogCore, options...)
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

func convertLevel(level common.Level) zapcore.Level {
	switch level {
	case common.Fatal:
		return zapcore.FatalLevel
	case common.Error:
		return zapcore.ErrorLevel
	case common.Warn:
		return zapcore.WarnLevel
	case common.Info:
		return zapcore.InfoLevel
	case common.Debug:
		return zapcore.DebugLevel
	default:
		return zapcore.WarnLevel
	}
}
