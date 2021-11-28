package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Symthy/PokeRest/pokeRest/adapters/di"
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/handler"
	"github.com/Symthy/PokeRest/pokeRest/cmd/migration"
	"github.com/Symthy/PokeRest/pokeRest/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	var isMigration = flag.Bool("migrate", false, "run migration")
	var isDropTables = flag.Bool("alldrop", false, "run drop all table")
	flag.Parse()

	conf, err := config.LoadConfigYaml()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if *isMigration {
		log.Print("Start Migration")
		migration.RunAutoMigration(conf.DbConfig)
		log.Print("End Migration")
		return
	} else if *isDropTables {
		log.Print("Start Drop All Tables")
		migration.RunDropTables(conf.DbConfig)
		log.Print("End Drop All Tables")
		return
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	e.Use(middleware.CSRF())

	pokeCon := di.InitPokemonController(conf.DbConfig)
	userCon := di.InitUserController(conf.DbConfig)
	handler := handler.NewPokeRestHandler(pokeCon, userCon)
	server.RegisterHandlers(e, handler)

	// JWT auth
	r := e.Group("/users")
	var jwtConfig = middleware.JWTConfig{
		Claims:     &config.JwtCustomClaims{},
		SigningKey: conf.AuthConfig.SecretKey,
	}
	r.Use(middleware.JWTWithConfig(jwtConfig))

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
