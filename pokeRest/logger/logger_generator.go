package logger

import (
	"io"

	"github.com/Symthy/PokeRest/pokeRest/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GenerateRotateErrorLogger(logsConf config.LogsConfig) io.Writer {
	errLogConf := logsConf.ErrorLogConfig
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logsConf.DirPath + errLogConf.Filename, // default: /var/log/pokeRest/error.log
		MaxSize:    errLogConf.MaxFileSizeMB,               // default: 200MB
		MaxBackups: errLogConf.MaxBackupNum,                // default: 5 file
		MaxAge:     errLogConf.MaxRetentionDays,            // default: 30 days
		Compress:   true,                                   // disabled by default
	}
	return lumberjackLogger
}

func GenerateAccessErrorLogger(logsConf config.LogsConfig) io.Writer {
	accLogConf := logsConf.AccessLogConfig
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logsConf.DirPath + accLogConf.Filename, // default: /var/log/pokeRest/error.log
		MaxSize:    accLogConf.MaxFileSizeMB,               // default: 200MB
		MaxBackups: accLogConf.MaxBackupNum,                // default: 5 file
		MaxAge:     accLogConf.MaxRetentionDays,            // default: 30 days
		Compress:   true,                                   // disabled by default
	}
	return lumberjackLogger
}
