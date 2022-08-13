package main

import (
	"path/filepath"
	"runtime"

	"github.com/Symthy/PokeRest/tools/codegen/internal/parser"
)

const (
	schemaPath    = "pokeRest/adapters/orm/gormio/schema/"
	entityPath    = "pokeRest/domain/entity"
	testInputPath = "tools/codegen/test/input"
)

var (
	_, thisPath, _, _ = runtime.Caller(0)
)

func main() {
	rootPath := filepath.Join(filepath.Dir(thisPath), "../../..")
	parser.Parse(filepath.Join(rootPath, testInputPath, "move.go"))
	// err := parser.Parse(filepath.Join(rootPath, entityPath, "users", "users.go"))

	// if err != nil {
	// 	fmt.Print(fields)
	// }

	// err := builder.BuildItemFactory(homePath)
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
