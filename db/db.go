package db

import (
	"fmt"
	"log"
	"server/types"

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

func (conn *DBConnection) SeedData() {
	groups := []types.Group{
		{
			ID:      "la_verte",
			Name:    "La Verte",
			Picture: "https://ik.imagekit.io/lddm/la_verte2.jpg?updatedAt=1756204812699",
		},
		{
			ID:      "atadi_fionfion",
			Name:    "Atadi Fionfion",
			Picture: "https://ik.imagekit.io/lddm/atadi_fionfion.jpg?updatedAt=1756204709945",
		},
		{
			ID:      "koklo_blend",
			Name:    "Koklo Blend",
			Picture: "https://ik.imagekit.io/lddm/koklo_blend.jpg?updatedAt=1756204709840",
		},
		{
			ID:      "la_totale",
			Name:    "La Totale",
			Picture: "https://ik.imagekit.io/lddm/la_totale2.jpg?updatedAt=1756204709814",
		},
		{
			ID:      "l_exotique",
			Name:    "L'Exotique",
			Picture: "https://ik.imagekit.io/lddm/exotique.jpg?updatedAt=1756204709634",
		},
		{
			ID:      "l_exquise",
			Name:    "L'Exquise",
			Picture: "https://ik.imagekit.io/lddm/exquise2.jpg?updatedAt=1756204709558",
		},
		{
			ID:      "persilmix",
			Name:    "Persilmix",
			Picture: "https://ik.imagekit.io/lddm/persilmix.jpg?updatedAt=1756204709441",
		},
		{
			ID:      "gbolan_blend",
			Name:    "Gbolan Blend",
			Picture: "https://ik.imagekit.io/lddm/gbolan_blend.jpg?updatedAt=1756204709171",
		},
		{
			ID:      "ail",
      Name:    "Ail",
      Picture: "https://ik.imagekit.io/lddm/ail.jpg?updatedAt=1756204709015",
		},
		{
			ID:      "atadi_delices",
			Name:    "Atadi Delices",
			Picture: "https://ik.imagekit.io/lddm/atadi_delices.jpg?updatedAt=1756204707833",
		},
	}
	const groupSeedQuery = `
    insert into groups (
      id, name, picture
    )
    values (
      :id, :name, :picture
    )
    on conflict(name) do nothing
  `
	log.Println("[INFO]: Seeding Groups")
	_, err := conn.db.NamedExec(
		groupSeedQuery, groups,
	)
	if err != nil {
		log.Panicf("Seeding failed for groups: %s\n", err.Error())
	}
	log.Println("[INFO]: Done!")
}
