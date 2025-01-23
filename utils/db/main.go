package db

import "database/sql"

type DB struct {
	db *sql.DB
}

func newDB() *DB {
	return &DB{}
}
