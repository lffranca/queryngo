package gmigrate

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PostgresService service

func (pkg PostgresService) Up() error {
	driver, err := postgres.WithInstance(pkg.client.db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", *pkg.client.path),
		"postgres", driver)
	if err != nil {
		return err
	}

	return m.Up()
}
