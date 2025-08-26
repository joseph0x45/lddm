package db

import (
	"fmt"
	"log"

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

func (conn *DBConnection) Migrate() {
	const schema = `
    PRAGMA foreign_keys = ON;

    create table if not exists groups (
      id text not null primary key,
      name text not null unique,
      picture text not null
    );

    create table if not exists products (
      id text not null primary key,
      group_id text not null references groups(id) on delete cascade,
      name text not null,
      variant text not null,
      picture text not null,
      in_stock integer not null default 0,
      base_price integer not null
    );

    create table if not exists product_bundle_prices (
      id text not null primary key,
      product_id text not null references products(id) on delete cascade,
      quantity integer not null,
      bundle_price integer not null,
      UNIQUE(product_id, quantity)
    );
  `
  log.Println("[INFO]: Running migrations")
  conn.db.MustExec(schema)
  log.Println("[INFO]: Done!")
}
