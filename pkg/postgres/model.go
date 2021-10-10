package postgres

import (
	"database/sql"
	"fmt"
)

type connectionDB struct {
	Host     sql.NullString
	Port     sql.NullInt64
	Username sql.NullString
	Password sql.NullString
	Database sql.NullString
	Encrypt  sql.NullString
}

func (item *connectionDB) String() *string {
	conn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		item.Host.String,
		item.Port.Int64,
		item.Username.String,
		item.Password.String,
		item.Database.String,
		item.Encrypt.String,
	)

	return &conn
}
