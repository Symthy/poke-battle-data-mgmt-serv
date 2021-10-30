package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/autogen/server"
	"github.com/Symthy/PokeRest/pokeRest/adapters/rest/handler"
	"github.com/Symthy/PokeRest/pokeRest/cmd/migration"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	var is_migration = flag.Bool("migration", false, "run migration")
	flag.Parse()

	if *is_migration {
		log.Print("Start Migration")
		migration.RunMigration()
		log.Print("End Migration")
		return
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())

	handler := handler.NewPokeRestHandler()
	server.RegisterHandlers(e, handler)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
