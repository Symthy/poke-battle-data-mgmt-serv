package orm

import (
	"log"
	"os"
	"time"

	"github.com/Symthy/PokeRest/pokeRest/config"
	"github.com/buildkite/interpolate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _ IDbClient = (*GormDbClient)(nil)

type GormDbClient struct {
	db *gorm.DB
}

func NewGormDbClient(dbConfig config.DbConfig, logger logger.Interface) *GormDbClient {
	dc := &GormDbClient{}
	dc.initialize(dbConfig, logger)
	return dc
}

func NewGormDbClientForStdOut(dbConfig config.DbConfig) *GormDbClient {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	dc := &GormDbClient{}
	dc.db = openDb(dbConfig, newLogger)
	return dc
}

func NewGormDbClientForTesting(db *gorm.DB) *GormDbClient {
	return &GormDbClient{db: db}
}

func (dc *GormDbClient) initialize(dbConfig config.DbConfig, logger logger.Interface) {
	dc.db = openDb(dbConfig, logger)
}

func openDb(dbConfig config.DbConfig, logger logger.Interface) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  resolveDsn(dbConfig),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: logger})
	if err != nil {
		// Todo: error handling
		panic("failed to connect database")
	}
	return db
}

func (dc GormDbClient) Db() *gorm.DB {
	return dc.db
}

func (dc GormDbClient) Close() {
	sqldb, _ := dc.db.DB()
	sqldb.Close()
}

func (dc GormDbClient) Paginate(next uint32, pageSize uint16) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if next == 0 {
			next = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 100
		}

		return db.Offset(int(next)).Limit(int(pageSize))
	}
}

func resolveDsn(dbConfig config.DbConfig) string {
	confArray := []string{
		"Host=" + dbConfig.Host,
		"User=" + dbConfig.User,
		"Password=" + dbConfig.Password,
		"DbName=" + dbConfig.DbName,
		"Port=" + dbConfig.Port,
	}

	env := interpolate.NewSliceEnv(confArray)

	dsn, err := interpolate.Interpolate(env,
		"host=${Host:-localhost} user=${User:-postgres} password=${Password:-postgres} dbname=${DbName:-pokedb} port=${Port:-5432} sslmode=disable TimeZone=Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	return dsn
}
