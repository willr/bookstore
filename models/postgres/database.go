package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {

	// db, err := NewDB("postgres://dbuser:db3057md@localhost/bookstore?sslmode=disable")
	db, err := sql.Open("postres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
