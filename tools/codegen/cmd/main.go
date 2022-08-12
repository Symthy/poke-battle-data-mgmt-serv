package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Symthy/PokeRest/tools/codegen/internal/parser"
)

const (
	schemaPath = "pokeRest/adapters/orm/gormio/schema/"
	entityPath = "pokeRest/domain/entity"
)

var (
	_, thisPath, _, _ = runtime.Caller(0)
)

func main() {
	rootPath := filepath.Join(filepath.Dir(thisPath), "../../..")
	err := parser.Parse(filepath.Join(rootPath, schemaPath, "user.go"))
	if err != nil {
		fmt.Print(err)
	}

	// err := builder.BuildItemFactory(homePath)
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
