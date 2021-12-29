package logs

import (
	"testing"

	"github.com/Symthy/PokeRest/pokeRest/config"
	"github.com/Symthy/PokeRest/pokeRest/logs"
	"github.com/Symthy/PokeRest/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoggerFactoriesTestSuite struct {
	suite.Suite
	logsConf config.LogsConfig
}

// Before
func (suite *LoggerFactoriesTestSuite) SetupTest() {
	suite.logsConf = config.LogsConfig{
		DirPath: "test/path/",
		ServerLogConfig: config.ServerLogConfig{
			Filename:         "server.log.test",
			MaxFileSizeMB:    1,
			MaxBackupNum:     2,
			MaxRetentionDays: 3,
		},
		AccessLogConfig: config.AccessLogConfig{
			Filename:         "access.log.test",
			MaxFileSizeMB:    4,
			MaxBackupNum:     5,
			MaxRetentionDays: 6,
		},
		DbLogConfig: config.DbLogConfig{
			Filename:         "db.log.test",
			MaxFileSizeMB:    7,
			MaxBackupNum:     8,
			MaxRetentionDays: 9,
		},
	}
}

func TestLoggerFactoriesTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerFactoriesTestSuite))
}

func (suite LoggerFactoriesTestSuite) TestLoggerFactories() {
	suite.Run("new logger factories", func() {
		absPath := test.GetAbsolutePath()
		configYamlModelPath := absPath + "conf/config.yml.model"
		config, err := config.LoadConfig(configYamlModelPath)
		if err != nil {
			assert.Fail(suite.T(), "load failure config.yml.model.", err.Error())
		}
		logsConf := config.LogsConfig
		serverLoggerFactory, accessLoggerFactory, dbLoggerFactory := logs.NewLoggerFactories(logsConf)
		// Todo: check log writing
		assert.NotNil(suite.T(), serverLoggerFactory)
		assert.NotNil(suite.T(), accessLoggerFactory)
		assert.NotNil(suite.T(), dbLoggerFactory)
	})

	suite.Run("new server logger", func() {
		factory := logs.NewServerLoggerFactory(suite.logsConf)
		rotateLogger := factory.BuildRotateServerLogger()
		assert.Equal(suite.T(), "test/path/server.log.test", rotateLogger.Filename)
		assert.Equal(suite.T(), 1, rotateLogger.MaxSize)
		assert.Equal(suite.T(), 2, rotateLogger.MaxBackups)
		assert.Equal(suite.T(), 3, rotateLogger.MaxAge)
	})

	suite.Run("new access logger", func() {
		factory := logs.NewAccessLoggerFactory(suite.logsConf)
		rotateLogger := factory.BuildRotateAccessLogger()
		assert.Equal(suite.T(), "test/path/access.log.test", rotateLogger.Filename)
		assert.Equal(suite.T(), 4, rotateLogger.MaxSize)
		assert.Equal(suite.T(), 5, rotateLogger.MaxBackups)
		assert.Equal(suite.T(), 6, rotateLogger.MaxAge)
	})

	suite.Run("new db logger", func() {
		factory := logs.NewDbLoggerFactory(suite.logsConf)
		rotateLogger := factory.BuildRotateDbLogger()
		assert.Equal(suite.T(), "test/path/db.log.test", rotateLogger.Filename)
		assert.Equal(suite.T(), 7, rotateLogger.MaxSize)
		assert.Equal(suite.T(), 8, rotateLogger.MaxBackups)
		assert.Equal(suite.T(), 9, rotateLogger.MaxAge)
	})
}
