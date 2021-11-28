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

type GormDbClient struct {
	db *gorm.DB
}

func NewGormDbClient(dbConfig config.DbConfig) *GormDbClient {
	dc := &GormDbClient{}
	dc.initialize(dbConfig)
	return dc
}

func NewGormDbClientForTesting(db *gorm.DB) *GormDbClient {
	return &GormDbClient{db: db}
}

func (dc *GormDbClient) initialize(dbConfig config.DbConfig) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	//	db, err := gorm.Open(postgres.Open(resolveDsn()), &gorm.Config{})
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  resolveDsn(dbConfig),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	dc.db = db
}

func (dc GormDbClient) Db() *gorm.DB {
	return dc.db
}

func (dc GormDbClient) Close() {
	sqldb, _ := dc.db.DB()
	sqldb.Close()
}

func (dc GormDbClient) Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func resolveDsn(dbConfig config.DbConfig) string {
	var confArray = []string{}
	confArray = []string{
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
