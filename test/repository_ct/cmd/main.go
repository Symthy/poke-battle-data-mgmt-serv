//go:build test

package main

import (
	"flag"

	"github.com/Symthy/PokeRest/test/repository_ct/migration"
)

func main() {
	var isMigration = flag.Bool("migrate", false, "run migration")
	var isDropTables = flag.Bool("alldrop", false, "run drop all table")
	flag.Parse()

	if *isMigration {
		migration.CreateTable()
	} else if *isDropTables {
		migration.AllTableDrop()
	}
}
