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
	var isMigration = flag.Bool("migrate", false, "run migration")
	var isDropTables = flag.Bool("alldrop", false, "run drop all table")
	flag.Parse()

	if *isMigration {
		log.Print("Start Migration")
		migration.RunAutoMigration()
		log.Print("End Migration")
		return
	} else if *isDropTables {
		log.Print("Start Drop All Tables")
		migration.RunDropTables()
		log.Print("End Drop All Tables")
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
