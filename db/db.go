package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBConnection struct {
	db *sqlx.DB
}

func NewDBConnection(dbURL string) *DBConnection {
	conn := &DBConnection{}
	err := conn.Init(dbURL)
	if err != nil {
		panic(err)
	}
	return conn
}

func (conn *DBConnection) Init(dbURL string) error {
	db, err := sqlx.Connect("sqlite3", dbURL)
	if err != nil {
		return err
	}
	conn.db = db
	_, err = conn.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}
	return nil
}
