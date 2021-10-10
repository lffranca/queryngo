package format

import "database/sql"

type DB interface {
	DB() *sql.DB
}
