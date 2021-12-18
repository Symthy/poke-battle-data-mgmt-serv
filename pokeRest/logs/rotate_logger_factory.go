package logs

import (
	"io"

	"github.com/Symthy/PokeRest/pokeRest/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLoggerFactories(conf config.LogsConfig) (IServerLoggerFactory, IAccessLoggerFactory, IDbLoggerFactory) {
	serverLoggerFactory := NewRotateServerLoggerFactory(conf)
	accessLoggerFactory := NewRotateAccessLoggerFactory(conf)
	dbLoggerFactory := NewRotateDbLoggerFactory(conf)
	return serverLoggerFactory, accessLoggerFactory, dbLoggerFactory
}

type RotateServerLoggerFactory struct {
	dirPath string
	conf    config.ErrorLogConfig
}

func NewRotateServerLoggerFactory(conf config.LogsConfig) IServerLoggerFactory {
	return RotateServerLoggerFactory{
		dirPath: conf.DirPath,
		conf:    conf.ErrorLogConfig,
	}
}

func (f RotateServerLoggerFactory) BuildRotateServerLogger() io.Writer {
	var fileName = f.conf.Filename
	if fileName == "" {
		fileName = "server.log"
	}
	rotateLogger := buildRotateLogger(
		f.dirPath+fileName,      // default: /var/log/pokeRest/error.log
		f.conf.MaxFileSizeMB,    // default: 200MB
		f.conf.MaxBackupNum,     // default: 5 file
		f.conf.MaxRetentionDays, // default: 30 days
		true,                    // disabled by default
	)
	return rotateLogger
}

type RotateAccessLoggerFactory struct {
	dirPath string
	conf    config.AccessLogConfig
}

func NewRotateAccessLoggerFactory(conf config.LogsConfig) IAccessLoggerFactory {
	return RotateAccessLoggerFactory{
		dirPath: conf.DirPath,
		conf:    conf.AccessLogConfig,
	}
}

func (f RotateAccessLoggerFactory) BuildRotateAccessLogger() io.Writer {
	var fileName = f.conf.Filename
	if fileName == "" {
		fileName = ""
	}
	rotateLogger := buildRotateLogger(
		f.dirPath+fileName,
		f.conf.MaxFileSizeMB,    // default: 200MB
		f.conf.MaxBackupNum,     // default: 5 file
		f.conf.MaxRetentionDays, // default: 30 days
		true,                    // disabled by default
	)
	return rotateLogger
}

type RotateDbLoggerFactory struct {
	dirPath string
	conf    config.DbLogConfig
}

func NewRotateDbLoggerFactory(conf config.LogsConfig) IDbLoggerFactory {
	return RotateDbLoggerFactory{
		dirPath: conf.DirPath,
		conf:    conf.DbLogConfig,
	}
}

func (f RotateDbLoggerFactory) BuildRotateDbLogger() io.Writer {
	var fileName = f.conf.Filename
	if fileName == "" {
		fileName = "db.log" // default
	}
	rotateLogger := buildRotateLogger(
		f.dirPath+fileName,
		f.conf.MaxFileSizeMB,    // default: 200MB
		f.conf.MaxBackupNum,     // default: 5 file
		f.conf.MaxRetentionDays, // default: 30 days
		true,                    // disabled by default
	)
	return rotateLogger
}

func resolveDirPath(dirPath string) string {
	if dirPath == "" {
		dirPath = "/var/log/pokeRest/" // default
	}
	return dirPath
}

const (
	defaultMaxFileSizeMB    int = 200
	defaultMaxBackupFileNum int = 5
	defaultMaxRetentionDays int = 30
)

func buildRotateLogger(filePath string, maxFileSizeMB int, maxBackupFileNum int, maxRetentionDays int, isCompress bool) io.Writer {
	var fileSizeMB = defaultMaxFileSizeMB
	var backupFileNum = defaultMaxBackupFileNum
	var retentionDays = defaultMaxRetentionDays
	if maxFileSizeMB > 0 {
		fileSizeMB = maxFileSizeMB
	}
	if maxBackupFileNum > 0 {
		backupFileNum = maxBackupFileNum
	}
	if maxRetentionDays > 0 {
		retentionDays = maxRetentionDays
	}
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    fileSizeMB,
		MaxBackups: backupFileNum,
		MaxAge:     retentionDays,
		Compress:   isCompress,
	}
}
