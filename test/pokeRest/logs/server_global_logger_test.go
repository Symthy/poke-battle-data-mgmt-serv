package logs

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/Symthy/PokeRest/pokeRest/common"
	"github.com/Symthy/PokeRest/pokeRest/logs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ServerGlobalLoggerTestSuite struct {
	suite.Suite
	buf    *bytes.Buffer
	logger logs.IServerGlobalLogger
}

type constantClock time.Time

func (c constantClock) Now() time.Time {
	return time.Time(c)
}
func (c constantClock) NewTicker(duration time.Duration) *time.Ticker {
	// return time.NewTicker(duration)
	return &time.Ticker{}
}

// Before
func (suite *ServerGlobalLoggerTestSuite) SetupTest() {
	date := time.Date(2021, 12, 29, 17, 05, 22, 1, time.UTC)
	clock := constantClock(date)
	suite.buf = new(bytes.Buffer)
	logger := logs.BuildZapLogger(suite.buf, common.Debug, zap.WithClock(clock))
	logs.InitGlobalServerLogger(logger)
	suite.logger = logs.GetGlobalServerLogger()
}

func TestServerGlobalLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerGlobalLoggerTestSuite))
}

func (suite ServerGlobalLoggerTestSuite) TestServerGlobalLogger() {
	suite.Run("logger level methods output", func() {
		logger := suite.logger
		tests := []struct {
			method          func(string, ...logs.ILogField)
			expectedMessage string
		}{
			{logger.Debug, "2021/12/29 17:05:22.000(UTCZ)	DEBUG	test message"},
			{logger.Info, "2021/12/29 17:05:22.000(UTCZ)	INFO	test message"},
			{logger.Warn, "2021/12/29 17:05:22.000(UTCZ)	WARN	test message"},
			{logger.Error, "2021/12/29 17:05:22.000(UTCZ)	ERROR	test message"},
			// time zone could not check
			// Fatal is no check. because process abort.
		}

		for _, tt := range tests {
			tt.method("test message")
			assert.Equal(suite.T(), tt.expectedMessage, strings.TrimRight(suite.buf.String(), "\n"))
			suite.buf.Reset()
		}
	})
}
