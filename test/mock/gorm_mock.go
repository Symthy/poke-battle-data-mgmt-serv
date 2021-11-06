package mock

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqldb, mock, _ := sqlmock.New()

	gdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	return gdb, mock, nil
}
