package logs

import (
	"io"

	"github.com/Symthy/PokeRest/internal/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLoggerFactories(conf config.LogsConfig) (IServerLoggerFactory, IAccessLoggerFactory, IDbLoggerFactory) {
	serverLoggerFactory := NewServerLoggerFactory(conf)
	accessLoggerFactory := NewAccessLoggerFactory(conf)
	dbLoggerFactory := NewDbLoggerFactory(conf)
	return serverLoggerFactory, accessLoggerFactory, dbLoggerFactory
}

type ServerLoggerFactory struct {
	dirPath  string
	logLevel Level
	conf     config.ServerLogConfig
}

// public for testing
func NewServerLoggerFactory(conf config.LogsConfig) ServerLoggerFactory {
	return ServerLoggerFactory{
		dirPath: conf.DirPath,
		conf:    conf.ServerLogConfig,
	}
}

func (f ServerLoggerFactory) ResolveLogLevel() Level {
	return f.logLevel
}

func (f ServerLoggerFactory) BuildBaseServerLogger() io.Writer {
	rotateLogger := f.BuildRotateServerLogger()
	return rotateLogger
}

// public for testing
func (f ServerLoggerFactory) BuildRotateServerLogger() *lumberjack.Logger {
	var fileName = f.conf.Filename
	if fileName == "" {
		fileName = "server.log"
	}
	rotateLogger := buildRotateLogger(
		resolveDirPath(f.dirPath)+fileName, // default: /var/log/internal/server.log
		f.conf.MaxFileSizeMB,               // default: 200MB
		f.conf.MaxBackupNum,                // default: 5 file
		f.conf.MaxRetentionDays,            // default: 30 days
		true,                               // disabled by default
	)
	return rotateLogger
}

type AccessLoggerFactory struct {
	dirPath  string
	logLevel Level
	conf     config.AccessLogConfig
}

// public for testing
func NewAccessLoggerFactory(conf config.LogsConfig) AccessLoggerFactory {
	return AccessLoggerFactory{
		dirPath: conf.DirPath,
		conf:    conf.AccessLogConfig,
	}
}

func (f AccessLoggerFactory) ResolveLogLevel() Level {
	return f.logLevel
}

func (f AccessLoggerFactory) BuildBaseAccessLogger() io.Writer {
	rotateLogger := f.BuildRotateAccessLogger()
	return rotateLogger
}

// public for testing
func (f AccessLoggerFactory) BuildRotateAccessLogger() *lumberjack.Logger {
	var fileName = f.conf.Filename
	if fileName == "" {
		fileName = ""
	}
	rotateLogger := buildRotateLogger(
		resolveDirPath(f.dirPath)+fileName,
		f.conf.MaxFileSizeMB,    // default: 200MB
		f.conf.MaxBackupNum,     // default: 5 file
		f.conf.MaxRetentionDays, // default: 30 days
		true,                    // disabled by default
	)
	return rotateLogger
}

type DbLoggerFactory struct {
	dirPath  string
	logLevel Level
	conf     config.DbLogConfig
}

// public for testing
func NewDbLoggerFactory(conf config.LogsConfig) DbLoggerFactory {
	return DbLoggerFactory{
		dirPath: conf.DirPath,
		conf:    conf.DbLogConfig,
	}
}

func (f DbLoggerFactory) ResolveLogLevel() Level {
	return f.logLevel
}

func (f DbLoggerFactory) BuildBaseDbLogger() io.Writer {
	rotateLogger := f.BuildRotateDbLogger()
	return rotateLogger
}

// public for testing
func (f DbLoggerFactory) BuildRotateDbLogger() *lumberjack.Logger {
	var fileName = f.conf.Filename
	if fileName == "" {
		fileName = "db.log" // default
	}
	rotateLogger := buildRotateLogger(
		resolveDirPath(f.dirPath)+fileName,
		f.conf.MaxFileSizeMB,    // default: 200MB
		f.conf.MaxBackupNum,     // default: 5 file
		f.conf.MaxRetentionDays, // default: 30 days
		true,                    // disabled by default
	)
	return rotateLogger
}

func resolveDirPath(dirPath string) string {
	if dirPath == "" {
		dirPath = "/var/log/internal/" // default
	}
	return dirPath
}

const (
	defaultMaxFileSizeMB    int = 200
	defaultMaxBackupFileNum int = 5
	defaultMaxRetentionDays int = 30
)

func buildRotateLogger(filePath string, maxFileSizeMB int, maxBackupFileNum int, maxRetentionDays int, isCompress bool) *lumberjack.Logger {
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
