package repositoryct

import (
	"database/sql"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	db, err = sql.Open("postgres", "dbname=pokedb")
	if err != nil {
		log.Fatalf("Could not connect to db: %s", err)
	}

	setupFixtures()

	os.Exit(m.Run())
}

func setupFixtures() {
	_, pwd, _, _ := runtime.Caller(0)

	dir := filepath.Join(path.Dir(pwd), "testdata")
	fix, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(dir),
	)
	if err != nil {
		log.Fatalf("setup fixtures error: %v", err)
	}
	fixtures = fix
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatalf("test data load error: %v", err)
	}
}

func TestPokemonRepository(t *testing.T) {

}
