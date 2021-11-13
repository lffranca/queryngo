package postgres

import "database/sql"

func newDB(conn *string) (*sql.DB, error) {
	db, err := sql.Open("postgres", *conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
