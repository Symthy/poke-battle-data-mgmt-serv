package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Symthy/PokeRest/internal/adapters/di"
	"github.com/Symthy/PokeRest/internal/adapters/orm"
	"github.com/Symthy/PokeRest/internal/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/internal/adapters/rest/handler"
	"github.com/Symthy/PokeRest/internal/config"
	"github.com/Symthy/PokeRest/internal/logs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")

	conf, err := config.LoadConfigYaml()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.

	// This is how you set up a basic Echo router
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())

	// Logging settings
	serverLoggerFactory, accessLoggerFactory, dbLoggerFactory := logs.NewLoggerFactories(conf.LogsConfig)
	serverLogInitializer := logs.NewServerGlobalLoggerInitializer(serverLoggerFactory)
	accessLogInitializer := logs.NewAccessLoggerMiddlewareInitializer(accessLoggerFactory)
	// if l, ok := e.Logger.(*_labstacklog.Logger); ok {
	// 	l.SetHeader("${time_rfc3339} ${level}")
	// 	l.SetOutput(logs.BuildRotateErrorLogger(conf.LogsConfig))
	// } else {
	// 	e.Logger.Fatalf("failure logging settings. start abort.")
	// }
	// e.Use(middleware.Logger()) // Log all requests
	serverLogInitializer.AcceptLogger()
	accessLogInitializer.AcceptLogger(e)

	// DB
	dbClientInitializer := orm.NewDbClientInitializer(conf.DbConfig, dbLoggerFactory)
	dbClient := dbClientInitializer.InitializeDbClient()

	// JWT auth
	r := e.Group("/users")
	var jwtConfig = middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: conf.AuthConfig.SecretKey,
	}
	r.Use(middleware.JWTWithConfig(jwtConfig))

	// controller initialization
	pokeRestHandler := initPokeRestHandler(dbClient)
	server.RegisterHandlers(e, pokeRestHandler)

	// custom http error handler
	e.HTTPErrorHandler = handler.PokeRestHttpErrorHandler

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}

func initPokeRestHandler(dbClient orm.IDbClient) *handler.PokeRestHandler {
	return handler.NewPokeRestHandler(
		di.InitPokemonController(dbClient),
		di.InitAbilityController(dbClient),
		di.InitMoveController(dbClient),
		di.InitItemController(dbClient),
		di.InitTypeController(),
		di.InitPartyTagController(dbClient),
		di.InitPartyController(dbClient),
		di.InitBattleRecordController(dbClient),
		di.InitTrainedPokemonController(dbClient),
		di.InitUserController(dbClient))
}
