package models

import (
	"bookstore/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (Datastore, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func SetupDB(conf *config.RuntimeConfig) Datastore {
	db, err := NewDB(config.BuildConnectionString(conf))
	// db, err := NewDB("postgres://dbuser:db3057md@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	return db
}
