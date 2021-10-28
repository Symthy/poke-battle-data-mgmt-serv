package main

import (
	"log"
	"os"
	"time"

	"github.com/Symthy/PokeRest/pokeRest/adapters/orm/gormio/schema"
	"github.com/Symthy/PokeRest/pokeRest/config"
	"github.com/buildkite/interpolate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

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
		DSN:                  resolveDsn(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&schema.Ability{},
		&schema.BattleOpponentParty{},
		&schema.BattleRecord{},
		&schema.Form{},
		&schema.HeldItem{},
		&schema.Move{},
		&schema.Party{},
		&schema.PartyResult{},
		&schema.Pokemon{},
		&schema.Tag{},
		&schema.TrainedPokemon{},
	)
}

func resolveDsn() string {
	conf, _ := config.LoadConfigYaml()

	var confArray = []string{}
	if conf != nil {
		confArray = []string{
			"Host=" + conf.DbConfig.Host,
			"User=" + conf.DbConfig.User,
			"Password=" + conf.DbConfig.Password,
			"DbName=" + conf.DbConfig.DbName,
			"Port=" + conf.DbConfig.Port,
		}
	}
	env := interpolate.NewSliceEnv(confArray)

	dsn, err := interpolate.Interpolate(env,
		"host=${Host:-localhost} user=${User:-postgres} password=${Password:-postgres} dbname=${DbName:-pokedb} port=${Port:-5432} sslmode=disable TimeZone=Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	return dsn
}
