package models

import (
	"bookstore/config"
	"database/sql"
	"bookstore/models/postgres"
	"bookstore/models/sqlite"
	"log"
)

type DB struct {
	*sql.DB
}

func NewDB(driverName string, dataSourceName string) (Datastore, error) {

	handled := false
	var db *sql.DB
	var err error
	switch driverName {
	case "postgres":
		db, err = postgres.NewDB(dataSourceName)
		if err == nil {
			handled = true
		}
	case "sqlite3":
		db, err = sqlite.NewDB(dataSourceName)
		if err == nil {
			handled = true
		}
	}
	if handled {
		return &DB{db}, nil
	} else {
		return nil, err
	}
}

func SetupDB(conf *config.RuntimeConfig) Datastore {
	db, err := NewDB(conf.DatabaseDriver, config.BuildConnectionString(conf))
	if err != nil {
		log.Panic(err)
	}

	return db
}
