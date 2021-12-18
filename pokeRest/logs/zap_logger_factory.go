package logs

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildZapLogger(writer io.Writer) *zap.Logger {
	encoderConfig := buildZapCommonEncoderConfig()
	accessLogCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // フォーマット指定
		zapcore.AddSync(writer),
		zapcore.WarnLevel, // 出力対象
	)
	return zap.New(accessLogCore)
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
		EncodeTime:     zapcore.TimeEncoderOfLayout("2021/12/19 15:04:05.000(UTCZ0700)"), // 時刻フォーマット
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 出力元ソース出力形式
	}
}
